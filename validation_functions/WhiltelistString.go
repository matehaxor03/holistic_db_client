package validation_functions

import (
	"fmt"
)

func WhiteListString(map_values map[string]interface{}, value string, label string, data_type string) []error {
	var errors []error

	if _, found := map_values[value]; !found {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: did not find value", data_type, label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetWhitelistStringFunc() *func(map_values map[string]interface{}, value string, label string, data_type string) []error {
	function := WhiteListString
	return &function
}
