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
	ExecuteUnsafeCommand func(database Database, sql_command *string, options *json.Map) (*json.Array, []error)
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
		ExecuteUnsafeCommand: func(database Database, sql_command *string, options *json.Map) (*json.Array, []error) {
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

			host := database.GetHost()
			if host_errors := host.Validate(); host_errors != nil {
				return nil, host_errors
			}

			if len(errors) > 0 {
				return nil, errors
			}

			database_username := database.GetDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is not set"))
			} else if *database_username == "" {
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
			
			if len(errors) > 0 {
				return nil, errors
			}

			database_name := database.GetDatabaseName()
			database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
			if database_name_escaped_errors != nil {
				errors = append(errors, database_name_escaped_errors)
			}

			if len(errors) > 0 {
				return nil, errors
			}

			if *database_username != "root" {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "#" + database_name + "#" + (*database_username) + ".config"
			} else {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "##" + (*database_username) + ".config"
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

			if options.IsBoolTrue("use_mysql_database") {
				sql += fmt.Sprintf("USE %s;\n", "mysql")
			} else {
				if !(options.IsBoolTrue("creating_database") || options.IsBoolTrue("deleting_database") || options.IsBoolTrue("checking_database_exists") || options.IsBoolTrue("updating_database_global_settings")) {
					sql += fmt.Sprintf("USE %s;\n", database_name_escaped)
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

			//fmt.Println(command)
			shell_output, bash_errors := bashCommand.ExecuteUnsafeCommand(command, nil, nil)

			if sql_command_use_file {
				os.Remove(filename)
			}

			if bash_errors != nil {
				errors = append(errors, bash_errors...)
			}

			/*
			if shell_output != nil {
				fmt.Println(*shell_output)
			}

			if len(errors) > 0 {
				fmt.Println(errors)
			}*/

			if len(errors) > 0 {
				//fmt.Println(command)
				return nil, errors
			}

			records := json.NewArrayValue()

			if shell_output == nil || len(*shell_output) == 0 {
				return &records, nil
			}

			if options.IsBoolTrue("read_no_records") {
				return &records, nil
			}

			reading_columns := true
			value := ""
			columns_count := 0
			columns := json.NewArray()
			record := json.NewMap()
			for _, shell_row := range *shell_output {
				shell_row = strings.TrimSpace(shell_row)
				current_row_rune := []rune(shell_row)
				current_row_length := len(current_row_rune)
				for i := 0; i < current_row_length; i++ {
					current_value := string(current_row_rune[i])
					if reading_columns {
						if i == current_row_length - 1 {
							value = value + current_value
							columns.AppendStringValue(value)
							value = ""
							reading_columns = false
						} else if current_value == "\t" {
							columns.AppendStringValue(value)
							value = ""
						} else {
							value = value + current_value
						}
					} else {
						if i == current_row_length - 1  {
							value = value + current_value
							column_name, column_name_errors := columns.GetStringValue(columns_count)
							if column_name_errors != nil {
								errors = append(errors, column_name_errors...)
								continue
							} else if common.IsNil(column_name) {
								errors = append(errors,	fmt.Errorf("column_name is nil"))
								continue
							}
							record.SetStringValue(column_name, value)
							records.AppendMap(record)
							record = json.NewMap()
							value = ""
							columns_count = 0
						} else if current_value == "\t" {
							column_name, column_name_errors := columns.GetStringValue(columns_count)
							if column_name_errors != nil {
								errors = append(errors, column_name_errors...)
								continue
							} else if common.IsNil(column_name) {
								errors = append(errors,	fmt.Errorf("column name is nil"))
								continue
							}
							record.SetStringValue(column_name, value)
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
