package class

import (
	json "github.com/matehaxor03/holistic_json/json"

)

type DatabaseNameCharacterWhitelist struct {
	GetDatabaseNameCharacterWhitelist func() (*json.Map)
}

func newDatabaseNameCharacterWhitelist() (*DatabaseNameCharacterWhitelist) {
	database_name_character_whitelist := GetMySQLDatabaseNameWhitelistCharacters()

	x := DatabaseNameCharacterWhitelist {
		GetDatabaseNameCharacterWhitelist: func() (*json.Map) {
			return &database_name_character_whitelist
		},
	}

	return &x
}
