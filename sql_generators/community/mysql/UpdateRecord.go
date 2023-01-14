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

type UpdateRecordSQL struct {
	GetUpdateRecordSQL func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error)
}

func newUpdateRecordSQL() (*UpdateRecordSQL) {
	get_update_record_sql := func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error
		
		table_name_validation_errors := verify.ValidateTableName(table_name)
		if table_name_validation_errors != nil {
			errors = append(errors, table_name_validation_errors...)
		}

		if len(errors) > 0 {
			return nil, options, errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return nil, options, errors
		}

		_, record_columns_errors := helper.GetRecordColumns(record_data)
		if record_columns_errors != nil {
			return nil, options, record_columns_errors
		}

		record_fields, record_fields_error := helper.GetFields(record_data, "[fields]")
		if record_fields_error != nil {
			return nil, options, record_fields_error
		} else if common.IsNil(record_fields) {
			errors = append(errors, fmt.Errorf("record fields is nil"))
			return nil, options, errors
		}

		primary_key_table_columns, primary_key_table_columns_errors := helper.GetTablePrimaryKeyColumnsForSchema(table_schema)
		if primary_key_table_columns_errors != nil {
			return nil, options, primary_key_table_columns_errors
		}

		foreign_key_table_columns, foreign_key_table_columns_errors := helper.GetTableForeignKeyColumnsForSchema(table_schema)
		if foreign_key_table_columns_errors != nil {
			return nil, options, foreign_key_table_columns_errors
		}

		table_non_primary_key_columns, table_non_primary_key_columns_errors := helper.GetTableNonPrimaryKeyColumnsForSchema(table_schema)
		if table_non_primary_key_columns_errors != nil {
			return nil, options, table_non_primary_key_columns_errors
		}

		record_primary_key_columns, record_primary_key_columns_errors := helper.GetRecordPrimaryKeyColumns(record_data, primary_key_table_columns)
		if record_primary_key_columns_errors != nil {
			return nil, options, record_primary_key_columns_errors
		}

		record_foreign_key_columns, record_foreign_key_columns_errors := helper.GetRecordForeignKeyColumns(record_data, foreign_key_table_columns)
		if record_foreign_key_columns_errors != nil {
			return nil, options, record_foreign_key_columns_errors
		}

		for primary_key_table_column, _ := range *primary_key_table_columns {
			if _, found := (*record_primary_key_columns)[primary_key_table_column]; !found {
				errors = append(errors, fmt.Errorf("error: record did not contain primary key column: %s", primary_key_table_column))
			} 
		}

		for foreign_key_table_column, _ := range *foreign_key_table_columns {
			if _, found := (*record_foreign_key_columns)[foreign_key_table_column]; found {
				if record_fields.IsNull(foreign_key_table_column) {
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
			return nil, options, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString("UPDATE ")

		Box(&sql_command,table_name_escaped,"`","`")

		sql_command.WriteString(" \n")

		sql_command.WriteString("SET ")

		length_record_non_primary_key_columns := len(*record_non_primary_key_columns)
		index := 0
		for record_non_primary_key_column, _  := range *record_non_primary_key_columns {
			record_non_identity_column_escaped, record_non_identity_column_escaped_errors := common.EscapeString(record_non_primary_key_column, "'")
			
			if record_non_identity_column_escaped_errors != nil {
				errors = append(errors, record_non_identity_column_escaped_errors)
				continue
			}

			column_data, paramter_errors := helper.GetField(record_data, "[schema]", "[fields]", record_non_primary_key_column, "self")
			if paramter_errors != nil {
				errors = append(errors, paramter_errors...)
				continue
			}

			column_definition, column_definition_errors := table_schema.GetMap(record_non_primary_key_column)
			if column_definition_errors != nil {
				errors = append(errors, column_definition_errors...) 
				continue
			} else if common.IsNil(column_definition) {
				errors = append(errors, fmt.Errorf("column_definition not found"))
				continue
			}
			
			Box(&sql_command,record_non_identity_column_escaped,"`","`")

			sql_command.WriteString("=")

			if common.IsNil(column_data) {
				sql_command.WriteString("NULL")
			} else {
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
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int64":
					value := column_data.(int64)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
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

							if value_escaped == "0001-01-01 00:00:00.000000" {
								value_escaped = "0000-00-00 00:00:00.000000"
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

							if value_escaped == "0001-01-01 00:00:00.000000" {
								value_escaped = "0000-00-00 00:00:00.000000"
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
					errors = append(errors, fmt.Errorf("error: Record.getUpdateSQL type: %s not supported for table please implement", rep))
				}
			}

			if index < length_record_non_primary_key_columns-1 {
				sql_command.WriteString(", \n")
			}
			index++
		}

		sql_command.WriteString(" WHERE ")
		index = 0
		number_of_primary_keys := len(*primary_key_table_columns)
		for primary_key_table_column, _ := range *primary_key_table_columns {
			primary_key_table_column_ecaped, primary_key_table_column_ecaped_errors := common.EscapeString(primary_key_table_column, "'")
			
			if primary_key_table_column_ecaped_errors != nil {
				errors = append(errors, primary_key_table_column_ecaped_errors)
				continue
			}

			column_data, paramter_errors := helper.GetField(record_data, "[schema]", "[fields]", primary_key_table_column, "self")
			if paramter_errors != nil {
				errors = append(errors, paramter_errors...)
				continue
			}

			column_definition, column_definition_errors := table_schema.GetMap(primary_key_table_column)
			if column_definition_errors != nil {
				errors = append(errors, column_definition_errors...) 
				continue
			} else if common.IsNil(column_definition) {
				errors = append(errors, fmt.Errorf("column_definition not found"))
				continue
			}

			Box(&sql_command, primary_key_table_column_ecaped,"`","`")

			sql_command.WriteString(" = ")

			if common.IsNil(column_data) {
				sql_command.WriteString("NULL")
			} else {
				record_non_identity_column_type := common.GetType(column_data)
				switch record_non_identity_column_type {
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
					errors = append(errors, fmt.Errorf("error: update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
				}
			}

			if index < (number_of_primary_keys - 1) {
				sql_command.WriteString(" AND ")
			}
			index++
		}
		sql_command.WriteString(" ;")

		if len(errors) > 0 {
			return nil, options, errors
		}

		return &sql_command, options, nil
	}

	return &UpdateRecordSQL{
		GetUpdateRecordSQL: func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_update_record_sql(verify, table_name, table_schema, valid_columns, record_data, options)
		},
	}
}
