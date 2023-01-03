package db_client

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func BlackListStringToUpper(m json.Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(map_values.GetKeys()) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if map_values.HasKey(strings.ToUpper(*str)) {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func getWhitelistCharactersFunc() *func(m json.Map) []error {
	funcValue := WhitelistCharacters
	return &funcValue
}

func WhitelistCharacters(m json.Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has nil map", *data_type, *label))
	} else if len(map_values.GetKeys()) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	var invalid_letters []string
	for _, letter_rune := range *str {
		letter_string := string(letter_rune)

		if !map_values.HasKey(letter_string) {
			invalid_letters = append(invalid_letters, letter_string)
		}
	}

	if len(invalid_letters) > 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhitelistCharacters: has invalid character(s): %s", *data_type, *label, invalid_letters))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func getWhitelistStringFunc() *func(m json.Map) []error {
	function := WhiteListString
	return &function
}

func getBlacklistStringFunc() *func(m json.Map) []error {
	function := BlackListString
	return &function
}

func getBlacklistStringToUpperFunc() *func(m json.Map) []error {
	function := BlackListStringToUpper
	return &function
}

func WhiteListString(m json.Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if  map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has nil map", *data_type, *label))
	} else if len(map_values.GetKeys()) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if !map_values.HasKey(*str) {
		errors = append(errors, fmt.Errorf("error: %s: %s: WhiteListString: did not find value", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func BlackListString(m json.Map) []error {
	var errors []error
	map_values, map_values_errors := m.GetMap("values")
	str, _ := m.GetString("value")
	label, _ := m.GetString("label")
	data_type, _ := m.GetString("data_type")

	if map_values_errors != nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has get map has errors %s", *data_type, *label, fmt.Sprintf("%s", map_values_errors)))
	} else if map_values == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has nil map", *data_type, *label))
	} else if len(map_values.GetKeys()) == 0 {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	if map_values.HasKey(*str) {
		errors = append(errors, fmt.Errorf("error: %s: %s: BlackListString: found value: %s", *data_type, *label, *str))
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateDatabaseColumnName(value string) []error {
	var errors []error
	column_name_params := json.NewMapValue()
	column_name_params.SetObjectForMap("values", GetMySQLColumnNameWhitelistCharacters())
	column_name_params.SetStringValue("value", value)
	column_name_params.SetStringValue("label", "column_name")
	column_name_params.SetStringValue("data_type", "Table")

	column_name_errors := WhitelistCharacters(column_name_params)
	if column_name_errors != nil {
		errors = append(errors, column_name_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateData(data *json.Map, struct_type string) []error {	
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

	if (struct_type == "*db_client.Table" || struct_type  == "db_client.Table") && data.HasKey("[schema_is_nil]") {
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
		for _, parameter := range (*schemas).GetKeys() {
			value_errors := ValidateParameterData(struct_type, schemas, "[schema]", field_parameters, "[fields]", parameter, nil, primary_key_count, auto_increment_count)

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
		for _, parameter := range (*system_schemas).GetKeys() {
			value_errors := ValidateParameterData(struct_type, system_schemas, "[system_schema]", system_field_parameters, "[system_fields]", parameter, nil, primary_key_count, auto_increment_count)
			if value_errors != nil {
				system_field_errors = append(system_field_errors, value_errors...)
			}
		}
	}

	if len(system_field_errors) > 0 {
		errors = append(errors, system_field_errors...)
	}

	if (struct_type == "*db_client.Table" || struct_type == "db_client.Table") && !ignore_identity_errors {
		if *primary_key_count <= 0 {
			errors = append(errors, fmt.Errorf("error: table: %s did not have any primary keys", struct_type))
		}

		if *auto_increment_count > 1 {
			errors = append(errors, fmt.Errorf("error: table: %s had more than one auto_increment attribute on a column", struct_type))
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ValidateParameterData(struct_type string, schemas *json.Map, schemas_type string, parameters *json.Map, parameters_type string, parameter string, value_to_validate interface{}, primary_key_count *int,  auto_increment_count *int) ([]error) {
	var errors []error

	schema_of_parameter, schema_of_parameter_errors := schemas.GetMap(parameter)
	if schema_of_parameter_errors != nil {
		errors = append(errors, fmt.Errorf("error: Common.ValidateParameterData: %s column: %s error getting parameter schema %s", struct_type, parameter, fmt.Sprintf("%s", schema_of_parameter_errors)))
	} else if common.IsNil(schema_of_parameter) {
		errors = append(errors, fmt.Errorf("error: Common.ValidateParameterData: %s column: %s had nil schema", struct_type, parameter))
	} else if !schemas.IsMap(parameter) {
		errors = append(errors, fmt.Errorf("error: Common.ValidateParameterData: %s column: %s is not a map", struct_type, parameter))
	}

	if len(errors) > 0 {
		return errors
	} 

	var value_is_mandatory bool
	var value_is_set bool
	var value_is_null bool

	value_is_mandatory = true
	value_is_set = true
	value_is_null = false

	if !common.IsNil(parameters) {
		if parameters.HasKey(parameter) {
			value_is_set = true
			if parameters.IsNil(parameter) {
				value_is_null = true
			} else {
				value_to_validate = parameters.GetObjectForMap(parameter)
				value_is_null = false
			}
		} else {
			value_is_set = false
			value_is_null = true
		}
	} else {
		if common.IsNil(value_to_validate) {
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
		errors = append(errors, fmt.Errorf("error: struct: %s column: %s error getting \"type\" attribute for schema %s", struct_type, parameter, fmt.Sprintf("%s", type_of_parameter_schema_value_errors)))
	} else if type_of_parameter_schema_value == nil {
		errors = append(errors, fmt.Errorf("error: struct: %s column: %s \"type\" attribute of schema is nil", struct_type, parameter))
	} else {
		if strings.HasPrefix(*type_of_parameter_schema_value, "*") {
			value_is_mandatory = false
		} else {
			value_is_mandatory = true
		}
	}
 
	if struct_type == "*db_client.Table" || struct_type == "db_client.Table" || struct_type == "*db_client.Record" || struct_type == "db_client.Record" {
		if schema_of_parameter.IsBoolTrue("primary_key") {
			value_is_mandatory = true
			*primary_key_count++ 

			if schema_of_parameter.IsBoolTrue("auto_increment") {
				value_is_mandatory = false
				*auto_increment_count++
			}
		}
	}

	if !common.IsNil(parameters) {
		if !schema_of_parameter.HasKey("validated") {
			bool_true := true
			schema_of_parameter.SetBool("validated", &bool_true)
		} else {
				if !schema_of_parameter.IsBool("validated") {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s does not have attribute: validated is not a bool", struct_type, parameter))
				return errors
			} else if schema_of_parameter.IsBoolTrue("validated") {
				return nil
			} else {
				bool_true := true
				schema_of_parameter.SetBool("validated", &bool_true)
			}
		}
	}

	if (struct_type == "*db_client.Table" || struct_type == "db_client.Table") && parameters_type == "[fields]" {
		value_is_mandatory = false
	}

	if len(errors) > 0 {
		return errors
	} 

	type_of_parameter_value := common.GetType(value_to_validate)
	if type_of_parameter_value == "*json.Value" {
		value_to_validate_unboxed := value_to_validate.(*json.Value).GetObject()
		if common.IsNil(value_to_validate_unboxed) {
			value_is_set = false
			value_is_null = true
		} else {
			value_is_set = true
			value_is_null = false
		}
	} else if type_of_parameter_value == "json.Value" {
		value_to_validate_unboxed := value_to_validate.(json.Value).GetObject()
		if common.IsNil(value_to_validate_unboxed) {
			value_is_set = false
			value_is_null = true
		} else {
			value_is_set = true
			value_is_null = false
		}
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
			value_to_validate = schema_of_parameter.GetObjectForMap("default")
		} else if !default_set { 
			if value_is_mandatory {
				errors = append(errors,  fmt.Errorf("error: struct: %s column: %s does not have a value or a default value, value_set=%t value_nil=%t default_set=%t default_nil=%t", struct_type, parameter, value_is_set, value_is_null, default_set, default_is_null))
			} else {
				value_to_validate = nil
			}
		}
	} else if !value_is_set {
		if default_set && default_is_null {
			value_to_validate = schema_of_parameter.GetObjectForMap("default")
		} else if default_set && !default_is_null {
			value_to_validate = schema_of_parameter.GetObjectForMap("default")
		} else if !default_set {
			if value_is_mandatory {
				errors = append(errors,  fmt.Errorf("error: struct: %s column: %s does not have a value or a default value, value_set=%t value_nil=%t default_set=%t default_nil=%t", struct_type, parameter, value_is_set, value_is_null, default_set, default_is_null))
			} else {
				value_to_validate = nil
			}
		}
	}

	if len(errors) > 0 {
		return errors
	} 


	type_of_parameter_value = common.GetType(value_to_validate)
	if type_of_parameter_value == "*json.Value" {
		type_of_parameter_value = value_to_validate.(*json.Value).GetType()
		value_to_validate = value_to_validate.(*json.Value).GetObject()
	} else if type_of_parameter_value == "json.Value" {
		type_of_parameter_value = value_to_validate.(json.Value).GetType()
		value_to_validate = value_to_validate.(json.Value).GetObject()
	}

	if strings.ReplaceAll(*type_of_parameter_schema_value, "*", "") == "time.Time" {
		decimal_places, decimal_places_error := schema_of_parameter.GetInt("decimal_places")
		if decimal_places_error != nil {
			errors = append(errors, decimal_places_error...)
		} else if decimal_places == nil {
			errors = append(errors, fmt.Errorf("decimal places is nil"))
		} else if common.IsTime(value_to_validate, *decimal_places) {
			type_of_parameter_value = "*time.Time"
		}
	}

	if len(errors) > 0 {
		return errors
	} 
	
	if strings.HasPrefix(*type_of_parameter_schema_value, "*") && common.IsNil(value_to_validate) {
		return nil
	}

	if !value_is_mandatory && common.IsNil(value_to_validate) {
		return nil
	}

	if !((struct_type == "*db_client.Table" || struct_type == "db_client.Table") && parameters_type == "[fields]") {
		if strings.ReplaceAll(*type_of_parameter_schema_value, "*", "") != strings.ReplaceAll(type_of_parameter_value, "*", "") {
			type_of_parameter_schema_value_simple := strings.ReplaceAll(*type_of_parameter_schema_value, "*", "")
			type_of_parameter_value_simple := strings.ReplaceAll(type_of_parameter_value, "*", "")
			if strings.Contains(type_of_parameter_schema_value_simple, "int") && strings.Contains(type_of_parameter_value_simple, "int") {

			} else if strings.Contains(type_of_parameter_schema_value_simple, "float") && strings.Contains(type_of_parameter_value_simple, "float"){

			} else {
				errors = append(errors, fmt.Errorf("error: Common.ValidateParameterData %s column: %s mismatched schema type expected: %s actual: %s", struct_type, parameter, *type_of_parameter_schema_value, type_of_parameter_value))
			}
		}
	}

	if len(errors) > 0 {
		return errors
	} 

	if (struct_type == "*db_client.Table" || struct_type == "db_client.Table") && parameters_type == "[fields]" && len(parameters.GetKeys()) == 0 {
		return nil
	}

	switch type_of_parameter_value {
	case "*string", "string":
		var string_value *string
		if type_of_parameter_value == "*string" {
			string_value = value_to_validate.(*string)
		} else {
			temp_string :=  value_to_validate.(string)
			string_value = &temp_string
		}
		

		if schema_of_parameter.IsInteger("min_length") {
			min_length, min_length_errors := schema_of_parameter.GetUInt64("min_length")
			if min_length_errors != nil {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s had an error parsing number", struct_type, parameter, "min_length"))
			} else {
				runes := []rune(*string_value)

				if uint64(len(runes)) < *min_length {
					errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: min_length did not meet minimum length requirements and had length: %d", struct_type, parameter, len(runes)))
				}
			}
		}


		if len(errors) > 0 {
			return errors
		} 


		if schema_of_parameter.IsBoolTrue("not_empty_string_value") {
			if *string_value == "" {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s was an empty string", struct_type, parameter, "not_empty_string_value"))
			}
		}


		if len(errors) > 0 {
			return errors
		} 


		if schema_of_parameter.IsNil("filters") {
			return nil
		}
		
		if !schema_of_parameter.IsArray("filters")  {
			errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s is not an array: %s", struct_type, parameter, "filters", schema_of_parameter.GetType("filters")))
			return errors
		}

		filters, filters_errors := schema_of_parameter.GetArray("filters")
		if filters_errors != nil {
			errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s had error getting array %s", struct_type, parameter, "filters", filters_errors))
		} else if filters == nil {
			return nil
		}

		if len(errors) > 0 {
			return errors
		}

		if len(*(filters.GetValues())) == 0 {
			errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s has no filters", struct_type, parameter, "filters"))
			return errors
		}

		for filter_index, filter := range *(filters.GetValues()) {
			filter_map, filter_map_errors := filter.GetMap()
			if filter_map_errors != nil {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d getting filter had errors %s", struct_type, parameter, "filters", filter_index, fmt.Sprintf("%s", filter_map_errors)))
				return errors
			} else if common.IsNil(filter_map) {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d getting filter is nil", struct_type, parameter, "filters", filter_index))
				return errors
			}

			if !filter_map.HasKey("function") {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d function is empty", struct_type, parameter, "filters", filter_index))
				return errors
			}

			function, function_errors := filter_map.GetFunc("function")
			if function_errors != nil {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d function had errors %s", struct_type, parameter, "filters", filter_index, fmt.Sprintf("%s", function_errors)))
				return errors
			} else if common.IsNil(function) {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d function is nil", struct_type, parameter, "filters", filter_index))
				return errors
			}

			if filter_map.GetType("values") == "nil" || filter_map.GetType("values") == "<nil>" {
				errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d values is nil", struct_type, parameter, "filters", filter_index))
				return errors
			}

			filter_map.SetString("value", string_value)
			filter_map.SetString("data_type", &struct_type)
			filter_map.SetString("label", &parameter)

			function_execution_errors := function(*filter_map)
			if function_execution_errors != nil {
				errors = append(errors, function_execution_errors...)
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
	case "*float32", "float32":
	case "*float64", "float64":
	case "*time.Time", "time.Time":
	case "*db_client.Database":
		database := value_to_validate.(*Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "db_client.Database":
		database := value_to_validate.(Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "*db_client.DomainName":
		domain_name := value_to_validate.(*DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "db_client.DomainName":
		domain_name := value_to_validate.(DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "*db_client.Host":
		host := value_to_validate.(*Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "db_client.Host":
		host := value_to_validate.(Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "*db_client.Credentials":
		credentials := value_to_validate.(*Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "db_client.Credentials":
		credentials := value_to_validate.(Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "*db_client.DatabaseCreateOptions":
		database_create_options := value_to_validate.(*DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "db_client.DatabaseCreateOptions":
		database_create_options := value_to_validate.(DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "*db_client.Client":
		client := value_to_validate.(*Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "db_client.Client":
		client := value_to_validate.(Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "*db_client.Grant":
		grant := value_to_validate.(*Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "db_client.Grant":
		grant := value_to_validate.(Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "*db_client.User":
		user := value_to_validate.(*User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "db_client.User":
		user := value_to_validate.(User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "*db_client.Table":
		table := value_to_validate.(*Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "db_client.Table":
		table := value_to_validate.(Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "*db_client.ClientManager":
		client_manager := value_to_validate.(*ClientManager)

		errors_for_client_manager := client_manager.Validate()
		if errors_for_client_manager != nil {
			errors = append(errors, errors_for_client_manager...)
		}

		break
	case "db_client.ClientManager":
		client_manager := value_to_validate.(ClientManager)

		errors_for_client_manager := client_manager.Validate()
		if errors_for_client_manager != nil {
			errors = append(errors, errors_for_client_manager...)
		}

		break
	default:
		errors = append(errors, fmt.Errorf("error: class: %s column: %s type: %s did not meet validation requirements please adjust either your data or table schema (value_nil=%t, value_mandatory=%t, default_nil=%t)", struct_type, parameter, type_of_parameter_value, value_is_null, value_is_mandatory, default_is_null))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func get_repository_name_characters() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil("0")
	valid_chars.SetNil("1")
	valid_chars.SetNil("2")
	valid_chars.SetNil("3")
	valid_chars.SetNil("4")
	valid_chars.SetNil("5")
	valid_chars.SetNil("6")
	valid_chars.SetNil("7")
	valid_chars.SetNil("8")
	valid_chars.SetNil("9")
	valid_chars.SetNil("A")
	valid_chars.SetNil("B")
	valid_chars.SetNil("C")
	valid_chars.SetNil("D")
	valid_chars.SetNil("E")
	valid_chars.SetNil("F")
	valid_chars.SetNil("G")
	valid_chars.SetNil("H")
	valid_chars.SetNil("I")
	valid_chars.SetNil("J")
	valid_chars.SetNil("K")
	valid_chars.SetNil("L")
	valid_chars.SetNil("M")
	valid_chars.SetNil("N")
	valid_chars.SetNil("O")
	valid_chars.SetNil("P")
	valid_chars.SetNil("Q")
	valid_chars.SetNil("R")
	valid_chars.SetNil("S")
	valid_chars.SetNil("T")
	valid_chars.SetNil("U")
	valid_chars.SetNil("V")
	valid_chars.SetNil("W")
	valid_chars.SetNil("X")
	valid_chars.SetNil("Y")
	valid_chars.SetNil("Z")
	valid_chars.SetNil("_")
	valid_chars.SetNil("-")
	valid_chars.SetNil("a")
	valid_chars.SetNil("b")
	valid_chars.SetNil("c")
	valid_chars.SetNil("d")
	valid_chars.SetNil("e")
	valid_chars.SetNil("f")
	valid_chars.SetNil("g")
	valid_chars.SetNil("h")
	valid_chars.SetNil("i")
	valid_chars.SetNil("j")
	valid_chars.SetNil("k")
	valid_chars.SetNil("l")
	valid_chars.SetNil("m")
	valid_chars.SetNil("n")
	valid_chars.SetNil("o")
	valid_chars.SetNil("p")
	valid_chars.SetNil("q")
	valid_chars.SetNil("r")
	valid_chars.SetNil("s")
	valid_chars.SetNil("t")
	valid_chars.SetNil("u")
	valid_chars.SetNil("v")
	valid_chars.SetNil("w")
	valid_chars.SetNil("x")
	valid_chars.SetNil("y")
	valid_chars.SetNil("z")
	return valid_chars
}

func get_repository_account_name_characters() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil("0")
	valid_chars.SetNil("1")
	valid_chars.SetNil("2")
	valid_chars.SetNil("3")
	valid_chars.SetNil("4")
	valid_chars.SetNil("5")
	valid_chars.SetNil("6")
	valid_chars.SetNil("7")
	valid_chars.SetNil("8")
	valid_chars.SetNil("9")
	valid_chars.SetNil("A")
	valid_chars.SetNil("B")
	valid_chars.SetNil("C")
	valid_chars.SetNil("D")
	valid_chars.SetNil("E")
	valid_chars.SetNil("F")
	valid_chars.SetNil("G")
	valid_chars.SetNil("H")
	valid_chars.SetNil("I")
	valid_chars.SetNil("J")
	valid_chars.SetNil("K")
	valid_chars.SetNil("L")
	valid_chars.SetNil("M")
	valid_chars.SetNil("N")
	valid_chars.SetNil("O")
	valid_chars.SetNil("P")
	valid_chars.SetNil("Q")
	valid_chars.SetNil("R")
	valid_chars.SetNil("S")
	valid_chars.SetNil("T")
	valid_chars.SetNil("U")
	valid_chars.SetNil("V")
	valid_chars.SetNil("W")
	valid_chars.SetNil("X")
	valid_chars.SetNil("Y")
	valid_chars.SetNil("Z")
	valid_chars.SetNil("_")
	valid_chars.SetNil("-")
	valid_chars.SetNil("a")
	valid_chars.SetNil("b")
	valid_chars.SetNil("c")
	valid_chars.SetNil("d")
	valid_chars.SetNil("e")
	valid_chars.SetNil("f")
	valid_chars.SetNil("g")
	valid_chars.SetNil("h")
	valid_chars.SetNil("i")
	valid_chars.SetNil("j")
	valid_chars.SetNil("k")
	valid_chars.SetNil("l")
	valid_chars.SetNil("m")
	valid_chars.SetNil("n")
	valid_chars.SetNil("o")
	valid_chars.SetNil("p")
	valid_chars.SetNil("q")
	valid_chars.SetNil("r")
	valid_chars.SetNil("s")
	valid_chars.SetNil("t")
	valid_chars.SetNil("u")
	valid_chars.SetNil("v")
	valid_chars.SetNil("w")
	valid_chars.SetNil("x")
	valid_chars.SetNil("y")
	valid_chars.SetNil("z")
	return valid_chars
}

func get_branch_name_characters() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil("0")
	valid_chars.SetNil("1")
	valid_chars.SetNil("2")
	valid_chars.SetNil("3")
	valid_chars.SetNil("4")
	valid_chars.SetNil("5")
	valid_chars.SetNil("6")
	valid_chars.SetNil("7")
	valid_chars.SetNil("8")
	valid_chars.SetNil("9")
	valid_chars.SetNil("A")
	valid_chars.SetNil("B")
	valid_chars.SetNil("C")
	valid_chars.SetNil("D")
	valid_chars.SetNil("E")
	valid_chars.SetNil("F")
	valid_chars.SetNil("G")
	valid_chars.SetNil("H")
	valid_chars.SetNil("I")
	valid_chars.SetNil("J")
	valid_chars.SetNil("K")
	valid_chars.SetNil("L")
	valid_chars.SetNil("M")
	valid_chars.SetNil("N")
	valid_chars.SetNil("O")
	valid_chars.SetNil("P")
	valid_chars.SetNil("Q")
	valid_chars.SetNil("R")
	valid_chars.SetNil("S")
	valid_chars.SetNil("T")
	valid_chars.SetNil("U")
	valid_chars.SetNil("V")
	valid_chars.SetNil("W")
	valid_chars.SetNil("X")
	valid_chars.SetNil("Y")
	valid_chars.SetNil("Z")
	valid_chars.SetNil("_")
	valid_chars.SetNil("-")
	valid_chars.SetNil("a")
	valid_chars.SetNil("b")
	valid_chars.SetNil("c")
	valid_chars.SetNil("d")
	valid_chars.SetNil("e")
	valid_chars.SetNil("f")
	valid_chars.SetNil("g")
	valid_chars.SetNil("h")
	valid_chars.SetNil("i")
	valid_chars.SetNil("j")
	valid_chars.SetNil("k")
	valid_chars.SetNil("l")
	valid_chars.SetNil("m")
	valid_chars.SetNil("n")
	valid_chars.SetNil("o")
	valid_chars.SetNil("p")
	valid_chars.SetNil("q")
	valid_chars.SetNil("r")
	valid_chars.SetNil("s")
	valid_chars.SetNil("t")
	valid_chars.SetNil("u")
	valid_chars.SetNil("v")
	valid_chars.SetNil("w")
	valid_chars.SetNil("x")
	valid_chars.SetNil("y")
	valid_chars.SetNil("z")
	return valid_chars
}





