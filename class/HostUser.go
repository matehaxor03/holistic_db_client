package class

import (
	"fmt"
	"bytes"
	"reflect"
	"strings"
	"os/exec"
)

type HostUser struct {
	host *Host
	credentials *Credentials
	user_credentials *Credentials
	domain *Domain

	extra_options map[string][]string
	validation_functions map[string]func() []error

	DATA_DEFINITION_STATEMENT_CREATE string
	DATA_DEFINITION_STATEMENTS []string

	LOGIC_OPTION_FIELD string
	LOGIC_OPTION_IF_NOT_EXISTS []string
	LOGIC_OPTION_CREATE_OPTIONS []string
}

func newHostUser(host *Host, credentials *Credentials, user_credentials *Credentials, domain *Domain, extra_options map[string][]string) (*HostUser) {
	x := HostUser{host: host,
				credentials: credentials,
				user_credentials: user_credentials,
				domain: domain}
				
	x.DATA_DEFINITION_STATEMENT_CREATE = "CREATE"
	x.DATA_DEFINITION_STATEMENTS = []string{x.DATA_DEFINITION_STATEMENT_CREATE}

	x.LOGIC_OPTION_CREATE_OPTIONS = append(x.LOGIC_OPTION_CREATE_OPTIONS, x.LOGIC_OPTION_IF_NOT_EXISTS...)
	
	x.validation_functions = make(map[string]func() []error)
	x.InitValidationFunctions()
	return &x
}

func (this *HostUser) CreateHostUser(client *Client, user_credentials *Credentials, domain *Domain, options map[string][]string) (*HostUser, *string, []error) {
	return newHostUser((*client).GetHost(), (*client).GetCredentials(), user_credentials, domain, options).Create()
}

func (this *HostUser) InitValidationFunctions() ()  {
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateHost"] = (*this).validateHost
	validation_functions["validateCredentials"] = (*this).validateCredentials
	validation_functions["validateUserCredentials"] = (*this).validateUserCredentials
	validation_functions["validateDomain"] = (*this).validateDomain
	validation_functions["validateExtraOptions"] = (*this).validateExtraOptions
	validation_functions["validateValidationFunctions"] = (*this).validateValidationFunctions
	validation_functions["validateConstants"] = (*this).validateConstants

	if validation_functions["validateValidationFunctions"] == nil|| 
	   GetFunctionName(validation_functions["validateValidationFunctions"]) != GetFunctionName((*this).validateValidationFunctions) {
		panic(fmt.Errorf("validateValidationFunctions validation method not found potential sql injection without it"))
	}

	if validation_functions["validateConstants"] == nil|| 
	   GetFunctionName(validation_functions["validateConstants"]) != GetFunctionName((*this).validateConstants) {
		panic(fmt.Errorf("validateConstants validation method not found potential sql injection without it"))
	}
}

func (this *HostUser) GetHost() *Host {
	return (*this).host
 }

func (this *HostUser) GetDomain() *Domain {
	return (*this).domain
}

func (this *HostUser) validateConstants()  ([]error) {
	var errors []error 
	VALID_CHARACTERS := GetConstantValueAllowedCharacters()
	reflected_value := reflect.ValueOf(this)
	refected_element := reflected_value.Elem()
	string_fieldValue := ""

	for i := 0; i < refected_element.NumField(); i++ {
		string_fieldValue = ""
		field := refected_element.Type().Field(i)
		fieldName := field.Name
		if !IsUpper(fieldName) {
			continue
		}

		fieldValue := refected_element.FieldByName(fieldName)		
		if fieldValue.Kind().String() == "string" {
			string_fieldValue = fmt.Sprintf("%s", fieldValue)	
		} else if fieldValue.Kind().String() == "slice" {
			var array = fieldValue.Interface().([]string)
			for _, value := range array {
				string_fieldValue += fmt.Sprintf("%s", value)
			}
		} else {
			panic(fmt.Sprintf("please implement validation for constant value %s", fieldName))
		}

		character_errors := ValidateCharacters(VALID_CHARACTERS, &string_fieldValue, fieldName)
		if character_errors != nil {
			errors = append(errors, character_errors...)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostUser) validateValidationFunctions() ([]error) {
	var errors []error 
	current := (*this).getValidationFunctions()
	compare := make(map[string]func() []error)
	found := false

    for current_key, current_value := range current {
		found = false
		for compare_key, compare_value := range compare {
			if GetFunctionName(current_value) == GetFunctionName(compare_value) && 
			   current_key != compare_key {
				found = true
				errors = append(errors, fmt.Errorf("key %s and key %s contain duplicate validation functions %s",  current_key, compare_key, current_value))
				break
			}
		}

		if !found {
			compare[current_key] = current_value
		}
    }

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostUser) validateExtraOptions()  ([]error) {
	var errors []error 
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	for key, value := range (*this).GetExtraOptions() {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options key %s", key))
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		var combined = strings.Join(value, "")
		value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options value %s", key))
		if value_extra_options_errors != nil {
			errors = append(errors, value_extra_options_errors...)	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostUser) validateDomain()  ([]error) {
	var errors []error 
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.1234567890"
	for key, value := range (*this).GetExtraOptions() {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options key %s", key))
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		var combined = strings.Join(value, "")
		value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options value %s", key))
		if value_extra_options_errors != nil {
			errors = append(errors, value_extra_options_errors...)	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostUser) Create() (*HostUser, *string, []error)  {
	this, result, errors := (*this).createHostUser()
	if errors != nil {
		return nil, result, errors
	}

	return this, result, nil
}

func (this *HostUser) createHostUser() (*HostUser, *string, []error) {
	var errors []error 
	crud_sql_command, crud_command_errors := (*this).getCLSCRUDHostUserCommand((*this).DATA_DEFINITION_STATEMENT_CREATE, (*this).GetExtraOptions())

	if crud_command_errors != nil {
		errors = append(errors, crud_command_errors...)	
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", *crud_sql_command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    command_err := cmd.Run()

    if command_err != nil {
		errors = append(errors, command_err)	
	}

	shell_ouput := ""

	if len(errors) > 0 {
		shell_ouput = stderr.String()
		return nil, &shell_ouput, errors
	}

	shell_ouput = stdout.String()
    return this, &shell_ouput, nil
}

func (this *HostUser) getCLSCRUDHostUserCommand(command string, options map[string][]string) (*string, []error) {
	var errors []error 

	command_errs := Contains((*this).DATA_DEFINITION_STATEMENTS, &command, "command")

	if command_errs != nil {
		errors = append(errors, command_errs...)	
	}

	database_errs := (*this).Validate()

	if database_errs != nil {
		errors = append(errors, database_errs...)	
	}

	logic_option := ""
	if options != nil {
	    logic_option_value, logic_option_exists := options[(*this).LOGIC_OPTION_FIELD]
		if command == (*this).DATA_DEFINITION_STATEMENT_CREATE &&
		   logic_option_exists {
		    logic_option_errors := ArrayContainsArray((*this).LOGIC_OPTION_CREATE_OPTIONS, logic_option_value, "LOGIC")
			if logic_option_errors != nil {
				errors = append(errors, logic_option_errors...)	
			} else {
				logic_option = strings.Join(logic_option_value, " ")
			}
		}
	}

	host_command, host_command_errors := (*(*this).GetHost()).GetCLSCommand()
	if host_command_errors != nil {
		errors = append(errors, host_command_errors...)	
	}

	credentials_command, credentials_command_errors := (*(*this).GetCredentials()).GetCLSCommand()
	if credentials_command_errors != nil {
		errors = append(errors, credentials_command_errors...)	
	}

	if len(errors) > 0 {
		return nil, errors
	}

	sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", *host_command, *credentials_command) 
	sql_command += fmt.Sprintf(" -e \"%s USER ", command)
	
	if logic_option != "" {
		sql_command += fmt.Sprintf("%s ", logic_option)
	}
	
	sql_command += fmt.Sprintf("'%s' ", (*(*this).GetUserCredentials().GetUsername()))
	sql_command += fmt.Sprintf("@'%s' ", (*(*this).GetDomain().GetDomainName()))
	sql_command += fmt.Sprintf("IDENTIFIED BY ")
	sql_command += fmt.Sprintf("'%s' ", ((*(*this).GetUserCredentials()).GetPassword()))

	sql_command += ";\""
	return &sql_command, nil
}

func (this *HostUser) Validate() []error {
	var errors []error 
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
				errors = append(errors, relection_errors...)
			}
		}
	}

	method, found_method := (*this).getValidationFunctions()["validateConstants"]
	if !found_method {
		errors = append(errors, fmt.Errorf("validation method: validateConstants not found please add to InitValidationFunctions"))
	} else {
		constant_errors := method()
		if constant_errors != nil{
			errors = append(errors, constant_errors...)
		}
	}

	for _, value := range fieldsNotFound {
		if !IsUpper(value) {
			errors = append(errors, fmt.Errorf("validation method: %s not found for %s please add to InitValidationFunctions", GetValidationMethodNameForFieldName(value), value))	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostUser) getValidationFunctions() map[string]func() []error {
	return (*this).validation_functions
}

func (this *HostUser) validateHost()  ([]error) {
	var errors []error 
	if (*this).GetHost() == nil {
		errors = append(errors, fmt.Errorf("host is nil"))
		return errors
	}

	return (*((*this).GetHost())).Validate()
}

func (this *HostUser) validateCredentials()  ([]error) {
	var errors []error 
	if (*this).GetCredentials() == nil {
		errors = append(errors, fmt.Errorf("credentials is nil"))
		return errors
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *HostUser) validateUserCredentials()  ([]error) {
	var errors []error 
	if (*this).GetUserCredentials() == nil {
		errors = append(errors, fmt.Errorf("credentials is nil"))
		return errors
	}

	return (*((*this).GetUserCredentials())).Validate()
}

func (this *HostUser) GetCredentials() *Credentials {
	return (*this).credentials
}

func (this *HostUser) GetUserCredentials() *Credentials {
	return (*this).user_credentials
}

func (this *HostUser) GetExtraOptions() map[string][]string {
	return (*this).extra_options
}
