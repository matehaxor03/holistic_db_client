package class

import (
	"fmt"
	"bytes"
	"reflect"
	"strings"
	"os/exec"
)

type User struct {
	client *Client
	credentials *Credentials
	domain_name *DomainName

	extra_options map[string][]string
	validation_functions map[string]func() []error

	DATA_DEFINITION_STATEMENT_CREATE string
	DATA_DEFINITION_STATEMENTS []string

	LOGIC_OPTION_FIELD string
	LOGIC_OPTION_IF_NOT_EXISTS []string
	LOGIC_OPTION_CREATE_OPTIONS []string
}

func newUser(client *Client, credentials *Credentials, domain_name *DomainName, extra_options map[string][]string) (*User) {
	x := User{client: client,
				credentials: credentials,
				domain_name: domain_name}
				
	x.DATA_DEFINITION_STATEMENT_CREATE = "CREATE"
	x.DATA_DEFINITION_STATEMENTS = []string{x.DATA_DEFINITION_STATEMENT_CREATE}

	x.LOGIC_OPTION_FIELD = "LOGIC"
	x.LOGIC_OPTION_IF_NOT_EXISTS = []string{"IF", "NOT", "EXISTS"}
	x.LOGIC_OPTION_CREATE_OPTIONS = append(x.LOGIC_OPTION_CREATE_OPTIONS, x.LOGIC_OPTION_IF_NOT_EXISTS...)
	
	x.validation_functions = make(map[string]func() []error)
	x.InitValidationFunctions()
	return &x
}

func (this *User) Create() (*User, *string, []error)  {
	this, result, errors := (*this).createUser()
	if errors != nil {
		return nil, result, errors
	}

	return this, result, nil
}

func (this *User) InitValidationFunctions() ()  {
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateClient"] = (*this).validateClient
	validation_functions["validateCredentials"] = (*this).validateCredentials
	validation_functions["validateDomainName"] = (*this).validateDomainName
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

func (this *User) GetDomainName() *DomainName {
	return (*this).domain_name
}

func (this *User) validateConstants()  ([]error) {
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

		character_errors := ValidateCharacters(VALID_CHARACTERS, &string_fieldValue, fieldName, reflect.ValueOf(*this).Kind())
		if character_errors != nil {
			errors = append(errors, character_errors...)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *User) validateValidationFunctions() ([]error) {
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

func (this *User) validateExtraOptions()  ([]error) {
	var errors []error 
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	for key, value := range (*this).GetExtraOptions() {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("%s extra_options key %s", reflect.ValueOf(*this).Kind(), key),  reflect.ValueOf(*this).Kind())
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		var combined = strings.Join(value, "")
		value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("%s extra_options value %s",reflect.ValueOf(*this).Kind(), key),  reflect.ValueOf(*this).Kind())
		if value_extra_options_errors != nil {
			errors = append(errors, value_extra_options_errors...)	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *User) validateDomainName()  ([]error) {
	var errors []error 
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.1234567890"
	for key, value := range (*this).GetExtraOptions() {
		key_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &key, fmt.Sprintf("extra_options key %s", key),  reflect.ValueOf(*this).Kind())
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)	
		}

		var combined = strings.Join(value, "")
		value_extra_options_errors := ValidateCharacters(VALID_CHARACTERS, &combined, fmt.Sprintf("extra_options value %s", key),  reflect.ValueOf(*this).Kind())
		if value_extra_options_errors != nil {
			errors = append(errors, value_extra_options_errors...)	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *User) createUser() (*User, *string, []error) {
	var errors []error 
	crud_sql_command, crud_command_errors := (*this).getCLSCRUDUserCommand((*this).DATA_DEFINITION_STATEMENT_CREATE, (*this).GetExtraOptions())

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

func (this *User) getCLSCRUDUserCommand(command string, options map[string][]string) (*string, []error) {
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

	host_command, host_command_errors := (*(*this).GetClient().GetHost()).GetCLSCommand()
	if host_command_errors != nil {
		errors = append(errors, host_command_errors...)	
	}

	credentials_command, credentials_command_errors := (*(*this).GetClient().GetCredentials()).GetCLSCommand()
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
	
	sql_command += fmt.Sprintf("'%s' ", (*(*this).GetCredentials().GetUsername()))
	sql_command += fmt.Sprintf("@'%s' ", (*(*this).GetDomainName().GetDomainName()))
	sql_command += fmt.Sprintf("IDENTIFIED BY ")
	sql_command += fmt.Sprintf("'%s' ", ((*(*this).GetCredentials()).GetPassword()))

	sql_command += ";\""
	return &sql_command, nil
}

func (this *User) Validate() []error {
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

func (this *User) getValidationFunctions() map[string]func() []error {
	return (*this).validation_functions
}

func (this *User) validateClient()  ([]error) {
	var errors []error 
	if (*this).GetClient() == nil {
		errors = append(errors, fmt.Errorf("client is nil"))
		return errors
	}

	return (*((*this).GetClient())).Validate()
}

func (this *User) validateCredentials()  ([]error) {
	var errors []error 
	if (*this).GetCredentials() == nil {
		errors = append(errors, fmt.Errorf("credentials is nil"))
		return errors
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *User) GetCredentials() *Credentials {
	return (*this).credentials
}

func (this *User) GetExtraOptions() map[string][]string {
	return (*this).extra_options
}

func (this *User) GetClient() *Client {
	return (*this).client
}
