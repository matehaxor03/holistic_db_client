package dao

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"time"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(database Database, raw_sql *string, options json.Map) (json.Array, []error)
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
		ExecuteUnsafeCommand: func(database Database, raw_sql *string, options json.Map) (json.Array, []error) {
			var errors []error
			const maxCapacity = 10*1024*1024  
			records := json.NewArrayValue()



			if common.IsNil(database) {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			if common.IsNil(raw_sql) {
				errors = append(errors, fmt.Errorf("sql is nil"))
			}

			if len(errors) > 0 {
				return records, errors
			}

			validate_errors := database.Validate()
			if validate_errors != nil {
				errors = append(errors, validate_errors...)
			}

			if len(errors) > 0 {
				return records, errors
			}

			client := database.GetClient()
			client_manager := client.GetClientManager()
			host := database.GetHost()

			database_username := database.GetDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is not set"))
			} else if *database_username == "" {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is empty string"))
			}

			if *raw_sql == "" {
				errors = append(errors, fmt.Errorf("error: sql is an empty string"))
			}

			if len(errors) > 0 {
				return records, errors
			}

			host_name := host.GetHostName()
			port_number := host.GetPortNumber()
			

			host_command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", host_name, port_number)
			credentials_command := ""
			
			if len(errors) > 0 {
				return records, errors
			}

			database_name := database.GetDatabaseName()
			database_name_escaped, database_name_escaped_errors := common.EscapeString(database_name, "'")
			if database_name_escaped_errors != nil {
				errors = append(errors, database_name_escaped_errors)
			}

			if len(errors) > 0 {
				return records, errors
			}

			
			if *database_username != "root" {
				temp_database_username := ""
				temp_database_username = *database_username
				if temp_database_username == "holistic_w" || temp_database_username == "holistic_r" {
					temp_database_username += fmt.Sprintf("%d", client_manager.GetNextUserCount())
				}
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "#" + database_name + "#" + temp_database_username + ".config"
			} else {
				credentials_command = "--defaults-extra-file=" + directory + "/holistic_db_config#" + host_name + "#" + port_number + "##" + *database_username + ".config"
			}

			if len(errors) > 0 {
				return records, errors
			}

			sql_header_command := fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s --wait --quick ", credentials_command, host_command)

			uuid, _ := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
			time_now := time.Now().UnixNano()
			filename := directory + "/" + fmt.Sprintf("%v%s.sql", time_now, string(uuid))
			filename_stdout := directory + "/" + fmt.Sprintf("%v%s-stdout.sql", time_now, string(uuid))
			filename_stderr := directory + "/" + fmt.Sprintf("%v%s-stderr.sql", time_now, string(uuid))


			command := ""

			var sql_command strings.Builder

			if options.IsBoolTrue("transactional") {
				sql_command.WriteString("START TRANSACTION;\n")
			}

			if options.IsBoolTrue("use_mysql_database") {
				sql_command.WriteString("USE mysql;\n")
			} else {
				if !(options.IsBoolTrue("creating_database") || options.IsBoolTrue("deleting_database") || options.IsBoolTrue("checking_database_exists") || options.IsBoolTrue("updating_database_global_settings")) {
					sql_command.WriteString("USE ")
					sql_generator_mysql.Box(&sql_command, database_name_escaped, "`","`")
					sql_command.WriteString(";\n")
				}
			}
			
			sql_command.WriteString(" ")
			sql_command.WriteString(*raw_sql)

			if options.IsBoolTrue("get_last_insert_id") {
				sql_command.WriteString(" SELECT LAST_INSERT_ID();")
			}

			if options.IsBoolTrue("transactional") {
				sql_command.WriteString("COMMIT;\n")
			}

			sql := sql_command.String()
			ioutil.WriteFile(filename, []byte(sql), 0600)
			command = sql_header_command + " < " + filename +  " > " + filename_stdout + " 2> " + filename_stderr
			//fmt.Println(command)
			defer os.Remove(filename)
			defer os.Remove(filename_stdout)
			defer os.Remove(filename_stderr)
			
			
			if len(errors) > 0 {
				return records, errors
			}

			//fmt.Println(command)
			//fmt.Println(sql)
			_, bash_errors := bashCommand.ExecuteUnsafeCommand(command, nil, nil)

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

			var stdout_lines []string
			file_stdout, file_stdout_errors := os.Open(filename_stdout)
			if file_stdout_errors != nil {
				errors = append(errors, file_stdout_errors)
			} else {
				defer file_stdout.Close()
				stdout_scanner := bufio.NewScanner(file_stdout)
				stdout_scanner_buffer := make([]byte, maxCapacity)
				stdout_scanner.Buffer(stdout_scanner_buffer, maxCapacity)
				stdout_scanner.Split(bufio.ScanLines)
				for stdout_scanner.Scan() {
					current_text := stdout_scanner.Text()
					if current_text != "" {
						stdout_lines = append(stdout_lines, current_text)
					}
				}
			}

			file_stderr, file_stderr_errors := os.Open(filename_stderr)
			if file_stderr_errors != nil {
				errors = append(errors, file_stderr_errors)
			} else {
				defer file_stderr.Close()
				stderr_scanner := bufio.NewScanner(file_stderr)
				stderr_scanner_buffer := make([]byte, maxCapacity)
				stderr_scanner.Buffer(stderr_scanner_buffer, maxCapacity)
				stderr_scanner.Split(bufio.ScanLines)
				for stderr_scanner.Scan() {
					current_text := stderr_scanner.Text()
					if current_text != "" {
						errors = append(errors, fmt.Errorf("%s", current_text))
					}
				}
			}

			if len(errors) > 0 {
				//fmt.Println(command)
				//fmt.Println(fmt.Errorf("%s", errors))
				return records, errors
			}

			if options.IsBoolTrue("read_no_records") {
				return records, nil
			}

			reading_columns := true
			value := ""
			columns_count := 0
			columns := json.NewArray()
			record := json.NewMap()
			for _, shell_row := range stdout_lines {
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
			return records, nil
		},
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &x, nil
}
