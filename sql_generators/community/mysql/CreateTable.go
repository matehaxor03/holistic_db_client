package mysql

import (
	"fmt"
	"strconv"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

func GetCreateTableSQL(verify *validate.Validator, table_data json.Map, options *json.Map) (*string, *json.Map, []error) {
	var errors []error

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
	}

	temp_table_name, temp_table_name_errors := helper.GetTableName(table_data)
	if temp_table_name_errors != nil {
		errors = append(errors, temp_table_name_errors...)
	} else if common.IsNil(temp_table_name) {
		errors = append(errors, fmt.Errorf("temp_table_name is nil"))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	validate_table_name_errors := verify.ValidateTableName(temp_table_name)
	if validate_table_name_errors != nil  {
		errors = append(errors, validate_table_name_errors...)
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
	} else if common.IsNil(temp_table_name) {
		errors = append(errors, fmt.Errorf("table_name_escaped is nil"))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	sql_command := ""
	sql_command += "CREATE TABLE "
	sql_command += fmt.Sprintf("%s ", table_name_escaped)

	valid_columns, valid_columns_errors := helper.GetTableColumns(table_data)
	if valid_columns_errors != nil {
		errors = append(errors, valid_columns_errors...)
	} else if common.IsNil(valid_columns) {
		errors = append(errors, fmt.Errorf("table_columns is nil"))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	for valid_column, _ := range *valid_columns {
		valid_columns_errors := verify.ValidateColumnName(valid_column)
		if valid_columns_errors != nil {
			errors = append(errors, valid_columns_errors...)
		}
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	schemas_map, schemas_map_errors := helper.GetSchemas(table_data, "[schema]")
	if schemas_map_errors != nil {
		return nil, nil, schemas_map_errors
	}

	primary_key_count := 0

	sql_command += "("
	number_of_valid_columns := len(*valid_columns)
	index := 0
	for column, _ := range *valid_columns {
		columnSchema, columnSchema_errors := schemas_map.GetMap(column)
		if columnSchema_errors != nil {
			errors = append(errors, columnSchema_errors...)
			continue
		} else if common.IsNil(columnSchema) {
			errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column schema for column: %s is nil", column))
			continue
		}

		column_escaped, column_escaped_errors := common.EscapeString(column, "'")
		if column_escaped_errors != nil {
			errors = append(errors, column_escaped_errors)
		} else if common.IsNil(column_escaped) {
			errors = append(errors, fmt.Errorf("column_escaped is nil"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command += column_escaped

		typeOf, type_of_errors := columnSchema.GetString("type")
		if type_of_errors != nil {
			errors = append(errors, type_of_errors...)
			continue
		} else if common.IsNil(typeOf) {
			errors = append(errors, fmt.Errorf("type is nil"))
		}

		switch *typeOf {
		case "*uint64", "uint64","*int64", "int64", "*uint32", "uint32", "*int32","int32", "*uint16", "uint16", "*int16","int16",  "*uint8", "uint8", "*int8","int8":
			switch *typeOf {
			case "*uint64", "*int64", "uint64", "int64":
				sql_command += " BIGINT"
			case "*uint32", "*int32", "uint32", "int32":
				sql_command += " INT"
			case "*uint16", "*int16", "uint16", "int16":
				sql_command += " SMALLINT"
			case "*uint8", "*int8", "uint8", "int8":
				sql_command += " TINYINT"
			default:
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL number type not mapped: %s", *typeOf))
			}

			unsigned_number := false
			switch *typeOf {
			case "*uint64", "uint64":
				unsigned_number = true
			case "*uint32", "uint32":
				unsigned_number = true
			case "*uint16", "uint16":
				unsigned_number = true
			case "*uint8","uint8":
				unsigned_number = true
			default:
			}

			if unsigned_number {
				sql_command += " UNSIGNED"
			}

			if !strings.HasPrefix(*typeOf, "*") {
				sql_command += " NOT NULL"
			}

			if columnSchema.HasKey("auto_increment") {
				if columnSchema.IsBool("auto_increment") && !columnSchema.IsNull("auto_increment") {
					if columnSchema.IsBoolTrue("auto_increment") {
						sql_command += " AUTO_INCREMENT"
					}
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.GetType("auto_increment")))
				}
			}

			if columnSchema.HasKey("primary_key") {
				if columnSchema.IsBool("primary_key") && !columnSchema.IsNull("primary_key") {
					if columnSchema.IsBoolTrue("primary_key") {
						sql_command += " PRIMARY KEY"
						primary_key_count += 1
					}
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
				}
			} 

			if columnSchema.HasKey("default") {
				if columnSchema.IsBoolTrue("primary_key") || columnSchema.IsBoolTrue("foreign_key") {
					sql_command += " "
				} else if columnSchema.IsInteger("default") {
					default_value, default_value_errors := columnSchema.GetInt64("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else {
						sql_command += " DEFAULT " + strconv.FormatInt(*default_value, 10)
					}
				} else if columnSchema.IsString("default") {
					default_value, default_value_errors := columnSchema.GetString("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else if common.IsNil(default_value) {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: default contained string value but was nil type: %s", column, columnSchema.GetType("default")))
					} else if *default_value == "nil" && strings.HasPrefix(*typeOf, "*") {
						sql_command += " DEFAULT 0 " 
					} else if columnSchema.IsNull("default") {
						sql_command += " DEFAULT 0 " 
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: default contained string value which is not supported: %s", column, columnSchema.GetType("default")))
					}
				} else if columnSchema.IsNull("default") && strings.HasPrefix(*typeOf, "*") {
					sql_command += " DEFAULT 0 " 
				} else if columnSchema.IsNull("default") {
					sql_command += " DEFAULT 0 "
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: default contained a value which is not supported: %s", column, columnSchema.GetType("default")))
				}
			}
		case "*time.Time", "time.Time":
			decimal_places, decimal_places_error := columnSchema.GetInt("decimal_places")
			if decimal_places_error != nil {
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained a value which is not supported %s", column, fmt.Sprintf("%s", decimal_places_error)))
			} else if common.IsNil(decimal_places) {
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained a value which is not supported: nil", column))
			} else if *decimal_places < 0  || *decimal_places > 6 {
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained invalid decimal range outside [0-6]: %d", column, *decimal_places))
			} else {
				if *decimal_places == 0 {
					sql_command += " TIMESTAMP"
				} else {
					sql_command += fmt.Sprintf(" TIMESTAMP(%d)", *decimal_places)
				}

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") {
					default_value, default_value_errors := columnSchema.GetString("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else if default_value == nil {
						sql_command += " DEFAULT NULL"
					} else if *default_value == "now" {
						if *decimal_places == 0 {
							sql_command += " DEFAULT CURRENT_TIMESTAMP"
						} else {
							sql_command += fmt.Sprintf(" DEFAULT CURRENT_TIMESTAMP(%d)", *decimal_places)
						}
					} else if *default_value == "zero" {
						default_time, default_time_errors := columnSchema.GetTimeWithDecimalPlaces("default", *decimal_places)
						if default_time_errors != nil {
							errors = append(errors, default_time_errors...)
						} else if common.IsNil(default_time) {
							errors = append(errors, fmt.Errorf("default zero time was nil"))
						} else {
							time_zero_as_string, time_zero_as_string_errors := common.GetTimeZeroStringSQL(*decimal_places)
							if time_zero_as_string_errors != nil {
								errors = append(errors, time_zero_as_string_errors...)
							} else {
								value_escaped, value_escaped_errors := common.EscapeString(*time_zero_as_string, "'")
								if value_escaped_errors != nil {
									errors = append(errors, value_escaped_errors)
								} else {
									sql_command += " DEFAULT "
									if options.IsBoolTrue("use_file") {
										sql_command += "'" + value_escaped + "'"
									} else {
										sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
									}
								}	
							}
						}
					}  else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had default value it did not understand", column))
					}
				}




			}

			

			
		case "*bool", "bool":
			sql_command += " BOOLEAN"

			if !strings.HasPrefix(*typeOf, "*") {
				sql_command += " NOT NULL"
			}

			if columnSchema.HasKey("default") {
				if columnSchema.IsNull("default") {
					sql_command += " DEFAULT 0"
				} else if !columnSchema.IsBool("default") {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
				} else if columnSchema.IsBoolTrue("default") {
					sql_command += " DEFAULT 1"
				} else if columnSchema.IsBoolFalse("default") {
					sql_command += " DEFAULT 0"
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
				}
			}
		case "*float32", "float32":
			sql_command += " FLOAT"

			if !strings.HasPrefix(*typeOf, "*") {
				sql_command += " NOT NULL"
			}

			if columnSchema.HasKey("default") {
				if columnSchema.IsNull("default") {
					sql_command += " DEFAULT 0"
				} else if !columnSchema.IsFloat("default") {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
				} else if columnSchema.IsFloat("default") {
					default_float_value, default_float_value_errors := columnSchema.GetFloat32("default")
					if default_float_value_errors != nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for float32 default value %s", column, fmt.Sprintf("%s", default_float_value_errors)))
					} else if default_float_value == nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s float32 default value returned nil", column))
					} else {
						sql_command += fmt.Sprintf(" DEFAULT %f", *default_float_value)
					}
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
				}
			}
		case "*float64", "float64":
			sql_command += " DOUBLE"

			if !strings.HasPrefix(*typeOf, "*") {
				sql_command += " NOT NULL"
			}

			if columnSchema.HasKey("default") {
				if columnSchema.IsNull("default") {
					sql_command += " DEFAULT 0"
				} else if !columnSchema.IsFloat("default") {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
				} else if columnSchema.IsFloat("default") {
					default_float_value, default_float_value_errors := columnSchema.GetFloat64("default")
					if default_float_value_errors != nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for float32 default value %s", column, fmt.Sprintf("%s", default_float_value_errors)))
					} else if default_float_value == nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s float32 default value returned nil", column))
					} else {
						sql_command += fmt.Sprintf(" DEFAULT %f", *default_float_value)
					}
				} else {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
				}
			}
		case "*string", "string":
			sql_command += " VARCHAR("
			if !columnSchema.HasKey("max_length") {
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s did not specify max_length attribute", column))
			} else if !columnSchema.IsInteger("max_length") {
				errors = append(errors, fmt.Errorf("error: column: %s specified max_length attribute however it's not an int", column))
			} else {
				max_length, max_length_errors := columnSchema.GetInt("max_length")
				if max_length_errors != nil {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified max_length attribute had errors %s", column, fmt.Sprintf("%s", max_length_errors)))
				} else if *max_length < 0 {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified max_length attribute was < 0 and had value: %d", column, max_length))
				} else {
					// utf-8 should use 4 bytes (maxiumum per character) but in mysql it's 3 bytes but to be consistent going to assume 4 bytes, 
					sql_command += fmt.Sprintf("%d", (4*(*max_length)))
				}
			}
			sql_command += ")"

			if !strings.HasPrefix(*typeOf, "*") {
				sql_command += " NOT NULL"
			}

			if columnSchema.HasKey("default") {
				if columnSchema.IsNull("default") {
					sql_command += " DEFAULT ''"
				} else if !columnSchema.IsString("default") {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-string default value", column))
				} else {
					default_value, default_value_errors := columnSchema.GetString("default")
					if default_value_errors != nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified default attribute had errors %s", column, fmt.Sprintf("%s", default_value_errors)))
					} else {

						default_value_escaped, default_value_escaped_errors := common.EscapeString(*default_value, "'")
						if default_value_escaped_errors != nil {
							errors = append(errors, default_value_escaped_errors)
						}

						sql_command += " DEFAULT "

						if options.IsBoolTrue("use_file") {
							sql_command += "'" + default_value_escaped + "'"
						} else {
							sql_command += strings.ReplaceAll("'" + default_value_escaped + "'", "`", "\\`")
						}

					}
				} 
			}

			
		default:
			errors = append(errors, fmt.Errorf("error: Table.getCreateSQL type: %s is not supported please implement for column %s", *typeOf, column))
		}

		if index < ( number_of_valid_columns - 1) {
			sql_command += ", "
		}
		index++
	}
	sql_command += ");"

	if primary_key_count == 0 {
		errors = append(errors, fmt.Errorf("error: Table.getCreateSQL: %s must have at least 1 primary key", table_name_escaped))
	}

	// todo: check that length of row for all columns does not exceed 65,535 bytes (it's not hard but low priority)

	if len(errors) > 0 {
		return nil, nil, errors
	}

	return &sql_command, options, nil
}

