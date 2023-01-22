package dao

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	sql_generator_mysql "github.com/matehaxor03/holistic_db_client/sql_generators/community/mysql"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(database Database, raw_sql strings.Builder, options json.Map) (json.Array, []error)
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
		ExecuteUnsafeCommand: func(database Database, raw_sql strings.Builder, options json.Map) (json.Array, []error) {
			var errors []error
			const maxCapacity = 10*1024*1024  
			records := json.NewArrayValue()

			client := database.GetClient()
			client_manager := client.GetClientManager()
			host := database.GetHost()

			database_username := database.GetDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is not set"))
			} else if *database_username == "" {
				errors = append(errors, fmt.Errorf("error: SQLCommand.ExecuteUnsafeCommand database_username is empty string"))
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
			sql_command.WriteString(raw_sql.String())

			if options.IsBoolTrue("get_last_insert_id") {
				sql_command.WriteString(" SELECT LAST_INSERT_ID();")
			}

			if options.IsBoolTrue("transactional") {
				sql_command.WriteString("COMMIT;\n")
			}
			
			if len(errors) > 0 {
				return records, errors
			}

			stdout_lines, bash_errors := bashCommand.ExecuteUnsafeCommandUsingFiles(sql_header_command, sql_command.String())

		
			if bash_errors != nil {
				errors = append(errors, bash_errors...)
			}

			if len(errors) > 0 {
				fmt.Println(sql_command.String());
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
