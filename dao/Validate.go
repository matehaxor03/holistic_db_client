
package dao

import (
    json "github.com/matehaxor03/holistic_json/json"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	common "github.com/matehaxor03/holistic_common/common"
	"strings"
	"fmt"
)

func ValidateData(data *json.Map, struct_type string) []error {	
	var errors []error
	
	var primary_key_count *int
	primary_key_count_value := 0 
	primary_key_count = &primary_key_count_value
	
	var auto_increment_count *int
	auto_increment_count_value := 0 
	auto_increment_count = &auto_increment_count_value

	if len(errors) > 0 {
		return errors
	}

	var field_errors []error
	field_parameters, field_parameters_errors := helper.GetFields(struct_type, data, "[fields]")
	if field_parameters_errors != nil {
		field_errors = append(field_errors, field_parameters_errors...)
	}

	schemas, schemas_errors := helper.GetSchemas(struct_type, data, "[schema]")
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
	system_field_parameters, system_field_parameters_errors := helper.GetFields(struct_type, data, "[system_fields]")
	if system_field_parameters_errors != nil {
		system_field_errors = append(system_field_errors, system_field_parameters_errors...)
	}

	system_schemas, system_schemas_errors := helper.GetSchemas(struct_type, data, "[system_schema]")
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

	if ((struct_type == "*dao.Table" || struct_type == "dao.Table")) {
		if *primary_key_count <= 0 {
			errors = append(errors, fmt.Errorf("error: table: %s did not have any primary keys and had keys", struct_type))
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
			if parameters.IsNull(parameter) {
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
		if schema_of_parameter.IsNull("default") {
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
 
	if struct_type == "*dao.Table" || struct_type == "dao.Table" || struct_type == "*dao.Record" || struct_type == "dao.Record" {
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

	if (struct_type == "*dao.Table" || struct_type == "dao.Table") && parameters_type == "[fields]" {
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

	if !((struct_type == "*dao.Table" || struct_type == "dao.Table") && parameters_type == "[fields]") {
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

	if (struct_type == "*dao.Table" || struct_type == "dao.Table") && parameters_type == "[fields]" && len(parameters.GetKeys()) == 0 {
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


		if schema_of_parameter.IsNull("filters") {
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

			function_uncast := filter_map.GetObjectForMap("function")

			//fmt.Println(common.GetType(*function))
			//if function_errors != nil {
			//	errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d function had errors %s", struct_type, parameter, "filters", filter_index, fmt.Sprintf("%s", function_errors)))
			//	return errors
			//} 
			
			// todo: fix npe check 
			function := function_uncast.(*func(string) []error)
			//fmt.Println(function)
			//fmt.Println(fmt.Sprintf("%s", function))
			//fmt.Println(fmt.Sprintf("%s", common.GetType(function)))



			//if common.IsNil(function) {
			//	errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d function is nil", struct_type, parameter, "filters", filter_index))
			//	return errors
			//}

			//if filter_map.GetType("values") == "nil" || filter_map.GetType("values") == "<nil>" {
			//	errors = append(errors, fmt.Errorf("error: table: %s column: %s attribute: %s at index: %d values is nil", struct_type, parameter, "filters", filter_index))
			//	return errors
			//}

			//filter_map.SetString("value", string_value)
			//filter_map.SetString("data_type", &struct_type)
			//filter_map.SetString("label", &parameter)

			function_execution_errors := (*function)(*string_value)
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
	case "*dao.Database":
		database := value_to_validate.(*Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "dao.Database":
		database := value_to_validate.(Database)

		errors_for_database := database.Validate()
		if errors_for_database != nil {
			errors = append(errors, errors_for_database...)
		}
		break
	case "*dao.DomainName":
		domain_name := value_to_validate.(*DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "dao.DomainName":
		domain_name := value_to_validate.(DomainName)

		errors_for_domain_name := domain_name.Validate()
		if errors_for_domain_name != nil {
			errors = append(errors, errors_for_domain_name...)
		}
		break
	case "*dao.Host":
		host := value_to_validate.(*Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "dao.Host":
		host := value_to_validate.(Host)

		errors_for_host := host.Validate()
		if errors_for_host != nil {
			errors = append(errors, errors_for_host...)
		}

		break
	case "*dao.Credentials":
		credentials := value_to_validate.(*Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "dao.Credentials":
		credentials := value_to_validate.(Credentials)

		errors_for_credentaials := credentials.Validate()
		if errors_for_credentaials != nil {
			errors = append(errors, errors_for_credentaials...)
		}

		break
	case "*dao.DatabaseCreateOptions":
		database_create_options := value_to_validate.(*DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "dao.DatabaseCreateOptions":
		database_create_options := value_to_validate.(DatabaseCreateOptions)

		errors_for_database_create_options := database_create_options.Validate()
		if errors_for_database_create_options != nil {
			errors = append(errors, errors_for_database_create_options...)
		}

		break
	case "*dao.Grant":
		grant := value_to_validate.(*Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "dao.Grant":
		grant := value_to_validate.(Grant)

		errors_for_grant := grant.Validate()
		if errors_for_grant != nil {
			errors = append(errors, errors_for_grant...)
		}

		break
	case "*dao.User":
		user := value_to_validate.(*User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "dao.User":
		user := value_to_validate.(User)

		errors_for_user := user.Validate()
		if errors_for_user != nil {
			errors = append(errors, errors_for_user...)
		}

		break
	case "*dao.Table":
		table := value_to_validate.(*Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "dao.Table":
		table := value_to_validate.(Table)

		errors_for_table := table.Validate()
		if errors_for_table != nil {
			errors = append(errors, errors_for_table...)
		}

		break
	case "*dao.Client":
		client := value_to_validate.(*Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "dao.Client":
		client := value_to_validate.(Client)

		errors_for_client := client.Validate()
		if errors_for_client != nil {
			errors = append(errors, errors_for_client...)
		}

		break
	case "*dao.ClientManager":
		client_manager := value_to_validate.(*ClientManager)

		errors_for_client_manager := client_manager.Validate()
		if errors_for_client_manager != nil {
			errors = append(errors, errors_for_client_manager...)
		}

		break
	case "dao.ClientManager":
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