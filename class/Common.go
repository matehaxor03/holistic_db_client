package class


import (
	"fmt"
	"unicode"
	"reflect"
	"strings"
)

func Examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		Examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	default:
		panic(fmt.Sprintf("kind %s is not supported please implement", t.Kind()))
	}
}


func FIELD_NAME_VALIDATION_FUNCTIONS() string {
	return "validation_functions"
}

func FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS() string {
	return "validation_functions_parameters"
}

func ContainsExactMatch(array []string, str *string, label string, data_type string) []error {
	var errors []error 
	if array == nil {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: has nil array"))
	}

	if len(array) == 0 {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: has empty array"))
	}

	for _, value := range array {
		if value == "" {
			errors = append(errors, fmt.Errorf("ContainsExactMatch: array has empty value"))
		}
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: compare value is nil"))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: compare value is empty"))
	}

	if len(errors) > 0 {
		return errors
	}
	
	for _, value := range array {
		if value == *str {
			return nil
		}
	}

    errors = append(errors, fmt.Errorf("%s: %s has value '%s' expected to have value in %s", data_type, label, *str , array))
	return errors
}

/*
func ValidateGeneric(payload Map) []error {	
	errors := []error{}
	if payload == nil {
		errors = append(errors, fmt.Errorf("ValidateGeneric received a nil payload"))
		return errors
	}

	var function, function_exists = payload["function"]
	if !function_exists {
		errors = append(errors, fmt.Errorf("ValidateGeneric: key: 'function' did not exist"))
	} else if function == nil {
		errors = append(errors, fmt.Errorf("ValidateGeneric: key: 'function' had a nil value"))
	} 

	var parameters = payload.M("parameters")
	
	var parameter_keys = strings.Join(parameters.Keys(), " ")
	if parameters == nil {
		errors = append(errors, fmt.Errorf("ValidateGeneric: key 'parameters' did not exist"))
	} 

	if len(errors) > 0 {
		return errors
	}
	
	whitelist, whitelist_err := parameters.A("whitelist|[]string")

	if whitelist_err != nil {
		errors = append(errors, fmt.Errorf("ValidateGeneric key: parameters->whiltelist did not exist"))
	} else if len(whitelist) == 0 {
		errors = append(errors, fmt.Errorf("ValidateGeneric args had key: parameters->whiltelist but it did have an empty value"))
	} 

	var data, data_err = parameters.S("data|string")
	if data_err != nil || data == "" {
		errors = append(errors, fmt.Errorf("ValidateGeneric had nil or empty data: parameters->data keys: %s", parameter_keys))
	}

	var data_type, data_type_err = parameters.S("type|data_type")
	if (data_type_err != nil || data_type == "") {
		errors = append(errors, fmt.Errorf("ValidateGeneric had nil or empty data: parameters->kind keys: %s", parameter_keys))
	} 

	var column_name, column_name_err = parameters.S("column_name|string")
	if ( column_name_err != nil || column_name == "") {
		errors = append(errors, fmt.Errorf("ValidateGeneric had nil or emtpy data parameters->column_name: %s", parameter_keys))
	} 

	return nil
}
*/

/*
func ContainsExactMatchz(payload Map) []error {
	errors := []error{}

	if ValidateGeneric(payload) != nil {
		return errors
	}
	
	parameters := payload.M("parameters")
	whitelist, _ := parameters.A("whitelist|[]string")
	data, _ := parameters.S("data|string")
	column_name, _ := parameters.S("column_name|string")
	data_type, _ := parameters.S("type|data_type")



	var containsExactMatchErrors = ContainsExactMatch(whitelist.ToPrimativeArray(), &data, column_name, data_type)

	if containsExactMatchErrors != nil {
		errors = append(errors, containsExactMatchErrors...)
	}

	if errors != nil && len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			fmt.Println(errors[i])
		}
		return errors
	}

	return nil
}


func Validate(args...interface{}) []error {
	var errors []error 


	return errors
}
*/

/*
func CloneStringPtr(s* string) (*string) {
	if s == nil {
		return nil
	}

	var builder strings.Builder
    if _, err := sb.WriteString(*s); err != nil {
        //log & return
    }
    newstring := sb.String()
	return &newstring
}*/

func ArraysContainsArraysOrdered(a [][]string, b [][]string, label string, reflect_value reflect.Value) []error {
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
			errors = append(errors, fmt.Errorf("%s %s has value '%s' expected to have value in '%s", reflect_value.Kind(), label, b_value, a))
		}
	}
	

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateCharacters(whitelist string, str *string, label string, typeOf string) ([]error) {
	var errors []error 

	if str == nil {
		errors = append(errors, fmt.Errorf("%s %s is nil", typeOf, label))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("%s %s is empty", typeOf, label))
		return errors
	}

	for _, letter := range *str {
		found := false

		for _, check := range whitelist {
			if check == letter {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("%s invalid letter %s for %s please use %s", typeOf, string(letter), label, whitelist))
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
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
}

func GetAllowedStringValues() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
}


func ValidateOptions(extra_options map[string]map[string][][]string, reflect_value reflect.Value) ([]error) {
	var errors []error 
	var VALID_CHARACTERS = GetConstantValueAllowedCharacters()
	for key, value := range extra_options {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options root key %s", key), fmt.Sprintf("%T", reflect_value))
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		for key2, value2 := range value {
			var combined = ""
			for _, array_value := range value2 {
				combined += strings.Join(array_value, "")
			}

			value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), fmt.Sprintf("%T", reflect_value))
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

func GetLogicCommand(command string, field_name string, allowed_options map[string]map[string][][]string, options map[string]map[string][][]string, reflect_value reflect.Value) (*string, []error){
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
		errors = append(errors, fmt.Errorf("%s field: %s is not supported but was provided", reflect_value.Kind(), field_name))
		return nil, errors
	}

	logic_option_command_value, logic_option_command_exists := logic_option_value[command]
	if !logic_option_command_exists {
		return &logic_option, nil
	}
	
	allowed_logic_option_command_value, allowed_logic_option_command_exists := allowed_logic_option_value[command]

	if !allowed_logic_option_command_exists {
		errors = append(errors, fmt.Errorf("%s field: %s is not supported for command: %s but was provided", reflect_value.Kind(), field_name, command))
		return nil, errors
	}


	logic_option_errors := ArraysContainsArraysOrdered(allowed_logic_option_command_value, logic_option_command_value, field_name + "->" + command, reflect_value)
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
	var errors[]error 
	var parameters = fields.Keys()
	for _, parameter := range parameters {
		var errors_for_param []error 
		isOptional := false
		error_on_value_of := false

		parameter_fields := fields.M(parameter)
		typeOf := parameter_fields.S("type|string")

		if typeOf == nil {
			errors = append(errors, fmt.Errorf("tyoe of not specified for %s->%s", structType, parameter))
			continue
		}
		
		switch *typeOf {
		case "string":
			valueOf := parameter_fields.S("value|string")
			
			if valueOf == nil {
				error_on_value_of = true
			}
			
			filters := parameter_fields.A(FILTERS())
			if filters != nil {
				all_compare_filters := GET_FILTERS()
				for _, filter := range filters {

					compare_filter := all_compare_filters.M(filter.(string))
					if compare_filter == nil {
						errors_for_param = append(errors_for_param, fmt.Errorf("please implement filter for %s", filter))
						continue
					}

					compare_filter_values := compare_filter.A(parameter)
					if compare_filter_values == nil {
						errors_for_param = append(errors_for_param, fmt.Errorf("please implement create compare values for %s->%s", filter, parameter))
						continue
					}

					errors_of_filter := ContainsExactMatch(compare_filter_values.ToPrimativeArray(), valueOf, parameter, structType)
					if errors_of_filter != nil {
						errors_for_param = append(errors_for_param, errors_of_filter...)
					}	




					
					/*
					switch filter {
					case WHITELIST_FILTER():
						
						

						errors_of_filter := ContainsExactMatch(allowed_values.ToPrimativeArray(), valueOf, parameter, structType)
						if errors_of_filter != nil {
							errors_for_param = append(errors_for_param, errors_of_filter...)
						}	
						
					default:
					}*/	
				}
			}
			

			default_errors_filter := ValidateCharacters(GetAllowedStringValues(), valueOf, parameter, structType)
			if default_errors_filter != nil {
				errors_for_param = append(errors_for_param, default_errors_filter...)
			}

			break
		default:
			panic(fmt.Sprintf("please implement type %s", typeOf))
		}

		if error_on_value_of && isOptional {
			errors_for_param = nil
		} else {
			errors = append(errors, errors_for_param...)
		}
	}

	return errors
}
