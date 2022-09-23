package common

import (
	"fmt"
)

type Array []interface{}

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
		case "common.Map":
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
