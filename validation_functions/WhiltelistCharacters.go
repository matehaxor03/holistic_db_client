package validation_functions

import (
	"fmt"
)

func WhitelistCharacters(map_values map[string]interface{}, value string, label string, data_type string) []error {
	var errors []error

	var invalid_letters []string
	for _, letter_rune := range value {
		letter_string := string(letter_rune)

		if _, found := map_values[letter_string]; !found {
			invalid_letters = append(invalid_letters, letter_string)
		}
	}

	if len(invalid_letters) > 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has invalid character(s): %s", data_type, label, invalid_letters))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func GetWhitelistCharactersFunc() (*func(map_values map[string]interface{}, value string, label string, data_type string) []error) {
	funcValue := WhitelistCharacters
	return &funcValue
}