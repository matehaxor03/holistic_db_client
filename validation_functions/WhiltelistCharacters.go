package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	"fmt"
)

func WhitelistCharacters(m json.Map) []error {
	var errors []error
	map_values := m.GetObjectForMap("values")
	map_values = map_values.(*map[string]interface{})
	
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	 if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has nil map", *data_type, *label))
	} else if len(*(map_values.(*map[string]interface{}))) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	var invalid_letters []string
	for _, letter_rune := range *str {
		letter_string := string(letter_rune)

		if _, found := (*(map_values.(*map[string]interface{})))[letter_string]; !found {
			invalid_letters = append(invalid_letters, letter_string)
		}
	}

	if len(invalid_letters) > 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has invalid character(s): %s", *data_type, *label, invalid_letters))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func GetWhitelistCharactersFunc() (*func(m json.Map) []error) {
	funcValue := WhitelistCharacters
	return &funcValue
}