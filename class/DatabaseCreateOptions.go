package class

import (
	"reflect"
	"fmt"
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
	data := common.Map {}
	
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = character_set
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] = collate

	x := DatabaseCreateOptions{data: data}
	return &x
}

func (this *DatabaseCreateOptions) getValidations() common.Map {	
	
	validations := common.Map{GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET(): common.Array{common.Map{"function": ContainsExactMatchz, "parameters": common.Map{"whitelist": consts.GET_CHARACTER_SETS(), "reflect.ValueOf":reflect.ValueOf(*this), "data":"utf8"}}}}
	

	return validations
}

func (this *DatabaseCreateOptions) getData() map[string]interface{} {
	return (*this).data
}

func (this *DatabaseCreateOptions) GetCharacterSet() *string {
	return (*this).character_set
}

func (this *DatabaseCreateOptions) GetCollate() *string {
	return (*this).collate
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	var array_of_validations = (*this).getValidations()
	var keys = common.KeysForMap(array_of_validations)
	for _, parameter := range keys {
		var method_signiture = array_of_validations[parameter].(common.Array)
		
		for _, validation := range method_signiture {
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

			err := output_array_map_result[0].Interface()
			if err == nil {
				fmt.Println("No error returned by", function)
			} else {
				fmt.Printf("Error calling %s: %v", function, err)
			}
			
			//result := common.ConvertPrimativeArrayToArray(reflect.ValueOf(output_array_map_result).Interface().([]reflect.Value))
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[0])))
			//errors = append(errors, result[1]["errors"]...)
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[1])))
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[2])))
		   
		}
	}

	return errors
}
