package class

import (
	"fmt"
)

func CloneDatabase(database *Database) *Database {
	if database == nil {
		return nil
	}

	return database.Clone()
}

type Database struct {
	Validate        func() []error
	Clone           func() *Database
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (*bool, []error)
	TableExists     func(table_name string) (*bool, []error)
	GetDatabaseName func() string
	SetDatabaseName func(database_name string) []error
	SetClient       func(client *Client) []error
	GetClient       func() *Client
	CreateTable     func(table_name string, schema Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() (*[]Table, []error)
	GetTableNames   func() (*[]string, []error)
	ToJSONString    func() (*string, []error)
	UseDatabase     func() []error
}

func NewDatabase(client *Client, database_name string, database_create_options *DatabaseCreateOptions) (*Database, []error) {
	SQLCommand := NewSQLCommand()
	var this_database *Database

	data := Map{
		"[client]": Map{"value": CloneClient(client), "mandatory": true},
		"[database_name]": Map{"value": CloneString(&database_name), "mandatory": true, "min_length":2, "not_empty_string_value": true,
			FILTERS(): Array{Map{"values": GetDatabaseNameWhitelistCharacters(), "function": getWhitelistCharactersFunc()},
							 Map{"values": GetMySQLKeywordsAndReservedWordsInvalidWords(), "function": getBlacklistStringToUpperFunc()}}},
		"[database_create_options]": Map{"value": database_create_options, "mandatory": false},
	}

	getData := func() (*Map, []error) {
		return data.Clone()
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Database")
	}

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	getDatabaseCreateOptions := func() *DatabaseCreateOptions {
		return data.M("[database_create_options]").GetObject("value").(*DatabaseCreateOptions)
	}

	getClient := func() *Client {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	setClient := func(client *Client) {
		(*(data.M("[client]")))["value"] = client
	}

	getDatabaseName := func() string {
		database_name, _ := data.M("[database_name]").GetString("value")
		n := CloneString(database_name)
		return *n
	}

	setDatabaseName := func(new_database_name string) []error {
		_, new_database_errors := NewDatabase(getClient(), new_database_name, getDatabaseCreateOptions())
		if new_database_errors != nil {
			return new_database_errors
		}

		(data["[database_name]"].(Map))["value"] = CloneString(&new_database_name)
		return nil
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

	createDatabase := func() []error {
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

	exists := func() (*bool, []error) {
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
	}

	getTableNames := func() (*[]string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command :=  fmt.Sprintf("SHOW TABLES IN %s;", EscapeString(getDatabaseName()))
		records, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false})

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
		column_name := "Tables_in_" + EscapeString(getDatabaseName())
		for _, record := range *records {
			table_name, table_name_errors := record.(*Map).GetString(column_name)
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

		sql_command := fmt.Sprintf("DROP DATABASE %s;", EscapeString(getDatabaseName()))
		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false})

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

		table, table_errors := NewTable(getDatabase(), table_name, nil)
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

		get_table, get_table_errors := NewTable(getDatabase(), table_name, *table_schema)

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
		Clone: func() *Database {
			clone_value, _ := NewDatabase(getClient(), getDatabaseName(), getDatabaseCreateOptions())
			return clone_value
		},
		Create: func() []error {
			errors := createDatabase()
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

			exists, exists_errors := exists()
			if exists_errors != nil {
				return exists_errors
			}

			if !(*exists) {
				return nil
			}

			return delete()
		},
		TableExists: func(table_name string) (*bool, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			table, table_errors := NewTable(getDatabase(), table_name, nil)

			if table_errors != nil {
				return nil, table_errors
			}

			return table.Exists()
		},
		CreateTable: func(table_name string, schema Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := NewTable(getDatabase(), table_name, schema)

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
		GetClient: func() *Client {
			return getClient()
		},
		GetDatabaseName: func() string {
			return getDatabaseName()
		},
		SetDatabaseName: func(database_name string) []error {
			return setDatabaseName(database_name)
		},
		UseDatabase: func() []error {
			return getClient().UseDatabase(getDatabase())
		},
		ToJSONString: func() (*string, []error) {
			data_cloned, data_cloned_errors := getData()
			if data_cloned_errors != nil {
				return nil, data_cloned_errors
			}

			return data_cloned.ToJSONString()
		},
	}
	setDatabase(&x)

	return &x, nil
}
