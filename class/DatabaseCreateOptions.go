package class

import (
	"fmt"
	"reflect"
	"strings"
	consts "github.com/matehaxor03/holistic_db_client/consts"
)

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET() string {
	return "character_set"
}

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE() string {
	return "collate"
}

type DatabaseCreateOptions struct {
	data Map
	validations Map
	
	character_set *string
	collate *string

	validation_functions map[string]func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	/*data := Map{"a":"apple", 
	"b":2}*/

	data := Map {}
	
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = character_set
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] = collate
	data["type"] = reflect.TypeOf(DatabaseCreateOptions{})

	validations := Map {}
	validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = Map{}
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = []func(...interface{}) []error {Containsy}
	validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = Array{Map{"function": Containsy, "parameters": Map{"whitelist": consts.GET_CHARACTER_SETS(), "data": func () string {return strings.Clone(*character_set)}}}}
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS()] = Array{consts.GET_CHARACTER_SETS()}

	//validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] =  (*this).validateCollate
	
	/*args := []string{"hello", "go"}
	vargs := make([]reflect.Value, len(args))
	
	for n, v := range args {
		vargs[n] = reflect.ValueOf(v)
	}*/


	//reflect.ValueOf(validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()).FA(FIELD_NAME_VALIDATION_FUNCTIONS())[0].(func(...interface{}) []error)).Call(vargs)

	x := DatabaseCreateOptions{data: data, validations: validations}


	

	//Containsy("hi")
	
	x.validation_functions = make(map[string]func() []error)
	x.InitValidationFunctions()
	return &x
}

func (this *DatabaseCreateOptions) getValidations(data Map) map[string]interface{} {	
	validations := make(map[string]interface{})
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = []func(...interface{}) []error {Containsy}
	validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = Array{Map{"function": Containsy, "parameters": Map{"whitelist": consts.GET_CHARACTER_SETS(), "data": func () string {return strings.Clone(data.S(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()))}}}}

	return validations
}

func (this *DatabaseCreateOptions) InitValidationFunctions() ()  {
	
	
	
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateCharacterSet"] = (*this).validateCharacterSet
	validation_functions["validateCollate"] = (*this).validateCollate
	validation_functions["validateValidationFunctions"] = (*this).validateValidationFunctions

	if validation_functions["validateValidationFunctions"] == nil|| 
	   GetFunctionName(validation_functions["validateValidationFunctions"]) != GetFunctionName((*this).validateValidationFunctions) {
		panic(fmt.Errorf("validateValidationFunctions validation method not found potential sql injection without it"))
	}
}

func (this *DatabaseCreateOptions) getData() map[string]interface{} {
	return (*this).data
}

func (this *DatabaseCreateOptions) validateCharacterSet() ([]error) {
	if (*this).character_set == nil {
		return nil
	}

	return Contains(consts.GET_CHARACTER_SETS(), (*this).character_set, "character_set")
}

func (this *DatabaseCreateOptions) validateCollate() ([]error) {
	if (*this).collate == nil {
		return nil
	}

	return Contains(consts.GET_COLLATES(), (*this).collate, "collate")
}

func (this *DatabaseCreateOptions) GetCharacterSet() *string {
	return (*this).character_set
}

func (this *DatabaseCreateOptions) GetCollate() *string {
	return (*this).collate
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	var mappy = (*this).getValidations((*this).data)
	var keys = KeysForMap(mappy)
	for _, parameter := range keys {
		var array_of_validations = mappy[parameter].(Array)
		
		for _, validation := range array_of_validations {
			
			//	//reflect.ValueOf(validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()).FA(FIELD_NAME_VALIDATION_FUNCTIONS())[0].(func(...interface{}) []error)).Call(vargs)
			var function = validation.(Map).Func("function")
			var parameters = validation.(Map).M("parameters")
			var keys_of_parameters = parameters.Keys()
			keys_of_parameters = append(keys_of_parameters, "_column_name")

			var vargs = make(map[string]interface{})
			
			for _, v := range keys_of_parameters {
				vargs[v] = reflect.ValueOf(parameters[v])
			}
			vargs["_column_name"] = parameter
			

			var vargsConvert = []reflect.Value{reflect.ValueOf(vargs)}

		    errors := reflect.ValueOf(function).Call(vargsConvert)
			panic("hi")
			panic(errors)
		}	
	}

	return errors
}

/*
func (this *DatabaseCreateOptions) Validate() []error {
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
}*/

func (this *DatabaseCreateOptions) getValidationFunctions() map[string]func() []error {
	return (*this).validation_functions
}

func (this *DatabaseCreateOptions) validateValidationFunctions() ([]error) {
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

