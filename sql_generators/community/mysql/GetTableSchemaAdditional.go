package mysql

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetTableSchemaAdditionalSQL(verify validate.Validator, struct_type string, database_name string, table_name string, options *json.Map) (*string, *json.Map, []error) {
	var errors []error
	
	database_name_validation_errors := verify.ValidateDatabaseName(database_name)
	if database_name_validation_errors != nil {
		errors = append(errors, database_name_validation_errors...)
	}

	table_name_validation_errors := verify.ValidateTableName(table_name)
	if table_name_validation_errors != nil {
		errors = append(errors, table_name_validation_errors...)
	}

	database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
	if database_name_escaped_errors != nil {
		errors = append(errors, database_name_escaped_errors)
	} else if common.IsNil(database_name_escaped) {
		errors = append(errors, fmt.Errorf("database_name_escaped is nil"))
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
	} else if common.IsNil(table_name_escaped) {
		errors = append(errors, fmt.Errorf("table_name_escaped is nil"))
	}

	if len(errors) > 0 {
		return nil, nil ,errors
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


func MapAdditionalSchemaFromDBToMap(json_array *json.Array) (*json.Map, []error) {
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
					*comment_value = strings.TrimSpace(*comment_value)
					if *comment_value != ""  && strings.HasPrefix(*comment_value, "{") && strings.HasSuffix(*comment_value, "}") {
						comment_as_map, comment_as_map_value_errors := json.Parse(strings.TrimSpace(*comment_value))
						if comment_as_map_value_errors != nil {
							errors = append(errors, comment_as_map_value_errors...)
						} else if common.IsNil(comment_as_map) {
							errors = append(errors, fmt.Errorf("comment is nil"))
						} else {
							additional_schema.SetMap("Comment", comment_as_map)
						}
					} else {
						comment_as_map_raw := json.NewMap()
						comment_as_map_raw.SetStringValue("raw", *comment_value)
						additional_schema.SetMap("Comment", comment_as_map_raw)
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
