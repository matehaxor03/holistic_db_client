package common

import (
	"fmt"
	"reflect"
)

type Map map[string]interface{}
type Array []interface{}

func ConvertPrimativeArrayToThing(a []interface{}) Array {
	var primitMap = make(map[string]interface{})
	var primitArray []interface{}
	var mappie = Map{}
	var copy = Array{}
	for _, array := range a {
		if reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Map) || 
		   reflect.ValueOf(array).Type() == reflect.TypeOf(mappie) ||
		   reflect.ValueOf(array).Type() == reflect.TypeOf(primitMap) {
			fmt.Println("ConvertPrimativeArrayToThing map")

			thing := ConvertPrimativeMapToMap(array.(map[string]interface{}))
			copy = append(copy, thing)
		} else if reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Slice) ||
		          reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Array) || 
				  reflect.ValueOf(array).Type() == reflect.TypeOf(copy) || 
				  reflect.ValueOf(array).Type() == reflect.TypeOf(primitArray)  {
					fmt.Println("ConvertPrimativeArrayToThing array")
			thing := ConvertPrimativeArrayToThing(array.([]interface{}))
			copy = append(copy, thing)
		} else {
			fmt.Println("ConvertPrimativeArrayToThing value")
			copy = append(copy, array)
		}	
	}
	return copy
}



func ConvertIntefaceArrayToStringArray(aInterface []interface{}) []string{
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}


func KeysForMap(m map[string]interface{}) []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}