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
			FILTERS(): Array{Map{"values": GetDatabaseNameValidCharacters(), "function": getWhitelistCharactersFunc()}}}
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
		return ValidateData(data.Clone(), "Grant")
	}

	getClient := func() *Client {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	getUser := func() *User {
		return CloneUser(data.M("[user]").GetObject("value").(*User))
	}

	getGrantValue := func() string {
		grant, _ := data.M("[grant]").GetString("value")
		g := CloneString(grant)
		return *g
	}

	getDatabaseFilter := func() *string {
		database_filter, _ := data.M("[database_filter]").GetString("value")
		return CloneString(database_filter)
	}

	getTableFilter := func() *string {
		table_filter, _ := data.M("[table_filter]").GetString("value")
		return CloneString(table_filter)
	}

	getSQL := func() (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}

		user := getUser()
		credentials := (*user).GetCredentials()
		domain_name := (*user).GetDomainName()

		grant_value := getGrantValue()
		username_value := *((*credentials).GetUsername())
		domain_name_value := *((*domain_name).GetDomainName())

		database_filter := getDatabaseFilter()
		table_filter := getTableFilter()

		sql := ""
		if database_filter != nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s.%s ", grant_value, *database_filter, *table_filter)
		} else if database_filter != nil && table_filter == nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", grant_value, *database_filter)
		} else if database_filter == nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", grant_value, *table_filter)
		} else {
			errors = append(errors, fmt.Errorf("Grant: getSQL: both database_filter and table_filter were nil"))
		}

		sql += fmt.Sprintf("To '%s'@'%s';", username_value, domain_name_value)

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

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": true})

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
			cloned, _ := NewGrant(getClient(), getUser(), getGrantValue(), getDatabaseFilter(), getTableFilter())
			return cloned
		},
		Grant: func() []error {
			return grant()
		},
	}

	return &x, nil
}
