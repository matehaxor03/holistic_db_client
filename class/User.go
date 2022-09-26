package class

import (
	"fmt"
	"bytes"
	"reflect"
	"os/exec"
)

type User struct {
	client *Client
	credentials *Credentials
	domain_name *DomainName

	options map[string]map[string][][]string
	validation_functions map[string]func() []error
}

func GET_USER_DATA_DEFINITION_STATEMENTS() Array {
	return Array{GET_DATA_DEFINTION_STATEMENT_CREATE()}
}

func GET_USER_LOGIC_OPTIONS_CREATE() ([][]string){
	return [][]string{GET_LOGIC_STATEMENT_IF_NOT_EXISTS()}
}

func GET_USER_EXTRA_OPTIONS() (map[string]map[string][][]string) {
	var root = make(map[string]map[string][][]string)
	
	var logic_options = make(map[string][][]string)
	logic_options[GET_DATA_DEFINTION_STATEMENT_CREATE()] = GET_USER_LOGIC_OPTIONS_CREATE()

	root[GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options

	return root
}

func newUser(client *Client, credentials *Credentials, domain_name *DomainName, options map[string]map[string][][]string) (*User, []error) {
	x := User{client: client,
				credentials: credentials,
				domain_name: domain_name,
			    options: options}
	x.validation_functions = make(map[string]func() []error)
	x.InitValidationFunctions()
	return &x, nil
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
	validation_functions["validateOptions"] = (*this).validateOptions
	validation_functions["validateValidationFunctions"] = (*this).validateValidationFunctions

	if validation_functions["validateValidationFunctions"] == nil|| 
	   GetFunctionName(validation_functions["validateValidationFunctions"]) != GetFunctionName((*this).validateValidationFunctions) {
		panic(fmt.Errorf("validateValidationFunctions validation method not found potential sql injection without it"))
	}
}

func (this *User) GetDomainName() *DomainName {
	return (*this).domain_name
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

func (this *User) validateOptions()  ([]error) {
	return ValidateOptions((*this).GetOptions(), reflect.ValueOf(*this))
}

func (this *User) createUser() (*User, *string, []error) {
	var errors []error 
	crud_sql_command, crud_command_errors := (*this).getCLSCRUDUserCommand(GET_DATA_DEFINTION_STATEMENT_CREATE(), (*this).GetOptions())

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

func (this *User) getCLSCRUDUserCommand(command string, options map[string]map[string][][]string) (*string, []error) {
	var errors []error 

	m := Map{}
	m.SetArray("values", GET_USER_DATA_DEFINITION_STATEMENTS())
	m.SetString("value", &command)
	commandTemp := "command"
	m.SetString("label", &commandTemp)
	rep :=  fmt.Sprintf("%T", *this)
	m.SetString("data_type", &rep)


	command_errs := ContainsExactMatch(m)

	if command_errs != nil {
		errors = append(errors, command_errs...)	
	}

	database_errs := (*this).Validate()

	if database_errs != nil {
		errors = append(errors, database_errs...)	
	}

	logic_option, logic_option_errs := GetLogicCommand(command, GET_LOGIC_STATEMENT_FIELD_NAME(), GET_USER_EXTRA_OPTIONS(), options, reflect.ValueOf(*this))
	if logic_option_errs != nil {
		errors = append(errors, logic_option_errs...)	
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
	
	if *logic_option != "" {
		sql_command += fmt.Sprintf("%s ", *logic_option)
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
		errors = append(errors, fmt.Errorf("Holistic.User: client is nil"))
		return errors
	}

	return (*((*this).GetClient())).Validate()
}

func (this *User) validateDomainName()  ([]error) {
	var errors []error 
	if (*this).GetClient() == nil {
		errors = append(errors, fmt.Errorf("Holistic.User: domain name is nil"))
		return errors
	}

	return (*((*this).GetDomainName())).Validate()
}

func (this *User) validateCredentials()  ([]error) {
	var errors []error 
	if (*this).GetCredentials() == nil {
		errors = append(errors, fmt.Errorf("Holistic.User: credentials is nil"))
		return errors
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *User) GetCredentials() *Credentials {
	return (*this).credentials
}

func (this *User) GetOptions() map[string]map[string][][]string {
	return (*this).options
}

func (this *User) GetClient() *Client {
	return (*this).client
}
