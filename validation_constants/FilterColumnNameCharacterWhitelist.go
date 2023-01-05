package validation_constants

import(
	json "github.com/matehaxor03/holistic_json/json"
)

type ColumnNameCharacterWhitelist struct {
	GetColumnNameCharacterWhitelist func() (*json.Map)
}

func NewColumnNameCharacterWhitelist() (*ColumnNameCharacterWhitelist) {
	column_name_character_whitelist := GetMySQLColumnNameWhitelistCharacters()

	x := ColumnNameCharacterWhitelist {
		GetColumnNameCharacterWhitelist: func() (*json.Map) {
			return &column_name_character_whitelist
		},
	}

	return &x
}
