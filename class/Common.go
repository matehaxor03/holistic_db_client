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

func ContainsExactMatch(m Map) []error {
	m = ConvertPrimitiveMapToMap(m)

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

func ValidateCharacters(m Map) ([]error) {
	var errors []error 
	m = ConvertPrimitiveMapToMap(m)
	whitelist := *(m.S("whitelist"))
	str := (m.S("str"))
	label := (m.S("label"))
	typeOf := (m.S("typeOf"))


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
		filters1 := Map{"whitelist":VALID_CHARACTERS, "str":&key, "label":fmt.Sprintf("extra_options root key %s", key), "typeOf":fmt.Sprintf("%T", reflect_value) }
		
		key_extra_options_errors := ValidateCharacters(filters1)
		//key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options root key %s", key), fmt.Sprintf("%T", reflect_value))
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		for key2, value2 := range value {
			var combined = ""
			for _, array_value := range value2 {
				combined += strings.Join(array_value, "")
			}

			filters2 := Map{"whitelist":VALID_CHARACTERS, "str":&combined, "label":fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), "typeOf":fmt.Sprintf("%T", reflect_value) }

			value_extra_options_errors := ValidateCharacters(filters2)
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
		typeOf := parameter_fields.S("type")
		mandatory_field := parameter_fields.B("mandatory")
		
		if mandatory_field != nil {
			fmt.Println(fmt.Sprintf("is mandatory: %b", *mandatory_field))
			value_is_mandatory = *mandatory_field
		}

		if typeOf == nil {
			errors = append(errors, fmt.Errorf("tyoe of not specified for %s->%s", structType, parameter))
			continue
		}
		
		switch *typeOf {
		case "map[string]map[string][][]string)":
			// todo: convert these to objects for validations
			break
		case "string":
		case "*string":
			valueOf := parameter_fields.S("value")
			
			if valueOf == nil {
				value_is_null = true
			}

			if value_is_null && !value_is_mandatory {
				continue
			}
			
			filters := parameter_fields.A(FILTERS())
			if filters != nil {
				for _, filter := range filters {
					fmt.Println("START")
					fmt.Println(parameter + " " + *valueOf)
					fmt.Println(filter.(Map).ToJSONString())

					function := filter.(Map).Func("function")

					filter.(Map).SetString("value", valueOf)
					filter.(Map).SetString("data_type", &structType)
					temp := "ValidateGenericSpecial"
					filter.(Map).SetString("label", &temp)
				
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
					}
				}
			}


			default_filter_params := Map{"whitelist":GetAllowedStringValues(), "str":valueOf, "label":parameter, "typeOf":structType }

			default_errors_filter := ValidateCharacters(default_filter_params)

			//default_errors_filter := ValidateCharacters(GetAllowedStringValues(), valueOf, parameter, structType)
			if default_errors_filter != nil {
				errors = append(errors, default_errors_filter...)
			}

			break
		case "*Host":
			errors_for_host := parameter_fields.GetObject("value").(*Host).Validate()
			if errors_for_host != nil {
				errors = append(errors, errors_for_host...)
			}
			break
		case "*Credentials":
			errors_for_credentaials := parameter_fields.GetObject("value").(*Credentials).Validate()
			if errors_for_credentaials != nil {
				errors = append(errors, errors_for_credentaials...)
			}
		case "*DatabaseCreateOptions":
			errors_for_database_create_options := parameter_fields.GetObject("value").(*DatabaseCreateOptions).Validate()
			if errors_for_database_create_options != nil {
				errors = append(errors, errors_for_database_create_options...)
			}
		default:
			panic(fmt.Sprintf("please implement type %s", *typeOf))
		}

		
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
