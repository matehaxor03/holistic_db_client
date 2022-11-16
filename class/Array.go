package class

import (
	"fmt"
	"time"
)

type Array []interface{}

func NewArrayOfStrings(a *[]string) *Array {
	if a == nil {
		return nil
	}

	array := Array{}
	for _, value := range *a {
		array = append(array, value)
	}
	return &array
}

func (a Array) ToJSONString() (*string, []error) {
	var errors []error
	json := "[\n"
	length := len(a)
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

func (a Array) ToPrimativeArray() []string {
	var results []string
	for _, value := range a {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			results = append(results, value.(string))
			break
		default:
			panic(fmt.Errorf("Array.ToPrimativeArray: type %s is not supported please implement", rep))
		}
	}
	return results
}

func (a Array) Clone() (*Array, []error) {
	var errors []error
	clone := Array{}
	for _, current := range a {
		if current == nil {
			clone = append(clone, nil)
			continue
		}
	
		string_value := fmt.Sprintf("%s", current)
		if string_value == "<nil>" {
			clone = append(clone, nil)
			continue
		}
	
		rep := fmt.Sprintf("%T", current)
	
		if string_value == "%!s("+rep+"=<nil>)" {
			clone = append(clone, nil)
			continue
		}

		switch rep {
		case "string":
			clone = append(clone, current)
			break
		case "*string":
			value := CloneString(current.(*string))
			clone = append(clone, *value)
			break
		case "class.Map":
			cloned_map, cloned_map_errors := current.(Map).Clone()
			if cloned_map_errors != nil {
				errors = append(errors, cloned_map_errors...)
			} else {
				clone = append(clone, cloned_map)
			}
			break
		case "*class.Map":
			cloned_map, cloned_map_errors := current.(*Map).Clone()
			if cloned_map_errors != nil {
				errors = append(errors, cloned_map_errors...)
			} else {
				clone = append(clone, cloned_map)
			}
			break
		case "class.Array":
			cloned_array, cloned_array_errors := current.(Array).Clone()
			if cloned_array_errors != nil {
				errors = append(errors, cloned_array_errors...)
			} else {
				clone = append(clone, cloned_array)
			}
			break
		case "*class.Array":
			cloned_array, cloned_array_errors := current.(*Array).Clone()
			if cloned_array_errors != nil {
				errors = append(errors, cloned_array_errors...)
			} else {
				clone = append(clone, cloned_array)
			}
			break
		case "bool":
			clone = append(clone, current.(bool))
		case "*bool":
			bool_value := *(current.(*bool))
			clone = append(clone, &bool_value)
		case "time.Time":
			clone = append(clone, current.(time.Time))
		case "*time.Time":
			time_value := *(current.(*time.Time))
			clone = append(clone, &time_value)
		case "int":
			clone = append(clone, current.(int))
		case "*int":
			int_value := *(current.(*int))
			clone = append(clone, &int_value)
		case "int8":
			clone = append(clone, current.(int8))
		case "*int8":
			int8_value := *(current.(*int8))
			clone = append(clone, &int8_value)
		case "int16":
			clone = append(clone, current.(int16))
		case "*int16":
			int16_value := *(current.(*int16))
			clone = append(clone,  &int16_value)
		case "int32":
			clone = append(clone, current.(int32))
		case "*int32":
			int32_value := *(current.(*int32))
			clone = append(clone, &int32_value)
		case "int64":
			clone = append(clone, current.(int64))
		case "*int64":
			int64_value := *(current.(*int64))
			clone = append(clone, &int64_value)
		case "uint":
			clone = append(clone, current.(uint))
		case "*uint":
			uint_value := *(current.(*uint))
			clone = append(clone, &uint_value)
		case "uint8":
			clone = append(clone, current.(uint8))
		case "*uint8":
			uint8_value := *(current.(*uint8))
			clone = append(clone, &uint8_value)
		case "uint16":
			clone = append(clone, current.(uint16))
		case "*uint16":
			uint16_value := *(current.(*uint16))
			clone = append(clone, &uint16_value)
		case "uint32":
			clone = append(clone, current.(uint32))
		case "*uint32":
			uint32_value := *(current.(*uint32))
			clone = append(clone, &uint32_value)
		case "uint64":
			clone = append(clone, current.(uint64))
		case "*uint64":
			uint64_value := *(current.(*uint64))
			clone = append(clone,  &uint64_value)
		case "[]error":
			clone = append(clone, current.([]error))
		case "*[]error":
			errs := *(current.(*[]error))
			clone = append(clone, &errs)
		default:
			errors = append(errors, fmt.Errorf("Array.Clone: type %s is not supported please implement", rep))
		}
	}

	if len(errors) > 0 { 
		return nil, errors
	}

	return &clone, nil
}
