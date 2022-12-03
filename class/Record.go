package class

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Validate  func() []error
	GetSQL    func(action string) (*string, []error)
	Create    func() []error
	Update    func() []error
	GetInt64  func(field string) (*int64, []error)
	SetInt64  func(field string, value *int64) []error
	GetUInt64 func(field string) (*uint64, []error)
	GetString func(field string) (*string, []error)
	SetString func(field string, value *string) []error 
	SetStringValue func(field string, value string) []error 
	ToJSONString  func() (*string, []error)
}

func newRecord(table *Table, record_data Map, database_reserved_words_obj *DatabaseReservedWords, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Record, []error) {
	SQLCommand := newSQLCommand()
	var errors []error
	struct_type := "*class.Record"

	if record_data == nil {
		errors = append(errors, fmt.Errorf("record_data is nil"))
	}

	if table == nil {
		errors = append(errors, fmt.Errorf("table is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	table_schema, table_schema_errors := table.GetSchema()
	if table_schema_errors != nil {
		errors = append(errors, table_schema_errors...)
	} else if IsNil(table_schema) {
		errors = append(errors, fmt.Errorf("table schema is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	database_reserved_words := database_reserved_words_obj.GetDatabaseReservedWords()
	column_name_whitelist_characters := column_name_whitelist_characters_obj.GetColumnNameCharacterWhitelist()
	

	data := Map{"[fields]": record_data}
	record_data.SetObject("[table]", table)
	data["[schema]"] = table_schema
	table_schema.SetMapValue("[table]", Map{"type":"*class.Table", "mandatory": true})


	getData := func() (*Map) {
		return &data
	}

	getRecordColumns := func() (*[]string, []error) {
		var columns []string
		column_name_whitelist_params := Map{"values": column_name_whitelist_characters, "value": nil, "label": "column_name_character", "data_type": "Column"}
		column_name_blacklist_params := Map{"values": database_reserved_words, "value": nil, "label": "column_name", "data_type": "Column"}

		fields_map, fields_map_errors := GetFields(getData())
		if fields_map_errors != nil {
			return nil, fields_map_errors
		}

		for _, column := range (*fields_map).Keys() {
			column_name_whitelist_params.SetString("value", &column)
			column_name_whiltelist_errors := WhitelistCharacters(column_name_whitelist_params)
			if column_name_whiltelist_errors != nil {
				continue
			}
			
			column_name_blacklist_params.SetString("value", &column)
			column_name_blacklist_errors := BlackListStringToUpper(column_name_blacklist_params)
			if column_name_blacklist_errors != nil {
				continue
			}	

			columns = append(columns, column)
		}
		return &columns, nil
	}

	getTable := func() (*Table, []error) {
		return GetTableField(getData(), "[table]")
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
			errors = append(errors, fmt.Errorf("table: %s can only have 1 auto_increment primary_key column, found: %s", table_name, auto_increment_columns))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command := fmt.Sprintf("INSERT INTO %s ", EscapeString(table_name))
		sql_command += "("
		for index, record_column := range *record_columns {
			sql_command += EscapeString(record_column)
			if index < (len(*record_columns) - 1) {
				sql_command += ", "
			}
		}

		sql_command += ") VALUES ("
		for index, record_column := range *record_columns {
			parameter, paramter_errors := GetField(getData(), record_column)
			if paramter_errors != nil {
				errors = append(errors, paramter_errors...)
				continue
			}

			rep := GetType(parameter)
			switch rep {
			case "string":
				sql_command += "\"" + EscapeString(parameter.(string)) + "\""
			case "*string":
				sql_command += "\"" + EscapeString(*(parameter.(*string))) + "\""
			default:
				errors = append(errors, fmt.Errorf("type: %s not supported for table please implement", rep))
			}

			if index < (len(*record_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ");"

		if len(errors) > 0 {
			return nil, nil, errors
		}

		return &sql_command, options, nil
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
				errors = append(errors, fmt.Errorf("record did not contain identify column: %s", identity_column))
			}
		}

		SetField(struct_type, getData(), "last_modified_date", GetTimeNow())

		record_non_identity_columns, record_non_identity_columns_errors := getNonIdentityColumnsUpdate()
		if record_non_identity_columns_errors != nil {
			return nil, nil, record_non_identity_columns_errors
		}

		if len(*record_non_identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("no non-identity columns detected in record to update"))
		}

		if len(*identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("table schema has no identity columns"))
		}

		if !Contains(*table_non_identity_columns, "last_modified_date") {
			errors = append(errors, fmt.Errorf("table schema does not have last_modified_date"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command := fmt.Sprintf("UPDATE %s \n", EscapeString(table_name))

		sql_command += "SET "

		for index, record_non_identity_column := range *record_non_identity_columns {
			sql_command += EscapeString(record_non_identity_column) + "="
			column_data, column_data_errors := GetField(getData(), record_non_identity_column)

			if column_data_errors != nil {
				errors = append(errors, column_data_errors...)
				continue
			}

			if IsNil(column_data) {
				sql_command += "NULL"
			} else {
				record_non_identity_column_type := GetType(column_data)
				switch record_non_identity_column_type {
				case "*uint64":
					value := column_data.(*uint64)
					sql_command += strconv.FormatUint(*value, 10)
				case "uint64":
					value := column_data.(uint64)
					sql_command += strconv.FormatUint(value, 10)
				case "*int64":
					value := column_data.(*int64)
					sql_command += strconv.FormatInt(*value, 10)
				case "int64":
					value := column_data.(int64)
					sql_command += strconv.FormatInt(value, 10)
				case "*int":
					value := column_data.(*int)
					sql_command += strconv.FormatInt(int64(*value), 10)
				case "int":
					value := column_data.(int)
					sql_command += strconv.FormatInt(int64(value), 10)
				case "*time.Time":
					value := column_data.(*time.Time)
					sql_command += "'" + EscapeString(FormatTime(*value)) + "'"
				case "time.Time":
					value := column_data.(time.Time)
					sql_command += "'" + EscapeString(FormatTime(value)) + "'"
				case "*string":
					value := column_data.(*string)
					sql_command += "'" + EscapeString(*value) + "'"
				case "string":
					value := column_data.(string)
					sql_command += "'" + EscapeString(value) + "'"
				default:
					errors = append(errors, fmt.Errorf("update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
				}
			}

			if index < len(*record_non_identity_columns)-1 {
				sql_command += ", \n"
			}
		}

		sql_command += " WHERE "
		for index, identity_column := range *identity_columns {
			sql_command += EscapeString(identity_column) + " = "
			column_data, column_data_errors := GetField(getData(), identity_column)

			if column_data_errors != nil {
				errors = append(errors, column_data_errors...)
				continue
			}

			if IsNil(column_data) {
				sql_command += "NULL"
			} else {
				record_non_identity_column_type := GetType(column_data)
				switch record_non_identity_column_type {
				case "*uint64":
					value := column_data.(*uint64)
					sql_command += strconv.FormatUint(*value, 10)
				case "uint64":
					value := column_data.(uint64)
					sql_command += strconv.FormatUint(value, 10)
				case "*int64":
					value := column_data.(*int64)
					sql_command += strconv.FormatInt(*value, 10)
				case "int64":
					value := column_data.(int64)
					sql_command += strconv.FormatInt(value, 10)
				case "*int":
					value := column_data.(*int)
					sql_command += strconv.FormatInt(int64(*value), 10)
				case "int":
					value := column_data.(int)
					sql_command += strconv.FormatInt(int64(value), 10)
				case "*time.Time":
					value := column_data.(*time.Time)
					sql_command += "'" + EscapeString(FormatTime(*value)) + "'"
				case "time.Time":
					value := column_data.(time.Time)
					sql_command += "'" + EscapeString(FormatTime(value)) + "'"
				case "*string":
					value := column_data.(*string)
					sql_command += "'" + EscapeString(*value) + "'"
				case "string":
					value := column_data.(string)
					sql_command += "'" + EscapeString(value) + "'"
				default:
					errors = append(errors, fmt.Errorf("update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
				}
			}

			if index < (len(*identity_columns) - 1) {
				sql_command += " AND "
			}
		}
		sql_command += " ;"

		if len(errors) > 0 {
			return nil, nil, errors
		}

		return &sql_command, options, nil
	}

	x := Record{
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

			json_array, errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql, options)

			if len(errors) > 0 {
				return errors
			}

			if options["get_last_insert_id"].(bool) && options["auto_increment_column_name"] != "" {
				if len(*json_array) != 1 {
					errors = append(errors, fmt.Errorf("get_last_insert_id not found "))
					return errors
				}

				last_insert_id, last_insert_id_errors := (*json_array)[0].(Map).GetString("LAST_INSERT_ID()")
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
						errors = append(errors, fmt.Errorf("auto_increment_column_name is nil"))
					}

					fields_map, fields_map_errors := GetFields(getData())
					if fields_map_errors != nil { 
						errors = append(errors, fields_map_errors...)
					}

					if len(errors) > 0 {
						return errors
					}

					fields_map.SetUInt64(*auto_increment_column_name, &last_insert_id_value)
				}
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

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(temp_client, sql, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			field_value, field_value_errors := GetField(getData(), field)
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*int64), nil
		},
		SetInt64: func(field string, value *int64) []error {
			return SetField(struct_type, getData(), field, value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			field_value, field_value_errors := GetField(getData(), field)
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*uint64), nil
		},
		GetString: func(field string) (*string, []error) {
			field_value, field_value_errors := GetField(getData(), field)
			if field_value_errors != nil {
				return nil, field_value_errors
			}
			return field_value.(*string), nil
		},
		SetString: func(field string, value *string) []error {
			return SetField(struct_type, getData(), field, value)
		},
		SetStringValue: func(field string, value string) []error {
			return SetField(struct_type, getData(), field, value)
		},
		ToJSONString: func() (*string, []error) {
			return getData().ToJSONString()
		},
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
