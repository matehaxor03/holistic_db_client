package class

import (
	"fmt"
	"strings"
	"strconv"
)

func GET_TABLE_DATA_DEFINITION_STATEMENTS() Array {
	return Array {GET_DATA_DEFINTION_STATEMENT_CREATE()}
}

func GET_TABLE_LOGIC_OPTIONS_CREATE() ([][]string){
	return [][]string{GET_LOGIC_STATEMENT_IF_NOT_EXISTS()}
}

func GET_TABLE_OPTIONS() (map[string]map[string][][]string) {
	var root = make(map[string]map[string][][]string)
	
	var logic_options = make(map[string][][]string)
	logic_options[GET_DATA_DEFINTION_STATEMENT_CREATE()] = GET_TABLE_LOGIC_OPTIONS_CREATE()

	root[GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options

	return root
}

func GetTableNameValidCharacters() *string {
	values :=  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._"
	return &values
}

func GetColumnNameValidCharacters() *string {
	values :=  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	return &values
}

func CloneTable(table *Table) (*Table) {
	if table == nil {
		return nil
	}

	return table.Clone()
}

type Table struct {
	Validate func() ([]error)
	Clone func() (*Table)
	GetSQL func(action string) (*string, []error)
	Create func() ([]error)
	GetTableName func() (*string)
	GetTableColumns func() ([]string)
	GetIdentityColumns func() ([]string)
	GetNonIdentityColumns func() ([]string)
	Count func() (*uint64, []error)
	GetData func() (Map)
	CreateRecord func(record Map) (*Record, []error)
	Select func(filter Map, limit *uint64, offset *uint64) (*[]Record, []error)
	GetDatabase func() (*Database)
	ToJSONString func() string
}

func NewTable(database *Database, schema Map, options map[string]map[string][][]string) (*Table, []error) {
	var this_table *Table
	SQLCommand := NewSQLCommand()
	var errors []error

	if schema == nil {
		errors = append(errors, fmt.Errorf("schema is nil"))
		return nil, errors
	}

	if !schema.HasKey("[table_name]") {
		errors = append(errors, fmt.Errorf("table_name field is nil"))
	} else if schema.GetType("[table_name]") != "class.Map" {
		errors = append(errors, fmt.Errorf("table_name field is not a map"))
	} else {
		boolean_value := true
		schema.M("[table_name]").SetBool("mandatory", &boolean_value)
		schema.M("[table_name]")[FILTERS()] = Array{ Map {"values":GetTableNameValidCharacters(),"function":getWhitelistCharactersFunc()}}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	data := schema.Clone()
	data["[database]"] = Map{"value":CloneDatabase(database),"mandatory":true}
	data["[options]"] = Map{"value":options,"mandatory":false}
	data["active"] = Map{"type":"*bool", "mandatory":true, "default":true}
	data["created_date"] = Map{"type":"*time.Time", "mandatory":true, "default":"now"}
	data["last_modified_date"] = Map{"type":"*time.Time", "mandatory":true, "default":"now"}
	data["archieved_date"] = Map{"type":"*time.Time", "mandatory":true, "default":"now"}

	getData := func() (Map) {
		return data.Clone()
	}

	getTableName := func() (*string) {
		return CloneString(data.M("[table_name]").S("value"))
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

	getIdentityColumns := func() ([]string) {
		var columns []string
		for _, column := range getTableColumns() {
			columnSchema := data[column].(Map)

			if columnSchema.IsBoolFalse("primary_key") {
				continue
			}
			
			columns = append(columns, column)
		}
		return columns
	}

	getNonIdentityColumns := func() ([]string) {
		var columns []string
		for _, column := range getTableColumns() {
			columnSchema := data[column].(Map)

			if columnSchema.IsBoolTrue("primary_key") {
				continue
			}
			
			columns = append(columns, column)
		}
		return columns
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(getData(), "Table")
	}

	getDatabase := func() (*Database) {
		return CloneDatabase(data.M("[database]").GetObject("value").(*Database))
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("[options]").GetObject("value").(map[string]map[string][][]string)
	}

	setTable := func(table *Table) {
		this_table = table
	}

	getTable := func() *Table{
		return this_table
	}

	getSQL := func(command string) (*string, []error) {
		errors := validate()

		m := Map{}
		m.SetArray("values", GET_TABLE_DATA_DEFINITION_STATEMENTS())
		m.SetString("value", &command)
		commandTemp := "command"
		m.SetString("label", &commandTemp)
		someValue :=  "Table"
		m.SetString("data_type", &someValue)

		command_errs := WhiteListString(m)


		if command_errs != nil {
			errors = append(errors, command_errs...)	
		}

		logic_option, logic_options_errs := GetLogicCommand(command, GET_LOGIC_STATEMENT_FIELD_NAME(), GET_TABLE_OPTIONS(), options, "Table")
		if logic_options_errs != nil {
			errors = append(errors, logic_options_errs...)	
		}
		
		if len(errors) > 0 {
			return nil, errors
		}

		sql_command := fmt.Sprintf("%s TABLE ", command)
		
		if *logic_option != "" {
			sql_command += fmt.Sprintf("%s ", *logic_option)
		}
		
		sql_command += fmt.Sprintf("%s ", EscapeString(*getTableName()))

		valid_columns := getTableColumns()
		primary_key_count := 0

		sql_command += "("
		for index, column := range valid_columns {
			columnSchema := data[column].(Map)

			if !columnSchema.HasKey("type") {
				errors = append(errors, fmt.Errorf("column: %s does not have type attribute", column))
				continue
			}

			typeOf := columnSchema.S("type")
			switch *typeOf {
			case "*uint64", "*int64", "*int", "uint64", "uint", "int64", "int":
				sql_command += EscapeString(column) + " BIGINT"
				

				if *typeOf == "*uint64" || 
				   *typeOf == "*uint" ||
				   *typeOf == "uint64" ||
				   *typeOf == "uint"  {
					sql_command += " UNSIGNED"
				}

				if columnSchema.HasKey("auto_increment") {
					if columnSchema.IsBool("auto_increment") && !columnSchema.IsNil("auto_increment") {
						if columnSchema.IsBoolTrue("auto_increment") {
							sql_command += " AUTO_INCREMENT"
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.GetType("auto_increment")))
					}
				}

				if columnSchema.HasKey("primary_key") {
					if columnSchema.IsBool("primary_key") && !columnSchema.IsNil("primary_key") {
						if columnSchema.IsBoolTrue("primary_key") {
							sql_command += " PRIMARY KEY NOT NULL"
							primary_key_count += 1
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
					}		
				} else if !strings.HasPrefix(*typeOf, "*") {
					sql_command += " NOT NULL"
				}

				if columnSchema.HasKey("default") && columnSchema.GetType("default") == "int" {
					default_value, default_value_errors := columnSchema.GetInt64("default")
					if default_value_errors != nil {
						errors = append(errors, default_value_errors...)
					} else {
						sql_command += " DEFAULT " + strconv.FormatInt(*default_value, 10)
					}
				}
			case "*time.Time":
				sql_command += EscapeString(column) + " TIMESTAMP(6)"
				if columnSchema.HasKey("default") {
					if columnSchema.S("default") == nil {
						errors = append(errors, fmt.Errorf("column: %s had nil default value", column))
						continue
					} else if *(columnSchema.S("default")) != "now" {
						errors = append(errors, fmt.Errorf("column: %s had default value it did not understand", column))
						continue
					}
					
					sql_command += " DEFAULT CURRENT_TIMESTAMP(6)"
				} 
				case "*bool", "bool": 
				sql_command += EscapeString(column) + " BOOLEAN"
				if columnSchema.HasKey("default") {
					if columnSchema.IsNil("default") {
						errors = append(errors, fmt.Errorf("column: %s had nil default value", column))
					} else if !columnSchema.IsBool("default") {
						errors = append(errors, fmt.Errorf("column: %s had non-boolean default value", column))
					} else if columnSchema.IsBoolTrue("default") {
						sql_command += " DEFAULT 1"
					} else if columnSchema.IsBoolFalse("default") {
						sql_command += " DEFAULT 0"
					} else {
						errors = append(errors, fmt.Errorf("column: %s had unknown error for boolean default value", column))
					}
				} 
			default:
				errors = append(errors, fmt.Errorf("Table.getSQL type: %s is not supported please implement for column %s", *typeOf, column))
			}

			if index < (len(valid_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ")"

		if primary_key_count == 0 {
			errors = append(errors, fmt.Errorf("table: %s must have at least 1 primary key", EscapeString(*getTableName())))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command += ";"

		return &sql_command, nil
	}

	createTable := func() ([]error) {
		sql_command, sql_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if sql_command_errors != nil {
			return sql_command_errors
		}
	
		_, execute_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), sql_command, Map{"use_file": false})
	
		if execute_errors != nil {
			return execute_errors
		}
	
		return nil
	}

	validate_errors := validate()

	if validate_errors != nil {
		errors = append(errors, validate_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	x := Table{
		Validate: func() ([]error) {
			return validate()
		},
		GetSQL: func(action string) (*string, []error) {
			return getSQL(action)
		},
		GetDatabase: func() (*Database) {
			return getDatabase()
		},
		Clone: func() (*Table) {
			clone_value, _ := NewTable(getDatabase(), getData(), getOptions())
			return clone_value
		},
		GetTableColumns: func() ([]string) {
			return getTableColumns()
		},
		GetIdentityColumns: func() ([]string) {
			return getIdentityColumns()
		},
		GetNonIdentityColumns: func() ([]string) {
			return getNonIdentityColumns()
		},
		Create: func() ([]error) {
			errors := createTable()
			if errors != nil {
				return errors
			}

			return nil
		},
		Count: func() (*uint64, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			sql :=  fmt.Sprintf("SELECT COUNT(*) FROM %s;", EscapeString((*getTableName())))
			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false})

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}
		
			if len(errors) > 0 {
				return nil, errors
			}

			if len(*json_array) != 1 {
				errors = append(errors, fmt.Errorf("count record does not exist"))
				return nil, errors
			}

			count, count_err := strconv.ParseUint(*((*json_array)[0].(Map).S("COUNT(*)")), 10, 64)
			if count_err != nil {
				errors = append(errors, count_err)
				return nil, errors
			}

			return &count, nil
		},
		CreateRecord: func(new_record_data Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := NewRecord(getTable(), new_record_data)
			if record_errors != nil {
				return nil, record_errors
			}

			create_record_errors := record.Create()
			if create_record_errors != nil {
				return  nil, create_record_errors
			}

			return record, nil
		},
		Select: func(filters Map, limit *uint64, offset *uint64) (*[]Record, []error) {
			var errors []error
			validate_errors := validate()
			if errors != nil {
				errors = append(errors, validate_errors...)
				return nil, errors
			}

			table_schema := getData()
			if filters != nil {
				filter_errors := ValidateGenericSpecial(filters, "SelectRecords")
				if filter_errors != nil {
					errors = append(errors, filter_errors...)
					return nil, errors
				}

				table_columns := getTableColumns()
				filter_columns := filters.Keys()
				for _, filter_column := range filter_columns {
					if Contains(table_columns, filter_column) {
						errors = append(errors, fmt.Errorf("SelectRecords: column: %s not found for table: %s", filter_column, *getTableName()))
					}
				}

				if len(errors) > 0 {
					return nil, errors
				}

				for _, filter_column := range filter_columns {
					filter_column_type := filters.GetType(filter_column)
					table_column_type := *((table_schema.M(filter_column)).S("type"))
					if table_column_type != filter_column_type {
						errors = append(errors, fmt.Errorf("SelectRecords: column filter: %s has data type: %s however table: %s has data type: %s", filter_column, filter_column_type, *getTableName(), table_column_type))

						//todo ignore if filter data_type is nil and table column allows nil
					}
				}
			}
			
			if len(errors) > 0 {
				return nil, errors
			}

			sql := fmt.Sprintf("SELECT * FROM %s ", EscapeString(*getTableName()))

			// todo add where clause

			if limit != nil {
				limit_value := strconv.FormatUint(*limit, 10)
				sql += fmt.Sprintf("LIMIT %s ", limit_value)
			}

			if offset != nil {
				offset_value := strconv.FormatUint(*offset, 10)
				sql += fmt.Sprintf("OFFSET %s ", offset_value)
			}
			sql += ";"

			if len(errors) > 0 {
				return nil, errors
			}

			json_array, sql_errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false})

			if sql_errors != nil {
				errors = append(errors, sql_errors...)
			}
		
			if len(errors) > 0 {
				return nil, errors
			}

			var mapped_records []Record
			for _, json := range (*json_array) {
				current_record := json.(Map)
				columns := current_record.Keys()
				mapped_record := Map{}
				for _, column := range columns {
					table_data_type := *((table_schema.M(column)).S("type"))
					switch table_data_type {
					case "*uint64", "uint64":
						value, value_errors := current_record.GetUInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetUInt64Value(column, *value)
						}
					case "*int64", "int64":
						value, value_errors := current_record.GetInt64(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt64(column, value)
						}
					case "*int", "int":
						value, value_errors := current_record.GetInt(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetInt(column, value)
						}
					case "*time.Time":
						value, value_errors := current_record.GetTime(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetTime(column, value)
						}
					case "*bool", "bool":
						value, value_errors := current_record.GetBool(column)
						if value_errors != nil {
							errors = append(errors, value_errors...)
						} else {
							mapped_record.SetBool(column, value)
						}
					default:
						errors = append(errors, fmt.Errorf("SelectRecords: table: %s column: %s mapping of data type: %s not supported please implement", *getTableName(), column, table_data_type))
					}
				}

				mapped_record_obj, mapped_record_obj_errors := NewRecord(getTable(), mapped_record)
				if mapped_record_obj_errors != nil {
					errors = append(errors, mapped_record_obj_errors...)
				} else {
					mapped_records = append(mapped_records, *mapped_record_obj)
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &mapped_records, nil
		},
		GetData: func() (Map) {
			return getData()
		},
		GetTableName: func() (*string) {
			return getTableName()
		},
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
    }
	setTable(&x)

	return &x, nil
}
