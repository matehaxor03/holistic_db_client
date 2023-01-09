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

	executeUnsafeCommand := func(sql_command *string, options *json.Map) (*json.Array, []error) {
		errors := validate()
		if errors != nil {
			return nil, errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}
		
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(temp_database, sql_command, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		} else if common.IsNil(sql_command_results) {
			errors = append(errors, fmt.Errorf("records from db was nil"))	
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return sql_command_results, nil
	}

	getCreateSQL := func(options *json.Map) (*string, *json.Map, []error) {
		if common.IsNil(options) {
			options := json.NewMap()
			options.SetBoolValue("use_file", true)
		}
		
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

		temp_password_escaped, temp_password_escaped_errors := common.EscapeString(temp_password, "'")
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
			options := json.NewMap()
			options.SetBoolValue("use_file", true)
			options.SetBoolValue("use_mysql_database", true)
			sql_command, new_options, sql_command_errors := getCreateSQL(options)

			if sql_command_errors != nil {
				return sql_command_errors
			}

			_, execute_errors := executeUnsafeCommand(sql_command, new_options)

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
			options := json.NewMap()
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


			temp_host, temp_host_errors := database.GetHost()
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

			options_update := json.NewMap()
			options_update.SetBoolValue("use_file", true)
			_, execute_errors := executeUnsafeCommand(&sql_command, options_update)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
	}, nil
}
