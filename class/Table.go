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
	Create func() (*string, []error)
	GetTableName func() (*string)
	Count func() (*uint64, []error)
	CreateRecord func(record Map) (*Map, []error)
}

func NewTable(client *Client, schema Map, options map[string]map[string][][]string) (*Table, []error) {
	SQLCommand := newSQLCommand()
	var errors []error

	if schema == nil {
		errors = append(errors, fmt.Errorf("schema is nil"))
	}

	if !schema.HasKey("[table_name]") {
		errors = append(errors, fmt.Errorf("[table_name] field is nil"))
	} else if schema.GetType("[table_name]") != "class.Map" {
		errors = append(errors, fmt.Errorf("[table_name] field is not a map"))
	} else {
		boolean_value := true
		schema.M("[table_name]").SetBool("mandatory", &boolean_value)
		schema.M("[table_name]")[FILTERS()] = Array{ Map {"values":GetTableNameValidCharacters(),"function":getWhitelistCharactersFunc()}}
	}

	data := schema.Clone()

	getData := func() (Map) {
		return data.Clone()
	}

	getTableColumns := func() ([]string) {
		columns := getData().Keys() 
		var valid_columns []string
		for _, column := range columns {
			if strings.HasPrefix(column, "[") && strings.HasSuffix(column, "]") {
				continue
			}

			valid_columns = append(valid_columns, column)
		}
		return valid_columns
	}

	getTableName := func() (*string) {
		return CloneString(data.M("[table_name]").S("value"))
	}

	data["[client]"] = Map{"value":CloneClient(client),"mandatory":true}
	data["[options]"] = Map{"value":options,"mandatory":false}
	data["created_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}
	data["last_modified_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}

	{
		for _, column_name := range getTableColumns() {
			params := Map{"values": GetColumnNameValidCharacters(), "value":column_name, "label": column_name, "data_type": getTableName() }
			column_name_errors := WhitelistCharacters(params)
			if column_name_errors != nil {
				errors = append(errors, column_name_errors...)
			}

			if data.GetType(column_name) != "class.Map" {
				errors = append(errors, fmt.Errorf("column: %s for table: %s is not of type class.Map", column_name, *getTableName()))
			}
		}
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Table")
	}

	getClient := func() (*Client) {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("[options]").GetObject("value").(map[string]map[string][][]string)
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
		
		sql_command += fmt.Sprintf("%s ", *getTableName())

		valid_columns := getTableColumns()

		sql_command += "("
		for index, column := range valid_columns {
			columnSchema := data[column].(Map)

			if !columnSchema.HasKey("type") {
				errors = append(errors, fmt.Errorf("type field not found for column: " + column))
				continue
			}
			
			typeOf := columnSchema.S("type")
			if typeOf == nil {
				errors = append(errors, fmt.Errorf("type field had nil value for column: " + column))
				continue
			}
			
			switch *typeOf {
			case "*uint64", "*int64", "*int", "uint64", "uint", "int64", "int":
				sql_command += column + " BIGINT"
				

				if *typeOf == "*uint64" || 
				   *typeOf == "*uint" ||
				   *typeOf == "uint64" ||
				   *typeOf == "uint"  {
					sql_command += " UNSIGNED"
				}

				if columnSchema.HasKey("auto_increment") {
					if columnSchema.GetType("auto_increment") == "bool" {
						if *(columnSchema.B("auto_increment")) == true {
							sql_command += " AUTO_INCREMENT"
						} else if *(columnSchema.B("auto_increment")) == false {

						} else {
							errors = append(errors, fmt.Errorf("column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.B("auto_increment")))
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: auto_increment contained a value which is not a bool: %s", column, columnSchema.GetType("auto_increment")))
					}
				}

				if columnSchema.HasKey("primary_key") {
					if columnSchema.GetType("primary_key") == "bool" {
						if *(columnSchema.B("primary_key")) == true {
							sql_command += " PRIMARY KEY"
						} else if *(columnSchema.B("primary_key")) == false {

						} else {
							errors = append(errors, fmt.Errorf("column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.B("primary_key")))
						}
					} else {
						errors = append(errors, fmt.Errorf("column: %s for attribute: primary_key contained a value which is not a bool: %s", column, columnSchema.GetType("primary_key")))
					}		
				}

				if columnSchema.HasKey("default") && columnSchema.GetType("default") == "int" {
					sql_command += " DEFAULT " + strconv.FormatInt(*(columnSchema.GetInt64("default")), 10)
				}
			case "*time.Time":
				sql_command += column + " TIMESTAMP"
				if columnSchema.HasKey("default") {
					if columnSchema.S("default") == nil {
						errors = append(errors, fmt.Errorf("column: %s had nil default value", column))
						continue
					} else if *(columnSchema.S("default")) != "now" {
						errors = append(errors, fmt.Errorf("column: %s had default value it did not understand", column))
						continue
					}
					
					sql_command += " DEFAULT CURRENT_TIMESTAMP"
				} 
			default:
				errors = append(errors, fmt.Errorf("Table.getSQL type: %s is not supported please implement for column %s", *typeOf, column))
			}

			if index < (len(valid_columns) - 1) {
				sql_command += ", "
			}
		}
		sql_command += ")"

		if len(errors) > 0 {
			return nil, errors
		}

		sql_command += ";"

		return &sql_command, nil
	}

	createTable := func() (*string, []error) {
		var errors []error 
		sql_command, sql_command_errors := getSQL(GET_DATA_DEFINTION_STATEMENT_CREATE())
	
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}
	
		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": false})
	
		if *stderr != "" {
			if strings.Contains(*stderr, " table exists") {
				errors = append(errors, fmt.Errorf("create table failed most likely the table already exists"))
			} else {
				errors = append(errors, fmt.Errorf(*stderr))
			}
		}
	
		if len(errors) > 0 {
			return stdout, errors
		}
	
		return stdout, nil
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
		Clone: func() (*Table) {
			clone_value, _ := NewTable(getClient(), getData(), getOptions())
			return clone_value
		},
		Create: func() (*string, []error) {
			result, errors := createTable()
			if errors != nil {
				return result, errors
			}

			return result, nil
		},
		Count: func() (*uint64, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			sql :=  fmt.Sprintf("SELECT COUNT(*) FROM %s;", (*getTableName()))
			stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql, Map{"use_file": false, "no_column_headers": true})
						
			if *stderr != "" {
				if strings.Contains(*stderr, " table exists") {
					errors = append(errors, fmt.Errorf("create table failed most likely the table already exists"))
				} else {
					errors = append(errors, fmt.Errorf(*stderr))
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}

			count, count_err := strconv.ParseUint(string(strings.TrimSuffix(*stdout, "\n")), 10, 64)
			if count_err != nil {
				errors = append(errors, count_err)
				return nil, errors
			}

			return &count, nil
		},
		CreateRecord: func(record Map) (*Map, []error) {
			options := Map{"use_file": false, "no_column_headers": true, "get_last_insert_id": false}

			errors := validate()

			if record != nil {
				// add custom validation to fields


				record_errors := ValidateGenericSpecial(record, "Record")
				if record_errors != nil {
					errors = append(errors, record_errors...)
				}
			} else {
				errors = append(errors, fmt.Errorf("record is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			valid_columns := getTableColumns()
			record_columns := record.Keys()
			for _, record_column := range record_columns {
				if !Contains(valid_columns, record_column) {
					errors = append(errors, fmt.Errorf("column: %s does not exist for table: %s valid column names are: %s", record_column, *getTableName(), valid_columns))
				} else {
					if strings.HasPrefix(record_column, "credential_") {
						options["use_file"] = true
					}
				}

				type_of_schema_column := *((data.M(record_column)).S("type"))
				type_of_record_column := record.GetType(record_column)
				if type_of_record_column != type_of_schema_column {
					errors = append(errors, fmt.Errorf("table schema for column: %s has type: %s however record has type: %s", record_column, type_of_schema_column, type_of_record_column))
				}
			}

			auto_increment_column_name := ""
			auto_increment_columns := 0
			for _, valid_column := range valid_columns {
				column_definition := data.M(valid_column)
				
				if column_definition.HasKey("primary_key") &&
				   column_definition.GetType("primary_key") == "bool" &&
				   *(column_definition.B("primary_key")) &&
				   column_definition.HasKey("auto_increment") && 
			       column_definition.GetType("auto_increment") == "bool" &&
				   *(column_definition.B("auto_increment")) {
					options["get_last_insert_id"] = true
					auto_increment_column_name = valid_column
					auto_increment_columns += 1
				}
			}

			if auto_increment_columns > 1 {
				errors = append(errors, fmt.Errorf("table: %s can only have 1 auto_increment primary_key column, found: %s", *getTableName(), auto_increment_columns))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			sql := fmt.Sprintf("INSERT INTO %s ", *getTableName())
			sql += "("
			for index, record_column := range record_columns {
				sql += record_column
				if index < (len(record_columns) - 1) {
					sql += ", "
				}
			}
			sql += ") VALUES ("
			for index, record_column := range record_columns {
				rep := record.GetType(record_column)
				switch rep {
				default:
					errors = append(errors, fmt.Errorf("type: %s not supported for table please implement", rep))
				}
				
				if index < (len(record_columns) - 1) {
					sql += ", "
				}
			}
			sql += ");"

			if len(errors) > 0 {
				return nil, errors
			}

			stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), &sql, options)
						
			if *stderr != "" {
				if strings.Contains(*stderr, " some error") {
					errors = append(errors, fmt.Errorf("insert record failed"))
				} else {
					errors = append(errors, fmt.Errorf(*stderr))
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}

			if options["get_last_insert_id"].(bool) && auto_increment_column_name != "" {
				count, count_err := strconv.ParseUint(string(strings.TrimSuffix(*stdout, "\n")), 10, 64)
				if count_err != nil {
					errors = append(errors, count_err)
					return nil, errors
				}

				if auto_increment_column_name != "" {
					record[auto_increment_column_name] = count
				}
			}
			return &record, nil
		},
		GetTableName: func() (*string) {
			return getTableName()
		},
    }

	return &x, nil
}
