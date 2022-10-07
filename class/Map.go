package class

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Map map[string]interface{}

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
		panic(fmt.Errorf("Map.SetMap: type %s is not supported please implement for %s", rep, s))
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
		case "func(string, *string, string, string) []error", "func(class.Map) []error", "*func(class.Map) []error":
			json = json + fmt.Sprintf("\"%s\"", rep)
		case "bool": 
			boolValue := fmt.Sprintf("%B", reflect.ValueOf(value).Interface())
			if boolValue  == "%!B(bool=false)" {
				json = json + fmt.Sprintf("false")
			} else if boolValue == "%!B(bool=true)" {
				json = json + fmt.Sprintf("true")
			} else if boolValue == "%!B(bool=nil)" {
				json = json + fmt.Sprintf("null")
			} else {
				panic(fmt.Errorf("Map.ToJSONString: type %s is not supported please implement for %s %s", rep, key, boolValue))
			}
		case "*class.Host":
			json += (*(value.(*Host))).ToJSONString()
		case "*class.Credentials":
			json += (*(value.(*Credentials))).ToJSONString()
		case "*class.DatabaseCreateOptions":
			json += (*(value.(*DatabaseCreateOptions))).ToJSONString()
		case "map[string]map[string][][]string":
			json = json + "\"map[string]map[string][][]string\""
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
	if m[s] == nil {
		return nil
	}
	
	rep := fmt.Sprintf("%T", m[s])
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
		panic(fmt.Errorf("Map.SetArray: type %s is not supported please implement for field: %s", rep, s))
	}
}

func (m Map) GetType(s string) (string) {
	return fmt.Sprintf("%T", m[s])
}

func (m Map) Func(s string) (func(Map) []error) {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
		case "func(class.Map) []error":
			return m[s].(func(Map) []error)
		case "*func(class.Map) []error":
			value :=  m[s].(*func(Map) []error)
			return *value
	default:
		panic(fmt.Errorf("Map.Func: type %s is not supported please implement for field: %s", rep, s))
	}
	
	return nil
}

func (m Map) SetFunc(s string, function func(Map) ([]error)) {
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
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			return &s
		} else {
			return nil
		}
	default:
		panic(fmt.Errorf("Map.S: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) GetObject(s string) (interface{}) {
	return m[s]
}

func (m Map) B(s string) (*bool) {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "bool":
		value := m[s].(bool)
		newValue := value
		return &newValue
		break
	case "*bool":
		if fmt.Sprintf("%s", m[s]) != "%!s(*bool=<nil>)" {
			newValue := *((m[s]).(*bool))
			return &newValue
		} else {
			return nil
		}
		break
	default:
		panic(fmt.Errorf("Map.B: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) SetBool(s string, value *bool) {
	m[s] = value
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
		panic(fmt.Errorf("Map.SetString: type %s is not supported please implement", rep))
	}
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}

func (m Map) HasKey(key string) bool {
	keys := m.Keys()
	for _, compare_key := range keys {
		if key == compare_key {
			return true
		}
	}
	return false
}

func (m Map) GetInt64(s string) *int64 {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*int64":
		return m[s].(*int64)
	case "int":
		value := int64(m[s].(int))
		return &value
	case "*int":
		value := int64(*(m[s].(*int)))
		return &value
	default:
		panic(fmt.Errorf("Map.GetInt64: type %s is not supported please implement", rep))
	}

	return nil
}

func (m Map) GetTime(s string) *time.Time {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*time.Time":
		return m[s].(*time.Time)
	}
	
	return nil
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
			cloneString := strings.Clone(*(m.S(key)))
			clone.SetString(key, &cloneString)
			break
		case "class.Map":
			clone[key] = current.(Map).Clone()
			break
		case "class.Array":
			clone[key] = current.(Array).Clone()
			break
		case "func(class.Map) []error", "map[string]map[string][][]string", "*func(class.Map) []error":
			clone[key] = current
			break	
		case "*class.Credentials":
			clone[key] = (*(current.(*Credentials))).Clone()
			break
		case "*class.DatabaseCreateOptions":
			clone[key] = (*(current.(*DatabaseCreateOptions))).Clone()
			break
		case "*class.Host":
			clone[key] = (*(current.(*Host))).Clone()
			break
		case "*class.Client":
			clone[key] = (*(current.(*Client))).Clone()
			break
		case "*class.Grant":
			clone[key] = (*(current.(*Grant))).Clone()
			break
		case "*class.User":
			clone[key] = (*(current.(*User))).Clone()
			break
		case "*class.DomainName":
			clone[key] = (*(current.(*DomainName))).Clone()
			break
		case "bool": 
			clone[key] = current.(bool)	
		case "*bool": 
			bool_value := *(current.(*bool))
			clone[key] = &bool_value
		case "*time.Time": 
			clone[key] = current.(*time.Time)	
		case "int": 
			clone[key] = current.(int)	
		case "<nil>":
			clone[key] = nil
		default:
			panic(fmt.Errorf("Map.Clone: type %s is not supported please implement", rep))
		}
	}

	return clone
}


