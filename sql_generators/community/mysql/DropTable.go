package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_validator/validate"
	"strings"
)

type DropTableSQL struct {
	GetDropTableIfExistsSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropTableSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
}

func newDropTableSQL() (*DropTableSQL) {
	get_drop_table_if_exists_sql := func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error

		validation_errors := verify.ValidateTableName(table_name)
		if validation_errors != nil {
			return nil, options, validation_errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return  nil, options, errors 
		}


		var sql_command strings.Builder
		sql_command.WriteString("DROP TABLE IF EXISTS ")
		Box(&sql_command, table_name_escaped,"`","`")
		sql_command.WriteString(";")

		return &sql_command, options, nil
	}

	get_drop_table_sql := func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error
		validation_errors := verify.ValidateTableName(table_name)
		if validation_errors != nil {
			return nil, options, validation_errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return  nil, options, errors 
		}


		var sql_command strings.Builder
		sql_command.WriteString("DROP TABLE ")
		Box(&sql_command, table_name_escaped,"`","`")
		sql_command.WriteString(";")

		return &sql_command, options, nil
	}

	return &DropTableSQL{
		GetDropTableIfExistsSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_drop_table_if_exists_sql(verify, table_name, options)
		},
		GetDropTableSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_drop_table_sql(verify, table_name, options)
		},
	}
}
