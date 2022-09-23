package common

import (
	"fmt"
	"reflect"
)

func ConvertPrimativeMapToMap(m map[string]interface{}) Map {
	var primitMap = make(map[string]interface{})
	var primitArray []interface{}
	var arrayise = Array{}
	var copy = Map{}
	for key, value := range m {
		fmt.Println("valueof: " + reflect.ValueOf(value).String() + " type: " + reflect.TypeOf(value).String())

		if reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Map) || 
		   reflect.ValueOf(value).Type() == reflect.TypeOf(copy) ||
		   reflect.ValueOf(value).Type() == reflect.TypeOf(primitMap) {
			fmt.Println("ConvertPrimativeMapToMap map")
			copy.setMap(key, ConvertPrimativeMapToMap(value.(map[string]interface{})))
		} else if reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Slice) ||
		          reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Array) ||
				  reflect.ValueOf(value).Type() == reflect.TypeOf(arrayise) ||
				  reflect.ValueOf(value).Type() == reflect.TypeOf(primitArray) {
					fmt.Println("ConvertPrimativeMapToMap array")
			copy.setArray(key, ConvertPrimativeArrayToThing(value.([]interface{})))
		} else {
			fmt.Println("ConvertPrimativeMapToMap value")
			copy.setValue(key, value)
		}		
	}
	return copy
}


func (m Map) M(s string) Map {
	//maapp := Map{}
	
	//if reflect.ValueOf(value).Type() == 

	fmt.Println(s)
	fmt.Println("valueof: " + reflect.ValueOf(m[s]).String() + " type: " + reflect.TypeOf(m[s]).String())

	return m[s].(Map)
}

func (m Map) M_self() Map {
	return m
}

func (m Map) setMap(key string, value Map) Map {
	m[key] = value
	return m
}

func (m Map) setArray(key string, value Array) Map {
	m[key] = value
	return m
}

func (m Map) setValue(key string, value interface{}) Map {
	m[key] = value
	return m
}


func (m Map) Func(s string) func(...map[string]interface{}) (map[string]interface{}) {
	return m[s].(func(...map[string]interface{}) (map[string]interface{}))
}

func (m Map) Array(s string) []interface{} {
	
	return m[s].([]interface{})
}

func (m Map) S(s string) string {
	return m[s].(string)
}

func (m Map) InterfaceString(s string) interface{} {
	return m[s].(interface{})
}

func (m Map) Spt(s string) *string {
	return m[s].(*string)
}

func (m Map) PrimArray(s string) []string {
	return m[s].([]string)
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}

func (m Map) M_Value(s string) (interface{}, bool) {
	if m[s] != nil {
		return m[s].(interface{}), true
	}
	return nil, false
}