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
	Create func() (*string, []error)
	GetDatabaseName func() (*string)
}

func NewDatabase(host *Host, credentials *Credentials, database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, []error) {
	bashCommand := newBashCommand()
	mapType := "map[string]map[string][][]string)"
	databaseCreateOptionsType := "*DatabaseCreateOptions"
	database_name_whitelist := GetDatabasenameValidCharacters()

	data := Map {
		"host":Map{"type":"*Host","value":CloneHost(host),"mandatory":true},
		"credentials":Map{"type":"*Credentials","value":CloneCredentials(credentials),"mandatory":true},
		"database_name":Map{"type":"*string","value":database_name,"mandatory":true,
		FILTERS(): Array{ Map {"values":&database_name_whitelist,"function":ValidateCharacters }}},
		"database_create_options":Map{"type":&databaseCreateOptionsType,"value":database_create_options,"mandatory":false},
		"options":Map{"type":&mapType,"value":options,"mandatory":false},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "Database")
	}

	getHost := func() (*Host) {
		return CloneHost(data.M("host").GetObject("value").(*Host))
	}

	getCredentials := func() (*Credentials) {
		return CloneCredentials(data.M("credentials").GetObject("value").(*Credentials))
	}

	getDatabaseName := func() (*string) {
		return CloneString(data.M("database_name").GetObject("value").(*string))
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions) {
		return data.M("database_create_options").GetObject("value").(*DatabaseCreateOptions)
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("options").GetObject("value").(map[string]map[string][][]string)
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

		command_errs := ContainsExactMatch(m)


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

		host_command, host_command_errors := (*host).GetCLSCommand()
		if host_command_errors != nil {
			errors = append(errors, host_command_errors...)	
		}

		host := getHost()
		credentials = getCredentials()

		credentials_command := "--defaults-extra-file=./holistic-db-config-" +  *(host.GetHostName()) + "-" + *(host.GetPortNumber()) + "-" + *(credentials.GetUsername()) + "-" + *(getDatabaseName()) + ".config"

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", credentials_command, *host_command) 
		sql_command += fmt.Sprintf(" -e \"%s DATABASE ", command)
		
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

		sql_command += ";\""
		fmt.Println(sql_command)
		return &sql_command, nil
	}

	createDatabase := func() (*string, []error) {
		var errors []error 
		crud_sql_command, crud_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if crud_command_errors != nil {
			return nil, crud_command_errors
		}
	
		stdout, stderr, errors := bashCommand.ExecuteUnsafeCommand(crud_sql_command)
	
		if *stderr != "" {
			if strings.Contains(*stderr, " database exists") {
				errors = append(errors, fmt.Errorf("create database failed most likely the database already exists"))
			} else {
				errors = append(errors, fmt.Errorf("unknown create user error"))
			}
		}
	
		if len(errors) > 0 {
			return stdout, errors
		}
	
		return stdout, nil
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
			clone_value, _ := NewDatabase(getHost(), getCredentials(), getDatabaseName(), getDatabaseCreateOptions(), getOptions())
			return clone_value
		},
		Create: func() (*string, []error) {
			result, errors := createDatabase()
			if errors != nil {
				return result, errors
			}

			return result, nil
		},
		GetDatabaseName: func() (*string) {
			return getDatabaseName()
		},
    }

	return &x, nil
}
