package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/ericlagergren/decimal"

	"github.com/dmji/go-jikan/types"
)

func PopulateQueryParams(_ context.Context, req *http.Request, queryParams interface{}, globals interface{}) error {
	// Query parameters may already be present from overriding URL
	if req.URL.RawQuery != "" {
		return nil
	}

	values := url.Values{}

	globalsAlreadyPopulated, err := populateQueryParams(queryParams, globals, values, []string{})
	if err != nil {
		return err
	}

	if globals != nil {
		_, err = populateQueryParams(globals, nil, values, globalsAlreadyPopulated)
		if err != nil {
			return err
		}
	}

	req.URL.RawQuery = values.Encode()

	return nil
}

func populateQueryParams(queryParams interface{}, globals interface{}, values url.Values, skipFields []string) ([]string, error) {
	queryParamsStructType, queryParamsValType := dereferencePointers(reflect.TypeOf(queryParams), reflect.ValueOf(queryParams))

	globalsAlreadyPopulated := []string{}

	for i := 0; i < queryParamsStructType.NumField(); i++ {
		fieldType := queryParamsStructType.Field(i)
		valType := queryParamsValType.Field(i)

		if contains(skipFields, fieldType.Name) {
			continue
		}

		requestTag := getRequestTag(fieldType)
		if requestTag != nil {
			continue
		}

		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			continue
		}

		constValue := parseConstTag(fieldType)
		if constValue != nil {
			values.Add(qpTag.ParamName, *constValue)
			continue
		}

		defaultValue := parseDefaultTag(fieldType)

		if globals != nil {
			var globalFound bool
			fieldType, valType, globalFound = populateFromGlobals(fieldType, valType, queryParamTagKey, globals)
			if globalFound {
				globalsAlreadyPopulated = append(globalsAlreadyPopulated, fieldType.Name)
			}
		}

		if qpTag.Serialization != "" {
			vals, err := populateSerializedParams(qpTag, fieldType.Type, valType)
			if err != nil {
				return nil, err
			}
			for k, v := range vals {
				values.Add(k, v)
			}
		} else {
			switch qpTag.Style {
			case "deepObject":
				vals := populateDeepObjectParams(qpTag, fieldType.Type, valType)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			case "form":
				vals := populateFormParams(qpTag, fieldType.Type, valType, ",", defaultValue)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			case "pipeDelimited":
				vals := populateFormParams(qpTag, fieldType.Type, valType, "|", defaultValue)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			default:
				return nil, fmt.Errorf("unsupported style: %s", qpTag.Style)
			}
		}
	}

	return globalsAlreadyPopulated, nil
}

func populateSerializedParams(tag *paramTag, objType reflect.Type, objValue reflect.Value) (map[string]string, error) {
	if isNil(objType, objValue) {
		return nil, nil
	}

	if objType.Kind() == reflect.Pointer {
		objValue = objValue.Elem()
	}

	values := map[string]string{}

	switch tag.Serialization {
	case "json":
		data, err := json.Marshal(objValue.Interface())
		if err != nil {
			return nil, fmt.Errorf("error marshaling json: %v", err)
		}
		values[tag.ParamName] = string(data)
	}

	return values, nil
}

func populateDeepObjectParams(tag *paramTag, objType reflect.Type, objValue reflect.Value) url.Values {
	values := url.Values{}

	if isNil(objType, objValue) {
		return values
	}

	if objValue.Kind() == reflect.Pointer {
		objValue = objValue.Elem()
	}

	switch objValue.Kind() {
	case reflect.Map:
		populateDeepObjectParamsMap(values, tag.ParamName, objValue)
	case reflect.Struct:
		populateDeepObjectParamsStruct(values, tag.ParamName, objValue)
	}

	return values
}

func populateDeepObjectParamsArray(qsValues url.Values, priorScope string, value reflect.Value) {
	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < value.Len(); i++ {
		qsValues.Add(priorScope, valToString(value.Index(i).Interface()))
	}
}

func populateDeepObjectParamsMap(qsValues url.Values, priorScope string, mapValue reflect.Value) {
	if mapValue.Kind() != reflect.Map {
		return
	}

	iter := mapValue.MapRange()

	for iter.Next() {
		scope := priorScope + "[" + iter.Key().String() + "]"
		iterValue := iter.Value()

		switch iterValue.Kind() {
		case reflect.Array, reflect.Slice:
			populateDeepObjectParamsArray(qsValues, scope, iterValue)
		case reflect.Map:
			populateDeepObjectParamsMap(qsValues, scope, iterValue)
		default:
			qsValues.Add(scope, valToString(iterValue.Interface()))
		}
	}
}

func populateDeepObjectParamsStruct(qsValues url.Values, priorScope string, structValue reflect.Value) {
	if structValue.Kind() != reflect.Struct {
		return
	}

	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		if isNil(field.Type, fieldValue) {
			continue
		}

		if fieldValue.Kind() == reflect.Pointer {
			fieldValue = fieldValue.Elem()
		}

		qpTag := parseQueryParamTag(field)

		if qpTag == nil {
			continue
		}

		scope := priorScope

		if !qpTag.Inline {
			scope = priorScope + "[" + qpTag.ParamName + "]"
		}

		switch fieldValue.Kind() {
		case reflect.Array, reflect.Slice:
			populateDeepObjectParamsArray(qsValues, scope, fieldValue)
		case reflect.Map:
			populateDeepObjectParamsMap(qsValues, scope, fieldValue)
		case reflect.Struct:
			switch fieldValue.Type() {
			case reflect.TypeOf(big.Int{}), reflect.TypeOf(decimal.Big{}), reflect.TypeOf(time.Time{}), reflect.TypeOf(types.Date{}):
				qsValues.Add(scope, valToString(fieldValue.Interface()))

				continue
			}

			populateDeepObjectParamsStruct(qsValues, scope, fieldValue)
		default:
			qsValues.Add(scope, valToString(fieldValue.Interface()))
		}
	}
}

func populateFormParams(tag *paramTag, objType reflect.Type, objValue reflect.Value, delimiter string, defaultValue *string) url.Values {
	return populateForm(tag.ParamName, tag.Explode, objType, objValue, delimiter, defaultValue, func(fieldType reflect.StructField) string {
		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			return ""
		}

		return qpTag.ParamName
	})
}

type paramTag struct {
	Style         string
	Explode       bool
	ParamName     string
	Serialization string

	// Inline is a special case for union/oneOf. When a wrapper struct type is
	// used, each union/oneOf value field should be inlined (e.g. not appended
	// in deepObject style with the name) as if the value was directly on the
	// parent struct field. Without this annotation, the value would not be
	// encoded by downstream logic that requires the struct field tag.
	Inline bool
}

func parseQueryParamTag(field reflect.StructField) *paramTag {
	return parseParamTag(queryParamTagKey, field, "form", true)
}
