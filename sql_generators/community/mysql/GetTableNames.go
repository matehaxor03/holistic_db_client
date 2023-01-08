package mysql

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetTableNamesSQL(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	validation_errors := verify.ValidateDatabaseName(database_name)
	if validation_errors != nil {
		return nil, nil, validation_errors
	}

	database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
	if database_name_escaped_errors != nil {
		var errors []error
		errors = append(errors, database_name_escaped_errors)
		return nil, nil, errors
	}

	sql_command := "SHOW TABLES IN "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
	}
	return &sql_command, options, nil
}

