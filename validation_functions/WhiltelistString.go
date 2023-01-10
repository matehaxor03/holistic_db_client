package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	"fmt"
)

func WhiteListString(m json.Map) []error {
	var errors []error
	map_values := m.GetObjectForMap("values")
	map_values = map_values.(*map[string]interface{})
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if  map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has nil map", *data_type, *label))
	} else if len(*(map_values.(*map[string]interface{}))) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if _, found := (*(map_values.(*map[string]interface{})))[*str]; !found {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: did not find value", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetWhitelistStringFunc() *func(m json.Map) []error {
	function := WhiteListString
	return &function
}
