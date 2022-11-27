package class

import (
	"fmt"
)

type Array []interface{}

func ToArray(a interface{}) (*Array, []error) {
	if a == nil {
		return nil, nil
	}

	var errors []error
	array := Array{}
	rep := fmt.Sprintf("%T", a)
	switch rep {
	case "*[]string": 
		for _, value := range *(a.(*[]string)) {
			array = append(array, value)
		}
	case "*[]class.Record":
		for _, value := range *(a.(*[]Record)) {
			array = append(array, value)
		}
	default:
		errors = append(errors, fmt.Errorf("Array.ToArray: type is not supported please implement: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &array, nil
}

func (a Array) ToJSONString() (*string, []error) {
	var errors []error
	length := len(a)

	json := ""
	if length == 0 {
		json = "[]"
		return &json, nil
	}

	json += "[\n"
	for i, value := range a {
		string_conversion, string_conversion_error := ConvertInterfaceValueToJSONStringValue(value)
		if string_conversion_error != nil {
			errors = append(errors, string_conversion_error...)
		} else {
			json += *string_conversion
		}

		if i < length - 1 {
			json += ","
		}

		json += "\n"
	}
	json += "]"

	if len(errors) > 0 {
		return nil, errors
	}

	return &json, nil
}
