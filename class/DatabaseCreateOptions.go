package class

import (
	"fmt"
	"strings"
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
	Validate func() ([]error)
	GetSQL func() (*string, []error)
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	data := Map {
		"character_set":Map{"type|string":"string","value|string":character_set, 
		FILTERS(): Array{ Map {"values|array":GET_CHARACTER_SETS(),"function|func":ContainsExactMatch } }},
		"collate":Map{"type|string":"string","value|string":collate,
		FILTERS(): Array{ Map {"values|array":GET_COLLATES(),"function|func":ContainsExactMatch } }},
	}

	//panic(data.ToJSONString())
	
	getData := func() Map {
		return data
    }

	getCharacterSet := func() (*string) {
		v := data.M("character_set").S("value|string")
		if v == nil {
			return nil
		}
		clone := strings.Clone(*v)
		return &clone
	}

	getCollate := func() (*string) {
		v := data.M("collate").S("value|string")
		if v == nil {
			return nil
		}
		clone := strings.Clone(*v)
		return &clone
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(getData(), "DatabaseCreateOptions")
	}

	getSQL := func() (*string, []error) {
		errs := validate()
		if errs != nil {
			return nil, errs
		}
		
		sql_command := ""

		character_set := getCharacterSet()
		if character_set != nil {
			sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
		}
		
		collate := getCollate()
		if collate != nil {
			sql_command += fmt.Sprintf("COLLATE %s ", *collate)
		}

		return &sql_command, nil
	}
	
	return &DatabaseCreateOptions{
		Validate: func() ([]error) {
			return validate()
		},
		GetSQL: func() (*string, []error) {
			return getSQL()
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
