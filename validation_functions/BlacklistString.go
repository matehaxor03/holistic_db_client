package validation_functions

import (
	"fmt"
	"strings"
)

func BlackListString(map_values map[string]interface{}, value string, label string, data_type string) []error {
	var errors []error
	if _, found := map_values[value]; found {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString found: %s", data_type, label, value))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func BlackListStringToUpper(map_values map[string]interface{}, value string, label string, data_type string) []error {
	var errors []error

	if _, found := map_values[strings.ToUpper(value)]; found {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: found value: %s", data_type, label, value))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetBlacklistStringFunc() *func(map_values map[string]interface{}, value string, label string, data_type string) []error {
	function := BlackListString
	return &function
}

func GetBlacklistStringToUpperFunc() *func(map_values map[string]interface{}, value string, label string, data_type string) []error {
	function := BlackListStringToUpper
	return &function
}


