package class

import (
	"fmt"
	"strconv"
	"strings"
)

func CloneRecord(record *Record) (*Record, []error) {
	if record == nil {
		return nil, nil
	}

	return record.Clone()
}

type Record struct {
	Validate  func() []error
	Clone     func() (*Record, []error)
	GetSQL    func(action string) (*string, []error)
	Create    func() []error
	Update    func() []error
	GetInt64  func(field string) (*int64, []error)
	SetInt64  func(field string, value *int64)
	GetUInt64 func(field string) (*uint64, []error)
}

func NewRecord(table *Table, record_data Map) (*Record, []error) {
	SQLCommand := NewSQLCommand()
	var errors []error

	if record_data == nil {
		errors = append(errors, fmt.Errorf("record_data is nil"))
	}

	if table == nil {
		errors = append(errors, fmt.Errorf("table is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	table_data, table_data_errors := table.GetData()
	if table_data_errors != nil {
		return nil, table_data_errors
	}

	table_columns, table_columns_errors := (*table).GetTableColumns()
	if table_columns_errors != nil {
		return nil, table_columns_errors
	}
	expanded_record := Map{}
	for _, column := range record_data.Keys() {
		for _, schema_column := range *table_columns {
			if column == schema_column {
				column_data := Map{"value": record_data[column]}

				if table_data.GetType(column) != "class.Map" {
					errors = append(errors, fmt.Errorf("Record.newRecord table schema column: %s is not a map", column))
					continue
				}

				for _, schema_column_data := range table_data.M(column).Keys() {
					switch schema_column_data {
					case "type", "default", "filters", "mandatory", "primary_key", "auto_increment", "unsigned":
						column_data[schema_column_data] = (*(table_data).M(column))[schema_column_data]
					case "value":
					default:
						errors = append(errors, fmt.Errorf("Record.newRecord table schema column: attribute not supported please implement: %s", schema_column_data))
					}
				}
				expanded_record[column] = column_data
				break
			}
		}
	}

	data, data_errors := expanded_record.Clone()
	if data_errors != nil {
		return nil, data_errors
	}
	(*data)["[table]"] = Map{"value": CloneTable(table), "mandatory": true}

	getData := func() (*Map, []error) {
		return data.Clone()
	}

	getTableColumns := func() (*[]string, []error) {
		var columns []string
		column_name_whitelist_params := Map{"values": GetColumnNameValidCharacters(), "value": nil, "label": "column_name_character", "data_type": "Column"}
		column_name_blacklist_params := Map{"values": GetMySQLKeywordsAndReservedWordsInvalidWords(), "value": nil, "label": "column_name", "data_type": "Column"}

		data_clone, data_clone_errors := getData()
		if data_clone_errors != nil {
			return nil, data_clone_errors
		}

		for _, column := range (*data_clone).Keys() {
			if data.GetType(column) != "class.Map" {
				continue
			}

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

	getTable := func() *Table {
		return CloneTable(data.M("[table]").GetObject("value").(*Table))
	}

	getNonIdentityColumnsUpdate := func() (*[]string, []error) {
		record_columns, record_columns_errors := getTableColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		non_identity_columns, non_identity_columns_errors := getTable().GetNonIdentityColumns()
		if non_identity_columns_errors != nil {
			return nil, non_identity_columns_errors
		}

		var record_non_identity_columns []string
		for _, record_column := range *record_columns {
			if record_column == "created_date" ||
				record_column == "archieved_date" ||
				record_column == "active" {
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
		record_columns, record_columns_errors := getTableColumns()
		if record_columns_errors != nil {
			return nil, record_columns_errors
		}

		identity_columns, identity_columns_errors := getTable().GetIdentityColumns()
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
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Record")
	}

	getInsertSQL := func() (*string, Map, []error) {
		options := Map{"use_file": false, "no_column_headers": true, "get_last_insert_id": false}
		errors := validate()

		if len(errors) > 0 {
			return nil, nil, errors
		}

		table := getTable()
		table_schema, table_schema_errors := table.GetData()
		
		if table_schema_errors != nil {
			return nil, nil, table_schema_errors
		}

		record, record_errors := getData()
		if record_errors != nil {
			return nil, nil, record_errors
		}

		valid_columns, valid_columns_errors := table.GetTableColumns()
		if valid_columns_errors != nil {
			return nil, nil, valid_columns_errors
		}
		record_columns, record_columns_errors := getTableColumns()
		if record_columns_errors != nil {
			return nil, nil, record_columns_errors
		}

		for _, record_column := range *record_columns {
			if !Contains(*valid_columns, record_column) {
				errors = append(errors, fmt.Errorf("column: %s does not exist for table: %s valid column names are: %s", record_column, table.GetTableName(), valid_columns))
			} else {
				if strings.HasPrefix(record_column, "credential") {
					options["use_file"] = true
				}
			}

			type_of_schema_column, _ := (table_schema.M(record_column)).GetString("type")
			type_of_record_column := record.M(record_column).GetType("value")
			if strings.Replace(type_of_record_column, "*", "", -1) != strings.Replace(*type_of_schema_column, "*", "", -1) {
				errors = append(errors, fmt.Errorf("table schema for column: %s has type: %s however record has type: %s", record_column, type_of_schema_column, type_of_record_column))
			}
		}

		auto_increment_columns := 0
		for _, valid_column := range *valid_columns {
			column_definition := table_schema.M(valid_column)

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
			errors = append(errors, fmt.Errorf("table: %s can only have 1 auto_increment primary_key column, found: %s", table.GetTableName(), auto_increment_columns))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command := fmt.Sprintf("INSERT INTO %s ", EscapeString(table.GetTableName()))
		sql_command += "("
		for index, record_column := range *record_columns {
			sql_command += EscapeString(record_column)
			if index < (len(*record_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ") VALUES ("
		for index, record_column := range *record_columns {
			rep := record.M(record_column).GetType("value")
			switch rep {
			default:
				//EscapeString
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

		table := getTable()
		table_schema, table_schema_errors := table.GetData()

		if table_schema_errors != nil {
			return nil, nil, table_schema_errors
		}

		record, record_errors := getData()
		if record_errors != nil {
			return nil, nil, record_errors
		}

		valid_columns, valid_columns_errors := table.GetTableColumns()
		if valid_columns_errors != nil {
			return nil, nil, valid_columns_errors
		}

		record_columns, record_columns_errors := getTableColumns()
		if record_columns_errors != nil {
			return nil, nil, record_columns_errors
		}

		for _, record_column := range *record_columns {
			if !Contains(*valid_columns, record_column) {
				errors = append(errors, fmt.Errorf("column: %s does not exist for table: %s valid column names are: %s", record_column, table.GetTableName(), valid_columns))
			} else {
				if strings.HasPrefix(record_column, "credential") {
					options["use_file"] = true
				}
			}

			type_of_schema_column, _ := ((*table_schema).M(record_column)).GetString("type")
			type_of_record_column := record.M(record_column).GetType("value")
			if strings.Replace(type_of_record_column, "*", "", -1) != strings.Replace(*type_of_schema_column, "*", "", -1) {
				errors = append(errors, fmt.Errorf("table schema for column: %s has type: %s however record has type: %s", record_column, type_of_schema_column, type_of_record_column))
			}
		}

		identity_columns, identity_columns_errors := table.GetIdentityColumns()
		if identity_columns_errors != nil {
			return nil, nil, identity_columns_errors
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

		if !Contains(*record_non_identity_columns, "last_modified_date") {
			errors = append(errors, fmt.Errorf("table record does not have last_modified_date"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		(*(record.M("last_modified_date")))["value"] = GetTimeNow()

		sql_command := fmt.Sprintf("UPDATE %s \n", EscapeString(table.GetTableName()))

		sql_command += "SET "

		for index, record_non_identity_column := range *record_non_identity_columns {
			sql_command += EscapeString(record_non_identity_column) + "="
			column_data := record.M(record_non_identity_column)
			record_non_identity_column_type := column_data.GetType("value")

			if column_data.IsNil("value") {
				sql_command += "NULL"
			} else {
				switch record_non_identity_column_type {
				case "uint64", "*uint64":
					value, value_errors := column_data.GetUInt64("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatUint(*value, 10)
					}
				case "*int64", "int64":
					value, value_errors := column_data.GetInt64("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatInt(*value, 10)
					}
				case "*int", "int":
					value, value_errors := column_data.GetInt("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				case "*time.Time", "time.Time":
					value, value_errors := column_data.GetTime("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += "'" + FormatTime(*value) + "'"
					}
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

			column_data := record.M(identity_column)
			record_identity_column_type := column_data.GetType("value")

			if column_data.IsNil("value") {
				errors = append(errors, fmt.Errorf("identity column is nil %s", identity_column))
			} else {
				switch record_identity_column_type {
				case "uint64", "*uint64":
					value, value_errors := column_data.GetUInt64("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatUint(*value, 10)
					}
				case "*int64", "int64":
					value, value_errors := column_data.GetInt64("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatInt(*value, 10)
					}
				case "*int", "int":
					value, value_errors := column_data.GetInt("value")
					if value_errors != nil {
						errors = append(errors, value_errors...)
					} else {
						sql_command += strconv.FormatInt(int64(*value), 10)
					}
				default:
					errors = append(errors, fmt.Errorf("update record type is not supported please implement for where clause: %s", record_identity_column_type))
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
		Clone: func() (*Record, []error) {
			cloned_data, cloned_data_errors := getData()
			if cloned_data_errors != nil {
				return nil, cloned_data_errors
			}
			return NewRecord(getTable(), *cloned_data)
		},
		Create: func() []error {
			sql, options, errors := getInsertSQL()
			if errors != nil {
				return errors
			}

			json_array, errors := SQLCommand.ExecuteUnsafeCommand(getTable().GetDatabase().GetClient(), sql, options)

			if len(errors) > 0 {
				return errors
			}

			if options["get_last_insert_id"].(bool) && options["auto_increment_column_name"] != "" {
				if len(*json_array) != 1 {
					errors = append(errors, fmt.Errorf("get_last_insert_id not found "))
					return errors
				}

				last_insert_id, _ := (*json_array)[0].(*Map).GetString("LAST_INSERT_ID()")
				count, count_err := strconv.ParseUint(*last_insert_id, 10, 64)
				if count_err != nil {
					errors = append(errors, count_err)
					return errors
				}


				if !options.IsNil("auto_increment_column_name") && !options.IsEmptyString("auto_increment_column_name") {
					auto_increment_column_name, _ := options.GetString("auto_increment_column_name")
					auto_increment_column_schema := Map{"type": "uint64", "value": count, "auto_increment": true, "primary_key": true}
					data.SetMap(*auto_increment_column_name, &auto_increment_column_schema)
				}
			}
			return nil
		},
		Update: func() []error {
			sql, options, generate_sql_errors := getUpdateSQL()
			if generate_sql_errors != nil {
				return generate_sql_errors
			}

			_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getTable().GetDatabase().GetClient(), sql, options)

			if execute_errors != nil {
				return execute_errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			cloned_data, cloned_data_errors := getData()
			if cloned_data_errors != nil {
				return nil, cloned_data_errors
			}
			return cloned_data.M(field).GetInt64("value")
		},
		SetInt64: func(field string, value *int64) {
			data.M(field).SetInt64("value", value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			cloned_data, cloned_data_errors := getData()
			if cloned_data_errors != nil {
				return nil, cloned_data_errors
			}
			return cloned_data.M(field).GetUInt64("value")
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
