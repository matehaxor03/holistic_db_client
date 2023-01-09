package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type DatabaseNameCharacterWhitelist struct {
	GetDatabaseNameCharacterWhitelist func() (*json.Map)
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

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", database_name)
		parameters.SetMap("values", &database_name_character_whitelist)
		parameters.SetStringValue("label", "Validator.ValidateDatabaseName")
		parameters.SetStringValue("data_type", "dao.Database.database_name")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
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
		GetDatabaseNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetMySQLDatabaseNameWhitelistCharacters()
			return &v
		},
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
