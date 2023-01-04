package db_client

import (
	"fmt"
	"strconv"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getSelectRecordsSQLMySQL(table *Table, select_fields *json.Array, filters *json.Map, filters_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64, options *json.Map, column_name_whitelist_characters *json.Map) (*string, *json.Map, []error) {
	var errors []error
	if common.IsNil(table) {
		errors = append(errors, fmt.Errorf("table is nil"))
		return nil, nil, errors
	}

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
	}
	
	validate_errors := table.Validate()
	if errors != nil {
		errors = append(errors, validate_errors...)
		return nil, nil, errors
	}

	table_schema, table_schema_errors := table.GetSchema()
	if table_schema_errors != nil {
		return nil, nil, table_schema_errors
	}

	temp_table_name, temp_table_name_errors := table.GetTableName()
	if temp_table_name_errors != nil {
		return nil, nil, temp_table_name_errors
	}

	table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_errors != nil {
		errors = append(errors, table_name_escaped_errors)
		return nil, nil, errors
	}

	if filters != nil {
		table_columns, table_columns_errors := table.GetTableColumns()
		if table_columns_errors != nil {
			return nil, nil, table_columns_errors
		}

		filter_columns := filters.GetKeys()
		for _, filter_column := range filter_columns {
			if !common.Contains(*table_columns, filter_column) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", filter_column, temp_table_name, *table_columns))
			}
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		for _, filter_column := range filter_columns {
			filter_column_type := filters.GetType(filter_column)

			if !filters.IsNil(filter_column) && !strings.HasPrefix(filter_column_type, "*") {
				filter_column_type = "*" + filter_column_type
			}
				
			if table_schema.IsNil(filter_column) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s does not exist however filter had the value, table has columns: %s", filter_column, temp_table_name, table_schema.GetKeys()))
				continue
			}

			table_schema_column, table_schema_column_errors := table_schema.GetMap(filter_column)
			if table_schema_column_errors != nil {
				errors = append(errors, table_schema_column_errors...)
				continue
			}

			if table_schema_column.IsNil("type") {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s did not have atrribute: type", filter_column, temp_table_name))
				continue
			}


			table_column_type, _ := (*table_schema_column).GetString("type")
			if strings.Replace(*table_column_type, "*", "", -1) != strings.Replace(filter_column_type, "*", "", -1) {
				table_column_type_simple := strings.Replace(*table_column_type, "*", "", -1)
				filter_column_type_simple := strings.Replace(filter_column_type, "*", "", -1)
				if strings.Contains(table_column_type_simple, "int") && strings.Contains(filter_column_type_simple, "int") {

				} else if strings.Contains(table_column_type_simple, "float") && strings.Contains(filter_column_type_simple, "float"){

				} else {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, temp_table_name, *table_column_type))
				}
				

				//todo ignore if filter data_type is nil and table column allows nil
			}
		}
	}

	if select_fields != nil {
		table_columns, table_columns_errors := table.GetTableColumns()
		if table_columns_errors != nil {
			return nil, nil, table_columns_errors
		}

		for _, select_field := range *(select_fields.GetValues()) {
			select_field_value, select_field_value_errors := select_field.GetString()
			if select_field_value_errors != nil {
				return nil, nil, select_field_value_errors
			} else if common.IsNil(select_field_value) {
				errors = append(errors, fmt.Errorf("select_field_value is nil"))
				return nil, nil, errors
			}

			if !common.Contains(*table_columns, *select_field_value) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", select_field, temp_table_name, *table_columns))
			}
		}
	}

	if filters_logic  != nil {
		table_columns, table_columns_errors := table.GetTableColumns()
		if table_columns_errors != nil {
			return nil, nil, table_columns_errors
		}

		temp_filters_fields := filters_logic.GetKeys()
		for _, temp_filters_field := range temp_filters_fields {
			if !common.Contains(*table_columns, temp_filters_field) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", temp_filters_field, temp_table_name, *table_columns))
				continue
			}

			if !filters_logic.IsString(temp_filters_field) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter logic is not a string", temp_filters_field, temp_table_name))
			} else {
				temp_logic_value, temp_logic_value_errors := filters_logic.GetString(temp_filters_field)
				if temp_logic_value_errors != nil {
					errors = append(errors, temp_logic_value_errors...)
				} else if common.IsNil(temp_logic_value) {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value had error getting string value", temp_filters_field, temp_table_name))
				} else {
					if !(*temp_logic_value != "=" ||
						*temp_logic_value != ">" ||
						*temp_logic_value != "<") {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value logic however it's not supported please implement: %s", temp_filters_field, temp_table_name, *temp_logic_value))
					} 
				}
			}
		}
	}

	order_by_clause := ""
	if !common.IsNil(order_by) && len(*(order_by.GetValues())) > 0 {
		table_columns, table_columns_errors := table.GetTableColumns()
		if table_columns_errors != nil {
			return nil, nil, table_columns_errors
		}

		order_by_columns := len(*(order_by.GetValues()))
		for order_by_index, order_by_field := range *(order_by.GetValues()) {

			if common.IsNil(order_by_field) {
				errors = append(errors, fmt.Errorf("order_by_field is nil"))
				continue
			} 

			order_by_map, order_by_map_errors := order_by_field.GetMap()
			if order_by_map_errors != nil {
				errors = append(errors, order_by_map_errors...)
				continue
			} else if common.IsNil(order_by_map) {
				errors = append(errors, fmt.Errorf("order_by_map is nil"))
				continue
			}
			
			order_by_map_column_names := order_by_map.GetKeys()
			if len(order_by_map_column_names) != 1 {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: order by field at index %d was a map however did not have a column name", order_by_index))
				continue
			}

			order_by_column_name := order_by_map_column_names[0]
			
			if !common.Contains(*table_columns, order_by_column_name) {
				errors = append(errors, fmt.Errorf("error: Table.ReadRecords: order by column: %s not found for table: %s available columns are: %s", order_by_column_name, temp_table_name, *table_columns))
				continue
			}

			order_by_string_value, order_by_string_value_errors := order_by_map.GetString(order_by_column_name)
			if order_by_string_value_errors != nil {
				errors = append(errors, order_by_string_value_errors...)
				continue
			} else if common.IsNil(order_by_string_value) {
				errors = append(errors, fmt.Errorf("order by value is nil"))
				continue
			}

			order_by_string_value_validated := ""
			if *order_by_string_value == "ascending" {
				order_by_string_value_validated = "ASC"
			} else if *order_by_string_value == "decending" {
				order_by_string_value_validated = "DESC"
			} else {
				errors = append(errors, fmt.Errorf("order by value is is not valid %s must be ascending or decending", *order_by_string_value))
				continue
			}
			
			if options.IsBoolTrue("use_file") {
				order_by_clause += "`"
			} else {
				order_by_clause += "\\`"
			}
			escaped_order_by_column_name, escaped_order_by_column_name_errors := common.EscapeString(order_by_column_name, "'")
			if escaped_order_by_column_name_errors != nil {
				errors = append(errors, escaped_order_by_column_name_errors)
				continue
			}

			order_by_clause += escaped_order_by_column_name
			if options.IsBoolTrue("use_file") {
				order_by_clause += "`"
			} else {
				order_by_clause += "\\`"
			}

			order_by_clause += " "

			order_by_clause += order_by_string_value_validated

			if order_by_index < (order_by_columns - 1) {
				order_by_clause += ", "
			} else {
				order_by_clause += " "
			}
		}
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	sql_command := "SELECT "
	if select_fields != nil && len(*(select_fields.GetValues())) > 0 {
		select_fields_values_length := len(*(select_fields.GetValues()))
		for i, _ := range *(select_fields.GetValues()) {
			select_fields_value, select_fields_value_errors := select_fields.GetStringValue(i)
			if select_fields_value_errors != nil {
				errors = append(errors, select_fields_value_errors...)
				continue
			}

			escape_string_value, escape_string_value_errors := common.EscapeString(select_fields_value, "'")
			if escape_string_value_errors != nil {
				errors = append(errors, escape_string_value_errors)
			} else {
				sql_command += escape_string_value
				if i < (select_fields_values_length - 1) {
					sql_command += ", "
				} else {
					sql_command += " "
				}
			}
		}
	} else {
		sql_command += "* "
	}

	sql_command += "FROM "
	
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s` ", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\` ", table_name_escaped)
	}

	if filters != nil {
		if len(filters.GetKeys()) > 0 {
			sql_command += "WHERE "
		}

		column_name_params := json.NewMapValue()
		column_name_params.SetObjectForMap("values", column_name_whitelist_characters)
		column_name_params.SetNil("value")
		column_name_params.SetStringValue("label","column_name")
		column_name_params.SetStringValue("data_type", "Table")
		for index, column_filter := range filters.GetKeys() {
			
			column_definition, column_definition_errors := table_schema.GetMap(column_filter)
			if column_definition_errors != nil {
				errors = append(errors, column_definition_errors...) 
				continue
			}
			
			column_name_params.SetString("value", &column_filter)
			column_name_errors := WhitelistCharacters(column_name_params)
			if column_name_errors != nil {
				errors = append(errors, column_name_errors...)
			}	

			column_filter_escaped, column_filter_escaped_errors := common.EscapeString(column_filter, "'")
			if table_name_escaped_errors != nil {
				errors = append(errors, column_filter_escaped_errors)
			}

			if options.IsBoolTrue("use_file") {
				sql_command += "`"
			} else {
				sql_command += "\\`"
			}
			sql_command += column_filter_escaped
			if options.IsBoolTrue("use_file") {
				sql_command += "`"
			} else {
				sql_command += "\\`"
			}

			if common.IsNil(filters_logic) {
				sql_command += " = "
			} else if !filters_logic.HasKey(column_filter) {
				sql_command += " = "
			} else {
				filters_logic_temp, filters_logic_temp_errors := filters_logic.GetString(column_filter)
				if filters_logic_temp_errors != nil {
					errors = append(errors, filters_logic_temp_errors...)
				} else if common.IsNil(filters_logic_temp) {
					errors = append(errors, fmt.Errorf("filters logic is nil"))
				} else {
					sql_command += " "
					sql_command += *filters_logic_temp
					sql_command += " "
				}
			}

			if filters.IsNil(column_filter) {
				sql_command += "NULL "
			} else {
				//todo check data type with schema
				type_of := filters.GetType(column_filter)
				column_data := filters.GetValue(column_filter)
				switch type_of {
				case "*uint64":
					value, value_errors := column_data.GetUInt64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(*value, 10)
					}
				case "uint64":
					value, value_errors := column_data.GetUInt64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(*value, 10)
					}
				case "*int64":
					value, value_errors := column_data.GetInt64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "int64":
					value, value_errors := column_data.GetInt64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "*uint32":
					value, value_errors := column_data.GetUInt32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "uint32":
					value, value_errors := column_data.GetUInt32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "*int32":
					value, value_errors := column_data.GetInt32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=strconv.FormatInt(int64(*value), 10)
					}
				case "int32":
					value, value_errors := column_data.GetInt32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=strconv.FormatInt(int64(*value), 10)
					}
				case "*uint16":
					value, value_errors := column_data.GetUInt16()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "uint16":
					value, value_errors := column_data.GetUInt16()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "*int16":
					value, value_errors := column_data.GetInt16()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "int16":
					value, value_errors := column_data.GetInt16()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "*uint8":
					value, value_errors := column_data.GetUInt8()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "uint8":
					value, value_errors := column_data.GetUInt8()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatUint(uint64(*value), 10)
					}
				case "*int8":
					value, value_errors := column_data.GetInt8()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "int8":
					value, value_errors := column_data.GetInt8()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "*int":
					value, value_errors := column_data.GetInt()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "int":
					value, value_errors := column_data.GetInt()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "float32":
					value, value_errors := column_data.GetFloat32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=  fmt.Sprintf("%f", *value)
					}
				case "*float32":
					value, value_errors := column_data.GetFloat32()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=  fmt.Sprintf("%f", *value)
					}
				case "float64":
					value, value_errors := column_data.GetFloat64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=  fmt.Sprintf("%f", *value)
					}
				case "*float64":
					value, value_errors := column_data.GetFloat64()
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else if common.IsNil(value) {
						errors = append(errors,fmt.Errorf("*uint64 is nil"))
					} else {
						sql_command +=  fmt.Sprintf("%f", *value)
					}
				case "*time.Time":
					decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
					if decimal_places_error != nil {
						errors = append(errors, decimal_places_error...)
					} else if decimal_places == nil {
						errors = append(errors, fmt.Errorf("decimal_places is nil"))
					} else {
						value, value_errors := column_data.GetTimeWithDecimalPlaces(*decimal_places)
						if value_errors != nil {
							errors = append(errors, value_errors...)
							continue
						} else if common.IsNil(value) {
							errors = append(errors, fmt.Errorf("time is nil"))
							continue
						}

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
					decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
					if decimal_places_error != nil {
						errors = append(errors, decimal_places_error...)
					} else if decimal_places == nil {
						errors = append(errors, fmt.Errorf("decimal_places is nil"))
					} else {
						value, value_errors := column_data.GetTimeWithDecimalPlaces(*decimal_places)
						if value_errors != nil {
							errors = append(errors, value_errors...)
							continue
						} else if common.IsNil(value) {
							errors = append(errors, fmt.Errorf("time is nil"))
							continue
						}

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
				case "string":
					column_data_value, column_data_value_errors := column_data.GetString()
					if column_data_value_errors != nil {
						errors = append(errors, column_data_value_errors...)
						continue
					} else if common.IsNil(column_data_value) {
						errors = append(errors, fmt.Errorf("column data is nil"))
						continue
					}

					value_escaped, value_escaped_errors := common.EscapeString(*column_data_value, "'")
					if value_escaped_errors != nil {
						errors = append(errors, value_escaped_errors)
					}
					
					if options.IsBoolTrue("use_file") {
						sql_command += "'" + value_escaped + "'"
					} else {
						sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
					}
					
				case "*string":
					column_data_value, column_data_value_errors := column_data.GetString()
					if column_data_value_errors != nil {
						errors = append(errors, column_data_value_errors...)
						continue
					} else if common.IsNil(column_data_value) {
						errors = append(errors, fmt.Errorf("column data is nil"))
						continue
					}

					value_escaped, value_escaped_errors := common.EscapeString(*column_data_value, "'")
					if value_escaped_errors != nil {
						errors = append(errors, value_escaped_errors)
					}

					if options.IsBoolTrue("use_file") {
						sql_command += "'" + value_escaped + "'"
					} else {
						sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
					}
				
				case "bool":

					if column_data.IsBoolTrue() {
						sql_command += "1"
					} else {
						sql_command += "0"
					}
				case "*bool":
					if column_data.IsBoolTrue() {
						sql_command += "1"
					} else {
						sql_command += "0"
					}
				default:
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: filter type not supported please implement: %s", type_of))
				}
				sql_command += " "
			}

			if index < len(filters.GetKeys()) - 1 {
				sql_command += "AND "
			}
		}
	}

	if order_by_clause != "" {
		sql_command += "ORDER BY " + (order_by_clause + " ")
	}

	if limit != nil {
		limit_value := strconv.FormatUint(*limit, 10)
		sql_command += fmt.Sprintf("LIMIT %s ", limit_value)
	}

	if offset != nil {
		offset_value := strconv.FormatUint(*offset, 10)
		sql_command += fmt.Sprintf("OFFSET %s ", offset_value)
	}
	sql_command += ";"

	if len(errors) > 0 {
		return nil, nil, errors
	}

	return &sql_command, options, nil
}

