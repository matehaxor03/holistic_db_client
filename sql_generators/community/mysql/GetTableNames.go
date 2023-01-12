package mysql

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetTableNamesSQL(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
	if options == nil {
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

	var sql_command strings.Builder
	sql_command.WriteString("SHOW TABLES IN ")
	Box(options, &sql_command, database_name_escaped,"`","`")

	//sql_command.WriteString(database_name_escaped)
	sql_command.WriteString(";")
	sql_command_result := sql_command.String()
	return &sql_command_result, options, nil
}

