package class

import (
	"fmt"
)

func CloneUser(user *User) *User {
	if user == nil {
		return user
	}

	return user.Clone()
}

type User struct {
	Validate       func() []error
	Create         func() []error
	GetCredentials func() *Credentials
	GetDomainName  func() *DomainName
	Clone          func() *User
	UpdatePassword func(new_password string) []error
}

func NewUser(client *Client, credentials *Credentials, domain_name *DomainName) (*User, []error) {
	SQLCommand := NewSQLCommand()

	data := Map{
		"[client]":      Map{"value": CloneClient(client), "mandatory": true},
		"[credentials]": Map{"value": CloneCredentials(credentials), "mandatory": true},
		"[domain_name]": Map{"value": CloneDomainName(domain_name), "mandatory": true},
	}

	validate := func() []error {
		return ValidateData(data, "User")
	}

	getClient := func() *Client {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	getCredentials := func() *Credentials {
		return CloneCredentials(data.M("[credentials]").GetObject("value").(*Credentials))
	}

	getDomainName := func() *DomainName {
		return CloneDomainName(data.M("[domain_name]").GetObject("value").(*DomainName))
	}

	getCreateSQL := func() (*string, Map, []error) {
		options := Map{"use_file": true}

		errors := validate()
		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command := "CREATE USER "
		sql_command += fmt.Sprintf("'%s'", EscapeString(*((*getCredentials()).GetUsername())))
		sql_command += fmt.Sprintf("@'%s' ", EscapeString(*((*getDomainName()).GetDomainName())))
		sql_command += fmt.Sprintf("IDENTIFIED BY ")
		sql_command += fmt.Sprintf("'%s'", EscapeString(*((*getCredentials()).GetPassword())))

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

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		Clone: func() *User {
			cloned, _ := NewUser(getClient(), getCredentials(), getDomainName())
			return cloned
		},
		GetCredentials: func() *Credentials {
			return getCredentials()
		},
		GetDomainName: func() *DomainName {
			return getDomainName()
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

			data := Map{
				"password": Map{"value": CloneString(&new_password), "mandatory": true,
					FILTERS(): Array{Map{"values": GetCredentialPasswordValidCharacters(), "function": getWhitelistCharactersFunc()}}},
			}

			validate_password_errors := ValidateData(data.Clone(), "NewUserPassword")
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

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(client, &sql_command, Map{"use_file": true})

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
	}, nil
}
