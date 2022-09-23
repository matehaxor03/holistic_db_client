package class

import (
	"fmt"
	"reflect"
	"strings"
	consts "github.com/matehaxor03/holistic_db_client/consts"
	common "github.com/matehaxor03/holistic_db_client/common"
)

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET() string {
	return "character_set"
}

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE() string {
	return "collate"
}

type DatabaseCreateOptions struct {
	data common.Map
	validations common.Map
	
	character_set *string
	collate *string

	validation_functions map[string]func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	/*data := Map{"a":"apple", 
	"b":2}*/

	data := common.Map {}
	
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = character_set
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] = collate

	//validations := Map {}
	//validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = Map{}
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = Array{Map{"function": ContainsExactMatchz, "parameters": Map{"whitelist": consts.GET_CHARACTER_SETS(), "kind":reflect.ValueOf(*this), "data": func () string {return strings.Clone(*character_set)}}}}


	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = []func(...interface{}) []error {Containsy}
	
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS()] = Array{consts.GET_CHARACTER_SETS()}

	//validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] =  (*this).validateCollate
	
	/*args := []string{"hello", "go"}
	vargs := make([]reflect.Value, len(args))
	
	for n, v := range args {
		vargs[n] = reflect.ValueOf(v)
	}*/

	//reflect.ValueOf(validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()).FA(FIELD_NAME_VALIDATION_FUNCTIONS())[0].(func(...interface{}) []error)).Call(vargs)

	x := DatabaseCreateOptions{data: data}
	return &x
}

func (this *DatabaseCreateOptions) getValidations() common.Map {	
	//data := (*this).getData()
	
	validations := common.Map{GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET(): common.Array{common.Map{"function": ContainsExactMatchz, "parameters": common.Map{"whitelist": consts.GET_CHARACTER_SETS(), "reflect.ValueOf":reflect.ValueOf(*this), "data":"gfgff"}}}}
	//validations.M(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET())[FIELD_NAME_VALIDATION_FUNCTIONS()] = []func(...interface{}) []error {Containsy}
	/*var whiltlistCustome = Array{}
	for _, value := range consts.GET_CHARACTER_SETS() {
		whiltlistCustome = append(whiltlistCustome, value)
	}*/

	//validations[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = Array{Map{"function": ContainsExactMatchz, "parameters": Map{"whitelist": consts.GET_CHARACTER_SETS(), "kind":reflect.ValueOf(*this), "data": func () string {return strings.Clone(data.S(GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()))}}}}

	return validations
}

func (this *DatabaseCreateOptions) getData() map[string]interface{} {
	return (*this).data
}

/*
func (this *DatabaseCreateOptions) validateCharacterSet() ([]error) {
	if (*this).character_set == nil {
		return nil
	}

	return ContainsExactMatch(ConvertIntefaceArrayToStringArray(consts.GET_CHARACTER_SETS()), (*this).character_set, "character_set", reflect.ValueOf(*this))
}

func (this *DatabaseCreateOptions) validateCollate() ([]error) {
	if (*this).collate == nil {
		return nil
	}

	return ContainsExactMatch(consts.GET_COLLATES(), (*this).collate, "collate", reflect.ValueOf(*this))
}
*/

func (this *DatabaseCreateOptions) GetCharacterSet() *string {
	return (*this).character_set
}

func (this *DatabaseCreateOptions) GetCollate() *string {
	return (*this).collate
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	var mappy = (*this).getValidations()
	var keys = common.KeysForMap(mappy)
	for _, parameter := range keys {
		var array_of_validations = mappy[parameter].(common.Array)
		
		for _, validation := range array_of_validations {
			fmt.Println(validation)
			var function = validation.(common.Map).Func("function")
			var parameters = validation.(common.Map).M("parameters")
			var keys_of_parameters = parameters.Keys()
			keys_of_parameters = append(keys_of_parameters, "column_name")

			var vargs = make(map[string]interface{})
			
			for _, v := range keys_of_parameters {
				vargs[v] = reflect.ValueOf(parameters[v])
			}
			vargs["column_name"] = parameter

			var root = make(map[string]interface{})
			root["function"] = function
			root["parameters"] = vargs
			
			var vargsConvert = []reflect.Value{reflect.ValueOf(root)}

		   var output_array_map_result = reflect.ValueOf(function).Call(vargsConvert)
		   
		   var singleResult = make([]interface{}, len(output_array_map_result))
		   var counttit = 0
		   singleResult[counttit] = output_array_map_result
		   
		   fmt.Println("##############################")
		   //fmt.Println(output_array_map_result)
		   for key, value := range singleResult {
			fmt.Println(fmt.Sprintf("key: %s value: %s", key, value))
			mappoutput := value
			for key2, value2 := range mappoutput.([]reflect.Value) {
				fmt.Println(fmt.Sprintf("%s %s", key2, value2))
				if strings.Contains(fmt.Sprintf("%s", key2), "POTENTIAL SQL INJECTION:") || 
				   strings.Contains(fmt.Sprintf("%s", value2), "POTENTIAL SQL INJECTION:") {
					errors = append(errors, fmt.Errorf("WWWWWWWWWWWW%s", key2))
					errors = append(errors, fmt.Errorf("WWWWWWWWWWWW%s", value2))
				}
				
				
				
				/*var valueOfit = reflect.ValueOf(value2).(map[string]interface{})


				errorrssdfd := valueOfit["errors"]
				if errorrssdfd != nil && len(errorrssdfd) > 0 {
					errors = append(errors, errorrssdfd...)
				}
				panic(errorrssdfd)*/
			}
		   }
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


/*
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
}*/

