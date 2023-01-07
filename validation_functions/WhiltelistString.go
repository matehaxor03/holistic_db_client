package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	"fmt"
)

func WhiteListString(m *json.Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if  map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has nil map", *data_type, *label))
	} else if len(map_values.GetKeys()) == 0 {
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

	if !map_values.HasKey(*str) {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: did not find value", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetWhitelistStringFunc() *func(m *json.Map) []error {
	function := WhiteListString
	return &function
}
