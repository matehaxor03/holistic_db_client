package db_client

import(
	json "github.com/matehaxor03/holistic_json/json"
)

type ColumnNameCharacterWhitelist struct {
	GetColumnNameCharacterWhitelist func() (*json.Map)
}

func newColumnNameCharacterWhitelist() (*ColumnNameCharacterWhitelist) {
	column_name_character_whitelist := GetMySQLColumnNameWhitelistCharacters()

	x := ColumnNameCharacterWhitelist {
		GetColumnNameCharacterWhitelist: func() (*json.Map) {
			return &column_name_character_whitelist
		},
	}

	return &x
}
