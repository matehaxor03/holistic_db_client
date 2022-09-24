package class

import (
	"fmt"
	"reflect"
	"strings"
)

type Map map[string]interface{}

func ConvertPrimitiveValueToMap(f interface{}) Map {
	m := Map{}
	rep := fmt.Sprintf("%T", f)
	switch rep {
		case "reflect.Value":
			reflect_value := reflect.ValueOf(f)
			m["value"] = reflect_value
			return m
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement", rep))
	}

	return nil
}


func ConvertPrimitiveMapToMap(m map[string]interface{}) Map {
	if m == nil {
		return nil
	}

	newMap := Map{}
	for key, value := range m {
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "class.Map":
		case "map[string]interface {}":
			clone := Map{}
			for clone_key, clone_value := range value.(map[string]interface{}) {
				clone_rep := fmt.Sprintf("%T", clone_value)
				switch clone_rep {
				case "class.Map":
					clone[clone_key] = ConvertPrimitiveMapToMap(clone_value.(map[string]interface{}))
					break
				case "string": 
					clone[clone_key] = clone_value
				case "reflect.Value":
					if strings.Contains(clone_key, "|") && len(strings.Split(clone_key, "|")) == 2 {
						parts := strings.Split(clone_key, "|")
						//inner_clone_key := parts[0]
						inner_clone_key_data_type := parts[1]
						switch inner_clone_key_data_type {
						case "string":
						case "data_type":	
							fmt.Println(clone_key)
							clone[clone_key] = fmt.Sprintf("%s", reflect.ValueOf(clone_value).Interface())
							break
						case "[]string":
							raw_data := fmt.Sprintf("%s",  reflect.ValueOf(clone_value).Interface())
							if strings.HasPrefix(raw_data, "[") &&
							   strings.HasSuffix(raw_data, "]") {
								raw_data = raw_data[1:len(raw_data)-1]
								array_to_copy := strings.Split(raw_data, " ")
								string_array := make([]string, len(array_to_copy))
								for i, _  := range array_to_copy {
									string_array[i] = array_to_copy[i]
								}
								clone[clone_key] = string_array
							} else {
								panic(fmt.Errorf("Map.M: data for data type: '%s' for %s->%s was in the wrong format and neds to be [data1 data2 ...]", inner_clone_key_data_type, key, clone_key))
							}
							//fmt.Println(fmt.Sprintf("%s %s %s", clone_key, reflect.ValueOf(clone_value).Interface()))
							break
						default:
							panic(fmt.Errorf("Map.M: data type: '%s' not supported for %s->%s please implement", inner_clone_key_data_type, key, clone_key))
						}

					} else {
						panic(fmt.Errorf("Map.M: cannot determine field struct for field '%s' key needs to be in format fieldName|datatype", clone_key))
					}
					
					inner_rep := fmt.Sprintf("%T", reflect.ValueOf(clone_value).Interface())
					
					//zype := fmt.Sprintf("%T", clone_value)
					//panic(clone_key)
					fmt.Println(fmt.Sprintf("%s %s %s", clone_key, reflect.ValueOf(clone_value).Interface(), inner_rep))
					//clone[clone_key] = reflect.ValueOf(clone_value).Interface()
				default:
					panic(fmt.Errorf("Map.M: type %s is not supported please implement", clone_rep))
				}
			}
			newMap[key] = clone
			break
		case "func(...map[string]interface {}) *class.Result":
		case "*class.Result":
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
	case "class.Map":
		return m[s].(Map)
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) SetMap(s string, zap Map) {
	rep := fmt.Sprintf("%T", zap)
	
	switch rep {
	case "class.Map":
		m[s] = zap
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement for %s", rep, s))
	}
}

func (m Map) ToJSONString() string {
	json := "{\n"
	keys := m.Keys()
	length := len(keys)
	for i, key := range keys {
		json = json + "\"" + key + "\":"

		value := m[key]
		rep := fmt.Sprintf("%T", value)
		switch rep {
		case "string":
			json = json + "\"" + value.(string) + "\""
		case "class.Map":
			json += value.(Map).ToJSONString()
		case "class.Array":
			json += value.(Array).ToJSONString()
		case "[]string":
			json += "["
			array_length := len(m[key].([]string))
			for array_index, array_value := range m[key].([]string) {
				json += "\"" + array_value + "\""
				if array_index < array_length {
					json += ","
				}
			}
			json += "]"			
		default:
			panic(fmt.Errorf("Map.ToJSONString: type %s is not supported please implement for %s", rep, key))
		}

		if i < length {
			json += ","
		} 
		json += "\n"
	}
	json += "}"
	return json
}

func (m Map) A(s string) Array {
	rep := fmt.Sprintf("%T", m[s])
	if m[s] == nil {
		return nil
	}

	switch rep {
		case "class.Array":
			return m[s].(Array)
		case "[]string":
			newArray := Array{}
			for _, v := range m[s].([]string) {
				newArray = append(newArray, v)
			}
			return newArray

		/*case "reflect.Value":
			reflect_value := reflect.ValueOf(m[s])
			array := Array{reflect.ValueOf(reflect_value)}
			return array*/
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement for field: %s", rep, s))
	}
}

func (m Map) SetArray(s string, array Array) {
	rep := fmt.Sprintf("%T", array)
	switch rep {
		case "class.Array":
		 m[s] = array
	default:
		panic(fmt.Errorf("Map.A: type %s is not supported please implement for field: %s", rep, s))
	}
}

/*
func (m Map) Func(s string) func(...map[string]interface{}) (map[string]interface{}) {
	return m[s].(func(...map[string]interface{}) (map[string]interface{}))
}*/

func (m Map) Func(s string) func(...map[string]interface{}) (*Result) {
	return m[s].(func(...map[string]interface{}) (*Result))
}

func (m Map) SetFunc(s string, function func(...map[string]interface{}) (*Result)) {
	m[s] = function
}

func (m Map) Array(s string) []interface{} {
	
	return m[s].([]interface{})
}

func (m Map) S(s string) (string, error) {
	rep := fmt.Sprintf("%T", m[s])
	if m[s] == nil {
		return "",fmt.Errorf("Map.S(s string): field: %s is not set", s)
	}
	
	switch rep {
	case "string":
		return m[s].(string), nil
		break
	case "reflect.Value":
		return fmt.Sprintf("%s", reflect.ValueOf(m[s]).Interface()), nil
	}

	return "", fmt.Errorf("Map.S(s string): datatype: '%s' is not supported please implement when fetching field: %s", rep, s)
}

func (m Map) SetString(s string, value string) {
	rep := fmt.Sprintf("%T", value)
	
	switch rep {
	case "string":
		m[s] = value
		break
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}
}

func (m Map) GetResult(s string) *Result {
	rep := fmt.Sprintf("%T", m[s])
	
	switch rep {
	case "*class.Result":
		return m[s].(*Result)
		break
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}
