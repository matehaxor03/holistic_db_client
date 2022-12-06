package class

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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
}

func newRecord(table Table, record_data Map, database_reserved_words_obj *DatabaseReservedWords, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Record, []error) {
	var errors []error
	SQLCommand, SQLCommand_errors := newSQLCommand()
	if SQLCommand_errors != nil {
		errors = append(errors, SQLCommand_errors...)
	}

	struct_type := "*class.Record"

	if record_data == nil {
		errors = append(errors, fmt.Errorf("error: record_data is nil"))
	}

	table_schema, table_schema_errors := table.GetSchema()
	if table_schema_errors != nil {
		errors = append(errors, table_schema_errors...)
	} else if IsNil(table_schema) {
		errors = append(errors, fmt.Errorf("error: table schema is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	//database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	//column_name_whitelist_characters := column_name_whitelist_characters_obj.GetColumnNameCharacterWhitelist()
	

	data := Map{"[fields]": record_data, "[system_fields]": Map{"[table]": table}}
	data["[schema]"] = table_schema
	data["[system_schema]"] = Map{"[table]": Map{"type":"class.Table"}}

	schema_column_names := table_schema.Keys()
	for _, schema_column_name := range schema_column_names {
		validate_database_column_name_errors := ValidateDatabaseColumnName(schema_column_name)
		if validate_database_column_name_errors != nil {
			errors = append(errors, validate_database_column_name_errors...)
		}
	}

	getData := func() (*Map) {
		return &data
	}

	getRecordColumns := func() (*[]string, []error) {
		fields_map, fields_map_errors := GetFields(struct_type, getData(), "[fields]")
		if fields_map_errors != nil {
			return nil, fields_map_errors
		}
		columns := fields_map.Keys()
		return &columns, nil
	}

	getTable := func() (*Table, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[table]", "*class.Table")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*Table), temp_value_errors
	}

	getNonIdentityColumnsUpdate := func() (*[]string, []error) {
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, temp_table_errors
		}

		non_identity_columns, non_identity_columns_errors := temp_table.GetNonIdentityColumns()
		if non_identity_columns_errors != nil {
			return nil, non_identity_columns_errors
		}

		var record_non_identity_columns []string
		for _, record_column := range *record_columns {
			if record_column == "created_date" ||
				record_column == "archieved_date" {
				continue
			}

			for _, non_identity_column := range *non_identity_columns {
				if non_identity_column == record_column {
					record_non_identity_columns = append(record_non_identity_columns, non_identity_column)
					break
				}
			}
		}
		return &record_non_identity_columns, nil
	}

	getIdentityColumns := func() (*[]string, []error) {
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		temp_table, temp_table_errors := getTable()
		if temp_table_errors != nil {
			return nil, temp_table_errors
		}

		identity_columns, identity_columns_errors := temp_table.GetIdentityColumns()
		if identity_columns_errors != nil {
			return nil, identity_columns_errors
		}

		var record_identity_columns []string
		for _, record_column := range *record_columns {
			for _, identity_column := range *identity_columns {
				if identity_column == record_column {
					record_identity_columns = append(record_identity_columns, identity_column)
					break
				}
			}
		}
		return &record_identity_columns, nil
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getInsertSQL := func() (*string, Map, []error) {
		options := Map{"use_file": false, "no_column_headers": true, "get_last_insert_id": false}
		errors := validate()

		if len(errors) > 0 {
			return nil, nil, errors
		}

		table, table_errors := getTable()
		if table_errors != nil {
			return nil, nil, table_errors
		}

		table_schema, table_schema_errors := table.GetSchema()
		
		if table_schema_errors != nil {
			return nil, nil, table_schema_errors
		}

		valid_columns, valid_columns_errors := table.GetTableColumns()
		if valid_columns_errors != nil {
			return nil, nil, valid_columns_errors
		}
		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, nil, record_columns_errors
		}

		table_name, table_name_errors := table.GetTableName() 
		if table_name_errors != nil {
			return nil, nil, table_name_errors
		}

		for _, record_column := range *record_columns {
			if strings.HasPrefix(record_column, "credential") {
				options["use_file"] = true
			}
		}

		auto_increment_columns := 0
		for _, valid_column := range *valid_columns {
			column_definition, column_definition_errors := table_schema.GetMap(valid_column)
			if column_definition_errors != nil {
				errors = append(errors, column_definition_errors...) 
				continue
			}

			if !column_definition.IsBool("primary_key") || !column_definition.IsBool("auto_increment") {
				continue
			}

			primary_key_value, primary_key_value_errors := column_definition.GetBool("primary_key")
			if primary_key_value_errors != nil {
				errors = append(errors, primary_key_value_errors...)
				continue
			}

			if *primary_key_value == false {
				continue
			}

			auto_increment_value, auto_increment_value_errors := column_definition.GetBool("auto_increment")
			if auto_increment_value_errors != nil {
				errors = append(errors, auto_increment_value_errors...)
				continue
			}

			if *auto_increment_value == false {
				continue
			}

			options["get_last_insert_id"] = true
			options["auto_increment_column_name"] = valid_column
			auto_increment_columns += 1
		}

		if auto_increment_columns > 1 {
			errors = append(errors, fmt.Errorf("error: table: %s can only have 1 auto_increment primary_key column, found: %s", table_name, auto_increment_columns))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString(fmt.Sprintf("INSERT INTO "))
		sql_command.WriteString(EscapeString(table_name))
		sql_command.WriteString(" (")
		for index, record_column := range *record_columns {
			sql_command.WriteString(EscapeString(record_column))
			if index < (len(*record_columns) - 1) {
				sql_command.WriteString(", ")
			}
		}

		sql_command.WriteString(") VALUES (")
		for index, record_column := range *record_columns {
			column_data, paramter_errors := GetField(struct_type, getData(), "[schema]", "[fields]", record_column, "self")
			if paramter_errors != nil {
				errors = append(errors, paramter_errors...)
				continue
			}

			rep := GetType(column_data)
			switch rep {
			case "*uint64":
				value := column_data.(*uint64)
				sql_command.WriteString(strconv.FormatUint(*value, 10))
			case "uint64":
				value := column_data.(uint64)
				sql_command.WriteString(strconv.FormatUint(value, 10))
			case "*int64":
				value := column_data.(*int64)
				sql_command.WriteString(strconv.FormatInt(*value, 10))
			case "int64":
				value := column_data.(int64)
				sql_command.WriteString(strconv.FormatInt(value, 10))
			case "*uint32":
				value := column_data.(*uint32)
				sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
			case "uint32":
				value := column_data.(uint32)
				sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
			case "*int32":
				value := column_data.(*int32)
				sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
			case "int32":
				value := column_data.(int32)
				sql_command.WriteString(strconv.FormatInt(int64(value), 10))
			case "*uint16":
				value := column_data.(*uint16)
				sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
			case "uint16":
				value := column_data.(uint16)
				sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
			case "*int16":
				value := column_data.(*int16)
				sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
			case "int16":
				value := column_data.(int16)
				sql_command.WriteString(strconv.FormatInt(int64(value), 10))
			case "*uint8":
				value := column_data.(*uint8)
				sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
			case "uint8":
				value := column_data.(uint8)
				sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
			case "*int8":
				value := column_data.(*int8)
				sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
			case "int8":
				value := column_data.(int8)
				sql_command.WriteString(strconv.FormatInt(int64(value), 10))
			case "*int":
				value := column_data.(*int)
				sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
			case "int":
				value := column_data.(int)
				sql_command.WriteString(strconv.FormatInt(int64(value), 10))
			case "float32":
				sql_command.WriteString(fmt.Sprintf("%f", column_data.(float32)))
			case "*float32":
				sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float32))))
			case "float64":
				sql_command.WriteString(fmt.Sprintf("%f", column_data.(float64)))
			case "*float64":
				sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float64))))
			case "*time.Time":
				value := column_data.(*time.Time)
				sql_command.WriteString("'")
				sql_command.WriteString(EscapeString(FormatTime(*value)))
				sql_command.WriteString("'")
			case "time.Time":
				value := column_data.(time.Time)
				sql_command.WriteString("'")
				sql_command.WriteString(EscapeString(FormatTime(value)))
				sql_command.WriteString("'")
			case "string":
				sql_command.WriteString("'")
				sql_command.WriteString(EscapeString(column_data.(string)))
				sql_command.WriteString("'")
			case "*string":
				sql_command.WriteString("'")
				sql_command.WriteString(EscapeString(*(column_data.(*string))))
				sql_command.WriteString("'")
			case "bool":
				if column_data.(bool) {
					sql_command.WriteString("1")
				} else {
					sql_command.WriteString("0")
				}
			case "*bool":
				if *(column_data.(*bool)) {
					sql_command.WriteString("1")
				} else {
					sql_command.WriteString("0")
				}
			default:
				errors = append(errors, fmt.Errorf("error: %s Record.getInsertSQL type: %s not supported for table please implement", struct_type, rep))
			}

			if index < (len(*record_columns) - 1) {
				sql_command.WriteString(", ")
			}
		}
		sql_command.WriteString(");")

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command_string := sql_command.String()
		return &sql_command_string, options, nil
	}

	getUpdateSQL := func() (*string, Map, []error) {
		options := Map{"use_file": false}
		errors := validate()

		if len(errors) > 0 {
			return nil, nil, errors
		}

		table, table_errors := getTable()
		if table_errors != nil {
			return nil, nil, table_errors
		}

		table_name, table_name_errors := table.GetTableName()
		if table_name_errors != nil {
			return nil, nil, table_name_errors
		}

		_, table_schema_errors := table.GetSchema()

		if table_schema_errors != nil {
			return nil, nil, table_schema_errors
		}

		_, valid_columns_errors := table.GetTableColumns()
		if valid_columns_errors != nil {
			return nil, nil, valid_columns_errors
		}

		record_columns, record_columns_errors := getRecordColumns()
		if record_columns_errors != nil {
			return nil, nil, record_columns_errors
		}

		for _, record_column := range *record_columns {
			if strings.HasPrefix(record_column, "credential") {
				options["use_file"] = true
			}
		}

		identity_columns, identity_columns_errors := table.GetIdentityColumns()
		if identity_columns_errors != nil {
			return nil, nil, identity_columns_errors
		}

		table_non_identity_columns, table_non_identity_columns_errors := table.GetNonIdentityColumns()
		if table_non_identity_columns_errors != nil {
			return nil, nil, table_non_identity_columns_errors
		}

		record_identity_columns, record_identity_columns_errors := getIdentityColumns()
		if record_identity_columns_errors != nil {
			return nil, nil, record_identity_columns_errors
		}

		for _, identity_column := range *identity_columns {
			found_identity_column := false
			for _, record_identity_column := range *record_identity_columns {
				if identity_column == record_identity_column {
					found_identity_column = true
				}
			}

			if !found_identity_column {
				errors = append(errors, fmt.Errorf("error: record did not contain identify column: %s", identity_column))
			}
		}

		SetField(struct_type, getData(), "[schema]", "[fields]", "last_modified_date", GetTimeNow())

		record_non_identity_columns, record_non_identity_columns_errors := getNonIdentityColumnsUpdate()
		if record_non_identity_columns_errors != nil {
			return nil, nil, record_non_identity_columns_errors
		}

		if len(*record_non_identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("error: no non-identity columns detected in record to update"))
		}

		if len(*identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("error: table schema has no identity columns"))
		}

		if !Contains(*table_non_identity_columns, "last_modified_date") {
			errors = append(errors, fmt.Errorf("error: table schema does not have last_modified_date"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		var sql_command strings.Builder
		sql_command.WriteString(fmt.Sprintf("UPDATE %s \n", EscapeString(table_name)))

		sql_command.WriteString("SET ")

		for index, record_non_identity_column := range *record_non_identity_columns {
			sql_command.WriteString(EscapeString(record_non_identity_column))
			sql_command.WriteString("=")
			column_data, column_data_errors := GetField(struct_type, getData(), "[schema]", "[fields]", record_non_identity_column, "self")

			if column_data_errors != nil {
				errors = append(errors, column_data_errors...)
				continue
			}

			if IsNil(column_data) {
				sql_command.WriteString("NULL")
			} else {
				rep := GetType(column_data)
				switch rep {
				case "*uint64":
					value := column_data.(*uint64)
					sql_command.WriteString(strconv.FormatUint(*value, 10))
				case "uint64":
					value := column_data.(uint64)
					sql_command.WriteString(strconv.FormatUint(value, 10))
				case "*int64":
					value := column_data.(*int64)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int64":
					value := column_data.(int64)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*uint32":
					value := column_data.(*uint32)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint32":
					value := column_data.(uint32)
					sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
				case "*int32":
					value := column_data.(*int32)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int32":
					value := column_data.(int32)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*uint16":
					value := column_data.(*uint16)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint16":
					value := column_data.(uint16)
					sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
				case "*int16":
					value := column_data.(*int16)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int16":
					value := column_data.(int16)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*uint8":
					value := column_data.(*uint8)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint8":
					value := column_data.(uint8)
					sql_command.WriteString( strconv.FormatUint(uint64(value), 10))
				case "*int8":
					value := column_data.(*int8)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int8":
					value := column_data.(int8)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*int":
					value := column_data.(*int)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int":
					value := column_data.(int)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "float32":
					sql_command.WriteString(fmt.Sprintf("%f", column_data.(float32)))
				case "*float32":
					sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float32))))
				case "float64":
					sql_command.WriteString(fmt.Sprintf("%f", column_data.(float64)))
				case "*float64":
					sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float64))))
				case "*time.Time":
					value := column_data.(*time.Time)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(FormatTime(*value)))
					sql_command.WriteString("'")
				case "time.Time":
					value := column_data.(time.Time)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(FormatTime(value)))
					sql_command.WriteString("'")
				case "*string":
					value := column_data.(*string)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(*value))
					sql_command.WriteString("'")
				case "string":
					value := column_data.(string)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(value))
					sql_command.WriteString("'")
				case "bool":
					if column_data.(bool) {
						sql_command.WriteString("1")
					} else {
						sql_command.WriteString("0")
					}
				case "*bool":
					if *(column_data.(*bool)) {
						sql_command.WriteString("1")
					} else {
						sql_command.WriteString("0")
					}
				default:
					errors = append(errors, fmt.Errorf("error: %s Record.getUpdateSQL type: %s not supported for table please implement", struct_type, rep))
				}
			}

			if index < len(*record_non_identity_columns)-1 {
				sql_command.WriteString(", \n")
			}
		}

		sql_command.WriteString(" WHERE ")
		for index, identity_column := range *identity_columns {
			sql_command.WriteString(EscapeString(identity_column))
			sql_command.WriteString(" = ")
			column_data, column_data_errors := GetField(struct_type, getData(), "[schema]", "[fields]", identity_column, "self")

			if column_data_errors != nil {
				errors = append(errors, column_data_errors...)
				continue
			}

			if IsNil(column_data) {
				sql_command.WriteString("NULL")
			} else {
				record_non_identity_column_type := GetType(column_data)
				switch record_non_identity_column_type {
				case "*uint64":
					value := column_data.(*uint64)
					sql_command.WriteString(strconv.FormatUint(*value, 10))
				case "uint64":
					value := column_data.(uint64)
					sql_command.WriteString(strconv.FormatUint(value, 10))
				case "*int64":
					value := column_data.(*int64)
					sql_command.WriteString(strconv.FormatInt(*value, 10))
				case "int64":
					value := column_data.(int64)
					sql_command.WriteString(strconv.FormatInt(value, 10))
				case "*uint32":
					value := column_data.(*uint32)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint32":
					value := column_data.(uint32)
					sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
				case "*int32":
					value := column_data.(*int32)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int32":
					value := column_data.(int32)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*uint16":
					value := column_data.(*uint16)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint16":
					value := column_data.(uint16)
					sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
				case "*int16":
					value := column_data.(*int16)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int16":
					value := column_data.(int16)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*uint8":
					value := column_data.(*uint8)
					sql_command.WriteString(strconv.FormatUint(uint64(*value), 10))
				case "uint8":
					value := column_data.(uint8)
					sql_command.WriteString(strconv.FormatUint(uint64(value), 10))
				case "*int8":
					value := column_data.(*int8)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int8":
					value := column_data.(int8)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "*int":
					value := column_data.(*int)
					sql_command.WriteString(strconv.FormatInt(int64(*value), 10))
				case "int":
					value := column_data.(int)
					sql_command.WriteString(strconv.FormatInt(int64(value), 10))
				case "float32":
					sql_command.WriteString(fmt.Sprintf("%f", column_data.(float32)))
				case "*float32":
					sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float32))))
				case "float64":
					sql_command.WriteString(fmt.Sprintf("%f", column_data.(float64)))
				case "*float64":
					sql_command.WriteString(fmt.Sprintf("%f", *(column_data.(*float64))))
				case "*time.Time":
					value := column_data.(*time.Time)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(FormatTime(*value)))
					sql_command.WriteString("'")
				case "time.Time":
					value := column_data.(time.Time)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(FormatTime(value)))
					sql_command.WriteString("'")
				case "*string":
					value := column_data.(*string)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(*value))
					sql_command.WriteString("'")
				case "string":
					value := column_data.(string)
					sql_command.WriteString("'")
					sql_command.WriteString(EscapeString(value))
					sql_command.WriteString("'")
				case "bool":
					if column_data.(bool) {
						sql_command.WriteString("1")
					} else {
						sql_command.WriteString("0")
					}
				case "*bool":
					if *(column_data.(*bool)) {
						sql_command.WriteString("1")
					} else {
						sql_command.WriteString("0")
					}
				default:
					errors = append(errors, fmt.Errorf("error: update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
				}
			}

			if index < (len(*identity_columns) - 1) {
				sql_command.WriteString(" AND ")
			}
		}
		sql_command.WriteString(" ;")

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command_string := sql_command.String()
		return &sql_command_string, options, nil
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}


	return &Record{
		Validate: func() []error {
			return validate()
		},
		Create: func() []error {
			sql, options, errors := getInsertSQL()
			if errors != nil {
				return errors
			}

			temp_table, temp_table_errors := getTable()
			if temp_table_errors != nil {
				return temp_table_errors
			}

			temp_database, temp_database_errors := temp_table.GetDatabase()
			if temp_database_errors != nil {
				return temp_database_errors
			}

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			json_array, errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql, options)

			if len(errors) > 0 {
				return errors
			}

			if options["get_last_insert_id"].(bool) && options["auto_increment_column_name"] != "" {
				if len(*json_array) != 1 {
					errors = append(errors, fmt.Errorf("error: get_last_insert_id not found "))
					return errors
				}

				record_from_db := (*json_array)[0].(Map)

				last_insert_id, last_insert_id_errors := record_from_db.GetString("LAST_INSERT_ID()")
				if last_insert_id_errors != nil {
					errors = append(errors, last_insert_id_errors...)
					return errors
				}

				last_insert_id_value, count_err := strconv.ParseUint(*last_insert_id, 10, 64)
				if count_err != nil {
					errors = append(errors, count_err)
					return errors
				}

				if !options.IsNil("auto_increment_column_name") && !options.IsEmptyString("auto_increment_column_name") {
					auto_increment_column_name, auto_increment_column_name_errors := options.GetString("auto_increment_column_name")
					if auto_increment_column_name_errors != nil {
						errors = append(errors, auto_increment_column_name_errors...)
					} else if IsNil(auto_increment_column_name) {
						errors = append(errors, fmt.Errorf("error: auto_increment_column_name is nil"))
					}

					set_auto_field_errors := SetField(struct_type, getData(), "[schema]", "[fields]", *auto_increment_column_name, &last_insert_id_value)
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

			temp_client, temp_client_errors := temp_database.GetClient()
			if temp_client_errors != nil {
				return temp_client_errors
			}

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(*temp_client, sql, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int64), nil
		},
		GetInt64Value: func(field string) (int64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "int64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int64), nil
		},
		GetInt32: func(field string) (*int32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int32), nil
		},
		GetInt32Value: func(field string) (int32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "int32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int32), nil
		},
		GetInt16: func(field string) (*int16, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int16), nil
		},
		GetInt16Value: func(field string) (int16, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "int16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int16), nil
		},
		GetInt8: func(field string) (*int8, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*int8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*int8), nil
		},
		GetInt8Value: func(field string) (int8, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "int8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(int8), nil
		},
		SetInt64: func(field string, value *int64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt64Value: func(field string, value int64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32: func(field string, value *int32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt32Value: func(field string, value int32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16: func(field string, value *int16) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt16Value: func(field string, value int16) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8: func(field string, value *int8) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetInt8Value: func(field string, value int8) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint64")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint64), nil
		},
		GetUInt64Value: func(field string) (uint64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint64")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint64), nil
		},
		GetUInt32: func(field string) (*uint32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint32")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint32), nil
		},
		GetUInt32Value: func(field string) (uint32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint32")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint32), nil
		},
		GetUInt16: func(field string) (*uint16, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint16")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint16), nil
		},
		GetUInt16Value: func(field string) (uint16, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint16")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint16), nil
		},
		GetUInt8: func(field string) (*uint8, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*uint8")
			if field_value_errors != nil {
				return nil, field_value_errors
			} else if field_value == nil {
				return nil, nil
			}
			return field_value.(*uint8), nil
		},
		GetUInt8Value: func(field string) (uint8, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "uint8")
			if field_value_errors != nil {
				return 0, field_value_errors
			}
			return field_value.(uint8), nil
		},
		SetUInt64: func(field string, value *uint64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt64Value: func(field string, value uint64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32: func(field string, value *uint32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt32Value: func(field string, value uint32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16: func(field string, value *uint16) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt16Value: func(field string, value uint16) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8: func(field string, value *uint8) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetUInt8Value: func(field string, value uint8) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetString: func(field string) (*string, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*string")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*string), nil
		},
		GetStringValue: func(field string) (string, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "string")
			if field_value_errors != nil {
				return "", field_value_errors
			}
			return field_value.(string), nil
		},
		SetString: func(field string, value *string) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetStringValue: func(field string, value string) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetBool: func(field string) (*bool, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*bool")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*bool), nil
		},
		GetBoolValue: func(field string) (bool, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "bool")
			if field_value_errors != nil {
				return false, field_value_errors
			}
			return field_value.(bool), nil
		},
		SetBool: func(field string, value *bool) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetBoolValue: func(field string, value bool) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat32: func(field string) (*float32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*float32")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float32), nil
		},
		GetFloat32Value: func(field string) (float32, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "float32")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float32), nil
		},
		SetFloat32: func(field string, value *float32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat32Value: func(field string, value float32) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		GetFloat64: func(field string) (*float64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "*float64")
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*float64), nil
		},
		GetFloat64Value: func(field string) (float64, []error) {
			field_value, field_value_errors := GetField(struct_type, getData(), "[schema]", "[fields]", field, "float64")
			if field_value_errors != nil {
				return 0.0, field_value_errors
			}
			return field_value.(float64), nil
		},
		SetFloat64: func(field string, value *float64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		SetFloat64Value: func(field string, value float64) []error {
			return SetField(struct_type, getData(), "[schema]", "[fields]", field, value)
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			fields_map, fields_map_errors := GetFields(struct_type, getData(), "[fields]")
			if fields_map_errors != nil {
				return fields_map_errors
			}
			return fields_map.ToJSONString(json)
		},
	}, nil
}
