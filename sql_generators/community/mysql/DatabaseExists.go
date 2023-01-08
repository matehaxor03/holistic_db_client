package mysql

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetDatabaseExistsSQL(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("checking_database_exists", true)
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

	sql_command := "USE "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
	}

	return &sql_command, options, nil
}

