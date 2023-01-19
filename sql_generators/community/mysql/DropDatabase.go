package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_validator/validate"
	"strings"
)

type DropDatabaseSQL struct {
	GetDropDatabaseIfExistsSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropDatabaseSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
}

func newDropDatabaseSQL() (*DropDatabaseSQL) {
	get_drop_database_if_exists_sql := func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error
	
		validation_errors := verify.ValidateDatabaseName(database_name)
		if validation_errors != nil {
			return nil, options, validation_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, options, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString("DROP DATABASE IF EXISTS ")
		Box(&sql_command, database_name_escaped,"`","`")

		
		sql_command.WriteString(";")
		return &sql_command, options, nil
	}

	get_drop_database_sql := func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error

		validation_errors := verify.ValidateDatabaseName(database_name)
		if validation_errors != nil {
			return nil, options, validation_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, options, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString("DROP DATABASE ")
		Box(&sql_command, database_name_escaped,"`","`")

		sql_command.WriteString(";")
		return &sql_command, options, nil
	}

	return &DropDatabaseSQL{
		GetDropDatabaseIfExistsSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_drop_database_if_exists_sql(verify, database_name, options)
		},
		GetDropDatabaseSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_drop_database_sql(verify, database_name, options)
		},
	}
}
