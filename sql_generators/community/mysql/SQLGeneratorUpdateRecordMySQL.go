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

func GetUpdateRecordSQLMySQL(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	table_name_validation_errors := verify.ValidateTableName(table_name)
	if table_name_validation_errors != nil {
		errors = append(errors, table_name_validation_errors...)
	}
	
	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
	if table_name_escaped_errors != nil {
		errors = append(errors, table_name_escaped_errors)
		return nil, nil, errors
	}

	record_columns, record_columns_errors := helper.GetRecordColumns(record_data)
	if record_columns_errors != nil {
		return nil, nil, record_columns_errors
	}

	for record_column, _   := range *record_columns {
		if strings.HasPrefix(record_column, "credential") {
			options.SetBoolValue("use_file", true)
		}
	}

	primary_key_table_columns, primary_key_table_columns_errors := helper.GetTablePrimaryKeyColumnsForSchema(table_schema)
	if primary_key_table_columns_errors != nil {
		return nil, nil, primary_key_table_columns_errors
	}

	foreign_key_table_columns, foreign_key_table_columns_errors := helper.GetTableForeignKeyColumnsForSchema(table_schema)
	if foreign_key_table_columns_errors != nil {
		return nil, nil, foreign_key_table_columns_errors
	}

	table_non_primary_key_columns, table_non_primary_key_columns_errors := helper.GetTableNonPrimaryKeyColumnsForSchema(table_schema)
	if table_non_primary_key_columns_errors != nil {
		return nil, nil, table_non_primary_key_columns_errors
	}

	record_primary_key_columns, record_primary_key_columns_errors := helper.GetRecordPrimaryKeyColumns(record_data, primary_key_table_columns)
	if record_primary_key_columns_errors != nil {
		return nil, nil, record_primary_key_columns_errors
	}

	record_foreign_key_columns, record_foreign_key_columns_errors := helper.GetRecordForeignKeyColumns(record_data, foreign_key_table_columns)
	if record_foreign_key_columns_errors != nil {
		return nil, nil, record_foreign_key_columns_errors
	}

	for _, record_primary_key_column := range *record_primary_key_columns {
		if _, found := (*primary_key_table_columns)[record_primary_key_column]; !found {
			errors = append(errors, fmt.Errorf("error: record did not contain primary key column: %s", record_primary_key_column))
		}
	}

	for foreign_key_table_column, _ := range *foreign_key_table_columns {
		if _, found := (*record_foreign_key_columns)[foreign_key_table_column]; found {
			record_forign_key_column_data := record_data.GetObjectForMap(foreign_key_table_column)
			if common.IsNil(record_forign_key_column_data) {
				errors = append(errors, fmt.Errorf("error: record had foreign key set however was null: %s", foreign_key_table_column))
			}
		}
	}

	record_data.SetTime("last_modified_date", common.GetTimeNow())

	archieved, archieved_errors := record_data.GetBool("archieved")
	if archieved_errors != nil {
		errors = append(errors, archieved_errors...)
	} else if !common.IsNil(archieved) {
		if *archieved {
			record_data.SetTime("archieved_date", common.GetTimeNow())
		} else {
			record_data.SetStringValue("archieved_date", "0000-00-00 00:00:00.000000")
		}
	}

	record_non_primary_key_columns, record_non_primary_key_columns_errors := helper.GetRecordNonPrimaryKeyColumnsUpdate(record_data, table_non_primary_key_columns)
	if record_non_primary_key_columns_errors != nil {
		return nil, options, record_non_primary_key_columns_errors
	}

	if len(*record_non_primary_key_columns) == 0 {
		errors = append(errors, fmt.Errorf("error: no non-primary key columns detected in record to update"))
	}

	if len(*primary_key_table_columns) == 0 {
		errors = append(errors, fmt.Errorf("error: table schema has no identity columns"))
	}

	if _, found := (*table_non_primary_key_columns)["last_modified_date"]; !found {
		errors = append(errors, fmt.Errorf("error: table schema does not have last_modified_date"))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	sql_command := "UPDATE "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s` \n", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\` \n", table_name_escaped)
	}

	sql_command += "SET "

	length_record_non_primary_key_columns := len(*record_non_primary_key_columns)
	index := 0
	for record_non_primary_key_column, _  := range *record_non_primary_key_columns {
		record_non_identity_column_escaped, record_non_identity_column_escaped_errors := common.EscapeString(record_non_primary_key_column, "'")
		
		if record_non_identity_column_escaped_errors != nil {
			errors = append(errors, record_non_identity_column_escaped_errors)
			continue
		}

		column_definition, column_definition_errors := table_schema.GetMap(record_non_primary_key_column)
		if column_definition_errors != nil {
			errors = append(errors, column_definition_errors...) 
			continue
		} else if common.IsNil(column_definition) {
			errors = append(errors, fmt.Errorf("column definition not found")) 
			continue
		}
		
		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}
		
		sql_command += record_non_identity_column_escaped
		
		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}

		sql_command += "="
		column_data := record_data.GetObjectForMap(record_non_primary_key_column)

		if common.IsNil(column_data) {
			sql_command += "NULL"
		} else {
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
				sql_command += strconv.FormatInt(int64(*value), 10)
			case "int64":
				value := column_data.(int64)
				sql_command += strconv.FormatInt(int64(value), 10)
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
				sql_command +=  strconv.FormatUint(uint64(value), 10)
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
				errors = append(errors, fmt.Errorf("error: Record.getUpdateSQL type: %s not supported for table please implement", rep))
			}
		}

		if index < length_record_non_primary_key_columns-1 {
			sql_command += ", \n"
		}
		index++
	}

	sql_command += " WHERE "
	index = 0
	number_of_primary_keys := len(*primary_key_table_columns)
	for primary_key_table_column, _ := range *primary_key_table_columns {
		primary_key_table_column_ecaped, primary_key_table_column_ecaped_errors := common.EscapeString(primary_key_table_column, "'")
		
		if primary_key_table_column_ecaped_errors != nil {
			errors = append(errors, primary_key_table_column_ecaped_errors)
			continue
		}

		column_definition, column_definition_errors := table_schema.GetMap(primary_key_table_column_ecaped)
		if column_definition_errors != nil {
			errors = append(errors, column_definition_errors...) 
			continue
		}

		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}
		
		sql_command += primary_key_table_column_ecaped
		
		if options.IsBoolTrue("use_file") {
			sql_command += "`"
		} else {
			sql_command += "\\`"
		}

		sql_command += " = "
		column_data := record_data.GetObjectForMap(primary_key_table_column)

		if common.IsNil(column_data) {
			sql_command += "NULL"
		} else {
			record_non_identity_column_type := common.GetType(column_data)
			switch record_non_identity_column_type {
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
				sql_command += fmt.Sprintf("%f", column_data.(float32))
			case "*float32":
				sql_command += fmt.Sprintf("%f", *(column_data.(*float32)))
			case "float64":
				sql_command += fmt.Sprintf("%f", column_data.(float64))
			case "*float64":
				sql_command += fmt.Sprintf("%f", *(column_data.(*float64)))
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
				errors = append(errors, fmt.Errorf("error: update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
			}
		}

		if index < (number_of_primary_keys - 1) {
			sql_command += " AND "
		}
		index++
	}
	sql_command += " ;"

	if len(errors) > 0 {
		return nil, nil, errors
	}

	return &sql_command, options, nil
}

