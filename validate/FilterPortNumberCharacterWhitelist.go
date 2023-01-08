package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type PortNumberCharacterWhitelist struct {
	GetPortNumberCharacterWhitelist func() (*json.Map)
}

func NewPortNumberCharacterWhitelist() (*PortNumberCharacterWhitelist) {
	//valid_characters := validation_constants.GetValidPortNumberCharacters()

	x := PortNumberCharacterWhitelist {
		GetPortNumberCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidPortNumberCharacters()
			return &v
		},
	}

	return &x
}
