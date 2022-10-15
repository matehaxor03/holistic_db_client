package class

import (
	"fmt"
	"strconv"
)

func GET_DATABASE_DATA_DEFINITION_STATEMENTS() Array {
	return Array {GET_DATA_DEFINTION_STATEMENT_CREATE()}
}

func GET_DATABASE_LOGIC_OPTIONS_CREATE() ([][]string){
	return [][]string{GET_LOGIC_STATEMENT_IF_NOT_EXISTS()}
}

func GET_DATABASE_OPTIONS() (map[string]map[string][][]string) {
	var root = make(map[string]map[string][][]string)
	
	var logic_options = make(map[string][][]string)
	logic_options[GET_DATA_DEFINTION_STATEMENT_CREATE()] = GET_DATABASE_LOGIC_OPTIONS_CREATE()

	root[GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options

	return root
}

func GetDatabasenameValidCharacters() string {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.="
}

func CloneDatabase(database *Database) (*Database) {
	if database == nil {
		return nil
	}

	return database.Clone()
}

type Database struct {
	Validate func() ([]error)
	Clone func() (*Database)
	Create func() ([]error)
	Delete func() ([]error)
	Exists func() (*bool, []error)
	GetDatabaseName func() (string)
	SetClient func(client *Client) ([]error)
	GetClient func() (*Client)
	CreateTable func(schema Map, options map[string]map[string][][]string) (*Table, []error)
	GetTable func(table_name string) (*Table, []error)
	ToJSONString func() string
}

func NewDatabase(client *Client, database_name string, database_create_options *DatabaseCreateOptions) (*Database, []error) {
	var this_database *Database

	SQLCommand := NewSQLCommand()
	database_name_whitelist := GetDatabasenameValidCharacters()

	data := Map {
		"[client]":Map{"value":CloneClient(client),"mandatory":true},
		"[database_name]":Map{"value":CloneString(&database_name),"mandatory":true,
		FILTERS(): Array{ Map {"values":&database_name_whitelist,"function":getWhitelistCharactersFunc() }}},
		"[database_create_options]":Map{"value":database_create_options,"mandatory":false},
	}

	getData := func() (Map) {
		return data.Clone()
	}

	validate := func() ([]error) {
		return ValidateData(data, "Database")
	}

	getClient := func() (*Client) {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	setClient := func(client *Client) {
		data.M("[client]")["value"] = client
	}

	getDatabaseName := func() (string) {
		n := CloneString(data.M("[database_name]").S("value"))
		return *n
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions) {
		return data.M("[database_create_options]").GetObject("value").(*DatabaseCreateOptions)
	}

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	getCreateSQL := func() (*string, []error) {
		errors := validate()
		
		if len(errors) > 0 {
			return nil, errors
		}

		sql_command := fmt.Sprintf("CREATE DATABASE %s ", getDatabaseName())
	
		databaseCreateOptions := getDatabaseCreateOptions()
		if databaseCreateOptions != nil {
			database_create_options_command, database_create_options_command_errs := (*databaseCreateOptions).GetSQL()
			if database_create_options_command_errs != nil {
				errors = append(errors, database_create_options_command_errs...)
			} else {
				sql_command += *database_create_options_command
			}
		}
		sql_command += ";"

		if len(errors) > 0 {
			return nil, errors
		}

		return &sql_command, nil
	}

	createDatabase := func() ([]error) {
		sql_command, generate_sql_errors := getCreateSQL()
	
		if generate_sql_errors != nil {
			return generate_sql_errors
		}
	
		_, execute_sql_command_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": false})

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}
	
		return nil
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}
	
	x := Database{
		Validate: func() ([]error) {
			return validate()
		},
		Clone: func() (*Database) {
			clone_value, _ := NewDatabase(getClient(), getDatabaseName(), getDatabaseCreateOptions())
			return clone_value
		},
		Create: func() ([]error) {
			errors := createDatabase()
			if errors != nil {
				return errors
			}

			return nil
		},
		Delete: func() ([]error) {
			errors := validate()
		
			if len(errors) > 0 {
				return errors
			}

			sql_command := fmt.Sprintf("DROP DATABASE %s;", EscapeString(getDatabaseName()))
			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false})
			
			if execute_errors != nil {
				errors = append(errors, execute_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		Exists: func() (*bool, []error) {
			errors := validate()
		
			if len(errors) > 0 {
				return nil, errors
			}

			sql_command := fmt.Sprintf("USE %s;", EscapeString(getDatabaseName()))
			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false})
			
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
		},
		CreateTable: func(schema Map, options map[string]map[string][][]string) (*Table, []error) {
			table, new_table_errors := NewTable(getDatabase(), schema, options)
			
			if new_table_errors != nil {
				return nil, new_table_errors
			}

			create_table_errors := table.Create()
			if create_table_errors != nil {
				return nil, create_table_errors
			}

			return table, nil
		},
		GetTable: func(table_name string) (*Table, []error) {
			var errors []error
			database := getDatabase()
			if database == nil {
				errors = append(errors, fmt.Errorf("database is nil"))
			} else {
				database_errors := database.Validate()
				if database_errors != nil {
					errors = append(errors, database_errors...)
				}
			}

			if table_name == "" {
				errors = append(errors, fmt.Errorf("table_name is empty"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			data_type := "Table"
			params := Map{"values": GetTableNameValidCharacters(), "value": &table_name, "data_type": &data_type, "label": table_name}
			table_name_errors := WhitelistCharacters(params)
			if table_name_errors != nil {
				errors = append(errors, table_name_errors...)
				return nil, errors 
			}
			
			sql_command := fmt.Sprintf("SHOW COLUMNS FROM %s;", EscapeString(table_name))
			
			json_array, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false, "json_output": true})
		
			if len(errors) > 0 {
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
				extra_value := ""
				for _, column_attribute := range column_attributes {
					switch column_attribute {
					case "Key":
						key_value := *(column_map.S("Key"))
						switch key_value {
							case "PRI":
								primary_key := true
								is_nullable = false
								column_schema.SetBool("primary_key", &primary_key)
							case "":
							default:
								errors = append(errors, fmt.Errorf("Key not implemented please implement: %s", key_value))
						}
					case "Field": 
						field_name = (*column_map.S("Field"))
					case "Type":
						type_of_value := (*column_map.S("Type"))
						switch type_of_value {
							case "bigint unsigned":
								data_type := "uint64"
								unsigned := true
								column_schema.SetString("type", &data_type)
								column_schema.SetBool("unsigned", &unsigned)
							case "bigint":
								data_type := "int64"
								column_schema.SetString("type", &data_type)
							case "timestamp(6)":
								data_type := "time.Time"
								column_schema.SetString("type", &data_type)
							case "tinyint(1)":
								data_type := "bool"
								column_schema.SetString("type", &data_type)
							default:
							errors = append(errors, fmt.Errorf("type not implement please implement: %s", type_of_value))
						}
					case "Null":
						null_value := *(column_map.S("Null"))
						switch null_value {
							case "YES":
								is_nullable = true
							case "NO":
								is_nullable = false
							default:
								errors = append(errors, fmt.Errorf("Null value not supported please implement: %s", null_value))
						}
					case "Default":
						default_value = *(column_map.S("Default"))
					case "Extra":
						extra_value = *(column_map.S("Extra"))
						switch extra_value {
							case "auto_increment":
								auto_increment := true
								column_schema.SetBool("auto_increment", &auto_increment)
							case "DEFAULT_GENERATED":
							case "":
							default:
								errors = append(errors, fmt.Errorf("Extra value not supported please implement: %s", extra_value))
						}	
					default:
						errors = append(errors, fmt.Errorf("column attribute not supported please implement: %s", column_attribute))
					}
				}

				if len(errors) > 0 {
					continue
				}

				if default_value != "" {
					if default_value == "NULL" {
					} else if default_value == "CURRENT_TIMESTAMP(6)" && extra_value == "DEFAULT_GENERATED" {
						now := "now"
						column_schema.SetString("default", &now)
					} else {
						if (*(column_schema.S("type")) == "uint64") {
							number, err := strconv.ParseUint(default_value, 10, 64)
							if err != nil {
								errors = append(errors, err)
							} else {
								column_schema.SetUInt64("default", &number)
							}
						} else if (*(column_schema.S("type")) == "int64") {
							number, err := strconv.ParseInt(default_value, 10, 64)
							if err != nil {
								errors = append(errors, err)
							} else {
								column_schema.SetInt64("default", &number)
							}
						} else if (*(column_schema.S("type")) == "bool") {
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
									errors = append(errors, fmt.Errorf("default value not supported %s for type: %s can only be 1 or 0", default_value, *(column_schema.S("type"))))
								}
							}
						} else {
							errors = append(errors, fmt.Errorf("default value not supported please implement: %s for type: %s", default_value, *(column_schema.S("type"))))
						}
					}
				}

				if is_nullable {
					adjusted_type := "*" + *(column_schema.S("type"))
					column_schema.SetString("type", &adjusted_type)
				}

				schema[field_name] = column_schema
			}

			table_name_schema := Map{"type":"*string", "value": &table_name}
			schema["[table_name]"] = table_name_schema

			if len(errors) > 0 {
				return nil, errors
			}

			table, new_table_errors := NewTable(getDatabase(), schema, nil)
			
			if new_table_errors != nil {
				return nil, new_table_errors
			}

			return table, nil			
		},
		SetClient: func(client *Client) ([]error) {
			var errors []error
			if client == nil {
				errors = append(errors, fmt.Errorf("client is nil"))
				return errors
			}

			client_errors := client.Validate()
			if client_errors != nil {
				errors = append(errors, client_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			setClient(client)
			return nil
		},
		GetClient: func() (*Client) {
			return getClient()
		},
		GetDatabaseName: func() (string) {
			return getDatabaseName()
		},
		ToJSONString: func() string {
			return getData().ToJSONString()
		},
    }
	setDatabase(&x)

	return &x, nil
}
