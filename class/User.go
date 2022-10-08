package class

import (
	"fmt"
	"strings"
)

func CloneUser(user *User) *User {
	if user == nil {
		return user
	}

	return user.Clone()
}

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
	Validate func() ([]error)
	Create func() (*string, []error)
	GetCredentials func() (*Credentials)
	GetDomainName func() (*DomainName)
	Clone func() (*User)
	UpdatePassword func(new_password string) ([]error)
}

func NewUser(client *Client, credentials *Credentials, domain_name *DomainName, options map[string]map[string][][]string) (*User, []error) {
	SQLCommand := newSQLCommand()
	
	data := Map {
		"client":Map{"value":CloneClient(client),"mandatory":true},
		"credentials":Map{"value":CloneCredentials(credentials),"mandatory":true},
		"domain_name":Map{"value":CloneDomainName(domain_name),"mandatory":true},
		"options":Map{"value":options,"mandatory":false},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "User")
	}

	getClient := func() (*Client) {
		return CloneClient(data.M("client").GetObject("value").(*Client))
	}

	getCredentials := func() (*Credentials) {
		return CloneCredentials(data.M("credentials").GetObject("value").(*Credentials))
	}

	getDomainName := func() (*DomainName) {
		return CloneDomainName(data.M("domain_name").GetObject("value").(*DomainName))
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("options").GetObject("value").(map[string]map[string][][]string)
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
		
		sql_command += fmt.Sprintf("'%s'", *(*(data.M("credentials").GetObject("value").(*Credentials))).GetUsername())
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

		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": true})
		
		if *stderr != "" {
			if strings.Contains(*stderr, "Operation CREATE USER failed for") {
				errors = append(errors, fmt.Errorf("create user failed most likely the user already exists"))
			} else {
				errors = append(errors, fmt.Errorf(*stderr))
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
			Validate: func() ([]error) {
				return validate()
			},
			Create: func() (*string, []error) {
				return create()
			},
			Clone: func() *User {
				cloned, _ := NewUser(getClient(), getCredentials(), getDomainName(), getOptions())
				return cloned
			},
			GetCredentials: func() *Credentials {
				return getCredentials()
			},
			GetDomainName: func() *DomainName {
				return getDomainName()
			},
			UpdatePassword: func(new_password string) ([]error) {
				var errors []error

				if new_password == "" {
					errors = append(errors, fmt.Errorf("new password is empty"))
				}
				
				validate_errors := validate()
				if validate_errors != nil {
					errors = append(errors, validate_errors...)
				}
				
				data := Map {
					"password":Map{"value":CloneString(&new_password),"mandatory":true,
					FILTERS(): Array{ Map {"values":GetCredentialPasswordValidCharacters(),"function":getValidateCharacters() }}},
				}

				validate_password_errors := ValidateGenericSpecial(data.Clone(), "NewUserPassword")
				if validate_password_errors != nil {
					errors = append(errors, validate_password_errors...)
				}

				if len(errors) > 0 {
					return errors
				}

				client := getClient()
				host := client.GetHost()
				host_name := host.GetHostName()
				credentials := getCredentials()
				username := credentials.GetUsername()


				sql_command := fmt.Sprintf("ALTER USER '%s'@'%s' IDENTIFIED BY '%s'", *username, *host_name, new_password)

				_, stderr, errors := SQLCommand.ExecuteUnsafeCommand(client, &sql_command, Map{"use_file": true})
		
				if *stderr != "" {
					if strings.Contains(*stderr, "Operation CREATE USER failed for") {
						errors = append(errors, fmt.Errorf("create user failed most likely the user already exists"))
					} else {
						errors = append(errors, fmt.Errorf(*stderr))
					}
				}

				if len(errors) > 0 {
					return errors
				}

				return nil
			},
		}, nil
}
