package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type DatabaseNameCharacterWhitelist struct {
	GetDatabaseNameCharacterWhitelist func() (*json.Map)
}

func NewDatabaseNameCharacterWhitelist() (*DatabaseNameCharacterWhitelist) {
	database_name_character_whitelist := validation_constants.GetMySQLDatabaseNameWhitelistCharacters()

	x := DatabaseNameCharacterWhitelist {
		GetDatabaseNameCharacterWhitelist: func() (*json.Map) {
			return &database_name_character_whitelist
		},
	}

	return &x
}
