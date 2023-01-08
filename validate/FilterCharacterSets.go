package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
)

type CharacterSetWordWhitelist struct {
	GetCharacterSetWordWhitelist func() (*json.Map)
	ValidateCharacterSet func(character_set string) ([]error)
}

func NewCharacterSetWordWhitelist() (*CharacterSetWordWhitelist) {
	valid_words := validation_constants.GET_CHARACTER_SETS()
	cache := make(map[string]interface{})

	x := CharacterSetWordWhitelist {
		GetCharacterSetWordWhitelist: func() (*json.Map) {
			v := validation_constants.GET_CHARACTER_SETS()
			return &v
		},
		ValidateCharacterSet: func(character_set string) ([]error) {
			if _, found := cache[character_set]; found {
				return nil
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", character_set)
			parameters.SetMap("values", &valid_words)
			parameters.SetStringValue("label", "Validator.ValidateCharacterSet")
			parameters.SetStringValue("data_type", "database.character_set")
			whitelist_errors := validation_functions.WhiteListString(parameters)
			if whitelist_errors != nil {
				return whitelist_errors
			}

			cache[character_set] = nil
			return nil
		},
	}

	return &x
}
