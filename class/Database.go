package class

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type Database struct {
	Validate        func() []error
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (*bool, []error)
	TableExists     func(table_name string) (*bool, []error)
	GetDatabaseName func() (string, []error)
	SetDatabaseName func(database_name string) []error
	SetClient       func(client *Client) []error
	GetClient       func() (*Client, []error)
	CreateTable     func(table_name string, schema json.Map) (*Table, []error)
	GetTableInterface func(table_name string, schema json.Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() (*[]Table, []error)
	GetTableNames   func() (*[]string, []error)
	ToJSONString    func(json *strings.Builder) ([]error)
	UseDatabase     func() []error
	ExecuteUnsafeCommand func(sql_command *string, options json.Map) (*json.Array, []error)
}

func newDatabase(client Client, database_name string, database_create_options *DatabaseCreateOptions, database_reserved_words_obj *DatabaseReservedWords, database_name_whitelist_characters_obj *DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *TableNameCharacterWhitelist, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Database, []error) {
	var errors []error
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	var this_database *Database
	struct_type := "*class.Database"

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	//database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	database_name_whitelist_characters := database_name_whitelist_characters_obj.GetDatabaseNameCharacterWhitelist()

	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]": json.Map{
			"[client]":client, "[database_name]":database_name,"[database_create_options]":database_create_options},
		"[system_schema]": json.Map{
			"[client]": json.Map{"type":"class.Client"},
			"[database_name]": json.Map{"type":"string","min_length":2,"not_empty_string_value":true,
			"filters": json.Array{ json.Map{"values":database_name_whitelist_characters,"function":getWhitelistCharactersFunc()}}},
			"[database_create_options]": json.Map{"type":"*class.DatabaseCreateOptions"}},
	}

	getData := func() (*json.Map) {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[database_create_options]", "*class.DatabaseCreateOptions")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*DatabaseCreateOptions), nil
	}

	getClient := func() (*Client, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client]", "*class.Client")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*Client), nil
	}

	setClient := func(new_client *Client) []error {
		return SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client]", new_client)
	}

	getDatabaseName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} else if common.IsNil(temp_value) {
			return "", nil
		}
		return temp_value.(string), nil
	}

	setDatabaseName := func(new_database_name string) []error {
		return SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_name]", &new_database_name)
	}

	getCreateSQL := func(options json.Map) (*string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, errors
		}

		sql_command := "CREATE DATABASE "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s` ", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\` ", database_name_escaped)
		}

		databaseCreateOptions, databaseCreateOptions_errors := getDatabaseCreateOptions()
		if databaseCreateOptions_errors != nil {
			return nil, databaseCreateOptions_errors
		}

		if databaseCreateOptions != nil {
			character_set, character_set_errors := databaseCreateOptions.GetCharacterSet()
			if character_set_errors != nil {
				return nil, character_set_errors
			}

			collate, collate_errors := databaseCreateOptions.GetCollate()
			if collate_errors != nil {
				return nil, collate_errors
			}

			if character_set != nil && *character_set != "" {
				sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
			}

			if collate != nil && *collate != "" {
				sql_command += fmt.Sprintf("COLLATE %s ", *collate)
			}
		}
		sql_command += ";"

		if len(errors) > 0 {
			return nil, errors
		}

		return &sql_command, nil
	}

	create := func() []error {
		options := json.Map{"use_file":false, "creating_database":true}

		sql_command, generate_sql_errors := getCreateSQL(options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}
		_, execute_sql_command_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql_command, options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}

		return nil
	}

	exists := func() (*bool, []error) {
		options := json.Map{"use_file":false, "checking_database_exists":true}

		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, errors
		}

		sql_command := "USE "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		boolean_value := false
		if len(errors) > 0 {
			//todo: check error message e.g database does not exist
			boolean_value = false
			return &boolean_value, nil
		}

		boolean_value = true
		return &boolean_value, nil
	}

	getTableNames := func() (*[]string, []error) {
		options := json.Map{"use_file": false}

		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, errors
		}

		sql_command := "SHOW TABLES IN "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		records, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if records == nil {
			errors = append(errors, fmt.Errorf("error: Database: getTableNames returned nil records"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		var table_names []string
		column_name := "Tables_in_" + database_name_escaped
		for _, record := range *records {
			table_name, table_name_errors := record.(json.Map).GetString(column_name)
			if table_name_errors != nil {
				errors = append(errors, table_name_errors...)
				continue
			} else if table_name == nil {
				errors = append(errors, fmt.Errorf("error: Database: getTableNames(%s) was nil available fields are: %s", column_name, record.(json.Map).Keys()))
				continue
			}
			table_names = append(table_names, *table_name)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &table_names, nil
	}

	delete := func() ([]error) {
		options := json.Map{"use_file": false, "deleting_database":true}
		
		errors := validate()

		if len(errors) > 0 {
			return errors
		}


		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}


		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return errors
		}

		sql_command := "DROP DATABASE "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		options := json.Map{"use_file": false, "deleting_database":true}
		
		errors := validate()

		if len(errors) > 0 {
			return errors
		}


		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return errors
		}

		sql_command := "DROP DATABASE IF EXISTS "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`;", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`;", database_name_escaped)
		}

		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getTable := func(table_name string) (*Table, []error) {
		var errors []error
		database := getDatabase()
		
		database_errors := database.Validate()
		if database_errors != nil {
			errors = append(errors, database_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		table, table_errors := newTable(*getDatabase(), table_name, nil, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if table_errors != nil {
			errors = append(errors, table_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		table_schema, schema_errors := table.GetSchema()

		if schema_errors != nil {
			errors = append(errors, schema_errors...)
		}

		if table_schema == nil {
			errors = append(errors, fmt.Errorf("error: Database.getTable schema is nil for table: %s", table_name))
		}

		if len(errors) > 0 {
			return nil, errors
		}		

		get_table, get_table_errors := newTable(*getDatabase(), table_name, *table_schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

		if get_table_errors != nil {
			return nil, get_table_errors
		}

		return get_table, nil
	}

	x := Database{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			errors := create()
			if errors != nil {
				return errors
			}

			return nil
		},
		Delete: func() []error {
			return delete()
		},
		Exists: func() (*bool, []error) {
			return exists()
		},
		DeleteIfExists: func() []error {
			errors := validate()

			if len(errors) > 0 {
				return errors
			}

			return deleteIfExists()
		},
		TableExists: func(table_name string) (*bool, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			table, table_errors := newTable(*getDatabase(), table_name, nil, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if table_errors != nil {
				return nil, table_errors
			}

			return table.Exists()
		},
		GetTableInterface: func(table_name string, schema json.Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := newTable(*getDatabase(), table_name, schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if new_table_errors != nil {
				return nil, new_table_errors
			}

			return table, nil
		},
		CreateTable: func(table_name string, schema json.Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := newTable(*getDatabase(), table_name, schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if new_table_errors != nil {
				return nil, new_table_errors
			}

			create_table_errors := table.Create()
			if create_table_errors != nil {
				return nil, create_table_errors
			}

			return table, nil
		},
		GetTable: func(table_name string) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			return getTable(table_name)
		},
		GetTableNames: func()  (*[]string, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			return getTableNames()
		},
		GetTables: func() (*[]Table, []error) {
			var errors []error
			database := getDatabase()
			
			database_errors := database.Validate()
			if database_errors != nil {
				errors = append(errors, database_errors...)
			}
			

			if len(errors) > 0 {
				return nil, errors
			}

			var tables []Table
			table_names, table_names_errors := getTableNames()
			if table_names_errors != nil {
				errors = append(errors, table_names_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			for _, table_name := range *table_names {
				table, table_errors := getTable(table_name)
				if table_errors != nil {
					errors = append(errors, table_errors...)
					continue
				}

				tables = append(tables, *table)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &tables, nil
		},
		SetClient: func(client *Client) []error {
			var errors []error
			if client == nil {
				errors = append(errors, fmt.Errorf("error: client is nil"))
				return errors
			}

			client_errors := client.Validate()
			if client_errors != nil {
				errors = append(errors, client_errors...)
			}

			if len(errors) > 0 {
				return errors
			}

			setClient(client)
			return nil
		},
		GetClient: func() (*Client, []error) {
			return getClient()
		},
		GetDatabaseName: func() (string, []error) {
			return getDatabaseName()
		},
		SetDatabaseName: func(database_name string) []error {
			return setDatabaseName(database_name)
		},
		ExecuteUnsafeCommand: func(sql_command *string, options json.Map) (*json.Array, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			return SQLCommand.ExecuteUnsafeCommand(*temp_client, sql_command, options)
		},
		UseDatabase: func() []error {
			temp_database := getDatabase()

			temp_client, temp_client_errors := getClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			return temp_client.UseDatabase(*temp_database)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
	}
	setDatabase(&x)

	valiation_errors := validate()
	if valiation_errors != nil {
		errors = append(errors, valiation_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
