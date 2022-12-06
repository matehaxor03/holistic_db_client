package class

import (
	json "github.com/matehaxor03/holistic_json/json"
)

type TableNameCharacterWhitelist struct {
	GetTableNameCharacterWhitelist func() (*json.Map)
}

func newTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	table_name_character_whitelist := GetMySQLTableNameWhitelistCharacters()

	x := TableNameCharacterWhitelist {
		GetTableNameCharacterWhitelist: func() (*json.Map) {
			return &table_name_character_whitelist
		},
	}

	return &x
}
