package validate

import(
	json "github.com/matehaxor03/holistic_json/json"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type Validator struct {
	ValidateTableName func(table_name string) ([]error)
	ValidateDatabaseName func(database_name string) ([]error)
	ValidateColumnName  func(column_name string) ([]error)

	GetDatabaseReservedWordsBlackList func() *json.Map 
	GetDatabaseNameWhitelistCharacters func() *json.Map
	GetColumnNameCharacterWhitelist func() *json.Map
	GetTableNameCharacterWhitelist func() *json.Map
	
	GetUsernameCharacterWhitelist func() *json.Map 
	GetBranchNameCharacterWhitelist func() *json.Map 
	GetRepositoryNameCharacterWhitelist func() *json.Map 
	GetRepositoryAccountNameCharacterWhitelist func() *json.Map 

	GetDomainNameCharacterWhitelist func() *json.Map 
	GetPortNumberCharacterWhitelist func() *json.Map 

	GetCharacterSetWordWhitelist func() *json.Map
	GetCollateWordWhitelist func() *json.Map
}

func NewValidator() (*Validator) {
	database_reserved_words_blacklist := NewDatabaseReservedWordsBlackList()
	valid_database_name_characters := NewDatabaseNameCharacterWhitelist()
	valid_table_name_characters := NewTableNameCharacterWhitelist()
	valid_column_name_characters := NewColumnNameCharacterWhitelist()
	
	valid_username_characters := NewUsernameCharacterWhitelist()
	valid_branch_name_characters := NewBranchNameCharacterWhitelist()
	valid_repository_name_characters := NewRepositoryNameCharacterWhitelist()
	valid_repository_account_name_characters := NewRepositoryAccountNameCharacterWhitelist()
	valid_domain_name_characters := NewDomainNameCharacterWhitelist()
	valid_port_number_characters := NewPortNumberCharacterWhitelist()

	valid_character_set_words := NewCharacterSetWordWhitelist()
	valid_collate_words := NewCollateWordWhitelist()

	

	x := Validator {
		GetDatabaseNameWhitelistCharacters: func() *json.Map {
			return valid_database_name_characters.GetDatabaseNameCharacterWhitelist()
		},
		GetTableNameCharacterWhitelist: func() *json.Map {
			return valid_table_name_characters.GetTableNameCharacterWhitelist()
		},
		GetColumnNameCharacterWhitelist: func() *json.Map {
			return valid_column_name_characters.GetColumnNameCharacterWhitelist()
		},
		GetUsernameCharacterWhitelist: func() *json.Map {
			return valid_username_characters.GetUsernameCharacterWhitelist()
		},
		GetDatabaseReservedWordsBlackList: func() *json.Map {
			return database_reserved_words_blacklist.GetDatabaseReservedWordsBlackList()
		},
		GetBranchNameCharacterWhitelist: func() *json.Map {
			return valid_branch_name_characters.GetBranchNameCharacterWhitelist()
		},
		GetRepositoryNameCharacterWhitelist: func() *json.Map {
			return valid_repository_name_characters.GetRepositoryNameCharacterWhitelist()
		},
		GetRepositoryAccountNameCharacterWhitelist: func() *json.Map {
			return valid_repository_account_name_characters.GetRepositoryAccountNameCharacterWhitelist()
		},
		GetDomainNameCharacterWhitelist: func() *json.Map {
			return valid_domain_name_characters.GetDomainNameCharacterWhitelist()
		},
		GetPortNumberCharacterWhitelist: func() *json.Map {
			return valid_port_number_characters.GetPortNumberCharacterWhitelist()
		},
		GetCharacterSetWordWhitelist: func() *json.Map {
			return valid_character_set_words.GetCharacterSetWordWhitelist()
		},
		GetCollateWordWhitelist: func() *json.Map {
			return valid_collate_words.GetCollateWordWhitelist()
		},
		ValidateTableName: func(table_name string) ([]error) {
			var errors []error
			if table_name == "" {
				errors = append(errors, fmt.Errorf("table_name is empty"))
			}

			if len(table_name) < 2 {
				errors = append(errors, fmt.Errorf("table_name is too short must be at least 2 characters"))
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", table_name)
			parameters.SetMap("values", valid_table_name_characters.GetTableNameCharacterWhitelist())
			parameters.SetStringValue("label", "Validator.ValidateTableName")
			parameters.SetStringValue("data_type", "dao.Table.table_name")
			table_name_character_whitelist_errors := validation_functions.WhitelistCharacters(parameters)
			if table_name_character_whitelist_errors != nil {
				errors = append(errors, table_name_character_whitelist_errors...)
			}

			parameters.SetMap("values", database_reserved_words_blacklist.GetDatabaseReservedWordsBlackList())
			reserved_words_errors := validation_functions.BlackListStringToUpper(parameters)
			if reserved_words_errors != nil {
				errors = append(errors, reserved_words_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		ValidateColumnName: func(column_name string) ([]error) {
			var errors []error
			
			if column_name == "" {
				errors = append(errors, fmt.Errorf("column_name is empty"))
			}

			if len(column_name) < 2 {
				errors = append(errors, fmt.Errorf("column_name is too short must be at least 2 characters"))
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", column_name)
			parameters.SetMap("values", valid_column_name_characters.GetColumnNameCharacterWhitelist())
			parameters.SetStringValue("label", "Validator.ValidateTableName")
			parameters.SetStringValue("data_type", "dao.Table.table_name")
			table_name_character_whitelist_errors := validation_functions.WhitelistCharacters(parameters)
			if table_name_character_whitelist_errors != nil {
				errors = append(errors, table_name_character_whitelist_errors...)
			}

			parameters.SetMap("values", database_reserved_words_blacklist.GetDatabaseReservedWordsBlackList())
			reserved_words_errors := validation_functions.BlackListStringToUpper(parameters)
			if reserved_words_errors != nil {
				errors = append(errors, reserved_words_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		ValidateDatabaseName: func(database_name string) ([]error) {
			var errors []error
			
			if database_name == "" {
				errors = append(errors, fmt.Errorf("database_name is empty"))
			}

			if len(database_name) < 2 {
				errors = append(errors, fmt.Errorf("database_name is too short must be at least 2 characters"))
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", database_name)
			parameters.SetMap("values", valid_database_name_characters.GetDatabaseNameCharacterWhitelist())
			parameters.SetStringValue("label", "Validator.ValidateDatabaseName")
			parameters.SetStringValue("data_type", "dao.Database.database_name")
			table_name_character_whitelist_errors := validation_functions.WhitelistCharacters(parameters)
			if table_name_character_whitelist_errors != nil {
				errors = append(errors, table_name_character_whitelist_errors...)
			}

			parameters.SetMap("values", database_reserved_words_blacklist.GetDatabaseReservedWordsBlackList())
			reserved_words_errors := validation_functions.BlackListStringToUpper(parameters)
			if reserved_words_errors != nil {
				errors = append(errors, reserved_words_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			return nil
		},
	}

	return &x
}
