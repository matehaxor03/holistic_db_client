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
	getTableStatus        func() (*json.Map, []error)
	GetTableName          func() (string, []error)
	SetTableName          func(table_name string) []error
	GetSchemaColumns      func() (*[]string, []error)
	GetTableColumns       func() (*[]string, []error)
	GetIdentityColumns    func() (*[]string, []error)
	GetPrimaryKeyColumns  func() (*[]string, []error)
	GetForeignKeyColumns  func() (*[]string, []error)

	GetNonPrimaryKeyColumns func() (*[]string, []error)
	Count                 func() (*uint64, []error)
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
		options.SetBoolValue("use_file", false)
		
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
		
		errors := validate()
		if errors != nil {
			return errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return temp_table_name_errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return errors
		}

		sql := "DROP TABLE "
		if options.IsBoolTrue("use_file") {
			sql += fmt.Sprintf("`%s`;", table_name_escaped)
		} else {
			sql += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
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

	deleteIfExists := func() ([]error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		
		errors := validate()
		if errors != nil {
			return errors
		}

		temp_table_name, temp_table_name_errors := getTableName()
		if temp_table_name_errors != nil {
			return temp_table_name_errors
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return errors
		}

		sql := "DROP TABLE IF EXISTS "
		if options.IsBoolTrue("use_file") {
			sql += fmt.Sprintf("`%s`;", table_name_escaped)
		} else {
			sql += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
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

	getTableStatus := func() (*json.Map, []error) {
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

		temp_database_name, temp_database_name_errors := temp_database.GetDatabaseName()
		if temp_database_name_errors != nil {
			return nil, temp_database_name_errors
		}

		temp_client, temp_client_errors := temp_database.GetClient()
		if temp_client_errors != nil {
			return nil, temp_client_errors
		}

		temp_client_manager, temp_client_manager_errors := temp_client.GetClientManager()
		if temp_client_manager_errors != nil {
			return nil, temp_client_manager_errors
		}

		cached_table_status, cached_table_status_errors := temp_client_manager.GetOrSetTableStatus(*temp_database, temp_table_name, nil)
		if cached_table_status_errors != nil {
			return nil, cached_table_status_errors
		} else if !common.IsNil(cached_table_status) {
			return cached_table_status, nil
		}

		table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return nil, errors
		}

		database_name_escaped, database_name_escaped_errors := common.EscapeString(temp_database_name, "'")
		if database_name_escaped_errors != nil {
			errors = append(errors, database_name_escaped_errors)
			return nil, errors
		}
		
		sql_command := "SHOW TABLE STATUS FROM "
		
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s` ", database_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\` ", database_name_escaped)
		}

		sql_command += "WHERE name='" + table_name_escaped + "';"

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, options)

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

		table_status := json.NewMapValue()
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
								table_status.SetMap("Comment", comment_as_map)
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
						table_status.SetStringValue(column_attribute, *column_attribute_value)
					}
				}
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}


		temp_client_manager.GetOrSetTableStatus(*temp_database, temp_table_name, &table_status)
		return &table_status, nil
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

		table_name_escaped, table_name_escaped_errors := common.EscapeString(temp_table_name, "'")
		if table_name_escaped_errors != nil {
			errors = append(errors, table_name_escaped_errors)
			return nil, errors
		}
		
		sql_command := "SHOW FULL COLUMNS FROM "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s`;", table_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
		}

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql_command, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return nil, errors
		}

		if json_array == nil {
			errors = append(errors, fmt.Errorf("error: show columns returned nil records"))
			return nil, errors
		}

		if len(*(json_array.GetValues())) == 0 {
			errors = append(errors, fmt.Errorf("error: show columns did not return any records"))
			return nil, errors
		}

		schema := json.NewMapValue()
		for _, column_details := range *(json_array.GetValues()) {
			column_map, column_map_errors := column_details.GetMap()
			if column_map_errors != nil {
				return nil, column_map_errors
			} else if common.IsNil(column_map) {
				errors = append(errors, fmt.Errorf("column_map is nil"))
				return nil, errors
			}
			column_attributes := column_map.GetKeys()

			column_schema := json.NewMapValue()
			default_value := ""
			field_name := ""
			is_nullable := false
			is_primary_key := false
			is_unique := false
			extra_value := ""
			comment_value := ""
			for _, column_attribute := range column_attributes {
				switch column_attribute {
				case "Key":
					key_value, _ := column_map.GetString("Key")
					switch *key_value {
					case "PRI":
						is_primary_key = true
						is_nullable = false
						column_schema.SetBool("primary_key", &is_primary_key)
					case "", "MUL":
					case "UNI":
						is_unique = true
						column_schema.SetBool("unique", &is_unique)
					default:
						errors = append(errors, fmt.Errorf("error: Table: GetSchema: Key not implemented please implement: %s", *key_value))
					}
				case "Field":
					field_name_value, _ := column_map.GetString("Field")
					field_name = *field_name_value
				case "Type":
					type_of_value, _ := column_map.GetString("Type")
					switch *type_of_value {
					case "bigint unsigned":
						data_type := "uint64"
						unsigned := true
						column_schema.SetString("type", &data_type)
						column_schema.SetBool("unsigned", &unsigned)
					case "int unsigned":
						data_type := "uint32"
						unsigned := true
						column_schema.SetString("type", &data_type)
						column_schema.SetBool("unsigned", &unsigned)
					case "mediumint unsigned":
						data_type := "uint32"
						unsigned := true
						column_schema.SetString("type", &data_type)
						column_schema.SetBool("unsigned", &unsigned)
					case "smallint unsigned":
						data_type := "uint16"
						unsigned := true
						column_schema.SetString("type", &data_type)
						column_schema.SetBool("unsigned", &unsigned)
					case "tinyint unsigned":
						data_type := "uint8"
						unsigned := true
						column_schema.SetString("type", &data_type)
						column_schema.SetBool("unsigned", &unsigned)
					case "bigint":
						data_type := "int64"
						column_schema.SetString("type", &data_type)
					case "int":
						data_type := "int32"
						column_schema.SetString("type", &data_type)
					case "mediumint":
						data_type := "int32"
						column_schema.SetString("type", &data_type)
					case "smallint":
						data_type := "int16"
						column_schema.SetString("type", &data_type)
					case "tinyint":
						data_type := "int8"
						column_schema.SetString("type", &data_type)
					case "timestamp":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(0))
					case "timestamp(1)":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(1))
					case "timestamp(2)":						
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(2))
					case "timestamp(3)":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(3))
					case "timestamp(4)":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(4))
					case "timestamp(5)":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(5))
					case "timestamp(6)":
						data_type := "time.Time"
						column_schema.SetString("type", &data_type)
						column_schema.SetUInt8Value("decimal_places", uint8(6))
					case "tinyint(1)":
						data_type := "bool"
						column_schema.SetString("type", &data_type)
					case "text", "blob", "json":
						data_type := "string"
						column_schema.SetString("type", &data_type)
					case "float":
						data_type := "float32"
						column_schema.SetString("type", &data_type)
					case "double":
						data_type := "float64"
						column_schema.SetString("type", &data_type)
					default:
						if strings.HasPrefix(*type_of_value, "char(") && strings.HasSuffix(*type_of_value, ")") {
							data_type := "string"
							column_schema.SetString("type", &data_type)
						} else if strings.HasPrefix(*type_of_value, "varchar(") && strings.HasSuffix(*type_of_value, ")") {
							data_type := "string"
							column_schema.SetString("type", &data_type)
						} else if strings.HasPrefix(*type_of_value, "enum(")  && strings.HasSuffix(*type_of_value, ")") {
							type_of_value_values := (*type_of_value)[5:len(*type_of_value)-1]
							parts := strings.Split(type_of_value_values, ",")
							if len(parts) == 0 {
								errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum had length of zero: %s", *type_of_value))
							} else {
								part := parts[0]
								if strings.HasPrefix(part, "'")  && strings.HasSuffix(part, "'") {
									data_type := "string"
									column_schema.SetString("type", &data_type)
								} else {
									errors = append(errors, fmt.Errorf("error: Table: GetSchema: could not determine parts of enum for data type: %s", *type_of_value))
								}
							}
						} else {
							errors = append(errors, fmt.Errorf("error: Table: GetSchema: type not implemented please implement: %s", *type_of_value))
						}
					}
				case "Null":
					null_value, _ := column_map.GetString("Null")
					switch *null_value {
					case "YES":
						if !is_primary_key {
							is_nullable = true
						}
					case "NO":
						is_nullable = false
					default:
						errors = append(errors, fmt.Errorf("error: Table: GetSchema: Null value not supported please implement: %s", *null_value))
					}
				case "Default":
					default_val, _ := column_map.GetString("Default")
					default_value = *default_val
				case "Extra":
					extra_val, _ := column_map.GetString("Extra")
					extra_value = *extra_val
					switch extra_value {
					case "auto_increment":
						auto_increment := true
						column_schema.SetBool("auto_increment", &auto_increment)
					case "DEFAULT_GENERATED":
					case "":
					default:
						errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: Extra value not supported please implement: %s", temp_table_name, extra_value))
					}
				case "Privileges":
				case "Collation":
				case "Comment":
					comment_val, comment_errors := column_map.GetString("Comment")
					if comment_errors != nil {
						errors = append(errors, comment_errors...)
					} else {
						comment_value = *comment_val
						if strings.TrimSpace(comment_value) != "" {
							comment_as_map, comment_as_map_value_errors := json.Parse(strings.TrimSpace(comment_value))
							if comment_as_map_value_errors != nil {
								errors = append(errors, comment_as_map_value_errors...)
							} else if common.IsNil(comment_as_map) {
								errors = append(errors, fmt.Errorf("comment is nil"))
							} else {
								rules_array, rules_array_errors := comment_as_map.GetArray("rules")
								if rules_array_errors != nil {
									errors = append(errors, rules_array_errors...)
								} else if !common.IsNil(rules_array) {
									filters := json.NewArrayValue()
									for _, rule := range *(rules_array.GetValues()) {
										rule_value, rule_value_errors := rule.GetString()
										if rule_value_errors != nil {
											return nil, rule_value_errors
										} else if common.IsNil(rule_value) {
											errors = append(errors, fmt.Errorf("rule value is nil"))
											return nil, errors
										}

										switch *rule_value {
										case "domain_name":
											domain_name_filter := json.NewMapValue()
											domain_name_filter.SetObjectForMap("values", get_domain_name_characters())
											domain_name_filter.SetObjectForMap("function", getWhitelistCharactersFunc())
											filters.AppendMapValue(domain_name_filter)
										case "repository_name":
											repostiory_name_filter := json.NewMapValue()
											repostiory_name_filter.SetObjectForMap("values", get_repository_name_characters())
											repostiory_name_filter.SetObjectForMap("function", getWhitelistCharactersFunc())
											filters.AppendMapValue(repostiory_name_filter)
										case "repository_account_name":
											repository_account_name_filter := json.NewMapValue()
											repository_account_name_filter.SetObjectForMap("values", get_repository_account_name_characters())
											repository_account_name_filter.SetObjectForMap("function", getWhitelistCharactersFunc())
											filters.AppendMapValue(repository_account_name_filter)
										case "branch_name":
											branch_name_filter := json.NewMapValue()
											branch_name_filter.SetObjectForMap("values", get_branch_name_characters())
											branch_name_filter.SetObjectForMap("function", getWhitelistCharactersFunc())
											filters.AppendMapValue(branch_name_filter)
										default:
											errors = append(errors, fmt.Errorf("rule not supported %s", rule_value))
										}
									}
									column_schema.SetArray("filters", &filters)
								}

								foreign_key_map, foreign_key_map_errors := comment_as_map.GetMap("foreign_key")
								if foreign_key_map_errors != nil {
									errors = append(errors, foreign_key_map_errors...)
								} else if !common.IsNil(foreign_key_map) {
									foreign_key_true := true
									column_schema.SetBool("foreign_key", &foreign_key_true)

									foreign_key_table_name, foreign_key_table_name_errors := foreign_key_map.GetString("table_name")
									if foreign_key_table_name_errors != nil {
										errors = append(errors, foreign_key_table_name_errors...)
									} else if common.IsNil(foreign_key_table_name) {
										errors = append(errors, fmt.Errorf("foreign_key table_name is nil"))
									} else {
										column_schema.SetString("foreign_key_table_name", foreign_key_table_name)
									}

									foreign_key_column_name, foreign_key_column_name_errors := foreign_key_map.GetString("column_name")
									if foreign_key_column_name_errors != nil {
										errors = append(errors, foreign_key_column_name_errors...)
									} else if common.IsNil(foreign_key_column_name) {
										errors = append(errors, fmt.Errorf("foreign_key column_name is nil"))
									} else {
										column_schema.SetString("foreign_key_column_name", foreign_key_column_name)
									}

									foreign_key_type, foreign_key_type_errors := foreign_key_map.GetString("type")
									if foreign_key_type_errors != nil {
										errors = append(errors, foreign_key_type_errors...)
									} else if common.IsNil(foreign_key_type) {
										errors = append(errors, fmt.Errorf("foreign_key type is nil"))
									} else {
										column_schema.SetString("foreign_key_type", foreign_key_type)
									}
								}
							}
						}
					}
				default:
					errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: %s not supported please implement", temp_table_name, field_name, column_attribute))
				}
			}

			if column_schema.IsNil("type") {
				errors = append(errors, fmt.Errorf("error: Table: %s GetSchema: column: %s attribute: type is nill", temp_table_name, field_name))
			}

			if len(errors) > 0 {
				continue
			}

			dt, _ := column_schema.GetString("type")

		
			if default_value == "NULL" {
				column_schema.SetNil("default")
			} else {
				if *dt == "string" {
					column_schema.SetString("default", &default_value)
				} else if *dt == "uint64" && default_value != "" {
					number, err := strconv.ParseUint(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						column_schema.SetUInt64("default", &number)
					}
				} else if *dt == "int64" && default_value != "" {
					number, err := strconv.ParseInt(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						column_schema.SetInt64("default", &number)
					}
				} else if *dt == "uint32" && default_value != "" {
					number, err := strconv.ParseUint(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := uint32(number)
						column_schema.SetUInt32("default", &converted)
					}
				} else if *dt == "int32" && default_value != "" {
					number, err := strconv.ParseInt(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := int32(number)
						column_schema.SetInt32("default", &converted)
					}
				} else if *dt == "uint16" && default_value != "" {
					number, err := strconv.ParseUint(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := uint16(number)
						column_schema.SetUInt16("default", &converted)
					}
				} else if *dt == "int16" && default_value != "" {
					number, err := strconv.ParseInt(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := int16(number)
						column_schema.SetInt16("default", &converted)
					}
				} else if *dt == "uint8" && default_value != "" {
					number, err := strconv.ParseUint(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := uint8(number)
						column_schema.SetUInt8("default", &converted)
					}
				} else if *dt == "int8" && default_value != "" {
					number, err := strconv.ParseInt(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := int8(number)
						column_schema.SetInt8("default", &converted)
					}
				} else if *dt == "float32" && default_value != "" {
					number, err := strconv.ParseFloat(default_value, 32)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := float32(number)
						column_schema.SetFloat32("default", &converted)
					}
				} else if *dt == "float64" && default_value != "" {
					number, err := strconv.ParseFloat(default_value, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						converted := float64(number)
						column_schema.SetFloat64("default", &converted)
					}
				}  else if *dt == "bool" && default_value != "" {
					number, err := strconv.ParseInt(default_value, 10, 64)
					if err != nil {
						errors = append(errors, err)
					} else {
						if number == 0 {
							boolean_value := false
							column_schema.SetBool("default", &boolean_value)
						} else if number == 1 {
							boolean_value := true
							column_schema.SetBool("default", &boolean_value)
						} else {
							errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 1 or 0", default_value, *dt))
						}
					}
				} else if *dt == "time.Time" && default_value != "" {
					if extra_value == "DEFAULT_GENERATED" && strings.HasPrefix(default_value, "CURRENT_TIMESTAMP") {
						if default_value == "CURRENT_TIMESTAMP" {
						} else if default_value == "CURRENT_TIMESTAMP(1)" {
						} else if default_value == "CURRENT_TIMESTAMP(2)" {
						} else if default_value == "CURRENT_TIMESTAMP(3)" {
						} else if default_value == "CURRENT_TIMESTAMP(4)" {
						} else if default_value == "CURRENT_TIMESTAMP(5)" {
						} else if default_value == "CURRENT_TIMESTAMP(6)" {
						} else {
							errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be 0-6 decimal places", default_value, *dt))
						}
						
						default_value = "now"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.0" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.00" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.000" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.0000" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.00000" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					} else if default_value == "0000-00-00 00:00:00.000000" {
						default_value = "zero"
						column_schema.SetString("default", &default_value)
					}  else {
						errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported %s for type: %s can only be DEFAULT_GENERATED or 0000-00-00 00:00:00", default_value, *dt))
					}
				} else if !(*dt == "time.Time" || *dt == "bool" || *dt == "int64" || *dt == "uint64" ||  *dt == "int32" || *dt == "uint32" ||  *dt == "int16" || *dt == "uint16" ||  *dt == "int8" || *dt == "uint8" || *dt == "string" || *dt == "float32" || *dt == "float64") && default_value != "" {
					errors = append(errors, fmt.Errorf("error: Table.GetSchema default value not supported please implement: %s for type: %s", default_value, *dt))
				}
			}
			

			if is_nullable {
				adjusted_type := "*" + *dt
				column_schema.SetString("type", &adjusted_type)
			}

			schema.SetMapValue(field_name, column_schema)
		}

		if len(errors) > 0 {
			return nil, errors
		}


		temp_client_manager.GetOrSetSchema(*temp_database, temp_table_name, &schema)
		return &schema, nil
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
		Count: func() (*uint64, []error) {
			options := json.NewMap()
			options.SetBoolValue("use_file", false)
			errors := validate()
			if errors != nil {
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

			sql := "SELECT COUNT(*) FROM "
			if options.IsBoolTrue("use_file") {
				sql += fmt.Sprintf("`%s`;", table_name_escaped)
			} else {
				sql += fmt.Sprintf("\\`%s\\`;", table_name_escaped)
			}

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				return nil, temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return nil, temp_client_errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, &sql, options)

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

			table_status, table_status_errors := getTableStatus()
			if table_status_errors != nil {
				return nil, table_status_errors
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
			table_status_comment, table_status_comment_errors := table_status.GetMap("Comment")
			if table_status_comment_errors != nil {
				return nil, table_status_comment_errors
			} else if !common.IsNil(table_status_comment) {
				cache, cache_errors := table_status_comment.GetBool("cache")
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
		getTableStatus: func() (*json.Map, []error) {
			return getTableStatus()
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
