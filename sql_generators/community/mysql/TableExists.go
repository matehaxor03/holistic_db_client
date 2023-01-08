package mysql

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetCheckTableExistsSQL(verify *validate.Validator, struct_type string, table_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	if common.IsNil(table_name) {
		errors = append(errors, fmt.Errorf("table_name is nil"))
		return nil, nil, errors
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
	
	sql_command := fmt.Sprintf("SELECT 0 FROM ")
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`", table_name_escaped)
	}
	sql_command += " LIMIT 1;"

	return &sql_command, options, nil
}
