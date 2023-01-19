package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_validator/validate"
	"strings"
)

type CreateDatabaseSQL struct {
	GetCreateDatabaseSQL func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options json.Map) (*strings.Builder, json.Map, []error)
}

func newCreateDatabaseSQL() (*CreateDatabaseSQL) {
	get_create_database_sql := func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options json.Map) (*strings.Builder, json.Map, []error) {
		var sql_command strings.Builder
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

		sql_command.WriteString("CREATE DATABASE ")
		Box(&sql_command, database_name_escaped,"`","`")


		if character_set != nil {
			character_set_errors := verify.ValidateCharacterSet(*character_set)
			if character_set_errors != nil {
				return nil, options, character_set_errors
			}
			sql_command.WriteString(" CHARACTER SET ")
			sql_command.WriteString(*character_set)
		}

		if collate != nil {
			collate_errors := verify.ValidateCollate(*collate)
			if collate_errors != nil {
				return nil, options, collate_errors
			} 
			sql_command.WriteString(" COLLATE ")
			sql_command.WriteString(*collate)
		}
		sql_command.WriteString(";")

		if len(errors) > 0 {
			return nil, options, errors
		}

		return &sql_command, options, nil
	}

	return &CreateDatabaseSQL{
		GetCreateDatabaseSQL: func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options json.Map) (*strings.Builder, json.Map, []error) {	
			return get_create_database_sql(verify, database_name, character_set, collate, options)
		},
	}
}