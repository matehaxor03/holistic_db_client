package validate

import(
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type ColumnNameCharacterWhitelist struct {
	ValidateColumnName func(column_name string) ([]error)
	GetValidateColumnNameFunc func() (*func(string) []error)
}

func NewColumnNameCharacterWhitelist() (*ColumnNameCharacterWhitelist) {
	column_name_character_whitelist := validation_constants.GetMySQLColumnNameWhitelistCharacters()
	valid_column_names_cache := make(map[string]interface{})

	validateColumnName := func(column_name string) ([]error) {
		if _, found := valid_column_names_cache[column_name]; found {
			return nil
		}
		
		var errors []error
		
		if column_name == "" {
			errors = append(errors, fmt.Errorf("column_name is empty"))
		}

		if len(column_name) < 2 {
			errors = append(errors, fmt.Errorf("column_name is too short must be at least 2 characters"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(column_name_character_whitelist, column_name, "Validator.ValidateTableName", "dao.Table.table_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_column_names_cache[column_name] = nil
		return nil
	}


	x := ColumnNameCharacterWhitelist {
		ValidateColumnName: func(column_name string) ([]error) {
			return validateColumnName(column_name)
		},
		GetValidateColumnNameFunc: func() (*func(column_name string) []error) {
			function := validateColumnName
			return &function
		},
	}

	return &x
}
