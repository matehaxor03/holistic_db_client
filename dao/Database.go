package dao

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type Database struct {
	Validate        func() []error
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (bool, []error)
	TableExists     func(table_name string) (bool, []error)
	GetDatabaseName func() (string)
	GetDatabaseUsername func() (*string)
	SetDatabaseUsername func(database_username string) []error
	SetDatabaseName func(database_name string) []error
	DeleteTableByTableNameIfExists func(table_name string) []error
	DeleteTableByTableName func(table_name string) []error
	GetHost       func() (Host)
	CreateTable     func(table_name string, schema json.Map) (*Table, []error)
	GetTableInterface func(table_name string, schema json.Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() ([]Table, []error)
	GetTableNames   func() ([]string, []error)
	GetTableSchema func(table_name string) (*json.Map, []error)
	GetAdditionalTableSchema func(table_name string) (*json.Map, []error)
	GetClient func() (Client)

	GlobalGeneralLogDisable	func() []error
	GlobalGeneralLogEnable	func() []error
	GlobalSetTimeZoneUTC    func() []error
	GlobalSetSQLMode func() []error
}

func newDatabase(verify *validate.Validator, client Client, host Host, database_username *string, database_name string, database_create_options *DatabaseCreateOptions) (*Database, []error) {	
	var errors []error
	var this_database Database
	table_schema_cache := newTableSchemaCache()
	table_additional_schema_cache := newTableAdditionalSchemaCache()
	table_exists_cache := newTableExistsCache()
	table_cache := newTableCache()
	mysql_wrapper := sql_generator_mysql.NewMySQL()
	
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	setDatabase := func(database Database) {
		this_database = database
	}

	getDatabase := func() Database {
		return this_database
	}

	getClient := func() Client {
		return client
	}

	validate := func() []error {
		var errors []error
		if host_errors := host.Validate(); host_errors != nil {
			errors = append(errors, host_errors...)
		}
		
		if database_username != nil {
			if database_username_errors := verify.ValidateUsername(*database_username); database_username_errors != nil {
				errors = append(errors, database_username_errors...)
			}
		}
		
		if database_name_errors := verify.ValidateDatabaseName(database_name); database_name_errors != nil {
			errors = append(errors, database_name_errors...)
		}

		if database_create_options != nil {
			if database_create_options_errors := database_create_options.Validate(); database_create_options_errors != nil {
				errors = append(errors, database_create_options_errors...)
			}
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getDatabaseCreateOptions := func() (*DatabaseCreateOptions) {
		return database_create_options
	}

	getHost := func() (Host) {
		return host
	}

	getDatabaseName := func() (string) {
		return database_name
	}

	setDatabaseName := func(new_database_name string) []error {
		database_name_errors := verify.ValidateDatabaseName(new_database_name) 
		if database_name_errors != nil {
			return database_name_errors
		}
		database_name = new_database_name
		return nil
	}

	getDatabaseUsername := func() (*string) {
		return database_username
	}
	
	setDatabaseUsername := func(new_database_username string) []error {
		if database_username_errors := verify.ValidateUsername(new_database_username); database_username_errors != nil {
			return database_username_errors
		}

		database_username = &new_database_username
		return nil
	}

	executeUnsafeCommand := func(sql_command strings.Builder, options json.Map) (json.Array, []error) {
		var errors []error
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase(), sql_command, options)
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

	getOrSetAdditionalTableSchema := func(table_name string, additional_schema *json.Map) (*json.Map, []error) {
		// todo clone schema
		return table_additional_schema_cache.GetOrSet(getDatabase(), table_name, additional_schema)
	}

	create := func() []error {
		options := json.NewMapValue()
		options.SetBoolValue("creating_database", true)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)
		
		databaseCreateOptions := getDatabaseCreateOptions()

		var collate *string = nil
		var character_set *string = nil
		if databaseCreateOptions != nil {
			character_set = databaseCreateOptions.GetCharacterSet()
			collate = databaseCreateOptions.GetCollate()
		}

		sql_command, new_options, generate_sql_errors :=  mysql_wrapper.GetCreateDatabaseSQL(verify, database_name, character_set, collate, options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}

		return nil
	}

	exists := func() (*bool, []error) {
		var errors []error
		options := json.NewMapValue()
		options.SetBoolValue("creating_database", true)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		sql_command, new_options, sql_command_errors := mysql_wrapper.GetDatabaseExistsSQL(verify, getDatabaseName(), options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		_, execute_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		exists := true
		if errors == nil {
			exists = true
			return &exists, nil
		}
	
		error_string := fmt.Sprintf("%s", errors)
		if strings.Contains(error_string, "Unknown database") {
			exists = false
			return &exists, nil
		}

		return nil, errors
	}

	getTableNames := func() (*[]string, []error) {
		var errors []error
		options := json.NewMapValue()
		options.SetBoolValue("creating_database", true)
		options.SetBoolValue("get_last_insert_id", false)

		sql_command, new_options, sql_command_errors := mysql_wrapper.GetTableNamesSQL(verify, database_name, options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		records, execute_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		} else if common.IsNil(records) {
			errors = append(errors, fmt.Errorf("error: Database: getTableNames returned nil records"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
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

			table_name, table_name_errors := record_map.GetStringValue(column_name)
			if table_name_errors != nil {
				errors = append(errors, table_name_errors...)
				continue
			}
			table_names = append(table_names, table_name)
		}
		return &table_names, nil
	}

	delete := func() ([]error) {
		options := json.NewMapValue()
		options.SetBoolValue("deleting_database", true)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)


		sql_command, new_options, sql_command_errors := mysql_wrapper.GetDropDatabaseSQL(verify, database_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors :=  executeUnsafeCommand(*sql_command, new_options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		options := json.NewMapValue()
		options.SetBoolValue("deleting_database", true)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)


		sql_command, new_options, sql_command_errors := mysql_wrapper.GetDropDatabaseIfExistsSQL(verify, database_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors :=  executeUnsafeCommand(*sql_command, new_options)

		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	tableExists :=  func(table_name string) (bool, []error) {
		exists_cache, exists_cache_errors := table_exists_cache.GetOrSet(getDatabase(), table_name, "get")
		if exists_cache_errors != nil {
			return false, exists_cache_errors
		} else if exists_cache {
			return true, nil
		}

		options := json.NewMapValue()
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		var errors []error
		sql_command, new_options, sql_command_errors := mysql_wrapper.GetTableExistsSQL(verify, table_name, options)
		
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
			return false, errors
		}
		
		_, execute_errors := executeUnsafeCommand(*sql_command, new_options)
		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if errors == nil {
			table_exists_cache.GetOrSet(getDatabase(), table_name, "set")
			return true, nil
		}

		error_string := fmt.Sprintf("%s", errors)
		if strings.Contains(error_string, "doesn't exist") {
			table_exists_cache.GetOrSet(getDatabase(), table_name, "delete")
			return false, nil
		}

		return false, errors
	}
	
	getOrSetTableSchema := func(table_name string, schema *json.Map, mode string) (*json.Map, []error) {
		return table_schema_cache.GetOrSet(getDatabase(), table_name, schema, mode)
	}

	getTableSchema := func(table_name string) (*json.Map, []error) {
		var errors []error
	
		options := json.NewMapValue()
		options.SetBoolValue("get_last_insert_id", false)
		options.SetBoolValue("read_no_records", false)

		cached_schema, cached_schema_errors := getOrSetTableSchema(table_name, nil, "get")
		if cached_schema_errors != nil {
			if fmt.Sprintf("%s", cached_schema_errors[0]) != "cache is not there" {
				return nil, cached_schema_errors
			}
		} else if !common.IsNil(cached_schema) {
			return cached_schema, nil
		} 
		
		sql_command, new_options, sql_command_errors := mysql_wrapper.GetTableSchemaSQL(verify, table_name, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		}

		json_array, sql_errors := executeUnsafeCommand(*sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return  nil, errors
		}

		temp_schema, schem_errors := mysql_wrapper.MapTableSchemaFromDB(verify, table_name, json_array)
		if schem_errors != nil {
			errors = append(errors, schem_errors...)
			return nil, errors
		}

		if !temp_schema.HasKey("[no_schema_cache_on_creation]") {
			getOrSetTableSchema(table_name, temp_schema, "set")
		}

		return temp_schema, nil
	}

	getTable := func(table_name string) (*Table, []error) {
		temp_table, temp_table_errors := table_cache.GetOrSet(getDatabase(), table_name, nil, "get")
		if temp_table_errors != nil {
			return nil, temp_table_errors
		} else if !common.IsNil(temp_table) {
			return temp_table, nil
		}

		table_exists, table_exists_errors := tableExists(table_name)
		if table_exists_errors != nil {
			errors = append(errors, table_exists_errors...)
		} else if !table_exists {
			return nil, nil
		}

		if len(errors) > 0 {
			return nil, errors
		}

		table_schema, schema_errors := getTableSchema(table_name)

		if schema_errors != nil {
			errors = append(errors, schema_errors...)
		} else if common.IsNil(table_schema) {
			errors = append(errors, fmt.Errorf("table schema is nil"))
		}

		if len(errors) > 0 {
			return nil, errors
		}		

		get_table, get_table_errors := newTable(verify, getDatabase(), table_name, nil, table_schema)

		if get_table_errors != nil {
			return nil, get_table_errors
		}
		table_cache.GetOrSet(getDatabase(), table_name, get_table, "set")
		return get_table, nil
	}

	deleteTableByTableNameIfExists := func(table_name string) []error {
		options := json.NewMapValue()
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		getOrSetTableSchema(table_name, nil, "delete")
		table_exists_cache.GetOrSet(getDatabase(), table_name, "delete")
		table_cache.GetOrSet(getDatabase(), table_name, nil, "delete")
		sql_command, new_options, generate_sql_errors := mysql_wrapper.GetDropTableIfExistsSQL(verify, table_name, options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}
		return nil
	}

	deleteTableByTableName := func(table_name string) []error {
		options := json.NewMapValue()
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		getOrSetTableSchema(table_name, nil, "delete")
		table_exists_cache.GetOrSet(getDatabase(), table_name, "delete")
		table_cache.GetOrSet(getDatabase(), table_name, nil, "delete")
		sql_command, new_options, generate_sql_errors := mysql_wrapper.GetDropTableSQL(verify, table_name, options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}
		return nil
	}


	getTableInterface := func(table_name string, user_defined_schema json.Map) (*Table, []error) {
		var errors []error
		table, new_table_errors := newTable(verify, getDatabase(), table_name, &user_defined_schema, nil)

		if new_table_errors != nil {
			errors = append(errors, new_table_errors...)
		} else if common.IsNil(table) {
			errors = append(errors, fmt.Errorf("table interface is nil"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return table, nil
	}

	getAdditionalTableSchema := func(table_name string) (*json.Map, []error) {
		var errors []error
		options := json.NewMapValue()
		options.SetBoolValue("get_last_insert_id", false)
		options.SetBoolValue("read_no_records", false)

		cached_additonal_schema, cached_additonal_schema_errors := getOrSetAdditionalTableSchema(table_name, nil)
		if cached_additonal_schema_errors != nil {
			return nil, cached_additonal_schema_errors
		} else if !common.IsNil(cached_additonal_schema) {
			return cached_additonal_schema, nil
		}
		
		sql_command, new_options,  sql_command_errors := mysql_wrapper.GetTableSchemaAdditionalSQL(verify, database_name, table_name, options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		json_array, sql_errors := executeUnsafeCommand(*sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return nil, errors
		}

		additional_schema, additional_schema_errors := mysql_wrapper.MapAdditionalSchemaFromDBToMap(json_array)
		if additional_schema_errors != nil {
			errors = append(errors, additional_schema_errors...)
		} else if common.IsNil(additional_schema) {
			errors = append(errors, fmt.Errorf("additional schema is nil"))
		}

		if len(errors) > 0 {
			return nil , errors
		}


		getOrSetAdditionalTableSchema(table_name, additional_schema)
		return additional_schema, nil
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
		Exists: func() (bool, []error) {
			temp_exists, temp_exists_errors := exists()
			if temp_exists_errors != nil {
				return false, temp_exists_errors
			}
			return *temp_exists, nil
		},
		DeleteIfExists: func() []error {
			return deleteIfExists()
		},
		TableExists: func(table_name string) (bool, []error) {
			return tableExists(table_name)
		},
		GetTableInterface: func(table_name string, user_defined_schema json.Map) (*Table, []error) {
			return getTableInterface(table_name, user_defined_schema)
		},
		CreateTable: func(table_name string, user_defined_schema json.Map) (*Table, []error) {
			table, new_table_errors := getTableInterface(table_name, user_defined_schema)

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
			return getTable(table_name)
		},
		GetTableNames: func()  ([]string, []error) {
			temp_table_names, temp_table_names_errors := getTableNames()
			if temp_table_names_errors != nil {
				var empty_strings []string
				return empty_strings, temp_table_names_errors
			} else if common.IsNil(temp_table_names) {
				var errors []error
				var empty_strings []string
				errors = append(errors, fmt.Errorf("table_names is nil"))
				return empty_strings, errors
			}
			return *temp_table_names, nil
		},
		GetTables: func() ([]Table, []error) {
			var errors []error
			var empty_tables []Table
			database := getDatabase()
			
			database_errors := database.Validate()
			if database_errors != nil {
				errors = append(errors, database_errors...)
			}
			

			if len(errors) > 0 {
				return empty_tables, errors
			}

			var tables []Table
			table_names, table_names_errors := getTableNames()
			if table_names_errors != nil {
				errors = append(errors, table_names_errors...)
			}

			if len(errors) > 0 {
				return empty_tables, errors
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
				return empty_tables, errors
			}

			return tables, nil
		},
		GetHost: func() (Host) {
			return getHost()
		},
		GetDatabaseName: func() (string) {
			return getDatabaseName()
		},
		SetDatabaseName: func(database_name string) ([]error) {
			return setDatabaseName(database_name)
		},
		SetDatabaseUsername: func(database_username string) []error {
			return setDatabaseUsername(database_username)
		},
		GetDatabaseUsername: func() (*string) {
			return getDatabaseUsername()
		},
		GetTableSchema: func(table_name string) (*json.Map, []error) {
			return getTableSchema(table_name)
		},
		GetAdditionalTableSchema: func(table_name string) (*json.Map, []error) {
			// todo clone schema
			return getAdditionalTableSchema(table_name)
		},
		GetClient: func() Client {
			return getClient()
		},
		GlobalGeneralLogDisable: func() []error {
			options := json.NewMapValue()
			options.SetBoolValue("read_no_records", true)
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("updating_database_global_settings", true)
			var command strings.Builder
			command.WriteString("SET GLOBAL general_log = 'OFF';")
			_, command_errors :=  executeUnsafeCommand(command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalGeneralLogEnable: func() []error {
			options := json.NewMapValue()
			options.SetBoolValue("read_no_records", true)
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("updating_database_global_settings", true)
			var command strings.Builder
			command.WriteString("SET GLOBAL general_log = 'ON';")
			_, command_errors := executeUnsafeCommand(command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalSetTimeZoneUTC: func() []error {
			options := json.NewMapValue()
			options.SetBoolValue("read_no_records", true)
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("updating_database_global_settings", true)
			var command strings.Builder
			command.WriteString("SET GLOBAL time_zone = '+00:00';")
			_, command_errors :=  executeUnsafeCommand(command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		GlobalSetSQLMode: func() []error {
			options := json.NewMapValue()
			options.SetBoolValue("read_no_records", true)
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("updating_database_global_settings", true)
			var command strings.Builder
			command.WriteString("SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';")
			_, command_errors := executeUnsafeCommand(command, options)
			if command_errors != nil {
				return command_errors
			}
			return nil
		},
		DeleteTableByTableNameIfExists: func(table_name string) []error {
			return deleteTableByTableNameIfExists(table_name)
		},
		DeleteTableByTableName: func(table_name string) []error {
			return deleteTableByTableName(table_name)
		},
	}
	setDatabase(x)

	valiation_errors := validate()
	if valiation_errors != nil {
		errors = append(errors, valiation_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
