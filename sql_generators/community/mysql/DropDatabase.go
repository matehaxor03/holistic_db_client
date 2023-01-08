package mysql

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetDropDatabaseIfExistsSQL(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("deleting_database", true)
		options.SetBoolValue("read_no_records", true)
	}

	validation_errors := verify.ValidateDatabaseName(database_name)
	if validation_errors != nil {
		return nil, nil, validation_errors
	}

	database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
	if database_name_escaped_errors != nil {
		errors = append(errors, database_name_escaped_errors)
		return nil, nil, errors
	}

	sql_command := "DROP DATABASE IF EXISTS "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
	}
	return &sql_command, options, nil
}


func GetDropDatabaseSQL(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("deleting_database", true)
		options.SetBoolValue("read_no_records", true)
	}

	validation_errors := verify.ValidateDatabaseName(database_name)
	if validation_errors != nil {
		return nil, nil, validation_errors
	}

	database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
	if database_name_escaped_errors != nil {
		errors = append(errors, database_name_escaped_errors)
		return nil, nil, errors
	}

	sql_command := "DROP DATABASE "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
	}
	return &sql_command, options, nil
}

