package class

import (
	"fmt"
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
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "*string":
			if fmt.Sprintf("%s", value) != "%!s(*string=<nil>)" {
				json = json + "null"
			} else {
				json = json + "\"" + (*(value).(*string)) + "\""
			}
		case "string":
			json = json + "\"" + value.(string) + "\""
		case "class.Map":
			x, x_errors := value.(Map).ToJSONString()
			if x_errors != nil {
				errors = append(errors, x_errors...)
			} else {
				json += *x
			}
		case "class.Array":
			x, x_errors := value.(Array).ToJSONString()
			if x_errors != nil {
				errors = append(errors, x_errors...)
			} else {
				json += *x
			}
		case "*class.Array":
			x, x_errors := (*(value.(*Array))).ToJSONString()
			if x_errors != nil {
				errors = append(errors, x_errors...)
			} else {
				json += *x
			}
		case "reflect.Value":
			json = json + fmt.Sprintf("\"%s\"", value)
		case "func(class.Map) []error":
			json = json + fmt.Sprintf("\"func(class.Map) []error\"")
		case "<nil>":
			json = json + fmt.Sprintf("null")
		default:
			errors = append(errors, (fmt.Errorf("Array.ToJSONString: type %s is not supported please implement", rep)))
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

func (a Array) Clone() *Array {
	clone := Array{}
	for _, current := range a {
		rep := fmt.Sprintf("%T", current)
		switch rep {
		case "string":
			clone = append(clone, current)
			break
		case "class.Map":
			clone = append(clone, current.(Map).Clone())
			break
		case "class.Array":
			clone = append(clone, current.(Array).Clone())
			break
		default:
			panic(fmt.Errorf("Array.Clone: type %s is not supported please implement", rep))
		}
	}
	return &clone
}
