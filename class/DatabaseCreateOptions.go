package class

import (
	"reflect"
	"fmt"
	//"unsafe"
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
	data := Map {}
	
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()] = character_set
	data[GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE()] = collate

	x := DatabaseCreateOptions{data: data}
	return &x
}

func (this *DatabaseCreateOptions) getValidations() Map {	
	
	validations := Map{GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET(): Array{Map{"function": ContainsExactMatchz, "parameters": Map{"whitelist": GET_CHARACTER_SETS(), "reflect.ValueOf":reflect.ValueOf(*this), "data":"utjjf8"}}}}
	

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
	var keys = KeysForMap(array_of_validations)
	for _, parameter := range keys {
		var method_signiture = array_of_validations[parameter].(Array)
		
		for _, validation := range method_signiture {
			var result_obj = NewResult()
			var function = validation.(Map).Func("function")
			var parameters = validation.(Map).M("parameters")
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
			root["result"] = result_obj

			
			var vargsConvert = []reflect.Value{reflect.ValueOf(root)}

		    var output_array_map_result = reflect.ValueOf(function).Call(vargsConvert)

			Examiner(reflect.TypeOf(output_array_map_result), 0)
			
			//var functionReturnResult = output_array_map_result[0].Interface()
			return_result_of_func := ConvertPrimitiveValueToMap(output_array_map_result[0])
			fmt.Println(fmt.Sprintf("%s", return_result_of_func))
			return_result_of_map_func := ConvertPrimitiveValueToMap(return_result_of_func["value"])
			fmt.Println(fmt.Sprintf("%s", return_result_of_map_func))
			return_result_of_map_func_value := reflect.ValueOf(return_result_of_map_func["value"])
			fmt.Println(fmt.Sprintf("%s", return_result_of_map_func_value))

			//return_result_of_obj_value := return_result_of_map_func_value.Interface().(Result)
			//Examiner(reflect.TypeOf(return_result_of_obj_value), 5)
			fmt.Println(fmt.Sprintf("%s", return_result_of_map_func_value))
			
			//Examiner(reflect.TypeOf(return_result_of_map_func_value), 5)

			/*
			reflect_map_ptr := return_result_of_map_func_value.FieldByName("ptr")
			fmt.Println(fmt.Sprintf("%s", reflect_map_ptr))

			reflect_map := 	*(*(map[string]interface{}))(unsafe.Pointer(&(reflect.ValueOf(return_result_of_func))))
			fmt.Println(fmt.Sprintf("%s", reflect_map))*/


			//mapArray := (*[n]float32)(unsafe.Pointer(&mmap[0]))

			//return_result_of_map_func_value_map := reflect.MapOf(reflect.TypeOf("map[string]interface{}"), reflect.TypeOf(&return_result_of_map_func_value).Elem())

			
			//fmt.Println(fmt.Sprintf("%s", reflect.ValueOf(return_result_of_map_func["value"]).Interface()))
			/*
			errorsOfFunc := common.ConvertPrimitiveMapToMap(functionReturnResult)["errors"].([]error)
			if errorsOfFunc != nil && len(errorsOfFunc) > 0 {
				errors = append(errors, errorsOfFunc)
			}*/
			

			/*
			err := output_array_map_result[0].Interface()
			if err == nil {
				fmt.Println("No error returned by", function)
			} else {
				fmt.Printf("Error calling %s: %v", function, err)
			}*/
			
			//result := common.ConvertPrimativeArrayToArray(reflect.ValueOf(output_array_map_result).Interface().([]reflect.Value))
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[0])))
			//errors = append(errors, result[1]["errors"]...)
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[1])))
			//panic(fmt.Sprintf("%s",reflect.ValueOf(result[2])))
		   
		}
	}

	return errors
}
