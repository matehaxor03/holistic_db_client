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
	GetInt64 func(field string) (*int64, []error)
	GetUInt64 func(field string) (*uint64, []error)
}

func NewRecord(table *Table, record_data Map) (*Record, []error) {
	SQLCommand := newSQLCommand()
	var errors []error

	if record_data == nil {
		errors = append(errors, fmt.Errorf("record_data is nil"))
		return nil, errors
	}

	data := record_data.Clone()
	data["table"] = Map{"value":CloneTable(table),"mandatory":true}

	getData := func() (Map) {
		return data.Clone()
	}

	getTableColumns := func() ([]string) {
		var columns []string
		for _, column := range getData().Keys() {
			if column == "table_name" {
				continue
			}

			if data.GetType(column) != "class.Map" {
				continue
			}

			columnSchema := data[column].(Map)

			rep := fmt.Sprintf("%T", columnSchema["value"])
			switch rep {
				case "*uint64", "*int64", "*int", "uint64", "uint", "int64", "int", "*string", "string", "*time.Time", "time.Time", "*bool", "bool", "<nil>":
				default:
				continue
			}
			columns = append(columns, column)
		}
		return columns
	}

	getTable := func() (*Table) {
		return CloneTable(data.M("table").GetObject("value").(*Table))
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Table")
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
				if strings.HasPrefix(record_column, "credential_") {
					options["use_file"] = true
				}
			}

			type_of_schema_column := *((table_schema.M(record_column)).S("type"))
			type_of_record_column := record.GetType(record_column)
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
			rep := record.GetType(record_column)
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
					data[*auto_increment_column_name] = count
				}
			}

			return nil
		},
		GetInt64: func(field string) (*int64, []error) {
			return getData().M(field).GetInt64("value")
		},
		GetUInt64: func(field string) (*uint64, []error) {
			return getData().M(field).GetUInt64("value")
		},
    }

	return &x, nil
}
