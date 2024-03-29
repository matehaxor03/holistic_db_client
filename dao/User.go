package dao

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

type User struct {
	Validate       func() []error
	Create         func() []error
	GetCredentials func() (Credentials, []error)
	GetDatabase func() (Database, []error)
	GetDomainName  func() (DomainName, []error)
	UpdatePassword func(new_password string) []error
}

func newUser(database Database, credentials Credentials, domain_name DomainName) (*User, []error) {
	var errors []error

	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[database]", database)
	map_system_fields.SetObjectForMap("[credentials]", credentials)
	map_system_fields.SetObjectForMap("[domain_name]", domain_name)
	data.SetMapValue("[system_fields]", map_system_fields)

	map_system_schema := json.NewMapValue()
	
	map_client_schema := json.NewMapValue()
	map_client_schema.SetStringValue("type", "dao.Database")
	map_system_schema.SetMapValue("[database]", map_client_schema)


	map_credentials_schema := json.NewMapValue()
	map_credentials_schema.SetStringValue("type", "dao.Credentials")
	map_system_schema.SetMapValue("[credentials]", map_credentials_schema)

	map_domain_name_schema := json.NewMapValue()
	map_domain_name_schema.SetStringValue("type", "dao.DomainName")
	map_system_schema.SetMapValue("[domain_name]", map_domain_name_schema)

	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "User")
	}

	getDatabase := func() (Database, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[database]", "dao.Database")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("database is nil"))
		}
		if len(errors) > 0 {
			return Database{}, errors
		}
		return temp_value.(Database), temp_value_errors
	}

	getCredentials := func() (Credentials, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[credentials]", "dao.Credentials")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("credentials is nil"))
		}

		if len(errors) > 0 {
			return Credentials{}, errors
		}
		return temp_value.(Credentials), nil
	}

	getDomainName := func() (DomainName, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[domain_name]", "dao.DomainName")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("domain name is nil"))
		}

		if len(errors) > 0 {
			return DomainName{}, errors
		}
		return temp_value.(DomainName), nil
	}

	executeUnsafeCommand := func(sql_command strings.Builder, options json.Map) (json.Array, []error) {
		var errors []error
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return json.NewArrayValue(), temp_database_errors
		}
		
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(temp_database, sql_command, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		} else if common.IsNil(sql_command_results) {
			errors = append(errors, fmt.Errorf("records from db was nil"))	
		}

		if len(errors) > 0 {
			return sql_command_results, errors
		}

		return sql_command_results, nil
	}

	getCreateSQL := func(options json.Map) (*strings.Builder, json.Map, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, options, errors
		}

		temp_credentials, temp_credentials_errors := getCredentials()
		if temp_credentials_errors != nil {
			return nil, options, temp_credentials_errors
		}

		temp_username := temp_credentials.GetUsername()
		temp_password := temp_credentials.GetPassword()
	
		if len(errors) > 0 {
			return nil, options, errors
		}

		temp_domain_name, temp_domain_name_errors := getDomainName()
		if temp_domain_name_errors != nil {
			return nil, options, temp_domain_name_errors
		}

		temp_domain_name_value := temp_domain_name.GetDomainName()

		username_escaped, username_escaped_errors := common.EscapeString(temp_username, "'")
		if username_escaped_errors != nil {
			errors = append(errors, username_escaped_errors)
		}

		temp_domain_name_value_escaped, temp_domain_name_value_escaped_errors := common.EscapeString(temp_domain_name_value, "'")
		if temp_domain_name_value_escaped_errors != nil {
			errors = append(errors, temp_domain_name_value_escaped_errors)
		}

		temp_password_escaped, temp_password_escaped_errors := common.EscapeString(temp_password, "'")
		if temp_password_escaped_errors != nil {
			errors = append(errors, temp_password_escaped_errors)
		}

		if len(errors) > 0 {
			return nil, options, errors
		}

		username_escaped = "'" + username_escaped + "'"
		temp_domain_name_value_escaped = "'" + temp_domain_name_value_escaped + "'"
		temp_password_escaped = "'" + temp_password_escaped + "'"

		var sql_command strings.Builder
		sql_command.WriteString("CREATE USER ")
		sql_command.WriteString(fmt.Sprintf("%s", username_escaped))
		sql_command.WriteString(fmt.Sprintf("@%s ", temp_domain_name_value_escaped))
		sql_command.WriteString(fmt.Sprintf("IDENTIFIED BY "))
		sql_command.WriteString(fmt.Sprintf("%s", temp_password_escaped))
		sql_command.WriteString(";")
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
			options := json.NewMapValue()
			options.SetBoolValue("use_mysql_database", true)
			sql_command, new_options, sql_command_errors := getCreateSQL(options)

			if sql_command_errors != nil {
				return sql_command_errors
			}

			_, execute_errors := executeUnsafeCommand(*sql_command, new_options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetCredentials: func() (Credentials, []error) {
			return getCredentials()
		},
		GetDomainName: func() (DomainName, []error) {
			return getDomainName()
		},
		UpdatePassword: func(new_password string) []error {
			var errors []error
			options := json.NewMapValue()

			if len(new_password) == 0 {
				errors = append(errors, fmt.Errorf("password cannot be empty"))
			}

			if len(errors) > 0 {
				return errors
			}

			temp_host := database.GetHost()

			temp_host_name := temp_host.GetHostName()
			
			temp_credentials, temp_credentials_errors := getCredentials()
			if temp_credentials_errors != nil {
				return temp_credentials_errors
			}

			temp_username := temp_credentials.GetUsername()

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

			temp_username_escaped = "'" + temp_username_escaped + "'"
			temp_host_name_escaped = "'" + temp_host_name_escaped + "'"
			new_password_escaped = "'" + new_password_escaped + "'"
	
			var sql_command strings.Builder
			sql_command.WriteString(fmt.Sprintf("ALTER USER %s@%s IDENTIFIED BY %s", temp_username_escaped, temp_host_name_escaped, new_password_escaped))
			if len(errors) > 0 {
				return errors
			}

			_, execute_errors := executeUnsafeCommand(sql_command, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
	}, nil
}
