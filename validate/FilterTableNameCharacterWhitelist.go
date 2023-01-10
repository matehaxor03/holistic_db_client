package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type TableNameCharacterWhitelist struct {
	ValidateTableName func(table_name string) ([]error)
	GetValidateTableNameFunc func() (*func(string) []error)
}

func NewTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	table_name_character_whitelist := validation_constants.GetMySQLTableNameWhitelistCharacters()
	valid_table_names_cache := make(map[string]interface{})

	validateTableName := func(table_name string) ([]error) {
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

		whitelist_errors := validation_functions.WhitelistCharacters(table_name_character_whitelist, table_name, "Validator.ValidateTableName", "dao.Table.table_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_table_names_cache[table_name] = nil
		return nil
	}

	x := TableNameCharacterWhitelist {
		ValidateTableName: func(table_name string) ([]error) {
			return validateTableName(table_name)
		},
		GetValidateTableNameFunc: func() (*func(table_name string) []error) {
			function := validateTableName
			return &function
		},
	}

	return &x
}
