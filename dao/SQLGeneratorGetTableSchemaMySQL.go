package dao

import (
	"fmt"
	"strings"
	"strconv"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

func getTableSchemaSQLMySQL(struct_type string, table *Table, options *json.Map) (*string, *json.Map, []error) {
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

	if common.IsNil(options) {
		options = json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)
	}

	temp_table_name, temp_table_name_errors := table.GetTableName()
	if temp_table_name_errors != nil {
		return nil, nil, temp_table_name_errors
	}

	table_name_escaped, table_name_escaped_error := common.EscapeString(temp_table_name, "'")
	if table_name_escaped_error != nil {
		errors = append(errors, table_name_escaped_error)
		return nil, nil, errors
	}

	sql_command := "SHOW FULL COLUMNS FROM "
	if options.IsBoolTrue("use_file") {
		sql_command += fmt.Sprintf("`%s`;", table_name_escaped)
	} else {
		sql_command += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
	}

	return &sql_command, options, nil
}


func mapTableSchemaFromDBMySQL(struct_type string, table *Table, json_array *json.Array) (*json.Map, []error) {
	var errors []error

	if common.IsNil(table) {
		errors = append(errors, fmt.Errorf("table is nil"))
	} else {
		table_validation_errors := table.Validate() 
		if table_validation_errors != nil {
			errors = append(errors, table_validation_errors...)
		}
	}

	if common.IsNil(json_array) {
		errors = append(errors, fmt.Errorf("error: show columns returned nil records"))
	}

	if len(errors) > 0 {
		return nil, errors 
	}

	if len(*(json_array.GetValues())) == 0 {
		errors = append(errors, fmt.Errorf("error: show columns did not return any records"))
		return nil, errors
	}

	table_name, table_name_errors := table.GetTableName()
	if table_name_errors != nil {
		return nil, table_name_errors
	}

	schema := json.NewMapValue()
	for _, column_details := range *(json_array.GetValues()) {
		column_map, column_map_errors := column_details.GetMap()
		if column_map_errors != nil {
			return nil, column_map_errors
		} else if common.IsNil(column_map) {
			errors = append(errors, fmt.Errorf("column_map is nil"))
			return nil, errors
		}
		column_attributes := column_map.GetKeys()

		column_schema := json.NewMapValue()
		default_value := ""
		field_name := ""
		is_nullable := false
		is_primary_key := false
		is_unique := false
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
					column_schema.SetBool("primary_key", &is_primary_key)
				case "", "MUL":
				case "UNI":
					is_unique = true
					column_schema.SetBool("unique", &is_unique)
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
					data_type := "uint64"
					unsigned := true
					column_schema.SetString("type", &data_type)
					column_schema.SetBool("unsigned", &unsigned)
				case "int unsigned":
					data_type := "uint32"
					unsigned := true
					column_schema.SetString("type", &data_type)
					column_schema.SetBool("unsigned", &unsigned)
				case "mediumint unsigned":
					data_type := "uint32"
					unsigned := true
					column_schema.SetString("type", &data_type)
					column_schema.SetBool("unsigned", &unsigned)
				case "smallint unsigned":
					data_type := "uint16"
					unsigned := true
					column_schema.SetString("type", &data_type)
					column_schema.SetBool("unsigned", &unsigned)
				case "tinyint unsigned":
					data_type := "uint8"
					unsigned := true
					column_schema.SetString("type", &data_type)
					column_schema.SetBool("unsigned", &unsigned)
				case "bigint":
					data_type := "int64"
					column_schema.SetString("type", &data_type)
				case "int":
					data_type := "int32"
					column_schema.SetString("type", &data_type)
				case "mediumint":
					data_type := "int32"
					column_schema.SetString("type", &data_type)
				case "smallint":
					data_type := "int16"
					column_schema.SetString("type", &data_type)
				case "tinyint":
					data_type := "int8"
					column_schema.SetString("type", &data_type)
				case "timestamp":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(0))
				case "timestamp(1)":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(1))
				case "timestamp(2)":						
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(2))
				case "timestamp(3)":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(3))
				case "timestamp(4)":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(4))
				case "timestamp(5)":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(5))
				case "timestamp(6)":
					data_type := "time.Time"
					column_schema.SetString("type", &data_type)
					column_schema.SetUInt8Value("decimal_places", uint8(6))
				case "tinyint(1)":
					data_type := "bool"
					column_schema.SetString("type", &data_type)
				case "text", "blob", "json":
					data_type := "string"
					column_schema.SetString("type", &data_type)
				case "float":
					data_type := "float32"
					column_schema.SetString("type", &data_type)
				case "double":
					data_type := "float64"
					column_schema.SetString("type", &data_type)
				default:
					if strings.HasPrefix(*type_of_value, "char(") && strings.HasSuffix(*type_of_value, ")") {
						data_type := "string"
						column_schema.SetString("type", &data_type)
					} else if strings.HasPrefix(*type_of_value, "varchar(") && strings.HasSuffix(*type_of_value, ")") {
						data_type := "string"
						column_schema.SetString("type", &data_type)
					} else if strings.HasPrefix(*type_of_value, "enum(")  && strings.HasSuffix(*type_of_value, ")") {
						type_of_value_values := (*type_of_value)[5:len(*type_of_value)-1]
						parts := strings.Split(type_of_value_values, ",")
						if len(parts) == 0 {
							errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum had length of zero: %s", *type_of_value))
						} else {
							part := parts[0]
							if strings.HasPrefix(part, "'")  && strings.HasSuffix(part, "'") {
								data_type := "string"
								column_schema.SetString("type", &data_type)
							} else {
								errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum for data type: %s", *type_of_value))
							}
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table: GetSchema: type not implemented please implement: %s", *type_of_value))
					}
				}
			case "Null":
				null_value, _ := column_map.GetString("Null")
				switch *null_value {
				case "YES":
					if !is_primary_key {
						is_nullable = true
					}
				case "NO":
					is_nullable = false
				default:
					errors = append(errors, fmt.Errorf("error: Table: GetSchema: Null value not supported please implement: %s", *null_value))
				}
			case "Default":
				default_val, _ := column_map.GetString("Default")
				default_value = *default_val
			case "Extra":
				extra_val, _ := column_map.GetString("Extra")
				extra_value = *extra_val
				switch extra_value {
				case "auto_increment":
					auto_increment := true
					column_schema.SetBool("auto_increment", &auto_increment)
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
										domain_name_filter := json.NewMapValue()
										domain_name_filter.SetObjectForMap("values", validation_constants.GetValidDomainNameCharacters())
										domain_name_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
										filters.AppendMapValue(domain_name_filter)
									case "repository_name":
										repostiory_name_filter := json.NewMapValue()
										repostiory_name_filter.SetObjectForMap("values", validation_constants.GetValidRepositoryNameCharacters())
										repostiory_name_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
										filters.AppendMapValue(repostiory_name_filter)
									case "repository_account_name":
										repository_account_name_filter := json.NewMapValue()
										repository_account_name_filter.SetObjectForMap("values", validation_constants.GetValidRepositoryAccountNameCharacters())
										repository_account_name_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
										filters.AppendMapValue(repository_account_name_filter)
									case "branch_name":
										branch_name_filter := json.NewMapValue()
										branch_name_filter.SetObjectForMap("values", validation_constants.GetValidBranchNameCharacters())
										branch_name_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
										filters.AppendMapValue(branch_name_filter)
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
								foreign_key_true := true
								column_schema.SetBool("foreign_key", &foreign_key_true)

								foreign_key_table_name, foreign_key_table_name_errors := foreign_key_map.GetString("table_name")
								if foreign_key_table_name_errors != nil {
									errors = append(errors, foreign_key_table_name_errors...)
								} else if common.IsNil(foreign_key_table_name) {
									errors = append(errors, fmt.Errorf("foreign_key table_name is nil"))
								} else {
									column_schema.SetString("foreign_key_table_name", foreign_key_table_name)
								}

								foreign_key_column_name, foreign_key_column_name_errors := foreign_key_map.GetString("column_name")
								if foreign_key_column_name_errors != nil {
									errors = append(errors, foreign_key_column_name_errors...)
								} else if common.IsNil(foreign_key_column_name) {
									errors = append(errors, fmt.Errorf("foreign_key column_name is nil"))
								} else {
									column_schema.SetString("foreign_key_column_name", foreign_key_column_name)
								}

								foreign_key_type, foreign_key_type_errors := foreign_key_map.GetString("type")
								if foreign_key_type_errors != nil {
									errors = append(errors, foreign_key_type_errors...)
								} else if common.IsNil(foreign_key_type) {
									errors = append(errors, fmt.Errorf("foreign_key type is nil"))
								} else {
									column_schema.SetString("foreign_key_type", foreign_key_type)
								}
							}
						}
					}
				}
			default:
				errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: %s not supported please implement", table_name, field_name, column_attribute))
			}
		}

		if column_schema.IsNil("type") {
			errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: type is nill", table_name, field_name))
		}

		if len(errors) > 0 {
			continue
		}

		dt, _ := column_schema.GetString("type")

	
		if default_value == "NULL" {
			column_schema.SetNil("default")
		} else {
			if *dt == "string" {
				column_schema.SetString("default", &default_value)
			} else if *dt == "uint64" && default_value != "" {
				number, err := strconv.ParseUint(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					column_schema.SetUInt64("default", &number)
				}
			} else if *dt == "int64" && default_value != "" {
				number, err := strconv.ParseInt(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					column_schema.SetInt64("default", &number)
				}
			} else if *dt == "uint32" && default_value != "" {
				number, err := strconv.ParseUint(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := uint32(number)
					column_schema.SetUInt32("default", &converted)
				}
			} else if *dt == "int32" && default_value != "" {
				number, err := strconv.ParseInt(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := int32(number)
					column_schema.SetInt32("default", &converted)
				}
			} else if *dt == "uint16" && default_value != "" {
				number, err := strconv.ParseUint(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := uint16(number)
					column_schema.SetUInt16("default", &converted)
				}
			} else if *dt == "int16" && default_value != "" {
				number, err := strconv.ParseInt(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := int16(number)
					column_schema.SetInt16("default", &converted)
				}
			} else if *dt == "uint8" && default_value != "" {
				number, err := strconv.ParseUint(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := uint8(number)
					column_schema.SetUInt8("default", &converted)
				}
			} else if *dt == "int8" && default_value != "" {
				number, err := strconv.ParseInt(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := int8(number)
					column_schema.SetInt8("default", &converted)
				}
			} else if *dt == "float32" && default_value != "" {
				number, err := strconv.ParseFloat(default_value, 32)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := float32(number)
					column_schema.SetFloat32("default", &converted)
				}
			} else if *dt == "float64" && default_value != "" {
				number, err := strconv.ParseFloat(default_value, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					converted := float64(number)
					column_schema.SetFloat64("default", &converted)
				}
			}  else if *dt == "bool" && default_value != "" {
				number, err := strconv.ParseInt(default_value, 10, 64)
				if err != nil {
					errors = append(errors, err)
				} else {
					if number == 0 {
						boolean_value := false
						column_schema.SetBool("default", &boolean_value)
					} else if number == 1 {
						boolean_value := true
						column_schema.SetBool("default", &boolean_value)
					} else {
						errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 1 or 0", default_value, *dt))
					}
				}
			} else if *dt == "time.Time" && default_value != "" {
				if extra_value == "DEFAULT_GENERATED" && strings.HasPrefix(default_value, "CURRENT_TIMESTAMP") {
					if default_value == "CURRENT_TIMESTAMP" {
					} else if default_value == "CURRENT_TIMESTAMP(1)" {
					} else if default_value == "CURRENT_TIMESTAMP(2)" {
					} else if default_value == "CURRENT_TIMESTAMP(3)" {
					} else if default_value == "CURRENT_TIMESTAMP(4)" {
					} else if default_value == "CURRENT_TIMESTAMP(5)" {
					} else if default_value == "CURRENT_TIMESTAMP(6)" {
					} else {
						errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 0-6 decimal places", default_value, *dt))
					}
					
					default_value = "now"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.0" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.00" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.000" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.0000" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.00000" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				} else if default_value == "0000-00-00 00:00:00.000000" {
					default_value = "zero"
					column_schema.SetString("default", &default_value)
				}  else {
					errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be DEFAULT_GENERATED or 0000-00-00 00:00:00", default_value, *dt))
				}
			} else if !(*dt == "time.Time" || *dt == "bool" || *dt == "int64" || *dt == "uint64" ||  *dt == "int32" || *dt == "uint32" ||  *dt == "int16" || *dt == "uint16" ||  *dt == "int8" || *dt == "uint8" || *dt == "string" || *dt == "float32" || *dt == "float64") && default_value != "" {
				errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported please implement: %s for type: %s", default_value, *dt))
			}
		}
		

		if is_nullable {
			adjusted_type := "*" + *dt
			column_schema.SetString("type", &adjusted_type)
		}

		schema.SetMapValue(field_name, column_schema)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &schema, nil
}

