package class

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(client *Client, sql_command *string, options Map) (*Array, []error)
}

func NewSQLCommand() *SQLCommand {
	bashCommand := newBashCommand()
	x := SQLCommand{
		ExecuteUnsafeCommand: func(client *Client, sql_command *string, options Map) (*Array, []error) {
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
				return nil, errors
			}

			host, host_errors := client.GetHost()
			if host_errors != nil {
				errors = append(errors, host_errors...)
			} else if host == nil {
				errors = append(errors, fmt.Errorf("host is nil"))
			} else {
				host_errs := host.Validate()
				if host_errs != nil {
					errors = append(errors, host_errs...)
				}
			}

			database_username, _ := client.GetDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("database_username is nil"))
			}

			database, database_errors := client.GetDatabase()
			if database_errors != nil {
				errors = append(errors, database_errors...)
			} else if database != nil {
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

			directory, directory_errors := GetDirectoryOfExecutable()
			if directory_errors != nil {
				errors = append(errors, directory_errors)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			host_name, host_name_errors := (*host).GetHostName()
			if host_name_errors != nil {
				errors = append(errors, host_name_errors...)
			}

			port_number, port_number_errors := (*host).GetPortNumber()
			if port_number_errors != nil {
				errors = append(errors, port_number_errors...)
			}

			host_command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", host_name, port_number)
			credentials_command := ""

			if database != nil {
				database_name, database_name_errors := (*database).GetDatabaseName()
				if database_name_errors != nil {
					errors = append(errors, database_name_errors...)
				} else {
					credentials_command = "--defaults-extra-file=" + *directory + "/holistic_db_config:" + host_name + ":" + port_number + ":" + database_name + ":" + (*database_username) + ".config"
				}
			} else {
				credentials_command = "--defaults-extra-file=" + *directory + "/holistic_db_config:" + host_name + ":" + port_number + "::" + (*database_username) + ".config"
			}

			if len(errors) > 0 {
				return nil, errors
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
				database_name, database_name_errors := (*database).GetDatabaseName()
				if database_name_errors != nil {
					errors = append(errors, database_name_errors...)
				} else {
					if !(options.IsBoolTrue("creating_database") || options.IsBoolTrue("deleting_database") || options.IsBoolTrue("checking_database_exists")) {
						sql = fmt.Sprintf("USE %s;\n", database_name)
					}
				}
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

			if sql_command_use_file {
				ioutil.WriteFile(filename, []byte(sql), 0600)
				command = sql_header_command + " < " + filename
			} else {
				command = sql_header_command + " <<[END]\n " + sql + "\n[END]"
			}

			if len(errors) > 0 {
				return nil, errors
			}

			shell_output, bash_errors := bashCommand.ExecuteUnsafeCommand(&command)

			if sql_command_use_file {
				os.Remove(filename)
			}

			if bash_errors != nil {
				errors = append(errors, bash_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			records := Array{}

			if shell_output == nil || strings.TrimSpace(*shell_output) == "" {
				return &records, nil
			}

			rune_array := []rune(*shell_output)
			reading_columns := true
			value := ""
			columns_count := 0
			columns := Array{}
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
						records = append(records, record)
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
			return &records, nil
		},
	}

	return &x
}
