package class

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
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

func GET_ALLOWED_GRANTS() json.Map {
	valid := json.NewMapValue()
	valid.SetNil(GRANT_ALL())
	valid.SetNil(GRANT_INSERT())
	valid.SetNil(GRANT_UPDATE())
	valid.SetNil(GRANT_SELECT())
	return valid
}

func GET_ALLOWED_FILTERS() json.Map {
	valid := json.NewMapValue()
	valid.SetNil("*")
	return valid
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
	

data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_schema := json.NewMapValue()

	// Start Client
	map_system_fields.SetObjectForMap("[client]", client)
	map_client_schema := json.NewMapValue()
	map_client_schema.SetStringValue("type", "class.Client")
	map_system_schema.SetMapValue("[client]", map_client_schema)
	// End Client


	// Start User
	map_system_fields.SetObjectForMap("[user]", user)
	map_user_schema := json.NewMapValue()
	map_user_schema.SetStringValue("type", "class.User")
	map_system_schema.SetMapValue("[user]", map_user_schema)
	// End User


	// Start Grant
	map_system_fields.SetObjectForMap("[grant]", grant)
	map_grant_schema := json.NewMapValue()
	map_grant_schema.SetStringValue("type", "string")

	map_grant_schema_filters := json.NewArrayValue()
	map_grant_schema_filter := json.NewMapValue()
	map_grant_schema_filter.SetObjectForMap("values", GET_ALLOWED_GRANTS())
	map_grant_schema_filter.SetObjectForMap("function",  getWhitelistStringFunc())
	map_grant_schema_filters.AppendMapValue(map_grant_schema_filter)
	map_grant_schema.SetArrayValue("filters", map_grant_schema_filters)
	map_system_schema.SetMapValue("[grant]", map_grant_schema)
	// End Grant


	// Start Database Filter
	if database_filter != nil {
		map_system_fields.SetObjectForMap("[database_filter]", database_filter)
		map_database_filter_schema := json.NewMapValue()
		map_database_filter_schema.SetStringValue("type", "string")

		map_database_filter_schema_filters := json.NewArrayValue()
		if *database_filter == "*" {
			map_database_filter_schema_filter := json.NewMapValue()
			map_database_filter_schema_filter.SetObjectForMap("values", GET_ALLOWED_FILTERS())
			map_database_filter_schema_filter.SetObjectForMap("function",  getWhitelistCharactersFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter)
		} else {
			map_database_filter_schema_filter1 := json.NewMapValue()
			map_database_filter_schema_filter1.SetObjectForMap("values", database_name_whitelist_characters)
			map_database_filter_schema_filter1.SetObjectForMap("function",  getWhitelistCharactersFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter1)

			map_database_filter_schema_filter2 := json.NewMapValue()
			map_database_filter_schema_filter2.SetObjectForMap("values", database_reserved_words)
			map_database_filter_schema_filter2.SetObjectForMap("function",  getBlacklistStringToUpperFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter2)
		}
		map_database_filter_schema.SetArrayValue("filters", map_database_filter_schema_filters)
		map_system_schema.SetMapValue("[database_filter]", map_database_filter_schema)
	}
	// End Database Filter


	// Start Table Filter
	if table_filter != nil {
		map_system_fields.SetObjectForMap("[table_filter]", table_filter)
		map_table_filter_schema := json.NewMapValue()
		map_table_filter_schema.SetStringValue("type", "string")

		map_table_filter_schema_filters := json.NewArrayValue()
		if *database_filter == "*" {
			map_table_filter_schema_filter := json.NewMapValue()
			map_table_filter_schema_filter.SetObjectForMap("values", GET_ALLOWED_FILTERS())
			map_table_filter_schema_filter.SetObjectForMap("function",  getWhitelistCharactersFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter)
		} else {
			map_table_filter_schema_filter1 := json.NewMapValue()
			map_table_filter_schema_filter1.SetObjectForMap("values", table_name_whitelist_characters_obj)
			map_table_filter_schema_filter1.SetObjectForMap("function",  getWhitelistCharactersFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter1)

			map_table_filter_schema_filter2 := json.NewMapValue()
			map_table_filter_schema_filter2.SetObjectForMap("values", database_reserved_words)
			map_table_filter_schema_filter2.SetObjectForMap("function",  getBlacklistStringToUpperFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter2)
		}
		map_table_filter_schema.SetArrayValue("filters", map_table_filter_schema_filters)
		map_system_schema.SetMapValue("[table_filter]", map_table_filter_schema)
	}
	// End Table Filter

	

	data.SetMapValue("[system_fields]", map_system_fields)
	data.SetMapValue("[system_schema]", map_system_schema)


	/*
	data := json.Map{
		"[fields]": json.NewMapValue(),
		"[schema]": json.NewMapValue(),
		"[system_fields]": json.Map{"[client]":client, "[user]":user, "[grant]":grant},
		"[system_schema]": json.Map{"[client]": json.Map{"type":"class.Client"},
						"[user]": json.Map{"type":"class.User"},
						"[grant]": json.Map{"type":"string","filters": json.Array{json.Map{"values": GET_ALLOWED_GRANTS(), "function": getWhitelistStringFunc()}}},
		},
	}

	if database_filter != nil {
		data["[system_fields]"].(json.Map)["[database_filter]"] = database_filter
		if *database_filter == "*" {
			data["[system_schema]"].(json.Map)["[database_filter]"] = json.Map{"type":"string", "filters": json.Array{json.Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[system_schema]"].(json.Map)["[database_filter]"] = json.Map{"type":"string", "filters": json.Array{json.Map{"values": database_name_whitelist_characters, "function": getWhitelistCharactersFunc()}, json.Map{"values":database_reserved_words,"function":getBlacklistStringToUpperFunc()}}}
		}
	}

	if table_filter != nil {
		data["[system_fields]"].(json.Map)["[table_filter]"] = table_filter
		if *table_filter == "*" {
			data["[system_schema]"].(json.Map)["[table_filter]"] = json.Map{"type":"string", "filters": json.Array{json.Map{"values": GET_ALLOWED_FILTERS(), "function": getWhitelistCharactersFunc()}}}
		} else {
			data["[system_schema]"].(json.Map)["[table_filter]"] = json.Map{"type":"string", "filters": json.Array{json.Map{"values": table_name_whitelist_characters_obj, "function": getWhitelistCharactersFunc()}}}
		}
	}*/


	if table_filter == nil && database_filter == nil {
		errors = append(errors, fmt.Errorf("error: Grant: database_filter and table_filter are both nil"))
	}

	getData := func() *json.Map {
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
		if table_filter_errors != nil {
			return nil, table_filter_errors
		}

		sql := ""
		if database_filter != nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s.%s ", grant_value, *database_filter, *table_filter)
		} else if database_filter != nil && table_filter == nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", grant_value, *database_filter)
		} else if database_filter == nil && table_filter != nil {
			sql = fmt.Sprintf("GRANT %s ON %s ", grant_value, *table_filter)
		} else {
			errors = append(errors, fmt.Errorf("error: Grant: getSQL: both database_filter and table_filter were nil"))
		}

		username_value_escaped, username_value_escaped_errors := common.EscapeString(username_value, "'")
		if username_value_escaped_errors != nil {
			errors = append(errors, username_value_escaped_errors)
			return nil, errors
		}

		domain_name_value_escaped, domain_name_value_escaped_errors := common.EscapeString(domain_name_value, "'")
		if domain_name_value_escaped_errors != nil {
			errors = append(errors, domain_name_value_escaped_errors)
			return nil, errors
		}

		sql += fmt.Sprintf("To '%s'@'%s';", username_value_escaped, domain_name_value_escaped)

		if len(errors) > 0 {
			return nil, errors
		}
		
		return &sql, nil
	}

	executeGrant := func() []error {
		options := json.NewMapValue()
		options.SetBoolValue("use_file", true)
		sql_command, sql_command_errors := getSQL()

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
