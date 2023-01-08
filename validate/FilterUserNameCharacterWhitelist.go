package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type UsernameCharacterWhitelist struct {
	GetUsernameCharacterWhitelist func() (*json.Map)
}

func NewUsernameCharacterWhitelist() (*UsernameCharacterWhitelist) {
	//valid_characters := validation_constants.GetValidUsernameCharacters()

	x := UsernameCharacterWhitelist {
		GetUsernameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidUsernameCharacters()
			return &v
		},
	}

	return &x
}
