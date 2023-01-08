package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type RepositoryAccountNameCharacterWhitelist struct {
	GetRepositoryAccountNameCharacterWhitelist func() (*json.Map)
}

func NewRepositoryAccountNameCharacterWhitelist() (*RepositoryAccountNameCharacterWhitelist) {
	//valid_characters := validation_constants.GetValidRepositoryAccountNameCharacters()

	x := RepositoryAccountNameCharacterWhitelist {
		GetRepositoryAccountNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidRepositoryAccountNameCharacters()
			return &v
		},
	}

	return &x
}
