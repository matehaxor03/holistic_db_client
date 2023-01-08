package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type TableNameCharacterWhitelist struct {
	GetTableNameCharacterWhitelist func() (*json.Map)
}

func NewTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	//table_name_character_whitelist := validation_constants.GetMySQLTableNameWhitelistCharacters()

	x := TableNameCharacterWhitelist {
		GetTableNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetMySQLTableNameWhitelistCharacters()
			return &v
		},
	}

	return &x
}
