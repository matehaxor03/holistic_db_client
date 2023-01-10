package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type DatabaseNameCharacterWhitelist struct {
	ValidateDatabaseName func(database_name string) ([]error)
	GetValidateDatabaseNameFunc func() (*func(string) []error)
}

func NewDatabaseNameCharacterWhitelist() (*DatabaseNameCharacterWhitelist) {
	database_name_character_whitelist := validation_constants.GetMySQLDatabaseNameWhitelistCharacters()
	valid_database_names_cache := make(map[string]interface{})

	validateDatabaseName := func(database_name string) ([]error) {
		if _, found := valid_database_names_cache[database_name]; found {
			return nil
		}
		
		var errors []error
		if database_name == "" {
			errors = append(errors, fmt.Errorf("database_name is empty"))
		}

		if len(database_name) < 2 {
			errors = append(errors, fmt.Errorf("database_name is too short must be at least 2 characters"))
		}

		whitelist_errors := validation_functions.WhitelistCharacters(database_name_character_whitelist, database_name, "Validator.ValidateDatabaseName",  "dao.Database.database_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_database_names_cache[database_name] = nil
		return nil
	}

	x := DatabaseNameCharacterWhitelist {
		ValidateDatabaseName: func(database_name string) ([]error) {
			return validateDatabaseName(database_name)
		},
		GetValidateDatabaseNameFunc: func() (*func(database_name string) []error) {
			function := validateDatabaseName
			return &function
		},
	}

	return &x
}
