package class

import (
	"fmt"
	"strings"
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

func (a Array) ToJSONString(json *strings.Builder) ([]error) {
	var errors []error
	
	if json == nil {
		errors = append(errors, fmt.Errorf("*strings.Builder is nil"))
		return errors
	}
	
	length := len(a)

	if length == 0 {
		json.WriteString("[]")
		return nil
	}

	json.WriteString("[\n")
	for i, value := range a {
		string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, value)
		if string_conversion_error != nil {
			errors = append(errors, string_conversion_error...)
		} 

		if i < length - 1 {
			json.WriteString(",")
		}

		json.WriteString("\n")
	}
	json.WriteString("]")

	if len(errors) > 0 {
		return errors
	}

	return nil
}
