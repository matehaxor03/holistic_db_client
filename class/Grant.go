package class

import (
	"fmt"
)

func CloneGrant(grant *Grant) *Grant {
	if grant == nil {
		return nil
	}

	return grant.Clone()
}

func GRANT_ALL() string {
	return "ALL"
}

func GRANT_INSERT() string {
	return "INSERT"
}

func GRANT_UPDATE() string {
	return "UPDATE"
}

func GRANT_SELECT() string {
	return "SELECT"
}

func GET_ALLOWED_GRANTS() Map {
	return Map{GRANT_ALL(): nil, GRANT_INSERT(): nil, GRANT_UPDATE(): nil, GRANT_SELECT(): nil}
}

func GET_ALLOWED_FILTERS() Map {
	return Map{"*": nil}
}

type Grant struct {
	Clone         func() *Grant
	Validate      func() []error
	Grant         func() []error
}

func NewGrant(client *Client, user *User, grant_value string, database_filter *string, table_filter *string) (*Grant, []error) {
	var errors []error
	SQLCommand := NewSQLCommand()

	data := Map{
		"[client]": Map{"value": CloneClient(client), "mandatory": true},
		"[user]":   Map{"value": CloneUser(user), "mandatory": true},
		"[grant]": Map{"value": CloneString(&grant_value), "mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_GRANTS(), "function": getWhitelistStringFunc()}}},
	}

	if database_filter != nil {
		if *database_filter == "*" {
			data["[database_filter]"] = Map{"value": CloneString(database_filter), "mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[database_filter]"] = Map{"type": "*string", "value":CloneString(database_filter), "mandatory": true,
			FILTERS(): Array{Map{"values": GetDatabaseNameWhitelistCharacters(), "function": getWhitelistCharactersFunc()}}}
		}
	}

	if table_filter != nil {
		if *table_filter == "*" {
			data["[table_filter]"] = Map{"value": CloneString(table_filter), "mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[table_filter]"] = Map{"type": "*string", "value":CloneString(table_filter), "mandatory": true,
			FILTERS(): Array{Map{"values": GetTableNameValidCharacters(), "function": getWhitelistCharactersFunc()}}}
		}
	}


	if table_filter == nil && database_filter == nil {
		errors = append(errors, fmt.Errorf("Grant: database_filter and table_filter are both nil"))
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Grant")
	}

	getClient := func() (*Client, []error) {
		temp_client_map, temp_client_map_errors := data.GetMap("[client]")
		if temp_client_map_errors != nil {
			return nil, temp_client_map_errors
		}

		temp_client_value := temp_client_map.GetObject("value").(*Client)
		return CloneClient(temp_client_value), nil
	}

	getUser := func() (*User, []error) {
		temp_user_map, temp_user_map_errors := data.GetMap("[user]")
		if temp_user_map_errors != nil {
			return nil, temp_user_map_errors
		}

		temp_user_value := temp_user_map.GetObject("value").(*User)
		return CloneUser(temp_user_value), nil
		//return CloneUser(data.M("[user]").GetObject("value").(*User))
	}

	getGrantValue := func() (string, []error) {
		temp_grant_map, temp_grant_map_errors := data.GetMap("[grant]")
		if temp_grant_map_errors != nil {
			return "", temp_grant_map_errors
		}

		temp_grant_value, temp_grant_value_errors := temp_grant_map.GetString("value")
		if temp_grant_value_errors != nil {
			return "", temp_grant_value_errors
		}
		//grant, _ := data.M("[grant]").GetString("value")
		g := CloneString(temp_grant_value)
		return *g, nil
	}

	getDatabaseFilter := func() (*string, []error) {
		database_filter_map, database_filter_map_errors := data.GetMap("[database_filter]")
		if database_filter_map_errors != nil {
			return nil, database_filter_map_errors
		}
		
		database_filter_value, database_filter_value_errors := database_filter_map.GetString("value")
		if database_filter_value_errors != nil {
			return nil, database_filter_value_errors
		}
		//database_filter, _ := data.M("[database_filter]").GetString("value")
		return CloneString(database_filter_value), nil
	}

	getTableFilter := func() (*string, []error) {
		table_filter_map, table_filter_map_errors := data.GetMap("[table_filter]")
		if table_filter_map_errors != nil {
			return nil, table_filter_map_errors
		}

		table_filter_value, table_filter_value_errors := table_filter_map.GetString("value")
		if table_filter_value_errors != nil {
			return nil, table_filter_value_errors
		}
		
		//table_filter, _ := data.M("[table_filter]").GetString("value")
		return CloneString(table_filter_value), nil
	}

	getSQL := func() (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}

		user, user_errors := getUser()
		if user_errors != nil {
			return nil, user_errors
		}

		credentials, credentials_errors := (*user).GetCredentials()
		if credentials_errors != nil {
			return nil, credentials_errors
		}
		
		domain_name, domain_name_errors := (*user).GetDomainName()
		if domain_name_errors != nil {
			return nil, domain_name_errors
		}

		grant_value, grant_value_errors := getGrantValue()
		if grant_value_errors != nil {
			return nil, grant_value_errors
		}

		username_value, username_value_errors := (*credentials).GetUsername()
		if username_value_errors != nil {
			return nil, username_value_errors
		}

		domain_name_value, domain_name_value_errors := (*domain_name).GetDomainName()
		if domain_name_value_errors != nil {
			return nil, domain_name_value_errors
		}

		database_filter, database_filter_errors := getDatabaseFilter()
		if database_filter_errors != nil {
			return nil, database_filter_errors
		}
		
		table_filter, table_filter_errors := getTableFilter()
		if table_filter_errors != nil{
			return nil, table_filter_errors
		}

		sql := ""
		if database_filter != nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s.%s ", EscapeString(grant_value), EscapeString(*database_filter), EscapeString(*table_filter))
		} else if database_filter != nil && table_filter == nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", EscapeString(grant_value), EscapeString(*database_filter))
		} else if database_filter == nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", EscapeString(grant_value), EscapeString(*table_filter))
		} else {
			errors = append(errors, fmt.Errorf("Grant: getSQL: both database_filter and table_filter were nil"))
		}

		sql += fmt.Sprintf("To '%s'@'%s';", EscapeString(username_value), EscapeString(domain_name_value))

		if len(errors) > 0 {
			return nil, errors
		}
		
		return &sql, nil
	}

	grant := func() []error {
		sql_command, sql_command_errors := getSQL()

		if sql_command_errors != nil {
			return sql_command_errors
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, Map{"use_file": true})

		if execute_errors != nil {
			return execute_errors
		}

		return nil
	}

	validation_errors := validate()

	if validation_errors != nil {
		errors = append(errors, validation_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	x := Grant{
		Validate: func() []error {
			return validate()
		},
		Clone: func() *Grant {
			temp_client, _ := getClient()
			temp_user, _ := getUser()
			temp_grant_value, _ := getGrantValue()
			temp_database_filter, _ := getDatabaseFilter()
			temp_table_filter, _ := getTableFilter()
			cloned, _ := NewGrant(temp_client, temp_user, temp_grant_value, temp_database_filter, temp_table_filter)
			return cloned
		},
		Grant: func() []error {
			return grant()
		},
	}

	return &x, nil
}
