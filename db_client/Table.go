package db_client

import (
	"fmt"
	"strconv"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type Table struct {
	Validate              func() []error
	Exists                func() (*bool, []error)
	Create                func() []error
	Read                func() []error
	Delete                func() []error
	DeleteIfExists        func() []error
	GetSchema             func() (*json.Map, []error)
	GetAdditionalSchema   func() (*json.Map, []error)
	GetTableName          func() (string, []error)
	SetTableName          func(table_name string) []error
	GetSchemaColumns      func() (*[]string, []error)
	GetTableColumns       func() (*[]string, []error)
	GetIdentityColumns    func() (*[]string, []error)
	GetPrimaryKeyColumns  func() (*[]string, []error)
	GetForeignKeyColumns  func() (*[]string, []error)

	GetNonPrimaryKeyColumns func() (*[]string, []error)
	Count                 func(filter *json.Map, filter_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*uint64, []error)
	CreateRecord          func(record json.Map) (*Record, []error)
	CreateRecords          func(records json.Array) ([]error)
	UpdateRecords          func(records json.Array) ([]error)
	UpdateRecord          func(record json.Map) ([]error)
	ReadRecords         func(select_fields *json.Array, filter *json.Map, filter_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase           func() (*Database, []error)
	ToJSONString          func(json *strings.Builder) ([]error)
}

func newTable(database Database, table_name string, schema *json.Map, database_reserved_words_obj *DatabaseReservedWords, table_name_whitelist_characters_obj *TableNameCharacterWhitelist, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Table, []error) {
	struct_type := "*Table"

	var errors []error
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	var this_table *Table

	setTable := func(table *Table) {
		this_table = table
	}

	getTable := func() *Table {
		return this_table
	}

	//database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	table_name_whitelist_characters := table_name_whitelist_characters_obj.GetTableNameCharacterWhitelist()
	column_name_whitelist_characters := column_name_whitelist_characters_obj.GetColumnNameCharacterWhitelist()
	
	
	setupData := func(b Database, n string, schema_from_db *json.Map) (json.Map) {
		schema_is_nil := false
		
		merged_schema := json.NewMapValue()
		if common.IsNil(schema_from_db) {
			schema_is_nil = true
			schema_from_db = json.NewMap()
		}

		d := json.NewMapValue()
		d.SetMapValue("[fields]", json.NewMapValue())
		d.SetMapValue("[schema]", json.NewMapValue())

		map_system_fields := json.NewMapValue()
		map_system_fields.SetObjectForMap("[database]", b)
		map_system_fields.SetObjectForMap("[table_name]", n)
		d.SetMapValue("[system_fields]", map_system_fields)

		///

		map_system_schema := json.NewMapValue()

		// Start database
		map_database_schema := json.NewMapValue()
		map_database_schema.SetStringValue("type", "db_client.Database")
		map_system_schema.SetMapValue("[database]", map_database_schema)
		// End database

		// Start table_name
		map_table_name_schema := json.NewMapValue()
		map_table_name_schema.SetStringValue("type", "string")
		map_table_name_schema.SetBoolValue("not_empty_string_value", true)
		map_table_name_schema.SetIntValue("min_length", 2)

		map_table_name_schema_filters := json.NewArrayValue()
		map_table_name_schema_filter := json.NewMapValue()
		map_table_name_schema_filter.SetObjectForMap("values", table_name_whitelist_characters)
		map_table_name_schema_filter.SetObjectForMap("function",  getWhitelistCharactersFunc())
		map_table_name_schema_filters.AppendMapValue(map_table_name_schema_filter)
		map_table_name_schema.SetArrayValue("filters", map_table_name_schema_filters)
		map_system_schema.SetMapValue("[table_name]", map_table_name_schema)
		// End table_name


		d.SetMapValue("[system_schema]", map_system_schema)

		// Start active
		map_active_schema := json.NewMapValue()
		map_active_schema.SetStringValue("type", "bool")
		map_active_schema.SetBoolValue("default", true)
		merged_schema.SetMapValue("active",map_active_schema)
		// End active


		// Start archieved
		map_archieved_schema := json.NewMapValue()
		map_archieved_schema.SetStringValue("type", "bool")
		map_archieved_schema.SetBoolValue("default", false)
		merged_schema.SetMapValue("archieved", map_archieved_schema)
		// End archieved

		// Start created_date
		map_created_date_schema := json.NewMapValue()
		map_created_date_schema.SetStringValue("type", "time.Time")
		map_created_date_schema.SetStringValue("default", "now")
		map_created_date_schema.SetUInt8Value("decimal_places", uint8(6))
		merged_schema.SetMapValue("created_date", map_created_date_schema)
		// End created_date

		// Start last_modified_date
		map_last_modified_date_schema := json.NewMapValue()
		map_last_modified_date_schema.SetStringValue("type", "time.Time")
		map_last_modified_date_schema.SetStringValue("default", "now")
		map_last_modified_date_schema.SetUInt8Value("decimal_places", uint8(6))
		merged_schema.SetMapValue("last_modified_date", map_last_modified_date_schema)
		// End last_modified_date


		// Start archieved_date
		map_archieved_date_date_schema := json.NewMapValue()
		map_archieved_date_date_schema.SetStringValue("type", "time.Time")
		map_archieved_date_date_schema.SetStringValue("default", "zero")
		map_archieved_date_date_schema.SetUInt8Value("decimal_places", uint8(6))
		merged_schema.SetMapValue("archieved_date", map_archieved_date_date_schema)
		// End archieved_date

	
		if schema_is_nil {
			d.SetBoolValue("[schema_is_nil]", true)
		} else {
			d.SetBoolValue("[schema_is_nil]", false)
		}
	
		for _, schema_key_from_db := range schema_from_db.GetKeys() {
			current_schema_from_db, current_schema_error_from_db := schema_from_db.GetMap(schema_key_from_db)
			if current_schema_error_from_db != nil {
				errors = append(errors, current_schema_error_from_db...)
			} else if common.IsNil(current_schema_from_db) {
				errors = append(errors, fmt.Errorf("schema is nil for key from db %s", schema_key_from_db))
			} else {
				if !merged_schema.HasKey(schema_key_from_db) {
					merged_schema.SetMapValue(schema_key_from_db, *current_schema_from_db)
				} else if current_schema_from_db.IsArray("filters") {
					filters_from_db, filters_from_db_errors := current_schema_from_db.GetArray("filters")
					if filters_from_db_errors != nil {
						errors = append(errors, filters_from_db_errors...)
					} else if common.IsNil(filters_from_db) {
						errors = append(errors, fmt.Errorf("filters from db is nil"))
					} else if merged_schema.IsMap(schema_key_from_db) {
						merged_schema_map, merged_schema_map_errors := merged_schema.GetMap(schema_key_from_db)
						if merged_schema_map_errors != nil {
							errors = append(errors, merged_schema_map_errors...)
						} else {
							filters_array, filters_array_errors := merged_schema_map.GetArray("filters")
							if filters_array_errors != nil {
								errors = append(errors, filters_array_errors...)
							} else {
								if common.IsNil(filters_array) {
									new_filters_array := json.NewArrayValue()
									merged_schema_map.SetArray("filters", &new_filters_array)
									filters_array = &new_filters_array
								}
								for _, filter_from_db := range *(filters_from_db.GetValues()) {
									filters_array.AppendValue(filter_from_db)
								}
							}
						}
					} 
				}
			}
		}
		
		d.SetMapValue("[schema]", merged_schema)

		return d
	}

	data := setupData(database, table_name, schema)

	getData := func() (*json.Map) {
		return &data
	}

	getTableName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[table_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} else if temp_value == nil {
			return "", nil
		}
		
		return temp_value.(string), temp_value_errors
	}

	getTableColumns := func() (*[]string, []error) {
		temp_schemas, temp_schemas_error := GetSchemas(struct_type, getData(), "[schema]")
		if temp_schemas_error != nil {
			return nil, temp_schemas_error
		}
		columns := temp_schemas.GetKeys()
		return &columns, nil
	}

	getPrimaryKeyColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.GetKeys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", struct_type, column))
				continue
			}

			if column_schema.IsBoolTrue("primary_key") {
				columns = append(columns, column)
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &columns, nil
	}

	getForeignKeyColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.GetKeys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", struct_type, column))
				continue
			}

			if column_schema.IsBoolTrue("foreign_key") {
				columns = append(columns, column)
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &columns, nil
	}

	getIdentityColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.GetKeys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", struct_type, column))
				continue
			}

			if column_schema.IsBoolTrue("primary_key") || column_schema.IsBoolTrue("foreign_key") {
				columns = append(columns, column)
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return &columns, nil
	}

	getNonPrimaryKeyColumns := func() (*[]string, []error) {
		var errors []error
		var columns []string

		schema_map, schema_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schema_map_errors != nil {
			errors = append(errors, schema_map_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		for _, column := range schema_map.GetKeys() {
			column_schema, column_schema_errors := schema_map.GetMap(column)
			if column_schema_errors != nil {
				errors = append(errors, column_schema_errors...)
				continue
			} else if column_schema == nil {
				errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", struct_type, column))
				continue
			}

			if !(column_schema.IsBoolTrue("primary_key")) {
				columns = append(columns, column)
			}
		}
		return &columns, nil
	}

	validate := func() []error {
		return ValidateData(getData(), "*db_client.Table")
	}

	getDatabase := func() (*Database, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database]", "*db_client.Database")
		return temp_value.(*Database), temp_value_errors
	}

	exists := func() (*bool, []error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", true)
		
		var errors []error
		validate_errors := validate()
		if errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return nil, errors
		}
		
		sql_command := fmt.Sprintf("SELECT 0 FROM ")
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`", table_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`", table_name_escaped)
		}
		sql_command += " LIMIT 1;"
		
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}
		
		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, options)

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

	delete := func() ([]error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		drop_table_if_exists := false
		sql_command, new_options, sql_command_errors := getDropTableSQLMySQL(struct_type, getTable(), &drop_table_if_exists, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}
		
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		drop_table_if_exists := true
		sql_command, new_options, sql_command_errors := getDropTableSQLMySQL(struct_type, getTable(), &drop_table_if_exists, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}
	
	updateRecords := func(records json.Array) []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)

		errors := validate()
		if errors != nil {
			return errors
		}

		if len(*(records.GetValues())) == 0 {
			return nil
		}

		for _, record := range *(records.GetValues()) {
			if !record.IsMap() {
				errors = append(errors, fmt.Errorf("record is not a map"))
			}
		}

		if len(errors) > 0 {
			return errors
		}

		var records_obj []Record
		for _, record := range *(records.GetValues()) {
			current_map, current_map_errors := record.GetMap()
			if current_map_errors != nil {
				errors = append(errors, current_map_errors...)
			} else if common.IsNil(current_map) {
				errors = append(errors, fmt.Errorf("record is nil"))
			}
			
			if len(errors) > 0 {
				continue
			}

			record_obj, record_errors := newRecord(*getTable(), *current_map, database_reserved_words_obj,  column_name_whitelist_characters_obj)
			if record_errors != nil {
				errors = append(errors, record_errors...)
			} else if common.IsNil(record_obj) {
				errors = append(errors, fmt.Errorf("record_obj is nil"))
			} else {
				records_obj = append(records_obj, *record_obj)
			}
		}

		if len(errors) > 0 {
			return errors
		}

		sql := ""
		for _, record_obj := range records_obj {
			sql_update_snippet, sql_update_snippet_errors := record_obj.GetUpdateSQL()
			if sql_update_snippet_errors != nil {
				errors = append(errors, sql_update_snippet_errors...)
			} else {
				sql += *sql_update_snippet + "\n"
			}
		}

		if len(errors) > 0 {
			return errors
		}


		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	updateRecord := func(record json.Map) []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)

		errors := validate()
		if errors != nil {
			return errors
		}

		record_obj, record_errors := newRecord(*getTable(), record, database_reserved_words_obj,  column_name_whitelist_characters_obj)
		if record_errors != nil {
			return record_errors
		} else if common.IsNil(record_obj) {
			errors = append(errors, fmt.Errorf("newRecord is nil"))
			return errors
		}

		sql, update_sql_errors := record_obj.GetUpdateSQL()
		if update_sql_errors != nil {
			return update_sql_errors
		} else if common.IsNil(sql) {
			errors = append(errors, fmt.Errorf("generated sql is nil"))
			return errors
		}
		
		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	createRecords := func(records json.Array) []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)

		errors := validate()
		if errors != nil {
			return errors
		}

		if len(*(records.GetValues())) == 0 {
			return nil
		}

		for _, record := range *(records.GetValues()) {
			if !record.IsMap() {
				errors = append(errors, fmt.Errorf("record is not a map"))
			}
		}

		if len(errors) > 0 {
			return errors
		}

		var records_obj []Record
		for _, record := range *(records.GetValues()) {
			current_map, current_map_errors := record.GetMap()
			if current_map_errors != nil {
				errors = append(errors, current_map_errors...)
			} else if common.IsNil(current_map) {
				errors = append(errors, fmt.Errorf("record is nil"))
			}
			
			if len(errors) > 0 {
				continue
			}

			record_obj, record_errors := newRecord(*getTable(), *current_map, database_reserved_words_obj,  column_name_whitelist_characters_obj)
			if record_errors != nil {
				errors = append(errors, record_errors...)
			} else if common.IsNil(record_obj) {
				errors = append(errors, fmt.Errorf("record_obj is nil"))
			} else {
				records_obj = append(records_obj, *record_obj)
			}
		}

		if len(errors) > 0 {
			return errors
		}

		sql := ""
		for _, record_obj := range records_obj {
			sql_update_snippet, _, sql_update_snippet_errors := record_obj.GetCreateSQL()
			if sql_update_snippet_errors != nil {
				errors = append(errors, sql_update_snippet_errors...)
			} else {
				sql += *sql_update_snippet + "\n"
			}
		}

		if len(errors) > 0 {
			return errors
		}


		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getAdditionalSchema := func() (*json.Map, []error) {
		var errors []error
		validate_errors := validate()
		
		if validate_errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		temp_client_manager, temp_client_manager_errors := temp_client.GetClientManager()
		if temp_client_manager_errors != nil {
			return nil, temp_client_manager_errors
		}

		cached_additonal_schema, cached_additonal_schema_errors := temp_client_manager.GetOrSetAdditonalSchema(*temp_database, temp_table_name, nil)
		if cached_additonal_schema_errors != nil {
			return nil, cached_additonal_schema_errors
		} else if !common.IsNil(cached_additonal_schema) {
			return cached_additonal_schema, nil
		}
		
		sql_command, new_options,  sql_command_errors := getTableSchemaAdditionalSQLMySQL(struct_type, getTable(), options)
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return nil, errors
		}

		if json_array == nil {
			errors = append(errors, fmt.Errorf("error: show table status returned nil records"))
			return nil, errors
		}

		if len(*(json_array.GetValues())) == 0 {
			errors = append(errors, fmt.Errorf("error:  show table status did not return any records"))
			return nil, errors
		}

		additional_schema := json.NewMapValue()
		for _, column_details := range *(json_array.GetValues()) {
			column_map, column_map_errors := column_details.GetMap()
			if column_map_errors != nil {
				return nil, column_map_errors
			} else if common.IsNil(column_map) {
				errors = append(errors, fmt.Errorf("column_map is nil"))
				return nil, errors
			}
			column_attributes := column_map.GetKeys()

			for _, column_attribute := range column_attributes {
				switch column_attribute {
				case "Comment":
					comment_value, comment_errors := column_map.GetString("Comment")
					if comment_errors != nil {
						errors = append(errors, comment_errors...)
					} else if common.IsNil(comment_value) {
						errors = append(errors, fmt.Errorf("comment is nil"))
					} else {
						if strings.TrimSpace(*comment_value) != "" {
							comment_as_map, comment_as_map_value_errors := json.Parse(strings.TrimSpace(*comment_value))
							if comment_as_map_value_errors != nil {
								errors = append(errors, comment_as_map_value_errors...)
							} else if common.IsNil(comment_as_map) {
								errors = append(errors, fmt.Errorf("comment is nil"))
							} else {
								additional_schema.SetMap("Comment", comment_as_map)
							}
						}
					}
				default:
					column_attribute_value, column_attribute_value_errors := column_map.GetString(column_attribute)
					if column_attribute_value_errors != nil {
						errors = append(errors, column_attribute_value_errors...)
					} else if common.IsNil(column_attribute_value) {
						errors = append(errors, fmt.Errorf("%s is nil", column_attribute))
					} else {
						additional_schema.SetStringValue(column_attribute, *column_attribute_value)
					}
				}
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}


		temp_client_manager.GetOrSetAdditonalSchema(*temp_database, temp_table_name, &additional_schema)
		return &additional_schema, nil
	}


	getSchema := func() (*json.Map, []error) {
		var errors []error
		validate_errors := validate()
		
		if validate_errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return nil, temp_table_name_errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		temp_client_manager, temp_client_manager_errors := temp_client.GetClientManager()
		if temp_client_manager_errors != nil {
			return nil, temp_client_manager_errors
		}

		cached_schema, cached_schema_errors := temp_client_manager.GetOrSetSchema(*temp_database, temp_table_name, nil)
		if cached_schema_errors != nil {
			return nil, cached_schema_errors
		} else if !common.IsNil(cached_schema) {
			return cached_schema, nil
		}
		
		sql_command, new_options, sql_command_errors := getTableSchemaSQLMySQL(struct_type, getTable(), options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		}

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return nil, errors
		}

		schema, schem_errors := mapTableSchemaFromDBMySQL(struct_type, getTable(), json_array)
		if schem_errors != nil {
			errors = append(errors, schem_errors...)
			return nil, errors
		}

		temp_client_manager.GetOrSetSchema(*temp_database, temp_table_name, schema)
		return schema, nil
	}

	setTableName := func(new_table_name string) []error {
		return SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[table_name]", new_table_name)
	}

	getCreateTableSQL := func(options *json.Map) (*string, *json.Map, []error) {	
		return getCreateTableSQLMySQL(struct_type, getTable(), getData(), options)
	}

	createTable := func() []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		sql_command, new_options, sql_command_errors := getCreateTableSQL(options)

		if sql_command_errors != nil {
			return sql_command_errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return temp_database_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return temp_client_errors
		}

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

		if execute_errors != nil {
			return execute_errors
		}

		return nil
	}

	read := func() []error {
		errors := validate()

		if len(errors) > 0 {
			return errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			errors = append(errors, temp_database_errors...)
		} else if common.IsNil(temp_database) {
			errors = append(errors, fmt.Errorf("error: Table.read database is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			errors = append(errors, temp_table_name_errors...)
		} else if common.IsNil(temp_table_name) {
			errors = append(errors, fmt.Errorf("error: Table.read table_name is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		temp_schema, temp_schema_errors := getSchema()
		if temp_schema_errors != nil {
			errors = append(errors, temp_schema_errors...)
		} else if common.IsNil(temp_schema) {
			errors = append(errors, fmt.Errorf("error: Table.read schema is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		data = setupData(*temp_database, temp_table_name, temp_schema)
		return nil
	}

	getSchemaColumns := func()  (*[]string, []error) {
		errors := validate()

		if len(errors) > 0 {
			return nil, errors
		}

		schemas_map, schemas_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schemas_map_errors != nil {
			return nil, schemas_map_errors
		}

		schema_column_names := schemas_map.GetKeys()
		return &schema_column_names, nil
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	x := Table{
		Validate: func() []error {
			return validate()
		},
		GetDatabase: func() (*Database, []error) {
			return getDatabase()
		},
		GetTableColumns: func() (*[]string, []error) {
			return getTableColumns()
		},
		GetSchemaColumns: func() (*[]string, []error) {
			return getSchemaColumns()
		},
		GetIdentityColumns: func() (*[]string, []error) {
			return getIdentityColumns()
		},
		GetPrimaryKeyColumns: func() (*[]string, []error) {
			return getPrimaryKeyColumns()
		},
		GetForeignKeyColumns: func() (*[]string, []error) {
			return getForeignKeyColumns()
		},
		GetNonPrimaryKeyColumns: func() (*[]string, []error) {
			return getNonPrimaryKeyColumns()
		},
		Create: func() []error {
			errors := createTable()
			if errors != nil {
				return errors
			}

			return nil
		},
		Count: func(filters *json.Map, filters_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*uint64, []error) {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			select_fields := json.NewArray()
			select_fields.AppendStringValue("COUNT(*)")
			sql_command, new_options, sql_command_errors := getSelectRecordsSQLMySQL(getTable(), select_fields, filters, filters_logic, order_by, limit, offset, options, column_name_whitelist_characters)
			if sql_command_errors != nil {
				return nil, sql_command_errors
			}

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				return nil, temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, new_options)

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			if len(*(json_array.GetValues())) != 1 {
				errors = append(errors, fmt.Errorf("error: count record does not exist"))
				return nil, errors
			}

			map_record, map_record_errors := (*(json_array.GetValues()))[0].GetMap()
			if map_record_errors != nil {
				errors = append(errors, map_record_errors...)
				return nil, errors
			} else if common.IsNil(map_record) {
				errors = append(errors, fmt.Errorf("map_record is nil"))
				return nil, errors
			}

			count_value, count_value_error := map_record.GetString("COUNT(*)")
			if count_value_error != nil {
				errors = append(errors, count_value_error...)
				return nil, errors
			}

			count, count_err := strconv.ParseUint(*count_value, 10, 64)
			if count_err != nil {
				errors = append(errors, count_err)
				return nil, errors
			}

			return &count, nil
		},
		Delete: func() []error {
			return delete()
		},
		Read: func() []error {
			return read()
		},
		DeleteIfExists: func() []error {
			errors := validate()

			if len(errors) > 0 {
				return errors
			}

			return deleteIfExists()
		},
		CreateRecord: func(new_record_data json.Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := newRecord(*getTable(), new_record_data, database_reserved_words_obj,  column_name_whitelist_characters_obj)
			if record_errors != nil {
				return nil, record_errors
			}

			create_record_errors := record.Create()
			if create_record_errors != nil {
				return nil, create_record_errors
			}

			return record, nil
		},
		ReadRecords: func(select_fields *json.Array, filters *json.Map, filters_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*[]Record, []error) {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			sql_command, options, sql_command_errors := getSelectRecordsSQLMySQL(getTable(), select_fields, filters, filters_logic, order_by, limit, offset, options, column_name_whitelist_characters)
			if sql_command_errors != nil {
				return nil, sql_command_errors
			}

			additional_schema, additional_schema_errors := getAdditionalSchema()
			if additional_schema_errors != nil {
				return nil, additional_schema_errors
			}

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				return nil, temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			temp_client_manager, temp_client_manager_errors := temp_client.GetClientManager()
			if temp_client_manager_errors != nil {
				return nil, temp_client_manager_errors
			}

			cacheable := false
			additional_schema_comment, additional_schema_comment_errors := additional_schema.GetMap("Comment")
			if additional_schema_comment_errors != nil {
				return nil, additional_schema_comment_errors
			} else if !common.IsNil(additional_schema_comment) {
				cache, cache_errors := additional_schema_comment.GetBool("cache")
				if cache_errors != nil {
					errors = append(errors, cache_errors...)
				} else if !common.IsNil(cache) {
					cacheable = *cache
				}
			}

			if cacheable {
				cachable_records, cachable_records_errors := temp_client_manager.GetOrSetReadRecords(*temp_database, *sql_command, nil)
				if cachable_records_errors != nil {
					return nil, cachable_records_errors
				} else if !common.IsNil(cachable_records) {
					return cachable_records, nil
				}
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql_command, options)

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, current_json := range *(json_array.GetValues()) {
				mapped_record_obj, mapped_record_obj_errors := mapValueFromDBToRecord(getTable(), current_json, database_reserved_words_obj, column_name_whitelist_characters_obj)
				if mapped_record_obj_errors != nil {
					errors = append(errors, mapped_record_obj_errors...)
				} else if common.IsNil(mapped_record_obj){
					errors = append(errors, fmt.Errorf("mapped record is nil"))
				} else {
					mapped_records = append(mapped_records, *mapped_record_obj)
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			if cacheable {
				temp_client_manager.GetOrSetReadRecords(*temp_database, *sql_command, &mapped_records)
			}

			return &mapped_records, nil
		},
		UpdateRecords: func(records json.Array) ([]error) {
			return updateRecords(records)
		},
		UpdateRecord: func(record json.Map) ([]error) {
			return updateRecord(record)
		},
		CreateRecords: func(records json.Array) ([]error) {
			return createRecords(records)
		},
		Exists: func() (*bool, []error) {
			return exists()
		},
		GetSchema: func() (*json.Map, []error) {
			return getSchema()
		},
		GetAdditionalSchema: func() (*json.Map, []error) {
			return getAdditionalSchema()
		},
		GetTableName: func() (string, []error) {
			return getTableName()
		},
		SetTableName: func(table_name string) []error {
			return setTableName(table_name)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
	}
	setTable(&x)

	
	return &x, nil
}
