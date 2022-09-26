package class

import (
	"fmt"
	"bytes"
	"os/exec"
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

type Database struct {
	GetSQL func(action string) (*string, []error)
}

func NewDatabase(host *Host, credentials *Credentials, database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, []error) {
	//x := Database{host: host, credentials: credentials, database_name: database_name, database_create_options: database_create_options, options: options}

	data := Map {
		"host":Map{"type":"*Host","value":host,"mandatory":true},
		"credentials":Map{"type":"*Credentials","value":credentials,"mandatory":true},
		"database_name":Map{"type":"*string","value":database_name,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetDatabasenameValidCharacters(),"function":ValidateCharacters }}},
		"database_create_options":Map{"type":"*DatabaseCreateOptions","value":database_create_options,"mandatory":false},
		"options":Map{"type":"map[string]map[string][][]string)","value":options,"mandatory":false},
	}

	getSQL := func(command string) (*string, []error) {
		errors := ValidateGenericSpecial(data.Clone(), "Database")

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

		mapHost := data.M("host")
		if mapHost == nil {
			errors = append(errors, fmt.Errorf("host field not found in data"))	
		}

		host := mapHost.GetObject("value").(*Host)
		if host == nil {
			errors = append(errors, fmt.Errorf("host field is nil in data"))	
		}

		host_command, host_command_errors := (*host).GetCLSCommand()
		if host_command_errors != nil {
			errors = append(errors, host_command_errors...)	
		}

		mapCredentials := data.M("credentials")
		if mapCredentials == nil {
			errors = append(errors, fmt.Errorf("credentials field not found in data"))	
		}

		credentials = mapCredentials.GetObject("value").(*Credentials)
		if credentials == nil {
			errors = append(errors, fmt.Errorf("credentials field is nil in data"))	
		}

		credentials_command, credentials_command_errors := (*credentials).GetCLSCommand()
		if credentials_command_errors != nil {
			errors = append(errors, credentials_command_errors...)	
		}

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", *host_command, *credentials_command) 
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
	
	x := Database{
		GetSQL: func(action string) (*string, []error) {
			return getSQL(action)
		},
    }

	return &x, nil
}

func (this *Database) Create() (*Database, *string, []error)  {
	this, result, errors := (*this).createDatabase()
	if errors != nil {
		return nil, result, errors
	}

	return this, result, nil
}

func (this *Database) createDatabase() (*Database, *string, []error) {
	var errors []error 
	crud_sql_command, crud_command_errors := (*this).GetSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())

	if crud_command_errors != nil {
		errors = append(errors, crud_command_errors...)	
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", *crud_sql_command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    command_err := cmd.Run()

    if command_err != nil {
		errors = append(errors, command_err)	
	}

	shell_ouput := ""

	if len(errors) > 0 {
		shell_ouput = stderr.String()
		return nil, &shell_ouput, errors
	}

	shell_ouput = stdout.String()
    return this, &shell_ouput, nil
}
