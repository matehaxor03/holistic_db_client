package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type BranchNameCharacterWhitelist struct {
	GetBranchNameCharacterWhitelist func() (*json.Map)
}

func NewBranchNameCharacterWhitelist() (*BranchNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidBranchNameCharacters()

	x := BranchNameCharacterWhitelist {
		GetBranchNameCharacterWhitelist: func() (*json.Map) {
			return &valid_characters
		},
	}

	return &x
}
