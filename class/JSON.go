package class

import (
	"fmt"
	"strconv"
	"strings"
	"container/list"
	"net/url"
	"time"
)

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
	metrics := Map{"{":0, "}":0, "[":0, "]":0, "opening_quote":0,"closing_quote":0}
	mode := "looking_for_keys"
	parent_map := Map{}
	list := list.New()
	list.PushFront(&parent_map)
	index := uint64(0)
	// parent map array and current map array etc
	result_error :=  parseJSONMap(&runes, &index, &mode, list, &metrics)

	opening_bracket_count, opening_bracket_count_errors := metrics.GetInt("{")
	closing_bracket_count, closing_bracket_count_errors := metrics.GetInt("}")
	opening_square_count, opening_square_count_errors := metrics.GetInt("[")
	closing_square_count, closing_square_count_errors := metrics.GetInt("]")
	opening_quote_count, opening_quote_count_errors := metrics.GetInt("opening_quote")
	closing_quote_count, closing_quote_count_errors := metrics.GetInt("closing_quote")


	if opening_bracket_count_errors != nil {
		errors = append(errors, opening_bracket_count_errors...)
	}

	if closing_bracket_count_errors != nil {
		errors = append(errors, closing_bracket_count_errors...)
	}

	if opening_square_count_errors != nil {
		errors = append(errors, opening_square_count_errors...)
	}

	if closing_square_count_errors != nil {
		errors = append(errors, closing_square_count_errors...)
	}

	if opening_quote_count_errors != nil {
		errors = append(errors, opening_quote_count_errors...)
	}

	if closing_quote_count_errors != nil {
		errors = append(errors, closing_quote_count_errors...)
	}

	if result_error != nil {
		errors = append(errors, result_error...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *opening_bracket_count != *closing_bracket_count {
		errors = append(errors, fmt.Errorf("opening and closing brackets {} do not match, opening: %d closing: %d", *opening_bracket_count, *closing_bracket_count))
	}

	if *opening_square_count != *closing_square_count {
		errors = append(errors, fmt.Errorf("opening and closing squares [] do not match, opening: %d closing: %d", *opening_square_count, *closing_square_count))
	}

	if *opening_quote_count != *closing_quote_count {
		errors = append(errors, fmt.Errorf("opening and closing quotes do not match, opening: %d closing: %d", *opening_quote_count, *closing_quote_count))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &parent_map, nil
}


func parseJSONMap(runes *[]rune, index *uint64, mode *string, list *list.List, metrics *Map) ([]error) {
	var errors []error
	if list == nil {
		errors = append(errors, fmt.Errorf("list is nil"))
	} else if (*list).Len() == 0 {
		errors = append(errors, fmt.Errorf("list is empty"))
	}

	if len(errors) > 0 {
		return errors
	}
	
	
	mode_looking_for_keys := "looking_for_keys"
	mode_looking_for_key_name := "looking_for_key_name"
	mode_looking_for_key_name_column := "looking_for_key_name_column"
	mode_looking_for_value := "looking_for_value"
	
	temp_key := ""
	temp_value := ""
	parsing_string := false
	found_value := false
	temp_mode := CloneString(mode)
	current_mode := *temp_mode

	for *index < uint64(len(*runes)) {
		i := *index
		value := (*runes)[*index]
		
		//fmt.Println(current_mode + " " + string(value))
		//fmt.Println(temp_key + " " + temp_value + " parsing:" + fmt.Sprintf("%s", parsing_string) + "found value: " +  fmt.Sprintf("%s", found_value))

		if !parsing_string {
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

			if string(value) == "[" {
				opening_count, _ := metrics.GetInt("[")
				*opening_count++
				metrics.SetInt("[", opening_count)
			}

			if string(value) == "]" {
				closing_count, _ := metrics.GetInt("]")
				*closing_count++
				metrics.SetInt("]", closing_count)
			}
		}
		

		if current_mode == mode_looking_for_keys {
			if string(value) == "\"" && string((*runes)[i-1]) != "\\"{
				current_mode = mode_looking_for_key_name
				parsing_string = true

				opening_count, _ := metrics.GetInt("opening_quote")
				*opening_count++
				metrics.SetInt("opening_quote", opening_count)
			} else if string(value) == "}" {
				if list.Len() > 1 {
					list.Remove(list.Front())
					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					}
				}	
			} else if string(value) == "]" {
				// this should not occur for valid json todo throw error
				/*if list.Len() > 1 {
					list.Remove(list.Front())
					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					}
				}*/
			}
		} else if current_mode == mode_looking_for_key_name {
			if string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				current_mode = mode_looking_for_key_name_column
				parsing_string = false

				opening_count, _ := metrics.GetInt("closing_quote")
				*opening_count++
				metrics.SetInt("closing_quote", opening_count)
			} else {
				temp_key += string(value)
			}
		} else if current_mode == mode_looking_for_key_name_column {
			if string(value) == ":" {
				current_mode = mode_looking_for_value
			}
		} else if current_mode == mode_looking_for_value {
			
			if !found_value && (string(value) == " " || string(value) == "\r" || string(value) == "\n" || string(value) == "\t") {
				*index++
				continue
			} else {
				found_value = true
			}

			if !parsing_string {
				if i > 0 && !parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
					temp_value += string(value)
					parsing_string = true
					
					opening_count, _ := metrics.GetInt("opening_quote")
					*opening_count++
					metrics.SetInt("opening_quote", opening_count)

				} else if i == 0 && !parsing_string && string(value) == "\"" {
					temp_value += string(value)
					parsing_string = true
					
					opening_count, _ := metrics.GetInt("opening_quote")
					*opening_count++
					metrics.SetInt("opening_quote", opening_count)

				} else if string(value) == "{" {
					new_mode := mode_looking_for_keys
					new_map := Map{}

					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						list.Front().Value.(*Map).SetMap(temp_key, &new_map)	
					} else {
						temp_array := list.Front().Value.(*Array)
						*temp_array = append(*temp_array, &new_map)
					} 
					list.PushFront(&new_map)


					found_value = false
					*index++
					parseJSONMap(runes, index, &new_mode, list, metrics)
				} else if string(value) == "[" {
					new_mode := mode_looking_for_value
					new_array := Array{}

					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						list.Front().Value.(*Map).SetArray(temp_key, &new_array)	
					} else {
						temp_array := list.Front().Value.(*Array)
						*temp_array = append(*temp_array, &new_array)
					} 
					list.PushFront(&new_array)

					
					found_value = false
					*index++
					parseJSONMap(runes, index, &new_mode, list, metrics)
				} else if string(value) == "}" {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}

					temp_key = ""
					temp_value = ""

					found_value = false
					list.Remove(list.Front())

					if list.Len() > 0 {
						if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
							current_mode = mode_looking_for_keys
						} else {
							current_mode = mode_looking_for_value
						}
					} 

				} else if string(value) == "]" {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}

					temp_key = ""
					temp_value = ""

					found_value = false
					list.Remove(list.Front())

					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					}
					
				} else if string(value) == "," {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}
					
					temp_key = ""
					temp_value = ""

					if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					}
		
					found_value = false
				} else {
					temp_value += string(value)
				}
			} else if !parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				temp_value += string(value)
				parsing_string = true

				opening_count, _ := metrics.GetInt("opening_quote")
				*opening_count++
				metrics.SetInt("opening_quote", opening_count)

			} else if parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				temp_value += string(value)
				parsing_string = false

				opening_count, _ := metrics.GetInt("closing_quote")
				*opening_count++
				metrics.SetInt("closing_quote", opening_count)
			} else {
				temp_value += string(value)
			}

			
		}

		if len(errors) > 0 {
			return errors
		}

		*index++
	}

	return nil
}

func parseJSONValue(temp_key string, temp_value string, list *list.List) []error {
	var errors []error
	
	if list == nil {
		errors = append(errors, fmt.Errorf("list is nil"))
	} else if (*list).Len() == 0 {
		errors = append(errors, fmt.Errorf("list is empty"))
	}

	if len(errors) > 0 {
		return errors
	}

	temp_key = strings.ReplaceAll(temp_key, "\\\"", "\"")
	temp_value = strings.ReplaceAll(temp_value, "\\\"", "\"")
	
	data_type := ""
	string_value := CloneString(&temp_value)
	
	var boolean_value *bool
	var float64_value *float64
	var float32_value *float32
	var int64_value *int64
	var int32_value *int32
	var int16_value *int16
	var int8_value *int8
	var uint64_value *uint64
	var uint32_value *uint32
	var uint16_value *uint16
	var uint8_value *uint8

	//if strings.HasPrefix(*string_value, "\"") || strings.HasSuffix(*string_value, "\"") {
	string_temp := strings.TrimSpace(*string_value)
	string_value = &string_temp
	//}


	if strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		data_type = "string"
		dequoted_value := (*string_value)[1:(len(*string_value)-1)]
		string_value = &dequoted_value	
	} else if strings.HasPrefix(*string_value, "\"") && !strings.HasSuffix(*string_value, "\"") {
		errors = append(errors, fmt.Errorf("value has \" as prefix but not \" as suffix"))
	} else if !strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		errors = append(errors, fmt.Errorf("value has \" as suffix but not \" as prefix"))
	} else {
		// when parsing emtpy array []
		if *string_value == "" {
			return nil
		}

		if *string_value == "true" {
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
				errors = append(errors, fmt.Errorf("value contained %d negative symbols expected 1", negative_number_count))
			}

			var decimal_number bool
			decimal_count := strings.Count(*string_value, ".")
			if decimal_count == 1 {
				decimal_number = true
			} else if decimal_count == 0 {
				decimal_number = false
			} else {
				errors = append(errors, fmt.Errorf("value contained %d decimal points expected 1", decimal_count))
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
					
					float32_temp, float32_temp_error := strconv.ParseFloat(*string_value, 32)
					if float32_temp_error != nil {
					} else {
						data_type = "float32"
						float32_conv := float32(float32_temp)
						float32_value = &float32_conv
					}
				}

				if len(errors) > 0 {
					return errors
				}
			} else {
				if negative_number {
					data_type = "int64"
					int64_temp, int64_temp_error := strconv.ParseInt(*string_value, 10, 64)
					if int64_temp_error != nil {
						errors = append(errors, fmt.Errorf("strconv.ParseInt(*string_value, 10, 64) error"))
					} else {
						int64_value = &int64_temp
						if *int64_value >= -128 && *int64_value <= 127 {
							data_type = "int8"
							int8_temp, int8_temp_error := strconv.ParseInt(*string_value, 10, 8)
							if int8_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseInt(*string_value, 10, 8) error"))
							} else {
								int8_conv := int8(int8_temp)
								int8_value = &int8_conv
							}
						} else if *int64_value >= -32768 && *int64_value <= 32767 {
							data_type = "int16"
							int16_temp, int16_temp_error := strconv.ParseInt(*string_value, 10, 16)
							if int16_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseInt(*string_value, 10, 16) error"))
							} else {
								int16_conv := int16(int16_temp)
								int16_value = &int16_conv
							}
						} else if *int64_value >= -2147483648 && *int64_value <= 2147483647 {
							data_type = "int32"
							int32_temp, int32_temp_error := strconv.ParseInt(*string_value, 10, 32)
							if int32_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseInt(*string_value, 10, 32) error"))
							} else {
								int32_conv := int32(int32_temp)
								int32_value = &int32_conv
							}
						}
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
						if *uint64_value >= 0 && *uint64_value <= 255 {
							data_type = "uint8"
							int8_temp, int8_temp_error := strconv.ParseUint(*string_value, 10, 8)
							if int8_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseUInt(*string_value, 10, 8) error"))
							} else {
								int8_conv := uint8(int8_temp)
								uint8_value = &int8_conv
							}
						} else if *uint64_value >= 256 && *uint64_value <= 65535 {
							data_type = "uint16"
							int16_temp, int16_temp_error := strconv.ParseUint(*string_value, 10, 16)
							if int16_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseUInt(*string_value, 10, 16) error"))
							} else {
								int16_conv := uint16(int16_temp)
								uint16_value = &int16_conv
							}
						} else if *uint64_value >= 65536  && *uint64_value <= 4294967295 {
							data_type = "uint32"
							int32_temp, int32_temp_error := strconv.ParseUint(*string_value, 10, 32)
							if int32_temp_error != nil {
								errors = append(errors, fmt.Errorf("strconv.ParseUInt(*string_value, 10, 32) error"))
							} else {
								int32_conv := uint32(int32_temp)
								uint32_value = &int32_conv
							}
						}
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
	}

	if data_type == "" {
		errors = append(errors, fmt.Errorf("data_type is unknown please implement"))
	}


	if len(errors) > 0 {
		return errors
	}

	if fmt.Sprintf("%T", list.Front().Value) == "*class.Array" {
		data_array := list.Front().Value.(*Array)
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
		} else if data_type == "int8" {
			*data_array = append(*data_array, int8_value)
		} else if data_type == "int16" {
			*data_array = append(*data_array, int16_value)
		} else if data_type == "int32" {
			*data_array = append(*data_array, int32_value)
		}  else if data_type == "int64" {
			*data_array = append(*data_array, int64_value)
		} else if data_type == "uint8" {
			*data_array = append(*data_array, uint8_value)
		} else if data_type == "uint16" {
			*data_array = append(*data_array, uint16_value)
		} else if data_type == "uint32" {
			*data_array = append(*data_array, uint32_value)
		} else if data_type == "uint64" {
			*data_array = append(*data_array, uint64_value)
		}
	}

	if fmt.Sprintf("%T", list.Front().Value) == "*class.Map" {
		if data_type == "string" {
			list.Front().Value.(*Map).SetString(temp_key, string_value)
		} else if data_type == "bool" {
			list.Front().Value.(*Map).SetBool(temp_key, boolean_value)
		} else if data_type == "null" {
			list.Front().Value.(*Map).SetNil(temp_key)
		} else if data_type == "float64" {
			list.Front().Value.(*Map).SetFloat64(temp_key, float64_value)
		} else if data_type == "float32" {
			list.Front().Value.(*Map).SetFloat32(temp_key, float32_value)
		} else if data_type == "int8" {
			list.Front().Value.(*Map).SetInt8(temp_key, int8_value)
		} else if data_type == "int16" {
			list.Front().Value.(*Map).SetInt16(temp_key, int16_value)
		} else if data_type == "int32" {
			list.Front().Value.(*Map).SetInt32(temp_key, int32_value)
		} else if data_type == "int64" {
			list.Front().Value.(*Map).SetInt64(temp_key, int64_value)
		} else if data_type == "uint8" {
			list.Front().Value.(*Map).SetUInt8(temp_key, uint8_value)
		} else if data_type == "uint16" {
			list.Front().Value.(*Map).SetUInt16(temp_key, uint16_value)
		} else if data_type == "uint32" {
			list.Front().Value.(*Map).SetUInt32(temp_key, uint32_value)
		} else if data_type == "uint64" {
			list.Front().Value.(*Map).SetUInt64(temp_key, uint64_value)
		}
	}

	if len(errors) > 0 {
		return errors
	}


	return nil
}

func ConvertInterfaceValueToJSONStringValue(value interface{}) (*string, []error) {
	var errors []error
	result := "null"

	if value == nil {
		return &result, nil
	}

	string_value := fmt.Sprintf("%s", value)

	if string_value == "<nil>" {
		return &result, nil
	}

	rep := fmt.Sprintf("%T", value)

	if string_value == "%!s("+rep+"=<nil>)" {
		return &result, nil
	}
	
	switch rep {
	case "string":
		result = "\"" + strings.ReplaceAll(value.(string), "\"", "\\\"") + "\""
	case "*string":
		result = "\"" + strings.ReplaceAll(*(value.(*string)), "\"", "\\\"") + "\""
	case "error":
		result = "\"" + strings.ReplaceAll(value.(error).Error(), "\"", "\\\"") + "\""
	case "*error":
		result = "\"" + strings.ReplaceAll((*(value.(*error))).Error(), "\"", "\\\"") + "\""
	case "*url.Error":
		result = "\"" + strings.ReplaceAll((*(value.(*url.Error))).Error(), "\"", "\\\"") + "\""
	case "exec.ExitError":
		result = "\"" + strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\"") + "\""
	case "*exec.ExitError":
		result = "\"" + strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\"") + "\""
	case "errors.errorString":
		result = "\"" + strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\"") + "\""
	case "*errors.errorString":
		result = "\"" + strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\"") + "\""
	case "bool":
		temp := value.(bool)
		if temp {
			result = "true"
		} else {
			result = "false"
		}
	case "*bool":
		temp := *(value.(*bool))
		if temp {
			result = "true"
		} else {
			result = "false"
		}
	case "class.Map":
		x, x_error := value.(Map).ToJSONString()
		if x_error != nil {
			errors = append(errors, x_error...)
		} else {
			result = *x
		}
	case "*class.Map":
		x, x_error := (*value.(*Map)).ToJSONString()
		if x_error != nil {
			errors = append(errors, x_error...)
		} else {
			result = *x
		}
	case "class.Array":
		x, x_error := value.(Array).ToJSONString()
		if x_error != nil {
			errors = append(errors, x_error...)
		} else {
			result = *x
		}
	case "*class.Array":
		x, x_error := (*(value.(*Array))).ToJSONString()
		if x_error != nil {
			errors = append(errors, x_error...)
		} else {
			result = *x
		}
	case "[]string":
		result = "["
		array_length := len(value.([]string))
		for array_index, array_value := range value.([]string) {
			string_conversion, string_conversion_error := ConvertInterfaceValueToJSONStringValue(array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} else {
				result += *string_conversion
			}
			if array_index < (array_length - 1) {
				result += ","
			}
		}
		result += "]"
	case "*[]string":
		result = "["
		array_length := len(*(value.(*[]string)))
		for array_index, array_value := range *(value.(*[]string)) {
			string_conversion, string_conversion_error := ConvertInterfaceValueToJSONStringValue(array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} else {
				result += *string_conversion
			}
			if array_index < (array_length - 1) {
				result += ","
			}
		}
		result += "]"
	case "[]error":
		result = "["
		array_length := len(value.([]error))
		for array_index, array_value := range value.([]error) {
			string_conversion, string_conversion_error := ConvertInterfaceValueToJSONStringValue(array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} else {
				result += *string_conversion
			}
			if array_index < (array_length - 1) {
				result += ","
			}
		}
		result += "]"
	case "*[]error":
		result = "["
		array_length := len(*(value.(*[]error)))
		for array_index, array_value := range *(value.(*[]error)) {
			string_conversion, string_conversion_error := ConvertInterfaceValueToJSONStringValue(array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} else {
				result += *string_conversion
			}
			if array_index < (array_length - 1) {
				result += ","
			}
		}
		result += "]"
	case "func(string, *string, string, string) []error", "func(class.Map) []error", "*func(class.Map) []error":
		result = fmt.Sprintf("\"%s\"", rep)
	case "*class.Host":
		x, x_errors := (*(value.(*Host))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.Credentials":
		x, x_errors := (*(value.(*Credentials))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.DatabaseCreateOptions":
		x, x_errors := (*(value.(*DatabaseCreateOptions))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.Database":
		x, x_errors := (*(value.(*Database))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.Client":
		x, x_errors := (*(value.(*Client))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.Table":
		x, x_errors := (*(value.(*Table))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*class.Record":
		x, x_errors := (*(value.(*Record))).ToJSONString()
		if x_errors != nil {
			errors = append(errors, x_errors...)
		} else {
			result = *x
		}
	case "*time.Time":
		result = "\"" + (*(value.(*time.Time))).Format("2006-01-02 15:04:05.000000") + "\""
	case "map[string]map[string][][]string":
		result = "\"map[string]map[string][][]string\""
	case "*uint64":
		result = strconv.FormatUint(*(value.(*uint64)), 10)
	case "uint64":
		result = strconv.FormatUint(value.(uint64), 10)
	case "*uint32":
		result = strconv.FormatUint(uint64(*(value.(*uint32))), 10)
	case "uint32":
		result = strconv.FormatUint(uint64(value.(uint32)), 10)
	case "*uint16":
		result = strconv.FormatUint(uint64(*(value.(*uint16))), 10)
	case "uint16":
		result = strconv.FormatUint(uint64(value.(uint16)), 10)
	case "*uint8":
		result = strconv.FormatUint(uint64(*(value.(*uint8))), 10)
	case "uint8":
		result = strconv.FormatUint(uint64(value.(uint8)), 10)
	case "*uint":
		result = strconv.FormatUint(uint64(*(value.(*uint))), 10)
	case "uint":
		result = strconv.FormatUint(uint64(value.(uint)), 10)
	case "*int64":
		result = strconv.FormatInt(*(value.(*int64)), 10)
	case "int64":
		result = strconv.FormatInt(value.(int64), 10)
	case "*int32":
		result = strconv.FormatInt(int64(*(value.(*int32))), 10)
	case "int32":
		result = strconv.FormatInt(int64((value.(int32))), 10)
	case "*int16":
		result = strconv.FormatInt(int64(*(value.(*int16))), 10)
	case "int16":
		result = strconv.FormatInt(int64((value.(int16))), 10)
	case "*int8":
		result = strconv.FormatInt(int64(*(value.(*int8))), 10)
	case "int8":
		result = strconv.FormatInt(int64((value.(int8))), 10)
	case "*int":
		result = strconv.FormatInt(int64(*(value.(*int))), 10)
	case "int":
		result = strconv.FormatInt(int64(value.(int)), 10)
	case "*float64":
		result = fmt.Sprintf("%f", *(value.(*float64)))
	case "float64":
		result = fmt.Sprintf("%f", (value.(float64)))
	case "*float32":
		result = fmt.Sprintf("%f", *(value.(*float32)))
	case "float32":
		result = fmt.Sprintf("%f", (value.(float32)))
	default:
		errors = append(errors, fmt.Errorf("JSON.ConvertInterfaceValueToJSONStringValue: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}


	return &result, nil
}