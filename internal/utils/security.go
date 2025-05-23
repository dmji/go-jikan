package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	securityTagKey = "security"
)

type securityTag struct {
	Option  bool
	Scheme  bool
	Name    string
	Type    string
	SubType string
	Env     string
}

func PopulateSecurity(ctx context.Context, req *http.Request, securitySource func(context.Context) (interface{}, error)) error {
	if securitySource == nil {
		return nil
	}

	security, err := securitySource(ctx)
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	queryParams := make(map[string]string)

	securityValType := trueReflectValue(reflect.ValueOf(security))
	securityStructType := securityValType.Type()

	if isNil(securityStructType, securityValType) {
		return nil
	}

	if securityStructType.Kind() == reflect.Ptr {
		securityStructType = securityStructType.Elem()
		securityValType = securityValType.Elem()
	}

	for i := 0; i < securityStructType.NumField(); i++ {
		fieldType := securityStructType.Field(i)
		valType := securityValType.Field(i)

		kind := valType.Kind()

		if isNil(fieldType.Type, valType) {
			continue
		}

		if fieldType.Type.Kind() == reflect.Pointer {
			kind = valType.Elem().Kind()
		}

		secTag := parseSecurityTag(fieldType)
		if secTag != nil {
			if secTag.Option {
				handleSecurityOption(headers, queryParams, valType.Interface())
			} else if secTag.Scheme {
				// Special case for basic auth which could be a flattened struct
				if secTag.SubType == "basic" && kind != reflect.Struct {
					parseSecurityScheme(headers, queryParams, secTag, security)
				} else {
					parseSecurityScheme(headers, queryParams, secTag, valType.Interface())
				}
			}
		}
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	query := req.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	return nil
}

func PopulateSecurityFromEnv(security interface{}) bool {
	securityValType := trueReflectValue(reflect.ValueOf(security))
	securityStructType := securityValType.Type()

	if securityStructType.Kind() == reflect.Ptr {
		securityStructType = securityStructType.Elem()
		securityValType = securityValType.Elem()
	}

	return populateStructFromEnv(securityStructType, securityValType)
}

func populateStructFromEnv(structType reflect.Type, structVal reflect.Value) bool {
	fieldPopulated := false
	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		secTag := parseSecurityTag(fieldType)
		if fieldType.Type.Kind() == reflect.Struct ||
			(fieldType.Type.Kind() == reflect.Ptr && fieldType.Type.Elem().Kind() == reflect.Struct) {
			if fieldType.Type.Kind() == reflect.Ptr {
				if hasEnvVarsForSecurityStruct(fieldType.Type.Elem(), secTag != nil && secTag.Option) {
					fieldVal := reflect.New(fieldType.Type.Elem())
					if populateStructFromEnv(fieldType.Type.Elem(), fieldVal.Elem()) {
						structVal.Field(i).Set(fieldVal)
						fieldPopulated = true
					}
				}
			} else {
				fieldVal := structVal.Field(i)
				if populateStructFromEnv(fieldType.Type.Elem(), fieldVal.Elem()) {
					fieldPopulated = true
				}
			}
			continue
		}

		if secTag == nil || secTag.Env == "" {
			continue
		}

		envValue := getCaseInsensitiveEnvVar(secTag.Env)
		if envValue == "" {
			continue
		}

		if fieldType.Type.Kind() == reflect.Ptr {
			fieldVal := reflect.New(fieldType.Type.Elem())
			if setFieldValue(fieldVal.Elem(), envValue) {
				structVal.Field(i).Set(fieldVal)
				fieldPopulated = true
			}
		} else {
			fieldVal := structVal.Field(i)
			if setFieldValue(fieldVal, envValue) {
				fieldPopulated = true
			}
		}
	}

	return fieldPopulated
}

func hasEnvVarsForSecurityStruct(structType reflect.Type, isSecurityOption bool) bool {
	valid := false
	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		secTag := parseSecurityTag(fieldType)
		if secTag == nil {
			continue
		}

		if secTag.Env != "" || getCaseInsensitiveEnvVar(secTag.Env) != "" {
			valid = true
		} else if isSecurityOption {
			return false
		}

		if fieldType.Type.Kind() == reflect.Struct ||
			(fieldType.Type.Kind() == reflect.Ptr && fieldType.Type.Elem().Kind() == reflect.Struct) {
			if hasEnvVarsForSecurityStruct(fieldType.Type, secTag.Option) {
				valid = true
			} else if isSecurityOption {
				return false
			}
		}
	}

	return valid
}

func handleSecurityOption(headers, queryParams map[string]string, option interface{}) {
	optionValType := trueReflectValue(reflect.ValueOf(option))
	optionStructType := optionValType.Type()

	if isNil(optionStructType, optionValType) {
		return
	}

	for i := 0; i < optionStructType.NumField(); i++ {
		fieldType := optionStructType.Field(i)
		valType := optionValType.Field(i)

		secTag := parseSecurityTag(fieldType)
		if secTag != nil && secTag.Scheme {
			parseSecurityScheme(headers, queryParams, secTag, valType.Interface())
		}
	}
}

func parseSecurityScheme(headers, queryParams map[string]string, schemeTag *securityTag, scheme interface{}) {
	schemeVal := trueReflectValue(reflect.ValueOf(scheme))
	schemeType := schemeVal.Type()

	if isNil(schemeType, schemeVal) {
		return
	}

	if schemeType.Kind() == reflect.Struct {
		if schemeTag.Type == "http" {
			switch schemeTag.SubType {
			case "basic":
				handleBasicAuthScheme(headers, schemeVal.Interface())
				return
			case "custom":
				return
			}
		}

		for i := 0; i < schemeType.NumField(); i++ {
			fieldType := schemeType.Field(i)
			valType := schemeVal.Field(i)

			if isNil(fieldType.Type, valType) {
				continue
			}

			if fieldType.Type.Kind() == reflect.Ptr {
				valType = valType.Elem()
			}

			secTag := parseSecurityTag(fieldType)
			if secTag == nil || secTag.Name == "" {
				return
			}

			parseSecuritySchemeValue(headers, queryParams, schemeTag, secTag, valType.Interface())
		}
	} else {
		parseSecuritySchemeValue(headers, queryParams, schemeTag, schemeTag, schemeVal.Interface())
	}
}

func parseSecuritySchemeValue(headers, queryParams map[string]string, schemeTag *securityTag, secTag *securityTag, val interface{}) {
	switch schemeTag.Type {
	case "apiKey":
		switch schemeTag.SubType {
		case "header":
			headers[secTag.Name] = valToString(val)
		case "query":
			queryParams[secTag.Name] = valToString(val)
		case "cookie":
			headers["Cookie"] = fmt.Sprintf("%s=%s", secTag.Name, valToString(val))
		default:
			panic("not supported")
		}
	case "openIdConnect":
		headers[secTag.Name] = prefixBearer(valToString(val))
	case "oauth2":
		if schemeTag.SubType != "client_credentials" {
			headers[secTag.Name] = prefixBearer(valToString(val))
		}
	case "http":
		switch schemeTag.SubType {
		case "bearer":
			headers[secTag.Name] = prefixBearer(valToString(val))
		case "custom":
		default:
			panic("not supported")
		}
	default:
		panic("not supported")
	}
}

func prefixBearer(authHeaderValue string) string {
	if strings.HasPrefix(strings.ToLower(authHeaderValue), "bearer ") {
		return authHeaderValue
	}

	return fmt.Sprintf("Bearer %s", authHeaderValue)
}

func handleBasicAuthScheme(headers map[string]string, scheme interface{}) {
	schemeStructType := reflect.TypeOf(scheme)
	schemeValType := reflect.ValueOf(scheme)

	var username, password string

	for i := 0; i < schemeStructType.NumField(); i++ {
		fieldType := schemeStructType.Field(i)
		valType := schemeValType.Field(i)

		if fieldType.Type.Kind() == reflect.Ptr {
			valType = valType.Elem()
		}

		secTag := parseSecurityTag(fieldType)
		if secTag == nil || secTag.Name == "" {
			continue
		}

		switch secTag.Name {
		case "username":
			username = valType.String()
		case "password":
			password = valType.String()
		}
	}

	headers["Authorization"] = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
}

func parseSecurityTag(field reflect.StructField) *securityTag {
	tag := field.Tag.Get(securityTagKey)
	if tag == "" {
		return nil
	}

	option := false
	scheme := false
	name := ""
	securityType := ""
	securitySubType := ""
	env := ""

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")
		if len(parts) < 1 || len(parts) > 2 {
			continue
		}

		switch parts[0] {
		case "name":
			name = parts[1]
		case "type":
			securityType = parts[1]
		case "subtype":
			securitySubType = parts[1]
		case "option":
			option = true
		case "scheme":
			scheme = true
		case "env":
			env = parts[1]
		}
	}

	// TODO: validate tag?

	return &securityTag{
		Option:  option,
		Scheme:  scheme,
		Name:    name,
		Type:    securityType,
		SubType: securitySubType,
		Env:     env,
	}
}

func trueReflectValue(val reflect.Value) reflect.Value {
	kind := val.Type().Kind()
	for kind == reflect.Interface || kind == reflect.Ptr {
		innerVal := val.Elem()
		if !innerVal.IsValid() {
			break
		}
		val = innerVal
		kind = val.Type().Kind()
	}
	return val
}
