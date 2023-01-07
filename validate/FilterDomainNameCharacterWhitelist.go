package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type DomainNameCharacterWhitelist struct {
	GetDomainNameCharacterWhitelist func() (*json.Map)
}

func NewDomainNameCharacterWhitelist() (*DomainNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidDomainNameCharacters()

	x := DomainNameCharacterWhitelist {
		GetDomainNameCharacterWhitelist: func() (*json.Map) {
			return &valid_characters
		},
	}

	return &x
}
