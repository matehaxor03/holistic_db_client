package class

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Map map[string]interface{}

func ParseJSON(s string) (*Map, []error) {
	var errors []error
	if s == "" {
		errors = append(errors, fmt.Errorf("value empty string"))
	}

	if !strings.HasPrefix(s, "{") {
		errors = append(errors, fmt.Errorf("json does not start with {"))
	}

	if !strings.HasSuffix(s, "}") {
		errors = append(errors, fmt.Errorf("json does not end with }"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	runes := []rune(s)
	metrics := Map{"{":0, "}":0}
	mode := "looking_for_keys"
	parent_map := Map{}
	// parent map array and current map array etc
	result, result_error :=  parseJSONMap(&runes, &mode, &parent_map, nil, &metrics)

	opening_bracket_count, opening_bracket_count_errors := metrics.GetInt("{")
	closing_bracket_count, closing_bracket_count_errors := metrics.GetInt("}")

	if opening_bracket_count_errors != nil {
		errors = append(errors, opening_bracket_count_errors...)
	}

	if closing_bracket_count_errors != nil {
		errors = append(errors, closing_bracket_count_errors...)
	}

	if result_error != nil {
		errors = append(errors, result_error...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *opening_bracket_count != *closing_bracket_count {
		errors = append(errors, fmt.Errorf("opening and closing brackets do not match, opening: %d closing: %d", *opening_bracket_count, *closing_bracket_count))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}


func parseJSONMap(runes *[]rune, mode *string, data_map *Map, data_array *Array, metrics *Map) (*Map, []error) {
	var errors []error
	if data_map == nil && data_array == nil {
		errors = append(errors, fmt.Errorf("parent map or array cannot both be nil"))
	}

	if data_map != nil && data_array != nil {
		errors = append(errors, fmt.Errorf("parent map or array cannot both not be nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	
	mode_looking_for_keys := "looking_for_keys"
	mode_looking_for_key_name := "looking_for_key_name"
	mode_looking_for_key_name_column := "looking_for_key_name_column"
	mode_looking_for_value := "looking_for_value"
	mode_unknown := "unknown"
	
	temp_key := ""
	temp_value := ""
	temp_mode := CloneString(mode)
	current_mode := *temp_mode

	for i, value := range *runes {
		if string(value) == "{" {
			opening_count, _ := metrics.GetInt("{")
			*opening_count++
			metrics.SetInt("{", opening_count)
		}

		if string(value) == "}" {
			closing_count, _ := metrics.GetInt("}")
			*closing_count++
			metrics.SetInt("}", closing_count)
		}

		if current_mode == mode_unknown {
			if string(value) == "\n" {

			} else if string(value) == "{" {
				new_mode := mode_looking_for_keys
				new_s := string((*runes)[i+1:])
				new_runes := []rune(new_s)
				new_map := Map{}

				if data_map != nil {
					data_map.SetMap(temp_key, &new_map)	
				} else if data_array != nil {
					*data_array = append(*data_array, &data_map)
				}
				return parseJSONMap(&new_runes, &new_mode, &new_map, nil, metrics)
			} else if string(value) == "[" {
				new_mode := mode_looking_for_value
				new_s := string((*runes)[i+1:])
				new_runes := []rune(new_s)
				new_array := Array{}

				if data_map != nil {
					data_map.SetArray(temp_key, &new_array)	
				} else if data_array != nil {
					*data_array = append(*data_array, &new_array)
				}
				return parseJSONMap(&new_runes, &new_mode, nil, &new_array, metrics)
			} else if string(value) == "," {
				if data_map != nil {
					current_mode = mode_looking_for_keys
				} else if data_array != nil {
					current_mode = mode_looking_for_value
				}
			}
		} else if current_mode == mode_looking_for_keys {
			if string(value) == "\"" {
				current_mode = mode_looking_for_key_name
			}
		} else if current_mode == mode_looking_for_key_name {
			if string(value) == "\"" {
				current_mode = mode_looking_for_key_name_column
			} else {
				temp_key += string(value)
			}
		} else if current_mode == mode_looking_for_key_name_column {
			if string(value) == ":" {
				current_mode = mode_looking_for_value
			}
		} else if current_mode == mode_looking_for_value {
			if string(value) == "{" {
				new_mode := mode_looking_for_keys
				new_s := string((*runes)[i+1:])
				new_runes := []rune(new_s)
				new_map := Map{}
				data_map.SetMap(temp_key, &new_map)
				
				return parseJSONMap(&new_runes, &new_mode, &new_map, nil, metrics)
			} else if string(value) == "[" {
				new_mode := mode_looking_for_value
				new_s := string((*runes)[i+1:])
				new_runes := []rune(new_s)
				new_array := Array{}
				data_map.SetArray(temp_key, &new_array)
				return parseJSONMap(&new_runes, &new_mode, nil, &new_array, metrics)
			} else if string(value) == "}" {
				parse_errors := parseJSONValue(temp_key, temp_value, data_map, data_array)
				if parse_errors != nil {
					errors = append(errors, parse_errors...)
				}

				temp_key = ""
				temp_value = ""

				current_mode = mode_unknown
			} else if string(value) == "]" {
				parse_errors := parseJSONValue(temp_key, temp_value, data_map, data_array)
				if parse_errors != nil {
					errors = append(errors, parse_errors...)
				}

				temp_key = ""
				temp_value = ""

				current_mode = mode_unknown
			} else if string(value) == "," {
				parse_errors := parseJSONValue(temp_key, temp_value, data_map, data_array)
				if parse_errors != nil {
					errors = append(errors, parse_errors...)
				}
				
				temp_key = ""
				temp_value = ""

				if data_map != nil {
					current_mode = mode_looking_for_keys
				} else if data_array != nil {
					current_mode = mode_looking_for_value
				}
			} else {
				temp_value += string(value)
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}
	}

	return data_map, nil
}

func parseJSONValue(temp_key string, temp_value string, data_map *Map, data_array *Array) []error {
	var errors []error

	if data_map == nil && data_array == nil {
		errors = append(errors, fmt.Errorf("map or array cannot both be nil"))
	}

	if data_map != nil && data_array != nil {
		errors = append(errors, fmt.Errorf("map or array cannot both not be nil"))
	}

	if len(errors) > 0 {
		return errors
	}
	
	data_type := ""
	string_value := CloneString(&temp_value)
	
	var boolean_value *bool
	var float64_value *float64
	var float32_value *float32
	var int64_value *int64
	var uint64_value *uint64
	if strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		data_type = "string"
		dequoted_value := (*string_value)[1:(len(*string_value)-1)]
		string_value = &dequoted_value	
	} else if strings.HasPrefix(*string_value, "\"") && !strings.HasSuffix(*string_value, "\"") {
		errors = append(errors, fmt.Errorf("value has \" as prefix but not \" as suffix"))
	} else if !strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		errors = append(errors, fmt.Errorf("value has \" as suffix but not \" as prefix"))
	} else if *string_value == "true" {
		data_type = "bool"
		boolean_value_true := true 
		boolean_value = &boolean_value_true
	} else if *string_value == "false" {
		data_type = "bool"
		boolean_value_false := false 
		boolean_value = &boolean_value_false
	} else if *string_value == "null" {
		data_type = "null"
	} else {
		var negative_number bool
		negative_number_count := strings.Count(*string_value, "-")
		if negative_number_count == 1 {
			negative_number = true
			if !strings.HasPrefix(*string_value, "-") {
				errors = append(errors, fmt.Errorf("negative symbol is not at the start of the number"))
			}
		} else if negative_number_count == 0 {
			negative_number = false
		} else {
			errors = append(errors, fmt.Errorf("value contained %d - characters", negative_number_count))
		}

		var decimal_number bool
		decimal_count := strings.Count(*string_value, ".")
		if decimal_count == 1 {
			decimal_number = true
		} else if decimal_count == 0 {
			decimal_number = false
		} else {
			errors = append(errors, fmt.Errorf("value contained %d decimal points", decimal_count))
		}

		whitelist_characters := Map{"0":nil,"1":nil,"2":nil,"3":nil,"4":nil,"5":nil,"6":nil,"7":nil,"8":nil,"9":nil,".":nil,"-":nil}
		parameters := Map{"values":&whitelist_characters,"value":string_value,"label":"parseJSONValue","data_type":"number"}
		whitelist_chracter_errors := WhitelistCharacters(parameters)
		if whitelist_chracter_errors != nil {
			errors = append(errors, whitelist_chracter_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		if decimal_number {
			data_type = "float64"
			float64_temp, float64_temp_error := strconv.ParseFloat(*string_value, 64)
			if float64_temp_error != nil {
				errors = append(errors, fmt.Errorf("strconv.ParseFloat(*string_value, 64) error"))
			} else {
				float64_value = &float64_temp
			}

			if len(errors) > 0 {
				return errors
			}

			float32_temp, float32_temp_error := strconv.ParseFloat(*string_value, 32)
			if float32_temp_error != nil {
			} else {
				data_type = "float32"
				float32_conv := float32(float32_temp)
				float32_value = &float32_conv
			}


		} else {
			if negative_number {
				data_type = "int64"
				int64_temp, int64_temp_error := strconv.ParseInt(*string_value, 10, 64)
				if int64_temp_error != nil {
					errors = append(errors, fmt.Errorf("strconv.ParseInt(*string_value, 10, 64) error"))
				} else {
					int64_value = &int64_temp
				}

				if len(errors) > 0 {
					return errors
				}

			} else {
				data_type = "uint64"
				uint64_temp, uint64_temp_error := strconv.ParseUint(*string_value, 10, 64)
				if uint64_temp_error != nil {
					errors = append(errors, fmt.Errorf("strconv.ParseUint(*string_value, 10, 64) error"))
				} else {
					uint64_value = &uint64_temp
				}

				if len(errors) > 0 {
					return errors
				}

			}
		}

		if len(errors) > 0 {
			return errors
		}
	}

	if data_type == "" {
		errors = append(errors, fmt.Errorf("data_type is unknown please implement"))
	}


	if len(errors) > 0 {
		return errors
	}

	if data_array != nil {
		if data_type == "string" {
			*data_array = append(*data_array, string_value)
		} else if data_type == "bool" {
			*data_array = append(*data_array, boolean_value)
		} else if data_type == "null" {
			*data_array = append(*data_array, nil)
		} else if data_type == "float64" {
			*data_array = append(*data_array, float64_value)
		} else if data_type == "float32" {
			*data_array = append(*data_array, float32_value)
		}  else if data_type == "int64" {
			*data_array = append(*data_array, int64_value)
		} else if data_type == "uint64" {
			*data_array = append(*data_array, uint64_value)
		}
	}

	if data_map != nil {
		if data_type == "string" {
			(*data_map).SetString(temp_key, string_value)
		} else if data_type == "bool" {
			(*data_map).SetBool(temp_key, boolean_value)
		} else if data_type == "null" {
			(*data_map).SetNil(temp_key)
		} else if data_type == "float64" {
			(*data_map).SetFloat64(temp_key, float64_value)
		} else if data_type == "float32" {
			(*data_map).SetFloat32(temp_key, float32_value)
		} else if data_type == "int64" {
			(*data_map).SetInt64(temp_key, int64_value)
		} else if data_type == "uint64" {
			(*data_map).SetUInt64(temp_key, uint64_value)
		}
	}

	if len(errors) > 0 {
		return errors
	}


	return nil
}

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
				x, x_error := value.(Map).ToJSONString()
				if x_error != nil {
					errors = append(errors, x_error...)
				} else {
					json += *x
				}
			case "class.Array":
				x, x_error := value.(Array).ToJSONString()
				if x_error != nil {
					errors = append(errors, x_error...)
				} else {
					json += *x
				}
			case "*class.Array":
				x, x_error := (*(value.(*Array))).ToJSONString()
				if x_error != nil {
					errors = append(errors, x_error...)
				} else {
					json += *x
				}
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
			case "[]error":
				json += "["
				array_length := len(m[key].([]error))
				for array_index, array_value := range m[key].([]error) {
					json += "\"" + array_value.Error() + "\""
					if array_index < array_length {
						json += ","
					}
				}
				json += "]"
			case "*[]error":
				json += "["
				array_length := len(*(m[key].(*[]error)))
				for array_index, array_value := range *(m[key].(*[]error)) {
					json += "\"" + array_value.Error() + "\""
					if array_index < array_length {
						json += ","
					}
				}
				json += "]"
			case "func(string, *string, string, string) []error", "func(class.Map) []error", "*func(class.Map) []error":
				json = json + fmt.Sprintf("\"%s\"", rep)
			case "*class.Host":
				x, x_errors := (*(value.(*Host))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*class.Credentials":
				x, x_errors := (*(value.(*Credentials))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*class.DatabaseCreateOptions":
				x, x_errors := (*(value.(*DatabaseCreateOptions))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*class.Database":
				x, x_errors := (*(value.(*Database))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*class.Client":
				x, x_errors := (*(value.(*Client))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*class.Table":
				x, x_errors := (*(value.(*Table))).ToJSONString()
				if x_errors != nil {
					errors = append(errors, x_errors...)
				} else {
					json += *x
				}
			case "*time.Time":
				json += "\"" + (*(value.(*time.Time))).Format("2006-01-02 15:04:05.000000") + "\""
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
			case "*float64":
				json = json + fmt.Sprintf("%f", *(value.(*float64)))
			case "float64":
				json = json + fmt.Sprintf("%f", (value.(float64)))
			case "*float32":
				json = json + fmt.Sprintf("%f", *(value.(*float32)))
			case "float32":
				json = json + fmt.Sprintf("%f", (value.(float32)))
			default:
				errors = append(errors, fmt.Errorf("Map.ToJSONString: type %s is not supported please implement for %s", rep, key))
			}
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
}

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

func (m Map) Array(s string) []interface{} {
	return m[s].([]interface{})
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
			errors = append(errors, fmt.Errorf("Map.GetInt64: cannot convert *string value to int64"))
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

func (m Map) SetFloat64(s string, v *float64) {
	m[s] = v
}

func (m Map) SetFloat32(s string, v *float32) {
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
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		value := (m[s].(int))
		if value >= 0 {
			temp := uint64(value)
			result = &temp
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		value := *(m[s].(*int))
		if value >= 0 {
			temp := uint64(value)
			result = &temp
		} else {
			errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*uint64":
		value := *(m[s].(*uint64))
		result = &value
	case "uint64":
		value := (m[s].(uint64))
		result = &value
	case "*string":
		string_value := (m[s].(*string))
		if string_value == nil || *string_value == "NULL" {
			result = nil
		} else {
			value, value_error := strconv.ParseUint(*string_value, 10, 64)
			if value_error != nil {
				errors = append(errors, fmt.Errorf("Map.GetUInt64: cannot convert *string value to uint64"))
			} else {
				result = &value
			}
		}
	default:
		errors = append(errors, fmt.Errorf("Map.GetUInt64: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetUInt64(s string, v *uint64) {
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
		case "*[]error":
			errs := *(current.(*[]error))
			clone[key] = &errs
		case "[]error":
			clone[key] = current.([]error)
		case "<nil>":
			clone[key] = nil
		default:
			panic(fmt.Errorf("Map.Clone: type %s is not supported please implement", rep))
		}
	}

	return clone
}
