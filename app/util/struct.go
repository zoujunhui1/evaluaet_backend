package util

import (
	"reflect"
	"strings"
)

func GetJsonFields(v interface{}) []string {
	var jsonFields []string
	val := reflect.ValueOf(v)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		if jsonTag := t.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
				fieldName = jsonTag[:commaIdx]
			} else {
				fieldName = jsonTag
			}
		}
		jsonFields = append(jsonFields, fieldName)

	}

	return jsonFields
}
