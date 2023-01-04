package db_client

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getTableSchemaAdditionalSQLMySQL(struct_type string, table *Table, options *json.Map) (*string, *json.Map, []error) {
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

	temp_database, temp_database_errors := table.GetDatabase()
	if temp_database_errors != nil {
		return nil, nil, temp_database_errors
	} else if common.IsNil(temp_database) {
		errors = append(errors, fmt.Errorf("database is nil"))
		return nil, nil, errors
	}

	temp_database_name, temp_database_name_errors := temp_database.GetDatabaseName()
	if temp_database_name_errors != nil {
		return nil, nil, temp_database_name_errors
	} else if common.IsNil(temp_database_name) {
		errors = append(errors, fmt.Errorf("database_name is nil"))
		return nil, nil, errors
	}


	database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
	if database_name_escaped_errors != nil {
		errors = append(errors, database_name_escaped_errors)
		return nil, nil, errors
	} else if common.IsNil(database_name_escaped) {
		errors = append(errors, fmt.Errorf("database_name_escaped is nil"))
		return nil, nil, errors
	}
	
	temp_table_name, temp_table_name_errors := table.GetTableName()
	if temp_table_name_errors != nil {
		return nil, nil, temp_table_name_errors
	} else if common.IsNil(temp_table_name) {
		errors = append(errors, fmt.Errorf("table_name is nil"))
		return nil, nil, errors
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
		return nil, nil, errors
	} else if common.IsNil(table_name_escaped) {
		errors = append(errors, fmt.Errorf("table_name_escaped is nil"))
		return nil, nil, errors
	}

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)
	}

	sql_command := "SHOW TABLE STATUS FROM "
		
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s` ", database_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\` ", database_name_escaped)
	}
	sql_command += "WHERE name='" + table_name_escaped + "';"

	return &sql_command, options, nil
}
