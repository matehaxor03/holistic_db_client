package class


import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type SQLCommand struct {
	ExecuteUnsafeCommand func(host *Host, database *Database, credentials *Credentials, sql_command *string, sql_command_use_file bool) (*string, *string, []error)
}

func newSQLCommand() (*SQLCommand) {			    
	bashCommand := newBashCommand()
	x := SQLCommand{
		ExecuteUnsafeCommand: func(host *Host, database *Database, credentials *Credentials, sql_command *string, sql_command_use_file bool) (*string, *string, []error) {
			var errors []error 

			host_errs := host.Validate()
			if host_errs != nil {
				errors = append(errors, host_errs...)
			}

			database_errs := database.Validate()
			if database_errs != nil {
				errors = append(errors, database_errs...)
			}

			credential_errs := credentials.Validate()
			if credential_errs != nil {
				errors = append(errors, credential_errs...)
			}

			if sql_command == nil {
				errors = append(errors, fmt.Errorf("sql command is nil"))
			} else if *sql_command == "" {
				errors = append(errors, fmt.Errorf("sql command is an empty string"))
			}

			host_command, host_command_errs := host.GetCLSCommand()
			if host_command_errs != nil {
				errors = append(errors, host_command_errs...)
			}

			if len(errors) > 0 {
				return nil, nil, errors
			}

			credentials_command := "--defaults-extra-file=./holistic-db-config-" +  *(host.GetHostName()) + "-" + *(host.GetPortNumber()) + "-" + *(credentials.GetUsername()) + "-" + *((*database).GetDatabaseName()) + ".config"
			sql_header_command := fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", credentials_command, *host_command) 

			uuid, _ := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
			filename := fmt.Sprintf("%v%s.sql", time.Now().UnixNano(), string(uuid))
			command := ""

			if sql_command_use_file {
				ioutil.WriteFile(filename, []byte(*sql_command), 0600)
				command = sql_header_command + " < " + filename
			} else {
				command = sql_header_command + " -e \"" + *sql_command + "\""
			}

			shell_output, shell_output_errs, bash_errors := bashCommand.ExecuteUnsafeCommand(&command)

			if sql_command_use_file {
				os.Remove(filename)
			}
			
			if bash_errors != nil {
				errors = append(errors, bash_errors...)	
				return shell_output, shell_output_errs, errors
			}

			return shell_output, shell_output_errs, nil
		},
    }

	return &x
}