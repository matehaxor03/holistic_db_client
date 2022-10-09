package class


import (
	"fmt"
	"unicode"
	"reflect"
	"strings"
)

func EscapeString(value string) (string) { 
	value = strings.ReplaceAll(value, "\\", "\\\\")
    value = strings.ReplaceAll(value, "'", "\\'")
	return value
}

func CloneString(value *string) (*string) {
	if value == nil {
		return nil
	}

	temp := strings.Clone(*value)
	return &temp
}

func FIELD_NAME_VALIDATION_FUNCTIONS() string {
	return "validation_functions"
}

func FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS() string {
	return "validation_functions_parameters"
}

func getWhitelistStringFunc() (*func(m Map) []error) {
	function := WhiteListString
	return &function
}

func WhiteListString(m Map) []error {
	array := m.A("values")
	str := m.S("value")
	label := m.S("label")
	data_type := m.S("data_type")
		
	var errors []error 
	if array == nil {
		errors = append(errors, fmt.Errorf("%s: %s: ContainsExactMatch: has nil array", *data_type, *label))
	}

	if len(array) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: ContainsExactMatch: has empty array", *data_type, *label))
	}

	for _, value := range array {
		if value == "" {
			errors = append(errors, fmt.Errorf("%s: %s: ContainsExactMatch: array has empty value", *data_type, *label))
		}
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: ContainsExactMatch: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: ContainsExactMatch: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}
	
	for _, value := range array {
		if value == *str {
			return nil
		}
	}

    errors = append(errors, fmt.Errorf("%s: %s has value '%s' expected to have value in %s", *data_type, *label, *str , array))
	return errors
}

func ArraysContainsArraysOrdered(a [][]string, b [][]string, label string, typeof string) []error {
	var errors []error 
	
	for _, b_value := range b {
		var current = strings.Join(b_value, "")
		var found = false
		for _, a_value := range a {
			var compare =  strings.Join(a_value, "")

			if current == compare {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("%s %s has value '%s' expected to have value in '%s", typeof, label, b_value, a))
		}
	}
	

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func getWhitelistCharactersFunc() (*func(m Map) []error) {
	funcValue := WhitelistCharacters
	return &funcValue
}

func WhitelistCharacters(m Map) ([]error) {
	var errors []error 

	whitelist := m.S("values")
	str := (m.S("value"))
	label := (m.S("label"))
	typeOf := (m.S("data_type"))

	if str == nil {
		errors = append(errors, fmt.Errorf("%s %s value is nil", *typeOf, *label))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("%s %s value is empty", *typeOf, *label))
		return errors
	}

	if whitelist == nil {
		errors = append(errors, fmt.Errorf("%s %s values is nil", *typeOf, *label))
		return errors
	}

	if *whitelist == "" {
		errors = append(errors, fmt.Errorf("%s %s values is empty", *typeOf, *label))
		return errors
	}

	for _, letter := range *str {
		found := false

		for _, check := range *whitelist {
			if check == letter {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("%s invalid letter %s for %s please use %s", *typeOf, string(letter), *label, *whitelist))
		}
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
 }

 func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func GetConstantValueAllowedCharacters() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_=.*"
}

func GetAllowedStringValues() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_=.*"
}


func ValidateOptions(extra_options map[string]map[string][][]string, reflect_value reflect.Value) ([]error) {
	var errors []error 
	var VALID_CHARACTERS = GetConstantValueAllowedCharacters()
	for key, value := range extra_options {
		filters1 := Map{"values":VALID_CHARACTERS, "value":&key, "label":fmt.Sprintf("extra_options root key %s", key), "typeOf":fmt.Sprintf("%T", reflect_value) }
		
		key_extra_options_errors := WhitelistCharacters(filters1)
		//key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options root key %s", key), fmt.Sprintf("%T", reflect_value))
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		for key2, value2 := range value {
			var combined = ""
			for _, array_value := range value2 {
				combined += strings.Join(array_value, "")
			}

			filters2 := Map{"values":VALID_CHARACTERS, "value":&combined, "label":fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), "typeOf":fmt.Sprintf("%T", reflect_value) }

			value_extra_options_errors := WhitelistCharacters(filters2)
			//value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), fmt.Sprintf("%T", reflect_value))
			if value_extra_options_errors != nil {
				errors = append(errors, value_extra_options_errors...)	
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func GetLogicCommand(command string, field_name string, allowed_options map[string]map[string][][]string, options map[string]map[string][][]string, typeof string) (*string, []error){
	var errors []error 
	logic_option := ""

	if options == nil {
		return &logic_option, nil
	}
	
	logic_option_value, logic_option_exists := options[field_name]
	if !logic_option_exists {
		return &logic_option, nil
	}

	allowed_logic_option_value, allowed_logic_option_exists := allowed_options[field_name]

	if !allowed_logic_option_exists {
		errors = append(errors, fmt.Errorf("%s field: %s is not supported but was provided", typeof, field_name))
		return nil, errors
	}

	logic_option_command_value, logic_option_command_exists := logic_option_value[command]
	if !logic_option_command_exists {
		return &logic_option, nil
	}
	
	allowed_logic_option_command_value, allowed_logic_option_command_exists := allowed_logic_option_value[command]

	if !allowed_logic_option_command_exists {
		errors = append(errors, fmt.Errorf("%s field: %s is not supported for command: %s but was provided", typeof, field_name, command))
		return nil, errors
	}


	logic_option_errors := ArraysContainsArraysOrdered(allowed_logic_option_command_value, logic_option_command_value, field_name + "->" + command, typeof)
	if logic_option_errors != nil {
		errors = append(errors, logic_option_errors...)	
		return nil, errors
	} 

	for _, statements := range logic_option_command_value {
		logic_option += strings.Join(statements, " ") + " "
	}

	return &logic_option, nil
}

func ValidateGenericSpecial(fields Map, structType string) []error {
	var errors []error 
	var parameters = fields.Keys()
	for _, parameter := range parameters {
		value_is_mandatory := true
		value_is_null := false

		parameter_fields := fields.M(parameter)

		typeOf := fmt.Sprintf("%T", parameter_fields["value"])
		if typeOf == "nil" {
			value_is_null = true
		}

		mandatory_field := parameter_fields.B("mandatory")
		
		if mandatory_field != nil {
			value_is_mandatory = *mandatory_field
		}

		if value_is_null && !value_is_mandatory {
			continue
		}
		
		switch typeOf {
		case "map[string]map[string][][]string":
			// todo: convert these to objects for validations
			break
		case "string":
		case "*string":
			valueOf := parameter_fields.S("value")

			filters := parameter_fields.A(FILTERS())
			if filters != nil {
				for filter_index, filter := range filters {
					filter_map := filter.(Map)

					if !filter_map.HasKey("function") {
						panic(fmt.Sprintf("parameter: %s does not have function parameter for filter index: %d " + filter_map.ToJSONString(), parameter, filter_index))
					}

					function := filter_map.Func("function")	
					if function == nil {
						panic(fmt.Sprintf("parameter: %s has function parameter for filter index: %d but it has nil value " + filter_map.ToJSONString(), parameter, filter_index))
					}
										
					filter_map.SetString("value", valueOf)
					filter_map.SetString("data_type", &structType)
					temp := "ValidateGenericSpecial"
					filter_map.SetString("label", &temp)

				
					function_errors := function(filter.(Map))
					if function_errors != nil {
						errors = append(errors, function_errors...)
					}	
					/*
					var vargsConvert = []reflect.Value{reflect.ValueOf(filter)}

					var output_array_map_result = reflect.ValueOf(function).Call(vargsConvert)

					validation_errors := ConvertPrimitiveReflectValueArrayToArray(output_array_map_result)
					outer_array_length := len(validation_errors)
					for i := 0; i < outer_array_length; i++ {
						validation_error := validation_errors[i]
						error_value := fmt.Sprintf("%s", reflect.ValueOf(validation_error).Interface())
						if error_value == "[]" {
							continue
						}
						errors = append(errors, fmt.Errorf(error_value))
					}*/
				}
			}

			default_allowed_values := GetAllowedStringValues()
			default_allowed_valuesAddress := &default_allowed_values
			default_filter_params := Map{"values":default_allowed_valuesAddress, "value":valueOf, "label": &parameter, "data_type":&structType }

			default_errors_filter := WhitelistCharacters(default_filter_params)

			if default_errors_filter != nil {
				errors = append(errors, default_errors_filter...)
			}
			break
		case "*int64":
			valueOf := parameter_fields.GetInt64("value")
			
			if valueOf == nil {
				value_is_null = true
			}

			if value_is_null && !value_is_mandatory {
				continue
			}


		case "*time.Time":
			valueOf := parameter_fields.GetTime("value")
			
			if valueOf == nil {
				value_is_null = true
			}

			if value_is_null && !value_is_mandatory {
				continue
			}
		case "*class.Database":
			database := parameter_fields.GetObject("value").(*Database)
			if database != nil {
				errors_for_database := database.Validate()
				if errors_for_database != nil {
					errors = append(errors, errors_for_database...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.DomainName":
			domain_name := parameter_fields.GetObject("value").(*DomainName)
			if domain_name != nil {
				errors_for_domain_name := domain_name.Validate()
				if errors_for_domain_name != nil {
					errors = append(errors, errors_for_domain_name...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.Host":
			host := parameter_fields.GetObject("value").(*Host)
			if host != nil {
				errors_for_host := host.Validate()
				if errors_for_host != nil {
					errors = append(errors, errors_for_host...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.Credentials":
			credentials := parameter_fields.GetObject("value").(*Credentials)
			if credentials != nil {
				errors_for_credentaials := credentials.Validate()
				if errors_for_credentaials != nil {
					errors = append(errors, errors_for_credentaials...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.DatabaseCreateOptions":
			database_create_options := parameter_fields.GetObject("value").(*DatabaseCreateOptions)
			if database_create_options != nil {
				errors_for_database_create_options := database_create_options.Validate()
				if errors_for_database_create_options != nil {
					errors = append(errors, errors_for_database_create_options...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.Client":
			client := parameter_fields.GetObject("value").(*Client)
			if client != nil {
				errors_for_client := client.Validate()
				if errors_for_client != nil {
					errors = append(errors, errors_for_client...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.Grant":
			grant := parameter_fields.GetObject("value").(*Grant)
			if grant != nil {
				errors_for_grant := grant.Validate()
				if errors_for_grant != nil {
					errors = append(errors, errors_for_grant...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "*class.User":
			user := parameter_fields.GetObject("value").(*User)
			if user != nil {
				errors_for_user := user.Validate()
				if errors_for_user != nil {
					errors = append(errors, errors_for_user...)
				}
			} else if value_is_mandatory {
				errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
			}
			break
		case "<nil>":
			if !parameter_fields.HasKey("default") {
				if value_is_null && value_is_mandatory {
					errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
				}
			} else if !parameter_fields.HasKey("type") {
				if value_is_null && value_is_mandatory {
					errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil and did not have type", parameter))
				}
			} else {
				typeOf := parameter_fields.S("type")
				switch *typeOf {
				case "*time.Time":
					type_of_default := parameter_fields.GetType("default")
					switch type_of_default {
					case "*string":
						type_of_default_value := parameter_fields.S("default")
						if *type_of_default_value != "now" {
							errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s and had default: %s please implement default value: %s", parameter, *typeOf, type_of_default, *type_of_default_value))
						}
					default:
						errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s please implement default for %s", parameter, *typeOf, type_of_default))
					}
				case "*int64":
					type_of_default := parameter_fields.GetType("default")
					switch type_of_default {
					case "*int":
					case "int":
						type_of_default_value := parameter_fields.GetInt64("default")
						if type_of_default_value == nil {
							errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s and had default: %s please implement default value: %s", parameter, *typeOf, type_of_default, "nil"))
						}
					default:
						errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s please implement default for %s", parameter, *typeOf, type_of_default))
					}
				default:
					errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil please implement for type: %s", parameter, *typeOf))
				}
			}
		default:
			panic(fmt.Sprintf("please implement type %s", typeOf))
		}

		
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
	   if value == name {
		  return true
	   }
	}
	return false
 }
