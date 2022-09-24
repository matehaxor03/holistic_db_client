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

func ConvertPrimitiveReflectValueArrayToArray(a []reflect.Value) Array {
	array := Array{}
	rep := fmt.Sprintf("%T", a)
	switch rep {
		case "[]reflect.Value":
			length := len(a)
			for i := 0; i < length; i++ {
				array = append(array, ConvertPrimitiveReflectValueToValue(a[i]))
			}
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement", rep))
	}

	return array
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

func (a Array) ToJSONString() string {
	json := "[\n"
	length := len(a)
	for i, value := range a {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			json = json + "\"" + value.(string) + "\""
		case "class.Map":
			json += value.(Map).ToJSONString()
		case "class.Array":
			json += value.(Array).ToJSONString()
		case "reflect.Value":
			fmt.Println("trying to draw refelect array")
			json = json + fmt.Sprintf("\"%s\"", value)
		default:
			panic(fmt.Errorf("Array.ToJSONString: type %s is not supported please implement", rep))
		}

		if i < length {
			json += ","
		} 

		json += "\n"
	}
	json += "]"
	return json
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
			reflect_result_value := fmt.Sprintf("%s", reflect_value)
			if reflect_result_value == "reflect.Value" {
				panic("Array.ToPrimativeArray: failed to unpack primitive")
			}
			results = append(results, reflect_result_value)
			break
		default:
			panic(fmt.Errorf("Array.ToPrimativeArray: type %s is not supported please implement", rep))
		}
	}
	return results
}
