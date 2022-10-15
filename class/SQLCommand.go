package class


import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strings"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(client *Client, sql_command *string, options Map) (*Array, *string, []error)
}

func NewSQLCommand() (*SQLCommand) {			    
	bashCommand := newBashCommand()
	x := SQLCommand{
		ExecuteUnsafeCommand: func(client *Client, sql_command *string, options Map) (*Array, *string, []error) {
			var errors []error 

			if client == nil {
				errors = append(errors, fmt.Errorf("client is nil"))
			} else {
				client_errs := client.Validate()
				if client_errs != nil {
					errors = append(errors, client_errs...)
				}
			}

			if len(errors) > 0 {
				return nil, nil, errors
			}

			host := client.GetHost()
			if host == nil {
				errors = append(errors, fmt.Errorf("host is nil"))
			} else {
				host_errs := host.Validate()
				if host_errs != nil {
					errors = append(errors, host_errs...)
				}
			}

			database_username := client.GetDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("database_username is nil"))
			} 
			
			database := client.GetDatabase()
			if database != nil {
				database_errs := database.Validate()
				if database_errs != nil {
					errors = append(errors, database_errs...)
				}
			}
			
			if sql_command == nil {
				errors = append(errors, fmt.Errorf("sql command is nil"))
			} else if *sql_command == "" {
				errors = append(errors, fmt.Errorf("sql command is an empty string"))
			}

			if len(errors) > 0 {
				return nil, nil, errors
			}

			host_command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", *(*(host)).GetHostName(), *(*(host)).GetPortNumber())
			credentials_command := ""
			
			if database != nil {
				credentials_command = "--defaults-extra-file=./holistic_db_config:" +  *((*host).GetHostName()) + ":" + *((*host).GetPortNumber()) + ":" + *((*database).GetDatabaseName()) + ":" + (*database_username) + ".config"
			} else {
				credentials_command = "--defaults-extra-file=./holistic_db_config:" +  *((*host).GetHostName()) + ":" + *((*host).GetPortNumber()) + "::" + (*database_username) + ".config"
			}
			
			sql_header_command := fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", credentials_command, host_command) 

			uuid, _ := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
			filename := fmt.Sprintf("%v%s.sql", time.Now().UnixNano(), string(uuid))
			command := ""

			sql := ""


			if options.IsBoolTrue("transactional") {
				sql += "START TRANSACTION;\n"
			}

			if database != nil {
				sql = fmt.Sprintf("USE %s;\n", (*(*database).GetDatabaseName()))
			}

			sql += " " + *sql_command

			sql_command_use_file := true
			if options.IsBoolFalse("use_file") {
				sql_command_use_file = false
			}

			if options.IsBoolTrue("get_last_insert_id") {
				sql += " SELECT LAST_INSERT_ID();"
			}

			if options.IsBoolTrue("transactional") {
				sql += "COMMIT;\n"
			}

			fmt.Println(sql)

			if sql_command_use_file {
				ioutil.WriteFile(filename, []byte(sql), 0600)
				command = sql_header_command + " < " + filename
			} else {
				command = sql_header_command + " -e '" + sql + "'"
			}

			fmt.Println(command)

			shell_output, shell_output_errs, bash_errors := bashCommand.ExecuteUnsafeCommand(&command)

			
			if sql_command_use_file {
				os.Remove(filename)
			}

			fmt.Println(*shell_output)
			fmt.Println(*shell_output_errs)
			
			if bash_errors != nil {
				errors = append(errors, bash_errors...)	
				return nil, shell_output_errs, errors
			}

			if shell_output_errs != nil && *shell_output_errs != "" {
				return nil, shell_output_errs, errors
			}

			if shell_output == nil || strings.TrimSpace(*shell_output) == "" {
				return nil, nil, nil
			}
			
			rune_array := []rune(*shell_output)
			reading_columns := true
			value := ""
			columns_count := 0
			columns := Array{}
			records := Array{}
			record := Map{}
			for i := 0; i < len(rune_array); i++ {
				current_value := string(rune_array[i])
				if reading_columns {
					if current_value == "\n" {
						column_name := CloneString(&value)
						columns = append(columns, *column_name)
						value = ""
						reading_columns = false
					} else if current_value == "\t" {
						column_name := CloneString(&value)
						columns = append(columns, *column_name)
						value = ""
					} else {
						value = value + current_value
					} 
				} else {
					if current_value == "\n" {
						column_value := CloneString(&value)
						record.SetString(columns[columns_count].(string), column_value)
						records = append(records, record.Clone())
						record = Map{}
						value = ""
						columns_count = 0
					} else if current_value == "\t" {
						column_value := CloneString(&value)
						record.SetString(columns[columns_count].(string), column_value)
						columns_count += 1
						value = ""
					} else {
						value = value + current_value
					} 
				}
			}
			return &records, shell_output_errs, nil
		},
    }

	return &x
}