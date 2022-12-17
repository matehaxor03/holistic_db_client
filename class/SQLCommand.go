package class

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(client Client, sql_command *string, options json.Map) (*json.Array, []error)
}

func newSQLCommand() (*SQLCommand, []error) {
	var errors []error
	bashCommand := newBashCommand()

	directory := "/Volumes/ramdisk"
	
	x := SQLCommand{
		ExecuteUnsafeCommand: func(client Client, sql_command *string, options json.Map) (*json.Array, []error) {
			var errors []error

			client_errs := client.Validate()
			if client_errs != nil {
				errors = append(errors, client_errs...)
			}
			
			if len(errors) > 0 {
				return nil, errors
			}

			host, host_errors := client.GetHost()
			if host_errors != nil {
				errors = append(errors, host_errors...)
			} else if host == nil {
				errors = append(errors, fmt.Errorf("error: host is nil"))
			} else {
				host_errs := host.Validate()
				if host_errs != nil {
					errors = append(errors, host_errs...)
				}
			}

			database_username, database_username_errors := client.GetDatabaseUsername()
			if database_username_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting database username: %s", fmt.Sprintf("%s", database_username_errors)))
			} else if database_username == nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is nil"))
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
				errors = append(errors, fmt.Errorf("error: sql command is nil"))
			} else if *sql_command == "" {
				errors = append(errors, fmt.Errorf("error: sql command is an empty string"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			host_name, host_name_errors := (*host).GetHostName()
			if host_name_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting hostname: %s", fmt.Sprintf("%s", host_name_errors)))
			}

			port_number, port_number_errors := (*host).GetPortNumber()
			if port_number_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting port number: %s", fmt.Sprintf("%s", port_number_errors)))
			}

			host_command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", host_name, port_number)
			credentials_command := ""

			if database != nil {
				database_name, database_name_errors := (*database).GetDatabaseName()
				if database_name_errors != nil {
					errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting database_name: %s", fmt.Sprintf("%s", database_name_errors)))
				} else {
					credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config:" + host_name + ":" + port_number + ":" + database_name + ":" + (*database_username) + ".config"
				}
			} else {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config:" + host_name + ":" + port_number + "::" + (*database_username) + ".config"
			}

			if len(errors) > 0 {
				return nil, errors
			}

			sql_command_use_file := true
			if options.IsBoolFalse("use_file") {
				sql_command_use_file = false
			}

			sql_header_command := fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", credentials_command, host_command)

			uuid, _ := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
			filename := directory + "/" + fmt.Sprintf("%v%s.sql", time.Now().UnixNano(), string(uuid))
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
					if !(options.IsBoolTrue("creating_database") || options.IsBoolTrue("deleting_database") || options.IsBoolTrue("checking_database_exists") || options.IsBoolTrue("updating_database_global_settings")) {
						if options.IsBoolTrue("use_file") {
							sql += fmt.Sprintf("USE `%s`;\n", database_name)
						} else {
							sql += fmt.Sprintf("USE \\`%s\\`;\n", database_name)
						}
					}
				}
			}

			sql += " " + *sql_command

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
				if sql_command_use_file {
					os.Remove(filename)
				}
				return nil, errors
			}

			shell_output, bash_errors := bashCommand.ExecuteUnsafeCommand(command)

			if sql_command_use_file {
				os.Remove(filename)
			}

			if bash_errors != nil {
				errors = append(errors, bash_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			records := json.Array{}

			if shell_output == nil || strings.TrimSpace(*shell_output) == "" {
				return &records, nil
			}

			rune_array := []rune(*shell_output)
			reading_columns := true
			value := ""
			columns_count := 0
			columns := json.Array{}
			record := json.Map{}
			for i := 0; i < len(rune_array); i++ {
				current_value := string(rune_array[i])
				if reading_columns {
					if current_value == "\n" {
						column_name := common.CloneString(&value)
						columns = append(columns, *column_name)
						value = ""
						reading_columns = false
					} else if current_value == "\t" {
						column_name := common.CloneString(&value)
						columns = append(columns, *column_name)
						value = ""
					} else {
						value = value + current_value
					}
				} else {
					if current_value == "\n" {
						column_value := common.CloneString(&value)
						record.SetString(columns[columns_count].(string), column_value)
						records = append(records, record)
						record = json.Map{}
						value = ""
						columns_count = 0
					} else if current_value == "\t" {
						column_value := common.CloneString(&value)
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

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
