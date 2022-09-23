package class


import (
	"fmt"
	"unicode"
	"reflect"
	"strings"
	"strconv"
)

type Map map[string]interface{}
type Array []interface{}

func ConvertPrimativeArrayToThing(a []interface{}) Array {
	var primitMap = make(map[string]interface{})
	var primitArray []interface{}
	var mappie = Map{}
	var copy = Array{}
	for _, array := range a {
		if reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Map) || 
		   reflect.ValueOf(array).Type() == reflect.TypeOf(mappie) ||
		   reflect.ValueOf(array).Type() == reflect.TypeOf(primitMap) {
			fmt.Println("ConvertPrimativeArrayToThing map")

			thing := ConvertPrimativeMapToMap(array.(map[string]interface{}))
			copy = append(copy, thing)
		} else if reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Slice) ||
		          reflect.ValueOf(array).Type() == reflect.TypeOf(reflect.Array) || 
				  reflect.ValueOf(array).Type() == reflect.TypeOf(copy) || 
				  reflect.ValueOf(array).Type() == reflect.TypeOf(primitArray)  {
					fmt.Println("ConvertPrimativeArrayToThing array")
			thing := ConvertPrimativeArrayToThing(array.([]interface{}))
			copy = append(copy, thing)
		} else {
			fmt.Println("ConvertPrimativeArrayToThing value")
			copy = append(copy, array)
		}	
	}
	return copy
}

func ConvertPrimativeMapToMap(m map[string]interface{}) Map {
	var primitMap = make(map[string]interface{})
	var primitArray []interface{}
	var arrayise = Array{}
	var copy = Map{}
	for key, value := range m {
		fmt.Println("valueof: " + reflect.ValueOf(value).String() + " type: " + reflect.TypeOf(value).String())

		if reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Map) || 
		   reflect.ValueOf(value).Type() == reflect.TypeOf(copy) ||
		   reflect.ValueOf(value).Type() == reflect.TypeOf(primitMap) {
			fmt.Println("ConvertPrimativeMapToMap map")
			copy.setMap(key, ConvertPrimativeMapToMap(value.(map[string]interface{})))
		} else if reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Slice) ||
		          reflect.ValueOf(value).Type() == reflect.TypeOf(reflect.Array) ||
				  reflect.ValueOf(value).Type() == reflect.TypeOf(arrayise) ||
				  reflect.ValueOf(value).Type() == reflect.TypeOf(primitArray) {
					fmt.Println("ConvertPrimativeMapToMap array")
			copy.setArray(key, ConvertPrimativeArrayToThing(value.([]interface{})))
		} else {
			fmt.Println("ConvertPrimativeMapToMap value")
			copy.setValue(key, value)
		}		
	}
	return copy
}

func ConvertIntefaceArrayToStringArray(aInterface []interface{}) []string{
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}

func (m Map) M(s string) Map {
	//maapp := Map{}
	
	//if reflect.ValueOf(value).Type() == 

	fmt.Println(s)
	fmt.Println("valueof: " + reflect.ValueOf(m[s]).String() + " type: " + reflect.TypeOf(m[s]).String())

	return m[s].(Map)
}

func (m Map) M_self() Map {
	return m
}

func (m Map) setMap(key string, value Map) Map {
	m[key] = value
	return m
}

func (m Map) setArray(key string, value Array) Map {
	m[key] = value
	return m
}

func (m Map) setValue(key string, value interface{}) Map {
	m[key] = value
	return m
}


func (m Map) Func(s string) func(...map[string]interface{}) (map[string]interface{}) {
	return m[s].(func(...map[string]interface{}) (map[string]interface{}))
}

func (m Map) Array(s string) []interface{} {
	
	return m[s].([]interface{})
}

func (m Map) S(s string) string {
	return m[s].(string)
}

func (m Map) InterfaceString(s string) interface{} {
	return m[s].(interface{})
}

func (m Map) Spt(s string) *string {
	return m[s].(*string)
}

func (m Map) PrimArray(s string) []string {
	return m[s].([]string)
}

func (m Map) Keys() []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}

func (m Map) M_Value(s string) (interface{}, bool) {
	if m[s] != nil {
		return m[s].(interface{}), true
	}
	return nil, false
}

func KeysForMap(m map[string]interface{}) []string {
	var keys []string
	for a, _ := range m {
		keys = append(keys, a)
	}
	return keys
}

func FIELD_NAME_VALIDATION_FUNCTIONS() string {
	return "validation_functions"
}

func FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS() string {
	return "validation_functions_parameters"
}


func ContainsExactMatch(array []string, str *string, label string, reflect_value reflect.Value) []error {
	for _, array_value := range array {
		if array_value == *str {
			return nil
		}
	}

	var errors []error 
    errors = append(errors, fmt.Errorf("%s: %s has value '%s' expected to have value in %s", reflect_value.Kind(), label, *str , array))
	return errors
}

/*func Containsy(args...interface{}) []error {
	
	/*for _, array_value := range array {
		if array_value == *str {
			return nil
		}
	}

	panic(args[0].(string) + " " + args[1].(string))

	var errors []error 
    //errors = append(errors, fmt.Errorf("%s has value '%s' expected to have value in %s", label, (*str) , array))
	return errors
}*/

func ValidateGeneric(args...[]map[string]interface{}) map[string]interface{} {
	//panic("validate generic")
	var result = make(map[string]interface{})
	result["errors"] = []error{}
	
	if len(args[0]) != 1 {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric received %s args however expected %s", len(args[0]), 1))
	}

	if len(result["errors"].([]error)) > 0 {
		return result
	}

	/*
	var argsArray = make([]map[string]interface{}, len(args))
	for index, currentArg := range args {
		fmt.Println(strings.Join(KeysForMap(currentArg[index]), " "))
		argsArray = append(argsArray, currentArg[index])
	}*/
	//var argsArray = make([]map[string]interface{}, 1)
	//fmt.Println("%s", args)
	//argsArray[0] = args[0].(map[string]interface{})
	var array = args[0]
	fmt.Println("%s", array)
	var array2 = array[0]
	fmt.Println("%s", array2)


	var payload = ConvertPrimativeMapToMap(array2)
	var tempKeys = strings.Join(payload.Keys(), " ")
	var function, function_exists = payload["function"]
	if !function_exists {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: function keys: %s", tempKeys))
	} else if function == nil {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args ha key: function but it had value nil"))
	} else {
		result["function"] = payload["function"]
	}

	var parameters = payload.M("parameters")
	var parameter_keys = strings.Join(parameters.Keys(), " ")
	if parameters == nil {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: parameters keys: %s", parameter_keys))
	} else {
		result["parameters"] = payload["parameters"]
	}

	if len(result["errors"].([]error)) > 0 {
		return result
	}

	var whitelistStringValue = fmt.Sprintf("%s", parameters["whitelist"])
	whitelistStringValue = strings.Replace(whitelistStringValue, "[", "", -1)
	whitelistStringValue = strings.Replace(whitelistStringValue, "]", "", -1)
	var whitelist = strings.Split(whitelistStringValue, " ")

	if whitelist == nil {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args did not have key: parameters->whiltelist keys: %s", parameter_keys))
	} else if len(whitelist) == 0 {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had key: parameters->whiltelist but it did have an empty value"))
	} else {
		result["whitelist"] = whitelist
	}
	
	var data = fmt.Sprintf("%s",parameters.InterfaceString("data"))
	if (data) == "" {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION:Common: ValidateGeneric args had empty data: parameters->data keys: %s", parameter_keys))
	} else {
		result["data"] = data
	}

	var kind = fmt.Sprintf("%s",parameters.InterfaceString("reflect.ValueOf"))
	if (data) == "" {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had empty data: parameters->kind keys: %s", parameter_keys))
	} else {
		result["reflect.ValueOf"] = kind
	}

	var column_name = fmt.Sprintf("%s",parameters["column_name"])
	if (column_name) == "" {
		result["errors"] = append(result["errors"].([]error), fmt.Errorf("POTENTIAL SQL INJECTION: Common: ValidateGeneric args had the key: parameters->_column_name however had an empty value" ))
	} else {
		result["column_name"] = column_name
	}

	return result
}

func ContainsExactMatchz(args...map[string]interface{}) map[string]interface{} {
	//panic(args)
	var result = ValidateGeneric(args)

	if len(result["errors"].([]error)) > 0 {
		var arrayresults = make(map[string]interface{})
		count := 0
		for key, _ := range result {
			arrayresults[key + strconv.Itoa(count)] = key
			count++
		}
		return arrayresults
	}

	var whitelist, _ = result["whitelist"]
	var data, _ = result["data"].(string)
	var kind = result["reflect.ValueOf"]
	var columnName = result["column_name"]


	//panic(fmt.Sprintf("%s %s %s %s", whitelist, data, columnName, kind))
	var containsExactMatchErrors = ContainsExactMatch(whitelist.([]string), &data, columnName.(string), reflect.ValueOf(kind))
	if containsExactMatchErrors != nil {
		result["errors"] = append(result["errors"].([]error), containsExactMatchErrors...)
		panic("die here")
	}

	

	
	//var map := args[0]




	/*for _, array_value := range array {
		if array_value == *str {
			return nil
		}
	}*/

	//panic(result)
	var arrayresults = make(map[string]interface{})
	count := 0
	for key, _ := range result {
		arrayresults[key + strconv.Itoa(count)] = key
		count++
	}

	
    //errors = append(errors, fmt.Errorf("%s has value '%s' expected to have value in %s", label, (*str) , array))
	return arrayresults
}

func Validate(args...interface{}) []error {
	var errors []error 

	//panic("hellowworld")


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



