package class

import (
	"fmt"
	"reflect"
	"strings"
)

type Map map[string]interface{}

func ConvertPrimitiveReflectValueToValue(v reflect.Value) any {
	rep := fmt.Sprintf("%T", v)
	switch rep {
		case "reflect.Value":
			return reflect.ValueOf(v).Interface()
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
				case "class.Array":
					clone[clone_key] = clone_value
					break
				case "string": 
					clone.SetString(clone_key, clone_value.(*string))
				case "reflect.Value":
					if strings.Contains(clone_key, "|") && len(strings.Split(clone_key, "|")) == 2 {
						parts := strings.Split(clone_key, "|")
						inner_clone_key_data_type := parts[1]
						switch inner_clone_key_data_type {
						case "string":
						case "data_type":
							s := strings.Clone(fmt.Sprintf("%s", value))
							clone[clone_key] = &s
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
							break
						default:
							panic(fmt.Errorf("Map.M: data type: '%s' not supported for %s->%s please implement", inner_clone_key_data_type, key, clone_key))
						}

					} else {
						panic(fmt.Errorf("Map.M: cannot determine field struct for field '%s' key needs to be in format fieldName|datatype", clone_key))
					}
				default:
					panic(fmt.Errorf("Map.M: type %s is not supported please implement", clone_rep))
				}
			}
			newMap[key] = clone
			break
		case "func(...map[string]interface {}) *class.Result":
		case "func(class.Map) []error":
			newMap[key] = value
			break
		case "*string":
			if fmt.Sprintf("%s", value) != "%!s(*string=<nil>)" {
				s := strings.Clone(*((value).(*string)))
				newMap[key] = &s
			} else {
				newMap[key] = nil
			}
			break
		case "string":
		case "data_type":
			s := strings.Clone(fmt.Sprintf("%s", value))
			newMap[key] = &s
			break
		case "class.Array":
			newMap[key] = value.(Array).Clone()
			break
		case "func(class.Map)":
			newMap[key] = value.(Map).Clone()
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
		case "<nil>":
			json = json + "null"
		case "string":
			json = json + "\"" + value.(string) + "\""
		case "*string":
			string_pt := (value).(*string)
			if string_pt == nil {
				json = json + "null"
			} else {
				json = json + "\"" + (*string_pt) + "\""
			}
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
		case "reflect.Value":
			json = json + fmt.Sprintf("\"%s\"", reflect.ValueOf(value).Interface())
		case "func(class.Map) []error": 
			json = json + fmt.Sprintf("\"%s\"", reflect.ValueOf(value).Interface())
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

func (m Map) A(s string) (Array) {
	rep := fmt.Sprintf("%T", m[s])
	//fmt.Println(rep)
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

func (m Map) Func(s string) func(Map) ([]error) {
	return m[s].(func(Map) ([]error))
}

func (m Map) SetFunc(s string, function func(...map[string]interface{}) ([]error)) {
	m[s] = function
}

func (m Map) Array(s string) []interface{} {
	
	return m[s].([]interface{})
}

func (m Map) S(s string) (*string) {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "string":
		value := m[s].(string)
		newValue := strings.Clone(value)
		
		return &newValue
		break
	case "reflect.Value":
		value := fmt.Sprintf("%s", reflect.ValueOf(m[s]).Interface())
		newValue := strings.Clone(value)
		return &newValue
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			return &s
		} else {
			return nil
		}
		break
	default:
		panic(fmt.Errorf("Map.S: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) SetString(s string, value *string) {
	rep := fmt.Sprintf("%T", value)
	
	switch rep {
	case "string":
		m[s] = value
		break
	case "*string":
		m[s] = value
		break
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
	}
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}

func (m Map) Values() Array {
	array := Array{}
	for _, f := range m {
		array = append(array, f)
	}
	return array
}

func (m Map) Clone() Map {
	clone := Map{}
	keys := m.Keys()
	fmt.Println(m.ToJSONString())
	for _, key := range keys {
		current := m[key] 
		rep := fmt.Sprintf("%T", current)
		switch rep {
		case "string":
			cloneString := strings.Clone(*(m.S(key)))
			clone.SetString(key, &cloneString)
			break	
		case "*string":
			if fmt.Sprintf("%s", m[key]) == "%!s(*string=<nil>)" {
				clone[key] = nil
				continue
			}
			
			parts := strings.Split(key, "|")
			keyValueType := parts[1]
			switch  keyValueType {
			case "string":
				value := *(m[key].(*string))
				cloneString := strings.Clone(value)
				clone.SetString(key, &cloneString)
				break
			default:
				panic(fmt.Errorf("Map.M: type %s is not supported please implement", keyValueType))
			}
			break
		case "class.Map":
			clone[key] = current.(Map).Clone()
			break
		case "class.Array":
			clone[key] = current.(Array).Clone()
			break
		case "func(class.Map) []error":
			clone[key] = current
			break			
		default:
			panic(fmt.Errorf("Map.M: type %s is not supported please implement", rep))
		}
	}

	return clone
}


