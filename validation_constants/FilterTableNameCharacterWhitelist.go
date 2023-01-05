package validation_constants

import (
	json "github.com/matehaxor03/holistic_json/json"
)

type TableNameCharacterWhitelist struct {
	GetTableNameCharacterWhitelist func() (*json.Map)
}

func NewTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	table_name_character_whitelist := GetMySQLTableNameWhitelistCharacters()

	x := TableNameCharacterWhitelist {
		GetTableNameCharacterWhitelist: func() (*json.Map) {
			return &table_name_character_whitelist
		},
	}

	return &x
}
