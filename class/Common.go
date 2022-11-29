package class

import (
	"fmt"
	"strings"
	"time"
	"unicode"
	"math/rand"
    "path/filepath"
	"runtime"
)

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

	rep := fmt.Sprintf("%T", object)

	if string_value == "%!s("+rep+"=<nil>)" {
		return true
	}

	return false
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


func GetFields(m *Map) (*Map, []error) {
	var errors []error
	if IsNil(m) {
		errors = append(errors, fmt.Errorf("data is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	fields_map, fields_map_errors := m.GetMap("[fields]")
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} else if IsNil(fields_map) {
		errors = append(errors, fmt.Errorf("[fields] is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return fields_map, nil
}

func GetSchemas(m *Map) (*Map, []error) {
	var errors []error
	if IsNil(m) {
		errors = append(errors, fmt.Errorf("data is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	fields_map, fields_map_errors := m.GetMap("[schema]")
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} else if IsNil(fields_map) {
		errors = append(errors, fmt.Errorf("[schema] is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return fields_map, nil
}

func GetField(m *Map, field string) (interface{}, []error) {
	var errors []error

	fields_map, fields_map_errors := GetFields(m)
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} else if !fields_map.HasKey(field) {
		available_fields := fields_map.Keys()
		errors = append(errors, fmt.Errorf("field does not exist: %s available fields are: %s", field, fmt.Sprintf("%s", available_fields)))
	}

	schemas_map, schemas_map_errors := GetSchemas(m)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if !schemas_map.HasKey(field) {
		available_fields := schemas_map.Keys()
		errors = append(errors, fmt.Errorf("field schema does not exist: %s available fields are: %s", field, fmt.Sprintf("%s", available_fields)))
	} else if !schemas_map.IsMap(field) {
		errors = append(errors, fmt.Errorf("field schema: %s is not a map", field))
	} 

	if len(errors) > 0 {
		return nil, errors
	}

	schema_map, schema_map_errors := schemas_map.GetMap(field)
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	} else if schema_map == nil {
		errors = append(errors, fmt.Errorf("schema map is nil"))
	} else if !schema_map.HasKey("type") {
		available_fields := schemas_map.Keys()
		errors = append(errors, fmt.Errorf("field: %s schema \"type\" attribute does not exist available fields are: %s", field, fmt.Sprintf("%s", available_fields)))
	} else if !schema_map.IsString("type") {
		errors = append(errors, fmt.Errorf("field: %s schema \"type\" attribute value is not a string it's %s", field, schema_map.GetType("type")))
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

	// todo check mandatory/default

	if fields_map.IsNil(field) {
		return nil, nil
	}

	object_type := fields_map.GetType(field)
	if object_type != *schema_type_value {
		errors = append(errors, fmt.Errorf("field: %s schema type: %s object type: %s are not a match", field, *schema_type_value, object_type))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return fields_map.GetObject(field), nil
}

func SetField(m *Map, field string, object interface{}) ([]error) {
	var errors []error

	fields_map, fields_map_errors := GetFields(m)
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} 

	schema_map, schema_map_errors := GetSchema(m, field)
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}
	
	if len(errors) > 0 {
		return errors
	}
	
	schema_type, schema_type_errors  := schema_map.GetString("type")
	if schema_type_errors != nil {
		errors = append(errors, schema_type_errors...)
	} else if schema_type == nil {
		errors = append(errors, fmt.Errorf("schema is nil for field: %s", field))
	}

	if len(errors) > 0 {
		return errors
	}

    // todo check mandatory/default


	object_type := GetType(object)
	if object_type != "nil" {
		if *schema_type != object_type {
			errors = append(errors, fmt.Errorf("field: %s schema type: %s object type: %s are not a match", field, *schema_type, object_type))
		}
	}

	if len(errors) > 0 {
		return errors
	}

	fields_map.SetObject(field, object)
	return nil
}

func GetSchema(m *Map, field string) (*Map, []error) {
	var errors []error

	schema_map, schema_map_errors := GetSchemas(m)
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	} else if !schema_map.HasKey(field) {
		available_fields := schema_map.Keys()
		errors = append(errors, fmt.Errorf("schema does not exist: %s available fields are: %s", field, fmt.Sprintf("%s", available_fields)))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if schema_map.IsNil(field) {
		errors = append(errors, fmt.Errorf("schema exists: %s however data of schema is nil", field))
		return nil, errors
	}

	schema_map_data, schema_map_data_errors := schema_map.GetMap(field)
	if schema_map_data_errors != nil {
		errors = append(errors, schema_map_data_errors...)
	} else if IsNil(schema_map_data) {
		errors = append(errors, fmt.Errorf("schema is nil: %s"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if !schema_map_data.HasKey("type") {
		errors = append(errors, fmt.Errorf("schema does not have type attribute"))
	} else if schema_map_data.IsString("type") {
		errors = append(errors, fmt.Errorf("schema type attribute is not a string"))
	} else {
		//todo validate types
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return schema_map_data, nil
}

func GetSchemaType(m *Map, field string) (*string, []error) {
	var errors []error
	schema_map, schema_map_errors := GetSchema(m, field)
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	} else if !schema_map.HasKey("type") {
		errors = append(errors, fmt.Errorf("field: %s schema does not have atrribute: type", field))
	} else if !schema_map.IsString("type") {
		errors = append(errors, fmt.Errorf("field: %s schema is not a string", field))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return schema_map.GetString("type")
}

func GetStringField(m *Map, field string) (*string, []error) {
	var errors []error
	var value *string
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*string":
			value = object_value.(*string)
		case "string":
		temp := object_value.(string)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetStringField: field: %s type: %s not in [*string, string, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetHostField(m *Map, field string) (*Host, []error) {
	var errors []error
	var value *Host
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.Host":
			value = object_value.(*Host)
		case "class.Host":
		temp := object_value.(Host)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetHostField: field: %s type: %s not in [*class.Host, class.Host, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetCredentialsField(m *Map, field string) (*Credentials, []error) {
	var errors []error
	var value *Credentials
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.Credentials":
			value = object_value.(*Credentials)
		case "class.Credentials":
		temp := object_value.(Credentials)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetCredentialsField: field: %s type: %s not in [*class.Credentials, class.Credentials, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetDomainNameField(m *Map, field string) (*DomainName, []error) {
	var errors []error
	var value *DomainName
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.DomainName":
			value = object_value.(*DomainName)
		case "class.DomainName":
		temp := object_value.(DomainName)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetDomainNameField: field: %s type: %s not in [*class.DomainName, class.DomainName, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetUserField(m *Map, field string) (*User, []error) {
	var errors []error
	var value *User
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.User":
			value = object_value.(*User)
		case "class.User":
		temp := object_value.(User)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetHostField: field: %s type: %s not in [*class.User, class.User, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetTableField(m *Map, field string) (*Table, []error) {
	var errors []error
	var value *Table
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.Table":
			value = object_value.(*Table)
		case "class.Table":
		temp := object_value.(Table)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetTableField: field: %s type: %s not in [*class.Table, class.Table, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetDatabaseField(m *Map, field string) (*Database, []error) {
	var errors []error
	var value *Database
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.Database":
			value = object_value.(*Database)
		case "class.Database":
		temp := object_value.(Database)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetHostField: field: %s type: %s not in [*class.Database, class.Database, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetDatabaseCreateOptionsField(m *Map, field string) (*DatabaseCreateOptions, []error) {
	var errors []error
	var value *DatabaseCreateOptions
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.DatabaseCreateOptions":
			value = object_value.(*DatabaseCreateOptions)
		case "class.DatabaseCreateOptions":
		temp := object_value.(DatabaseCreateOptions)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("DatabaseCreateOptions: field: %s type: %s not in [*class.DatabaseCreateOptions, class.DatabaseCreateOptions, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
}

func GetClientField(m *Map, field string) (*Client, []error) {
	var errors []error
	var value *Client
	
	object_value, object_value_errors := GetField(m, field)
	if object_value_errors != nil {
		return nil, object_value_errors
	} 

	type_of := GetType(object_value)
	
	switch type_of {
		case "*class.Client":
			value = object_value.(*Client)
		case "class.Client":
		temp := object_value.(Client)
		value = &temp
		case "nil":
		value = nil
		default:
			errors = append(errors, fmt.Errorf("GetClientField: field: %s type: %s not in [*class.Client, class.Client, nil]" , field, type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return value, nil
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

func ArraysContainsArraysOrdered(a [][]string, b [][]string, label string, typeof string) []error {
	var errors []error

	for _, b_value := range b {
		var current = strings.Join(b_value, "")
		var found = false
		for _, a_value := range a {
			var compare = strings.Join(a_value, "")

			if current == compare {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("%s %s has value '%s' expected to have value in '%s", typeof, label, b_value, a))
		}
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

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsDatabaseColumn(value string) bool {
	column_name_params := Map{"values": GetColumnNameValidCharacters(), "value": value, "label": "column_name", "data_type": "Table"}
	column_name_errors := WhitelistCharacters(column_name_params)
	return column_name_errors == nil
}

func ValidateData(data *Map, struct_type string) []error {	
	var errors []error
	var ignore_identity_errors = false
	primary_key_count := 0
	auto_increment_count := 0

	parameters, parameters_errors := GetFields(data)
	if parameters_errors != nil {
		errors = append(errors, parameters_errors...)
	}

	schemas, schemas_errors := GetSchemas(data)
	if schemas_errors != nil {
		errors = append(errors, schemas_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	if struct_type == "*class.Table" && data.HasKey("[schema_is_nil]") {
		if data.IsBoolTrue("[schema_is_nil]") {
			ignore_identity_errors = true
		}
	}

	//todo check schema mandatory fields not in parameters

	for _, parameter := range (*parameters).Keys() {
		if parameters.GetType(parameter) != "class.Map" {
			errors = append(errors, fmt.Errorf("struct: %s column: %s is not of type class.Map", struct_type, parameter))
			continue
		}

		paramter_value, paramter_value_errors := GetField(data, parameter)
		if paramter_value_errors != nil {
			errors = append(errors, fmt.Errorf("struct: %s column: %s error getting parameter value %s", struct_type, parameter, fmt.Sprintf("%s", paramter_value_errors)))
			continue
		}

		schema_of_parameter, schema_of_parameter_errors := GetSchema(data, parameter)
		if schema_of_parameter_errors != nil {
			errors = append(errors, fmt.Errorf("struct: %s column: %s error getting parameter schema %s", struct_type, parameter, fmt.Sprintf("%s", schema_of_parameter_errors)))
			continue
		}

		type_of_parameter := GetType(paramter_value)
		type_of_schema, type_of_schema_errors := schema_of_parameter.GetString("type")
		if type_of_schema_errors != nil {
			errors = append(errors, fmt.Errorf("struct: %s column: %s error getting type for schema %s", struct_type, parameter, fmt.Sprintf("%s", type_of_schema_errors)))
			continue
		} else if type_of_schema == nil {
			errors = append(errors, fmt.Errorf("struct: %s column: %s error type of schema is nil", struct_type, parameter))
			continue
		}

		value_is_mandatory := true
		value_is_null := parameters.IsNil(parameter)
		default_set := false
		default_is_null := false
		default_value := schema_of_parameter.GetObject("default")

		if schema_of_parameter.HasKey("default") {
			default_set = true
			if schema_of_parameter.IsNil("default") {
				default_is_null = true
			}
		}

		if schema_of_parameter.HasKey("mandatory") {
		    if !schema_of_parameter.IsBool("mandatory") {
				errors = append(errors, fmt.Errorf("struct: %s column: %s had mandatory schema field however was not a bool: %s", struct_type, parameter, schema_of_parameter.GetType("mandatory")))
				continue
			} else if schema_of_parameter.IsBoolFalse("mandatory") {
				value_is_mandatory = false
			}
		}  

		if struct_type == "*class.Table" {
			if schema_of_parameter.IsBoolTrue("primary_key") {
				value_is_mandatory = true
				primary_key_count += 1

				if schema_of_parameter.IsBoolTrue("auto_increment") {
					auto_increment_count += 1
				}
			}

			if IsDatabaseColumn(parameter) && (!schema_of_parameter.HasKey("type") || !schema_of_parameter.IsString("type")) {
				errors = append(errors, fmt.Errorf("table: %s column: %s is missing type attribute", struct_type, parameter))
				continue
			}
		}

		if !schema_of_parameter.HasKey("validated") {
			bool_true := true
			schema_of_parameter.SetBool("validated", &bool_true)
		} else {
		 	 if !schema_of_parameter.IsBool("validated") {
				errors = append(errors, fmt.Errorf("table: %s column: %s does not have attribute: validated is not a bool", struct_type, parameter))
			} else if schema_of_parameter.IsBoolTrue("validated") {
				continue
			} else {
				bool_true := true
				schema_of_parameter.SetBool("validated", &bool_true)
			}
		}
	
		if len(errors) > 0 {
			return errors
		}

		if value_is_null && default_is_null && (schema_of_parameter.IsBoolTrue("primary_key") && schema_of_parameter.IsBoolTrue("auto_increment")) {
			continue
		}

		var value_to_validate interface{}
		value_to_validate = parameters.GetObject(parameter)
		if value_is_null && !value_is_mandatory && default_is_null && !default_set {
			continue
		} else if value_is_null && !default_is_null && default_set {
			value_to_validate = schema_of_parameter.GetObject("default")
		} 

		typeOf := fmt.Sprintf("%T", value_to_validate)

		type_of_parameter_value, type_of_parameter_value_errors := schema_of_parameter.GetString("type")
		if type_of_parameter_value_errors != nil {
			errors = append(errors, type_of_parameter_value_errors...)
			continue
		}

		if typeOf != *type_of_parameter_value {
			errors = append(errors, fmt.Errorf("table: %s column: %s schema does not match expected: %s actual: %s", struct_type, parameter, type_of_parameter_value, typeOf))
			continue
		}

		switch typeOf {
		case "*string", "string":
			var string_value *string
			if typeOf == "*string" {
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
					var runes []rune
					for _, runeValue := range *string_value {
						runes = append(runes, runeValue)
					}

					if uint64(len(*runes)) < *min_length {
						errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s did not meet minimum length requirements and had length: %d", struct_type, parameter, "min_length", len(*runes)))
					}
				}
			}

			if schema_of_parameter.IsBoolTrue("not_empty_string_value") {
				if *string_value == "" {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s was an empty string", struct_type, parameter, "not_empty_string_value"))
				}
			}

			if schema_of_parameter.IsNil(FILTERS()) {
				continue
			}
			
			if !schema_of_parameter.IsArray(FILTERS())  {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s is not an array: %s", struct_type, parameter, FILTERS(), parameter_fields.GetType(FILTERS())))
				continue
			}

			filters, filters_errors := schema_of_parameter.GetArray(FILTERS())
			if filters_errors != nil {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s had error getting array %s", struct_type, parameter, FILTERS(), filters_errors))
				continue
			}

			if filters == nil {
				continue
			}

			if len(*filters) == 0 {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s has no filters", struct_type, parameter, FILTERS()))
				continue
			}

			for filter_index, filter := range *filters {
				filter_map := filter.(Map)

				if !filter_map.HasKey("function") {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is empty", struct_type, parameter, FILTERS(), filter_index))
					continue
				}

				function := filter_map.Func("function")
				if function == nil {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is nil", struct_type, parameter, FILTERS(), filter_index))
					continue
				}

				if filter_map.GetType("values") == "nil" || filter_map.GetType("values") == "<nil>" {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d values is nil", struct_type, parameter, FILTERS(), filter_index))
					continue
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
		case "*uint64", "uint64":
		case "*time.Time":
		case "*class.Database":
			database := value_to_validate.(*Database)

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
		case "*class.Host":
			host := value_to_validate.(*Host)

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
		case "*class.DatabaseCreateOptions":
			database_create_options := value_to_validate.(*DatabaseCreateOptions)

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
		case "*class.Grant":
			grant := value_to_validate.(*Grant)

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
		case "*class.Table":
			table := value_to_validate.(*Table)

			errors_for_table := table.Validate()
			if errors_for_table != nil {
				errors = append(errors, errors_for_table...)
			}

			break
		case "class.Map":
			json_string, json_string_errors := fields.ToJSONString()
			if json_string_errors != nil {
				errors = append(errors, json_string_errors...)
			} else {
				errors = append(errors, fmt.Errorf("raw: %s class: %s column: %s attribute: %s with type: %s did not meet validation requirements please adjust either your data or table schema (value_nil=%t, value_mandatory=%t, default_nil=%t)", *json_string, struct_type, parameter, attribute_to_validate, typeOf, value_is_null, value_is_mandatory, default_is_null))
			}
		default:
			errors = append(errors, fmt.Errorf("class: %s column: %s attribute: %s with type: %s did not meet validation requirements please adjust either your data or table schema (value_nil=%t, value_mandatory=%t, default_nil=%t)", struct_type, parameter, attribute_to_validate, typeOf, value_is_null, value_is_mandatory, default_is_null))
		}
	}

	if struct_type == "*class.Table" && !ignore_identity_errors {
		if primary_key_count <= 0 {
			errors = append(errors, fmt.Errorf("table: %s did not have any primary keys", struct_type))
		}

		if auto_increment_count > 1 {
			errors = append(errors, fmt.Errorf("table: %s had more than one auto_increment attribute on a column", struct_type))
		}
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

