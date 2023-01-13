package dao

import (
	"fmt"
	"sync"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type Grant struct {
	Validate      func() []error
	Grant         func() []error
}

func newGrant(verify *validate.Validator, database Database, user User, grant string, database_filter *string, table_filter *string, lock_sql_command *sync.RWMutex) (*Grant, []error) {
	var errors []error

	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_schema := json.NewMapValue()

	map_system_fields.SetObjectForMap("[database]", database)
	map_database_schema := json.NewMapValue()
	map_database_schema.SetStringValue("type", "dao.Database")
	map_system_schema.SetMapValue("[database]", map_database_schema)


	map_system_fields.SetObjectForMap("[user]", user)
	map_user_schema := json.NewMapValue()
	map_user_schema.SetStringValue("type", "dao.User")
	map_system_schema.SetMapValue("[user]", map_user_schema)


	map_system_fields.SetObjectForMap("[grant]", grant)
	map_grant_schema := json.NewMapValue()
	map_grant_schema.SetStringValue("type", "string")

	map_grant_schema_filters := json.NewArrayValue()
	map_grant_schema_filter := json.NewMapValue()
	map_grant_schema_filter.SetObjectForMap("function",  verify.GetValidateGrantFunc())
	map_grant_schema_filters.AppendMapValue(map_grant_schema_filter)
	map_grant_schema.SetArrayValue("filters", map_grant_schema_filters)
	map_system_schema.SetMapValue("[grant]", map_grant_schema)


	if database_filter != nil {
		map_system_fields.SetObjectForMap("[database_filter]", database_filter)
		map_database_filter_schema := json.NewMapValue()
		map_database_filter_schema.SetStringValue("type", "string")

		map_database_filter_schema_filters := json.NewArrayValue()
		
		if *database_filter == "*" {
			map_database_filter_schema_filter1 := json.NewMapValue()
			map_database_filter_schema_filter1.SetObjectForMap("function",  verify.GetValidateDatabaseNameFilterAllFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter1)
		} else {
			map_database_filter_schema_filter2 := json.NewMapValue()
			map_database_filter_schema_filter2.SetObjectForMap("function",  verify.GetValidateDatabaseNameFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter2)

			map_database_filter_schema_filter3 := json.NewMapValue()
			map_database_filter_schema_filter3.SetObjectForMap("function",  verify.GetValidateDatabaseReservedWordFunc())
			map_database_filter_schema_filters.AppendMapValue(map_database_filter_schema_filter3)
		}
		map_database_filter_schema.SetArrayValue("filters", map_database_filter_schema_filters)
		map_system_schema.SetMapValue("[database_filter]", map_database_filter_schema)
	}

	if table_filter != nil {
		map_system_fields.SetObjectForMap("[table_filter]", table_filter)
		map_table_filter_schema := json.NewMapValue()
		map_table_filter_schema.SetStringValue("type", "string")

		map_table_filter_schema_filters := json.NewArrayValue()
		if *table_filter == "*" {
			map_table_filter_schema_filter := json.NewMapValue()
			map_table_filter_schema_filter.SetObjectForMap("function",  verify.GetValidateTableNameFilterAllFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter)
		} else {
			map_table_filter_schema_filter1 := json.NewMapValue()
			map_table_filter_schema_filter1.SetObjectForMap("function",  verify.GetValidateTableNameFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter1)

			map_table_filter_schema_filter2 := json.NewMapValue()
			map_table_filter_schema_filter2.SetObjectForMap("function",  verify.GetValidateDatabaseReservedWordFunc())
			map_table_filter_schema_filters.AppendMapValue(map_table_filter_schema_filter2)
		}
		map_table_filter_schema.SetArrayValue("filters", map_table_filter_schema_filters)
		map_system_schema.SetMapValue("[table_filter]", map_table_filter_schema)
	}

	

	data.SetMapValue("[system_fields]", map_system_fields)
	data.SetMapValue("[system_schema]", map_system_schema)

	if table_filter == nil && database_filter == nil {
		errors = append(errors, fmt.Errorf("error: Grant: database_filter and table_filter are both nil"))
	}

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "Grant")
	}

	getDatabase := func() (Database, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]",  "[database]", "dao.Database")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("database is nil"))
		}
		if len(errors) > 0 {
			return Database{}, errors
		}
		return temp_value.(Database), nil
	}

	getUser := func() (User, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]",  "[user]", "dao.User")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("user is nil"))
		}
		if len(errors) > 0 {
			return User{}, errors
		}
		return temp_value.(User), nil
	}

	getGrantValue := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[grant]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("grant value is nil"))
		}
		if len(errors) > 0 {
			return "", errors
		}
		return temp_value.(string), nil
	}

	getDatabaseFilter := func() (*string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[database_filter]", "*string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			return nil, nil
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		return temp_value.(*string), nil
	}

	getTableFilter := func() (*string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[table_filter]", "*string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		
		if len(errors) > 0 {
			return nil, errors
		}
		
		return temp_value.(*string), nil
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
		
		lock_sql_command.Lock()
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(lock_sql_command, temp_database, sql_command, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		} else if common.IsNil(sql_command_results) {
			errors = append(errors, fmt.Errorf("records from db was nil"))	
		}

		if len(errors) > 0 {
			defer lock_sql_command.Unlock()
			return nil, errors
		}

		defer lock_sql_command.Unlock()
		return sql_command_results, nil
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

		credentials, credentials_errors := user.GetCredentials()
		if credentials_errors != nil {
			return nil, credentials_errors
		}
		
		domain_name, domain_name_errors := user.GetDomainName()
		if domain_name_errors != nil {
			return nil, domain_name_errors
		}

		grant_value, grant_value_errors := getGrantValue()
		if grant_value_errors != nil {
			return nil, grant_value_errors
		}

		username_value := credentials.GetUsername()

		domain_name_value := domain_name.GetDomainName()

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
		options := json.NewMap()
		options.SetBoolValue("use_file", true)
		sql_command, sql_command_errors := getSQL()

		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors := executeUnsafeCommand(sql_command, options)

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
