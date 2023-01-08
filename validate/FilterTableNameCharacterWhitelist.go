package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type TableNameCharacterWhitelist struct {
	GetTableNameCharacterWhitelist func() (*json.Map)
	ValidateTableName func(table_name string) ([]error)
}

func NewTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	table_name_character_whitelist := validation_constants.GetMySQLTableNameWhitelistCharacters()
	valid_table_names_cache := make(map[string]interface{})


	x := TableNameCharacterWhitelist {
		GetTableNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetMySQLTableNameWhitelistCharacters()
			return &v
		},
		ValidateTableName: func(table_name string) ([]error) {
			if _, found := valid_table_names_cache[table_name]; found {
				return nil
			}
			
			var errors []error
			if table_name == "" {
				errors = append(errors, fmt.Errorf("table_name is empty"))
			}

			if len(table_name) < 2 {
				errors = append(errors, fmt.Errorf("table_name is too short must be at least 2 characters"))
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", table_name)
			parameters.SetMap("values", &table_name_character_whitelist)
			parameters.SetStringValue("label", "Validator.ValidateTableName")
			parameters.SetStringValue("data_type", "dao.Table.table_name")
			whitelist_errors := validation_functions.WhitelistCharacters(parameters)
			if whitelist_errors != nil {
				errors = append(errors, whitelist_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			valid_table_names_cache[table_name] = nil
			return nil
		},
	}

	return &x
}
