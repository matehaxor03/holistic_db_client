package dao

import (
	"fmt"
	"strings"
	"sync"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"

	helper "github.com/matehaxor03/holistic_db_client/helper"

)

type Database struct {
	Validate        func() []error
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (*bool, []error)
	TableExists     func(table_name string) (*bool, []error)
	GetDatabaseName func() (string, []error)
	GetDatabaseUsername func() (string, []error)
	SetDatabaseUsername func(database_username string) []error
	SetDatabaseName func(database_name string) []error
	DeleteTableByTableNameIfExists func(table_name string, if_exists bool) []error
	GetHost       func() (Host, []error)
	CreateTable     func(table_name string, schema json.Map) (*Table, []error)
	GetTableInterface func(table_name string, schema json.Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() (*[]Table, []error)
	GetTableNames   func() (*[]string, []error)
	ToJSONString    func(json *strings.Builder) ([]error)
	ExecuteUnsafeCommand func(sql_command *string, options *json.Map) (*json.Array, []error)
	GetOrSetSchema func(table_name string, schema json.Map, mode string) (json.Map, []error)
	GetOrSetAdditonalSchema func(table_name string, additional_schema *json.Map) (*json.Map, []error)
	GetOrSetReadRecords func(sql string, records *[]Record) (*[]Record, []error)

	GlobalGeneralLogDisable	func() []error
	GlobalGeneralLogEnable	func() []error
	GlobalSetTimeZoneUTC    func() []error
	GlobalSetSQLMode func() []error
}

func NewDatabase(host Host, database_username string, database_name string, database_create_options *DatabaseCreateOptions, database_reserved_words_obj *validation_constants.DatabaseReservedWords, database_name_whitelist_characters_obj *validation_constants.DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *validation_constants.TableNameCharacterWhitelist, column_name_whitelist_characters_obj *validation_constants.ColumnNameCharacterWhitelist) (*Database, []error) {
	var lock_get_table_schema = &sync.Mutex{}

	var errors []error
	lock_table_schema_cache := &sync.Mutex{}
	table_schema_cache := newTableSchemaCache()

	lock_table_additional_schema_cache := &sync.Mutex{}
	table_additional_schema_cache := newTableAdditionalSchemaCache()

	lock_table_records_cache := &sync.Mutex{}
	table_read_records_cache := newTableReadRecordsCache()

	
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	var this_database *Database
	struct_type := "*dao.Database"

	setDatabase := func(database *Database) {
		this_database = database
	}

	getDatabase := func() *Database {
		return this_database
	}

	//database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	database_name_whitelist_characters := database_name_whitelist_characters_obj.GetDatabaseNameCharacterWhitelist()


	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[host]", host)
	map_system_fields.SetObjectForMap("[database_name]", database_name)
	map_system_fields.SetObjectForMap("[database_username]", database_username)
	map_system_fields.SetObjectForMap("[database_create_options]", database_create_options)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()
	
	map_host_schema := json.NewMapValue()
	map_host_schema.SetStringValue("type", "dao.Host")
	map_system_schema.SetMapValue("[host]", map_host_schema)


	// START database_name
	map_database_name_schema := json.NewMapValue()
	map_database_name_schema.SetStringValue("type", "string")
	map_database_name_schema.SetIntValue("min_length", 2)
	map_database_name_schema.SetBoolValue("not_empty_string_value", true)
	map_database_name_schema_filters := json.NewArrayValue()
	map_database_name_schema_filter := json.NewMapValue()
	map_database_name_schema_filter.SetObjectForMap("values", database_name_whitelist_characters)
	map_database_name_schema_filter.SetObjectForMap("function",  validation_functions.GetWhitelistCharactersFunc())
	map_database_name_schema_filters.AppendMapValue(map_database_name_schema_filter)
	map_database_name_schema.SetArrayValue("filters", map_database_name_schema_filters)
	map_system_schema.SetMapValue("[database_name]", map_database_name_schema)
	// END database_name

	// START database username
	map_database_username := json.NewMapValue()
	map_database_username.SetStringValue("type", "string")
	array_database_username_filters := json.NewArrayValue()
	map_database_username_filter := json.NewMapValue()
	map_database_username_filter.SetObjectForMap("values", validation_constants.GetValidUsernameCharacters())
	map_database_username_filter.SetObjectForMap("function",  validation_functions.GetWhitelistCharactersFunc())
	array_database_username_filters.AppendMapValue(map_database_username_filter)
	map_database_username.SetArrayValue("filters", array_database_username_filters)
	map_system_schema.SetMapValue("[database_username]", map_database_username)
	// End database username


	map_create_options_schema := json.NewMapValue()
	map_create_options_schema.SetStringValue("type", "*dao.DatabaseCreateOptions")
	map_system_schema.SetMapValue("[database_create_options]", map_create_options_schema)

	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() (*json.Map) {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[database_create_options]", "*dao.DatabaseCreateOptions")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*DatabaseCreateOptions), nil
	}

	getHost := func() (Host, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[host]", "dao.Host")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("host is nil"))
		}
		
		if len(errors) > 0 {
			return Host{}, errors
		}

		return temp_value.(Host), nil
	}

	getDatabaseName := func() (string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} else if common.IsNil(temp_value) {
			return "", nil
		}
		return temp_value.(string), nil
	}

	setDatabaseName := func(new_database_name string) []error {
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_name]", new_database_name)
	}

	getDatabaseUsername := func() (string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} else if common.IsNil(temp_value) {
			return "", nil
		}
		return temp_value.(string), nil
	}
	
	setDatabaseUsername := func(new_database_username string) []error {
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", new_database_username)
	}

	executeUnsafeCommand := func(sql_command *string, options *json.Map) (*json.Array, []error) {
		errors := validate()
		if errors != nil {
			return nil, errors
		}
		return SQLCommand.ExecuteUnsafeCommand(getDatabase(), sql_command, options)
	}

	getCreateSQL := func(options *json.Map) (*string, []error) {
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
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("creating_database", true)

		sql_command, generate_sql_errors := getCreateSQL(options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(sql_command, options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}

		return nil
	}

	exists := func() (*bool, []error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("checking_database_exists", true)

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

		/*
		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}*/

		_, execute_errors := executeUnsafeCommand(&sql_command, options)

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
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

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

		records, execute_errors := executeUnsafeCommand(&sql_command, options)

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
		for _, record := range *(records.GetValues()) {
			record_map, record_map_errors := record.GetMap()
			if record_map_errors != nil {
				errors = append(errors, record_map_errors...)
				continue
			} else if common.IsNil(record_map) {
				errors = append(errors, fmt.Errorf("error: Database: getTableNames(%s) record_map is nil"))
				continue
			}

			table_name, table_name_errors := record_map.GetString(column_name)
			if table_name_errors != nil {
				errors = append(errors, table_name_errors...)
				continue
			} else if common.IsNil(table_name) {
				errors = append(errors, fmt.Errorf("error: Database: getTableNames(%s) was nil available fields are: %s", column_name, record_map.GetKeys()))
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
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("deleting_database", true)
		
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

		/*
		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}*/

		_, execute_errors :=  executeUnsafeCommand(&sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("deleting_database", true)
		
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

		/*
		temp_client, temp_client_errors := getClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}*/

		_, execute_errors :=  executeUnsafeCommand(&sql_command, options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	tableExists :=  func(table_name string) (*bool, []error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", true)
		
		var errors []error
		validate_errors := validate()
		if errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}
		
		sql_command, new_options, sql_command_errors := getCheckTableExistsSQLMySQL(struct_type, table_name, options)
		
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
			return nil, errors
		}
		
		_, execute_errors := executeUnsafeCommand(sql_command, new_options)

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
	
	getOrSetTableSchema := func(table_name string, schema json.Map, mode string) (json.Map, []error) {
		lock_table_schema_cache.Lock()
		defer lock_table_schema_cache.Unlock()
		return table_schema_cache.GetOrSet(*getDatabase(), table_name, schema, mode)
	}

	getTableSchema := func(table_name string) (json.Map, []error) {
		lock_get_table_schema.Lock()
		defer lock_get_table_schema.Unlock()
		var errors []error
	
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)


		cached_schema, cached_schema_errors := getOrSetTableSchema(table_name, json.NewMapValue(), "get")
		if cached_schema_errors != nil {
			if fmt.Sprintf("%s", cached_schema_errors[0]) != "cache is not there" {
				return json.NewMapValue(), cached_schema_errors
			}
		} else if !common.IsNil(cached_schema) {
			return cached_schema, nil
		} 
		
		sql_command, new_options, sql_command_errors := getTableSchemaSQLMySQL(struct_type, table_name, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		}

		json_array, sql_errors := executeUnsafeCommand(sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return  json.NewMapValue(), errors
		}

		temp_schema, schem_errors := mapTableSchemaFromDBMySQL(struct_type, table_name, json_array)
		if schem_errors != nil {
			errors = append(errors, schem_errors...)
			return json.NewMapValue(), errors
		}

		if !temp_schema.HasKey("[no_schema_cache_on_creation]") {
			getOrSetTableSchema(table_name, temp_schema, "set")
		}

		return temp_schema, nil
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

		table_exists, table_exists_errors := tableExists(table_name)
		if table_exists_errors != nil {
			errors = append(errors, table_exists_errors...)
		} else if common.IsNil(table_exists) {
			errors = append(errors, fmt.Errorf("Database.getTable.tableExists returned nil bool"))
		} else if !(*table_exists) {
			return nil, nil
		}

		if len(errors) > 0 {
			return nil, errors
		}

		table_schema, schema_errors := getTableSchema(table_name)

		if schema_errors != nil {
			errors = append(errors, schema_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}		

		get_table, get_table_errors := newTable(*getDatabase(), table_name, table_schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

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

			return tableExists(table_name)
		},
		GetTableInterface: func(table_name string, schema json.Map) (*Table, []error) {
			errors := validate()

			if len(errors) > 0 {
				return nil, errors
			}
			
			table, new_table_errors := newTable(*getDatabase(), table_name, schema, database_reserved_words_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

			if new_table_errors != nil {
				errors = append(errors, new_table_errors...)
			} else if common.IsNil(table) {
				errors = append(errors, fmt.Errorf("table interface is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
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
		ExecuteUnsafeCommand: func(sql_command *string, options *json.Map) (*json.Array, []error) {			
			return executeUnsafeCommand(sql_command, options)
		},
		GetHost: func() (Host, []error) {
			return getHost()
		},
		GetDatabaseName: func() (string, []error) {
			return getDatabaseName()
		},
		SetDatabaseName: func(database_name string) ([]error) {
			return setDatabaseName(database_name)
		},
		SetDatabaseUsername: func(database_username string) []error {
			return setDatabaseUsername(database_username)
		},
		GetDatabaseUsername: func() (string, []error) {
			return getDatabaseUsername()
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
		GetOrSetSchema: func(table_name string, schema json.Map, mode string) (json.Map, []error) {
			// todo clone schema
			return getOrSetTableSchema(table_name, schema, mode)
		},
		GetOrSetAdditonalSchema: func(table_name string, additional_schema *json.Map) (*json.Map, []error) {
			// todo clone schema
			lock_table_additional_schema_cache.Lock()
			defer lock_table_additional_schema_cache.Unlock()
			return table_additional_schema_cache.GetOrSet(*getDatabase(), table_name, additional_schema)
		},
		GetOrSetReadRecords: func(sql string, records *[]Record) (*[]Record, []error) {
			// todo clone schema
			lock_table_records_cache.Lock()
			defer lock_table_records_cache.Unlock()
			return table_read_records_cache.GetOrSetReadRecords(*getDatabase(), sql, records)
		},
		GlobalGeneralLogDisable: func() []error {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			options.SetBoolValue("updating_database_global_settings", true)
			command := "SET GLOBAL general_log = 'OFF';"
			_, command_errors :=  executeUnsafeCommand(&command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalGeneralLogEnable: func() []error {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			options.SetBoolValue("updating_database_global_settings", true)
			command := "SET GLOBAL general_log = 'ON';"
			_, command_errors := executeUnsafeCommand(&command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalSetTimeZoneUTC: func() []error {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			options.SetBoolValue("updating_database_global_settings", true)
			command := "SET GLOBAL time_zone = '+00:00';"
			_, command_errors :=  executeUnsafeCommand(&command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalSetSQLMode: func() []error {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			options.SetBoolValue("updating_database_global_settings", true)
			command := "SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';"
			_, command_errors := executeUnsafeCommand(&command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		DeleteTableByTableNameIfExists: func(table_name string, if_exists bool) []error {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)

			getDatabase().GetOrSetSchema(table_name, json.NewMapValue(), "delete")
			sql_command, new_options, generate_sql_errors := getDropTableSQLMySQL(struct_type, table_name, if_exists, options)

			if generate_sql_errors != nil {
				return generate_sql_errors
			}

			_, execute_sql_command_errors := executeUnsafeCommand(sql_command, new_options)

			if execute_sql_command_errors != nil {
				return execute_sql_command_errors
			}
			return nil
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
