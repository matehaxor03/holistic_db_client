package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	"strings"
)

type DropDatabaseSQL struct {
	GetDropDatabaseIfExistsSQL func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error)
	GetDropDatabaseSQL func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error)
}

func newDropDatabaseSQL() (*DropDatabaseSQL) {
	get_drop_database_if_exists_sql := func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
		var errors []error
		if options == nil {
			options := json.NewMap()
			options.SetBoolValue("use_file", true)
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

		var sql_command strings.Builder
		sql_command.WriteString("DROP DATABASE IF EXISTS ")
		Box(options, &sql_command, database_name_escaped,"`","`")

		
		sql_command.WriteString(";")
		sql_command_result := sql_command.String()
		return &sql_command_result, options, nil
	}

	get_drop_database_sql := func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
		var errors []error
		if options == nil {
			options := json.NewMap()
			options.SetBoolValue("use_file", true)
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

		var sql_command strings.Builder
		sql_command.WriteString("DROP DATABASE ")
		Box(options, &sql_command, database_name_escaped,"`","`")

		sql_command.WriteString(";")
		sql_command_result := sql_command.String()
		return &sql_command_result, options, nil
	}

	return &DropDatabaseSQL{
		GetDropDatabaseIfExistsSQL: func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
			return get_drop_database_if_exists_sql(verify, database_name, options)
		},
		GetDropDatabaseSQL: func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
			return get_drop_database_sql(verify, database_name, options)
		},
	}
}
