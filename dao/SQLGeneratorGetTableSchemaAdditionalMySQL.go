package dao

import (
	"fmt"
	"strings"
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
		options.SetBoolValue("json_output", false)
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


func mapAdditionalSchemaFromDBToMap(json_array *json.Array) (*json.Map, []error) {
	var errors []error
	
	if json_array == nil {
		errors = append(errors, fmt.Errorf("error: show table status returned nil records"))
		return nil, errors
	}

	if len(*(json_array.GetValues())) == 0 {
		errors = append(errors, fmt.Errorf("error:  show table status did not return any records"))
		return nil, errors
	}

	additional_schema := json.NewMap()
	for _, column_details := range *(json_array.GetValues()) {
		column_map, column_map_errors := column_details.GetMap()
		if column_map_errors != nil {
			return nil, column_map_errors
		} else if common.IsNil(column_map) {
			errors = append(errors, fmt.Errorf("column_map is nil"))
			return nil, errors
		}
		column_attributes := column_map.GetKeys()

		for _, column_attribute := range column_attributes {
			switch column_attribute {
			case "Comment":
				comment_value, comment_errors := column_map.GetString("Comment")
				if comment_errors != nil {
					errors = append(errors, comment_errors...)
				} else if common.IsNil(comment_value) {
					errors = append(errors, fmt.Errorf("comment is nil"))
				} else {
					if strings.TrimSpace(*comment_value) != "" {
						comment_as_map, comment_as_map_value_errors := json.Parse(strings.TrimSpace(*comment_value))
						if comment_as_map_value_errors != nil {
							errors = append(errors, comment_as_map_value_errors...)
						} else if common.IsNil(comment_as_map) {
							errors = append(errors, fmt.Errorf("comment is nil"))
						} else {
							additional_schema.SetMap("Comment", comment_as_map)
						}
					}
				}
			default:
				column_attribute_value, column_attribute_value_errors := column_map.GetString(column_attribute)
				if column_attribute_value_errors != nil {
					errors = append(errors, column_attribute_value_errors...)
				} else if common.IsNil(column_attribute_value) {
					errors = append(errors, fmt.Errorf("%s is nil", column_attribute))
				} else {
					additional_schema.SetStringValue(column_attribute, *column_attribute_value)
				}
			}
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return additional_schema, nil
}
