package class

import (
	//"reflect"
	//"fmt"
)

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET() string {
	return "character_set"
}

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE() string {
	return "collate"
}

type DatabaseCreateOptions struct {
	GetData func() Map
	GetCharacterSet func() (*string, error)
	GetCollate func() (*string, error)
	Validate func() ([]error)
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	data := Map {"character_set":Map{"type|string":"string","value|string":character_set,"constraints|string":"optional:,white_list:GET_CHARACTER_SETS"},
	"collate":Map{"type|string":"string","value|string":collate,"constraints|string":"optional:,white_list:GET_COLLATES"}}
	
	getData := func() Map {
		return data
    }

	getCharacterSet := func() (*string, error) {
		return data.M("character_set").S("value|string")
	}

	getCollate := func() (*string, error) {
		return data.M("collate").S("value|string")
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "DatabaseCreateOptions")
	}
	
	return &DatabaseCreateOptions{
        GetData: func() Map {
            return getData()
        },
		GetCharacterSet: func() (*string, error) {
			return getCharacterSet()
		},
		GetCollate: func() (*string, error) {
			return getCollate()
		},
		Validate: func() ([]error) {
			return validate()
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
