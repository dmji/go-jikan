package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ericlagergren/decimal"
)

const (
	queryParamTagKey  = "queryParam"
	headerParamTagKey = "header"
	pathParamTagKey   = "pathParam"
)

var (
	paramRegex                       = regexp.MustCompile(`({.*?})`)
	SerializationMethodToContentType = map[string]string{
		"json":      "application/json",
		"form":      "application/x-www-form-urlencoded",
		"multipart": "multipart/form-data",
		"raw":       "application/octet-stream",
		"string":    "text/plain",
	}
)

func UnmarshalJsonFromResponseBody(body io.Reader, out interface{}, tag string) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := UnmarshalJSON(data, out, reflect.StructTag(tag), true, false); err != nil {
		return fmt.Errorf("error unmarshalling json response body: %w", err)
	}

	return nil
}

func ReplaceParameters(stringWithParams string, params map[string]string) string {
	if len(params) == 0 {
		return stringWithParams
	}

	return paramRegex.ReplaceAllStringFunc(stringWithParams, func(match string) string {
		match = match[1 : len(match)-1]
		return params[match]
	})
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func MatchStatusCodes(expectedCodes []string, statusCode int) bool {
	for _, codeStr := range expectedCodes {
		code, err := strconv.Atoi(codeStr)
		if err == nil {
			if code == statusCode {
				return true
			}
			continue
		}

		codeRange, err := strconv.Atoi(string(codeStr[0]))
		if err != nil {
			continue
		}

		if statusCode >= (codeRange*100) && statusCode < ((codeRange+1)*100) {
			return true
		}
	}

	return false
}

func AsSecuritySource(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

func getCaseInsensitiveEnvVar(envVar string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}

	if value := os.Getenv(strings.ToUpper(envVar)); value != "" {
		return value
	}

	return ""
}

func ValueFromEnvVar(envVar string, field interface{}) interface{} {
	value := getCaseInsensitiveEnvVar(envVar)
	if value == "" {
		return nil
	}

	t := reflect.TypeOf(field)
	if t.Kind() != reflect.Ptr {
		return nil
	}
	t = t.Elem()

	switch t.Kind() {
	case reflect.String:
		return value
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return i
		}
	case reflect.Float32, reflect.Float64:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		}
	case reflect.Bool:
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return nil
}

func parseConstTag(field reflect.StructField) *string {
	value := field.Tag.Get("const")

	if value == "" {
		return nil
	}

	return &value
}

func parseDefaultTag(field reflect.StructField) *string {
	value := field.Tag.Get("default")

	if value == "" {
		return nil
	}

	return &value
}

func parseStructTag(tagKey string, field reflect.StructField) map[string]string {
	tag := field.Tag.Get(tagKey)
	if tag == "" {
		return nil
	}

	values := map[string]string{}

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")

		switch len(parts) {
		case 1:
			// flag option
			parts = append(parts, "true")
		case 2:
			// key=value option
		default:
			// invalid option
			continue
		}

		values[parts[0]] = parts[1]
	}

	return values
}

func parseParamTag(tagKey string, field reflect.StructField, defaultStyle string, defaultExplode bool) *paramTag {
	// example `{tagKey}:"style=simple,explode=false,name=apiID"`
	// example `{tagKey}:"inline"`
	values := parseStructTag(tagKey, field)
	if values == nil {
		return nil
	}

	tag := &paramTag{
		Style:     defaultStyle,
		Explode:   defaultExplode,
		ParamName: strings.ToLower(field.Name),
	}

	for k, v := range values {
		switch k {
		case "inline":
			tag.Inline = v == "true"
		case "style":
			tag.Style = v
		case "explode":
			tag.Explode = v == "true"
		case "name":
			tag.ParamName = v
		case "serialization":
			tag.Serialization = v
		}
	}

	return tag
}

func valToString(val interface{}) string {
	switch v := val.(type) {
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case big.Int:
		return v.String()
	case decimal.Big:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func populateFromGlobals(fieldType reflect.StructField, valType reflect.Value, paramType string, globals interface{}) (reflect.StructField, reflect.Value, bool) {
	if globals == nil {
		return fieldType, valType, false
	}

	globalsStruct := reflect.TypeOf(globals)
	globalsStructVal := reflect.ValueOf(globals)

	globalsField, found := globalsStruct.FieldByName(fieldType.Name)
	if !found {
		return fieldType, valType, false
	}

	if fieldType.Type.Kind() != reflect.Ptr || !valType.IsNil() {
		return fieldType, valType, true
	}

	globalsVal := globalsStructVal.FieldByName(fieldType.Name)

	if !globalsVal.IsValid() {
		return fieldType, valType, false
	}

	switch paramType {
	case queryParamTagKey:
		qpTag := parseQueryParamTag(globalsField)
		if qpTag == nil {
			return fieldType, valType, false
		}
	default:
		tag := parseParamTag(paramType, fieldType, "simple", false)
		if tag == nil {
			return fieldType, valType, false
		}
	}

	return globalsField, globalsVal, true
}

func isNil(typ reflect.Type, val reflect.Value) bool {
	// `reflect.TypeOf(nil) == nil` so calling typ.Kind() will cause a nil pointer
	// dereference panic. Catch it and return early.
	// https://github.com/golang/go/issues/51649
	// https://github.com/golang/go/issues/54208
	if typ == nil {
		return true
	}

	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Map || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Interface {
		return val.IsNil()
	}

	return false
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return v.IsZero()
	default:
		return false
	}
}

func setFieldValue(fieldVal reflect.Value, value string) bool {
	if !fieldVal.CanSet() {
		return false
	}

	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(value)
		return true
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			fieldVal.SetBool(boolValue)
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			fieldVal.SetInt(intValue)
			return true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintValue, err := strconv.ParseUint(value, 10, 64); err == nil {
			fieldVal.SetUint(uintValue)
			return true
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			fieldVal.SetFloat(floatValue)
			return true
		}
	default:
		return false
	}

	return false
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ConsumeRawBody(res *http.Response) ([]byte, error) {
	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	return rawBody, nil
}
