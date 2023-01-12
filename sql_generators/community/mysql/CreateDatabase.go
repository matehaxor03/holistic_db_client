package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	"strings"
)

func GetCreateDatabaseSQL(verify *validate.Validator, database_name string, character_set *string, collate *string,  options *json.Map) (*string, *json.Map, []error) {
	var sql_command strings.Builder
	
	var errors []error
	if options == nil {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("creating_database", true)
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

	sql_command.WriteString("CREATE DATABASE ")
	Box(options, &sql_command, database_name_escaped,"`","`")

	//sql_command.WriteString(database_name_escaped)

	if character_set != nil {
		character_set_errors := verify.ValidateCharacterSet(*character_set)
		if character_set_errors != nil {
			return nil, nil, character_set_errors
		}
		sql_command.WriteString(" CHARACTER SET ")
		sql_command.WriteString(*character_set)
	}

	if collate != nil {
		collate_errors := verify.ValidateCollate(*collate)
		if collate_errors != nil {
			return nil, nil, collate_errors
		} 
		sql_command.WriteString(" COLLATE ")
		sql_command.WriteString(*collate)
	}
	sql_command.WriteString(";")

	if len(errors) > 0 {
		return nil, nil, errors
	}

	sql_command_result := sql_command.String()
	return &sql_command_result, options, nil
}