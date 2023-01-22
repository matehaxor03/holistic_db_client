package dao

import (
	"fmt"
	"strings"
	"strconv"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_validator/validate"
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
	Count                 func(filter *json.Map, filter_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64) (*uint64, []error)
	CreateRecord          func(record json.Map) (*Record, []error)
	CreateRecordAsync     func(record json.Map) ([]error)

	CreateRecords          func(records json.Array) ([]error)
	UpdateRecords          func(records json.Array) ([]error)
	UpdateRecord          func(record *json.Map) ([]error)
	ReadRecords         func(select_fields *json.Array, filter *json.Map, filter_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase           func() (Database)
}

func newTable(verify *validate.Validator, database Database, table_name string, user_defined_schema *json.Map, schema_from_database *json.Map,  sql_command SQLCommand) (*Table, []error) {
	var errors []error

	mysql_wrapper := sql_generator_mysql.NewMySQL()

	table_read_records_cache := newTableReadRecordsCache()
	var data *json.Map
	data = json.NewMap()
	
	var this_table *Table
	var this_schema_from_database *json.Map = nil
	var this_additional_schema_from_database *json.Map = nil

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

	executeUnsafeCommand := func(sql_command_builder strings.Builder, options json.Map) (json.Array, []error) {
		var errors []error
		sql_command_results, sql_command_errors := sql_command.ExecuteUnsafeCommand(database, sql_command_builder, options)
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

	exists := func() (bool, []error) {
		return database.TableExists(getTableName())
	}

	delete := func() ([]error) {
		return database.DeleteTableByTableName(getTableName())
	}

	deleteIfExists := func() ([]error) {
		return database.DeleteTableByTableNameIfExists(getTableName())
	}
	
	updateRecords := func(records json.Array) []error {
		options := json.NewMapValue()
		options.SetBoolValue("transactional", false)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

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

			record_obj, record_errors := newRecord(verify, *getTable(), *current_map, sql_command)
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

		var sql strings.Builder
		for _, record_obj := range records_obj {
			sql_update_snippet, sql_update_snippet_errors := record_obj.GetUpdateSQL()
			if sql_update_snippet_errors != nil {
				errors = append(errors, sql_update_snippet_errors...)
			} else {
				sql.WriteString(sql_update_snippet.String())
				sql.WriteString("\n")
			}
		}

		if len(errors) > 0 {
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

	updateRecord := func(record *json.Map) []error {
		options := json.NewMapValue()
		options.SetBoolValue("transactional", false)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		var errors []error
		record_obj, record_errors := newRecord(verify, *getTable(), *record, sql_command)
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
		

		_, sql_errors := executeUnsafeCommand(*sql, options)

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
		options.SetBoolValue("transactional", false)
		options.SetBoolValue("get_last_insert_id", false)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("no_column_headers", true)

		var errors []error
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

			record_obj, record_errors := newRecord(verify, *getTable(), *current_map, sql_command)
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

		var sql strings.Builder
		for _, record_obj := range records_obj {
			sql_update_snippet, _, sql_update_snippet_errors := record_obj.GetCreateSQLAsync()
			if sql_update_snippet_errors != nil {
				errors = append(errors, sql_update_snippet_errors...)
			} else {
				sql.WriteString(sql_update_snippet.String())
				sql.WriteString("\n")
			}
		}

		if len(errors) > 0 {
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

	getAdditionalSchema := func() (*json.Map, []error) {
		if this_additional_schema_from_database != nil {
			return this_additional_schema_from_database, nil
		}
		temp_addtional_schema_from_db, temp_addtional_schema_from_db_errors := database.GetAdditionalTableSchema(table_name)
		if temp_addtional_schema_from_db_errors != nil {
			return nil, temp_addtional_schema_from_db_errors
		}
		this_additional_schema_from_database = temp_addtional_schema_from_db
		return this_additional_schema_from_database, nil
	}

	getSchema := func() (*json.Map, []error) {
		if this_schema_from_database != nil {
			return this_schema_from_database, nil
		}

		temp_schema, temp_schema_errors := database.GetTableSchema(table_name)
		if temp_schema_errors != nil {
			return nil, temp_schema_errors
		}

		new_data, setup_data_errors := setupData(database, table_name, nil, temp_schema)
		if setup_data_errors != nil {
			errors = append(errors, setup_data_errors...)
			return nil, errors
		}
		setData(new_data)
		this_schema_from_database = temp_schema
		return temp_schema, nil
	}

	setTableName := func(new_table_name string) []error {
		if new_table_name_errors := verify.ValidateTableName(new_table_name); new_table_name_errors != nil {
			return new_table_name_errors
		}
		table_name = new_table_name
		return nil
	}

	getCreateTableSQL := func(options json.Map) (*strings.Builder, json.Map, []error) {	
		return mysql_wrapper.GetCreateTableSQL(verify, table_name, *getData(), options)
	}

	createTable := func() []error {
		options := json.NewMapValue()
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		sql_command, new_options, sql_command_errors := getCreateTableSQL(options)

		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, execute_errors := executeUnsafeCommand(*sql_command, new_options)

		if execute_errors != nil {
			return execute_errors
		}

		return nil
	}

	read := func() []error {
		var errors []error
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

	map_record_from_db_to_record := func(table Table, current_json *json.Value) (*Record, []error) {
		var errors []error
		if common.IsNil(current_json) {
			errors = append(errors, fmt.Errorf("current_json is nil"))
		}
		
		current_record, current_record_errors := current_json.GetMap()
		if current_record_errors != nil {
			errors = append(errors, current_record_errors...)
		} else if common.IsNil(current_record) {
			errors = append(errors, fmt.Errorf("record is nil"))
		}
	
		if len(errors) > 0 {
			return nil, errors
		}
	
		table_schema, table_schema_errors := table.GetSchema()
		if table_schema_errors != nil {
			errors = append(errors, table_schema_errors...)
		} else if common.IsNil(table_schema) {
			errors = append(errors, fmt.Errorf("table_schema is nil"))
		}
	
		table_name := table.GetTableName()
	
		if len(errors) > 0 {
			return nil, errors
		}
	
		columns := current_record.GetKeys()
		mapped_record := json.NewMap()
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
				errors = append(errors, fmt.Errorf("error: SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", table_name, column, *table_data_type))
			}
		}
	
		if len(errors) > 0 {
			return nil, errors
		}
	
		mapped_record_obj, mapped_record_obj_errors := newRecord(verify, table, *mapped_record, sql_command)
		if mapped_record_obj_errors != nil {
			errors = append(errors, mapped_record_obj_errors...)
		} else if common.IsNil(mapped_record_obj){
			errors = append(errors, fmt.Errorf("mapped record is nil"))
		} 
		
		if len(errors) > 0 {
			return nil, errors
		}
		
		return mapped_record_obj, nil
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
		Count: func(filters *json.Map, filters_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64) (*uint64, []error) {
			options := json.NewMapValue()
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("read_no_records", false)
			options.SetBoolValue("no_column_headers", false)

			select_fields := json.NewArray()
			select_fields.AppendStringValue("COUNT(*)")
			sql_command, new_options, sql_command_errors := mysql_wrapper.GetSelectRecordsSQL(verify, table_name, *getData(), select_fields, filters, filters_logic, group_by, order_by, limit, offset, options)
			if sql_command_errors != nil {
				return nil, sql_command_errors
			}

			json_array, sql_errors := executeUnsafeCommand(*sql_command, new_options)

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
			return deleteIfExists()
		},
		CreateRecord: func(new_record_data json.Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := newRecord(verify, *getTable(), new_record_data, sql_command)
			if record_errors != nil {
				return nil, record_errors
			}

			create_record_errors := record.Create()
			if create_record_errors != nil {
				return nil, create_record_errors
			}

			return record, nil
		},
		CreateRecordAsync: func(new_record_data json.Map) ([]error) {
			errors := validate()
			if errors != nil {
				return errors
			}

			record, record_errors := newRecord(verify, *getTable(), new_record_data, sql_command)
			if record_errors != nil {
				return record_errors
			}

			create_record_errors := record.CreateAsync()
			if create_record_errors != nil {
				return create_record_errors
			}

			return nil
		},
		ReadRecords: func(select_fields *json.Array, filters *json.Map, filters_logic *json.Map, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64) (*[]Record, []error) {
			options := json.NewMapValue()
			sql_command, options, sql_command_errors := mysql_wrapper.GetSelectRecordsSQL(verify, table_name, *getData(), select_fields, filters, filters_logic, group_by, order_by, limit, offset, options)
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
				cachable_records, cachable_records_errors := table_read_records_cache.GetOrSetReadRecords(*getTable(), sql_command.String(), nil)
				if cachable_records_errors != nil {
					return nil, cachable_records_errors
				} else if !common.IsNil(cachable_records) {
					return cachable_records, nil
				}
			}

			json_array, sql_errors := executeUnsafeCommand(*sql_command, options)

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, current_json := range *(json_array.GetValues()) {
				mapped_record_obj, mapped_record_obj_errors := map_record_from_db_to_record(*getTable(), current_json)
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
				table_read_records_cache.GetOrSetReadRecords(*getTable(), sql_command.String(), &mapped_records)
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
			return exists()
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
