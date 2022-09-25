package class

import (
	"fmt"
	//"reflect"
	//"fmt"
)

type Context struct {
	LogErrors func([]error)
	LogError func(error)
	HasErrors func() bool
}

func NewContext() (*Context) {
	errors := []error{}
	
	getErrors := func() []error {
		return errors
    }

	logError := func(err error) {
		errors := getErrors()
		errors = append(errors, err)
		fmt.Println(err)
    }
	
	logErrors := func(errs []error) {
		for i:=0; i < len((errs)); i++ {
			err := errs[i]
			logError(err)
		}
    }

	hasErrors := func() bool {
		return len(getErrors()) != 0
    }

	return &Context{
		LogErrors: func(errs []error) {
			logErrors(errs)
		},
		LogError: func(err error) {
			logError(err)
		},
		HasErrors: func() bool {
			return hasErrors()
		},
    }
}

/*
func (this *DatabaseCreateOptions) getValidations() Map {	
	typeOf := fmt.Sprintf("%T", *this)
	validations := Map{GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET():Array{Map{"function":ContainsExactMatchz,"parameters": Map{"whitelist|[]string":GET_CHARACTER_SETS(),"type|data_type":typeOf,"data|string":"utf8","column_name|string":GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET()}}}}
	

	return validations
}*/

/*
func (this *DatabaseCreateOptions) getData() map[string]interface{} {
	return (*this).data
}

func (this *DatabaseCreateOptions) GetCharacterSet() *string {
	return (*this).character_set
}

func (this *DatabaseCreateOptions) GetCollate() *string {
	return (*this).collate
}*/
/*
func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	var array_of_validations = (*this).getValidations()
	var keys = KeysForMap(array_of_validations)
	for _, parameter := range keys {
		var method_signiture = array_of_validations[parameter].(Array)
		
		for _, validation := range method_signiture {
	
			var vargsConvert = []reflect.Value{reflect.ValueOf(validation)}

		    var output_array_map_result = reflect.ValueOf(validation.(Map).Func("function")).Call(vargsConvert)

			validation_errors := ConvertPrimitiveReflectValueArrayToArray(output_array_map_result)
			outer_array_length := len(validation_errors)
			for i := 0; i < outer_array_length; i++ {
				validation_error := validation_errors[i]
				error_value := fmt.Sprintf("%s", reflect.ValueOf(validation_error).Interface())
				if error_value == "[]" {
					continue
				}
				errors = append(errors, fmt.Errorf(error_value))
			}

		}
	}

	return errors
}*/
