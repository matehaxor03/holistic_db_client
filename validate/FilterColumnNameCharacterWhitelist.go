package validate

import(
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"

)

type ColumnNameCharacterWhitelist struct {
	GetColumnNameCharacterWhitelist func() (*json.Map)
}

func NewColumnNameCharacterWhitelist() (*ColumnNameCharacterWhitelist) {
	column_name_character_whitelist := validation_constants.GetMySQLColumnNameWhitelistCharacters()

	x := ColumnNameCharacterWhitelist {
		GetColumnNameCharacterWhitelist: func() (*json.Map) {
			return &column_name_character_whitelist
		},
	}

	return &x
}
