package mysql

import (
	"fmt"
	"strconv"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_validator/validate"
)

type SelectRecordsSQL struct {
	GetSelectRecordsSQL func(verify *validate.Validator, table_name string, table_data json.Map, select_fields *json.Array, filters *json.Map, filters_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64, options json.Map) (*strings.Builder, json.Map, []error)
}

func newSelectRecordsSQL() (*SelectRecordsSQL) {
	get_select_records_sql := func(verify *validate.Validator, table_name string, table_data json.Map, select_fields *json.Array, filters *json.Map, filters_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error
	
		table_schema, table_schema_errors := helper.GetSchemas(table_data, "[schema]")
		if table_schema_errors != nil {
			errors = append(errors, table_schema_errors...)
		} else if common.IsNil(table_schema) {
			errors = append(errors, fmt.Errorf("table_schema is nil"))
		}

		table_columns, table_columns_errors := helper.GetTableColumns(table_data)
		if table_columns_errors != nil {
			errors = append(errors, table_columns_errors...)
		} else if common.IsNil(table_columns) {
			errors = append(errors, fmt.Errorf("table_columns is nil"))
		}

		if len(errors) > 0 {
			return nil, options, errors
		}
		
		for table_column, _  := range *table_columns {
			table_column_errors := verify.ValidateColumnName(table_column)
			if table_column_errors != nil {
				errors = append(errors, table_column_errors...)
			}
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
		}

		if len(errors) > 0 {
			return nil, options, errors
		}

		table_name_validation_errors := verify.ValidateTableName(table_name)
		if table_name_validation_errors != nil {
			errors = append(errors, table_name_validation_errors...)
		}

		if len(errors) > 0 {
			return nil, options, errors
		}

		if filters != nil {
			filter_columns := filters.GetKeys()
			for _, filter_column := range filter_columns {
				
				if filter_column == GetCountColumnNameSQLMySQL() {
					continue
				} 

				if _, found := (*table_columns)[filter_column]; !found {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", filter_column, table_name_escaped, *table_columns))
				}
			}

			if len(errors) > 0 {
				return nil, options, errors
			}


			for _, filter_column := range filter_columns {
				filter_column_type := filters.GetType(filter_column)

				if !filters.IsNull(filter_column) && !strings.HasPrefix(filter_column_type, "*") {
					filter_column_type = "*" + filter_column_type
				}
					
				if table_schema.IsNull(filter_column) {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s does not exist however filter had the value, table has columns: %s", filter_column, table_name_escaped, table_schema.GetKeys()))
					continue
				}

				table_schema_column, table_schema_column_errors := table_schema.GetMap(filter_column)
				if table_schema_column_errors != nil {
					errors = append(errors, table_schema_column_errors...)
					continue
				}

				if table_schema_column.IsNull("type") {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s did not have atrribute: type", filter_column, table_name_escaped))
					continue
				}


				table_column_type, _ := (*table_schema_column).GetString("type")
				if strings.Replace(*table_column_type, "*", "", -1) != strings.Replace(filter_column_type, "*", "", -1) {
					table_column_type_simple := strings.Replace(*table_column_type, "*", "", -1)
					filter_column_type_simple := strings.Replace(filter_column_type, "*", "", -1)
					if strings.Contains(table_column_type_simple, "int") && strings.Contains(filter_column_type_simple, "int") {

					} else if strings.Contains(table_column_type_simple, "float") && strings.Contains(filter_column_type_simple, "float"){

					} else if strings.Contains(table_column_type_simple, "int") && strings.Contains(filter_column_type_simple, "json.Array"){
						//todo validate array field values
					} else {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, table_name_escaped, *table_column_type))
					}
					

					//todo ignore if filter data_type is nil and table column allows nil
				}
			}
		}

		if select_fields != nil {
			for _, select_field := range *(select_fields.GetValues()) {
				select_field_value, select_field_value_errors := select_field.GetString()
				if select_field_value_errors != nil {
					return nil, options, select_field_value_errors
				} else if common.IsNil(select_field_value) {
					errors = append(errors, fmt.Errorf("select_field_value is nil"))
					return nil, options, errors
				}

				if *select_field_value == GetCountColumnNameSQLMySQL() {
					continue
				}

				if _, found := (*table_columns)[*select_field_value]; !found {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", *select_field_value, table_name_escaped, *table_columns))
				}			
			}
		}

		if filters_logic  != nil {
			temp_filters_fields := filters_logic.GetKeys()
			for _, temp_filters_field := range temp_filters_fields {

				if temp_filters_field != GetCountColumnNameSQLMySQL() {
					if _, found :=(*table_columns)[temp_filters_field]; !found {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", temp_filters_field, table_name_escaped, *table_columns))
					}
				}

				if !filters_logic.IsString(temp_filters_field) {
					errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter logic is not a string", temp_filters_field, table_name_escaped))
				} else {
					temp_logic_value, temp_logic_value_errors := filters_logic.GetString(temp_filters_field)
					if temp_logic_value_errors != nil {
						errors = append(errors, temp_logic_value_errors...)
					} else if common.IsNil(temp_logic_value) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value had error getting string value", temp_filters_field, table_name_escaped))
					} else {
						if !(*temp_logic_value != "=" ||
							*temp_logic_value != ">" ||
							*temp_logic_value != "<" ||
							*temp_logic_value != "in") {
							errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value logic however it's not supported please implement: %s", temp_filters_field, table_name_escaped, *temp_logic_value))
						} 
					}
				}
			}
		}

		order_by_clause_result := ""
		if !common.IsNil(order_by) && len(*(order_by.GetValues())) > 0 {
			var order_by_clause strings.Builder
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
				
				if order_by_column_name != GetCountColumnNameSQLMySQL() {
					if _, found := (*table_columns)[order_by_column_name]; !found {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: order by column: %s not found for table: %s available columns are: %s", order_by_column_name, table_name_escaped, *table_columns))
					}
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
				
				
				escaped_order_by_column_name, escaped_order_by_column_name_errors := common.EscapeString(order_by_column_name, "'")
				if escaped_order_by_column_name_errors != nil {
					errors = append(errors, escaped_order_by_column_name_errors)
					continue
				}

				Box(&order_by_clause, escaped_order_by_column_name,"`","`")
				order_by_clause.WriteString(" ")
				order_by_clause.WriteString(order_by_string_value_validated)

				if order_by_index < (order_by_columns - 1) {
					order_by_clause.WriteString(", ")
				} else {
					order_by_clause.WriteString(" ")
				}
			}
			order_by_clause_result = order_by_clause.String()
		}

		group_by_clause_result := ""
		if group_by != nil && len(*(group_by.GetValues())) > 0 {
			var group_by_clause strings.Builder
			group_by_columns := len(*(group_by.GetValues()))
			for group_by_index, group_by_field := range *(group_by.GetValues()) {
				group_by_field_value, group_by_field_value_errors := group_by_field.GetString()
				if group_by_field_value_errors != nil {
					return nil, options, group_by_field_value_errors
				} else if common.IsNil(group_by_field_value) {
					errors = append(errors, fmt.Errorf("group_by_field_value is nil"))
					return nil, options, errors
				}

				if *group_by_field_value != GetCountColumnNameSQLMySQL() {
					if _, found := (*table_columns)[*group_by_field_value]; !found {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", *group_by_field_value, table_name_escaped, *table_columns))
					}	
				}

				escaped_group_by_column_name, escaped_group_by_column_name_errors := common.EscapeString(*group_by_field_value, "'")
				if escaped_group_by_column_name_errors != nil {
					errors = append(errors, escaped_group_by_column_name_errors)
					continue
				}
				
				Box(&group_by_clause, escaped_group_by_column_name,"`","`")

				if group_by_index < (group_by_columns - 1) {
					group_by_clause.WriteString(", ")
				} else {
					group_by_clause.WriteString(" ")
				}
			}
			group_by_clause_result = group_by_clause.String()
		}


		if len(errors) > 0 {
			return nil, options, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString("SELECT ")
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
					if escape_string_value != GetCountColumnNameSQLMySQL() {
						Box(&sql_command, escape_string_value,"`","`")			
					} else {
						sql_command.WriteString(GetCountColumnNameSQLMySQL())		
					}
					if i < (select_fields_values_length - 1) {
						sql_command.WriteString(", ")
					} else {
						sql_command.WriteString(" ")
					}
				}
			}
		} else {
			sql_command.WriteString("* ")
		}

		sql_command.WriteString("FROM ")
		Box(&sql_command, table_name_escaped,"`","`")
		sql_command.WriteString(" ")

		if filters != nil && len(filters.GetKeys()) > 0 {
			sql_command.WriteString("WHERE ")
			
			for index, column_filter := range filters.GetKeys() {				
				column_definition, column_definition_errors := table_schema.GetMap(column_filter)
				if column_definition_errors != nil {
					errors = append(errors, column_definition_errors...) 
					continue
				}
				
				column_name_errors := verify.ValidateColumnName(column_filter)
				if column_name_errors != nil {
					errors = append(errors, column_name_errors...)
				}	

				column_filter_escaped, column_filter_escaped_errors := common.EscapeString(column_filter, "'")
				if table_name_escaped_errors != nil {
					errors = append(errors, column_filter_escaped_errors)
				}

				Box(&sql_command, column_filter_escaped,"`","`")			

				if common.IsNil(filters_logic) {
					sql_command.WriteString(" = ")
				} else if !filters_logic.HasKey(column_filter) {
					sql_command.WriteString(" = ")
				} else {
					filters_logic_temp, filters_logic_temp_errors := filters_logic.GetString(column_filter)
					if filters_logic_temp_errors != nil {
						errors = append(errors, filters_logic_temp_errors...)
					} else if common.IsNil(filters_logic_temp) {
						errors = append(errors, fmt.Errorf("filters logic is nil"))
					} else {
						sql_command.WriteString(" ")
						sql_command.WriteString(*filters_logic_temp)
						sql_command.WriteString(" ")
					}
				}

				if filters.IsNull(column_filter) {
					sql_command.WriteString("NULL ")
				} else {
					if filters.IsArray(column_filter) {
						sql_command.WriteString("( ")
					}

					var array_value_errors []error
					array_values := json.NewArray()
					
					if !filters.IsArray(column_filter) { 
						array_values.AppendValue(filters.GetValue(column_filter))
					} else {
						array_values, array_value_errors = filters.GetArray(column_filter)
						if len(array_value_errors) > 0 {
							errors = append(errors, array_value_errors...)
							continue
						}
					}

					for index, column_data := range *(array_values.GetValues()) {
						//todo check data type with schema
						type_of := column_data.GetType()
						//type_of := filters.GetType(column_filter)
						//column_data := filters.GetValue(column_filter)
						switch type_of {
						case "*uint64":
							value, value_errors := column_data.GetUInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(*value, 10))
							}
						case "uint64":
							value, value_errors := column_data.GetUInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(*value, 10))
							}
						case "*int64":
							value, value_errors := column_data.GetInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "int64":
							value, value_errors := column_data.GetInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "*uint32":
							value, value_errors := column_data.GetUInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "uint32":
							value, value_errors := column_data.GetUInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "*int32":
							value, value_errors := column_data.GetInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "int32":
							value, value_errors := column_data.GetInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "*uint16":
							value, value_errors := column_data.GetUInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "uint16":
							value, value_errors := column_data.GetUInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "*int16":
							value, value_errors := column_data.GetInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "int16":
							value, value_errors := column_data.GetInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "*uint8":
							value, value_errors := column_data.GetUInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "uint8":
							value, value_errors := column_data.GetUInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
							}
						case "*int8":
							value, value_errors := column_data.GetInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "int8":
							value, value_errors := column_data.GetInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "*int":
							value, value_errors := column_data.GetInt()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "int":
							value, value_errors := column_data.GetInt()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
							}
						case "float32":
							value, value_errors := column_data.GetFloat32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								value_string :=fmt.Sprintf("%f", *value)
								sql_command.WriteString(value_string)
								if !strings.Contains(value_string, ".") {
									sql_command.WriteString(".0")
								}
							}
						case "*float32":
							value, value_errors := column_data.GetFloat32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								value_string := fmt.Sprintf("%f", *value)
								sql_command.WriteString(value_string)
								if !strings.Contains(value_string, ".") {
									sql_command.WriteString(".0")
								}
							}
						case "float64":
							value, value_errors := column_data.GetFloat64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								value_string :=fmt.Sprintf("%f", *value)
								sql_command.WriteString(value_string)
								if !strings.Contains(value_string, ".") {
									sql_command.WriteString(".0")
								}
							}
						case "*float64":
							value, value_errors := column_data.GetFloat64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								value_string :=fmt.Sprintf("%f", *value)
								sql_command.WriteString(value_string)
								if !strings.Contains(value_string, ".") {
									sql_command.WriteString(".0")
								}
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
			
									sql_command.WriteString("'")
									sql_command.WriteString(value_escaped)
									sql_command.WriteString("'")
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
			
									sql_command.WriteString("'")
									sql_command.WriteString(value_escaped)
									sql_command.WriteString("'")
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
							
							sql_command.WriteString("'")
							sql_command.WriteString(value_escaped)
							sql_command.WriteString("'")
							
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

							sql_command.WriteString("'")
							sql_command.WriteString(value_escaped)
							sql_command.WriteString("'")
						
						case "bool":

							if column_data.IsBoolTrue() {
								sql_command.WriteString("1")
							} else {
								sql_command.WriteString("0")
							}
						case "*bool":
							if column_data.IsBoolTrue() {
								sql_command.WriteString("1")
							} else {
								sql_command.WriteString("0")
							}
						default:
							errors = append(errors, fmt.Errorf("error: Table.ReadRecords: filter type not supported please implement: %s", type_of))
						}
						sql_command.WriteString(" ")

						if index < len(*(array_values.GetValues())) - 1 {
							sql_command.WriteString(", ")
						}
					}
				}

				if filters.IsArray(column_filter) {
					sql_command.WriteString(") ")
				}

				if index < len(filters.GetKeys()) - 1 {
					sql_command.WriteString("AND ")
				}
			}
		}

		if group_by_clause_result != "" {
			sql_command.WriteString("GROUP BY ")
			sql_command.WriteString(group_by_clause_result)
			sql_command.WriteString(" ")
		}

		if order_by_clause_result != "" {
			sql_command.WriteString("ORDER BY ")
			sql_command.WriteString(order_by_clause_result)
			sql_command.WriteString(" ")
		}

		if limit != nil {
			limit_value := strconv.FormatUint(*limit, 10)
			sql_command.WriteString("LIMIT ")
			sql_command.WriteString(limit_value)
			sql_command.WriteString(" ")
		}

		if offset != nil {
			offset_value := strconv.FormatUint(*offset, 10)
			sql_command.WriteString("OFFSET ")
			sql_command.WriteString(offset_value)
			sql_command.WriteString(" ")
		}
		sql_command.WriteString(";")

		if len(errors) > 0 {
			return nil, options, errors
		}

		return &sql_command, options, nil
	}

	return &SelectRecordsSQL{
		GetSelectRecordsSQL: func(verify *validate.Validator, table_name string, table_data json.Map, select_fields *json.Array, filters *json.Map, filters_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_select_records_sql(verify, table_name, table_data, select_fields, filters, filters_logic, group_by, order_by, limit, offset, options)
		},
	}
}