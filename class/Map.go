package class

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Map map[string]interface{}

func (m Map) M(s string) *Map {
	var errors []error
	if m.IsNil(s) {
		return nil
	}

	var result *Map
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "class.Map":
		value := m[s].(Map)
		result = &value
	case "*class.Map":
		value := *(m[s].(*Map))
		result = &value
	default:
		errors = append(errors, fmt.Errorf("Map.M: type %s is not supported please implement for key %s", rep, s))
	}

	if len(errors) > 0 {
		return nil
	}

	return result
}

func (m Map) SetMap(s string, zap *Map) {
	m[s] = zap
}

func (m Map) IsNil(s string) bool {
	if m[s] == nil {
		return true
	}

	string_value := fmt.Sprintf("%s", m[s])

	if string_value == "<nil>" {
		return true
	}

	rep := fmt.Sprintf("%T", m[s])

	if string_value == "%!s("+rep+"=<nil>)" {
		return true
	}

	return false
}

func (m Map) IsBool(s string) bool {
	type_of := m.GetType(s)
	if type_of == "bool" || type_of == "*bool" {
		return true
	}

	return false
}

func (m Map) IsArray(s string) bool {
	type_of := m.GetType(s)
	if type_of == "class.Array" || type_of == "*class.Array" {
		return true
	}

	return false
}

func (m Map) IsEmptyString(s string) bool {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "string" || type_of == "*string" {
		string_value, _ := m.GetString(s)
		return *string_value == ""
	}

	return false
}

func (m Map) IsNumber(s string) bool {
	type_of := m.GetType(s)
	switch type_of {
	case "*int", "*int64", "*uint64","int", "int64", "uint64":
		return true
	default: 
		return false
	}
}

func (m Map) IsString(s string) bool {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "string" || type_of == "*string" {
		return true
	}

	return false
}

func (m Map) IsBoolTrue(s string) bool {
	if m.IsNil(s) {
		return false
	}

	if !m.IsBool(s) {
		return false
	}

	value, _ := m.GetBool(s)
	return *value == true
}

func (m Map) IsBoolFalse(s string) bool {
	if m.IsNil(s) {
		return true
	}

	if !m.IsBool(s) {
		return true
	}

	value, _ := m.GetBool(s)
	return *value == false
}

func (m Map) ToJSONString() (*string, []error) {
	var errors []error
	keys := m.Keys()
	length := len(keys)
	
	json := ""
	if length == 0 {
		json = "{}"
		return &json, nil
	}

	json += "{\n"
	for i, key := range keys {
		json = json + "\"" + strings.ReplaceAll(key, "\"", "\\\"") + "\":"
		string_conversion, string_conversion_errors := ConvertInterfaceValueToJSONStringValue(m[key])
		if string_conversion_errors != nil {
			errors = append(errors, string_conversion_errors...)
		} else {
			json += *string_conversion
		}

		if i < length - 1 {
			json += ","
		}
		json += "\n"
	}
	json += "}"

	if len(errors) > 0 {
		return nil, errors
	}
	return &json, nil
}

/*
func (m Map) A(s string) Array {
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
}*/

func (m Map) SetArray(s string, array *Array) {
	if array == nil {
		m[s] = nil
		return
	}

	rep := fmt.Sprintf("%T", array)
	switch rep {
	case "*class.Array":
		m[s] = array
	default:
		panic(fmt.Errorf("Map.SetArray: type %s is not supported please implement for field: %s", rep, s))
	}
}

func (m Map) SetErrors(s string, errors *[]error) {
	if errors == nil {
		m[s] = nil
		return
	}

	m[s] = errors
}

func (m Map) GetType(s string) string {
	if m.IsNil(s) {
		return "nil"
	}
	return fmt.Sprintf("%T", m[s])
}

func (m Map) Func(s string) func(Map) []error {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "func(class.Map) []error":
		return m[s].(func(Map) []error)
	case "*func(class.Map) []error":
		value := m[s].(*func(Map) []error)
		return *value
	default:
		panic(fmt.Errorf("Map.Func: type %s is not supported please implement for field: %s", rep, s))
	}

	return nil
}

func (m Map) SetFunc(s string, function func(Map) []error) {
	m[s] = function
}

func (m Map) GetArray(s string) (*Array, []error) {
	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	var errors []error
	var array *Array

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*class.Array":
		cloned_array, cloned_array_errors := m[s].(*Array).Clone()
		if cloned_array_errors != nil {
			errors = append(errors, cloned_array_errors...)
		} else {
			array = cloned_array
		}
	case "class.Array":
		cloned_array, cloned_array_errors := m[s].(Array).Clone()
		if cloned_array_errors != nil {
			errors = append(errors, cloned_array_errors...)
		} else {
			array = cloned_array
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetArray: type %s is not supported please implement for field: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return array, nil
}

func (m Map) GetString(s string) (*string, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "string":
		value := m[s].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("Map.GetString: *string value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetString: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetFloat64(s string) (*float64, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *float64
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "float64":
		value := m[s].(float64)
		result = &value
	case "*float64":
		if fmt.Sprintf("%s", m[s]) != "%!s(*float64=<nil>)" {
			value := m[s].(*float64)
			new_value := *value
			result = &new_value
		} else {
			errors = append(errors, fmt.Errorf("Map.GetFloat64: *float64 value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetFloat64: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetFloat32(s string) (*float32, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *float32
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "float32":
		value := m[s].(float32)
		result = &value
	case "*float32":
		if fmt.Sprintf("%s", m[s]) != "%!s(*float32=<nil>)" {
			value := m[s].(*float32)
			new_value := *value
			result = &new_value
		} else {
			errors = append(errors, fmt.Errorf("Map.GetFloat32: *float32 value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetFloat32: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetRunes(s string) (*[]rune, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "string":
		value := m[s].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("Map.GetString: *string value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetString: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var runes []rune
	for _, runeValue := range *result {
		runes = append(runes, runeValue)
	}

	return &runes, nil
}

func (m Map) GetObject(s string) interface{} {
	return m[s]
}

func (m Map) GetBool(s string) (*bool, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var result *bool
	var errors []error

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "bool":
		value := m[s].(bool)
		result = &value
		break
	case "*bool":
		if fmt.Sprintf("%s", m[s]) != "%!s(*bool=<nil>)" {
			value := *((m[s]).(*bool))
			result = &value
		} else {
			return nil, nil
		}
		break
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			value := *((m[s]).(*string))
			if value == "1" {
				boolean_result := true
				result = &boolean_result
			} else if value == "0" {
				boolean_result := false
				result = &boolean_result
			} else {
				errors = append(errors, fmt.Errorf("Map.GetBool: unknown value for *string: %s", value))
				result = nil
			}
		} else {
			return nil, nil
		}
		break
	case "string":
		value := ((m[s]).(string))
		if value == "1" {
			boolean_result := true
			result = &boolean_result
		} else if value == "0" {
			boolean_result := false
			result = &boolean_result
		} else {
			errors = append(errors, fmt.Errorf("Map.GetBool: unknown value for string: %s", value))
			result = nil
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetBool: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetBool(s string, value *bool) {
	m[s] = value
}

func (m Map) SetString(s string, value *string) {
	if value == nil {
		m[s] = nil
		return 
	}

	clone_string := CloneString(value)
	m[s] = clone_string
}

func (m Map) SetNil(s string) {
	m[s] = nil
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
	var temp_value int64

	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
		case "*int64":
			x := m[s].(*int64)
			temp_value = int64(*x)
		case "int64":
			x := m[s].(int64)
			temp_value = x
		case "*int32":
			x := m[s].(*int32)
			temp_value = int64(*x)
		case "int32":
			x := m[s].(int32)
			temp_value = int64(x)
		case "*int16":
			x := m[s].(*int16)
			temp_value = int64(*x)
		case "int16":
			x := m[s].(int16)
			temp_value = int64(x)
		case "*int8":
			x := m[s].(*int8)
			temp_value = int64(*x)
		case "int8":
			x := m[s].(int8)
			temp_value = int64(x)
		case "int":
			x := m[s].(int)
			temp_value = int64(x)
		case "*int":
			x := m[s].(*int)
			temp_value = int64(*x)
		case "*string":
			value, value_error := strconv.ParseInt((*(m[s].(*string))), 10, 64)
			if value_error != nil {
				errors = append(errors, fmt.Errorf("Map.GetInt64: cannot convert *string value to int64"))
			} else {
				temp_value = value
			}
		default:
			errors = append(errors, fmt.Errorf("Map.GetInt64: type %s is not supported please implement", rep))
	}
	
	if len(errors) > 0 {
		return nil, errors
	}

	result := &temp_value

	return result, nil
}

func (m Map) GetInt8(s string) (*int8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -128 || *int64_value > 127 {
		errors = append(errors, fmt.Errorf("value is not in range [-128, 127]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int8_conv := int8(*int64_value)
	result := &int8_conv

	return result, nil
}

func (m Map) GetUInt8(s string) (*uint8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetUInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 255 {
		errors = append(errors, fmt.Errorf("value is not in range [0, 255]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int8_conv := uint8(*int64_value)
	result := &int8_conv

	return result, nil
}

func (m Map) GetInt16(s string) (*int16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -32768 || *int64_value > 32767 {
		errors = append(errors, fmt.Errorf("value is not in range [-32768, 32767]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int16_conv := int16(*int64_value)
	result := &int16_conv

	return result, nil
}

func (m Map) GetUInt16(s string) (*uint16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetUInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 65535 {
		errors = append(errors, fmt.Errorf("value is not in range [0, 65535]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int16_conv := uint16(*int64_value)
	result := &int16_conv

	return result, nil
}


func (m Map) GetInt32(s string) (*int32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -2147483648 || *int64_value > 2147483647 {
		errors = append(errors, fmt.Errorf("value is not in range [-2147483648, 2147483647]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int32_conv := int32(*int64_value)
	result := &int32_conv

	return result, nil
}

func (m Map) GetUInt32(s string) (*uint32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf(" m.GetUInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 4294967295 {
		errors = append(errors, fmt.Errorf("value is not in range [0, 4294967295]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int32_conv := uint32(*int64_value)
	result := &int32_conv

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
			errors = append(errors, fmt.Errorf("Map.GetInt: cannot convert *string value to int"))
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

func (m Map) SetInt(s string, v *int) {
	m[s] = v
}

func (m Map) SetInt64(s string, v *int64) {
	m[s] = v
}

func (m Map) SetInt8(s string, v *int8) {
	m[s] = v
}

func (m Map) SetInt16(s string, v *int16) {
	m[s] = v
}

func (m Map) SetInt32(s string, v *int32) {
	m[s] = v
}

func (m Map) SetFloat64(s string, v *float64) {
	m[s] = v
}

func (m Map) SetFloat32(s string, v *float32) {
	m[s] = v
}

func (m Map) GetUInt64(s string) (*uint64, []error) {
	var errors []error
	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	var uint64_value uint64
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*int64":
		x := *(m[s].(*int64))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int64":
		x := m[s].(int64)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int32":
		x := *(m[s].(*int32))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int32":
		x := m[s].(int32)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int16":
		x := *(m[s].(*int16))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int16":
		x := m[s].(int16)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int8":
		x := *(m[s].(*int8))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int8":
		x := m[s].(int8)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		x := (m[s].(int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		x := *(m[s].(*int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*uint64":
		uint64_value = *(m[s].(*uint64))
	case "uint64":
		uint64_value = (m[s].(uint64))
	case "*uint32":
		x := *(m[s].(*uint32))
		uint64_value = uint64(x)
	case "uint32":
		x := (m[s].(uint32))
		uint64_value = uint64(x)
	case "*uint16":
		x := *(m[s].(*uint16))
		uint64_value = uint64(x)
	case "uint16":
		x := (m[s].(uint16))
		uint64_value = uint64(x)
	case "*uint8":
		x := *(m[s].(*uint8))
		uint64_value = uint64(x)
	case "uint8":
		x := (m[s].(uint8))
		uint64_value = uint64(x)
	case "*string":
		string_value := (m[s].(*string))
		if *string_value == "NULL" {
			return nil, nil
		} else {
			value, value_error := strconv.ParseUint(*string_value, 10, 64)
			if value_error != nil {
				errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert *string value to uint64"))
			} else {
				uint64_value = value
			}
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetUInt64: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &uint64_value, nil
}

func (m Map) SetUInt64(s string, v *uint64) {
	m[s] = v
}

func (m Map) SetUInt32(s string, v *uint32) {
	m[s] = v
}

func (m Map) SetUInt16(s string, v *uint16) {
	m[s] = v
}

func (m Map) SetUInt8(s string, v *uint8) {
	m[s] = v
}

func (m Map) SetUInt64Value(s string, v uint64) {
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
		value := m[s].(time.Time)
		result = &value
	case "*string":
		//todo: parse for null
		value1, value_errors1 := time.Parse("2006-01-02 15:04:05.000000", *(m[s].(*string)))
		value2, value_errors2 := time.Parse("2006-01-02 15:04:05", *(m[s].(*string)))

		if value_errors1 != nil && value_errors2 != nil {
			errors = append(errors, value_errors1)
		}

		if value_errors1 == nil {
			result = &value1
		}

		if value_errors2 == nil {
			result = &value2
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

func (m Map) Clone() (*Map, []error) {
	var errors []error
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
			x, _ := m.GetString(key)
			clone[key] = *(CloneString(x))
			break
		case "*string":
			if fmt.Sprintf("%s", m[key]) == "%!s(*string=<nil>)" {
				clone[key] = nil
			} else {
				x, _ := m.GetString(key)
				clone.SetString(key, CloneString(x))
			}
			break
		case "class.Map":
			cloned_map, cloned_map_error := current.(Map).Clone()
			if cloned_map_error != nil {
				errors = append(errors, cloned_map_error...)
			} else {
				clone[key] = *cloned_map
			}
			break
		case "*class.Map":
			cloned_map, cloned_map_error := current.(*Map).Clone()
			if cloned_map_error != nil {
				errors = append(errors, cloned_map_error...)
			} else {
				clone[key] = *cloned_map
			}
			break
		case "class.Array":
			cloned_array, cloned_array_errors := current.(Array).Clone()
			if cloned_array_errors != nil {
				errors = append(errors, cloned_array_errors...)
			} else {
				clone[key] = cloned_array
			}
			break
		case "*class.Array":
			cloned_array, cloned_array_errors := current.(*Array).Clone()
			if cloned_array_errors != nil {
				errors = append(errors, cloned_array_errors...)
			} else {
				clone[key] = cloned_array
			}
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
		case "*int":
			int_value := *(current.(*int))
			clone[key] = int_value
		case "int8":
			clone[key] = current.(int8)
		case "*int8":
			int8_value := *(current.(*int8))
			clone[key] = &int8_value
		case "int16":
			clone[key] = current.(int16)
		case "*int16":
			int16_value := *(current.(*int16))
			clone[key] = &int16_value
		case "int32":
			clone[key] = current.(int32)
		case "*int32":
			int32_value := *(current.(*int32))
			clone[key] = &int32_value
		case "int64":
			clone[key] = current.(int64)
		case "*int64":
			int64_value := *(current.(*int64))
			clone[key] = &int64_value
		case "uint":
			clone[key] = current.(uint)
		case "*uint":
			uint_value := *(current.(*uint))
			clone[key] = &uint_value
		case "uint8":
			clone[key] = current.(uint8)
		case "*uint8":
			uint8_value := *(current.(*uint8))
			clone[key] = &uint8_value
		case "uint16":
			clone[key] = current.(uint16)
		case "*uint16":
			uint16_value := *(current.(*uint16))
			clone[key] = &uint16_value
		case "uint32":
			clone[key] = current.(uint32)
		case "*uint32":
			uint32_value := *(current.(*uint32))
			clone[key] = &uint32_value
		case "uint64":
			clone[key] = current.(uint64)
		case "*uint64":
			uint64_value := *(current.(*uint64))
			clone[key] = &uint64_value
		case "float32":
			clone[key] = current.(float32)
		case "*float32":
			float32_value := *(current.(*float32))
			clone[key] = &float32_value
		case "float64":
			clone[key] = current.(float64)
		case "*float64":
			float64_value := *(current.(*float64))
			clone[key] = &float64_value
		case "*[]error":
			errs := *(current.(*[]error))
			clone[key] = &errs
		case "[]error":
			clone[key] = current.([]error)
		case "<nil>":
			clone[key] = nil
		default:
			errors = append(errors, fmt.Errorf("Map.Clone: type %s is not supported please implement", rep))
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &clone, nil
}
