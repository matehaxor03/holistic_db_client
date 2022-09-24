package class


import (
	"fmt"
	"unicode"
	"reflect"
	"strings"
	//common "github.com/matehaxor03/holistic_db_client/common"

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
		return errors
	}

	if len(array) == 0 {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: has empty array"))
		return errors
	}

	for _, value := range array {
		if value == "" {
			errors = append(errors, fmt.Errorf("ContainsExactMatch: array has empty value"))
		}
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: compare value is nil"))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("ContainsExactMatch: compare value is nil"))
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

func ValidateGeneric(args...[]map[string]interface{}) *Result {	
	
	if args == nil {
		result := NewResult()
		(*result).LogError(fmt.Errorf("ValidateGeneric received a nil args"))
		return result
	}

	if len(args) == 0 {
		result := NewResult()
		(*result).LogError(fmt.Errorf("ValidateGeneric received a empty args"))
		return result
	}
	
	if len(args[0]) != 1 {
		result := NewResult()
		(*result).LogError(fmt.Errorf("ValidateGeneric received %s args however expected %s", len(args[0]), 1))
		return result
	}
	
	payload := ConvertPrimitiveMapToMap(args[0][0])
	//payload := array_of_args[0].(Map)
	//panic(payload.ToJSONString())
	result, err := payload.GetResult("result")
	if result == nil || err != nil {
		result = NewResult()
		(*result).LogError(fmt.Errorf("ValidateGeneric no result reference received"))
		return result
	}

	top_level_keys := strings.Join(payload.Keys(), " ")
	var function, function_exists = payload["function"]
	if !function_exists {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: function keys: %s", top_level_keys))
	} else if function == nil {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args ha key: function but it had value nil"))
	} 

	var parameters = payload.M("parameters")
	var parameter_keys = strings.Join(parameters.Keys(), " ")
	if parameters == nil {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: parameters keys: %s", parameter_keys))
	} 

	if len((*result).GetErrors()) > 0 {
		return result
	}

	
	whitelist := parameters.A("whitelist|[]string")

	if whitelist == nil {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: parameters->whiltelist keys: %s", parameter_keys))
	} else if len(whitelist) == 0 {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had key: parameters->whiltelist but it did have an empty value"))
	} 

	var data, data_err = parameters.S("data|string")
	if data_err == nil || data == "" {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION:Common: ValidateGeneric args had empty data: parameters->data keys: %s", parameter_keys))
	}

	var data_type, data_type_err = parameters.S("type|data_type")
	if (data_type_err == nil || data_type == "") {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had empty data: parameters->kind keys: %s", parameter_keys))
	} 

	var column_name, column_name_err = parameters.S("column_name|string")
	if ( column_name_err == nil || column_name == "") {
		(*result).LogError(fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had the key: parameters->_column_name however had an empty value" ))
	} 

	if len((*result).GetErrors()) > 0 {
		return result
	}

	output := (*result).GetData()
	(*output).SetArray("whitelist|[]string", whitelist)
	(*output).SetString("data|string", data)
	(*output).SetString("type|data_type", data_type)
	(*output).SetString("column_name|string", column_name)

	panic("zzzzzz" + output.ToJSONString())

	return result
}

func ContainsExactMatchz(args...map[string]interface{}) *Result {
	var result = ValidateGeneric(args)

	if len((*result).GetErrors()) > 0 {
		return result
	}

	data_obj := (*result).GetData()
	var whitelist = data_obj.A("whitelist|[]string")
	
	var data, data_err = data_obj.S("data|string")
	if data_err != nil {
		(*result).LogError(data_err)
	}

	var data_type, data_type_err = data_obj.S("type|data_type")
	if data_type_err != nil {
		(*result).LogError(data_type_err)
	}
	
	var column_name, column_name_err = data_obj.S("column_name|string")
	if column_name_err != nil {
		(*result).LogError(column_name_err)
	}

	var containsExactMatchErrors = ContainsExactMatch(whitelist.ToPrimativeArray(), &data, column_name, data_type)

	if containsExactMatchErrors != nil {
		(*result).LogErrors(containsExactMatchErrors)
	}

	panic((*result).GetErrors()[0])

	return result
}

func Validate(args...interface{}) []error {
	var errors []error 


	return errors
}

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

func ValidateCharacters(whitelist string, str *string, label string, reflect_value reflect.Value) ([]error) {
	var errors []error 

	if str == nil {
		errors = append(errors, fmt.Errorf("%s %s is nil", reflect_value.Kind(), label))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("%s %s is empty", reflect_value.Kind(), label))
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
			errors = append(errors, fmt.Errorf("%s invalid letter %s for %s please use %s", reflect_value.Kind(), string(letter), label, whitelist))
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


func ValidateOptions(extra_options map[string]map[string][][]string, reflect_value reflect.Value) ([]error) {
	var errors []error 
	var VALID_CHARACTERS = GetConstantValueAllowedCharacters()
	for key, value := range extra_options {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options root key %s", key), reflect_value)
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		for key2, value2 := range value {
			var combined = ""
			for _, array_value := range value2 {
				combined += strings.Join(array_value, "")
			}

			value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), reflect_value)
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



