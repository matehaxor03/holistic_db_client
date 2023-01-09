package dao

import (
	"fmt"
	"strings"
	"sync"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type Database struct {
	Validate        func() []error
	Create          func() []error
	Delete          func() []error
	DeleteIfExists  func() []error
	Exists          func() (bool, []error)
	TableExists     func(table_name string) (bool, []error)
	GetDatabaseName func() (string, []error)
	GetDatabaseUsername func() (string, []error)
	SetDatabaseUsername func(database_username string) []error
	SetDatabaseName func(database_name string) []error
	DeleteTableByTableNameIfExists func(table_name string) []error
	GetHost       func() (Host, []error)
	CreateTable     func(table_name string, schema json.Map) (*Table, []error)
	GetTableInterface func(table_name string, schema json.Map) (*Table, []error)
	GetTable        func(table_name string) (*Table, []error)
	GetTables       func() ([]Table, []error)
	GetTableNames   func() ([]string, []error)
	GetTableSchema func(table_name string) (*json.Map, []error)
	GetAdditionalTableSchema func(table_name string) (*json.Map, []error)
	GetOrSetReadRecords func(sql string, records *[]Record) (*[]Record, []error)

	GlobalGeneralLogDisable	func() []error
	GlobalGeneralLogEnable	func() []error
	GlobalSetTimeZoneUTC    func() []error
	GlobalSetSQLMode func() []error
}

func newDatabase(verify *validate.Validator, host Host, database_username string, database_name string, database_create_options *DatabaseCreateOptions) (*Database, []error) {	
	var errors []error

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

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[host]", host)
	map_system_fields.SetObjectForMap("[database_name]", database_name)
	map_system_fields.SetObjectForMap("[database_username]", database_username)
	map_system_fields.SetObjectForMap("[database_create_options]", database_create_options)
	data.SetMapValue("[system_fields]", map_system_fields)

	map_system_schema := json.NewMapValue()
	
	map_host_schema := json.NewMapValue()
	map_host_schema.SetStringValue("type", "dao.Host")
	map_system_schema.SetMapValue("[host]", map_host_schema)

	map_database_name_schema := json.NewMapValue()
	map_database_name_schema.SetStringValue("type", "string")
	map_database_name_schema.SetIntValue("min_length", 2)
	map_database_name_schema.SetBoolValue("not_empty_string_value", true)
	map_database_name_schema_filters := json.NewArrayValue()
	map_database_name_schema_filter := json.NewMapValue()
	map_database_name_schema_filter.SetObjectForMap("function",  verify.GetValidateDatabaseNameFunc())
	map_database_name_schema_filters.AppendMapValue(map_database_name_schema_filter)
	map_database_name_schema.SetArrayValue("filters", map_database_name_schema_filters)
	map_system_schema.SetMapValue("[database_name]", map_database_name_schema)

	map_database_username := json.NewMapValue()
	map_database_username.SetStringValue("type", "string")
	array_database_username_filters := json.NewArrayValue()
	map_database_username_filter := json.NewMapValue()
	map_database_username_filter.SetObjectForMap("function",  verify.GetValidateUsernameFunc())
	array_database_username_filters.AppendMapValue(map_database_username_filter)
	map_database_username.SetArrayValue("filters", array_database_username_filters)
	map_system_schema.SetMapValue("[database_username]", map_database_username)

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
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("database name is nil"))
		}

		if len(errors) > 0 { 
			return "", errors
		}
		return temp_value.(string), nil
	}

	setDatabaseName := func(new_database_name string) []error {
		database_name_errors := verify.ValidateDatabaseName(new_database_name) 
		if database_name_errors != nil {
			return database_name_errors
		}
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_name]", new_database_name)
	}

	getDatabaseUsername := func() (*string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", "*string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} 

		if len(errors) > 0 { 
			return nil, errors
		}

		return temp_value.(*string), nil
	}
	
	setDatabaseUsername := func(new_database_username string) []error {
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", new_database_username)
	}

	executeUnsafeCommand := func(sql_command *string, options *json.Map) (*json.Array, []error) {
		errors := validate()
		if errors != nil {
			return nil, errors
		}
		
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(*getDatabase(), sql_command, options)
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

	getOrSetAdditionalTableSchema := func(table_name string, additional_schema *json.Map) (*json.Map, []error) {
		// todo clone schema
		lock_table_additional_schema_cache.Lock()
		defer lock_table_additional_schema_cache.Unlock()
		return table_additional_schema_cache.GetOrSet(*getDatabase(), table_name, additional_schema)
	}

	create := func() []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("creating_database", true)
		options.SetBoolValue("read_no_records", true)

		databaseCreateOptions, databaseCreateOptions_errors := getDatabaseCreateOptions()
		if databaseCreateOptions_errors != nil {
			return databaseCreateOptions_errors
		}

		var collate_value *string = nil
		var character_set_value *string = nil
		if databaseCreateOptions != nil {
			character_set, character_set_errors := databaseCreateOptions.GetCharacterSet()
			if character_set_errors != nil {
				return character_set_errors
			} else if !common.IsNil(character_set) {
				character_set_value = character_set
			}

			collate, collate_errors := databaseCreateOptions.GetCollate()
			if collate_errors != nil {
				return collate_errors
			} else if !common.IsNil(collate) {
				collate_value = collate
			}
		}

		sql_command, new_options, generate_sql_errors :=  sql_generator_mysql.GetCreateDatabaseSQL(verify, database_name, character_set_value, collate_value, options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(sql_command, new_options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}

		return nil
	}

	exists := func() (*bool, []error) {
		var errors []error
		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetDatabaseExistsSQL(verify, temp_database_name, nil)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		_, execute_errors := executeUnsafeCommand(sql_command, new_options)

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
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetTableNamesSQL(verify, temp_database_name, options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		records, execute_errors := executeUnsafeCommand(sql_command, new_options)

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
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("deleting_database", true)
		options.SetBoolValue("read_no_records", true)

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetDropDatabaseSQL(verify, temp_database_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors :=  executeUnsafeCommand(sql_command, new_options)

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
		options.SetBoolValue("read_no_records", true)

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return temp_database_name_errors
		}

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetDropDatabaseIfExistsSQL(verify, temp_database_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors :=  executeUnsafeCommand(sql_command, new_options)

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
		options.SetBoolValue("read_no_records", true)
		
		var errors []error
		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetCheckTableExistsSQL(verify, struct_type, table_name, options)
		
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
			return nil, errors
		}
		
		_, execute_errors := executeUnsafeCommand(sql_command, new_options)
		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		exists := false
		if errors == nil {
			exists = true
			return &exists, nil
		}

		error_string := fmt.Sprintf("%s", errors)
		if strings.Contains(error_string, "doesn't exist") {
			exists = false
			return &exists, nil
		}

		return nil, errors
	}
	
	getOrSetTableSchema := func(table_name string, schema *json.Map, mode string) (*json.Map, []error) {
		return table_schema_cache.GetOrSet(*getDatabase(), table_name, schema, mode)
	}

	getTableSchema := func(table_name string) (*json.Map, []error) {
		var errors []error
	
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)


		cached_schema, cached_schema_errors := getOrSetTableSchema(table_name, nil, "get")
		if cached_schema_errors != nil {
			if fmt.Sprintf("%s", cached_schema_errors[0]) != "cache is not there" {
				return nil, cached_schema_errors
			}
		} else if !common.IsNil(cached_schema) {
			return cached_schema, nil
		} 
		
		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetTableSchemaSQL(verify, struct_type, table_name, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		}

		json_array, sql_errors := executeUnsafeCommand(sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return  nil, errors
		}

		temp_schema, schem_errors := sql_generator_mysql.MapTableSchemaFromDB(verify, struct_type, table_name, json_array)
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
		} else if common.IsNil(table_schema) {
			errors = append(errors, fmt.Errorf("table schema is nil"))
		}

		if len(errors) > 0 {
			return nil, errors
		}		

		get_table, get_table_errors := newTable(verify, *getDatabase(), table_name, nil, table_schema)

		if get_table_errors != nil {
			return nil, get_table_errors
		}

		return get_table, nil
	}

	deleteTableByTableNameIfExists := func(table_name string) []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		getOrSetTableSchema(table_name, nil, "delete")
		sql_command, new_options, generate_sql_errors := sql_generator_mysql.GetDropTableSQL(verify, struct_type, table_name, true, options)

		if generate_sql_errors != nil {
			return generate_sql_errors
		}

		_, execute_sql_command_errors := executeUnsafeCommand(sql_command, new_options)

		if execute_sql_command_errors != nil {
			return execute_sql_command_errors
		}
		return nil
	}

	getTableInterface := func(table_name string, user_defined_schema json.Map) (*Table, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}
		
		table, new_table_errors := newTable(verify, *getDatabase(), table_name, &user_defined_schema, nil)

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
		validate_errors := validate()
		
		if validate_errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)

		cached_additonal_schema, cached_additonal_schema_errors := getOrSetAdditionalTableSchema(table_name, nil)
		if cached_additonal_schema_errors != nil {
			return nil, cached_additonal_schema_errors
		} else if !common.IsNil(cached_additonal_schema) {
			return cached_additonal_schema, nil
		}

		temp_database_name, temp_database_name_errors := getDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}
		
		sql_command, new_options,  sql_command_errors := sql_generator_mysql.GetTableSchemaAdditionalSQL(verify, struct_type, temp_database_name, table_name, options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		json_array, sql_errors := executeUnsafeCommand(sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return nil, errors
		}

		additional_schema, additional_schema_errors := sql_generator_mysql.MapAdditionalSchemaFromDBToMap(json_array)
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
			temp_table_exists, temp_table_exists_errors := tableExists(table_name)
			if temp_table_exists_errors != nil {
				return false, temp_table_exists_errors
			} else if common.IsNil(temp_table_exists) {
				var errors []error
				errors = append(errors, fmt.Errorf("table_exists result is nil"))
				return false, errors
			}
			return *temp_table_exists, nil
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
			database_username_temp, database_username_temp_errors := getDatabaseUsername()
			if database_username_temp_errors != nil {
				return "", database_username_temp_errors
			} else if common.IsNil(database_username_temp) {
				return "", nil
			}
			return *database_username_temp, nil
		},
		GetTableSchema: func(table_name string) (*json.Map, []error) {
			return getTableSchema(table_name)
		},
		GetAdditionalTableSchema: func(table_name string) (*json.Map, []error) {
			// todo clone schema
			return getAdditionalTableSchema(table_name)
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
		DeleteTableByTableNameIfExists: func(table_name string) []error {
			return deleteTableByTableNameIfExists(table_name)
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
