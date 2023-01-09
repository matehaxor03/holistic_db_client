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

func GetCreateRecordSQLMySQL(verify *validate.Validator, struct_type string, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	table_validation_errors := verify.ValidateTableName(table_name)
	if table_validation_errors != nil {
		return nil, nil, table_validation_errors
	}

	if common.IsNil(record_data) {
		errors = append(errors, fmt.Errorf("record_data is nil"))
		return nil, nil, errors
	}

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("no_column_headers", false)
		options.SetBoolValue("get_last_insert_id", true)
		options.SetBoolValue("transactional", false)
	}

	record_columns, record_columns_errors := helper.GetRecordColumns(struct_type, &record_data)
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

	
	sql_command := "INSERT INTO "
	
	if options.IsBoolTrue("use_file") {
		sql_command += "`"
	} else {
		sql_command += "\\`"
	}

	sql_command += table_name_escaped
	
	if options.IsBoolTrue("use_file") {
		sql_command += "`"
	} else {
		sql_command += "\\`"
	}

	sql_command += " ("
	for index, record_column := range *record_columns {
		if _, found := (valid_columns)[record_column]; !found {
			errors = append(errors, fmt.Errorf("column does not exist"))
			continue
		}
		
		record_column_escaped,record_column_escaped_errors := common.EscapeString(record_column, "'")
		if record_column_escaped_errors != nil {
			errors = append(errors, record_column_escaped_errors)
			continue
		}
		
		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}

		sql_command += record_column_escaped
		
		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}

		if index < (len(*record_columns) - 1) {
			sql_command += ", "
		}
	}

	sql_command += ") VALUES ("
	for index, record_column := range *record_columns {
		column_data, paramter_errors := helper.GetField(struct_type, &record_data, "[schema]", "[fields]", record_column, "self")
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
			sql_command += strconv.FormatUint(*value, 10)
		case "uint64":
			value := column_data.(uint64)
			sql_command += strconv.FormatUint(value, 10)
		case "*int64":
			value := column_data.(*int64)
			sql_command += strconv.FormatInt(*value, 10)
		case "int64":
			value := column_data.(int64)
			sql_command += strconv.FormatInt(value, 10)
		case "*uint32":
			value := column_data.(*uint32)
			sql_command += strconv.FormatUint(uint64(*value), 10)
		case "uint32":
			value := column_data.(uint32)
			sql_command += strconv.FormatUint(uint64(value), 10)
		case "*int32":
			value := column_data.(*int32)
			sql_command += strconv.FormatInt(int64(*value), 10)
		case "int32":
			value := column_data.(int32)
			sql_command += strconv.FormatInt(int64(value), 10)
		case "*uint16":
			value := column_data.(*uint16)
			sql_command += strconv.FormatUint(uint64(*value), 10)
		case "uint16":
			value := column_data.(uint16)
			sql_command += strconv.FormatUint(uint64(value), 10)
		case "*int16":
			value := column_data.(*int16)
			sql_command += strconv.FormatInt(int64(*value), 10)
		case "int16":
			value := column_data.(int16)
			sql_command += strconv.FormatInt(int64(value), 10)
		case "*uint8":
			value := column_data.(*uint8)
			sql_command += strconv.FormatUint(uint64(*value), 10)
		case "uint8":
			value := column_data.(uint8)
			sql_command += strconv.FormatUint(uint64(value), 10)
		case "*int8":
			value := column_data.(*int8)
			sql_command += strconv.FormatInt(int64(*value), 10)
		case "int8":
			value := column_data.(int8)
			sql_command += strconv.FormatInt(int64(value), 10)
		case "*int":
			value := column_data.(*int)
			sql_command += strconv.FormatInt(int64(*value), 10)
		case "int":
			value := column_data.(int)
			sql_command += strconv.FormatInt(int64(value), 10)
		case "float32":
			float_32_string_value := fmt.Sprintf("%f", column_data.(float32))
			if !strings.Contains(float_32_string_value, ".") {
				float_32_string_value += ".0"
			}
			sql_command += float_32_string_value
		case "*float32":
			float_32_string_value := fmt.Sprintf("%f", *(column_data.(*float32)))
			if !strings.Contains(float_32_string_value, ".") {
				float_32_string_value += ".0"
			}
			sql_command += float_32_string_value
		case "float64":
			float_64_string_value := fmt.Sprintf("%f", column_data.(float64))
			if !strings.Contains(float_64_string_value, ".") {
				float_64_string_value += ".0"
			}
			sql_command += float_64_string_value
		case "*float64":
			float_64_string_value := fmt.Sprintf("%f", *(column_data.(*float64)))
			if !strings.Contains(float_64_string_value, ".") {
				float_64_string_value += ".0"
			}
			sql_command += float_64_string_value
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

					if options.IsBoolTrue("use_file") {
						sql_command += "'" + value_escaped + "'"
					} else {
						sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
					}
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

					if options.IsBoolTrue("use_file") {
						sql_command += "'" + value_escaped + "'"
					} else {
						sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
					}
				}
			}
		case "string":
			value_escaped, value_escaped_errors := common.EscapeString(column_data.(string), "'")
			if value_escaped_errors != nil {
				errors = append(errors, value_escaped_errors)
			}

			if options.IsBoolTrue("use_file") {
				sql_command += "'" + value_escaped + "'"
			} else {
				sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
			}

		case "*string":
			value_escaped, value_escaped_errors := common.EscapeString(*(column_data.(*string)), "'")
			if value_escaped_errors != nil {
				errors = append(errors, value_escaped_errors)
			}

			if options.IsBoolTrue("use_file") {
				sql_command += "'" + value_escaped + "'"
			} else {
				sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
			}

		case "bool":
			if column_data.(bool) {
				sql_command += "1"
			} else {
				sql_command += "0"
			}
		case "*bool":
			if *(column_data.(*bool)) {
				sql_command += "1"
			} else {
				sql_command += "0"
			}
		default:
			errors = append(errors, fmt.Errorf("error: %s Record.getCreateSQL type: %s not supported for table please implement", struct_type, rep))
		}

		if index < (len(*record_columns) - 1) {
			sql_command += ", "
		}
	}
	sql_command += ");"

	if len(errors) > 0 {
		return nil, nil, errors
	}

	return &sql_command, options, nil
}
