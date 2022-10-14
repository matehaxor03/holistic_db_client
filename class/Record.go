package class

import (
	"fmt"
	"strings"
	"strconv"
)

func CloneRecord(record *Record) (*Record) {
	if record == nil {
		return nil
	}

	return record.Clone()
}

type Record struct {
	Validate func() ([]error)
	Clone func() (*Record)
	GetSQL func(action string) (*string, []error)
	Create func() ([]error)
	Update func() ([]error)
	GetInt64 func(field string) (*int64, []error)
	SetInt64 func(field string, value *int64)
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

	table_data := table.GetData()
	table_columns := (*table).GetTableColumns()
	expanded_record := Map{}
	for _, column := range record_data.Keys() {
		for _, schema_column := range table_columns {
			if column == schema_column {
				column_data := Map{"value": record_data[column]}

				if table_data.GetType(column) != "class.Map" {
					errors = append(errors, fmt.Errorf("Record.newRecord table schema column: %s is not a map", column))
					continue
				}

				for _, schema_column_data := range table_data.M(column).Keys() {
					switch schema_column_data {
					case "type", "default", "filters", "mandatory", "primary_key", "auto_increment", "unsigned": 
						column_data[schema_column_data] = table_data.M(column)[schema_column_data]
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

	data := expanded_record.Clone()
	data["[table]"] = Map{"value":CloneTable(table),"mandatory":true}

	getData := func() (Map) {
		return data.Clone()
	}

	getTableColumns := func() ([]string) {
		var columns []string
		for _, column := range getData().Keys() {
			if strings.HasPrefix(column, "[") {
				continue
			}

			if strings.HasSuffix(column, "]") {
				continue
			}

			if data.GetType(column) != "class.Map" {
				continue
			}

			columns = append(columns, column)
		}
		return columns
	}

	getTable := func() (*Table) {
		return CloneTable(data.M("[table]").GetObject("value").(*Table))
	}

	/*
	getNonIdentityColumns := func() ([]string) {
		record_columns := getTableColumns()
		non_identity_columns := getTable().GetNonIdentityColumns()
		var record_non_identity_columns []string
		for _, record_column := range record_columns {
			for _, non_identity_column := range non_identity_columns {
				if non_identity_column == record_column {
					record_non_identity_columns = append(record_non_identity_columns, non_identity_column)
					break
				}
			}
		}
		return record_non_identity_columns
	}*/

	getNonIdentityColumnsUpdate := func() ([]string) {
		record_columns := getTableColumns()
		non_identity_columns := getTable().GetNonIdentityColumns()
		var record_non_identity_columns []string
		for _, record_column := range record_columns {
			if record_column == "created_date" || 
			   record_column == "archieved_date" ||
			   record_column == "active" {
				continue
			}

			for _, non_identity_column := range non_identity_columns {
				if non_identity_column == record_column {
					record_non_identity_columns = append(record_non_identity_columns, non_identity_column)
					break
				}
			}
		}
		return record_non_identity_columns
	}

	getIdentityColumns := func() ([]string) {
		record_columns := getTableColumns()
		identity_columns := getTable().GetIdentityColumns()
		var record_identity_columns []string
		for _, record_column := range record_columns {
			for _, identity_column := range identity_columns {
				if identity_column == record_column {
					record_identity_columns = append(record_identity_columns, identity_column)
					break
				}
			}
		}
		return record_identity_columns
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Record")
	}

	getInsertSQL := func() (*string, Map, []error) {
		options := Map{"use_file": false, "no_column_headers": true, "get_last_insert_id": false}
		errors := validate()
		
		if len(errors) > 0 {
			return nil, nil, errors
		}
	
		table := getTable()
		table_schema := table.GetData()
		record := getData()
		valid_columns := table.GetTableColumns()
		record_columns := getTableColumns()
		for _, record_column := range record_columns {
			if !Contains(valid_columns, record_column) {
				errors = append(errors, fmt.Errorf("column: %s does not exist for table: %s valid column names are: %s", record_column, *(table.GetTableName()), valid_columns))
			} else {
				if strings.HasPrefix(record_column, "credential") {
					options["use_file"] = true
				}
			}

			type_of_schema_column := *((table_schema.M(record_column)).S("type"))
			type_of_record_column := record.M(record_column).GetType("value")
			if type_of_record_column != type_of_schema_column {
				errors = append(errors, fmt.Errorf("table schema for column: %s has type: %s however record has type: %s", record_column, type_of_schema_column, type_of_record_column))
			}
		}

		auto_increment_columns := 0
		for _, valid_column := range valid_columns {
			column_definition := table_schema.M(valid_column)
			
			if column_definition.HasKey("primary_key") &&
				column_definition.GetType("primary_key") == "bool" &&
				*(column_definition.B("primary_key")) &&
				column_definition.HasKey("auto_increment") && 
				column_definition.GetType("auto_increment") == "bool" &&
				*(column_definition.B("auto_increment")) {
				options["get_last_insert_id"] = true
				options["auto_increment_column_name"] = valid_column
				auto_increment_columns += 1
			}
		}

		if auto_increment_columns > 1 {
			errors = append(errors, fmt.Errorf("table: %s can only have 1 auto_increment primary_key column, found: %s", *(table.GetTableName()), auto_increment_columns))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		sql_command := fmt.Sprintf("INSERT INTO %s ", EscapeString(*getTable().GetTableName()))
		sql_command += "("
		for index, record_column := range record_columns {
			sql_command += EscapeString(record_column)
			if index < (len(record_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ") VALUES ("
		for index, record_column := range record_columns {
			rep := record.M(record_column).GetType("value")
			switch rep {
			default:
				//EscapeString
				errors = append(errors, fmt.Errorf("type: %s not supported for table please implement", rep))
			}
			
			if index < (len(record_columns) - 1) {
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
		table_schema := table.GetData()
		record := getData()
		valid_columns := table.GetTableColumns()
		record_columns := getTableColumns()
		for _, record_column := range record_columns {
			if !Contains(valid_columns, record_column) {
				errors = append(errors, fmt.Errorf("column: %s does not exist for table: %s valid column names are: %s", record_column, *(table.GetTableName()), valid_columns))
			} else {
				if strings.HasPrefix(record_column, "credential") {
					options["use_file"] = true
				}
			}

			type_of_schema_column := *((table_schema.M(record_column)).S("type"))
			type_of_record_column := record.M(record_column).GetType("value")
			if type_of_record_column != type_of_schema_column {
				errors = append(errors, fmt.Errorf("table schema for column: %s has type: %s however record has type: %s", record_column, type_of_schema_column, type_of_record_column))
			}
		}

		identity_columns := table.GetIdentityColumns()
		record_identity_columns := getIdentityColumns()
		for _, identity_column := range identity_columns {
			found_identity_column := false
			for _, record_identity_column := range record_identity_columns {
				if identity_column == record_identity_column {
					found_identity_column = true
				}
			}

			if !found_identity_column {
				errors = append(errors, fmt.Errorf("record did not contain identify column: %s", identity_column))
			}
		}

		record_non_identity_columns := getNonIdentityColumnsUpdate()

		if len(record_non_identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("no non-identity columns detected in record to update"))
		}

		if len(identity_columns) == 0 {
			errors = append(errors, fmt.Errorf("table schema has no identity columns"))
		}

		if !Contains(record_non_identity_columns, "last_modified_date") {
			errors = append(errors, fmt.Errorf("table record does not have last_modified_date"))
		}

		if len(errors) > 0 {
			return nil, nil, errors
		}

		record.M("last_modified_date")["value"] = GetTimeNow()

		sql_command := fmt.Sprintf("UPDATE %s \n", EscapeString(*getTable().GetTableName()))

		sql_command += "SET "

		for index, record_non_identity_column := range record_non_identity_columns {
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
				default:
					errors = append(errors, fmt.Errorf("update record type is not supported please implement for set clause: %s", record_non_identity_column_type))
				}
			}

			if index < len(record_non_identity_columns) - 1 {
				sql_command += ", \n"
			}
		}
		
		sql_command += " WHERE "
		for index, identity_column := range identity_columns {
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
			
			if index < (len(identity_columns) - 1) {
				sql_command += " AND "
			}
		}
		sql_command += " ;"

		if len(errors) > 0 {
			return nil, nil, errors
		}

		return &sql_command, options, nil
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	x := Record{
		Validate: func() ([]error) {
			return validate()
		},
		Clone: func() (*Record) {
			clone_value, _ := NewRecord(getTable(), getData())
			return clone_value
		},
		Create: func() ([]error) {
			sql, options, errors := getInsertSQL()
			if errors != nil {
				return errors
			}

			json_array, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getTable().GetDatabase().GetClient(), sql, options)
						
			if stderr != nil && *stderr != "" {
				if strings.Contains(*stderr, " some error") {
					errors = append(errors, fmt.Errorf("insert record failed"))
				} else {
					errors = append(errors, fmt.Errorf(*stderr))
				}
			}
		
			if len(errors) > 0 {
				return errors
			}

			if options["get_last_insert_id"].(bool) && options["auto_increment_column_name"] != "" {
				if len(*json_array) != 1 {
					errors = append(errors, fmt.Errorf("get_last_insert_id not found "))
					return errors
				}
				
				count, count_err := strconv.ParseUint(*((*json_array)[0].(Map).S("LAST_INSERT_ID()")), 10, 64)
				if count_err != nil {
					errors = append(errors, count_err)
					return errors
				}

				if options.S("auto_increment_column_name") != nil && *(options.S("auto_increment_column_name")) != "" {
					auto_increment_column_name := options.S("auto_increment_column_name")
					data.SetMap(*auto_increment_column_name, Map{"type":"uint64", "value":count, "auto_increment":true, "primary_key":true})
				}
			}
			return nil
		},
		Update: func() ([]error) {
			sql, options, errors := getUpdateSQL()
			if errors != nil {
				return errors
			}

			_, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getTable().GetDatabase().GetClient(), sql, options)
						
			if stderr != nil && *stderr != "" {
				if strings.Contains(*stderr, " some error") {
					errors = append(errors, fmt.Errorf("insert record failed"))
				} else {
					errors = append(errors, fmt.Errorf(*stderr))
				}
			}
		
			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			return getData().M(field).GetInt64("value")
		},
		SetInt64: func(field string, value *int64) {
			data.M(field).SetInt64("value", value)
		},
		GetUInt64: func(field string) (*uint64, []error) {
			return getData().M(field).GetUInt64("value")
		},
    }

	return &x, nil
}
