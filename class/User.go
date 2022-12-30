package class

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
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

func newUser(client Client, credentials Credentials, domain_name DomainName) (*User, []error) {
	struct_type := "*User"

	var errors []error
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	data := json.Map{}
	data.SetMapValue("[fields]", json.Map{})
	data.SetMapValue("[schema]", json.Map{})

	map_system_fields := json.Map{}
	map_system_fields.SetObject("[client]", client)
	map_system_fields.SetObject("[credentials]", credentials)
	map_system_fields.SetObject("[domain_name]", domain_name)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.Map{}
	
	map_client_schema := json.Map{}
	map_client_schema.SetStringValue("type", "class.Client")
	map_system_schema.SetMapValue("[client]", map_client_schema)


	map_credentials_schema := json.Map{}
	map_credentials_schema.SetStringValue("type", "class.Credentials")
	map_system_schema.SetMapValue("[credentials]", map_credentials_schema)

	map_domain_name_schema := json.Map{}
	map_domain_name_schema.SetStringValue("type", "class.DomainName")
	map_system_schema.SetMapValue("[domain_name]", map_domain_name_schema)

	data.SetMapValue("[system_schema]", map_system_schema)


	/*
	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]": json.Map{"[client]":client, "[credentials]":credentials, "[domain_name]":domain_name},
		"[system_schema]": json.Map{"[client]":json.Map{"type":"class.Client"},
		                "[credentials]":json.Map{"type":"class.Credentials"},
						"[domain_name]":json.Map{"type":"class.DomainName"},
		},
	}*/

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "User")
	}

	getClient := func() (*Client, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client]", "*class.Client")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		
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

	getCreateSQL := func(options json.Map) (*string, json.Map, []error) {
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
		} else if common.IsNil(temp_password) {
			errors = append(errors, fmt.Errorf("error: User.getCreateSQL password is nil"))
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

		username_escaped, username_escaped_errors := common.EscapeString(temp_username, "'")
		if username_escaped_errors != nil {
			errors = append(errors, username_escaped_errors)
		}

		temp_domain_name_value_escaped, temp_domain_name_value_escaped_errors := common.EscapeString(temp_domain_name_value, "'")
		if temp_domain_name_value_escaped_errors != nil {
			errors = append(errors, temp_domain_name_value_escaped_errors)
		}

		temp_password_escaped, temp_password_escaped_errors := common.EscapeString(*temp_password, "'")
		if temp_password_escaped_errors != nil {
			errors = append(errors, temp_password_escaped_errors)
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		if options.IsBoolTrue("use_file") {
			username_escaped = "'" + username_escaped + "'"
		} else {
			username_escaped =  strings.ReplaceAll("'" + username_escaped + "'", "`", "\\`")
		}

		if options.IsBoolTrue("use_file") {
			temp_domain_name_value_escaped = "'" + temp_domain_name_value_escaped + "'"
		} else {
			temp_domain_name_value_escaped =  strings.ReplaceAll("'" + temp_domain_name_value_escaped + "'", "`", "\\`")
		}

		if options.IsBoolTrue("use_file") {
			temp_password_escaped = "'" + temp_password_escaped + "'"
		} else {
			temp_password_escaped =  strings.ReplaceAll("'" + temp_password_escaped + "'", "`", "\\`")
		}


		sql_command := "CREATE USER "
		sql_command += fmt.Sprintf("%s", username_escaped)
		sql_command += fmt.Sprintf("@%s ", temp_domain_name_value_escaped)
		sql_command += fmt.Sprintf("IDENTIFIED BY ")
		sql_command += fmt.Sprintf("%s", temp_password_escaped)

		sql_command += ";"
		return &sql_command, options, nil
	}


	validation_errors := validate()

	if validation_errors != nil {
		errors = append(errors, validation_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &User{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			options := json.Map{}
			options.SetBoolValue("use_file", true)
			sql_command, options, sql_command_errors := getCreateSQL(options)

			if sql_command_errors != nil {
				return sql_command_errors
			}

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql_command, options)

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
			options := json.Map{}
			options.SetBoolValue("use_file", true)
			var errors []error

			validate_errors := validate()
			if validate_errors != nil {
				errors = append(errors, validate_errors...)
			}

			if len(new_password) == 0 {
				errors = append(errors, fmt.Errorf("password cannot be empty"))
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

			temp_username_escaped, temp_username_escaped_errors := common.EscapeString(temp_username, "'")
			if temp_username_escaped_errors != nil {
				errors = append(errors, temp_username_escaped_errors)
			}

			temp_host_name_escaped, temp_host_name_escaped_errors := common.EscapeString(temp_host_name, "'")
			if temp_host_name_escaped_errors != nil {
				errors = append(errors, temp_host_name_escaped_errors)
			}

			new_password_escaped, new_password_escaped_errors := common.EscapeString(new_password, "'")
			if new_password_escaped_errors != nil {
				errors = append(errors, new_password_escaped_errors)
			}

			if options.IsBoolTrue("use_file") {
				temp_username_escaped = "'" + temp_username_escaped + "'"
			} else {
				temp_username_escaped =  strings.ReplaceAll("'" + temp_username_escaped + "'", "`", "\\`")
			}

			if options.IsBoolTrue("use_file") {
				temp_host_name_escaped = "'" + temp_host_name_escaped + "'"
			} else {
				temp_host_name_escaped =  strings.ReplaceAll("'" + temp_host_name_escaped + "'", "`", "\\`")
			}

			if options.IsBoolTrue("use_file") {
				new_password_escaped = "'" + new_password_escaped + "'"
			} else {
				new_password_escaped =  strings.ReplaceAll("'" + new_password_escaped + "'", "`", "\\`")
			}

			sql_command := fmt.Sprintf("ALTER USER %s@%s IDENTIFIED BY %s", temp_username_escaped, temp_host_name_escaped, new_password_escaped)
			if len(errors) > 0 {
				return errors
			}

			options_update := json.Map{}
			options_update.SetBoolValue("use_file", true)
			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(client, &sql_command, options_update)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
	}, nil
}
