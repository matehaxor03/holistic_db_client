package mysql

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetCreateRecordSQLMySQL(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	table_validation_errors := verify.ValidateTableName(table_name)
	if table_validation_errors != nil {
		return nil, nil, table_validation_errors
	}

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("no_column_headers", false)
		options.SetBoolValue("get_last_insert_id", true)
		options.SetBoolValue("transactional", false)
	}

	record_columns, record_columns_errors := helper.GetRecordColumns(record_data)
	if record_columns_errors != nil {
		return nil, nil, record_columns_errors
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
		return nil, nil, errors
	}

	auto_increment_columns := 0
	for valid_column, _ := range valid_columns {
		column_definition, column_definition_errors := table_schema.GetMap(valid_column)
		if column_definition_errors != nil {
			errors = append(errors, column_definition_errors...) 
			continue
		} else if common.IsNil(column_definition) {
			errors = append(errors, fmt.Errorf("schema column definition is nil %s schema has keys %s and schema: %s", valid_column, table_schema.GetKeys())) 
			continue
		}

		if !column_definition.IsBool("primary_key") || !column_definition.IsBool("auto_increment") {
			continue
		}

		primary_key_value, primary_key_value_errors := column_definition.GetBool("primary_key")
		if primary_key_value_errors != nil {
			errors = append(errors, primary_key_value_errors...)
			continue
		}

		if *primary_key_value == false {
			continue
		}

		auto_increment_value, auto_increment_value_errors := column_definition.GetBool("auto_increment")
		if auto_increment_value_errors != nil {
			errors = append(errors, auto_increment_value_errors...)
			continue
		}

		if *auto_increment_value == false {
			continue
		}

		options.SetBoolValue("get_last_insert_id", true)
		options.SetStringValue("auto_increment_column_name", valid_column)
		auto_increment_columns += 1
	}

	if auto_increment_columns > 1 {
		errors = append(errors, fmt.Errorf("error: table: %s can only have 1 auto_increment primary_key column, found: %s", table_name, auto_increment_columns))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	var sql_command strings.Builder
	sql_command.WriteString("INSERT INTO ")
	
	
	sql_command.WriteString(table_name_escaped)
	

	sql_command.WriteString(" (")
	index := 0
	for record_column, _ := range *record_columns {
		if _, found := (valid_columns)[record_column]; !found {
			errors = append(errors, fmt.Errorf("column does not exist"))
			continue
		}
		
		record_column_escaped,record_column_escaped_errors := common.EscapeString(record_column, "'")
		if record_column_escaped_errors != nil {
			errors = append(errors, record_column_escaped_errors)
			continue
		}
		
		sql_command.WriteString(record_column_escaped)

		if index < (len(*record_columns) - 1) {
			sql_command.WriteString(", ")
		}
		index++
	}

	sql_command.WriteString(") VALUES (")
	index = 0
	for record_column, _  := range *record_columns {
		column_data, paramter_errors := helper.GetField(record_data, "[schema]", "[fields]", record_column, "self")
		if paramter_errors != nil {
			errors = append(errors, paramter_errors...)
			continue
		}

		column_definition, column_definition_errors := table_schema.GetMap(record_column)
		if column_definition_errors != nil {
			errors = append(errors, column_definition_errors...) 
			continue
		} else if common.IsNil(column_definition) {
			errors = append(errors, fmt.Errorf("column_definition not found"))
			continue
		}

		rep := common.GetType(column_data)
		switch rep {
		case "*uint64":
			value := column_data.(*uint64)
			sql_command.WriteString(strconv.FormatUint(*value, 10))
		case "uint64":
			value := column_data.(uint64)
			sql_command.WriteString(strconv.FormatUint(value, 10))
		case "*int64":
			value := column_data.(*int64)
			sql_command.WriteString(strconv.FormatInt(*value, 10))
		case "int64":
			value := column_data.(int64)
			sql_command.WriteString(strconv.FormatInt(value, 10))
		case "*uint32":
			value := column_data.(*uint32)
			sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
		case "uint32":
			value := column_data.(uint32)
			sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
		case "*int32":
			value := column_data.(*int32)
			sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
		case "int32":
			value := column_data.(int32)
			sql_command.WriteString(strconv.FormatInt(int64(value), 10))
		case "*uint16":
			value := column_data.(*uint16)
			sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
		case "uint16":
			value := column_data.(uint16)
			sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
		case "*int16":
			value := column_data.(*int16)
			sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
		case "int16":
			value := column_data.(int16)
			sql_command.WriteString(strconv.FormatInt(int64(value), 10))
		case "*uint8":
			value := column_data.(*uint8)
			sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
		case "uint8":
			value := column_data.(uint8)
			sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
		case "*int8":
			value := column_data.(*int8)
			sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
		case "int8":
			value := column_data.(int8)
			sql_command.WriteString(strconv.FormatInt(int64(value), 10))
		case "*int":
			value := column_data.(*int)
			sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
		case "int":
			value := column_data.(int)
			sql_command.WriteString(strconv.FormatInt(int64(value), 10))
		case "float32":
			float_32_string_value := fmt.Sprintf("%f", column_data.(float32))
			sql_command.WriteString(float_32_string_value)
			if !strings.Contains(float_32_string_value, ".") {
				sql_command.WriteString(".0")
			}
		case "*float32":
			float_32_string_value := fmt.Sprintf("%f", *(column_data.(*float32)))
			sql_command.WriteString(float_32_string_value)
			if !strings.Contains(float_32_string_value, ".") {
				sql_command.WriteString(".0")
			}
		case "float64":
			float_64_string_value := fmt.Sprintf("%f", column_data.(float64))
			sql_command.WriteString(float_64_string_value)
			if !strings.Contains(float_64_string_value, ".") {
				sql_command.WriteString(".0")
			}
		case "*float64":
			float_64_string_value := fmt.Sprintf("%f", *(column_data.(*float64)))
			sql_command.WriteString(float_64_string_value)
			if !strings.Contains(float_64_string_value, ".") {
				sql_command.WriteString(".0")
			}
		case "*time.Time":
			value := column_data.(*time.Time)
			decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
			if decimal_places_error != nil {
				errors = append(errors, decimal_places_error...)
			} else if decimal_places == nil {
				errors = append(errors, fmt.Errorf("decimal_places is nil"))
			} else {
				format_time, format_time_errors := common.FormatTime(*value, *decimal_places)
				if format_time_errors != nil {
					errors = append(errors, format_time_errors...)
				} else if format_time == nil { 
					errors = append(errors, fmt.Errorf("format time is nil"))
				} else {
					value_escaped, value_escaped_errors := common.EscapeString(*format_time, "'")
					if value_escaped_errors != nil {
						errors = append(errors, value_escaped_errors)
					}

					sql_command.WriteString("'")
					sql_command.WriteString(value_escaped)
					sql_command.WriteString("'")
				}
			}
		case "time.Time":
			value := column_data.(time.Time)
			decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
			if decimal_places_error != nil {
				errors = append(errors, decimal_places_error...)
			} else if decimal_places == nil {
				errors = append(errors, fmt.Errorf("decimal_places is nil"))
			} else {
				format_time, format_time_errors := common.FormatTime(value, *decimal_places)
				if format_time_errors != nil {
					errors = append(errors, format_time_errors...)
				} else if format_time == nil { 
					errors = append(errors, fmt.Errorf("format time is nil"))
				} else {
					value_escaped, value_escaped_errors := common.EscapeString(*format_time, "'")
					if value_escaped_errors != nil {
						errors = append(errors, value_escaped_errors)
					}

					sql_command.WriteString("'")
					sql_command.WriteString(value_escaped)
					sql_command.WriteString("'")
				}
			}
		case "string":
			value_escaped, value_escaped_errors := common.EscapeString(column_data.(string), "'")
			if value_escaped_errors != nil {
				errors = append(errors, value_escaped_errors)
			}

			sql_command.WriteString("'")
			sql_command.WriteString(value_escaped)
			sql_command.WriteString("'")

		case "*string":
			value_escaped, value_escaped_errors := common.EscapeString(*(column_data.(*string)), "'")
			if value_escaped_errors != nil {
				errors = append(errors, value_escaped_errors)
			}

			sql_command.WriteString("'")
			sql_command.WriteString(value_escaped)
			sql_command.WriteString("'")

		case "bool":
			if column_data.(bool) {
				sql_command.WriteString("1")
			} else {
				sql_command.WriteString("0")
			}
		case "*bool":
			if *(column_data.(*bool)) {
				sql_command.WriteString("1")
			} else {
				sql_command.WriteString("0")
			}
		default:
			errors = append(errors, fmt.Errorf("error: Record.getCreateSQL type: %s not supported for table please implement", rep))
		}

		if index < (len(*record_columns) - 1) {
			sql_command.WriteString(", ")
		}
		index++
	}
	sql_command.WriteString(");")

	if len(errors) > 0 {
		return nil, nil, errors
	}

	sql_command_result := sql_command.String()
	return &sql_command_result, options, nil
}
