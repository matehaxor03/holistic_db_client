package class

import (
	"fmt"
	"strconv"
	"strings"
)

type Table struct {
	Validate              func() []error
	Exists                func() (*bool, []error)
	Create                func() []error
	Read                func() []error
	Delete                func() []error
	DeleteIfExists        func() []error
	GetSchema             func() (*Map, []error)
	GetTableName          func() (string, []error)
	SetTableName          func(table_name string) []error
	GetSchemaColumns       func() (*[]string, []error)
	GetTableColumns       func() (*[]string, []error)
	GetIdentityColumns    func() (*[]string, []error)
	GetNonIdentityColumns func() (*[]string, []error)
	Count                 func() (*uint64, []error)
	CreateRecord          func(record Map) (*Record, []error)
	ReadRecords         func(filter Map, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase           func() (*Database, []error)
	ToJSONString          func(json *strings.Builder) ([]error)
}

func newTable(database *Database, table_name string, schema Map, database_reserved_words_obj *DatabaseReservedWords, table_name_whitelist_characters_obj *TableNameCharacterWhitelist, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Table, []error) {
	struct_type := "*Table"

	SQLCommand := newSQLCommand()
	var errors []error
	var this_table *Table

	setTable := func(table *Table) {
		this_table = table
	}

	getTable := func() *Table {
		return this_table
	}

	database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	table_name_whitelist_characters := table_name_whitelist_characters_obj.GetTableNameCharacterWhitelist()
	column_name_whitelist_characters := column_name_whitelist_characters_obj.GetColumnNameCharacterWhitelist()
	
	
	setupData := func(b *Database, n string, s Map) (Map) {
		schema_is_nil := false
		
		if IsNil(s) {
			schema_is_nil = true
			s = Map{}
		}
	
		d := Map{
			"[fields]": Map{},
			"[schema]": Map{},
			"[system_fields]":Map{"[database]":b, "[table_name]":n},
			"[system_schema]": Map{
				"[database]": Map{"type":"*class.Database", "mandatory": true},
				"[table_name]": Map{"type": "*string", "mandatory": true, "not_empty_string_value": true, "min_length": 2,
				FILTERS(): Array{Map{"values": table_name_whitelist_characters, "function": getWhitelistCharactersFunc()},
								 Map{"values": database_reserved_words, "function": getBlacklistStringToUpperFunc()}}},
			},
		}
	
		if schema_is_nil {
			d["[schema_is_nil]"] = true
		} else {
			d["[schema_is_nil]"] = false
		}
	
		s["enabled"] = Map{"type": "*bool", "mandatory": false, "default": true}
		s["archieved"] = Map{"type": "*bool", "mandatory": false, "default": false}
		s["created_date"] = Map{"type": "*time.Time", "mandatory": false, "default": "now"}
		s["last_modified_date"] = Map{"type": "*time.Time", "mandatory": false, "default": "now"}
		s["archieved_date"] = Map{"type": "*time.Time", "mandatory": false, "default":nil}
	
		d["[schema]"] = s
		return d
	}

	data := setupData(database, table_name, schema)

	getData := func() (*Map) {
		return &data
	}

	getTableName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[table_name]", "string")
		return temp_value.(string), temp_value_errors
	}

	getTableColumns := func() (*[]string, []error) {
		temp_schemas, temp_schemas_error := GetSchemas(struct_type, getData(), "[schema]")
		if temp_schemas_error != nil {
			return nil, temp_schemas_error
		}
		columns := temp_schemas.Keys()
		return &columns, nil
	}

	getIdentityColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.Keys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("%s schema: %s is nill", struct_type, column))
				continue
			}

			if column_schema.IsBoolFalse("primary_key") {
				continue
			}

			columns = append(columns, column)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &columns, nil
	}

	getNonIdentityColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.Keys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("%s schema: %s is nill", struct_type, column))
				continue
			}

			if column_schema.IsBoolTrue("primary_key") {
				continue
			}

			columns = append(columns, column)
		}
		return &columns, nil
	}

	validate := func() []error {
		return ValidateData(getData(), "*class.Table")
	}

	getDatabase := func() (*Database, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database]", "*class.Database")
		return temp_value.(*Database), temp_value_errors
	}

	exists := func() (*bool, []error) {
		var errors []error
		validate_errors := validate()
		if errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}
		
		sql_command := fmt.Sprintf("SELECT 0 FROM %s LIMIT 1;", EscapeString(temp_table_name))
		
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}
		
		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false})

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

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return temp_table_name_errors
		}

		sql := fmt.Sprintf("DROP TABLE %s;", EscapeString(temp_table_name))

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, Map{"use_file": false})

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		errors := validate()
		if errors != nil {
			return errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return temp_table_name_errors
		}

		sql := fmt.Sprintf("DROP TABLE IF EXISTS %s;", EscapeString(temp_table_name))

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, Map{"use_file": false})

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getSchema := func() (*Map, []error) {
		var errors []error
		validate_errors := validate()
		
		if validate_errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}
		
		sql_command := fmt.Sprintf("SHOW COLUMNS FROM %s;", EscapeString(temp_table_name))

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false, "json_output": true})

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
			column_map := column_details.(Map)
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
					case "timestamp(6)", "timestamp(5)", "timestamp(4)", "timestamp(3)", "timestamp(2)", "timestamp(1)", "timestamp":
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
					errors = append(errors, fmt.Errorf("Table: %s GetSchema: column: %s attribute: %s not supported please implement", temp_table_name, field_name, column_attribute))
				}
			}

			if column_schema.IsNil("type") {
				errors = append(errors, fmt.Errorf("Table: %s GetSchema: column: %s attribute: type is nill", temp_table_name, field_name))
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
							errors = append(errors, fmt.Errorf("default value not supported %s for type: %s can only be 1 or 0", default_value, *dt))
						}
					}
				} else if *dt == "time.Time" && default_value != "" {
					if default_value == "NULL" {
						column_schema.SetNil("default")
					} else if (default_value == "CURRENT_TIMESTAMP(6)" || 
						default_value == "CURRENT_TIMESTAMP(3)" ||
						default_value == "CURRENT_TIMESTAMP") && extra_value == "DEFAULT_GENERATED" {
						now := "now"
						column_schema.SetString("default", &now)
					} else {
						errors = append(errors, fmt.Errorf("default value not supported %s for type: %s please implement", default_value, *dt))
					}
				} else if !(*dt == "time.Time" || *dt == "bool" || *dt == "int64" || *dt == "uint64" ||  *dt == "int32" || *dt == "uint32" ||  *dt == "int16" || *dt == "uint16" ||  *dt == "int8" || *dt == "uint8" || *dt == "string") && default_value != "" {
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
	}

	setTableName := func(new_table_name string) []error {
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_schema, temp_schema_errors := getSchema()
		if temp_schema_errors != nil {
			return temp_schema_errors
		}

		_, new_table_errors := newTable(temp_database, new_table_name, *temp_schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if new_table_errors != nil {
			return new_table_errors
		}

		temp_database_name_map, temp_database_name_map_errors := getData().GetMap("[table_name]")
		if temp_database_name_map_errors != nil {
			return temp_database_name_map_errors
		}

		temp_database_name_map.SetObject("value", new_table_name)
		return nil
	}

	getCreateSQL := func() (*string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}

		sql_command := fmt.Sprintf("CREATE TABLE %s", EscapeString(temp_table_name))

		valid_columns, valid_columns_errors := getTableColumns()
		if valid_columns_errors != nil {
			return nil, valid_columns_errors
		}

		schemas_map, schemas_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schemas_map_errors != nil {
			return nil, schemas_map_errors
		}

		primary_key_count := 0

		sql_command += "("
		for index, column := range *valid_columns {
			columnSchema, columnSchema_errors := schemas_map.GetMap(column)
			if columnSchema_errors != nil {
				errors = append(errors, columnSchema_errors...)
				continue
			} else if IsNil(columnSchema) {
				errors = append(errors, fmt.Errorf("%s column schema for column: %s is nil", struct_type, column))
				continue
			}

			sql_command += EscapeString(column)

			typeOf, type_of_errors := columnSchema.GetString("type")
			if type_of_errors != nil {
				errors = append(errors, type_of_errors...)
				continue
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
					errors = append(errors, fmt.Errorf("Table.getCreateSQL number type not mapped: %s", *typeOf))
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
							sql_command += " PRIMARY KEY"
							primary_key_count += 1
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
					}
				} 

				if columnSchema.HasKey("default") && columnSchema.GetType("default") == "int" {
					default_value, default_value_errors := columnSchema.GetInt64("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else {
						sql_command += " DEFAULT " + strconv.FormatInt(*default_value, 10)
					}
				}
			case "*time.Time", "time.Time":
				sql_command += " TIMESTAMP(6)"

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
						sql_command += " DEFAULT CURRENT_TIMESTAMP(6)"
					} else {
						errors = append(errors, fmt.Errorf("column: %s had default value it did not understand", column))
					}
				}
			case "*bool", "bool":
				sql_command += " BOOLEAN"

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

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
			errors = append(errors, fmt.Errorf("Table.getSQL: %s must have at least 1 primary key", EscapeString(temp_table_name)))
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

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, Map{"use_file": false})

		if execute_errors != nil {
			return execute_errors
		}

		return nil
	}

	read := func() []error {
		errors := validate()

		if len(errors) > 0 {
			return errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			errors = append(errors, temp_database_errors...)
		} else if IsNil(temp_database) {
			errors = append(errors, fmt.Errorf("Table.read database is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			errors = append(errors, temp_table_name_errors...)
		} else if IsNil(temp_table_name) {
			errors = append(errors, fmt.Errorf("Table.read table_name is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		temp_schema, temp_schema_errors := getSchema()
		if temp_schema_errors != nil {
			errors = append(errors, temp_schema_errors...)
		} else if IsNil(temp_schema) {
			errors = append(errors, fmt.Errorf("Table.read schema is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		data = setupData(temp_database, temp_table_name, *temp_schema)
		return nil
	}

	getSchemaColumns := func()  (*[]string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		schemas_map, schemas_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schemas_map_errors != nil {
			return nil, schemas_map_errors
		}

		schema_column_names := schemas_map.Keys()
		return &schema_column_names, nil
	}

	x := Table{
		Validate: func() []error {
			return validate()
		},
		GetDatabase: func() (*Database, []error) {
			return getDatabase()
		},
		GetTableColumns: func() (*[]string, []error) {
			return getTableColumns()
		},
		GetSchemaColumns: func() (*[]string, []error) {
			return getSchemaColumns()
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

			temp_table_name, temp_table_name_errors := getTableName()
			if temp_table_name_errors != nil {
				return nil, temp_table_name_errors
			}

			sql := fmt.Sprintf("SELECT COUNT(*) FROM %s;", EscapeString(temp_table_name))

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				return nil, temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, Map{"use_file": false})

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
		Read: func() []error {
			return read()
		},
		DeleteIfExists: func() []error {
			errors := validate()

			if len(errors) > 0 {
				return errors
			}

			return deleteIfExists()
		},
		CreateRecord: func(new_record_data Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := newRecord(getTable(), new_record_data, database_reserved_words_obj,  column_name_whitelist_characters_obj)
			if record_errors != nil {
				return nil, record_errors
			}

			create_record_errors := record.Create()
			if create_record_errors != nil {
				return nil, create_record_errors
			}

			return record, nil
		},
		ReadRecords: func(filters Map, limit *uint64, offset *uint64) (*[]Record, []error) {
			var errors []error
			validate_errors := validate()
			if errors != nil {
				errors = append(errors, validate_errors...)
				return nil, errors
			}

			table_schema := getData()

			temp_table_name, temp_table_name_errors := getTableName()
			if temp_table_name_errors != nil {
				return nil, temp_table_name_errors
			}

			if filters != nil {
				table_columns, table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				filter_columns := filters.Keys()
				for _, filter_column := range filter_columns {
					if !Contains(*table_columns, filter_column) {
						errors = append(errors, fmt.Errorf("Table.SelectRecords: column: %s not found for table: %s available columns are: %s", filter_column, temp_table_name, *table_columns))
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
						errors = append(errors, fmt.Errorf("Table.SelectRecords: column filter: %s for table: %s does not exist however filter had the value, table has columns: %s", filter_column, temp_table_name, table_schema.Keys()))
						continue
					}

					table_schema_column, table_schema_column_errors := table_schema.GetMap(filter_column)
					if table_schema_column_errors != nil {
						errors = append(errors, table_schema_column_errors...)
						continue
					}

					if table_schema_column.IsNil("type") {
						errors = append(errors, fmt.Errorf("Table.SelectRecords: column filter: %s for table: %s did not have atrribute: type", filter_column, temp_table_name))
						continue
					}


					table_column_type, _ := (*table_schema_column).GetString("type")
					if strings.Replace(*table_column_type, "*", "", -1) != strings.Replace(filter_column_type, "*", "", -1) {
						errors = append(errors, fmt.Errorf("Table.SelectRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, temp_table_name, *table_column_type))

						//todo ignore if filter data_type is nil and table column allows nil
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			sql := fmt.Sprintf("SELECT * FROM %s ", EscapeString(temp_table_name))
			if filters != nil {
				if len(filters.Keys()) > 0 {
					sql += "WHERE "
				}

				column_name_params := Map{"values": column_name_whitelist_characters, "value": nil, "label": "column_name", "data_type": "Table"}
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
						errors = append(errors, fmt.Errorf("Table.SelectRecords: filter type not supported please implement: %s", type_of))
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

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				return nil, temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, Map{"use_file": false})

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
					table_schema_column_map, table_schema_column_map_errors := table_schema.GetMap(column)
					if table_schema_column_map_errors != nil {
						errors = append(errors, table_schema_column_map_errors...)
						continue
					}
					
					table_data_type, table_data_type_errors := table_schema_column_map.GetString("type")
					if table_data_type_errors != nil {
						errors = append(errors, table_data_type_errors...)
						continue
					}

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
						errors = append(errors, fmt.Errorf("SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", temp_table_name, column, *table_data_type))
					}
				}

				mapped_record_obj, mapped_record_obj_errors := newRecord(getTable(), mapped_record, database_reserved_words_obj, column_name_whitelist_characters_obj)
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
			return getSchema()
		},
		GetTableName: func() (string, []error) {
			return getTableName()
		},
		SetTableName: func(table_name string) []error {
			return setTableName(table_name)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
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
