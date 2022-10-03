package class

import (
	"fmt"
	"strings"
)

func GET_USER_DATA_DEFINITION_STATEMENTS() Array {
	return Array{GET_DATA_DEFINTION_STATEMENT_CREATE()}
}

func GET_USER_LOGIC_OPTIONS_CREATE() ([][]string){
	return [][]string{GET_LOGIC_STATEMENT_IF_NOT_EXISTS()}
}

func GET_USER_EXTRA_OPTIONS() (map[string]map[string][][]string) {
	var root = make(map[string]map[string][][]string)
	
	var logic_options = make(map[string][][]string)
	logic_options[GET_DATA_DEFINTION_STATEMENT_CREATE()] = GET_USER_LOGIC_OPTIONS_CREATE()

	root[GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options

	return root
}

type User struct {
	Create func() (*string, []error)
}

func NewUser(host *Host, host_credentials *Credentials, database *Database, credentials *Credentials, domain_name *DomainName, options map[string]map[string][][]string) (*User, []error) {
	SQLCommand := newSQLCommand()
	
	data := Map {
		"host":Map{"type":"*Host","value":CloneHost(host),"mandatory":true},
		"host_credentials":Map{"type":"*Credentials","value":CloneCredentials(host_credentials),"mandatory":true},
		"database":Map{"type":"*Database","value":CloneDatabase(database),"mandatory":true},
		"credentials":Map{"type":"*Credentials","value":CloneCredentials(credentials),"mandatory":true},
		"domain_name":Map{"type":"*DomainName","value":CloneDomainName(domain_name),"mandatory":true},
		"options":Map{"type":"map[string]map[string][][]string)","value":options,"mandatory":false},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "User")
	}


	getHost := func() (*Host) {
		return CloneHost(data.M("host").GetObject("value").(*Host))
	}

	getHostCredentials := func() (*Credentials) {
		return CloneCredentials(data.M("host_credentials").GetObject("value").(*Credentials))
	}

	getDatabase := func() (*Database) {
		return CloneDatabase(data.M("database").GetObject("value").(*Database))
	}

	getSQL := func(action string) (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}

		m := Map{}
		m.SetArray("values", GET_USER_DATA_DEFINITION_STATEMENTS())
		m.SetString("value", &action)
		commandTemp := "command"
		m.SetString("label", &commandTemp)
		dataTypeTemp := "User"
		m.SetString("data_type", &dataTypeTemp)


		command_errs := ContainsExactMatch(m)

		if command_errs != nil {
			errors = append(errors, command_errs...)	
		}

		database_errs := validate()

		if database_errs != nil {
			errors = append(errors, database_errs...)	
		}

		logic_option, logic_option_errs := GetLogicCommand(action, GET_LOGIC_STATEMENT_FIELD_NAME(), GET_USER_EXTRA_OPTIONS(), data.M("options").GetObject("value").(map[string]map[string][][]string), "User")
		if logic_option_errs != nil {
			errors = append(errors, logic_option_errs...)	
		}

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command := fmt.Sprintf("%s USER ", action)
		
		if *logic_option != "" {
			sql_command += fmt.Sprintf("%s ", *logic_option)
		}
		
		sql_command += fmt.Sprintf("'%s' ", *(*(data.M("credentials").GetObject("value").(*Credentials))).GetUsername())
		sql_command += fmt.Sprintf("@'%s' ",*(*(data.M("domain_name").GetObject("value").(*DomainName))).GetDomainName())
		sql_command += fmt.Sprintf("IDENTIFIED BY ")
		sql_command += fmt.Sprintf("'%s'",  *(*(data.M("credentials").GetObject("value").(*Credentials))).GetPassword())

		sql_command += ";"
		return &sql_command, nil
	}

	create := func () (*string, []error) {
		var errors []error 
		sql_command, sql_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getHost(), getDatabase(), getHostCredentials(), sql_command, true)
		
		if *stderr != "" {
			if strings.Contains(*stderr, "Operation CREATE USER failed for") {
				errors = append(errors, fmt.Errorf("create user failed most likely the user already exists"))
			} else {
				errors = append(errors, fmt.Errorf("unknown create user error" + *stderr))
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
		
	return &User{
			Create: func() (*string, []error) {
				return create()
			},
		}, nil
}
