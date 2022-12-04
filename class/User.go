package class

import (
	"fmt"
)

type User struct {
	Validate       func() []error
	Create         func() []error
	GetCredentials func() (*Credentials, []error)
	GetClient func() (*Client, []error)
	GetDomainName  func() (*DomainName, []error)
	UpdatePassword func(new_password string) []error
	Exists 		   func() (*bool, []error)
}

func newUser(client *Client, credentials *Credentials, domain_name *DomainName) (*User, []error) {
	struct_type := "*User"

	SQLCommand := newSQLCommand()

	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]": Map{"[client]":client, "[credentials]":credentials, "[domain_name]":domain_name},
		"[system_schema]": Map{"[client]":Map{"type":"*class.Client", "mandatory": true},
		                "[credentials]":Map{"type":"*class.Credentials", "mandatory": true},
						"[domain_name]":Map{"type":"*class.DomainName", "mandatory": true},
		},
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "User")
	}

	getClient := func() (*Client, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client]", "*class.Client")
		return temp_value.(*Client), temp_value_errors
	}

	getCredentials := func() (*Credentials, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[credentials]", "*class.Credentials")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*Credentials), nil
	}

	getDomainName := func() (*DomainName, []error) {
		temp_value, temp_value_errors :=  GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[domain_name]", "*class.DomainName")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*DomainName), nil
	}

	getCreateSQL := func() (*string, Map, []error) {
		options := Map{"use_file": true}

		errors := validate()
		if len(errors) > 0 {
			return nil, nil, errors
		}

		temp_credentials, temp_credentials_errors := getCredentials()
		if temp_credentials_errors != nil {
			return nil, nil, temp_credentials_errors
		}

		temp_username, temp_username_errors := temp_credentials.GetUsername()
		if temp_username_errors != nil {
			return nil, nil, temp_username_errors
		}

		temp_password, temp_password_errors := temp_credentials.GetPassword()
		if temp_password_errors != nil {
			return nil, nil, temp_password_errors
		} else if IsNil(temp_password) {
			errors = append(errors, fmt.Errorf("User.getCreateSQL password is nil"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		temp_domain_name, temp_domain_name_errors := getDomainName()
		if temp_domain_name_errors != nil {
			return nil, nil, temp_domain_name_errors
		}

		temp_domain_name_value, temp_domain_name_value_errors := temp_domain_name.GetDomainName()
		if temp_domain_name_value_errors != nil {
			return nil, nil, temp_domain_name_value_errors
		}

		sql_command := "CREATE USER "
		sql_command += fmt.Sprintf("'%s'", EscapeString(temp_username))
		sql_command += fmt.Sprintf("@'%s' ", EscapeString(temp_domain_name_value))
		sql_command += fmt.Sprintf("IDENTIFIED BY ")
		sql_command += fmt.Sprintf("'%s'", EscapeString(*temp_password))

		sql_command += ";"
		return &sql_command, options, nil
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &User{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			sql_command, options, sql_command_errors := getCreateSQL()

			if sql_command_errors != nil {
				return sql_command_errors
			}

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetCredentials: func() (*Credentials, []error) {
			return getCredentials()
		},
		GetClient: func() (*Client, []error) {
			return getClient()
		},
		GetDomainName: func() (*DomainName, []error) {
			return getDomainName()
		},
		Exists: func() (*bool, []error) {
			temp_client, temp_client_errors := getClient()
			if temp_client != nil {
				return nil, temp_client_errors
			}

			temp_credentials, temp_credentials_errors := getCredentials() 
			if temp_credentials_errors != nil {
				return nil, temp_credentials_errors
			}

			temp_username, temp_username_errors := temp_credentials.GetUsername()
			if temp_username_errors != nil {
				return nil, temp_username_errors
			}

			return temp_client.UserExists(temp_username)
		},
		UpdatePassword: func(new_password string) []error {
			var errors []error

			if new_password == "" {
				errors = append(errors, fmt.Errorf("new password is empty"))
			}

			validate_errors := validate()
			if validate_errors != nil {
				errors = append(errors, validate_errors...)
			}

			password_data := Map{
				"[fields]": Map{},
				"[schema]": Map{},
				"[system_fields]":Map{"[password]":new_password},
				"[system_schema]":Map{"[password]": Map{"type":"*string","mandatory": false, 
					FILTERS(): Array{Map{"values": GetCredentialPasswordValidCharacters(), "function": getWhitelistCharactersFunc()}}},
				},
			}

			validate_password_errors := ValidateData(&password_data, "NewUserPassword")
			if validate_password_errors != nil {
				errors = append(errors, validate_password_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			temp_host, temp_host_errors := temp_client.GetHost()
			if temp_host_errors != nil {
				return temp_host_errors
			}

			temp_host_name, temp_host_name_errors := temp_host.GetHostName()
			if temp_host_name_errors != nil {
				return temp_host_name_errors
			}

			temp_credentials, temp_credentials_errors := getCredentials()
			if temp_credentials_errors != nil {
				return temp_credentials_errors
			}

			temp_username, temp_username_errors := temp_credentials.GetUsername()
			if temp_username_errors != nil {
				return temp_username_errors
			}

			sql_command := fmt.Sprintf("ALTER USER '%s'@'%s' IDENTIFIED BY '%s'", EscapeString(temp_username), EscapeString(temp_host_name), EscapeString(new_password))

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(client, &sql_command, Map{"use_file": true})

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
	}, nil
}
