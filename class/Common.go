package class

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
    "path/filepath"
	"runtime"
	//"reflect"
)

func GetValidSchemaFields() Map {
	return Map{
		"type": nil,
		"primary_key": nil,
		"unsigned":nil,
		"auto_increment": nil,
		"default": nil,
		"validated": nil,
		"filters": nil,
		"not_empty_string_value":nil,
		"min_length": nil,
		"max_length": nil,
	}
}

func EscapeString(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "'", "\\'")
	return value
}

func CloneString(value *string) *string {
	if value == nil {
		return nil
	}

	temp := strings.Clone(*value)
	return &temp
}

func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}
	
	string_value := fmt.Sprintf("%s", object) 

	if string_value == "<nil>" {
		return true
	}

	if string_value == "map[]" {
		return true
	}

	rep := fmt.Sprintf("%T", object)

	if string_value == "%!s("+rep+"=<nil>)" {
		return true
	}

	/*
	switch reflect.TypeOf(object).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(object).IsNil()
	}*/

	return false
	/*
	string_value := fmt.Sprintf("%s", object)

	if string_value == "<nil>" {
		return true
	}

	rep := fmt.Sprintf("%T", object)

	if string_value == "%!s("+rep+"=<nil>)" {
		return true
	}

	return false*/
}

func IsTime(object interface{}) bool {
	if IsNil(object) {
		return false
	}

	time, time_errors := GetTime(object)
	if time_errors != nil {
		return false
	} else if IsNil(time) {
		return false
	}

	return true
}

func GetTime(object interface{}) (*time.Time, []error) {
	var errors []error
	var result *time.Time

	if object == nil {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", object)
	switch rep {
	case "*time.Time":
		value := *(object.(*time.Time))
		result = &value
	case "time.Time":
		value := object.(time.Time)
		result = &value
	case "*string":
		//todo: parse for null
		value1, value_errors1 := time.Parse("2006-01-02 15:04:05.000000", *(object.(*string)))
		value2, value_errors2 := time.Parse("2006-01-02 15:04:05", *(object.(*string)))
		var value3 *time.Time
		if *(object.(*string)) == "now" {
			value3 = GetTimeNow()
		} else {
			value3 = nil
		}

		if value_errors1 != nil && value_errors2 != nil && value3 == nil {
			errors = append(errors, value_errors1)
		}

		if value_errors1 == nil {
			result = &value1
		}

		if value_errors2 == nil {
			result = &value2
		}

		if value3 != nil {
			result = value3
		}

	case "string":
		//todo: parse for null
		value1, value_errors1 := time.Parse("2006-01-02 15:04:05.000000", (object.(string)))
		value2, value_errors2 := time.Parse("2006-01-02 15:04:05", (object.(string)))
		var value3 *time.Time
		if (object.(string)) == "now" {
			value3 = GetTimeNow()
		} else {
			value3 = nil
		}

		if value_errors1 != nil && value_errors2 != nil && value3 == nil {
			errors = append(errors, value_errors1)
		}

		if value_errors1 == nil {
			result = &value1
		}

		if value_errors2 == nil {
			result = &value2
		}

		if value3 != nil {
			result = value3
		}

	default:
		errors = append(errors, fmt.Errorf("Map.GetTime: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func GetType(object interface{}) string {
	if IsNil(object) {
		return "nil"
	}
	return fmt.Sprintf("%T", object)
}

func FIELD_NAME_VALIDATION_FUNCTIONS() string {
	return "validation_functions"
}

func FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS() string {
	return "validation_functions_parameters"
}

func getWhitelistStringFunc() *func(m Map) []error {
	function := WhiteListString
	return &function
}

func getBlacklistStringFunc() *func(m Map) []error {
	function := BlackListString
	return &function
}

func getBlacklistStringToUpperFunc() *func(m Map) []error {
	function := BlackListStringToUpper
	return &function
}

func WhiteListString(m Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if  map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	_, found := (*map_values)[*str]

	if !found {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: did not find value", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func BlackListString(m Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	_, found := (*map_values)[*str]

	if found {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}


func GetFields(struct_type string, m *Map, field_type string) (*Map, []error) {
	var errors []error
	if !(field_type == "[fields]" || field_type == "[system_fields]") {
		available_fields := m.Keys()
		errors = append(errors, fmt.Errorf("%s %s is not a valid root field, available root fields: %s", struct_type, field_type, available_fields))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	fields_map, fields_map_errors := m.GetMap(field_type)
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} else if IsNil(fields_map) {
		errors = append(errors, fmt.Errorf("%s %s is nil", struct_type, field_type))
	} 

	if len(errors) > 0 {
		return nil, errors
	}
	
	return fields_map, nil
}

func GetSchemas(struct_type string, m *Map, schema_type string) (*Map, []error) {
	var errors []error
	if !(schema_type == "[schema]" || schema_type == "[system_schema]") {
		available_fields := m.Keys()
		errors = append(errors, fmt.Errorf("%s, %s is not a valid root system schema, available root fields: %s", struct_type, schema_type, available_fields))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	schemas_map, schemas_map_errors := m.GetMap(schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if IsNil(schemas_map) {
		errors = append(errors, fmt.Errorf("%s %s is nil", struct_type, schema_type))
	} else {
		schema_paramters := schemas_map.Keys()
		for _, schema_paramter := range schema_paramters {
			if !schemas_map.IsMap(schema_paramter) {
				errors = append(errors, fmt.Errorf("%s %s %s is not a map", struct_type, schema_type, schema_paramter))
			} else {
				schema_paramter_map, schema_paramter_map_errors := schemas_map.GetMap(schema_paramter) 
				if schema_paramter_map_errors != nil {
					errors = append(errors, fmt.Errorf("%s %s %s had errors getting map: %s", struct_type, schema_type, schema_paramter, fmt.Sprintf("%s",schema_paramter_map_errors))) 
				} else {
					attributes := schema_paramter_map.Keys()
					valid_attributes_map := GetValidSchemaFields()
					for _, attribute := range attributes {
						if !valid_attributes_map.HasKey(attribute) {
							errors = append(errors, fmt.Errorf("%s %s %s has an invalid attribute: %s valid attributes are: %s", struct_type, schema_type, schema_paramter, attribute, valid_attributes_map.Keys()))
						}
					}
				}
			}
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return schemas_map, nil
}

func GetField(struct_type string, m *Map, schema_type string, field_type string, field string, desired_type string) (interface{}, []error) {
	var errors []error
	schemas_map, schemas_map_errors := GetSchemas(struct_type, m, schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if !schemas_map.HasKey(field) {
		available_fields := schemas_map.Keys()
		errors = append(errors, fmt.Errorf("Common.GetField %s schema_type: %s field: %s does not exist available fields are: %s", struct_type, schema_type, field, fmt.Sprintf("%s", available_fields)))
	} else if !schemas_map.IsMap(field) {
		errors = append(errors, fmt.Errorf("Common.GetField %s schema_type: %s field: %s is not a map and has type: %s", struct_type, schema_type, field, schemas_map.GetType(field)))
	} 

	if len(errors) > 0 { 
		return nil, errors
	}

	fields_map, fields_map_errors := GetFields(struct_type, m, field_type)
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} 

	if len(errors) > 0 {
		return nil, errors
	}

	schema_map, schema_map_errors := schemas_map.GetMap(field)
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	} else if schema_map == nil {
		errors = append(errors, fmt.Errorf("%s %s map is nil", struct_type, schema_type))
	} else if !schema_map.HasKey("type") {
		available_fields := schemas_map.Keys()
		errors = append(errors, fmt.Errorf("%s field: %s schema \"type\" attribute does not exist available fields are: %s", struct_type, field, fmt.Sprintf("%s", available_fields)))
	} else if !schema_map.IsString("type") {
		errors = append(errors, fmt.Errorf("%s field: %s schema \"type\" attribute value is not a string it's %s", struct_type, field, schema_map.GetType("type")))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	schema_type_value, schema_type_value_errors := schema_map.GetString("type")
	if schema_type_value_errors != nil {
		errors = append(errors, schema_type_value_errors...)
	} else if schema_type_value == nil {
		errors = append(errors, fmt.Errorf("field: %s schema type is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var result interface{}
	if fields_map.IsNil(field) {
		if schema_map.HasKey("default") && !schema_map.IsNil("default") {
			result = schema_map.GetObject("default")
		} else {
			result = nil
		}
	} else {
		result = fields_map.GetObject(field)
	}


	if IsNil(result) {
		if schema_map.IsBoolTrue("auto_increment") && schema_map.IsBoolTrue("primary_key") {
			return nil, nil
		}

		if schema_map.IsBoolTrue("primary_key") {
			errors = append(errors,	fmt.Errorf("field: %s had nil value and default value but is a primary key"))
			return nil, errors
		}

		if strings.HasPrefix(*schema_type_value, "*") {
			return nil, nil
		}

		errors = append(errors,	fmt.Errorf("field: %s had nil value and default value but is not nullable"))
		return nil, errors
	}
	

	object_type := fields_map.GetType(field)
	if strings.ReplaceAll(object_type, "*", "") != strings.ReplaceAll(*schema_type_value, "*", "") {
		errors = append(errors, fmt.Errorf("field: %s schema type: %s actual: %s are not a match", field, strings.ReplaceAll(*schema_type_value, "*", ""), strings.ReplaceAll(object_type, "*", "")))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if desired_type == "self" {
		return result, nil
	}

	type_of := GetType(result)

	if desired_type == type_of {
		return result, nil
	}

	switch type_of {
	case "string":
		switch desired_type {
		case "*string":
			temp_value := result.(string)
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*string":
		switch desired_type {
		case "string":
			return *(result.(*string)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*uint64":
		switch desired_type {
		case "uint64":
			return *(result.(*uint64)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*time.Time":
		switch desired_type {
		case "time.Time":
			return *(result.(*time.Time)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.ClientManager":
		switch desired_type {
		case "class.ClientManager":
			return *(result.(*ClientManager)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Client":
		switch desired_type {
		case "class.Client":
			return *(result.(*Client)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Host":
		switch desired_type {
		case "class.Host":
			return *(result.(*Host)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Database":
		switch desired_type {
		case "class.Database":
			return *(result.(*Database)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Table":
		switch desired_type {
		case "class.Table":
			return *(result.(*Table)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Record":
		switch desired_type {
		case "class.Record":
			return *(result.(*Record)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.DomainName":
		switch desired_type {
		case "class.DomainName":
			return *(result.(*DomainName)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Grant":
		switch desired_type {
		case "class.Grant":
			return *(result.(*Grant)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.Credentials":
		switch desired_type {
		case "class.Credentials":
			return *(result.(*Credentials)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.User":
		switch desired_type {
		case "class.User":
			return *(result.(*User)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*class.DatabaseCreateOptions":
		switch desired_type {
		case "class.DatabaseCreateOptions":
			return *(result.(*DatabaseCreateOptions)), nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.ClientManager":
		switch desired_type {
		case "*class.ClientManager":
			temp_result := result.(ClientManager)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Client":
		switch desired_type {
		case "*class.Client":
			temp_result := result.(Client)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Host":
		switch desired_type {
		case "*class.Host":
			temp_result := result.(Host)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Database":
		switch desired_type {
		case "*class.Database":
			temp_result := result.(Database)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Table":
		switch desired_type {
		case "*class.Table":
			temp_result := result.(Table)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Record":
		switch desired_type {
		case "*class.Record":
			temp_result := result.(Record)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.DomainName":
		switch desired_type {
		case "*class.DomainName":
			temp_result := result.(DomainName)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Grant":
		switch desired_type {
		case "*class.Grant":
			temp_result := result.(Grant)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.Credentials":
		switch desired_type {
		case "*class.Credentials":
			temp_result := result.(Credentials)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.User":
		switch desired_type {
		case "*class.User":
			temp_result := result.(User)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "class.DatabaseCreateOptions":
		switch desired_type {
		case "*class.DatabaseCreateOptions":
			temp_result := result.(DatabaseCreateOptions)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
	}
	default:
		errors = append(errors, fmt.Errorf("Common.GetField mapping not supported please implement: %s", type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return fields_map.GetObject(field), nil
}

func SetField(struct_type string, m *Map, schema_type string, field_type string, parameter string, object interface{}) ([]error) {
	var errors []error

	schemas_map, schemas_map_errors := GetSchemas(struct_type, m, schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	}

	fields_map, fields_map_errors := GetFields(struct_type, m, field_type) 
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	}

	if len(errors) > 0 {
		return errors
	}
	
	schema_of_parameter_map, schema_of_parameter_map_errors := schemas_map.GetMap(parameter)
	if schema_of_parameter_map_errors != nil {
		errors = append(errors, schema_of_parameter_map_errors...)
	} else if schema_of_parameter_map == nil {
		errors = append(errors, fmt.Errorf("field: %s schema map is nil", parameter))
	} else if !schema_of_parameter_map.HasKey("type") {
		available_fields := schema_of_parameter_map.Keys()
		errors = append(errors, fmt.Errorf("field: %s schema \"type\" attribute does not exist available fields are: %s", parameter, fmt.Sprintf("%s", available_fields)))
	} else if !schema_of_parameter_map.IsString("type") {
		errors = append(errors, fmt.Errorf("field: %s schema \"type\" attribute value is not a string it's %s", parameter, schema_of_parameter_map.GetType("type")))
	}

	if len(errors) > 0 {
		return errors
	}

	var primary_key_count *int
	primary_key_count_value := 0 
	primary_key_count = &primary_key_count_value
	
	var auto_increment_count *int
	auto_increment_count_value := 0 
	auto_increment_count = &auto_increment_count_value

	validate_parameters_errors := ValidateParameterData(struct_type, schemas_map, nil, parameter, object, primary_key_count, auto_increment_count)
	if validate_parameters_errors != nil {
		errors = append(errors, validate_parameters_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	fields_map.SetObject(parameter, object)
	validated_true := true
	schema_of_parameter_map.SetObject("validated", validated_true)
	return nil
}

func BlackListStringToUpper(m Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	_, found := (*map_values)[strings.ToUpper(*str)]

	if found {
		errors = append(errors, fmt.Errorf("%s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func getWhitelistCharactersFunc() *func(m Map) []error {
	funcValue := WhitelistCharacters
	return &funcValue
}

func WhitelistCharacters(m Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	var invalid_letters []string
	for _, letter_rune := range *str {
		letter_string := string(letter_rune)
		_, found := (*map_values)[letter_string]

		if !found {
			invalid_letters = append(invalid_letters, letter_string)
		}
	}

	if len(invalid_letters) > 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has invalid character(s): %s", *data_type, *label, invalid_letters))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func IsDatabaseColumn(value string) bool {
	column_name_params := Map{"values": GetMySQLColumnNameWhitelistCharacters(), "value": value, "label": "column_name", "data_type": "Table"}
	column_name_errors := WhitelistCharacters(column_name_params)
	if column_name_errors != nil {
		return false
	}

	blacklist_column_name_params := Map{"values": GetMySQLKeywordsAndReservedWordsInvalidWords(), "value": value, "label": "column_name", "data_type": "Table"}
	blacklist_column_name_errors := BlackListStringToUpper(blacklist_column_name_params)
	if blacklist_column_name_errors != nil {
		return false
	}

	return true
}

func ValidateData(data *Map, struct_type string) []error {	
	var errors []error
	var ignore_identity_errors = false
	
	var primary_key_count *int
	primary_key_count_value := 0 
	primary_key_count = &primary_key_count_value
	
	var auto_increment_count *int
	auto_increment_count_value := 0 
	auto_increment_count = &auto_increment_count_value

	if len(errors) > 0 {
		return errors
	}

	if (struct_type == "*class.Table" || struct_type  == "class.Table") && data.HasKey("[schema_is_nil]") {
		if data.IsBoolTrue("[schema_is_nil]") {
			ignore_identity_errors = true
		}
	}

	var field_errors []error
	field_parameters, field_parameters_errors := GetFields(struct_type, data, "[fields]")
	if field_parameters_errors != nil {
		field_errors = append(field_errors, field_parameters_errors...)
	}

	schemas, schemas_errors := GetSchemas(struct_type, data, "[schema]")
	if schemas_errors != nil {
		field_errors = append(field_errors, schemas_errors...)
	}

	if len(field_errors) == 0 {
		for _, parameter := range (*schemas).Keys() {
			value_errors := ValidateParameterData(struct_type, schemas, field_parameters, parameter, nil, primary_key_count, auto_increment_count)

			if value_errors != nil {
				field_errors = append(field_errors, value_errors...)
			}
		}
	}

	if len(field_errors) > 0 {
		errors = append(errors, field_errors...)
	}

	var system_field_errors []error
	system_field_parameters, system_field_parameters_errors := GetFields(struct_type, data, "[system_fields]")
	if system_field_parameters_errors != nil {
		system_field_errors = append(system_field_errors, system_field_parameters_errors...)
	}

	system_schemas, system_schemas_errors := GetSchemas(struct_type, data, "[system_schema]")
	if system_schemas_errors != nil {
		system_field_errors = append(system_field_errors, system_schemas_errors...)
	}
	
	if len(system_field_errors) == 0 {
		for _, parameter := range (*system_schemas).Keys() {
			value_errors := ValidateParameterData(struct_type, system_schemas, system_field_parameters, parameter, nil, primary_key_count, auto_increment_count)
			if value_errors != nil {
				system_field_errors = append(system_field_errors, value_errors...)
			}
		}
	}

	if len(system_field_errors) > 0 {
		errors = append(errors, system_field_errors...)
	}

	if (struct_type == "*class.Table" || struct_type == "class.Table") && !ignore_identity_errors {
		if *primary_key_count <= 0 {
			errors = append(errors, fmt.Errorf("table: %s did not have any primary keys", struct_type))
		}

		if *auto_increment_count > 1 {
			errors = append(errors, fmt.Errorf("table: %s had more than one auto_increment attribute on a column", struct_type))
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateParameterData(struct_type string, schemas *Map, parameters *Map, parameter string, value_to_validate interface{}, primary_key_count *int,  auto_increment_count *int) ([]error) {
	var errors []error

	schema_of_parameter, schema_of_parameter_errors := schemas.GetMap(parameter)
	if schema_of_parameter_errors != nil {
		errors = append(errors, fmt.Errorf("Common.ValidateParameterData: %s column: %s error getting parameter schema %s", struct_type, parameter, fmt.Sprintf("%s", schema_of_parameter_errors)))
	} else if IsNil(schema_of_parameter) {
		errors = append(errors, fmt.Errorf("Common.ValidateParameterData: %s column: %s had nil schema", struct_type, parameter))
	} else if !schemas.IsMap(parameter) {
		errors = append(errors, fmt.Errorf("Common.ValidateParameterData: %s column: %s is not a map", struct_type, parameter))
	}

	if len(errors) > 0 {
		return errors
	} 

	//fmt.Println(fmt.Sprintf("%s %s %s %s %s", struct_type, schemas, parameters, parameter, value_to_validate))

	var value_is_mandatory bool
	var value_is_set bool
	var value_is_null bool

	value_is_mandatory = true
	value_is_set = true
	value_is_null = false

	if !IsNil(parameters) {
		if parameters.HasKey(parameter) {
			value_is_set = true
			if parameters.IsNil(parameter) {
				value_is_null = true
			} else {
				value_to_validate = parameters.GetObject(parameter)
				value_is_null = false
			}
		} else {
			value_is_set = false
			value_is_null = true
		}
	} else {
		if IsNil(value_to_validate) {
			value_is_null = true
		} else {
			value_is_null = false
		}
	}

	var default_set bool
	var default_is_null bool

	if schema_of_parameter.HasKey("default") {
		default_set = true
		if schema_of_parameter.IsNil("default") {
			default_is_null = true
		} else {
			default_is_null = false
		}
	} else {
		default_is_null = true
		default_set = false
	}

	type_of_parameter_schema_value, type_of_parameter_schema_value_errors := schema_of_parameter.GetString("type")
	if type_of_parameter_schema_value_errors != nil {
		errors = append(errors, fmt.Errorf("struct: %s column: %s error getting \"type\" attribute for schema %s", struct_type, parameter, fmt.Sprintf("%s", type_of_parameter_schema_value_errors)))
	} else if type_of_parameter_schema_value == nil {
		errors = append(errors, fmt.Errorf("struct: %s column: %s \"type\" attribute of schema is nil", struct_type, parameter))
	} else {
		if strings.HasPrefix(*type_of_parameter_schema_value, "*") {
			value_is_mandatory = false
		} else {
			value_is_mandatory = true
		}
	}
 
	if struct_type == "*class.Table" || struct_type == "class.Table" || struct_type == "*class.Record" || struct_type == "class.Record" {
		if schema_of_parameter.IsBoolTrue("primary_key") {
			value_is_mandatory = true
			*primary_key_count++ 

			if schema_of_parameter.IsBoolTrue("auto_increment") {
				value_is_mandatory = false
				*auto_increment_count++
			}
		}
	}

	if !IsNil(parameters) {
		if !schema_of_parameter.HasKey("validated") {
			bool_true := true
			schema_of_parameter.SetBool("validated", &bool_true)
		} else {
				if !schema_of_parameter.IsBool("validated") {
				errors = append(errors, fmt.Errorf("table: %s column: %s does not have attribute: validated is not a bool", struct_type, parameter))
				return errors
			} else if schema_of_parameter.IsBoolTrue("validated") {
				return nil
			} else {
				bool_true := true
				schema_of_parameter.SetBool("validated", &bool_true)
			}
		}
	}

	if (struct_type == "*class.Table" || struct_type == "class.Table") && IsDatabaseColumn(parameter) {
		value_is_mandatory = false
	}

	if len(errors) > 0 {
		return errors
	} 

	if value_is_null && default_is_null && !value_is_mandatory {
		return nil
	}


	if value_is_set && !value_is_null {

	} else if value_is_set && value_is_null {
		if default_set && default_is_null {
			value_to_validate = nil
		} else if default_set && !default_is_null {
			value_to_validate = schema_of_parameter.GetObject("default")
		} else if !default_set { 
			if value_is_mandatory {
				errors = append(errors,  fmt.Errorf("struct: %s column: %s does not have a value or a default value, value_set=%t value_nil=%t default_set=%t default_nil=%t", struct_type, parameter, value_is_set, value_is_null, default_set, default_is_null))
			} else {
				value_to_validate = nil
			}
		}
	} else if !value_is_set {
		if default_set && default_is_null {
			value_to_validate = schema_of_parameter.GetObject("default")
		} else if default_set && !default_is_null {
			value_to_validate = schema_of_parameter.GetObject("default")
		} else if !default_set {
			if value_is_mandatory {
				errors = append(errors,  fmt.Errorf("struct: %s column: %s does not have a value or a default value, value_set=%t value_nil=%t default_set=%t default_nil=%t", struct_type, parameter, value_is_set, value_is_null, default_set, default_is_null))
			} else {
				value_to_validate = nil
			}
		}
	}

	if len(errors) > 0 {
		return errors
	} 

	type_of_parameter_value := GetType(value_to_validate)

	if strings.ReplaceAll(*type_of_parameter_schema_value, "*", "") == "time.Time" && IsTime(value_to_validate) {
		type_of_parameter_value = "*time.Time"
	}

	if !((struct_type == "*class.Table" || struct_type == "class.Table") && IsDatabaseColumn(parameter)) {
		if strings.ReplaceAll(*type_of_parameter_schema_value, "*", "") != strings.ReplaceAll(type_of_parameter_value, "*", "") {
			errors = append(errors, fmt.Errorf("table: %s column: %s mismatched schema type expected: %s actual: %s", struct_type, parameter, strings.ReplaceAll(*type_of_parameter_schema_value, "*", ""), strings.ReplaceAll(type_of_parameter_value, "*", "")))
		}
	}

	if len(errors) > 0 {
		return errors
	} 

	switch type_of_parameter_value {
	case "*string", "string":
		var string_value *string
		if type_of_parameter_value == "*string" {
			string_value = value_to_validate.(*string)
		} else {
			temp_string := value_to_validate.(string)
			string_value = &temp_string
		}

		if schema_of_parameter.IsNumber("min_length") {
			min_length, min_length_errors := schema_of_parameter.GetUInt64("min_length")
			if min_length_errors != nil {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s had an error parsing number", struct_type, parameter, "min_length"))
			} else {
				runes := []rune(*string_value)

				if uint64(len(runes)) < *min_length {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: min_length did not meet minimum length requirements and had length: %d", struct_type, parameter, len(runes)))
				}
			}
		}


		if len(errors) > 0 {
			return errors
		} 


		if schema_of_parameter.IsBoolTrue("not_empty_string_value") {
			if *string_value == "" {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s was an empty string", struct_type, parameter, "not_empty_string_value"))
			}
		}


		if len(errors) > 0 {
			return errors
		} 


		if schema_of_parameter.IsNil(FILTERS()) {
			return nil
		}
		
		if !schema_of_parameter.IsArray(FILTERS())  {
			errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s is not an array: %s", struct_type, parameter, FILTERS(), schema_of_parameter.GetType(FILTERS())))
			return errors
		}

		filters, filters_errors := schema_of_parameter.GetArray(FILTERS())
		if filters_errors != nil {
			errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s had error getting array %s", struct_type, parameter, FILTERS(), filters_errors))
		} else if filters == nil {
			return nil
		}

		if len(errors) > 0 {
			return errors
		}

		if len(*filters) == 0 {
			errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s has no filters", struct_type, parameter, FILTERS()))
			return errors
		}

		for filter_index, filter := range *filters {
			filter_map := filter.(Map)

			if !filter_map.HasKey("function") {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is empty", struct_type, parameter, FILTERS(), filter_index))
				return errors
			}

			function := filter_map.Func("function")
			if function == nil {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is nil", struct_type, parameter, FILTERS(), filter_index))
				return errors
			}

			if filter_map.GetType("values") == "nil" || filter_map.GetType("values") == "<nil>" {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d values is nil", struct_type, parameter, FILTERS(), filter_index))
				return errors
			}

			filter_map.SetString("value", string_value)
			filter_map.SetString("data_type", &struct_type)
			filter_map.SetString("label", &parameter)

			temp_map := filter.(Map)
			function_errors := function(temp_map)
			if function_errors != nil {
				errors = append(errors, function_errors...)
			}
		}

		break
	case "*int", "int":
	case "*bool", "bool":
	case "*int64", "int64":
	case "*int32", "int32":
	case "*int16", "int16":
	case "*int8", "int8":
	case "*uint64", "uint64":
	case "*uint32", "uint32":
	case "*uint16", "uint16":
	case "*uint8", "uint8":
	case "*time.Time":
	case "*class.Database":
		database := value_to_validate.(*Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "class.Database":
		database := value_to_validate.(Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "*class.DomainName":
		domain_name := value_to_validate.(*DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "class.DomainName":
		domain_name := value_to_validate.(DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "*class.Host":
		host := value_to_validate.(*Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "class.Host":
		host := value_to_validate.(Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "*class.Credentials":
		credentials := value_to_validate.(*Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "class.Credentials":
		credentials := value_to_validate.(Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "*class.DatabaseCreateOptions":
		database_create_options := value_to_validate.(*DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "class.DatabaseCreateOptions":
		database_create_options := value_to_validate.(DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "*class.Client":
		client := value_to_validate.(*Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "class.Client":
		client := value_to_validate.(Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "*class.Grant":
		grant := value_to_validate.(*Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "class.Grant":
		grant := value_to_validate.(Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "*class.User":
		user := value_to_validate.(*User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "class.User":
		user := value_to_validate.(User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "*class.Table":
		table := value_to_validate.(*Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "class.Table":
		table := value_to_validate.(Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "*class.ClientManager":
		client_manager := value_to_validate.(*ClientManager)

		errors_for_client_manager := client_manager.Validate()
		if errors_for_client_manager != nil {
			errors = append(errors, errors_for_client_manager...)
		}

		break
	case "class.ClientManager":
		client_manager := value_to_validate.(ClientManager)

		errors_for_client_manager := client_manager.Validate()
		if errors_for_client_manager != nil {
			errors = append(errors, errors_for_client_manager...)
		}

		break
	default:
		errors = append(errors, fmt.Errorf("class: %s column: %s type: %s did not meet validation requirements please adjust either your data or table schema (value_nil=%t, value_mandatory=%t, default_nil=%t)", struct_type, parameter, type_of_parameter_value, value_is_null, value_is_mandatory, default_is_null))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}


func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func GetTimeNow() *time.Time {
	now := time.Now()
	return &now
}

func FormatTime(value time.Time) string {
	return value.Format("2006-01-02 15:04:05.000000")
}

func GetTimeNowString() string {
	return (*GetTimeNow()).Format("2006-01-02 15:04:05.000000")
}

func GenerateRandomLetters(length uint64, upper_case *bool) (*string) {
	rand.Seed(time.Now().UnixNano())
	
	var letters_to_use string
	uppercase_letters :=  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase_letters := "abcdefghijklmnopqrstuvwxyz"

	if upper_case == nil {
		letters_to_use = uppercase_letters + lowercase_letters
	} else if *upper_case {
		letters_to_use = uppercase_letters
	} else {
		letters_to_use = lowercase_letters
	}

	var sb strings.Builder

	l := len(letters_to_use)

	for i := uint64(0); i < length; i++ {
		c := letters_to_use[rand.Intn(l)]
		sb.WriteByte(c)
	}

	value := sb.String()
	return &value
}

func GetDirectoryOfExecutable() (*string, error) {
    _, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("filename error")
	}
	directory_name := filepath.Dir(filename)
	directory_name = strings.Replace(directory_name, "/class","", 1)
	directory_name = strings.Replace(directory_name, "/tests","", 1)
	directory_name = strings.Replace(directory_name, "/queue","", 1)
	directory_name = strings.Replace(directory_name, "/go/pkg/mod/github.com/","/go/src/github.com/", 1)
	
	index_of_tag := strings.Index(directory_name, "@")
	if index_of_tag != -1 {
		directory_name = directory_name[:index_of_tag] + "/"
	}

	return &directory_name, nil
}

