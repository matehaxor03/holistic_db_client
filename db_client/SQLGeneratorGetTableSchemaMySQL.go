package db_client

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getTableSchemaSQLMySQL(struct_type string, table *Table, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(table) {
		errors = append(errors, fmt.Errorf("table is nil"))
		return nil, nil, errors
	} else {
		validation_errors := table.Validate()
		if validation_errors != nil {
			return nil, nil, validation_errors
		}
	}

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)
	}

	temp_table_name, temp_table_name_errors := table.GetTableName()
	if temp_table_name_errors != nil {
		return nil, nil, temp_table_name_errors
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
		return nil, nil, errors
	}

	sql_command := "SHOW FULL COLUMNS FROM "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
	}

	return &sql_command, options, nil
}
