package class

import (
	"fmt"
)

type Database struct {
	Validate        func() []error
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (*bool, []error)
	TableExists     func(table_name string) (*bool, []error)
	GetDatabaseName func() (string, []error)
	SetDatabaseName func(database_name string) []error
	SetClient       func(client *Client) []error
	GetClient       func() (*Client, []error)
	CreateTable     func(table_name string, schema Map) (*Table, []error)
	GetTableInterface func(table_name string, schema Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() (*[]Table, []error)
	GetTableNames   func() (*[]string, []error)
	ToJSONString    func() (*string, []error)
	UseDatabase     func() []error
}

func newDatabase(client *Client, database_name string, database_create_options *DatabaseCreateOptions, database_reserved_words_obj *DatabaseReservedWords, database_name_whitelist_characters_obj *DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *TableNameCharacterWhitelist, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Database, []error) {
	SQLCommand := newSQLCommand()
	var this_database *Database

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	database_name_whitelist_characters := database_name_whitelist_characters_obj.GetDatabaseNameCharacterWhitelist()

	data := Map{"[fields]": Map{
		"client":client, "database_name":&database_name,"database_create_options":database_create_options},
				"[schema]": Map{
		"client":Map{"type":"*class.Client","mandatory": true,"validated":false},
		"database_name":Map{"type":"*string", "mandatory": true,"min_length":2,"not_empty_string_value":true,"validated":false,
			FILTERS(): Array{Map{"values":database_name_whitelist_characters,"function":getWhitelistCharactersFunc()},
							 Map{"values":database_reserved_words,"function":getBlacklistStringToUpperFunc()}}},
		"database_create_options":Map{"type":"*class.DatabaseCreateOptions","mandatory":false,"validated":false}},
	}

	getData := func() (*Map) {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "Database")
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions, []error) {
		return GetDatabaseCreateOptionsField(getData(), "database_create_options")
	}

	getClient := func() (*Client, []error) {
		return GetClientField(getData(), "client")
	}

	setClient := func(new_client *Client) []error {
		return SetField(getData(), "client", new_client)
	}

	getDatabaseName := func() (string, []error) {
		var errors []error
		temp_database_name, temp_database_name_errors := GetStringField(getData(), "database_name")
		if temp_database_name_errors != nil {
			return "", temp_database_name_errors
		} else if IsNil(temp_database_name) {
			errors = append(errors, fmt.Errorf("database_name is nil"))
			return "", errors
		}

		return *temp_database_name, nil
	}

	setDatabaseName := func(new_database_name string) []error {
		temp_client, temp_client_errors := getClient()
		
		if temp_client_errors != nil {
			return temp_client_errors
		}

		temp_database_create_options, temp_database_create_options_errors := getDatabaseCreateOptions()
		if temp_database_create_options_errors != nil {
			return temp_database_create_options_errors
		}

		_, new_database_errors := newDatabase(temp_client, new_database_name, temp_database_create_options, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if new_database_errors != nil {
			return new_database_errors
		}

		temp_fields_map, temp_fields_map_errors := GetFields(getData())
		if temp_fields_map_errors != nil {
			return temp_fields_map_errors
		}

		temp_fields_map.SetString("database_name", &new_database_name)
		return nil
	}

	getCreateSQL := func() (*string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		sql_command := fmt.Sprintf("CREATE DATABASE %s ", EscapeString(temp_database_name))

		databaseCreateOptions, databaseCreateOptions_errors := getDatabaseCreateOptions()
		if databaseCreateOptions_errors != nil {
			return nil, databaseCreateOptions_errors
		}

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

	create := func() []error {
		sql_command, generate_sql_errors := getCreateSQL()

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}
		_, execute_sql_command_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, Map{"use_file":false, "creating_database":true})

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}

		return nil
	}

	exists := func() (*bool, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		sql_command := fmt.Sprintf("USE %s;", EscapeString(temp_database_name))

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false, "checking_database_exists":true})

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

	getTableNames := func() (*[]string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		sql_command :=  fmt.Sprintf("SHOW TABLES IN %s;", EscapeString(temp_database_name))

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		records, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false})

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if records == nil {
			errors = append(errors, fmt.Errorf("Database: getTableNames returned nil records"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		var table_names []string
		column_name := "Tables_in_" + EscapeString(temp_database_name)
		for _, record := range *records {
			table_name, table_name_errors := record.(Map).GetString(column_name)
			if table_name_errors != nil {
				errors = append(errors, table_name_errors...)
				continue
			}
			table_names = append(table_names, *table_name)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &table_names, nil
	}

	delete := func() ([]error) {
		errors := validate()

		if len(errors) > 0 {
			return errors
		}


		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}

		sql_command := fmt.Sprintf("DROP DATABASE %s;", EscapeString(temp_database_name))

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false, "deleting_database":true})

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		errors := validate()

		if len(errors) > 0 {
			return errors
		}


		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}

		sql_command := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", EscapeString(temp_database_name))

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, Map{"use_file": false, "deleting_database":true})

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getTable := func(table_name string) (*Table, []error) {
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

		if len(errors) > 0 {
			return nil, errors
		}

		table, table_errors := newTable(getDatabase(), table_name, nil, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if table_errors != nil {
			errors = append(errors, table_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		table_schema, schema_errors := table.GetSchema()

		if schema_errors != nil {
			errors = append(errors, schema_errors...)
		}

		if table_schema == nil {
			errors = append(errors, fmt.Errorf("Database.getTable schema is nil for table: %s", table_name))
		}

		if len(errors) > 0 {
			return nil, errors
		}		

		get_table, get_table_errors := newTable(getDatabase(), table_name, table_schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

		if get_table_errors != nil {
			return nil, get_table_errors
		}

		return get_table, nil
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	x := Database{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			errors := create()
			if errors != nil {
				return errors
			}

			return nil
		},
		Delete: func() []error {
			return delete()
		},
		Exists: func() (*bool, []error) {
			return exists()
		},
		DeleteIfExists: func() []error {
			errors := validate()

			if len(errors) > 0 {
				return errors
			}

			return deleteIfExists()
		},
		TableExists: func(table_name string) (*bool, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			table, table_errors := newTable(getDatabase(), table_name, nil, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if table_errors != nil {
				return nil, table_errors
			}

			return table.Exists()
		},
		GetTableInterface: func(table_name string, schema Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := newTable(getDatabase(), table_name, &schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if new_table_errors != nil {
				return nil, new_table_errors
			}

			return table, nil
		},
		CreateTable: func(table_name string, schema Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := newTable(getDatabase(), table_name, &schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

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
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			return getTable(table_name)
		},
		GetTableNames: func()  (*[]string, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			return getTableNames()
		},
		GetTables: func() (*[]Table, []error) {
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

			if len(errors) > 0 {
				return nil, errors
			}

			var tables []Table
			table_names, table_names_errors := getTableNames()
			if table_names_errors != nil {
				errors = append(errors, table_names_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			for _, table_name := range *table_names {
				table, table_errors := getTable(table_name)
				if table_errors != nil {
					errors = append(errors, table_errors...)
					continue
				}

				tables = append(tables, *table)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &tables, nil
		},
		SetClient: func(client *Client) []error {
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
		GetClient: func() (*Client, []error) {
			return getClient()
		},
		GetDatabaseName: func() (string, []error) {
			return getDatabaseName()
		},
		SetDatabaseName: func(database_name string) []error {
			return setDatabaseName(database_name)
		},
		UseDatabase: func() []error {
			temp_database := getDatabase()

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			return temp_client.UseDatabase(temp_database)
		},
		ToJSONString: func() (*string, []error) {
			return getData().ToJSONString()
		},
	}
	setDatabase(&x)

	return &x, nil
}
