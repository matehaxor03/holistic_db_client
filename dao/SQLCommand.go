package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(database *Database, sql_command *string, options *json.Map) (*json.Array, []error)
}

func newSQLCommand() (*SQLCommand, []error) {
	var errors []error
	bashCommand := common.NewBashCommand()

	directory_parts := common.GetDataDirectory()
	directory := "/" 
	for index, directory_part := range directory_parts {
		directory += directory_part
		if index < len(directory_parts) - 1 {
			directory += "/"
		}
	}

	x := SQLCommand{
		ExecuteUnsafeCommand: func(database *Database, sql_command *string, options *json.Map) (*json.Array, []error) {
			var errors []error

			if common.IsNil(database) {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			if common.IsNil(sql_command) {
				errors = append(errors, fmt.Errorf("sql_command is nil"))
			}

			if common.IsNil(options) {
				options = json.NewMap()
			}

			if len(errors) > 0 {
				return nil, errors
			}

			validate_errors := database.Validate()
			if validate_errors != nil {
				errors = append(errors, validate_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			host, host_errors := database.GetHost()
			if host_errors != nil {
				errors = append(errors, host_errors...)
			} else if common.IsNil(host) {
				errors = append(errors, fmt.Errorf("error: host is nil"))
			} else {
				host_errs := host.Validate()
				if host_errs != nil {
					errors = append(errors, host_errs...)
				}
			}

			database_username, database_username_errors := database.GetDatabaseUsername()
			if database_username_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting database username: %s", fmt.Sprintf("%s", database_username_errors)))
			} else if common.IsNil(database_username) {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is nil"))
			} else if database_username == "" {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is empty string"))
			}

			if sql_command == nil {
				errors = append(errors, fmt.Errorf("error: sql command is nil"))
			} else if *sql_command == "" {
				errors = append(errors, fmt.Errorf("error: sql command is an empty string"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			host_name, host_name_errors := host.GetHostName()
			if host_name_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting hostname: %s", fmt.Sprintf("%s", host_name_errors)))
			}

			port_number, port_number_errors := host.GetPortNumber()
			if port_number_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting port number: %s", fmt.Sprintf("%s", port_number_errors)))
			}

			host_command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", host_name, port_number)
			credentials_command := ""

			database_name, database_name_errors := (*database).GetDatabaseName()
			if database_name_errors != nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors getting database_name: %s", fmt.Sprintf("%s", database_name_errors)))
			} else if common.IsNil(database_name) {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors database_name is nil"))
			} else if database_name == "" {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand had errors database_name is empty string"))
			} 
			
			if len(errors) > 0 {
				return nil, errors
			}

			if database_username != "root" {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "#" + database_name + "#" + (database_username) + ".config"
			} else {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "##" + (database_username) + ".config"
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

			shell_output, bash_errors := bashCommand.ExecuteUnsafeCommand(command, nil, nil)

			if sql_command_use_file {
				os.Remove(filename)
			}

			if bash_errors != nil {
				errors = append(errors, bash_errors...)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			records := json.NewArrayValue()

			if shell_output == nil || len(*shell_output) == 0 {
				return &records, nil
			}


			//rune_array := []rune(strings.Join(*shell_output, "\n"))
			reading_columns := true
			value := ""
			columns_count := 0
			columns := json.NewArrayValue()
			record := json.NewMapValue()
			for _, shell_row := range *shell_output {
				shell_row = strings.TrimSpace(shell_row)
				current_row_rune := []rune(shell_row)
				current_row_length := len(current_row_rune)
				for i := 0; i < current_row_length; i++ {
					current_value := string(current_row_rune[i])
					if reading_columns {
						if i == current_row_length - 1 {
							value = value + current_value
							column_name := common.CloneString(&value)
							columns.AppendStringValue(*column_name)
							value = ""
							reading_columns = false
						} else if current_value == "\t" {
							column_name := common.CloneString(&value)
							columns.AppendStringValue(*column_name)
							value = ""
						} else {
							value = value + current_value
						}
					} else {
						if i == current_row_length - 1  {
							value = value + current_value
							column_value := common.CloneString(&value)
							x, x_errors := columns.GetStringValue(columns_count)
							if x_errors != nil {
								errors = append(errors, x_errors...)
								continue
							} else if common.IsNil(x) {
								errors = append(errors,	fmt.Errorf("SQLCommand x is nil"))
								continue
							}
							record.SetStringValue(x, *column_value)
							records.AppendMapValue(record)
							record = json.NewMapValue()
							value = ""
							columns_count = 0
						} else if current_value == "\t" {
							column_value := common.CloneString(&value)
							y, y_errors := columns.GetStringValue(columns_count)
							if y_errors != nil {
								errors = append(errors, y_errors...)
								continue
							} else if common.IsNil(y) {
								errors = append(errors,	fmt.Errorf("SQLCommand y is nil"))
								continue
							}

							record.SetStringValue(y, *column_value)
							columns_count += 1
							value = ""
						} else {
							value = value + current_value
						}
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
