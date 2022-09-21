package class

import (
	"fmt"
	"reflect"
)

type Client struct {
	host *Host
	credentials *Credentials
	database *Database

	validation_functions map[string]func() []error
}

func NewClient(host *Host, credentials *Credentials, database *Database) (*Client) {
	x := Client{host: host,
				credentials: credentials,
				database: database}

	x.validation_functions = make(map[string]func() []error)		
	x.InitValidationFunctions()
			    
	return &x
}

func (this *Client) validateConstants()  ([]error) {
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

func (this *Client) InitValidationFunctions() ()  {
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateHost"] = (*this).validateHost
	validation_functions["validateCredentials"] = (*this).validateCredentials
	validation_functions["validateDatabase"] = (*this).validateDatabase

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

func (this *Client) CreateDatabase(database_name *string, database_create_options *DatabaseCreateOptions, options map[string][]string) (*Database, *string, []error) {
	return NewDatabase((*this).GetHost(), (*this).GetCredentials(), database_name, database_create_options, options).Create()
}

func (this *Client) CreateUser(username *string, password *string, domain_name *string, options map[string][]string) (*User, *string, []error) {
	var errors []error 
	credentials := NewCredentials(username, password)
	domain := NewDomainName(domain_name)

	if (*this).GetHost() == nil {
		errors = append(errors, fmt.Errorf("holistic.Client: holistic.Host is nil, please set the host"))
	}

	if (*this).GetCredentials() == nil {
		errors = append(errors, fmt.Errorf("holistic.Client: holistic.Credentials is nil, please set the credentials"))
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	
	return newUser((this), credentials, domain, options).Create()
}

func (this *Client) GetHost() *Host {
	return (*this).host
 }

func (this *Client) GetCredentials() *Credentials {
	return (*this).credentials
 }

 func (this *Client) GetDatabase() *Database {
	return (*this).database
 }

 func (this *Client) validateValidationFunctions() ([]error) {
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

func (this *Client) getValidationFunctions() map[string]func() []error {
	return (*this).validation_functions
}

func (this *Client) Validate() []error {
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

func (this *Client) validateHost()  ([]error) {
	if (*this).GetHost() == nil {
		return nil
	}

	return (*((*this).GetHost())).Validate()
}

func (this *Client) validateCredentials()  ([]error) {
	if (*this).GetCredentials() == nil {
		return nil
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *Client) validateDatabase()  ([]error) {
	if (*this).GetDatabase() == nil {
		return nil
	}

	return (*((*this).GetDatabase())).Validate()
}
