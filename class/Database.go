package class

import (
	"fmt"
	"bytes"
	"os/exec"
	"reflect"
	"strings"
	"runtime"
)

func GET_DATABASE_DATA_DEFINITION_STATEMENTS() ([]string) {
	return []string{GET_DATA_DEFINTION_STATEMENT_CREATE()}
}

func GET_DATABASE_LOGIC_OPTIONS_CREATE() ([][]string){
	return [][]string{GET_LOGIC_STATEMENT_IF_NOT_EXISTS()}
}

func GET_DATABASE_OPTIONS() (map[string]map[string][][]string) {
	var root = make(map[string]map[string][][]string)
	
	var logic_options = make(map[string][][]string)
	logic_options[GET_DATA_DEFINTION_STATEMENT_CREATE()] = GET_DATABASE_LOGIC_OPTIONS_CREATE()

	root[GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options

	return root
}

type Database struct {
	host *Host
	credentials *Credentials
    database_name *string
	database_create_options *DatabaseCreateOptions
	options map[string]map[string][][]string
	validation_functions map[string]func() *[]error
}

func NewDatabase(host *Host, credentials *Credentials, database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database) {
	x := Database{host: host, credentials: credentials, database_name: database_name, database_create_options: database_create_options, options: options}
	
	x.validation_functions = make(map[string]func() *[]error)
	x.InitValidationFunctions()

	return &x
}

func (this *Database) InitValidationFunctions() ()  {
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateHost"] = (*this).validateHost
	validation_functions["validateCredentials"] = (*this).validateCredentials
	validation_functions["validateDatabaseName"] = (*this).validateDatabaseName
	validation_functions["validateDatabaseCreateOptions"] = (*this).validateDatabaseCreateOptions
	validation_functions["validateOptions"] = (*this).validateOptions
	validation_functions["validateValidationFunctions"] = (*this).validateValidationFunctions

	if validation_functions["validateValidationFunctions"] == nil|| 
	   GetFunctionName(validation_functions["validateValidationFunctions"]) != GetFunctionName((*this).validateValidationFunctions) {
		panic(fmt.Errorf("validateValidationFunctions validation method not found potential sql injection without it"))
	}
}

func (this *Database) Create() (*Database, *string, *[]error)  {
	this, result, errors := (*this).createDatabase()
	if errors != nil {
		return nil, result, errors
	}

	return this, result, nil
}

func (this *Database) getValidationFunctions() map[string]func() *[]error {
	return (*this).validation_functions
}

func GetValidationMethodNameForFieldName(fieldName string) string {
	varName := fieldName
	varName = strings.Replace(varName, "_", " ", -1)
	varName = strings.Title(strings.ToLower(varName))
	varName = strings.Replace(varName, " ", "", -1)
	varName = "validate" + varName
	return varName
}

func (this *Database) Validate() *[]error {
	var errors *[]error 
	var fieldsNotFound []string
	reflected_value := reflect.ValueOf(this)
	refected_element := reflected_value.Elem()
	
    for i := 0; i < refected_element.NumField(); i++ {
		field := refected_element.Type().Field(i)
		validationMethodName := GetValidationMethodNameForFieldName(field.Name)

		method, found_method := (*this).getValidationFunctions()[validationMethodName]
		if !found_method {
			fieldsNotFound = append(fieldsNotFound, field.Name)
		} else {
			relection_errors := method()
			if relection_errors != nil{
				*errors = append(*errors, *relection_errors...)
			}
		}
	}

	for _, value := range fieldsNotFound {
		if !IsUpper(value) {
			*errors = append(*errors, fmt.Errorf("validation method: %s not found for %s please add to InitValidationFunctions", GetValidationMethodNameForFieldName(value), value))	
		}
	}

	if len(*errors) > 0 {
		return errors
	}

	return nil
}

func (this *Database) validateDatabaseCreateOptions()  (*[]error) {
	if (*this).GetDatabaseCreateOptions() == nil {
		return nil
	}


	return (*(*this).GetDatabaseCreateOptions()).Validate()
}

func (this *Database) validateOptions() (*[]error) {
	return ValidateOptions((*this).GetOptions(), reflect.ValueOf(*this))
}


func (this *Database) validateHost()  (*[]error) {
	var errors *[]error 
	if (*this).GetHost() == nil {
		*errors = append(*errors, fmt.Errorf("host is nil"))
		return errors
	}

	return (*((*this).GetHost())).Validate()
}

func (this *Database) validateCredentials()  (*[]error) {
	var errors *[]error 
	if (*this).GetCredentials() == nil {
		*errors = append(*errors, fmt.Errorf("credentials is nil"))
		return errors
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *Database) validateDatabaseName() (*[]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return ValidateCharacters(VALID_CHARACTERS, (*this).database_name, "database_name",  reflect.ValueOf(*this))
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (this *Database) validateValidationFunctions() (*[]error) {
	var errors *[]error 
	current := (*this).getValidationFunctions()
	compare := make(map[string]func() *[]error)
	found := false

    for current_key, current_value := range current {
		found = false
		for compare_key, compare_value := range compare {
			if GetFunctionName(current_value) == GetFunctionName(compare_value) && 
			   current_key != compare_key {
				found = true
				*errors = append(*errors, fmt.Errorf("key %s and key %s contain duplicate validation functions %s",  current_key, compare_key, current_value))
				break
			}
		}

		if !found {
			compare[current_key] = current_value
		}
    }

	if len(*errors) > 0 {
		return errors
	}

	return nil
}

func (this *Database) GetDatabaseName() *string {
	return (*this).database_name
}

func (this *Database) createDatabase() (*Database, *string, *[]error) {
	var errors *[]error 
	crud_sql_command, crud_command_errors := (*this).getCLSCRUDDatabaseCommand(GET_DATA_DEFINTION_STATEMENT_CREATE(), (*this).GetOptions())

	if crud_command_errors != nil {
		*errors = append(*errors, *crud_command_errors...)	
	}

	if len(*errors) > 0 {
		return nil, nil, errors
	}

	var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", *crud_sql_command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    command_err := cmd.Run()

    if command_err != nil {
		*errors = append(*errors, command_err)	
	}

	shell_ouput := ""

	if len(*errors) > 0 {
		shell_ouput = stderr.String()
		return nil, &shell_ouput, errors
	}

	shell_ouput = stdout.String()
    return this, &shell_ouput, nil
}


func (this *Database) getCLSCRUDDatabaseCommand(command string, options map[string]map[string][][]string) (*string, *[]error) {
	var errors *[]error 

	command_errs := ContainsExactMatch(GET_DATABASE_DATA_DEFINITION_STATEMENTS(), &command, "command", fmt.Sprintf("%T", *this))

	if command_errs != nil {
		*errors = append(*errors, *command_errs...)	
	}

	database_errs := (*this).Validate()

	if database_errs != nil {
		*errors = append(*errors, *database_errs...)	
	}

	logic_option, logic_options_errs := GetLogicCommand(command, GET_LOGIC_STATEMENT_FIELD_NAME(), GET_DATABASE_OPTIONS(), options, reflect.ValueOf(*this))
	if logic_options_errs != nil {
		*errors = append(*errors, *logic_options_errs...)	
	}

	host_command, host_command_errors := (*(*this).GetHost()).GetCLSCommand()
	if host_command_errors != nil {
		*errors = append(*errors, *host_command_errors...)	
	}

	credentials_command, credentials_command_errors := (*(*this).GetCredentials()).GetCLSCommand()
	if credentials_command_errors != nil {
		*errors = append(*errors, *credentials_command_errors...)	
	}

	if len(*errors) > 0 {
		return nil, errors
	}

	sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", *host_command, *credentials_command) 
	sql_command += fmt.Sprintf(" -e \"%s DATABASE ", command)
	
	if *logic_option != "" {
		sql_command += fmt.Sprintf("%s ", *logic_option)
	}
	
	sql_command += fmt.Sprintf("%s ", (*(*this).GetDatabaseName()))
	
	database_create_options_command, database_create_options_command_errs := (*this).GetDatabaseCreateOptions().GetSQL()
	if database_create_options_command_errs != nil || len(*database_create_options_command_errs) > 0 {
		*errors = append(*errors, *database_create_options_command_errs...)
		return nil, errors
	}

	sql_command += *database_create_options_command

	sql_command += ";\""
	return &sql_command, nil
}

func (this *Database) GetHost() *Host {
	return (*this).host
}

func (this *Database) GetCredentials() *Credentials {
	return (*this).credentials
}

func (this *Database) GetDatabaseCreateOptions() *DatabaseCreateOptions {
	return (*this).database_create_options
}

func (this *Database) GetOptions() map[string]map[string][][]string {
	return (*this).options
}