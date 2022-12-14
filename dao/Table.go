package dao

import (
	"fmt"
	"strconv"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type Table struct {
	Validate              func() []error
	Exists                func() (bool, []error)
	Create                func() []error
	Read                func() []error
	Delete                func() []error
	DeleteIfExists        func() []error
	GetSchema             func() (*json.Map, []error)
	GetAdditionalSchema   func() (*json.Map, []error)
	GetTableName          func() (string)
	SetTableName          func(table_name string) []error
	GetTableColumns       func() (*map[string]bool, []error)
	GetIdentityColumns    func()  (*map[string]bool, []error)
	GetPrimaryKeyColumns  func()  (*map[string]bool, []error)
	GetForeignKeyColumns  func() (*map[string]bool, []error)
	GetNonPrimaryKeyColumns func() (*map[string]bool, []error)
	Count                 func(filter *json.Map, filter_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*uint64, []error)
	CreateRecord          func(record json.Map) (*Record, []error)
	CreateRecords          func(records json.Array) ([]error)
	UpdateRecords          func(records json.Array) ([]error)
	UpdateRecord          func(record *json.Map) ([]error)
	ReadRecords         func(select_fields *json.Array, filter *json.Map, filter_logic *json.Map, order_by *json.Array, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase           func() (Database)
}

func newTable(verify *validate.Validator, database Database, table_name string, user_defined_schema *json.Map, schema_from_database *json.Map) (*Table, []error) {
	var errors []error

	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	table_read_records_cache := newTableReadRecordsCache()
	var data *json.Map
	data = json.NewMap()
	
	var this_table *Table

	setTable := func(table *Table) {
		this_table = table
	}

	getTable := func() *Table {
		return this_table
	}

	getData := func() (*json.Map) {
		return data
	}

	setData := func(in_data *json.Map) {
		data = in_data
	}
	
	setupData := func(b Database, n string, user_defined_schema *json.Map, schema_from_database *json.Map) (*json.Map, []error) {
		var setup_errors []error
		d := json.NewMapValue()

		d.SetMap("[fields]", json.NewMap())
		d.SetMap("[system_fields]", json.NewMap())
		d.SetMap("[system_schema]", json.NewMap())

		if user_defined_schema != nil && schema_from_database == nil {
			default_schema := json.NewMap()
			
			map_active_schema := json.NewMap()
			map_active_schema.SetStringValue("type", "bool")
			map_active_schema.SetBoolValue("default", true)
			default_schema.SetMap("enabled", map_active_schema)

			map_archieved_schema := json.NewMap()
			map_archieved_schema.SetStringValue("type", "bool")
			map_archieved_schema.SetBoolValue("default", false)
			default_schema.SetMap("archieved", map_archieved_schema)

			map_created_date_schema := json.NewMap()
			map_created_date_schema.SetStringValue("type", "time.Time")
			map_created_date_schema.SetStringValue("default", "now")
			map_created_date_schema.SetUInt8Value("decimal_places", uint8(6))
			default_schema.SetMap("created_date", map_created_date_schema)

			map_last_modified_date_schema := json.NewMap()
			map_last_modified_date_schema.SetStringValue("type", "time.Time")
			map_last_modified_date_schema.SetStringValue("default", "now")
			map_last_modified_date_schema.SetUInt8Value("decimal_places", uint8(6))
			default_schema.SetMap("last_modified_date", map_last_modified_date_schema)

			map_archieved_date_date_schema := json.NewMap()
			map_archieved_date_date_schema.SetStringValue("type", "time.Time")
			map_archieved_date_date_schema.SetStringValue("default", "zero")
			map_archieved_date_date_schema.SetUInt8Value("decimal_places", uint8(6))
			default_schema.SetMap("archieved_date", map_archieved_date_date_schema)
		
			for _, column_name := range user_defined_schema.GetKeys() {
				current_schema, current_schema_errors := user_defined_schema.GetMapValue(column_name)
				if current_schema_errors != nil {
					return nil, current_schema_errors
				} 
			
				if !default_schema.HasKey(column_name) {
					default_schema.SetMapValue(column_name, current_schema)
					continue
				} 
			}
			d.SetMap("[schema]", default_schema)
		} else if user_defined_schema == nil &&  schema_from_database != nil {
			d.SetMap("[schema]", schema_from_database)
		} else if user_defined_schema != nil &&  schema_from_database != nil {
			errors = append(errors, fmt.Errorf("can either specify user defined schema or db schema but not both"))
		} else {
			errors = append(errors, fmt.Errorf("cannot create table without a schema"))
		}

		if setup_errors != nil {
			return nil, setup_errors
		}
		return &d, nil
	}

	new_data, new_data_errors := setupData(database, table_name, user_defined_schema, schema_from_database)
	if new_data_errors != nil {
		errors = append(errors, new_data_errors...)
	} else {
		setData(new_data)
	}

	var this_primary_key_columns *map[string]bool = nil
	var this_foreigin_key_columns *map[string]bool = nil
	var this_identity_key_columns *map[string]bool = nil
	var this_table_columns *map[string]bool = nil
	var this_non_primary_key_columns *map[string]bool = nil

	getTableName := func() (string) {
		return table_name
	}

	getTableColumns := func() (*map[string]bool, []error) {
		if !common.IsNil(this_table_columns) {
			return this_table_columns, nil
		}
		var errors []error
		primary_key_columns, primary_key_columns_errors := helper.GetTableColumns(*getData())
		if primary_key_columns_errors != nil {
			errors = append(errors, primary_key_columns_errors...)
		} else if common.IsNil(primary_key_columns) {
			errors = append(errors, fmt.Errorf("primary_key_columns is nil"))
		}
		if len(errors) > 0 {
			return nil, errors
		}
		this_table_columns = primary_key_columns
		return this_table_columns, nil
	}

	getPrimaryKeyColumns := func() (*map[string]bool, []error) {
		if !common.IsNil(this_primary_key_columns) {
			return this_primary_key_columns, nil
		}
		var errors []error
		primary_key_columns, primary_key_columns_errors := helper.GetTablePrimaryKeyColumns(*getData())
		if primary_key_columns_errors != nil {
			errors = append(errors, primary_key_columns_errors...)
		} else if common.IsNil(primary_key_columns) {
			errors = append(errors, fmt.Errorf("primary_key_columns is nil"))
		}
		if len(errors) > 0 {
			return nil, errors
		}
		this_primary_key_columns = primary_key_columns
		return this_primary_key_columns, nil
	}

	getForeignKeyColumns := func() (*map[string]bool, []error) {
		if !common.IsNil(this_foreigin_key_columns) {
			return this_foreigin_key_columns, nil
		}
		var errors []error
		foreigin_key_columns, foreigin_key_columns_errors := helper.GetTableForeignKeyColumns(*getData())
		if foreigin_key_columns_errors != nil {
			errors = append(errors, foreigin_key_columns_errors...)
		} else if common.IsNil(foreigin_key_columns) {
			errors = append(errors, fmt.Errorf("foreigin_key_columns is nil"))
		}
		if len(errors) > 0 {
			return nil, errors
		}
		this_foreigin_key_columns = foreigin_key_columns
		return this_foreigin_key_columns, nil
	}

	getIdentityColumns := func() (*map[string]bool, []error) {
		if !common.IsNil(this_identity_key_columns) {
			return this_identity_key_columns, nil
		}
		var errors []error
		identity_key_columns, identity_key_columns_errors := helper.GetTableIdentityColumns(*getData())
		if identity_key_columns_errors != nil {
			errors = append(errors, identity_key_columns_errors...)
		} else if common.IsNil(identity_key_columns) {
			errors = append(errors, fmt.Errorf("identity_key_columns is nil"))
		}
		if len(errors) > 0 {
			return nil, errors
		}
		this_identity_key_columns = identity_key_columns
		return this_identity_key_columns, nil
	}

	getNonPrimaryKeyColumns := func() (*map[string]bool, []error) {
		if !common.IsNil(this_non_primary_key_columns) {
			return this_non_primary_key_columns, nil
		}
		var errors []error
		non_primary_key_columns, non_primary_key_columns_errors := helper.GetTableNonPrimaryKeyColumns(*getData())
		if non_primary_key_columns_errors != nil {
			errors = append(errors, non_primary_key_columns_errors...)
		} else if common.IsNil(non_primary_key_columns) {
			errors = append(errors, fmt.Errorf("non_primary_key_columns is nil"))
		}
		if len(errors) > 0 {
			return nil, errors
		}
		this_non_primary_key_columns = non_primary_key_columns
		return this_non_primary_key_columns, nil
	}

	validate := func() []error {
		var errors []error
		if database_errors := database.Validate(); database_errors != nil {
			errors = append(errors, database_errors...)
		}

		if table_name_errors := verify.ValidateTableName(table_name); table_name_errors != nil {
			errors = append(errors, table_name_errors...)
		}

		if generic_validation_errors := ValidateData(getData(), "*dao.Table"); generic_validation_errors != nil {
			errors = append(errors, generic_validation_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getDatabase := func() (Database) {
		return database
	}

	executeUnsafeCommand := func(sql_command *string, options *json.Map) (*json.Array, []error) {
		errors := validate()
		if errors != nil {
			return nil, errors
		}
		
		sql_command_results, sql_command_errors := SQLCommand.ExecuteUnsafeCommand(database, sql_command, options)
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

	exists := func() (*bool, []error) {
		options := json.NewMap()
		options.SetBoolValue("use_file", true)
		
		var errors []error
		validate_errors := validate()
		if errors != nil {
			errors = append(errors, validate_errors...)
			return nil, errors
		}
		
		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetCheckTableExistsSQL(verify, table_name, options)
		
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

	delete := func() ([]error) {
		errors := validate()
		if errors != nil {
			return errors
		}
	
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		if len(errors) > 0 {
			return errors
		}

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetDropTableSQL(verify, table_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}
		
		_, sql_errors := executeUnsafeCommand(sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	deleteIfExists := func() ([]error) {
		errors := validate()
		if errors != nil {
			return errors
		}

		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetDropTableIfExistsSQL(verify, table_name, options)
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, sql_errors := executeUnsafeCommand(sql_command, new_options)

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

			record_obj, record_errors := newRecord(verify, *getTable(), *current_map)
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

		_, sql_errors := executeUnsafeCommand(&sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	updateRecord := func(record *json.Map) []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)

		errors := validate()
		if errors != nil {
			return errors
		}

		record_obj, record_errors := newRecord(verify, *getTable(), *record)
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
		

		_, sql_errors := executeUnsafeCommand(sql, options)

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

			record_obj, record_errors := newRecord(verify, *getTable(), *current_map)
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

		_, sql_errors := executeUnsafeCommand(&sql, options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getAdditionalSchema := func() (*json.Map, []error) {
		return database.GetAdditionalTableSchema(table_name)
	}


	getSchema := func() (*json.Map, []error) {
		if schema_from_database != nil {
			return schema_from_database, nil
		}

		var errors []error
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("json_output", true)
	
		sql_command, new_options, sql_command_errors := sql_generator_mysql.GetTableSchemaSQL(verify, table_name, options)
		if sql_command_errors != nil {
			errors = append(errors, sql_command_errors...)
		}

		json_array, sql_errors := executeUnsafeCommand(sql_command, new_options)

		if sql_errors != nil {
			errors = append(errors, sql_errors...)
			return  nil, errors
		}

		temp_schema, schem_errors := sql_generator_mysql.MapTableSchemaFromDB(verify, table_name, json_array)
		if schem_errors != nil {
			errors = append(errors, schem_errors...)
			return nil, errors
		}

		new_data, setup_data_errors := setupData(database, table_name, nil, temp_schema)
		if setup_data_errors != nil {
			errors = append(errors, setup_data_errors...)
			return nil, errors
		}
		setData(new_data)
		schema_from_database = temp_schema
		return temp_schema, nil
	}

	setTableName := func(new_table_name string) []error {
		if new_table_name_errors := verify.ValidateTableName(new_table_name); new_table_name_errors != nil {
			return new_table_name_errors
		}
		table_name = new_table_name
		return nil
	}

	getCreateTableSQL := func(options *json.Map) (*string, *json.Map, []error) {	
		return sql_generator_mysql.GetCreateTableSQL(verify, table_name, *getData(), options)
	}

	createTable := func() []error {
		options := json.NewMap()
		options.SetBoolValue("use_file", false)

		sql_command, new_options, sql_command_errors := getCreateTableSQL(options)

		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors := executeUnsafeCommand(sql_command, new_options)

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

		temp_schema, temp_schema_errors := getSchema()
		if temp_schema_errors != nil {
			errors = append(errors, temp_schema_errors...)
		} else if common.IsNil(temp_schema) {
			errors = append(errors, fmt.Errorf("error: Table.read schema is nil"))
		}

		if len(errors) > 0 {
			return errors
		}
		return nil
	}

	x := Table{
		Validate: func() []error {
			return validate()
		},
		GetDatabase: func() (Database) {
			return getDatabase()
		},
		GetTableColumns: func() (*map[string]bool, []error) {
			return getTableColumns()
		},
		GetIdentityColumns: func() (*map[string]bool, []error) {
			return getIdentityColumns()
		},
		GetPrimaryKeyColumns: func() (*map[string]bool, []error) {
			return getPrimaryKeyColumns()
		},
		GetForeignKeyColumns: func() (*map[string]bool, []error) {
			return getForeignKeyColumns()
		},
		GetNonPrimaryKeyColumns: func() (*map[string]bool, []error) {
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
			sql_command, new_options, sql_command_errors := sql_generator_mysql.GetSelectRecordsSQL(verify, table_name, *getData(), select_fields, filters, filters_logic, order_by, limit, offset, options)
			if sql_command_errors != nil {
				return nil, sql_command_errors
			}

			json_array, sql_errors := executeUnsafeCommand(sql_command, new_options)

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

			record, record_errors := newRecord(verify, *getTable(), new_record_data)
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
			sql_command, options, sql_command_errors := sql_generator_mysql.GetSelectRecordsSQL(verify, table_name, *getData(), select_fields, filters, filters_logic, order_by, limit, offset, options)
			if sql_command_errors != nil {
				return nil, sql_command_errors
			}

			additional_schema, additional_schema_errors := getAdditionalSchema()
			if additional_schema_errors != nil {
				return nil, additional_schema_errors
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
				cachable_records, cachable_records_errors := table_read_records_cache.GetOrSetReadRecords(*getTable(), *sql_command, nil)
				if cachable_records_errors != nil {
					return nil, cachable_records_errors
				} else if !common.IsNil(cachable_records) {
					return cachable_records, nil
				}
			}

			json_array, sql_errors := executeUnsafeCommand(sql_command, options)

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, current_json := range *(json_array.GetValues()) {
				mapped_record_obj, mapped_record_obj_errors := mapValueFromDBToRecord(verify, *getTable(), current_json)
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
				table_read_records_cache.GetOrSetReadRecords(*getTable(), *sql_command, &mapped_records)
			}

			return &mapped_records, nil
		},
		UpdateRecords: func(records json.Array) ([]error) {
			return updateRecords(records)
		},
		UpdateRecord: func(record *json.Map) ([]error) {
			return updateRecord(record)
		},
		CreateRecords: func(records json.Array) ([]error) {
			return createRecords(records)
		},
		Exists: func() (bool, []error) {
			table_exists, table_exists_errors := exists()
			if table_exists_errors != nil {
				return false, table_exists_errors
			} else if common.IsNil(table_exists) {
				var table_exists_errors []error
				errors = append(errors, fmt.Errorf("exists returned nil")) 
				return false, table_exists_errors
			}
			return *table_exists, nil
		},
		GetSchema: func() (*json.Map, []error) {
			return getSchema()
		},
		GetAdditionalSchema: func() (*json.Map, []error) {
			return getAdditionalSchema()
		},
		GetTableName: func() (string) {
			return getTableName()
		},
		SetTableName: func(table_name string) []error {
			return setTableName(table_name)
		},
	}
	setTable(&x)

	validate_errors := x.Validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	
	return &x, nil
}
