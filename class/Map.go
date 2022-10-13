package class

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

type Map map[string]interface{}

func (m Map) M(s string) Map {
	rep := fmt.Sprintf("%T", m[s])
	
	switch rep {
	case "class.Map":
		return m[s].(Map)
	default:
		panic(fmt.Errorf("Map.M: type %s is not supported please implement for key %s", rep, s))
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

func (m Map) IsNil(s string) (bool) {
	if m[s] == nil {
		return true
	}

	string_value := fmt.Sprintf("%s", m[s])
	
	if string_value == "<nil>" {
		return true
	}

	rep := fmt.Sprintf("%T", m[s])	
	
	if string_value == "%!s(" + rep + "=<nil>)" {
		return true
	}

	
	return false
}

func (m Map) ToJSONString() string {
	json := "{\n"
	keys := m.Keys()
	length := len(keys)
	for i, key := range keys {
		json = json + "\"" + key + "\":"
		if m.IsNil(key) {
			json = json + "null"
		} else {
			value := m[key]
			rep := fmt.Sprintf("%T", value)
			switch rep {
			case "string":
				json = json + "\"" + value.(string) + "\""
			case "*string":
				string_pt := (value).(*string)
				if string_pt == nil {
					json = json + "null"
				} else {
					json = json + "\"" + (*string_pt) + "\""
				}
			case "bool":
				temp := value.(bool)
				if temp {
					json = json + "true"
				} else {
					json = json + "false"
				}
			case "*bool":
				temp := *(value.(*bool))
				if temp {
					json = json + "true"
				} else {
					json = json + "false"
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
			case "*class.Host":
				json += (*(value.(*Host))).ToJSONString()
			case "*class.Credentials":
				json += (*(value.(*Credentials))).ToJSONString()
			case "*class.DatabaseCreateOptions":
				json += (*(value.(*DatabaseCreateOptions))).ToJSONString()
			case "*class.Database":
				json += (*(value.(*Database))).ToJSONString()
			case "*class.Client":
				json += (*(value.(*Client))).ToJSONString()
			case "map[string]map[string][][]string":
				json = json + "\"map[string]map[string][][]string\""
			case "*uint64":
				json = json + strconv.FormatUint(*(value.(*uint64)), 10)
			case "uint64":
				json = json + strconv.FormatUint(value.(uint64), 10)
			case "*int64":
				json = json + strconv.FormatInt(*(value.(*int64)), 10)
			case "int64":
				json = json + strconv.FormatInt(value.(int64), 10)
			case "*int":
				json = json + strconv.FormatInt(int64(*(value.(*int))), 10)
			case "int":
				json = json + strconv.FormatInt(int64(value.(int)), 10)
			default:
				panic(fmt.Errorf("Map.ToJSONString: type %s is not supported please implement for %s", rep, key))
			}
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

func (m Map) GetInt64(s string) (*int64, []error) {
	var errors []error
	var result *int64

	if m[s] == nil {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*int64":
		result = m[s].(*int64)
	case "int":
		value := int64(m[s].(int))
		result = &value
	case "*int":
		value := int64(*(m[s].(*int)))
		result = &value
	case "*string":
		value, value_error := strconv.ParseInt((*(m[s].(*string))), 10, 64)
		if value_error != nil {
			errors = append(errors,fmt.Errorf("Map.GetInt64: cannot convert *string value to int64"))
		} else {
			result = &value
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetInt64: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetInt(s string) (*int, []error) {
	var errors []error
	var result *int

	if m[s] == nil {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "int":
		value := m[s].(int)
		result = &value
	case "*int":
		value := *(m[s].(*int))
		result = &value
	case "*string":
		bit_size := strconv.IntSize
		value, value_error := strconv.ParseInt((*(m[s].(*string))), 10, bit_size)
		if value_error != nil {
			errors = append(errors,fmt.Errorf("Map.GetInt: cannot convert *string value to int"))
		} else {
			temp := int(value)
			result = &temp
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetInt: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetInt64(s string, v *int64) {
	m[s] = v
}

func (m Map) GetUInt64(s string) (*uint64, []error) {
	var errors []error
	var result *uint64
	if m[s] == nil {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*int64":
		value := *(m[s].(*int64))
		if value >= 0 {
			temp := uint64(value)
			result = &temp
		} else {
			errors = append(errors,fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		value := (m[s].(int))
		if value >= 0 {
			temp := uint64(value)
			result = &temp
		} else {
			errors = append(errors,fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		value := *(m[s].(*int))
		if value >= 0 {
			temp := uint64(value)
			result = &temp
		} else {
			errors = append(errors,fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*uint64":
		value := *(m[s].(*uint64))
		result = &value
	case "uint64":
		value := (m[s].(uint64))
		result = &value
	case "*string":
		value, value_error := strconv.ParseUint((*(m[s].(*string))), 10, 64)
		if value_error != nil {
			errors = append(errors,fmt.Errorf("Map.GetUInt64: cannot convert *string value to uint64"))
		} else {
			result = &value
		}
	default:
		errors = append(errors,fmt.Errorf("Map.GetUInt64: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetUInt64(s string, v *uint64) {
	m[s] = v
}

func (m Map) SetTime(s string, value *time.Time) {
	m[s] = value
}

func (m Map) GetTime(s string) (*time.Time, []error) {
	var errors []error
	var result *time.Time
	
	if m[s] == nil {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*time.Time":
		value := *(m[s].(*time.Time))
		result = &value
	case "time.Time":
		value :=  m[s].(time.Time)
		result = &value
	case "*string": 
		//todo: parse for null
		value, value_errors := time.Parse("2006-01-02 15:04:05.000000", *(m[s].(*string)))
		if value_errors != nil {
			errors = append(errors, value_errors)
		} else {
			result = &value
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetTime: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return result, nil
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
		if m.IsNil(key) {
			clone[key] = nil
			continue
		}

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
		case "*class.Database":
			clone[key] = (*(current.(*Database))).Clone()
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
		case "*class.Table":
			clone[key] = (*(current.(*Table))).Clone()
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
		case "uint64": 
			clone[key] = current.(uint64)	
		case "int64": 
			clone[key] = current.(int64)	
		case "*int64": 
			int64_value := *(current.(*int64))	
			clone[key] = &int64_value
		case "*int": 
			int_value := *(current.(*int))	
			clone[key] = &int_value
		case "*uint64": 
			uint64_value := *(current.(*uint64))	
			clone[key] = &uint64_value
		case "<nil>":
			clone[key] = nil
		default:
			panic(fmt.Errorf("Map.Clone: type %s is not supported please implement", rep))
		}
	}

	return clone
}


