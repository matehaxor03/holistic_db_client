package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type RepositoryNameCharacterWhitelist struct {
	GetRepositoryNameCharacterWhitelist func() (*json.Map)
}

func NewRepositoryNameCharacterWhitelist() (*RepositoryNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidRepositoryNameCharacters()

	x := RepositoryNameCharacterWhitelist {
		GetRepositoryNameCharacterWhitelist: func() (*json.Map) {
			return &valid_characters
		},
	}

	return &x
}
