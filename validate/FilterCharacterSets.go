package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
)

type CharacterSetWordWhitelist struct {
	ValidateCharacterSet func(character_set string) ([]error)
	GetValidateCharacterSetFunc func() (*func(string) []error)
}

func NewCharacterSetWordWhitelist() (*CharacterSetWordWhitelist) {
	valid_words := validation_constants.GET_CHARACTER_SETS()
	cache := make(map[string]interface{})

	validateCharacterSet := func(character_set string) ([]error) {
		if _, found := cache[character_set]; found {
			return nil
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", character_set)
		parameters.SetObjectForMap("values", &valid_words)
		parameters.SetStringValue("label", "Validator.ValidateCharacterSet")
		parameters.SetStringValue("data_type", "database.character_set")
		whitelist_errors := validation_functions.WhiteListString(parameters)
		if whitelist_errors != nil {
			return whitelist_errors
		}

		cache[character_set] = nil
		return nil
	}

	x := CharacterSetWordWhitelist {
		ValidateCharacterSet: func(character_set string) ([]error) {
			return validateCharacterSet(character_set)
		},
		GetValidateCharacterSetFunc: func() (*func(character_set string) []error) {
			function := validateCharacterSet
			return &function
		},
	}

	return &x
}
