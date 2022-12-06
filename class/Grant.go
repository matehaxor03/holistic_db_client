package class

import (
	"fmt"
)

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
	Validate      func() []error
	Grant         func() []error
}

func newGrant(client Client, user User, grant string, database_filter *string, table_filter *string, database_reserved_words_obj *DatabaseReservedWords, database_name_whitelist_characters_obj *DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *TableNameCharacterWhitelist) (*Grant, []error) {
	struct_type := "*Grant"

	var errors []error
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	database_name_whitelist_characters := database_name_whitelist_characters_obj.GetDatabaseNameCharacterWhitelist()
	
	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]": Map{"[client]":client, "[user]":user, "[grant]":grant},
		"[system_schema]": Map{"[client]":Map{"type":"class.Client"},
						"[user]":Map{"type":"class.User"},
						"[grant]": Map{"type":"string",
		"filters]": Array{Map{"values": GET_ALLOWED_GRANTS(), "function": getWhitelistStringFunc()}}},
		},
	}

	if database_filter != nil {
		data["[system_fields]"].(Map)["[database_filter]"] = database_filter
		if *database_filter == "*" {
			data["[ssystem_chema]"].(Map)["[database_filter]"] = Map{"type":"string", "filters": Array{Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[system_schema]"].(Map)["[database_filter]"] = Map{"type":"string", "filters": Array{Map{"values": database_name_whitelist_characters, "function": getWhitelistCharactersFunc()}, Map{"values":database_reserved_words,"function":getBlacklistStringToUpperFunc()}}}
		}
	}

	if table_filter != nil {
		data["[system_fields]"].(Map)["[table_filter]"] = table_filter
		if *table_filter == "*" {
			data["[system_schema]"].(Map)["[table_filter]"] = Map{"type":"string", "filters": Array{Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[system_schema]"].(Map)["[table_filter]"] = Map{"type":"string", "filters": Array{Map{"values": table_name_whitelist_characters_obj, "function": getWhitelistCharactersFunc()}}}
		}
	}


	if table_filter == nil && database_filter == nil {
		errors = append(errors, fmt.Errorf("error: Grant: database_filter and table_filter are both nil"))
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "Grant")
	}

	getClient := func() (*Client, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[client]", "*class.Client")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		}
		return temp_value.(*Client), temp_value_errors
	}

	getUser := func() (*User, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[user]", "*class.User")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		}
		return temp_value.(*User), temp_value_errors
	}

	getGrantValue := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[grant]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		}
		return temp_value.(string), temp_value_errors
	}

	getDatabaseFilter := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_filter]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		
		return temp_value.(*string), nil
	}

	getTableFilter := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[table_filter]", "*string")
		
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		
		return temp_value.(*string), nil
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
			errors = append(errors, fmt.Errorf("error: Grant: getSQL: both database_filter and table_filter were nil"))
		}

		sql += fmt.Sprintf("To '%s'@'%s';", EscapeString(username_value), EscapeString(domain_name_value))

		if len(errors) > 0 {
			return nil, errors
		}
		
		return &sql, nil
	}

	executeGrant := func() []error {
		sql_command, sql_command_errors := getSQL()

		if sql_command_errors != nil {
			return sql_command_errors
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql_command, Map{"use_file": true})

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

	return &Grant{
		Validate: func() []error {
			return validate()
		},
		Grant: func() []error {
			return executeGrant()
		},
	}, nil
}
