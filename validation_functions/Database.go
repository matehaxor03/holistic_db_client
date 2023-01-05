package validation_functions

import (
    json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

func ValidateDatabaseColumnName(value string) []error {
	var errors []error
	column_name_params := json.NewMapValue()
	column_name_params.SetObjectForMap("values", validation_constants.GetMySQLColumnNameWhitelistCharacters())
	column_name_params.SetStringValue("value", value)
	column_name_params.SetStringValue("label", "column_name")
	column_name_params.SetStringValue("data_type", "Table")

	column_name_errors := WhitelistCharacters(column_name_params)
	if column_name_errors != nil {
		errors = append(errors, column_name_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}