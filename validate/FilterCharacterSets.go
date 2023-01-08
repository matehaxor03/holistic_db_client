package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type CharacterSetWordWhitelist struct {
	GetCharacterSetWordWhitelist func() (*json.Map)
}

func NewCharacterSetWordWhitelist() (*CharacterSetWordWhitelist) {
	//valid_characters := validation_constants.GET_CHARACTER_SETS()

	x := CharacterSetWordWhitelist {
		GetCharacterSetWordWhitelist: func() (*json.Map) {
			v := validation_constants.GET_CHARACTER_SETS()
			return &v
		},
	}

	return &x
}
