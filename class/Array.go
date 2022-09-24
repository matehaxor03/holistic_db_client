package class

import (
	"fmt"
	"reflect"
)

type Array []interface{}

func ConvertReflectArrayToPrimativeArray(a []reflect.Value) []reflect.Value {
	length := len(a)
	copy := make([]reflect.Value, length)
	for i := 0; i < length; i++ {
		copy = append(copy, a[i])
	}

	return copy
}

func ConvertPrimativeArrayToArray(a []interface{}) Array {
	if a == nil {
		return nil
	}

	array := Array{}
	for _, value := range a {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			array = append(array, value)
			break
		case "class.Map":
			// todo deep copy map
			array = append(array, value)
		default:
			panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
		}
	}
	return array
}

func ConvertPrimativeArrayOfMapsToArray(a []map[string]interface{}) Array {
	if a == nil {
		return nil
	}
	
	array := Array{}
	for _, value := range a {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			array = append(array, value)
			break
		case "common.Map":
			array = append(array, value)
			break
		case "map[string]interface {}":
			array = append(array, ConvertPrimitiveMapToMap(value))
			break
		default:
			panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
		}
	}
	return array
}

func ConvertIntefaceArrayToStringArray(aInterface []interface{}) []string{
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}

func (a Array) ToPrimativeArray() []string {
	var results []string 
	for _, value := range a {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			results = append(results, value.(string))
			break
		case "reflect.Value":
			reflect_value := reflect.ValueOf(value)
			results = append(results, fmt.Sprintf("%s", reflect_value))
			break
		default:
			panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
		}
	}
	return results
}
