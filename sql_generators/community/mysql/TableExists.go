package mysql

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetCheckTableExistsSQL(verify *validate.Validator, table_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	validation_errors := verify.ValidateTableName(table_name)
	if validation_errors != nil {
		return nil, nil, validation_errors
	}

	table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
	if table_name_escaped_errors != nil {
		errors = append(errors, table_name_escaped_errors)
		return nil, nil, errors
	}
	
	var sql_command strings.Builder
	sql_command.WriteString("SELECT 0 FROM ")
	box(options, &sql_command, table_name_escaped,"`","`")
	sql_command.WriteString(" LIMIT 1;")

	sql_command_result := sql_command.String()
	return &sql_command_result, options, nil
}

