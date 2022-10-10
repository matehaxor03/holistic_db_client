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
	GetTableColumns func() ([]string)
	Count func() (*uint64, []error)
	GetData func() (Map)
	CreateRecord func(record Map) (*Record, []error)
	GetDatabase func() (*Database)
}

func NewTable(database *Database, schema Map, options map[string]map[string][][]string) (*Table, []error) {
	var this_table *Table
	SQLCommand := newSQLCommand()
	var errors []error

	if schema == nil {
		errors = append(errors, fmt.Errorf("schema is nil"))
		return nil, errors
	}

	if !schema.HasKey("table_name") {
		errors = append(errors, fmt.Errorf("table_name field is nil"))
	} else if schema.GetType("table_name") != "class.Map" {
		errors = append(errors, fmt.Errorf("table_name field is not a map"))
	} else {
		boolean_value := true
		schema.M("table_name").SetBool("mandatory", &boolean_value)
		schema.M("table_name")[FILTERS()] = Array{ Map {"values":GetTableNameValidCharacters(),"function":getWhitelistCharactersFunc()}}
	}

	data := schema.Clone()
	data["database"] = Map{"value":CloneDatabase(database),"mandatory":true}
	data["options"] = Map{"value":options,"mandatory":false}
	data["created_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}
	data["last_modified_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}

	getData := func() (Map) {
		return data.Clone()
	}

	getTableName := func() (*string) {
		return CloneString(data.M("table_name").S("value"))
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

			if columnSchema.HasKey("value") {
				rep := columnSchema.GetType("value")
				switch rep {
					case "*uint64", "*int64", "*int", "uint64", "uint", "int64", "int", "*string", "string", "*time.Time", "time.Time", "*bool", "bool", "<nil>":
					default:
					continue
				}
			}
			columns = append(columns, column)
		}
		return columns
	}

	

	validate := func() ([]error) {
		return ValidateGenericSpecial(getData(), "Table")
	}

	getDatabase := func() (*Database) {
		return CloneDatabase(data.M("database").GetObject("value").(*Database))
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("options").GetObject("value").(map[string]map[string][][]string)
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
				sql_command += EscapeString(column) + " BIGINT"
				

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
	
		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), sql_command, Map{"use_file": false})
	
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

			sql :=  fmt.Sprintf("SELECT COUNT(*) FROM %s;", EscapeString((*getTableName())))
			stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getDatabase().GetClient(), &sql, Map{"use_file": false, "no_column_headers": true})
						
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
		CreateRecord: func(data Map) (*Record, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}

			record, record_errors := NewRecord(getTable(), data)
			if record_errors != nil {
				return nil, record_errors
			}

			_, create_record_errors := record.Create()
			if create_record_errors != nil {
				return nil, create_record_errors
			}

			return record, nil
		},
		GetData: func() (Map) {
			return getData()
		},
		GetTableName: func() (*string) {
			return getTableName()
		},
    }
	setTable(&x)

	return &x, nil
}
