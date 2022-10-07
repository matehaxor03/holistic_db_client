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

func GetTableValidCharacters() *string {
	values :=  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._"
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
	GetTableName func() (string)
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
		schema.M("[table_name]")[FILTERS()] = Array{ Map {"values":GetTableValidCharacters(),"function":getValidateCharacters()}}
	}

	data := schema.Clone()
	data["[client]"] = Map{"value":CloneClient(client),"mandatory":true}
	data["[options]"] = Map{"value":options,"mandatory":false}
	data["created_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}
	data["last_modified_date"] = Map{"type":"*time.Time","value":nil,"mandatory":true, "default":"now"}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Table")
	}

	getClient := func() (*Client) {
		return CloneClient(data.M("[client]").GetObject("value").(*Client))
	}

	getTableName := func() (string) {
		return *(CloneString(data.M("[table_name]").S("value")))
	}

	getOptions := func() (map[string]map[string][][]string) {
		return data.M("[options]").GetObject("value").(map[string]map[string][][]string)
	}

	getData := func() (Map) {
		return data.Clone()
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

		command_errs := ContainsExactMatch(m)


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
		
		sql_command += fmt.Sprintf("%s ", getTableName())

		columns := data.Keys() 
		var valid_columns []string
		for _, column := range columns {
			if strings.HasPrefix(column, "[") && strings.HasSuffix(column, "]") {
				continue
			}

			valid_columns = append(valid_columns, column)
		}

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
			case "*string":


			case "*int64":
				sql_command += column + " BIGINT "
				if columnSchema.HasKey("primary_key") {
					sql_command += "  UNSIGNED AUTO_INCREMENT PRIMARY KEY"
				}

				if columnSchema.HasKey("default") {
					// todo check for safety that it's a number etc
					sql_command += " DEFAULT " + strconv.FormatInt(*(columnSchema.GetInt64("default")), 10)
				}
			case "*time.Time":
				sql_command += column + " TIMESTAMP "
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
				sql_command += ","
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
	
		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, true)
	
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
		GetTableName: func() (string) {
			return getTableName()
		},
    }

	return &x, nil
}
