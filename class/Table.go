package class

import (
	"fmt"
	"strconv"
	"strings"
)

func CloneTable(table *Table) *Table {
	if table == nil {
		return nil
	}

	return table.Clone()
}

type Table struct {
	Validate              func() []error
	Clone                 func() *Table
	Exists                func() (*bool, []error)
	Create                func() []error
	Delete                func() []error
	DeleteIfExists        func() []error
	GetSchema             func() (*Map, []error)
	GetTableName          func() string
	SetTableName          func(table_name string) []error
	GetTableColumns       func() (*[]string, []error)
	GetIdentityColumns    func() (*[]string, []error)
	GetNonIdentityColumns func() (*[]string, []error)
	Count                 func() (*uint64, []error)
	GetData               func() (*Map, []error)
	CreateRecord          func(record Map) (*Record, []error)
	Select                func(filter Map, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase           func() *Database
	ToJSONString          func() (*string, []error)
}

func NewTable(database *Database, table_name string, schema Map) (*Table, []error) {
	var this_table *Table
	SQLCommand := NewSQLCommand()
	var errors []error

	if schema == nil {
		schema = Map{}
		schema["[schema_is_nil]"] = Map{"type": "*bool", "value": true, "mandatory": true}
	} else {
		schema["[schema_is_nil]"] = Map{"type": "*bool", "value": false, "mandatory": true}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	data, data_errors := schema.Clone()
	if data_errors != nil {
		return nil, data_errors
	}
	
	(*data)["[database]"] = Map{"value": CloneDatabase(database), "mandatory": true}
	(*data)["[table_name]"] = Map{"type": "*string", "value": &table_name, "mandatory": true, "not_empty_string_value": true, "min_length": 2, 
		FILTERS(): Array{Map{"values": GetTableNameValidCharacters(), "function": getWhitelistCharactersFunc()},
						 Map{"values": GetMySQLKeywordsAndReservedWordsInvalidWords(), "function": getBlacklistStringToUpperFunc()}}}

	(*data)["active"] = Map{"type": "*bool", "mandatory": true, "default": true}
	(*data)["archieved"] = Map{"type": "*bool", "mandatory": true, "default": false}
	(*data)["created_date"] = Map{"type": "*time.Time", "mandatory": true, "default": "now"}
	(*data)["last_modified_date"] = Map{"type": "*time.Time", "mandatory": true, "default": "now"}
	(*data)["archieved_date"] = Map{"type": "*time.Time", "mandatory": true, "default": "now"}

	getData := func() (*Map, []error) {
		return data.Clone()
	}

	getTableName := func() string {
		table_name, _ := data.M("[table_name]").GetString("value")
		n := CloneString(table_name)
		return *n
	}

	getTableColumns := func() (*[]string, []error) {
		var columns []string
		column_name_whitelist_params := Map{"values": GetColumnNameValidCharacters(), "value": nil, "label": "column_name_character", "data_type": "Column"}
		column_name_blacklist_params := Map{"values": GetMySQLKeywordsAndReservedWordsInvalidWords(), "value": nil, "label": "column_name", "data_type": "Column"}

		clone_data, clone_data_errors := getData()
		if clone_data_errors != nil {
			return nil, clone_data_errors
		}

		for _, column := range clone_data.Keys() {
			if data.GetType(column) != "class.Map" {
				continue
			}

			column_name_whitelist_params.SetString("value", &column)
			column_name_whiltelist_errors := WhitelistCharacters(column_name_whitelist_params)
			if column_name_whiltelist_errors != nil {
				continue
			}
			
			column_name_blacklist_params.SetString("value", &column)
			column_name_blacklist_errors := BlackListStringToUpper(column_name_blacklist_params)
			if column_name_blacklist_errors != nil {
				continue
			}	

			columns = append(columns, column)
		}
		return &columns, nil
	}

	getIdentityColumns := func() (*[]string, []error) {
		var columns []string
		table_columns, table_columns_errors := getTableColumns()
		if table_columns_errors != nil {
			return nil, table_columns_errors
		}

		for _, column := range *table_columns {
			columnSchema := (*data)[column].(Map)

			if columnSchema.IsBoolFalse("primary_key") {
				continue
			}

			columns = append(columns, column)
		}
		return &columns, nil
	}

	getNonIdentityColumns := func() (*[]string, []error) {
		var columns []string
		
		table_columns, table_columns_errors := getTableColumns()
		if table_columns_errors != nil {
			return nil, table_columns_errors
		}

		for _, column := range *table_columns {
			columnSchema := (*data)[column].(Map)

			if columnSchema.IsBoolTrue("primary_key") {
				continue
			}

			columns = append(columns, column)
		}
		return &columns, nil
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := getData()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "*class.Table")
	}

	getDatabase := func() *Database {
		return CloneDatabase(data.M("[database]").GetObject("value").(*Database))
	}

	setTable := func(table *Table) {
		this_table = table
	}

	getTable := func() *Table {
		return this_table
	}

	setTableName := func(new_table_name string) []error {
		_, new_table_errors := NewTable(getDatabase(), new_table_name, schema)
		if new_table_errors != nil {
			return new_table_errors
		}

		((*data)["[table_name]"].(Map))["value"] = CloneString(&new_table_name)
		return nil
	}

	exists := func() (*bool, []error) {
		var errors []error
		validate_errors := validate()
		if errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}

		sql_command := fmt.Sprintf("SELECT 0 FROM %s LIMIT 1;", EscapeString(getTableName()))
		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql_command, Map{"use_file": false})

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		boolean_value := false
		if len(errors) > 0 {
			//todo: check error message e.g database does not exist
			boolean_value = false
			return &boolean_value, nil
		}

		boolean_value = true
		return &boolean_value, nil
	}

	delete := func() ([]error) {
		errors := validate()
		if errors != nil {
			return errors
		}

		sql := fmt.Sprintf("DROP TABLE %s;", EscapeString(getTableName()))
		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false})

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getCreateSQL := func() (*string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command := fmt.Sprintf("CREATE TABLE %s", EscapeString(getTableName()))

		valid_columns, valid_columns_errors := getTableColumns()
		if valid_columns_errors != nil {
			return nil, valid_columns_errors
		}

		primary_key_count := 0

		sql_command += "("
		for index, column := range *valid_columns {
			columnSchema := (*data)[column].(Map)
			sql_command += EscapeString(column)

			typeOf, _ := columnSchema.GetString("type")
			switch *typeOf {
			case "*uint64", "*int64", "*int", "uint64", "uint", "int64", "int":
				sql_command += " BIGINT"

				if *typeOf == "*uint64" ||
					*typeOf == "*uint" ||
					*typeOf == "uint64" ||
					*typeOf == "uint" {
					sql_command += " UNSIGNED"
				}

				if columnSchema.HasKey("auto_increment") {
					if columnSchema.IsBool("auto_increment") && !columnSchema.IsNil("auto_increment") {
						if columnSchema.IsBoolTrue("auto_increment") {
							sql_command += " AUTO_INCREMENT"
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.GetType("auto_increment")))
					}
				}

				if columnSchema.HasKey("primary_key") {
					if columnSchema.IsBool("primary_key") && !columnSchema.IsNil("primary_key") {
						if columnSchema.IsBoolTrue("primary_key") {
							sql_command += " PRIMARY KEY NOT NULL"
							primary_key_count += 1
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
					}
				} else if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") && columnSchema.GetType("default") == "int" {
					default_value, default_value_errors := columnSchema.GetInt64("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else {
						sql_command += " DEFAULT " + strconv.FormatInt(*default_value, 10)
					}
				}
			case "*time.Time":
				sql_command += " TIMESTAMP(6)"
				if columnSchema.HasKey("default") {
					default_value, _ := columnSchema.GetString("default")
					if columnSchema.IsNil("default") {
						errors = append(errors, fmt.Errorf("column: %s had nil default value", column))
						continue
					} else if *default_value != "now" {
						errors = append(errors, fmt.Errorf("column: %s had default value it did not understand", column))
						continue
					}

					sql_command += " DEFAULT CURRENT_TIMESTAMP(6)"
				}
			case "*bool", "bool":
				sql_command += " BOOLEAN"
				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						errors = append(errors, fmt.Errorf("column: %s had nil default value", column))
					} else if !columnSchema.IsBool("default") {
						errors = append(errors, fmt.Errorf("column: %s had non-boolean default value", column))
					} else if columnSchema.IsBoolTrue("default") {
						sql_command += " DEFAULT 1"
					} else if columnSchema.IsBoolFalse("default") {
						sql_command += " DEFAULT 0"
					} else {
						errors = append(errors, fmt.Errorf("column: %s had unknown error for boolean default value", column))
					}
				}
			case "*string", "string":
				sql_command += " VARCHAR("
				if !columnSchema.HasKey("max_length") {
					errors = append(errors, fmt.Errorf("column: %s did not specify length attribute", column))
				} else if columnSchema.GetType("max_length") != "int" {
					errors = append(errors, fmt.Errorf("column: %s specified length attribute however it's not an int", column))
				} else {
					max_length, max_length_errors := columnSchema.GetInt("max_length")
					if max_length_errors != nil {
						errors = append(errors, fmt.Errorf("column: %s specified max_length attribute had errors %s", column, fmt.Sprintf("%s", max_length_errors)))
					} else if *max_length <= 0 {
						errors = append(errors, fmt.Errorf("column: %s specified length attribute was <= 0 and had value: %d", column, max_length))
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
					if columnSchema.IsNil("default") {
						sql_command += " DEFAULT NULL"
					} else if !columnSchema.IsString("default") {
						errors = append(errors, fmt.Errorf("column: %s had non-string default value", column))
					} else {
						default_value, default_value_errors := columnSchema.GetString("default")
						if default_value_errors != nil {
							errors = append(errors, fmt.Errorf("column: %s specified default attribute had errors %s", column, fmt.Sprintf("%s", default_value_errors)))
						} else {
							sql_command += " DEFAULT \"" + EscapeString(*default_value) + "\""
						}
					} 
				}

				
			default:
				errors = append(errors, fmt.Errorf("Table.getSQL type: %s is not supported please implement for column %s", *typeOf, column))
			}

			if index < (len(*valid_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ");"

		if primary_key_count == 0 {
			errors = append(errors, fmt.Errorf("table: %s must have at least 1 primary key", EscapeString(getTableName())))
		}

		// todo: check that length of row for all columns does not exceed 65,535 bytes (it's not hard but low priority)

		if len(errors) > 0 {
			return nil, errors
		}

		return &sql_command, nil
	}

	createTable := func() []error {
		sql_command, sql_command_errors := getCreateSQL()

		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), sql_command, Map{"use_file": false})

		if execute_errors != nil {
			return execute_errors
		}

		return nil
	}

	x := Table{
		Validate: func() []error {
			return validate()
		},
		GetDatabase: func() *Database {
			return getDatabase()
		},
		Clone: func() *Table {
			schema_clone, _ :=  schema.Clone()
			clone_value, _ := NewTable(getDatabase(), getTableName(), *schema_clone)
			return clone_value
		},
		GetTableColumns: func() (*[]string, []error) {
			return getTableColumns()
		},
		GetIdentityColumns: func() (*[]string, []error) {
			return getIdentityColumns()
		},
		GetNonIdentityColumns: func() (*[]string, []error) {
			return getNonIdentityColumns()
		},
		Create: func() []error {
			errors := createTable()
			if errors != nil {
				return errors
			}

			return nil
		},
		Count: func() (*uint64, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			sql := fmt.Sprintf("SELECT COUNT(*) FROM %s;", EscapeString((getTableName())))
			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false})

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			if len(*json_array) != 1 {
				errors = append(errors, fmt.Errorf("count record does not exist"))
				return nil, errors
			}

			count_value, _ := (*json_array)[0].(*Map).GetString("COUNT(*)")
			count, count_err := strconv.ParseUint(*count_value, 10, 64)
			if count_err != nil {
				errors = append(errors, count_err)
				return nil, errors
			}

			return &count, nil
		},
		Delete: func() []error {
			return delete()
		},
		DeleteIfExists: func() []error {
			errors := validate()

			if len(errors) > 0 {
				return errors
			}

			exists, exists_errors := exists()
			if exists_errors != nil {
				return exists_errors
			}

			if !(*exists) {
				return nil
			}

			return delete()
		},
		CreateRecord: func(new_record_data Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := NewRecord(getTable(), new_record_data)
			if record_errors != nil {
				return nil, record_errors
			}

			create_record_errors := record.Create()
			if create_record_errors != nil {
				return nil, create_record_errors
			}

			return record, nil
		},
		Select: func(filters Map, limit *uint64, offset *uint64) (*[]Record, []error) {
			var errors []error
			validate_errors := validate()
			if errors != nil {
				errors = append(errors, validate_errors...)
				return nil, errors
			}

			table_schema, table_schema_errors := getData()
			if table_schema_errors != nil {
				return nil, table_schema_errors
			}

			if filters != nil {
				table_columns,table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				filter_columns := filters.Keys()
				for _, filter_column := range filter_columns {
					if !Contains(*table_columns, filter_column) {
						errors = append(errors, fmt.Errorf("SelectRecords: column: %s not found for table: %s available columns are: %s", filter_column, getTableName(), *table_columns))
					}
				}

				if len(errors) > 0 {
					return nil, errors
				}

				for _, filter_column := range filter_columns {
					filter_column_type := filters.GetType(filter_column)

					if !filters.IsNil(filter_column) && !strings.HasPrefix(filter_column_type, "*") {
						filter_column_type = "*" + filter_column_type
					}
					 
					if table_schema.IsNil(filter_column) {
						errors = append(errors, fmt.Errorf("SelectRecords: column filter: %s for table: %s does not exist however filter had the value, table has columns: %s", filter_column, getTableName(), table_schema.Keys()))
						continue
					}

					table_schema_column := table_schema.M(filter_column)

					if table_schema_column.IsNil("type") {
						errors = append(errors, fmt.Errorf("SelectRecords: column filter: %s for table: %s did not have atrribute: type", filter_column, getTableName()))
						continue
					}


					table_column_type, _ := (*table_schema_column).GetString("type")
					if strings.Replace(*table_column_type, "*", "", -1) != strings.Replace(filter_column_type, "*", "", -1) {
						errors = append(errors, fmt.Errorf("SelectRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, getTableName(), *table_column_type))

						//todo ignore if filter data_type is nil and table column allows nil
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			sql := fmt.Sprintf("SELECT * FROM %s ", EscapeString(getTableName()))
			if filters != nil {
				if len(filters.Keys()) > 0 {
					sql += "WHERE "
				}

				column_name_params := Map{"values": GetColumnNameValidCharacters(), "value": nil, "label": "column_name", "data_type": "Table"}
				for index, column_filter := range filters.Keys() {
					column_name_params.SetString("value", &column_filter)
					column_name_errors := WhitelistCharacters(column_name_params)
					if column_name_errors != nil {
						errors = append(errors, column_name_errors...)
					}	

					sql += EscapeString(column_filter) + " = "

					//todo check data type with schema
					type_of := filters.GetType(column_filter)
					switch type_of {


					case "*string", "string":
						filer_value, _ := filters.GetString(column_filter)
						sql += fmt.Sprintf("'%s' ", EscapeString(*filer_value))
					default:
						errors = append(errors, fmt.Errorf("Table: Select: filter type not supported please implement: %s", type_of))
					}

					if index < len(filters.Keys()) - 1 {
						sql += "AND "
					}
				}
			}

			if limit != nil {
				limit_value := strconv.FormatUint(*limit, 10)
				sql += fmt.Sprintf("LIMIT %s ", limit_value)
			}

			if offset != nil {
				offset_value := strconv.FormatUint(*offset, 10)
				sql += fmt.Sprintf("OFFSET %s ", offset_value)
			}
			sql += ";"

			if len(errors) > 0 {
				return nil, errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false})

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, json := range *json_array {
				current_record := json.(*Map)
				columns := current_record.Keys()
				mapped_record := Map{}
				for _, column := range columns {
					table_data_type, _ := table_schema.M(column).GetString("type")
					switch *table_data_type {
					case "*uint64", "uint64":
						value, value_errors := current_record.GetUInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt64(column, value)
						}
					case "*int64", "int64":
						value, value_errors := current_record.GetInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt64(column, value)
						}
					case "*int", "int":
						value, value_errors := current_record.GetInt(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt(column, value)
						}
					case "*time.Time":
						value, value_errors := current_record.GetTime(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetTime(column, value)
						}
					case "*bool", "bool":
						value, value_errors := current_record.GetBool(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetBool(column, value)
						}
					case "*string", "string":
						value, value_errors := current_record.GetString(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetString(column, value)
						}
					default:
						errors = append(errors, fmt.Errorf("SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", getTableName(), column, *table_data_type))
					}
				}

				mapped_record_obj, mapped_record_obj_errors := NewRecord(getTable(), mapped_record)
				if mapped_record_obj_errors != nil {
					errors = append(errors, mapped_record_obj_errors...)
				} else {
					mapped_records = append(mapped_records, *mapped_record_obj)
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &mapped_records, nil
		},
		Exists: func() (*bool, []error) {
			return exists()
		},
		GetSchema: func() (*Map, []error) {
			var errors []error
			validate_errors := validate()
			
			if validate_errors != nil {
				errors = append(errors, validate_errors...)
				return nil, errors
			}
			
			sql_command := fmt.Sprintf("SHOW COLUMNS FROM %s;", EscapeString(getTableName()))

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql_command, Map{"use_file": false, "json_output": true})

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
				return nil, errors
			}

			if json_array == nil {
				errors = append(errors, fmt.Errorf("show columns returned nil records"))
				return nil, errors
			}

			if len(*json_array) == 0 {
				errors = append(errors, fmt.Errorf("show columns did not return any records"))
				return nil, errors
			}

			schema := Map{}
			for _, column_details := range *json_array {
				column_map := column_details.(*Map)
				column_attributes := column_map.Keys()

				column_schema := Map{}
				default_value := ""
				field_name := ""
				is_nullable := true
				is_primary_key := false
				is_mandatory := false
				extra_value := ""
				for _, column_attribute := range column_attributes {
					switch column_attribute {
					case "Key":
						key_value, _ := column_map.GetString("Key")
						switch *key_value {
						case "PRI":
							is_primary_key = true
							is_mandatory = true
							is_nullable = false
							column_schema.SetBool("primary_key", &is_primary_key)
							column_schema.SetBool("mandatory", &is_mandatory)
						case "":
						default:
							errors = append(errors, fmt.Errorf("Table: GetSchema: Key not implemented please implement: %s", *key_value))
						}
					case "Field":
						field_name_value, _ := column_map.GetString("Field")
						field_name = *field_name_value
					case "Type":
						type_of_value, _ := column_map.GetString("Type")
						switch *type_of_value {
						case "bigint unsigned", "int unsigned", "smallint unsigned":
							data_type := "uint64"
							unsigned := true
							column_schema.SetString("type", &data_type)
							column_schema.SetBool("unsigned", &unsigned)
						case "bigint", "int", "smallint":
							data_type := "int64"
							column_schema.SetString("type", &data_type)
						case "timestamp(6)", "timestamp":
							data_type := "time.Time"
							column_schema.SetString("type", &data_type)
						case "tinyint(1)":
							data_type := "bool"
							column_schema.SetString("type", &data_type)
						case "text", "blob", "json":
							data_type := "string"
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
									errors = append(errors, fmt.Errorf("Table: GetSchema: could not determine parts of enum had length of zero: %s", *type_of_value))
								} else {
									part := parts[0]
									if strings.HasPrefix(part, "'")  && strings.HasSuffix(part, "'") {
										data_type := "string"
										column_schema.SetString("type", &data_type)
									} else {
										errors = append(errors, fmt.Errorf("Table: GetSchema: could not determine parts of enum for data type: %s", *type_of_value))
									}
								}
							} else {
								errors = append(errors, fmt.Errorf("Table: GetSchema: type not implemented please implement: %s", *type_of_value))
							}
						}
					case "Null":
						null_value, _ := column_map.GetString("Null")
						switch *null_value {
						case "YES":
							if !is_primary_key {
								is_mandatory = false
								is_nullable = true
								column_schema.SetBool("mandatory", &is_mandatory)
							}
						case "NO":
							is_nullable = false
							is_mandatory = true
							column_schema.SetBool("mandatory", &is_mandatory)
						default:
							errors = append(errors, fmt.Errorf("Table: GetSchema: Null value not supported please implement: %s", *null_value))
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
							errors = append(errors, fmt.Errorf("Table: GetSchema: Extra value not supported please implement: %s", extra_value))
						}
					default:
						errors = append(errors, fmt.Errorf("Table: %s GetSchema: column: %s attribute: %s not supported please implement", getTableName(), field_name, column_attribute))
					}
				}

				if column_schema.IsNil("type") {
					errors = append(errors, fmt.Errorf("Table: %s GetSchema: column: %s attribute: type is nill", getTableName(), field_name))
				}

				if len(errors) > 0 {
					continue
				}

				dt, _ := column_schema.GetString("type")

			
				if default_value == "NULL" {
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
					} else if *dt == "bool" && default_value != "" {
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
								errors = append(errors, fmt.Errorf("default value not supported %s for type: %s can only be 1 or 0", default_value, *dt))
							}
						}
					} else if *dt == "time.Time" && default_value != "" {
						if (default_value == "CURRENT_TIMESTAMP(6)" || 
							default_value == "CURRENT_TIMESTAMP(3)" ||
							default_value == "CURRENT_TIMESTAMP") && extra_value == "DEFAULT_GENERATED" {
							now := "now"
							column_schema.SetString("default", &now)
						} else {
							errors = append(errors, fmt.Errorf("default value not supported %s for type: %s please implement", default_value, *dt))
						}
					} else if !(*dt == "time.Time" || *dt == "bool" || *dt == "int64" || *dt == "string" || *dt == "uint64") && default_value != "" {
						errors = append(errors, fmt.Errorf("default value not supported please implement: %s for type: %s", default_value, *dt))
					}
				}
				

				if is_nullable {
					adjusted_type := "*" + *dt
					column_schema.SetString("type", &adjusted_type)
				}

				schema[field_name] = column_schema
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &schema, nil
		},
		GetData: func() (*Map, []error) {
			return getData()
		},
		GetTableName: func() string {
			return getTableName()
		},
		SetTableName: func(table_name string) []error {
			return setTableName(table_name)
		},
		ToJSONString: func() (*string, []error) {
			data_clone, data_clone_errors := data.Clone()
			if data_clone_errors != nil {
				return nil, data_clone_errors
			}
			return data_clone.ToJSONString()
		},
	}
	setTable(&x)

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
