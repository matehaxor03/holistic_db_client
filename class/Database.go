package class

import (
	"fmt"
	"strings"
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
	GetSQL func(action string) (*string, []error)
	Create func() ([]error)
	GetDatabaseName func() (*string)
	SetClient func(client *Client) ([]error)
	GetClient func() (*Client)
	CreateTable func(schema Map, options map[string]map[string][][]string) (*Table, []error)
	GetTable func(table_name string) (*Table, []error)
}

func NewDatabase(client *Client, database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, []error) {
	var this_database *Database

	SQLCommand := newSQLCommand()
	database_name_whitelist := GetDatabasenameValidCharacters()

	data := Map {
		"client":Map{"value":CloneClient(client),"mandatory":true},
		"database_name":Map{"value":CloneString(database_name),"mandatory":true,
		FILTERS(): Array{ Map {"values":&database_name_whitelist,"function":getWhitelistCharactersFunc() }}},
		"database_create_options":Map{"value":database_create_options,"mandatory":false},
		"options":Map{"value":options,"mandatory":false},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "Database")
	}

	getClient := func() (*Client) {
		return CloneClient(data.M("client").GetObject("value").(*Client))
	}

	setClient := func(client *Client) {
		data.M("client")["value"] = client
	}

	getDatabaseName := func() (*string) {
		return CloneString(data.M("database_name").S("value"))
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions) {
		return data.M("database_create_options").GetObject("value").(*DatabaseCreateOptions)
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("options").GetObject("value").(map[string]map[string][][]string)
	}

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	getSQL := func(command string) (*string, []error) {
		errors := validate()

		m := Map{}
		m.SetArray("values", GET_DATABASE_DATA_DEFINITION_STATEMENTS())
		m.SetString("value", &command)
		commandTemp := "command"
		m.SetString("label", &commandTemp)
		someValue :=  "dsfdf"
		m.SetString("data_type", &someValue)

		command_errs := WhiteListString(m)


		if command_errs != nil {
			errors = append(errors, command_errs...)	
		}

		database_errs := ValidateGenericSpecial(data, "Database")

		if database_errs != nil {
			errors = append(errors, database_errs...)	
		}

		logic_option, logic_options_errs := GetLogicCommand(command, GET_LOGIC_STATEMENT_FIELD_NAME(), GET_DATABASE_OPTIONS(), options, "Database")
		if logic_options_errs != nil {
			errors = append(errors, logic_options_errs...)	
		}
		
		if len(errors) > 0 {
			return nil, errors
		}

		sql_command := fmt.Sprintf("%s DATABASE ", command)
		
		if *logic_option != "" {
			sql_command += fmt.Sprintf("%s ", *logic_option)
		}
		
		sql_command += fmt.Sprintf("%s ", *database_name)

		mapDatabaseCreateOptions := data.M("database_create_options")
		if mapDatabaseCreateOptions == nil {
			errors = append(errors, fmt.Errorf("database_create_options field not found in data"))	
		}

		databaseCreateOptions := mapDatabaseCreateOptions.GetObject("value").(*DatabaseCreateOptions)
		if databaseCreateOptions == nil {
			errors = append(errors, fmt.Errorf("database_create_options field is nil in data"))	
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		database_create_options_command, database_create_options_command_errs := (*databaseCreateOptions).GetSQL()
		if database_create_options_command_errs != nil {
			errors = append(errors, database_create_options_command_errs...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command += *database_create_options_command

		sql_command += ";"

		return &sql_command, nil
	}

	createDatabase := func() ([]error) {
		var errors []error 
		sql_command, sql_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if sql_command_errors != nil {
			return sql_command_errors
		}
	
		_, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": true})
	
		if stderr != nil && *stderr != "" {
			if strings.Contains(*stderr, " database exists") {
				errors = append(errors, fmt.Errorf("create database failed most likely the database already exists"))
			} else {
				errors = append(errors, fmt.Errorf(*stderr))
			}
		}
	
		if len(errors) > 0 {
			return errors
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
		GetSQL: func(action string) (*string, []error) {
			return getSQL(action)
		},
		Clone: func() (*Database) {
			clone_value, _ := NewDatabase(getClient(), getDatabaseName(), getDatabaseCreateOptions(), getOptions())
			return clone_value
		},
		Create: func() ([]error) {
			errors := createDatabase()
			if errors != nil {
				return errors
			}

			return nil
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
			
			json_array, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql_command, Map{"use_file": false, "json_output": true})
	
			if stderr != nil && *stderr != "" {
				if strings.Contains(*stderr, " database exists") {
					errors = append(errors, fmt.Errorf("create database failed most likely the database already exists"))
				} else {
					errors = append(errors, fmt.Errorf(*stderr))
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}

			fmt.Println((*json_array).ToJSONString())
		
			return nil, nil
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
		GetDatabaseName: func() (*string) {
			return getDatabaseName()
		},
    }
	setDatabase(&x)

	return &x, nil
}
