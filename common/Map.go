package common

import (
	"fmt"
	"reflect"
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
		case "common.Map":
		case "map[string]interface {}":
			clone := Map{}
			for clone_key, clone_value := range value.(map[string]interface{}) {
				clone_rep := fmt.Sprintf("%T", clone_value)
				switch clone_rep {
				case "common.Map":
					clone[clone_key] = ConvertPrimitiveMapToMap(clone_value.(map[string]interface{}))
					break
				case "string": 
					clone[clone_key] = clone_value
				case "reflect.Value":
					clone[clone_key] = reflect.ValueOf(clone_value)
				default:
					panic(fmt.Errorf("Map.M: type %s is not supported please implement", clone_rep))
				}
			}
			newMap[key] = clone
			break
		case "func(...map[string]interface {}) map[string]interface {}":
			newMap[key] = value
			break
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
		case "reflect.Value":
			reflect_value := reflect.ValueOf(m[s])
			array := Array{reflect.ValueOf(reflect_value)}
			return array
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement", rep))
	}

	return nil
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
	case "reflect.Value":
		reflect_value := reflect.ValueOf(m[s])
		return fmt.Sprintf("%s", reflect_value)
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}

	return ""
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}
