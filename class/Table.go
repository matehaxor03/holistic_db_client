package class

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
		map_database_schema.SetStringValue("type", "class.Database")
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
	
		for _, schema_key_from_db := range schema_from_db.Keys() {
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
		columns := temp_schemas.Keys()
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

		for _, column := range schema_map.Keys() {
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

		for _, column := range schema_map.Keys() {
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

		for _, column := range schema_map.Keys() {
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

		for _, column := range schema_map.Keys() {
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
		return ValidateData(getData(), "*class.Table")
	}

	getDatabase := func() (*Database, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database]", "*class.Database")
		return temp_value.(*Database), temp_value_errors
	}

	exists := func() (*bool, []error) {
		options := json.NewMapValue()
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

	delete := func() ([]error) {
		options := json.NewMapValue()
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

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		options := json.NewMapValue()
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

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}
	
	updateRecords := func(records json.Array) []error {
		options := json.NewMapValue()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", true)

		errors := validate()
		if errors != nil {
			return errors
		}

		if len(*(records.GetValues())) == 0 {
			return nil
		}

		for _, record := range *(records.GetValues()) {
			if !common.IsMap(record) {
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

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	updateRecord := func(record json.Map) []error {
		options := json.NewMapValue()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", true)

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

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	createRecords := func(records json.Array) []error {
		options := json.NewMapValue()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", true)

		errors := validate()
		if errors != nil {
			return errors
		}

		if len(*(records.GetValues())) == 0 {
			return nil
		}

		for _, record := range *(records.GetValues()) {
			if !common.IsMap(record) {
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
			sql_update_snippet, sql_update_snippet_errors := record_obj.GetCreateSQL()
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

		_, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql, options)

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
		options := json.NewMapValue()
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

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

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
			column_attributes := column_map.Keys()

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
		options := json.NewMapValue()
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

		json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

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
			column_attributes := column_map.Keys()

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

	getCreateSQL := func(options json.Map) (*string, []error) {
		errors := validate()

		if len(errors) > 0 {
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

		sql_command := ""
		sql_command += "CREATE TABLE "
		if options.IsBoolTrue("use_file") {
			sql_command += fmt.Sprintf("`%s` ", table_name_escaped)
		} else {
			sql_command += fmt.Sprintf("\\`%s\\` ", table_name_escaped)
		}

		valid_columns, valid_columns_errors := getTableColumns()
		if valid_columns_errors != nil {
			return nil, valid_columns_errors
		}

		schemas_map, schemas_map_errors := GetSchemas(struct_type, getData(), "[schema]")
		if schemas_map_errors != nil {
			return nil, schemas_map_errors
		}

		primary_key_count := 0

		sql_command += "("
		for index, column := range *valid_columns {
			columnSchema, columnSchema_errors := schemas_map.GetMap(column)
			if columnSchema_errors != nil {
				errors = append(errors, columnSchema_errors...)
				continue
			} else if common.IsNil(columnSchema) {
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL %s column schema for column: %s is nil", struct_type, column))
				continue
			}

			column_escaped, column_escaped_errors := common.EscapeString(column, "'")
			if column_escaped_errors != nil {
				errors = append(errors, column_escaped_errors)
			}

			if options.IsBoolTrue("use_file") {
				sql_command += "`"
			} else {
				sql_command += "\\`"
			}
			sql_command += column_escaped
			
			if options.IsBoolTrue("use_file") {
				sql_command += "`"
			} else {
				sql_command += "\\`"
			}

			typeOf, type_of_errors := columnSchema.GetString("type")
			if type_of_errors != nil {
				errors = append(errors, type_of_errors...)
				continue
			}

			switch *typeOf {
			case "*uint64", "uint64","*int64", "int64", "*uint32", "uint32", "*int32","int32", "*uint16", "uint16", "*int16","int16",  "*uint8", "uint8", "*int8","int8":
				switch *typeOf {
				case "*uint64", "*int64", "uint64", "int64":
					sql_command += " BIGINT"
				case "*uint32", "*int32", "uint32", "int32":
					sql_command += " INT"
				case "*uint16", "*int16", "uint16", "int16":
					sql_command += " SMALLINT"
				case "*uint8", "*int8", "uint8", "int8":
					sql_command += " TINYINT"
				default:
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL number type not mapped: %s", *typeOf))
				}

				unsigned_number := false
				switch *typeOf {
				case "*uint64", "uint64":
					unsigned_number = true
				case "*uint32", "uint32":
					unsigned_number = true
				case "*uint16", "uint16":
					unsigned_number = true
				case "*uint8","uint8":
					unsigned_number = true
				default:
				}

				if unsigned_number {
					sql_command += " UNSIGNED"
				}

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("auto_increment") {
					if columnSchema.IsBool("auto_increment") && !columnSchema.IsNil("auto_increment") {
						if columnSchema.IsBoolTrue("auto_increment") {
							sql_command += " AUTO_INCREMENT"
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.GetType("auto_increment")))
					}
				}

				if columnSchema.HasKey("primary_key") {
					if columnSchema.IsBool("primary_key") && !columnSchema.IsNil("primary_key") {
						if columnSchema.IsBoolTrue("primary_key") {
							sql_command += " PRIMARY KEY"
							primary_key_count += 1
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
					}
				} 

				if columnSchema.HasKey("default") {
					if columnSchema.IsInteger("default") {
						default_value, default_value_errors := columnSchema.GetInt64("default")
						if default_value_errors != nil {
							errors = append(errors, default_value_errors...)
						} else {
							sql_command += " DEFAULT " + strconv.FormatInt(*default_value, 10)
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: default contained a value which is not supported: %s", column, columnSchema.GetType("default")))
					}
				}
			case "*time.Time", "time.Time":
				decimal_places, decimal_places_error := columnSchema.GetInt("decimal_places")
				if decimal_places_error != nil {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained a value which is not supported %s", column, fmt.Sprintf("%s", decimal_places_error)))
				} else if common.IsNil(decimal_places) {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained a value which is not supported: nil", column))
				} else if *decimal_places < 0  || *decimal_places > 6 {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s for attribute: decimal_places contained invalid decimal range outside [0-6]: %d", column, *decimal_places))
				} else {
					if *decimal_places == 0 {
						sql_command += " TIMESTAMP"
					} else {
						sql_command += fmt.Sprintf(" TIMESTAMP(%d)", *decimal_places)
					}

					if !strings.HasPrefix(*typeOf, "*") {
						sql_command += " NOT NULL"
					}

					if columnSchema.HasKey("default") {
						default_value, default_value_errors := columnSchema.GetString("default")
						if default_value_errors != nil {
							errors = append(errors, default_value_errors...)
						} else if default_value == nil {
							sql_command += " DEFAULT NULL"
						} else if *default_value == "now" {
							if *decimal_places == 0 {
								sql_command += " DEFAULT CURRENT_TIMESTAMP"
							} else {
								sql_command += fmt.Sprintf(" DEFAULT CURRENT_TIMESTAMP(%d)", *decimal_places)
							}
						} else if *default_value == "zero" {
							default_time, default_time_errors := columnSchema.GetTimeWithDecimalPlaces("default", *decimal_places)
							if default_time_errors != nil {
								errors = append(errors, default_time_errors...)
							} else if common.IsNil(default_time) {
								errors = append(errors, fmt.Errorf("default zero time was nil"))
							} else {
								time_zero_as_string, time_zero_as_string_errors := common.GetTimeZeroStringSQL(*decimal_places)
								if time_zero_as_string_errors != nil {
									errors = append(errors, time_zero_as_string_errors...)
								} else {
									value_escaped, value_escaped_errors := common.EscapeString(*time_zero_as_string, "'")
									if value_escaped_errors != nil {
										errors = append(errors, value_escaped_errors)
									} else {
										sql_command += " DEFAULT "
										if options.IsBoolTrue("use_file") {
											sql_command += "'" + value_escaped + "'"
										} else {
											sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
										}
									}	
								}
							}
						}  else {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had default value it did not understand", column))
						}
					}




				}

				

				
			case "*bool", "bool":
				sql_command += " BOOLEAN"

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had nil default value", column))
					} else if !columnSchema.IsBool("default") {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
					} else if columnSchema.IsBoolTrue("default") {
						sql_command += " DEFAULT 1"
					} else if columnSchema.IsBoolFalse("default") {
						sql_command += " DEFAULT 0"
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
					}
				}
			case "*float32", "float32":
				sql_command += " FLOAT"

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						sql_command += " DEFAULT NULL"
					} else if !columnSchema.IsFloat("default") {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
					} else if columnSchema.IsFloat("default") {
						default_float_value, default_float_value_errors := columnSchema.GetFloat32("default")
						if default_float_value_errors != nil {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for float32 default value %s", column, fmt.Sprintf("%s", default_float_value_errors)))
						} else if default_float_value == nil {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s float32 default value returned nil", column))
						} else {
							sql_command += fmt.Sprintf(" DEFAULT %f", *default_float_value)
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
					}
				}
			case "*float64", "float64":
				sql_command += " DOUBLE"

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						sql_command += " DEFAULT NULL"
					} else if !columnSchema.IsFloat("default") {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-boolean default value", column))
					} else if columnSchema.IsFloat("default") {
						default_float_value, default_float_value_errors := columnSchema.GetFloat64("default")
						if default_float_value_errors != nil {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for float32 default value %s", column, fmt.Sprintf("%s", default_float_value_errors)))
						} else if default_float_value == nil {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s float32 default value returned nil", column))
						} else {
							sql_command += fmt.Sprintf(" DEFAULT %f", *default_float_value)
						}
					} else {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had unknown error for boolean default value", column))
					}
				}
			case "*string", "string":
				sql_command += " VARCHAR("
				if !columnSchema.HasKey("max_length") {
					errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s did not specify max_length attribute", column))
				} else if columnSchema.GetType("max_length") != "int" {
					errors = append(errors, fmt.Errorf("error: column: %s specified length attribute however it's not an int", column))
				} else {
					max_length, max_length_errors := columnSchema.GetInt("max_length")
					if max_length_errors != nil {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified max_length attribute had errors %s", column, fmt.Sprintf("%s", max_length_errors)))
					} else if *max_length < 0 {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified max_length attribute was < 0 and had value: %d", column, max_length))
					} else {
						// utf-8 should use 4 bytes (maxiumum per character) but in mysql it's 3 bytes but to be consistent going to assume 4 bytes, 
						sql_command += fmt.Sprintf("%d", (4*(*max_length)))
					}
				}
				sql_command += ")"

				if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						sql_command += " DEFAULT NULL"
					} else if !columnSchema.IsString("default") {
						errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s had non-string default value", column))
					} else {
						default_value, default_value_errors := columnSchema.GetString("default")
						if default_value_errors != nil {
							errors = append(errors, fmt.Errorf("error: Table.getCreateSQL column: %s specified default attribute had errors %s", column, fmt.Sprintf("%s", default_value_errors)))
						} else {

							default_value_escaped, default_value_escaped_errors := common.EscapeString(*default_value, "'")
							if default_value_escaped_errors != nil {
								errors = append(errors, default_value_escaped_errors)
							}

							sql_command += " DEFAULT "

							if options.IsBoolTrue("use_file") {
								sql_command += "'" + default_value_escaped + "'"
							} else {
								sql_command += strings.ReplaceAll("'" + default_value_escaped + "'", "`", "\\`")
							}

						}
					} 
				}

				
			default:
				errors = append(errors, fmt.Errorf("error: Table.getCreateSQL type: %s is not supported please implement for column %s", *typeOf, column))
			}

			if index < (len(*valid_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ");"

		if primary_key_count == 0 {
			errors = append(errors, fmt.Errorf("error: Table.getCreateSQL: %s must have at least 1 primary key", table_name_escaped))
		}

		// todo: check that length of row for all columns does not exceed 65,535 bytes (it's not hard but low priority)

		if len(errors) > 0 {
			return nil, errors
		}

		return &sql_command, nil
	}

	createTable := func() []error {
		options := json.NewMapValue()
		options.SetBoolValue("use_file", false)

		sql_command, sql_command_errors := getCreateSQL(options)

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

		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql_command, options)

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

		schema_column_names := schemas_map.Keys()
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
			options := json.NewMapValue()
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

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql, options)

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
			options := json.NewMapValue()
			options.SetBoolValue("use_file", false)
			cacheable := false
			var errors []error
			validate_errors := validate()
			if errors != nil {
				errors = append(errors, validate_errors...)
				return nil, errors
			}


			table_schema, table_schema_errors := getSchema()
			if table_schema_errors != nil {
				return nil, table_schema_errors
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

			if filters != nil {
				table_columns, table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				filter_columns := filters.Keys()
				for _, filter_column := range filter_columns {
					if !common.Contains(*table_columns, filter_column) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", filter_column, temp_table_name, *table_columns))
					}
				}

				if len(errors) > 0 {
					return nil, errors
				}

				for _, filter_column := range filter_columns {
					filter_column_type := filters.GetType(filter_column)

					if !filters.IsNil(filter_column) && !strings.HasPrefix(filter_column_type, "*") {
						filter_column_type = "*" + filter_column_type
					}
					 
					if table_schema.IsNil(filter_column) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s does not exist however filter had the value, table has columns: %s", filter_column, temp_table_name, table_schema.Keys()))
						continue
					}

					table_schema_column, table_schema_column_errors := table_schema.GetMap(filter_column)
					if table_schema_column_errors != nil {
						errors = append(errors, table_schema_column_errors...)
						continue
					}

					if table_schema_column.IsNil("type") {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s for table: %s did not have atrribute: type", filter_column, temp_table_name))
						continue
					}


					table_column_type, _ := (*table_schema_column).GetString("type")
					if strings.Replace(*table_column_type, "*", "", -1) != strings.Replace(filter_column_type, "*", "", -1) {
						table_column_type_simple := strings.Replace(*table_column_type, "*", "", -1)
						filter_column_type_simple := strings.Replace(filter_column_type, "*", "", -1)
						if strings.Contains(table_column_type_simple, "int") && strings.Contains(filter_column_type_simple, "int") {

						} else if strings.Contains(table_column_type_simple, "float") && strings.Contains(filter_column_type_simple, "float"){

						} else {
							errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, temp_table_name, *table_column_type))
						}
						

						//todo ignore if filter data_type is nil and table column allows nil
					}
				}
			}

			if select_fields != nil {
				table_columns, table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				for _, select_field := range *(select_fields.GetValues()) {
					select_field_value, select_field_value_errors := select_field.GetString()
					if select_field_value_errors != nil {
						return nil, select_field_value_errors
					} else if common.IsNil(select_field_value) {
						errors = append(errors, fmt.Errorf("select_field_value is nil"))
						return nil, errors
					}

					if !common.Contains(*table_columns, *select_field_value) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", select_field, temp_table_name, *table_columns))
					}
				}
			}

			if filters_logic  != nil {
				table_columns, table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				temp_filters_fields := filters_logic.Keys()
				for _, temp_filters_field := range temp_filters_fields {
					if !common.Contains(*table_columns, temp_filters_field) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s not found for table: %s available columns are: %s", temp_filters_field, temp_table_name, *table_columns))
						continue
					}

					if !filters_logic.IsString(temp_filters_field) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter logic is not a string", temp_filters_field, temp_table_name))
					} else {
						temp_logic_value, temp_logic_value_errors := filters_logic.GetString(temp_filters_field)
						if temp_logic_value_errors != nil {
							errors = append(errors, temp_logic_value_errors...)
						} else if common.IsNil(temp_logic_value) {
							errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value had error getting string value", temp_filters_field, temp_table_name))
						} else {
							if !(*temp_logic_value != "=" ||
							   *temp_logic_value != ">" ||
							   *temp_logic_value != "<") {
								errors = append(errors, fmt.Errorf("error: Table.ReadRecords: column: %s found for table %s however filter value logic however it's not supported please implement: %s", temp_filters_field, temp_table_name, *temp_logic_value))
							} 
						}
					}
				}
			}

			order_by_clause := ""
			if !common.IsNil(order_by) && len(*(order_by.GetValues())) > 0 {
				table_columns, table_columns_errors := getTableColumns()
				if table_columns_errors != nil {
					return nil, table_columns_errors
				}

				order_by_columns := len(*(order_by.GetValues()))
				for order_by_index, order_by_field := range *(order_by.GetValues()) {

					if common.IsNil(order_by_field) {
						errors = append(errors, fmt.Errorf("order_by_field is nil"))
						continue
					} 
					
					order_by_map, order_by_map_errors := order_by_field.GetMap()
					if order_by_map_errors != nil {
						errors = append(errors, order_by_map_errors...)
						continue
					} else if common.IsNil(order_by_map) {
						errors = append(errors, fmt.Errorf("order_by_map is nil"))
						continue
					}
					
					order_by_map_column_names := order_by_map.Keys()
					if len(order_by_map_column_names) != 1 {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: order by field at index %d was a map however did not have a column name", order_by_index))
						continue
					}

					order_by_column_name := order_by_map_column_names[0]
					
					if !common.Contains(*table_columns, order_by_column_name) {
						errors = append(errors, fmt.Errorf("error: Table.ReadRecords: order by column: %s not found for table: %s available columns are: %s", order_by_column_name, temp_table_name, *table_columns))
						continue
					}

					order_by_string_value, order_by_string_value_errors := order_by_map.GetString(order_by_column_name)
					if order_by_string_value_errors != nil {
						errors = append(errors, order_by_string_value_errors...)
						continue
					} else if common.IsNil(order_by_string_value) {
						errors = append(errors, fmt.Errorf("order by value is nil"))
						continue
					}

					order_by_string_value_validated := ""
					if *order_by_string_value == "ascending" {
						order_by_string_value_validated = "ASC"
					} else if *order_by_string_value == "decending" {
						order_by_string_value_validated = "DESC"
					} else {
						errors = append(errors, fmt.Errorf("order by value is is not valid %s must be ascending or decending", *order_by_string_value))
						continue
					}
					
					if options.IsBoolTrue("use_file") {
						order_by_clause += "`"
					} else {
						order_by_clause += "\\`"
					}
					escaped_order_by_column_name, escaped_order_by_column_name_errors := common.EscapeString(order_by_column_name, "'")
					if escaped_order_by_column_name_errors != nil {
						errors = append(errors, escaped_order_by_column_name_errors)
						continue
					}

					order_by_clause += escaped_order_by_column_name
					if options.IsBoolTrue("use_file") {
						order_by_clause += "`"
					} else {
						order_by_clause += "\\`"
					}

					order_by_clause += " "

					order_by_clause += order_by_string_value_validated

					if order_by_index < (order_by_columns - 1) {
						order_by_clause += ", "
					} else {
						order_by_clause += " "
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			sql_command := "SELECT "
			if select_fields != nil && len(*(select_fields.GetValues())) > 0 {
				select_fields_values_length := len(*(select_fields.GetValues()))
				for i, _ := range *(select_fields.GetValues()) {
					select_fields_value, select_fields_value_errors := select_fields.GetStringValue(i)
					if select_fields_value_errors != nil {
						errors = append(errors, select_fields_value_errors...)
						continue
					}

					escape_string_value, escape_string_value_errors := common.EscapeString(select_fields_value, "'")
					if escape_string_value_errors != nil {
						errors = append(errors, escape_string_value_errors)
					} else {
						sql_command += escape_string_value
						if i < (select_fields_values_length - 1) {
							sql_command += ", "
						} else {
							sql_command += " "
						}
					}
				}
			} else {
				sql_command += "* "
			}

			sql_command += "FROM "
			
			if options.IsBoolTrue("use_file") {
				sql_command += fmt.Sprintf("`%s` ", table_name_escaped)
			} else {
				sql_command += fmt.Sprintf("\\`%s\\` ", table_name_escaped)
			}

			if filters != nil {
				if len(filters.Keys()) > 0 {
					sql_command += "WHERE "
				}

				column_name_params := json.NewMapValue()
				column_name_params.SetObjectForMap("values", column_name_whitelist_characters)
				column_name_params.SetNil("value")
				column_name_params.SetStringValue("label","column_name")
				column_name_params.SetStringValue("data_type", "Table")
				for index, column_filter := range filters.Keys() {
					
					column_definition, column_definition_errors := table_schema.GetMap(column_filter)
					if column_definition_errors != nil {
						errors = append(errors, column_definition_errors...) 
						continue
					}
					
					column_name_params.SetString("value", &column_filter)
					column_name_errors := WhitelistCharacters(column_name_params)
					if column_name_errors != nil {
						errors = append(errors, column_name_errors...)
					}	

					column_filter_escaped, column_filter_escaped_errors := common.EscapeString(column_filter, "'")
					if table_name_escaped_errors != nil {
						errors = append(errors, column_filter_escaped_errors)
					}

					if options.IsBoolTrue("use_file") {
						sql_command += "`"
					} else {
						sql_command += "\\`"
					}
					sql_command += column_filter_escaped
					if options.IsBoolTrue("use_file") {
						sql_command += "`"
					} else {
						sql_command += "\\`"
					}

					if common.IsNil(filters_logic) {
						sql_command += " = "
					} else if !filters_logic.HasKey(column_filter) {
						sql_command += " = "
					} else {
						filters_logic_temp, filters_logic_temp_errors := filters_logic.GetString(column_filter)
						if filters_logic_temp_errors != nil {
							errors = append(errors, filters_logic_temp_errors...)
						} else if common.IsNil(filters_logic_temp) {
							errors = append(errors, fmt.Errorf("filters logic is nil"))
						} else {
							sql_command += " "
							sql_command += *filters_logic_temp
							sql_command += " "
						}
					}

					if filters.IsNil(column_filter) {
						sql_command += "NULL "
					} else {
						//todo check data type with schema
						type_of := filters.GetType(column_filter)
						column_data := filters.GetValue(column_filter)
						switch type_of {
						case "*uint64":
							value, value_errors := column_data.GetUInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(*value, 10)
							}
						case "uint64":
							value, value_errors := column_data.GetUInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(*value, 10)
							}
						case "*int64":
							value, value_errors := column_data.GetInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "int64":
							value, value_errors := column_data.GetInt64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "*uint32":
							value, value_errors := column_data.GetUInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "uint32":
							value, value_errors := column_data.GetUInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "*int32":
							value, value_errors := column_data.GetInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=strconv.FormatInt(int64(*value), 10)
							}
						case "int32":
							value, value_errors := column_data.GetInt32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=strconv.FormatInt(int64(*value), 10)
							}
						case "*uint16":
							value, value_errors := column_data.GetUInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "uint16":
							value, value_errors := column_data.GetUInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "*int16":
							value, value_errors := column_data.GetInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "int16":
							value, value_errors := column_data.GetInt16()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "*uint8":
							value, value_errors := column_data.GetUInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "uint8":
							value, value_errors := column_data.GetUInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatUint(uint64(*value), 10)
							}
						case "*int8":
							value, value_errors := column_data.GetInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "int8":
							value, value_errors := column_data.GetInt8()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "*int":
							value, value_errors := column_data.GetInt()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "int":
							value, value_errors := column_data.GetInt()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command += strconv.FormatInt(int64(*value), 10)
							}
						case "float32":
							value, value_errors := column_data.GetFloat32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=  fmt.Sprintf("%f", *value)
							}
						case "*float32":
							value, value_errors := column_data.GetFloat32()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=  fmt.Sprintf("%f", *value)
							}
						case "float64":
							value, value_errors := column_data.GetFloat64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=  fmt.Sprintf("%f", *value)
							}
						case "*float64":
							value, value_errors := column_data.GetFloat64()
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else if common.IsNil(value) {
								errors = append(errors,fmt.Errorf("*uint64 is nil"))
							} else {
								sql_command +=  fmt.Sprintf("%f", *value)
							}
						case "*time.Time":
							decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
							if decimal_places_error != nil {
								errors = append(errors, decimal_places_error...)
							} else if decimal_places == nil {
								errors = append(errors, fmt.Errorf("decimal_places is nil"))
							} else {
								value, value_errors := column_data.GetTimeWithDecimalPlaces(*decimal_places)
								if value_errors != nil {
									errors = append(errors, value_errors...)
									continue
								} else if common.IsNil(value) {
									errors = append(errors, fmt.Errorf("time is nil"))
									continue
								}

								format_time, format_time_errors := common.FormatTime(*value, *decimal_places)
								if format_time_errors != nil {
									errors = append(errors, format_time_errors...)
								} else if format_time == nil { 
									errors = append(errors, fmt.Errorf("format time is nil"))
								} else {
									value_escaped, value_escaped_errors := common.EscapeString(*format_time, "'")
									if value_escaped_errors != nil {
										errors = append(errors, value_escaped_errors)
									}
			
									if options.IsBoolTrue("use_file") {
										sql_command += "'" + value_escaped + "'"
									} else {
										sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
									}
								}
							}
						case "time.Time":
							decimal_places, decimal_places_error := column_definition.GetInt("decimal_places")
							if decimal_places_error != nil {
								errors = append(errors, decimal_places_error...)
							} else if decimal_places == nil {
								errors = append(errors, fmt.Errorf("decimal_places is nil"))
							} else {
								value, value_errors := column_data.GetTimeWithDecimalPlaces(*decimal_places)
								if value_errors != nil {
									errors = append(errors, value_errors...)
									continue
								} else if common.IsNil(value) {
									errors = append(errors, fmt.Errorf("time is nil"))
									continue
								}

								format_time, format_time_errors := common.FormatTime(*value, *decimal_places)
								if format_time_errors != nil {
									errors = append(errors, format_time_errors...)
								} else if format_time == nil { 
									errors = append(errors, fmt.Errorf("format time is nil"))
								} else {
									value_escaped, value_escaped_errors := common.EscapeString(*format_time, "'")
									if value_escaped_errors != nil {
										errors = append(errors, value_escaped_errors)
									}
			
									if options.IsBoolTrue("use_file") {
										sql_command += "'" + value_escaped + "'"
									} else {
										sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
									}
								}
							}
						case "string":
							column_data_value, column_data_value_errors := column_data.GetString()
							if column_data_value_errors != nil {
								errors = append(errors, column_data_value_errors...)
								continue
							} else if common.IsNil(column_data_value) {
								errors = append(errors, fmt.Errorf("column data is nil"))
								continue
							}

							value_escaped, value_escaped_errors := common.EscapeString(*column_data_value, "'")
							if value_escaped_errors != nil {
								errors = append(errors, value_escaped_errors)
							}
							
							if options.IsBoolTrue("use_file") {
								sql_command += "'" + value_escaped + "'"
							} else {
								sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
							}
							
						case "*string":
							column_data_value, column_data_value_errors := column_data.GetString()
							if column_data_value_errors != nil {
								errors = append(errors, column_data_value_errors...)
								continue
							} else if common.IsNil(column_data_value) {
								errors = append(errors, fmt.Errorf("column data is nil"))
								continue
							}

							value_escaped, value_escaped_errors := common.EscapeString(*column_data_value, "'")
							if value_escaped_errors != nil {
								errors = append(errors, value_escaped_errors)
							}
		
							if options.IsBoolTrue("use_file") {
								sql_command += "'" + value_escaped + "'"
							} else {
								sql_command += strings.ReplaceAll("'" + value_escaped + "'", "`", "\\`")
							}
						
						case "bool":

							if column_data.IsBoolTrue() {
								sql_command += "1"
							} else {
								sql_command += "0"
							}
						case "*bool":
							if column_data.IsBoolTrue() {
								sql_command += "1"
							} else {
								sql_command += "0"
							}
						default:
							errors = append(errors, fmt.Errorf("error: Table.ReadRecords: filter type not supported please implement: %s", type_of))
						}
						sql_command += " "
					}

					if index < len(filters.Keys()) - 1 {
						sql_command += "AND "
					}
				}
			}

			if order_by_clause != "" {
				sql_command += "ORDER BY " + (order_by_clause + " ")
			}

			if limit != nil {
				limit_value := strconv.FormatUint(*limit, 10)
				sql_command += fmt.Sprintf("LIMIT %s ", limit_value)
			}

			if offset != nil {
				offset_value := strconv.FormatUint(*offset, 10)
				sql_command += fmt.Sprintf("OFFSET %s ", offset_value)
			}
			sql_command += ";"

			if len(errors) > 0 {
				return nil, errors
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
				cachable_records, cachable_records_errors := temp_client_manager.GetOrSetReadRecords(*temp_database, sql_command, nil)
				if cachable_records_errors != nil {
					return nil, cachable_records_errors
				} else if !common.IsNil(cachable_records) {
					return cachable_records, nil
				}
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, &sql_command, options)

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, current_json := range *(json_array.GetValues()) {
				current_record, current_record_errors := current_json.GetMap()
				if current_record_errors != nil {
					errors = append(errors, current_record_errors...)
					continue
				} else if common.IsNil(current_record) {
					errors = append(errors, fmt.Errorf("record is nil"))
					continue
				}

				columns := current_record.Keys()
				mapped_record := json.NewMapValue()
				for _, column := range columns {
					table_schema_column_map, table_schema_column_map_errors := table_schema.GetMap(column)
					if table_schema_column_map_errors != nil {
						errors = append(errors, table_schema_column_map_errors...)
						continue
					}
					
					table_data_type, table_data_type_errors := table_schema_column_map.GetString("type")
					if table_data_type_errors != nil {
						errors = append(errors, table_data_type_errors...)
						continue
					}

					switch *table_data_type {
					case "*uint64":
						value, value_errors := current_record.GetUInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetUInt64(column, value)
						}
					case "uint64":
						value, value_errors := current_record.GetUInt64Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt64Value(column, value)
						}
					case "*uint32":
						value, value_errors := current_record.GetUInt32(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetUInt32(column, value)
						}
					case "uint32":
						value, value_errors := current_record.GetUInt32Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt32Value(column, value)
						}
					case "*uint16":
						value, value_errors := current_record.GetUInt16(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetUInt16(column, value)
						}
					case "uint16":
						value, value_errors := current_record.GetUInt16Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt16Value(column, value)
						}
					case "*uint8":
						value, value_errors := current_record.GetUInt8(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetUInt8(column, value)
						}
					case "uint8":
						value, value_errors := current_record.GetUInt8Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt8Value(column, value)
						}
					case "*int64":
						value, value_errors := current_record.GetInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetInt64(column, value)
						}
					case "int64":
						value, value_errors := current_record.GetInt64Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt64Value(column, value)
						}
					case "*int32":
						value, value_errors := current_record.GetInt32(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetInt32(column, value)
						}
					case "int32":
						value, value_errors := current_record.GetInt32Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt32Value(column, value)
						}
					case "*int16":
						value, value_errors := current_record.GetInt16(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetInt16(column, value)
						}
					case "int16":
						value, value_errors := current_record.GetInt16Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt16Value(column, value)
						}
					case "*int8":
						value, value_errors := current_record.GetInt8(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetInt8(column, value)
						}
					case "int8":
						value, value_errors := current_record.GetInt8Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt8Value(column, value)
						}
					case "*time.Time", "time.Time":
						decimal_places, decimal_places_errors := table_schema_column_map.GetInt("decimal_places")
						if decimal_places_errors != nil {
							errors = append(errors, decimal_places_errors...)
						} else if common.IsNil(decimal_places) {
							errors = append(errors, fmt.Errorf("decimal places is nil"))
						} else {
							value, value_errors := current_record.GetTimeWithDecimalPlaces(column, *decimal_places)
							if value_errors != nil {
								errors = append(errors, value_errors...)
							} else {
								mapped_record.SetTime(column, value)
							}
						}
					case "*bool", "bool":
						value, value_errors := current_record.GetBool(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetBool(column, value)
						}
					case "*string":
						value, value_errors := current_record.GetString(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetString(column, value)
						}
					case "string":
						value, value_errors := current_record.GetStringValue(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetStringValue(column, value)
						}
					case "*float32":
						value, value_errors := current_record.GetFloat32(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetFloat32(column, value)
						}
					case "float32":
						value, value_errors := current_record.GetFloat32Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetFloat32Value(column, value)
						}
					case "*float64":
						value, value_errors := current_record.GetFloat64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else if value == nil {
							mapped_record.SetNil(column)
						} else {
							mapped_record.SetFloat64(column, value)
						}
					case "float64":
						value, value_errors := current_record.GetFloat64Value(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetFloat64Value(column, value)
						}
					default:
						errors = append(errors, fmt.Errorf("error: SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", temp_table_name, column, *table_data_type))
					}
				}

				mapped_record_obj, mapped_record_obj_errors := newRecord(*getTable(), mapped_record, database_reserved_words_obj, column_name_whitelist_characters_obj)
				if mapped_record_obj_errors != nil {
					errors = append(errors, mapped_record_obj_errors...)
				} else {
					mapped_records = append(mapped_records, *mapped_record_obj)
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			if cacheable {
				temp_client_manager.GetOrSetReadRecords(*temp_database, sql_command, &mapped_records)
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
