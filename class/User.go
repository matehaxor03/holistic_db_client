package class

import (
	"fmt"
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
	Create func() (*string, *string, []error)
}

func NewUser(host *Host, host_credentials *Credentials, credentials *Credentials, domain_name *DomainName, options map[string]map[string][][]string) (*User, []error) {
	bashCommand := newBashCommand()
	
	data := Map {
		"host":Map{"type":"*Host","value":CloneHost(host),"mandatory":true},
		"host_credentials":Map{"type":"*Credentials","value":CloneCredentials(host_credentials),"mandatory":true},
		"credentials":Map{"type":"*Credentials","value":CloneCredentials(credentials),"mandatory":true},
		"domain_name":Map{"type":"*DomainName","value":CloneDomainName(domain_name),"mandatory":true},
		"options":Map{"type":"map[string]map[string][][]string)","value":options,"mandatory":false},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "User")
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

		host_command, host_command_errors := (*(data.M("host").GetObject("value").(*Host))).GetCLSCommand()
		if host_command_errors != nil {
			errors = append(errors, host_command_errors...)	
		}

		credentials_command, credentials_command_errors := (*(data.M("host_credentials").GetObject("value").(*Credentials))).GetCLSCommand()
		if credentials_command_errors != nil {
			errors = append(errors, credentials_command_errors...)	
		}

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", *host_command, *credentials_command) 
		sql_command += fmt.Sprintf(" -e \"%s USER ", action)
		
		if *logic_option != "" {
			sql_command += fmt.Sprintf("%s ", *logic_option)
		}
		
		sql_command += fmt.Sprintf("'%s' ", *(*(data.M("credentials").GetObject("value").(*Credentials))).GetUsername())
		sql_command += fmt.Sprintf("@'%s' ",*(*(data.M("domain_name").GetObject("value").(*DomainName))).GetDomainName())
		sql_command += fmt.Sprintf("IDENTIFIED BY ")
		sql_command += fmt.Sprintf("'%s' ",  *(*(data.M("credentials").GetObject("value").(*Credentials))).GetPassword())

		sql_command += ";\""
		return &sql_command, nil
	}

	create := func () (*string, *string, []error) {
		var errors []error 
		crud_sql_command, crud_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if crud_command_errors != nil {
			errors = append(errors, crud_command_errors...)	
		}
	
		if errors != nil {
			return nil, nil, errors
		}
	
		return bashCommand.ExecuteUnsafeCommand(crud_sql_command)
	}	

	errors := validate()

	if errors != nil {
		return nil, errors
	}
		
	return &User{
			Create: func() (*string, *string, []error) {
				return create()
			},
		}, nil
}
