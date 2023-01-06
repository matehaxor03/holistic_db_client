package dao

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getDropTableSQLMySQL(struct_type string, table_name string, drop_table_if_exists bool, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	if common.IsNil(options) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
	if table_name_escaped_errors != nil {
		errors = append(errors, table_name_escaped_errors)
		return  nil, nil, errors 
	}

	sql_command := "DROP TABLE "
	if drop_table_if_exists {
		sql_command += "IF EXISTS "
	}

	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
	}

	return &sql_command, options, nil
}
