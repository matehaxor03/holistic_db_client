package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	"fmt"
	"strings"
)

func BlackListString(m json.Map) []error {
	var errors []error
	map_values := m.GetObjectForMap("values")
	map_values = map_values.(*map[string]interface{})
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(*(map_values.(*map[string]interface{}))) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if _, found := (*(map_values.(*map[string]interface{})))[*str]; found {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func BlackListStringToUpper(m json.Map) []error {
	var errors []error
	map_values := m.GetObjectForMap("values")
	map_values = map_values.(*map[string]interface{})
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(*(map_values.(*map[string]interface{}))) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if _, found := (*(map_values.(*map[string]interface{})))[strings.ToUpper(*str)]; found {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetBlacklistStringFunc() *func(m json.Map) []error {
	function := BlackListString
	return &function
}

func GetBlacklistStringToUpperFunc() *func(m json.Map) []error {
	function := BlackListStringToUpper
	return &function
}


