package dao

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

func mapValueFromDBToRecord(table *Table, current_json *json.Value, database_reserved_words_obj *validation_constants.DatabaseReservedWords, column_name_whitelist_characters_obj *validation_constants.ColumnNameCharacterWhitelist) (*Record, []error) {
	var errors []error

	if common.IsNil(table) {
		errors = append(errors, fmt.Errorf("table is nil"))
	}

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

	table_name, table_name_errors := table.GetTableName()
	if table_name_errors != nil {
		errors = append(errors, table_name_errors...)
	} else if common.IsNil(table_name) {
		errors = append(errors, fmt.Errorf("table_name is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := current_record.GetKeys()
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
			errors = append(errors, fmt.Errorf("error: SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", table_name, column, *table_data_type))
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	mapped_record_obj, mapped_record_obj_errors := newRecord(*table, mapped_record, database_reserved_words_obj, column_name_whitelist_characters_obj)
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

type Record struct {
	Validate  func() []error
	Create    func() []error
	Update    func() []error
	GetInt64  func(field string) (*int64, []error)
	SetInt64  func(field string, value *int64) []error
	GetInt32  func(field string) (*int32, []error)
	SetInt32  func(field string, value *int32) []error
	GetInt16  func(field string) (*int16, []error)
	SetInt16  func(field string, value *int16) []error
	GetInt8  func(field string) (*int8, []error)
	SetInt8  func(field string, value *int8) []error
	GetInt64Value  func(field string) (int64, []error)
	SetInt64Value  func(field string, value int64) []error
	GetInt32Value  func(field string) (int32, []error)
	SetInt32Value  func(field string, value int32) []error
	GetInt16Value  func(field string) (int16, []error)
	SetInt16Value  func(field string, value int16) []error
	GetInt8Value  func(field string) (int8, []error)
	SetInt8Value  func(field string, value int8) []error
	GetUInt64 func(field string) (*uint64, []error)
	SetUInt64 func(field string, value *uint64) []error
	GetUInt32 func(field string) (*uint32, []error)
	SetUInt32 func(field string, value *uint32) []error
	GetUInt16 func(field string) (*uint16, []error)
	SetUInt16 func(field string, value *uint16) []error
	GetUInt8 func(field string) (*uint8, []error)
	SetUInt8 func(field string, value *uint8) []error
	GetUInt64Value func(field string) (uint64, []error)
	SetUInt64Value func(field string, value uint64) []error
	GetUInt32Value func(field string) (uint32, []error)
	SetUInt32Value func(field string, value uint32) []error
	GetUInt16Value func(field string) (uint16, []error)
	SetUInt16Value func(field string, value uint16) []error
	GetUInt8Value func(field string) (uint8, []error)
	SetUInt8Value func(field string, value uint8) []error
	GetString func(field string) (*string, []error)
	GetStringValue func(field string) (string, []error)
	SetString func(field string, value *string) []error 
	SetStringValue func(field string, value string) []error 
	GetBool func(field string) (*bool, []error)
	GetBoolValue func(field string) (bool, []error)
	SetBool func(field string, value *bool) []error 
	SetBoolValue func(field string, value bool) []error 
	GetFloat32 func(field string) (*float32, []error)
	GetFloat32Value func(field string) (float32, []error)
	SetFloat32 func(field string, value *float32) []error 
	SetFloat32Value func(field string, value float32) []error 
	GetFloat64 func(field string) (*float64, []error)
	GetFloat64Value func(field string) (float64, []error)
	SetFloat64 func(field string, value *float64) []error 
	SetFloat64Value func(field string, value float64) []error 
	ToJSONString  func(json *strings.Builder) ([]error)
	GetFields func() (*json.Map, []error)
	GetField func(field string, return_type string) (interface{}, []error)
	SetField func(field string, value interface{}) ([]error)
	GetUpdateSQL func() (*string, []error)
	GetCreateSQL func() (*string, *json.Map, []error)
	GetRecordColumns func() (*[]string, []error)
	GetArchieved func() (*bool, []error)
	GetArchievedDate func() (*time.Time, []error)
	GetNonPrimaryKeyColumnsUpdate func() (*[]string, []error)
	GetPrimaryKeyColumns func() (*[]string, []error)
	GetForeignKeyColumns func() (*[]string, []error)
	SetLastModifiedDate func(value *time.Time) []error
	SetArchievedDate func(value *time.Time) []error
}

func newRecord(table Table, record_data json.Map, database_reserved_words_obj *validation_constants.DatabaseReservedWords, column_name_whitelist_characters_obj *validation_constants.ColumnNameCharacterWhitelist) (*Record, []error) {
	var errors []error
	var this *Record

	getThis := func() *Record {
		return this
	}

	setThis := func(r *Record) {
		this = r
	}

	struct_type := "*dao.Record"

	if common.IsNil(record_data) {
		errors = append(errors, fmt.Errorf("error: record_data is nil"))
	}

	table_schema, table_schema_errors := table.GetSchema()
	if table_schema_errors != nil {
		errors = append(errors, table_schema_errors...)
	} else if common.IsNil(table_schema) {
		errors = append(errors, fmt.Errorf("error: table schema is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	//database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	//column_name_whitelist_characters := column_name_whitelist_characters_obj.GetColumnNameCharacterWhitelist()
	

	data := json.NewMapValue()
	data.SetMapValue("[fields]", record_data)
	data.SetMapValue("[schema]", table_schema)

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[table]", table)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()

	// Start table
	
	map_table_schema := json.NewMapValue()
	map_table_schema.SetStringValue("type", "dao.Table")
	map_system_schema.SetMapValue("[table]", map_table_schema)
	// End table

	data.SetMapValue("[system_schema]", map_system_schema)

	schema_column_names := table_schema.GetKeys()
	for _, schema_column_name := range schema_column_names {
		validate_database_column_name_errors := validation_functions.ValidateDatabaseTableColumnName(schema_column_name)
		if validate_database_column_name_errors != nil {
			errors = append(errors, validate_database_column_name_errors...)
		}
	}

	getData := func() (*json.Map) {
		return &data
	}

	getRecordColumns := func() (*[]string, []error) {
		fields_map, fields_map_errors := helper.GetFields(struct_type, getData(), "[fields]")
		if fields_map_errors != nil {
			return nil, fields_map_errors
		}
		columns := fields_map.GetKeys()
		return &columns, nil
	}

	getTable := func() (Table, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[table]", "dao.Table")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("table is nil"))
		}
		if len(errors) > 0 {
			return Table{}, nil
		}
		return temp_value.(Table), nil
	}

	getArchieved := func() (*bool, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]",  "archieved", "*bool")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*bool), nil
	}

	getArchievedDate := func() (*time.Time, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]",  "archieved_date", "*time.Time")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*time.Time), nil
	}

	setLastModifiedDate := func(value *time.Time) []error {
		return helper.SetField(struct_type, getData(), "[schema]", "[fields]", "last_modified_date", value)
	}

	setArchievedDate := func(value *time.Time) []error {
		return helper.SetField(struct_type, getData(), "[schema]", "[fields]", "archieved_date", value)
	}

	getNonPrimaryKeyColumnsUpdate := func() (*[]string, []error) {
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, temp_table_errors
		}

		non_primary_key_columns, non_primary_key_columns_errors := temp_table.GetNonPrimaryKeyColumns()
		if non_primary_key_columns_errors != nil {
			return nil, non_primary_key_columns_errors
		}

		var record_non_primary_key_columns []string
		for _, record_column := range *record_columns {
			if record_column == "created_date" {
				continue
			}

			for _, non_primary_key_column := range *non_primary_key_columns {
				if non_primary_key_column == record_column {
					record_non_primary_key_columns = append(record_non_primary_key_columns, non_primary_key_column)
					break
				}
			}
		}
		return &record_non_primary_key_columns, nil
	}

	getPrimaryKeyColumns := func() (*[]string, []error) {
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, temp_table_errors
		}

		table_primary_key_columns, table_primary_key_columns_errors := temp_table.GetPrimaryKeyColumns()
		if table_primary_key_columns_errors != nil {
			return nil, table_primary_key_columns_errors
		}

		var record_primary_key_columns []string
		for _, record_column := range *record_columns {
			for _, table_primary_key_column := range *table_primary_key_columns {
				if table_primary_key_column == record_column {
					record_primary_key_columns = append(record_primary_key_columns, record_column)
					break
				}
			}
		}
		return &record_primary_key_columns, nil
	}

	getForeignKeyColumns := func() (*[]string, []error) {
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, temp_table_errors
		}

		table_foreign_key_columns, table_foreign_key_columns_errors := temp_table.GetForeignKeyColumns()
		if table_foreign_key_columns_errors != nil {
			return nil, table_foreign_key_columns_errors
		}

		var record_foreign_key_columns []string
		for _, record_column := range *record_columns {
			for _, table_foreign_key_column := range *table_foreign_key_columns {
				if table_foreign_key_column == record_column {
					record_foreign_key_columns = append(record_foreign_key_columns, record_column)
					break
				}
			}
		}
		return &record_foreign_key_columns, nil
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	getUpdateSQL := func() (*string, *json.Map, []error) {
		validate_errors := validate()
		if validate_errors != nil {
			return nil, nil, validate_errors
		}
		
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("transactional", false)
		errors := validate()
	
		if len(errors) > 0 {
			return nil, nil, errors
		}
	
		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, nil, temp_table_errors
		}
	
		return getUpdateRecordSQLMySQL("*dao.Record", &temp_table, getThis(), options)
	}

	getCreateSQL := func() (*string, *json.Map, []error) {
		validate_errors := validate()
		if validate_errors != nil {
			return nil, nil, validate_errors
		}
		
		options := json.NewMap()
		options.SetBoolValue("use_file", false)
		options.SetBoolValue("no_column_headers", true)
		options.SetBoolValue("transactional", false)
		
		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, nil, temp_table_errors
		}

		return getCreateRecordSQLMySQL("*dao.Record", &temp_table, getData(), options)
	}

	created_record := Record{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			errors := validate()
			if errors != nil {
				return errors
			}

			sql, options, create_sql_errors := getCreateSQL()
			if create_sql_errors != nil {
				return create_sql_errors
			}

			temp_table, temp_table_errors := getTable()
			if temp_table_errors != nil {
				return temp_table_errors
			}

			temp_database, temp_database_errors := temp_table.GetDatabase()
			if temp_database_errors != nil {
				return temp_database_errors
			}

			json_array, errors := temp_database.ExecuteUnsafeCommand(sql, options)

			if len(errors) > 0 {
				return errors
			}

			if options.IsBoolTrue("get_last_insert_id") && !options.IsEmptyString("auto_increment_column_name") {
				if len(*(json_array.GetValues())) != 1 {
					errors = append(errors, fmt.Errorf("error: get_last_insert_id not found "))
					return errors
				}

				record_from_db, record_from_db_errors := (*(json_array.GetValues()))[0].GetMap()
				if record_from_db_errors != nil {
					errors = append(errors, record_from_db_errors...)
					return errors
				} else if common.IsNil(record_from_db) {
					errors = append(errors, fmt.Errorf("Record.Create record_from_db is nil"))
					return errors
				}

				last_insert_id, last_insert_id_errors := record_from_db.GetString("LAST_INSERT_ID()")
				if last_insert_id_errors != nil {
					errors = append(errors, last_insert_id_errors...)
					return errors
				} else if common.IsNil(last_insert_id) {
					errors = append(errors, fmt.Errorf("LAST_INSERT_ID() was nil available columns are: %s", record_from_db.GetKeys()))
					return errors
				} 

				last_insert_id_value, count_err := strconv.ParseUint(strings.TrimSpace(*last_insert_id), 10, 64)
				if count_err != nil {
					errors = append(errors, count_err)
					return errors
				}

				if !options.IsNull("auto_increment_column_name") && !options.IsEmptyString("auto_increment_column_name") {
					auto_increment_column_name, auto_increment_column_name_errors := options.GetString("auto_increment_column_name")
					if auto_increment_column_name_errors != nil {
						errors = append(errors, auto_increment_column_name_errors...)
					} else if common.IsNil(auto_increment_column_name) {
						errors = append(errors, fmt.Errorf("error: auto_increment_column_name is nil"))
					}

					set_auto_field_errors := helper.SetField(struct_type, getData(), "[schema]", "[fields]", *auto_increment_column_name, &last_insert_id_value)
					if set_auto_field_errors != nil {
						errors = append(errors, set_auto_field_errors...)
					}
				}
			}
			
			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		GetUpdateSQL: func() (*string, []error) {
			//todo push options up higher to hide sensitive info if needed
			sql, _, generate_sql_errors := getUpdateSQL()
			if generate_sql_errors != nil {
				return nil, generate_sql_errors
			}
			return sql, nil
		},
		GetCreateSQL: func() (*string, *json.Map, []error) {
			return getCreateSQL()
		},
		Update: func() []error {
			sql, options, generate_sql_errors := getUpdateSQL()
			if generate_sql_errors != nil {
				return generate_sql_errors
			}

			temp_table, temp_table_errors := getTable()
			if temp_table_errors != nil {
				return temp_table_errors
			}

			temp_database, temp_database_errors := temp_table.GetDatabase()
			if temp_database_errors != nil {
				return temp_database_errors
			}

			_, execute_errors := temp_database.ExecuteUnsafeCommand(sql, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int64), nil
		},
		GetInt64Value: func(field string) (int64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "int64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int64), nil
		},
		GetInt32: func(field string) (*int32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int32), nil
		},
		GetInt32Value: func(field string) (int32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "int32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int32), nil
		},
		GetInt16: func(field string) (*int16, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int16), nil
		},
		GetInt16Value: func(field string) (int16, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "int16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int16), nil
		},
		GetInt8: func(field string) (*int8, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int8), nil
		},
		GetInt8Value: func(field string) (int8, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "int8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int8), nil
		},
		SetInt64: func(field string, value *int64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt64Value: func(field string, value int64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32: func(field string, value *int32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32Value: func(field string, value int32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16: func(field string, value *int16) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16Value: func(field string, value int16) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8: func(field string, value *int8) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8Value: func(field string, value int8) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint64), nil
		},
		GetUInt64Value: func(field string) (uint64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint64), nil
		},
		GetUInt32: func(field string) (*uint32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint32), nil
		},
		GetUInt32Value: func(field string) (uint32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint32), nil
		},
		GetUInt16: func(field string) (*uint16, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint16), nil
		},
		GetUInt16Value: func(field string) (uint16, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint16), nil
		},
		GetUInt8: func(field string) (*uint8, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint8), nil
		},
		GetUInt8Value: func(field string) (uint8, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint8), nil
		},
		SetUInt64: func(field string, value *uint64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt64Value: func(field string, value uint64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32: func(field string, value *uint32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32Value: func(field string, value uint32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16: func(field string, value *uint16) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16Value: func(field string, value uint16) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8: func(field string, value *uint8) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8Value: func(field string, value uint8) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetString: func(field string) (*string, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*string")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*string), nil
		},
		GetStringValue: func(field string) (string, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "string")
			if field_value_errors != nil {
				return "", field_value_errors
			}
			return field_value.(string), nil
		},
		SetString: func(field string, value *string) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetStringValue: func(field string, value string) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetBool: func(field string) (*bool, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*bool")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*bool), nil
		},
		GetBoolValue: func(field string) (bool, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "bool")
			if field_value_errors != nil {
				return false, field_value_errors
			}
			return field_value.(bool), nil
		},
		SetBool: func(field string, value *bool) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetBoolValue: func(field string, value bool) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat32: func(field string) (*float32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*float32")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float32), nil
		},
		GetFloat32Value: func(field string) (float32, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "float32")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float32), nil
		},
		SetFloat32: func(field string, value *float32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat32Value: func(field string, value float32) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat64: func(field string) (*float64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "*float64")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float64), nil
		},
		GetFloat64Value: func(field string) (float64, []error) {
			field_value, field_value_errors := helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, "float64")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float64), nil
		},
		SetFloat64: func(field string, value *float64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat64Value: func(field string, value float64) []error {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			fields_map, fields_map_errors := helper.GetFields(struct_type, getData(), "[fields]")
			if fields_map_errors != nil {
				return fields_map_errors
			}
			return fields_map.ToJSONString(json)
		},
		GetFields: func() (*json.Map, []error) {
			fields_map, fields_map_errors := helper.GetFields(struct_type, getData(), "[fields]")
			if fields_map_errors != nil {
				return nil, fields_map_errors
			}
			return fields_map, nil
		},
		GetRecordColumns: func() (*[]string, []error) {
			return getRecordColumns()
		},
		GetArchieved: func() (*bool, []error) {
			return getArchieved()
		},
		GetArchievedDate: func() (*time.Time, []error) {
			return getArchievedDate()
		},
		GetNonPrimaryKeyColumnsUpdate: func() (*[]string, []error) {
			return getNonPrimaryKeyColumnsUpdate()
		},
		GetPrimaryKeyColumns: func() (*[]string, []error) {
			return getPrimaryKeyColumns()
		},
		GetForeignKeyColumns: func() (*[]string, []error) {
			return getForeignKeyColumns()
		},
		GetField: func(field string, return_type string) (interface{}, []error) {
			return helper.GetField(struct_type, getData(), "[schema]", "[fields]", field, return_type)
		},
		SetField: func(field string, value interface{}) ([]error) {
			return helper.SetField(struct_type, getData(), "[schema]", "[fields]", field,  value)
		},
		SetLastModifiedDate: func(value *time.Time) []error {
			return setLastModifiedDate(value)
		},
		SetArchievedDate: func(value *time.Time) []error {
			return setArchievedDate(value)
		},
	}

	if len(errors) > 0 {
		return nil, errors
	}

	setThis(&created_record)
	return &created_record, nil
}
