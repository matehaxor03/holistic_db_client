package dao

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getDropTableSQLMySQL(struct_type string, table *Table, drop_table_if_exists *bool, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(table) {
		errors = append(errors, fmt.Errorf("table is nil"))
		return nil, nil, errors
	} else {
		validation_errors := table.Validate()
		if len(validation_errors) > 0 {
			return nil, nil, validation_errors 
		}
	}

	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	if common.IsNil(drop_table_if_exists) {
		temp_drop_table_if_exists := true
		drop_table_if_exists = &temp_drop_table_if_exists
	}

	temp_table_name, temp_table_name_errors := table.GetTableName()
	if temp_table_name_errors != nil {
		return  nil, nil, temp_table_name_errors 
	}

	table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_errors != nil {
		errors = append(errors, table_name_escaped_errors)
		return  nil, nil, errors 
	}

	sql_command := "DROP TABLE "
	if *drop_table_if_exists {
		sql_command += "IF EXISTS "
	}

	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
	}

	return &sql_command, options, nil
}

