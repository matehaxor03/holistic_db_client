package dao

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type Record struct {
	Validate  func() []error
	Create    func() []error
	CreateAsync func() []error
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
	GetUpdateSQL func() (*strings.Builder, []error)
	GetCreateSQL func() (*strings.Builder, json.Map, []error)
	GetCreateSQLAsync func() (*strings.Builder, json.Map, []error)
	GetRecordColumns func() (*map[string]bool, []error)
	GetArchieved func() (*bool, []error)
	GetArchievedDate func() (*time.Time, []error)
	GetNonPrimaryKeyColumnsUpdate func() (*map[string]bool, []error)
	GetPrimaryKeyColumns func() (*map[string]bool, []error)
	GetForeignKeyColumns func() (*map[string]bool, []error)
	SetLastModifiedDate func(value *time.Time) []error
	SetArchievedDate func(value *time.Time) []error
	GetTable func() (Table)
}

func newRecord(verify *validate.Validator, table Table, record_data json.Map, sql_command SQLCommand) (*Record, []error) {
	var errors []error
	//var this *Record
	
	
	mysql_wrapper := sql_generator_mysql.NewMySQL()
	/*
	getThis := func() *Record {
		return this
	}

	setThis := func(r *Record) {
		this = r
	}*/

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

	data := json.NewMap()
	for _, key := range record_data.GetKeys() {
		if !common.IsValue(record_data.GetObjectForMap(key)) {
			record_data.SetValue(key, json.NewValue(record_data.GetObjectForMap(key)))
		}
	}

	data.SetMap("[fields]", &record_data)
	data.SetMap("[schema]", table_schema)
	data.SetMap("[system_fields]", json.NewMap())
	data.SetMap("[system_schema]", json.NewMap())

	schema_column_names := table_schema.GetKeys()
	for _, schema_column_name := range schema_column_names {
		validate_database_column_name_errors := verify.ValidateColumnName(schema_column_name)
		if validate_database_column_name_errors != nil {
			errors = append(errors, validate_database_column_name_errors...)
		}
	}

	record_column_names := record_data.GetKeys()
	for _, record_column_name := range record_column_names {
		validate_record_column_name_errors := verify.ValidateColumnName(record_column_name)
		if validate_record_column_name_errors != nil {
			errors = append(errors, validate_record_column_name_errors...)
		}
	}

	getData := func() (*json.Map) {
		return data
	}

	getRecordColumns := func() (*map[string]bool, []error) {
		return helper.GetRecordColumns(*getData())
	}

	getTable := func() (Table) {
		return table
	}

	getArchieved := func() (*bool, []error) {
		temp_value, temp_value_errors := helper.GetField(*getData(), "[schema]", "[fields]",  "archieved", "*bool")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*bool), nil
	}

	getArchievedDate := func() (*time.Time, []error) {
		temp_value, temp_value_errors := helper.GetField(*getData(), "[schema]", "[fields]",  "archieved_date", "*time.Time")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*time.Time), nil
	}

	setLastModifiedDate := func(value *time.Time) []error {
		return helper.SetField(*getData(), "[schema]", "[fields]", "last_modified_date", value)
	}

	setArchievedDate := func(value *time.Time) []error {
		return helper.SetField(*getData(), "[schema]", "[fields]", "archieved_date", value)
	}

	getNonPrimaryKeyColumnsUpdate := func() (*map[string]bool, []error) {
		var errors []error
		table_non_primary_key_columns, table_non_primary_key_columns_errors := table.GetNonPrimaryKeyColumns()
		if table_non_primary_key_columns_errors != nil {
			errors = append(errors, table_non_primary_key_columns_errors...)
		} else if common.IsNil(table_non_primary_key_columns) {
			errors = append(errors, fmt.Errorf(struct_type + " table returned nil columns table.GetNonPrimaryKeyColumns()."))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return helper.GetRecordNonPrimaryKeyColumnsUpdate(*getData(), table_non_primary_key_columns)
	}

	getPrimaryKeyColumns := func() (*map[string]bool, []error) {
		var errors []error
		table_primary_key_columns, table_primary_key_columns_errors := table.GetPrimaryKeyColumns()
		if table_primary_key_columns_errors != nil {
			errors = append(errors, table_primary_key_columns_errors...)
		} else if common.IsNil(table_primary_key_columns) {
			errors = append(errors, fmt.Errorf(struct_type + " table returned nil columns table.GetPrimaryKeyColumns()."))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return helper.GetRecordPrimaryKeyColumns(*getData(), table_primary_key_columns)
	}

	getForeignKeyColumns := func() (*map[string]bool, []error) {
		var errors []error
		table_foreign_key_columns, table_foreign_key_columns_errors := table.GetForeignKeyColumns()
		if table_foreign_key_columns_errors != nil {
			errors = append(errors, table_foreign_key_columns_errors...)
		} else if common.IsNil(table_foreign_key_columns) {
			errors = append(errors, fmt.Errorf(struct_type + " table returned nil columns table.GetForeignKeyColumns()."))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return helper.GetRecordForeignKeyColumns(*getData(), table_foreign_key_columns)
	}

	validate := func() []error {
		var errors []error
		if table_validation_errors := table.Validate(); table_validation_errors != nil {
			errors = append(errors, table_validation_errors...)
		}

		if generic_validation_errors := ValidateData(getData(), struct_type); generic_validation_errors != nil {
			errors = append(errors, generic_validation_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	executeUnsafeCommand := func(sql_command_builder strings.Builder, options json.Map) (json.Array, []error) {
		var errors []error
		database := table.GetDatabase()
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

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	getUpdateSQL := func() (*strings.Builder, json.Map, []error) {
		var errors []error
		options := json.NewMapValue()
		options.SetBoolValue("transactional", false)
		options.SetBoolValue("read_no_records", true)
		options.SetBoolValue("get_last_insert_id", false)

		if validate_errors != nil {
			return nil, options, validate_errors
		}
	
		temp_table_schema, temp_table_schema_errors := table.GetSchema()
		if temp_table_schema_errors != nil {
			errors = append(errors, temp_table_schema_errors...)
		} else if common.IsNil(temp_table_schema) {
			errors = append(errors, fmt.Errorf("table schema is nil"))
		}

		temp_table_columns, temp_table_columns_errors := table.GetTableColumns()
		if temp_table_columns_errors != nil {
			errors = append(errors, temp_table_columns_errors...)
		} else if common.IsNil(temp_table_columns) {
			errors = append(errors, fmt.Errorf("table columns is nil"))
		}

		if len(errors) > 0 {
			return nil, options, errors
		}
	
		return mysql_wrapper.GetUpdateRecordSQL(verify, table.GetTableName(), *temp_table_schema, *temp_table_columns, *getData(), options)
	}

	getCreateSQL := func(async bool) (*strings.Builder, json.Map, []error) {
		var errors []error
		options := json.NewMapValue()
		options.SetBoolValue("no_column_headers", false)
		options.SetBoolValue("transactional", false)
		
		if async {
			options.SetBoolValue("get_last_insert_id", false)
			options.SetBoolValue("read_no_records", true)
		} else {
			options.SetBoolValue("get_last_insert_id", true)
			options.SetBoolValue("read_no_records", false)
		}

		temp_table_schema, temp_table_schema_errors := table.GetSchema()
		if temp_table_schema_errors != nil {
			errors = append(errors, temp_table_schema_errors...)
		} else if common.IsNil(temp_table_schema) {
			errors = append(errors, fmt.Errorf("table schema is nil"))
		}

		temp_table_columns, temp_table_columns_errors := table.GetTableColumns()
		if temp_table_columns_errors != nil {
			errors = append(errors, temp_table_columns_errors...)
		} else if common.IsNil(temp_table_columns) {
			errors = append(errors, fmt.Errorf("table columns is nil"))
		}

		if len(errors) > 0 {
			return nil, options, errors
		}
		
	

		return mysql_wrapper.GetCreateRecordSQL(verify, table.GetTableName(), *temp_table_schema, *temp_table_columns, *getData(), options)
	}

	create := func(async bool) []error {
		var errors []error
		sql, new_options, create_sql_errors := getCreateSQL(async)
		if create_sql_errors != nil {
			return create_sql_errors
		}

		json_array, execute_errors := executeUnsafeCommand(*sql, new_options)
		if execute_errors != nil {
			errors = append(errors, execute_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		if new_options.IsBoolFalse("read_no_records") && 
		   new_options.IsBoolTrue("get_last_insert_id") && 
		   !new_options.IsEmptyString("auto_increment_column_name") {
			
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

			if !new_options.IsNull("auto_increment_column_name") && 
			   !new_options.IsEmptyString("auto_increment_column_name") {
				auto_increment_column_name, auto_increment_column_name_errors := new_options.GetString("auto_increment_column_name")
				if auto_increment_column_name_errors != nil {
					errors = append(errors, auto_increment_column_name_errors...)
				} else if common.IsNil(auto_increment_column_name) {
					errors = append(errors, fmt.Errorf("error: auto_increment_column_name is nil"))
				}

				set_auto_field_errors := helper.SetField(*getData(), "[schema]", "[fields]", *auto_increment_column_name, &last_insert_id_value)
				if set_auto_field_errors != nil {
					errors = append(errors, set_auto_field_errors...)
				}
			}
		}
		
		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	created_record := Record{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			return create(false)
		},
		CreateAsync: func() []error {
			return create(true)
		},
		GetUpdateSQL: func() (*strings.Builder, []error) {
			//todo push options up higher to hide sensitive info if needed
			sql, _, generate_sql_errors := getUpdateSQL()
			if generate_sql_errors != nil {
				return nil, generate_sql_errors
			}
			return sql, nil
		},
		GetCreateSQL: func() (*strings.Builder, json.Map, []error) {
			return getCreateSQL(false)
		},
		GetCreateSQLAsync: func() (*strings.Builder, json.Map, []error) {
			return getCreateSQL(true)
		},
		Update: func() []error {
			sql, options, generate_sql_errors := getUpdateSQL()
			if generate_sql_errors != nil {
				return generate_sql_errors
			}

			_, execute_errors := executeUnsafeCommand(*sql, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*int64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int64), nil
		},
		GetInt64Value: func(field string) (int64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "int64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int64), nil
		},
		GetInt32: func(field string) (*int32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*int32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int32), nil
		},
		GetInt32Value: func(field string) (int32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "int32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int32), nil
		},
		GetInt16: func(field string) (*int16, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*int16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int16), nil
		},
		GetInt16Value: func(field string) (int16, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "int16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int16), nil
		},
		GetInt8: func(field string) (*int8, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*int8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int8), nil
		},
		GetInt8Value: func(field string) (int8, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "int8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int8), nil
		},
		SetInt64: func(field string, value *int64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt64Value: func(field string, value int64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32: func(field string, value *int32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32Value: func(field string, value int32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16: func(field string, value *int16) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16Value: func(field string, value int16) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8: func(field string, value *int8) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8Value: func(field string, value int8) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*uint64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint64), nil
		},
		GetUInt64Value: func(field string) (uint64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "uint64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint64), nil
		},
		GetUInt32: func(field string) (*uint32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*uint32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint32), nil
		},
		GetUInt32Value: func(field string) (uint32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "uint32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint32), nil
		},
		GetUInt16: func(field string) (*uint16, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*uint16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint16), nil
		},
		GetUInt16Value: func(field string) (uint16, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "uint16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint16), nil
		},
		GetUInt8: func(field string) (*uint8, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*uint8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint8), nil
		},
		GetUInt8Value: func(field string) (uint8, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "uint8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint8), nil
		},
		SetUInt64: func(field string, value *uint64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt64Value: func(field string, value uint64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32: func(field string, value *uint32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32Value: func(field string, value uint32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16: func(field string, value *uint16) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16Value: func(field string, value uint16) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8: func(field string, value *uint8) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8Value: func(field string, value uint8) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		GetString: func(field string) (*string, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*string")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*string), nil
		},
		GetStringValue: func(field string) (string, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "string")
			if field_value_errors != nil {
				return "", field_value_errors
			}
			return field_value.(string), nil
		},
		SetString: func(field string, value *string) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetStringValue: func(field string, value string) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		GetBool: func(field string) (*bool, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*bool")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*bool), nil
		},
		GetBoolValue: func(field string) (bool, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "bool")
			if field_value_errors != nil {
				return false, field_value_errors
			}
			return field_value.(bool), nil
		},
		SetBool: func(field string, value *bool) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetBoolValue: func(field string, value bool) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat32: func(field string) (*float32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*float32")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float32), nil
		},
		GetFloat32Value: func(field string) (float32, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "float32")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float32), nil
		},
		SetFloat32: func(field string, value *float32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat32Value: func(field string, value float32) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat64: func(field string) (*float64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "*float64")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float64), nil
		},
		GetFloat64Value: func(field string) (float64, []error) {
			field_value, field_value_errors := helper.GetField(*getData(), "[schema]", "[fields]", field, "float64")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float64), nil
		},
		SetFloat64: func(field string, value *float64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat64Value: func(field string, value float64) []error {
			return helper.SetField(*getData(), "[schema]", "[fields]", field, value)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			fields_map, fields_map_errors := helper.GetFields(*getData(), "[fields]")
			if fields_map_errors != nil {
				return fields_map_errors
			}
			return fields_map.ToJSONString(json)
		},
		GetFields: func() (*json.Map, []error) {
			fields_map, fields_map_errors := helper.GetFields(*getData(), "[fields]")
			if fields_map_errors != nil {
				return nil, fields_map_errors
			}
			return fields_map, nil
		},
		GetRecordColumns: func() (*map[string]bool, []error) {
			return getRecordColumns()
		},
		GetArchieved: func() (*bool, []error) {
			return getArchieved()
		},
		GetArchievedDate: func() (*time.Time, []error) {
			return getArchievedDate()
		},
		GetNonPrimaryKeyColumnsUpdate: func() (*map[string]bool, []error) {
			return getNonPrimaryKeyColumnsUpdate()
		},
		GetPrimaryKeyColumns: func() (*map[string]bool, []error) {
			return getPrimaryKeyColumns()
		},
		GetForeignKeyColumns: func() (*map[string]bool, []error) {
			return getForeignKeyColumns()
		},
		GetField: func(field string, return_type string) (interface{}, []error) {
			return helper.GetField(*getData(), "[schema]", "[fields]", field, return_type)
		},
		SetField: func(field string, value interface{}) ([]error) {
			return helper.SetField(*getData(), "[schema]", "[fields]", field,  value)
		},
		SetLastModifiedDate: func(value *time.Time) []error {
			return setLastModifiedDate(value)
		},
		SetArchievedDate: func(value *time.Time) []error {
			return setArchievedDate(value)
		},
		GetTable: func() (Table) {
			return getTable()
		},
	}

	if len(errors) > 0 {
		return nil, errors
	}

	//setThis(&created_record)
	return &created_record, nil
}
