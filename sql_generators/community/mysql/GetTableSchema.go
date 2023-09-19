package mysql

import (
	"fmt"
	"strings"
	"strconv"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_validator/validate"
)

type TableSchemaSQL struct {
	GetTableSchemaSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	MapTableSchemaFromDB func(verify *validate.Validator, table_name string, json_array json.Array) (*json.Map, []error)
}

func newTableSchemaSQL() (*TableSchemaSQL) {
	get_table_schema_sql := func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
		var errors []error

		validation_errors := verify.ValidateTableName(table_name)
		if validation_errors != nil {
			return nil, options, validation_errors
		}

		table_name_escaped, table_name_escaped_error := common.EscapeString(table_name, "'")
		if table_name_escaped_error != nil {
			errors = append(errors, table_name_escaped_error)
			return nil, options, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString("SHOW FULL COLUMNS FROM ")
		Box(&sql_command, table_name_escaped,"`","`")
		sql_command.WriteString(";")
		
		return &sql_command, options, nil
	}

	return &TableSchemaSQL{
		GetTableSchemaSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_table_schema_sql(verify, table_name, options)
		},
		MapTableSchemaFromDB: func(verify *validate.Validator, table_name string, json_array json.Array) (*json.Map, []error) {
			var errors []error
			if common.IsNil(table_name) {
				errors = append(errors, fmt.Errorf("error: table_name is nil"))
			}

			if len(errors) > 0 {
				return nil, errors 
			}

			if len(*(json_array.GetValues())) == 0 {
				errors = append(errors, fmt.Errorf("error: show columns did not return any records"))
				return nil, errors
			}

			schema := json.NewMap()
			for _, column_details := range *(json_array.GetValues()) {
				column_map, column_map_errors := column_details.GetMap()
				if column_map_errors != nil {
					return nil, column_map_errors
				} else if common.IsNil(column_map) {
					errors = append(errors, fmt.Errorf("column_map is nil"))
					return nil, errors
				}
				column_attributes := column_map.GetKeys()

				column_schema := json.NewMap()
				default_value := ""
				field_name := ""
				is_nullable := false
				is_primary_key := false
				extra_value := ""
				comment_value := ""
				for _, column_attribute := range column_attributes {
					switch column_attribute {
					case "Key":
						key_value, _ := column_map.GetString("Key")
						switch *key_value {
						case "PRI":
							is_primary_key = true
							is_nullable = false
							column_schema.SetBoolValue("primary_key", true)
						case "", "MUL":
						case "UNI":
							column_schema.SetBoolValue("unique", true)
						default:
							errors = append(errors, fmt.Errorf("error: Table: GetSchema: Key not implemented please implement: %s", *key_value))
						}
					case "Field":
						field_name_value, _ := column_map.GetString("Field")
						field_name = *field_name_value
					case "Type":
						type_of_value, _ := column_map.GetString("Type")
						switch *type_of_value {
						case "bigint unsigned":
							column_schema.SetStringValue("type", "uint64")
							column_schema.SetBoolValue("unsigned", true)
						case "int unsigned":
							column_schema.SetStringValue("type", "uint32")
							column_schema.SetBoolValue("unsigned", true)
						case "mediumint unsigned":
							column_schema.SetStringValue("type", "uint32")
							column_schema.SetBoolValue("unsigned", true)
						case "smallint unsigned":
							column_schema.SetStringValue("type", "uint16")
							column_schema.SetBoolValue("unsigned", true)
						case "tinyint unsigned":
							column_schema.SetStringValue("type",  "uint8")
							column_schema.SetBoolValue("unsigned", true)
						case "bigint":
							column_schema.SetStringValue("type", "int64")
						case "int":
							column_schema.SetStringValue("type","int32")
						case "mediumint":
							column_schema.SetStringValue("type", "int32")
						case "smallint":
							column_schema.SetStringValue("type", "int16")
						case "tinyint":
							column_schema.SetStringValue("type", "int8")
						case "timestamp":
							column_schema.SetStringValue("type",  "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(0))
						case "timestamp(1)":
							column_schema.SetStringValue("type","time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(1))
						case "timestamp(2)":						
							column_schema.SetStringValue("type", "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(2))
						case "timestamp(3)":
							column_schema.SetStringValue("type", "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(3))
						case "timestamp(4)":
							column_schema.SetStringValue("type", "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(4))
						case "timestamp(5)":
							column_schema.SetStringValue("type", "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(5))
						case "timestamp(6)":
							column_schema.SetStringValue("type", "time.Time")
							column_schema.SetUInt8Value("decimal_places", uint8(6))
						case "tinyint(1)":
							column_schema.SetStringValue("type", "bool")
						case "text", "blob", "json":
							column_schema.SetStringValue("type", "string")
						case "float":
							column_schema.SetStringValue("type", "float32")
						case "double":
							column_schema.SetStringValue("type", "float64")
						default:
							if strings.HasPrefix(*type_of_value, "char(") && strings.HasSuffix(*type_of_value, ")") {
								column_schema.SetStringValue("type","string")
								*type_of_value = strings.TrimPrefix(*type_of_value, "char(")
								*type_of_value = strings.TrimSuffix(*type_of_value, ")")
								temp_max_length, temp_max_length_error := strconv.ParseInt(strings.TrimSpace(*type_of_value), 10, 64)
								if temp_max_length_error != nil {
									errors = append(errors, fmt.Errorf("failed to parse string max_length"))
								} else if temp_max_length <= 0{
									errors = append(errors, fmt.Errorf("failed to parse string max_length had value less than or equal to zero: %d", temp_max_length))
								} else {
									column_schema.SetInt64Value("max_length", temp_max_length)
								}
							} else if strings.HasPrefix(*type_of_value, "varchar(") && strings.HasSuffix(*type_of_value, ")") {
								column_schema.SetStringValue("type", "string")
								*type_of_value = strings.TrimPrefix(*type_of_value, "varchar(")
								*type_of_value = strings.TrimSuffix(*type_of_value, ")")
								temp_max_length, temp_max_length_error := strconv.ParseInt(strings.TrimSpace(*type_of_value), 10, 64)
								if temp_max_length_error != nil {
									errors = append(errors, fmt.Errorf("failed to parse string max_length"))
								} else if temp_max_length <= 0{
									errors = append(errors, fmt.Errorf("failed to parse string max_length had value less than or equal to zero: %d", temp_max_length))
								} else {
									column_schema.SetInt64Value("max_length", temp_max_length)
								}
							} else if strings.HasPrefix(*type_of_value, "enum(")  && strings.HasSuffix(*type_of_value, ")") {
								type_of_value_values := (*type_of_value)[5:len(*type_of_value)-1]
								parts := strings.Split(type_of_value_values, ",")
								if len(parts) == 0 {
									errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum had length of zero: %s", *type_of_value))
								} else {
									part := parts[0]
									if strings.HasPrefix(part, "'")  && strings.HasSuffix(part, "'") {
										column_schema.SetStringValue("type", "string")
									} else {
										errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum for data type: %s", *type_of_value))
									}
								}
							} else {
								errors = append(errors, fmt.Errorf("error: Table: GetSchema: type not implemented please implement: %s", *type_of_value))
							}
						}
					case "Null":
						null_value, _ := column_map.GetStringValue("Null")
						switch null_value {
						case "YES":
							if !is_primary_key {
								is_nullable = true
							}
						case "NO":
							is_nullable = false
						default:
							errors = append(errors, fmt.Errorf("error: Table: GetSchema: Null value not supported please implement: %s", null_value))
						}
					case "Default":
						default_val, _ := column_map.GetStringValue("Default")
						default_value = default_val
					case "Extra":
						extra_val, _ := column_map.GetStringValue("Extra")
						extra_value = extra_val
						switch extra_value {
						case "auto_increment":
							column_schema.SetBoolValue("auto_increment",true)
						case "DEFAULT_GENERATED":
						case "":
						default:
							errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: Extra value not supported please implement: %s", table_name, extra_value))
						}
					case "Privileges":
					case "Collation":
					case "Comment":
						comment_val, comment_errors := column_map.GetString("Comment")
						if comment_errors != nil {
							errors = append(errors, comment_errors...)
						} else {
							comment_value = *comment_val
							if strings.TrimSpace(comment_value) != "" {
								comment_as_map, comment_as_map_value_errors := json.Parse(strings.TrimSpace(comment_value))
								if comment_as_map_value_errors != nil {
									errors = append(errors, comment_as_map_value_errors...)
								} else if common.IsNil(comment_as_map) {
									errors = append(errors, fmt.Errorf("comment is nil"))
								} else {
									rules_array, rules_array_errors := comment_as_map.GetArray("rules")
									if rules_array_errors != nil {
										errors = append(errors, rules_array_errors...)
									} else if !common.IsNil(rules_array) {
										filters := json.NewArrayValue()
										for _, rule := range *(rules_array.GetValues()) {
											rule_value, rule_value_errors := rule.GetString()
											if rule_value_errors != nil {
												return nil, rule_value_errors
											} else if common.IsNil(rule_value) {
												errors = append(errors, fmt.Errorf("rule value is nil"))
												return nil, errors
											}

											switch *rule_value {
											case "domain_name":
												domain_name_filter := json.NewMap()
												domain_name_filter.SetObjectForMap("function", verify.GetValidateDatabaseNameFunc())
												filters.AppendMap(domain_name_filter)
											case "repository_name":
												repostiory_name_filter := json.NewMap()
												repostiory_name_filter.SetObjectForMap("function", verify.GetValidateRepositoryNameFunc())
												filters.AppendMap(repostiory_name_filter)
											case "repository_account_name":
												repository_account_name_filter := json.NewMap()
												repository_account_name_filter.SetObjectForMap("function", verify.GetValidateRepositoryAccountNameFunc())
												filters.AppendMap(repository_account_name_filter)
											case "branch_name":
												branch_name_filter := json.NewMap()
												branch_name_filter.SetObjectForMap("function", verify.GetValidateBranchNameFunc())
												filters.AppendMap(branch_name_filter)
											default:
												errors = append(errors, fmt.Errorf("rule not supported %s", rule_value))
											}
										}
										column_schema.SetArray("filters", &filters)
									}

									foreign_key_map, foreign_key_map_errors := comment_as_map.GetMap("foreign_key")
									if foreign_key_map_errors != nil {
										errors = append(errors, foreign_key_map_errors...)
									} else if !common.IsNil(foreign_key_map) {
										column_schema.SetBoolValue("foreign_key", true)

										foreign_key_table_name, foreign_key_table_name_errors := foreign_key_map.GetStringValue("table_name")
										if foreign_key_table_name_errors != nil {
											errors = append(errors, foreign_key_table_name_errors...)
										} else if common.IsNil(foreign_key_table_name) {
											errors = append(errors, fmt.Errorf("foreign_key table_name is nil"))
										} else {
											column_schema.SetStringValue("foreign_key_table_name", foreign_key_table_name)
										}

										foreign_key_column_name, foreign_key_column_name_errors := foreign_key_map.GetStringValue("column_name")
										if foreign_key_column_name_errors != nil {
											errors = append(errors, foreign_key_column_name_errors...)
										} else if common.IsNil(foreign_key_column_name) {
											errors = append(errors, fmt.Errorf("foreign_key column_name is nil"))
										} else {
											column_schema.SetStringValue("foreign_key_column_name", foreign_key_column_name)
										}

										foreign_key_type, foreign_key_type_errors := foreign_key_map.GetStringValue("type")
										if foreign_key_type_errors != nil {
											errors = append(errors, foreign_key_type_errors...)
										} else if common.IsNil(foreign_key_type) {
											errors = append(errors, fmt.Errorf("foreign_key type is nil"))
										} else {
											column_schema.SetStringValue("foreign_key_type", foreign_key_type)
										}
									}

									foreign_keys_array, foreign_keys_array_errors := comment_as_map.GetArray("foreign_keys")
									if foreign_keys_array_errors != nil {
										errors = append(errors, foreign_keys_array_errors...)
									} else if !common.IsNil(foreign_keys_array) {
										column_schema.SetArray("foreign_keys", foreign_keys_array)

										foreign_keys_count := foreign_keys_array.Len()
										foreign_keys_count_index := 0
										for foreign_keys_count_index < foreign_keys_count {
											foreign_keys_map, foreign_keys_map_errors := foreign_keys_array.GetMap(foreign_keys_count_index)
											if foreign_keys_map_errors != nil {
												errors = append(errors, foreign_keys_map_errors...)
											} else if common.IsNil(foreign_keys_map) {
												errors = append(errors, fmt.Errorf("foreign_keys_map is nil"))
											} 

											if len(errors) > 0 {
												foreign_keys_count_index++
												continue
											}

											foreign_key_table_name, foreign_key_table_name_errors := foreign_keys_map.GetStringValue("table_name")

											if foreign_key_table_name_errors != nil {
												errors = append(errors, foreign_key_table_name_errors...)
											} else if common.IsNil(foreign_key_table_name) {
												errors = append(errors, fmt.Errorf("foreign_key table_name is nil"))
											} 

											foreign_key_column_name, foreign_key_column_name_errors := foreign_keys_map.GetStringValue("column_name")
											if foreign_key_column_name_errors != nil {
												errors = append(errors, foreign_key_column_name_errors...)
											} else if common.IsNil(foreign_key_column_name) {
												errors = append(errors, fmt.Errorf("foreign_key column_name is nil"))
											} 

											foreign_key_type, foreign_key_type_errors := foreign_keys_map.GetStringValue("type")
											if foreign_key_type_errors != nil {
												errors = append(errors, foreign_key_type_errors...)
											} else if common.IsNil(foreign_key_type) {
												errors = append(errors, fmt.Errorf("foreign_key type is nil"))
											} 
											foreign_keys_count_index++
										}

										
									}
								}
							}
						}
					default:
						errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: %s not supported please implement", table_name, field_name, column_attribute))
					}
				}

				if column_schema.IsNull("type") {
					errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: type is nill", table_name, field_name))
				}

				if len(errors) > 0 {
					continue
				}

				dt, _ := column_schema.GetStringValue("type")

			
				if default_value == "NULL" {
					if dt == "string" {
						column_schema.SetStringValue("default", "")
					} else {
						column_schema.SetNil("default")
					}
				} else {
					if dt == "string" {
						column_schema.SetStringValue("default", default_value)
					} else if dt == "uint64" && default_value != "" {
						number, err := strconv.ParseUint(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							column_schema.SetUInt64Value("default", number)
						}
					} else if dt == "int64" && default_value != "" {
						number, err := strconv.ParseInt(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							column_schema.SetInt64Value("default", number)
						}
					} else if dt == "uint32" && default_value != "" {
						number, err := strconv.ParseUint(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := uint32(number)
							column_schema.SetUInt32Value("default", converted)
						}
					} else if dt == "int32" && default_value != "" {
						number, err := strconv.ParseInt(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := int32(number)
							column_schema.SetInt32Value("default", converted)
						}
					} else if dt == "uint16" && default_value != "" {
						number, err := strconv.ParseUint(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := uint16(number)
							column_schema.SetUInt16Value("default", converted)
						}
					} else if dt == "int16" && default_value != "" {
						number, err := strconv.ParseInt(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := int16(number)
							column_schema.SetInt16Value("default", converted)
						}
					} else if dt == "uint8" && default_value != "" {
						number, err := strconv.ParseUint(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := uint8(number)
							column_schema.SetUInt8Value("default", converted)
						}
					} else if dt == "int8" && default_value != "" {
						number, err := strconv.ParseInt(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := int8(number)
							column_schema.SetInt8Value("default", converted)
						}
					} else if dt == "float32" && default_value != "" {
						number, err := strconv.ParseFloat(default_value, 32)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := float32(number)
							column_schema.SetFloat32Value("default", converted)
						}
					} else if dt == "float64" && default_value != "" {
						number, err := strconv.ParseFloat(default_value, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							converted := float64(number)
							column_schema.SetFloat64Value("default", converted)
						}
					}  else if dt == "bool" && default_value != "" {
						number, err := strconv.ParseInt(default_value, 10, 64)
						if err != nil {
							errors = append(errors, err)
						} else {
							if number == 0 {
								column_schema.SetBoolValue("default", false)
							} else if number == 1 {
								column_schema.SetBoolValue("default", true)
							} else {
								errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 1 or 0", default_value, dt))
							}
						}
					} else if dt == "time.Time" && default_value != "" {
						if extra_value == "DEFAULT_GENERATED" && strings.HasPrefix(default_value, "CURRENT_TIMESTAMP") {
							if default_value == "CURRENT_TIMESTAMP" {
							} else if default_value == "CURRENT_TIMESTAMP(1)" {
							} else if default_value == "CURRENT_TIMESTAMP(2)" {
							} else if default_value == "CURRENT_TIMESTAMP(3)" {
							} else if default_value == "CURRENT_TIMESTAMP(4)" {
							} else if default_value == "CURRENT_TIMESTAMP(5)" {
							} else if default_value == "CURRENT_TIMESTAMP(6)" {
							} else {
								errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 0-6 decimal places", default_value, dt))
							}
							column_schema.SetStringValue("default",  "now")
						} else if default_value == "0000-00-00 00:00:00" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.0" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.00" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.000" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.0000" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.00000" {
							column_schema.SetStringValue("default", "zero")
						} else if default_value == "0000-00-00 00:00:00.000000" {
							column_schema.SetStringValue("default", "zero")
						}  else {
							errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be DEFAULT_GENERATED or 0000-00-00 00:00:00", default_value, dt))
						}
					} else if !(dt == "time.Time" || dt == "bool" || dt == "int64" || dt == "uint64" ||  dt == "int32" || dt == "uint32" ||  dt == "int16" || dt == "uint16" ||  dt == "int8" || dt == "uint8" || dt == "string" || dt == "float32" || dt == "float64") && default_value != "" {
						errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported please implement: %s for type: %s", default_value, dt))
					}
				}
				

				if is_nullable {
					column_schema.SetStringValue("type",  "*" + dt)
				}

				schema.SetMap(field_name, column_schema)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return schema, nil
		},
	}
}
