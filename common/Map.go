package common

import (
	"fmt"
//	"reflect"
)

type Map map[string]interface{}

func ConvertPrimitiveMapToMap(m map[string]interface{}) Map {
	if m == nil {
		return nil
	}

	newMap := Map{}
	for key, value := range m {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			nestedMap := Map{}
			newMap[key] = nestedMap
			break
		case "common.Map":
			// todo deep clone
			clone := Map{}
			for clone_key, clone_value := range value.(map[string]interface{}) {
				clone[clone_key] = clone_value
			}
			newMap[key] = clone
		default:
			panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
		}
	}

	return newMap
}

func KeysForMap(m map[string]interface{}) []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}


func (m Map) M(s string) Map {
	rep := fmt.Sprintf("%T", m[s])
	
	switch rep {
	/*	
	case "string":
		key := m[s].(string)
		newMap := Map{}
		m[key] = newMap
		return newMap
		break*/
	case "common.Map":
		return m[s].(Map)
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) A(s string) Array {
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "common.Array":
		return m[s].(Array)
		break
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement", rep))
	}

	return nil
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
	rep := fmt.Sprintf("%T", m[s])
	
	switch rep {
	case "string":
		return m[s].(string)
		break
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}

	return ""
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