package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	"fmt"
)

func ValidateDatabaseName(database_name string) []error {
	var errors []error
	column_name_params := json.NewMapValue()
	column_name_params.SetStringValue("value", database_name)
	column_name_params.SetStringValue("label", "database_name")
	column_name_params.SetStringValue("data_type", "validation_functions.ValidateDatabaseName(database_name)")

	if len(database_name) < 2 {
		errors = append(errors, fmt.Errorf("database_name length is less than 2 characters. validation_functions.ValidateDatabaseName(database_name)"))
	}


	column_name_params.SetObjectForMap("values", validation_constants.GetMySQLColumnNameWhitelistCharacters())
	column_name_errors := WhitelistCharacters(column_name_params)
	if column_name_errors != nil {
		errors = append(errors, column_name_errors...)
	}

	column_name_params.SetObjectForMap("values", validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords())
	reserved_words_errors := BlackListStringToUpper(column_name_params)
	if reserved_words_errors != nil {
		errors = append(errors, reserved_words_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}