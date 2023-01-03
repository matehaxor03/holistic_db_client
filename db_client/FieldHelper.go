package db_client

import (
	"fmt"
	"strings"
	"time"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func getRecordColumns(struct_type string, m *json.Map) (*[]string, []error) {
	fields_map, fields_map_errors := GetFields(struct_type, m, "[fields]")
	if fields_map_errors != nil {
		return nil, fields_map_errors
	}
	columns := fields_map.GetKeys()
	return &columns, nil
}

func GetFields(struct_type string, m *json.Map, field_type string) (*json.Map, []error) {
	var errors []error
	if !(field_type == "[fields]" || field_type == "[system_fields]") {
		available_fields := m.GetKeys()
		errors = append(errors, fmt.Errorf("error: %s %s is not a valid root field, available root fields: %s", struct_type, field_type, available_fields))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	fields_map, fields_map_errors := m.GetMap(field_type)
	if fields_map_errors != nil {
		errors = append(errors, fields_map_errors...)
	} else if common.IsNil(fields_map) {
		errors = append(errors, fmt.Errorf("error: %s %s is nil", struct_type, field_type))
	} 

	if len(errors) > 0 {
		return nil, errors
	}
	
	return fields_map, nil
}

func GetField(struct_type string, m *json.Map, schema_type string, field_type string, field string, desired_type string) (interface{}, []error) {
	var errors []error
	schemas_map, schemas_map_errors := GetSchemas(struct_type, m, schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if !schemas_map.HasKey(field) {
		available_fields := schemas_map.GetKeys()
		errors = append(errors, fmt.Errorf("error: Common.GetField %s schema_type: %s field: %s does not exist available fields are: %s", struct_type, schema_type, field, fmt.Sprintf("%s", available_fields)))
	} else if !schemas_map.IsMap(field) {
		errors = append(errors, fmt.Errorf("error: Common.GetField %s schema_type: %s field: %s is not a map and has type: %s", struct_type, schema_type, field, schemas_map.GetType(field)))
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
		errors = append(errors, fmt.Errorf("error: %s %s map is nil", struct_type, schema_type))
	} else if !schema_map.HasKey("type") {
		available_fields := schemas_map.GetKeys()
		errors = append(errors, fmt.Errorf("error: %s field: %s schema \"type\" attribute does not exist available fields are: %s", struct_type, field, fmt.Sprintf("%s", available_fields)))
	} else if !schema_map.IsString("type") {
		errors = append(errors, fmt.Errorf("error: %s field: %s schema \"type\" attribute value is not a string it's %s", struct_type, field, schema_map.GetType("type")))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	schema_type_value, schema_type_value_errors := schema_map.GetString("type")
	if schema_type_value_errors != nil {
		errors = append(errors, schema_type_value_errors...)
	} else if schema_type_value == nil {
		errors = append(errors, fmt.Errorf("error: field: %s schema type is nil"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var result interface{}
	if fields_map.IsNil(field) {
		if schema_map.HasKey("default") && !schema_map.IsNil("default") {
			result = schema_map.GetObjectForMap("default")
		} else {
			result = nil
		}
	} else {
		result = fields_map.GetObjectForMap(field)
	}

	if common.IsNil(result) && strings.HasPrefix(*schema_type_value, "*") {
		if schema_map.IsBoolTrue("auto_increment") && schema_map.IsBoolTrue("primary_key") {
			return nil, nil
		}

		if schema_map.IsBoolTrue("primary_key") {
			errors = append(errors,	fmt.Errorf("error: field: %s had nil value and default value but is a primary key"))
			return nil, errors
		}

		if strings.HasPrefix(*schema_type_value, "*") {
			return nil, nil
		}

		errors = append(errors,	fmt.Errorf("error: field: %s had nil value and default value but is not nullable"))
		return nil, errors
	} else if common.IsNil(result) {
		errors = append(errors,	fmt.Errorf("error: field: %s had nil value and default value but is not nullable"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	object_type := common.GetType(result)
	if object_type == "*json.Value" {
		value_to_validate_unboxed := result.(*json.Value).GetObject()
		result = value_to_validate_unboxed
	} else if object_type == "json.Value" {
		value_to_validate_unboxed := result.(json.Value).GetObject()
		result = value_to_validate_unboxed
	} 
	object_type = common.GetType(result)


	if strings.ReplaceAll(object_type, "*", "") != strings.ReplaceAll(*schema_type_value, "*", "") {
		object_type_simple := strings.ReplaceAll(object_type, "*", "")
		schema_type_value_simple := strings.ReplaceAll(*schema_type_value, "*", "") 
		
		if strings.Contains(object_type_simple, "int") && strings.Contains(schema_type_value_simple, "int") {

		} else if strings.Contains(object_type_simple, "float") && strings.Contains(schema_type_value_simple, "float"){

		} else if strings.Contains(object_type_simple, "string") && strings.Contains(schema_type_value_simple, "time.Time") {
			var convert_default_time_string string
			if object_type == "*string" {
				convert_default_time_string = *(result.(*string))
			} else {
				convert_default_time_string = result.(string)
			}

			if convert_default_time_string == "zero" {
				return nil, nil
			} else if convert_default_time_string == "now" {
				return common.GetTimeNow(), nil
			}
		} else {
			errors = append(errors, fmt.Errorf("error: field: %s schema type: %s actual: %s are not a match", field, schema_type_value_simple, object_type_simple))
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if desired_type == "self" {
		return result, nil
	}

	type_of := common.GetType(result)

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
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*string":
		switch desired_type {
		case "string":
			return *(result.(*string)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*uint64":
		switch desired_type {
		case "uint64":
			return *(result.(*uint64)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "uint64":
		switch desired_type {
		case "*uint64":
			temp_value := (result.(uint64))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "int64":
		switch desired_type {
		case "*int64":
			temp_value := (result.(int64))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*int64":
		switch desired_type {
		case "int64":
			temp_value := *((result.(*int64)))
			return temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "int32":
		switch desired_type {
		case "*int32":
			temp_value := (result.(int32))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*int32":
		switch desired_type {
		case "int32":
			temp_value := *((result.(*int32)))
			return temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "int16":
		switch desired_type {
		case "*int16":
			temp_value := (result.(int16))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*int16":
		switch desired_type {
		case "int16":
			temp_value := *((result.(*int16)))
			return temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "int8":
		switch desired_type {
		case "*int8":
			temp_value := (result.(int8))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*int8":
		switch desired_type {
		case "int8":
			temp_value := *((result.(*int8)))
			return temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "int":
		switch desired_type {
		case "*int":
			temp_value := (result.(int))
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*int":
		switch desired_type {
		case "int":
			temp_value := *((result.(*int)))
			return temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*float32":
		switch desired_type {
		case "float32":
			return *(result.(*float32)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "float32":
		switch desired_type {
		case "*float32":
			temp_value := result.(float32)
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*float64":
		switch desired_type {
		case "float64":
			return *(result.(*float64)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "float64":
		switch desired_type {
		case "*float64":
			temp_value := result.(float64)
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "bool":
		switch desired_type {
		case "*bool":
			temp_value := result.(bool)
			return &temp_value, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*bool":
		switch desired_type {
		case "bool":
			return *(result.(*bool)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*time.Time":
		switch desired_type {
		case "time.Time":
			return *(result.(*time.Time)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.ClientManager":
		switch desired_type {
		case "db_client.ClientManager":
			return *(result.(*ClientManager)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Client":
		switch desired_type {
		case "db_client.Client":
			return *(result.(*Client)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Host":
		switch desired_type {
		case "db_client.Host":
			return *(result.(*Host)), nil
		default:
			errors = append(errors, fmt.Errorf("error: ommon.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Database":
		switch desired_type {
		case "db_client.Database":
			return *(result.(*Database)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Table":
		switch desired_type {
		case "db_client.Table":
			return *(result.(*Table)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Record":
		switch desired_type {
		case "db_client.Record":
			return *(result.(*Record)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.DomainName":
		switch desired_type {
		case "db_client.DomainName":
			return *(result.(*DomainName)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Grant":
		switch desired_type {
		case "db_client.Grant":
			return *(result.(*Grant)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.Credentials":
		switch desired_type {
		case "db_client.Credentials":
			return *(result.(*Credentials)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.User":
		switch desired_type {
		case "db_client.User":
			return *(result.(*User)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "*db_client.DatabaseCreateOptions":
		switch desired_type {
		case "db_client.DatabaseCreateOptions":
			return *(result.(*DatabaseCreateOptions)), nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.ClientManager":
		switch desired_type {
		case "*db_client.ClientManager":
			temp_result := result.(ClientManager)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Client":
		switch desired_type {
		case "*db_client.Client":
			temp_result := result.(Client)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Host":
		switch desired_type {
		case "*db_client.Host":
			temp_result := result.(Host)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Database":
		switch desired_type {
		case "*db_client.Database":
			temp_result := result.(Database)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Table":
		switch desired_type {
		case "*db_client.Table":
			temp_result := result.(Table)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Record":
		switch desired_type {
		case "*db_client.Record":
			temp_result := result.(Record)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.DomainName":
		switch desired_type {
		case "*db_client.DomainName":
			temp_result := result.(DomainName)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Grant":
		switch desired_type {
		case "*db_client.Grant":
			temp_result := result.(Grant)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.Credentials":
		switch desired_type {
		case "*db_client.Credentials":
			temp_result := result.(Credentials)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.User":
		switch desired_type {
		case "*db_client.User":
			temp_result := result.(User)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
		}
	case "db_client.DatabaseCreateOptions":
		switch desired_type {
		case "*db_client.DatabaseCreateOptions":
			temp_result := result.(DatabaseCreateOptions)
			return &temp_result, nil
		default:
			errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement %s->%s", type_of, desired_type))
	}
	default:
		errors = append(errors, fmt.Errorf("error: Common.GetField mapping not supported please implement: %s", type_of))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return fields_map.GetObjectForMap(field), nil
}

func SetField(struct_type string, m *json.Map, schema_type string, parameter_type string, parameter string, object interface{}) ([]error) {
	var errors []error

	schemas_map, schemas_map_errors := GetSchemas(struct_type, m, schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	}

	fields_map, fields_map_errors := GetFields(struct_type, m, parameter_type) 
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
		errors = append(errors, fmt.Errorf("error: field: %s schema map is nil", parameter))
	} else if !schema_of_parameter_map.HasKey("type") {
		available_fields := schema_of_parameter_map.GetKeys()
		errors = append(errors, fmt.Errorf("error: field: %s schema \"type\" attribute does not exist available fields are: %s", parameter, fmt.Sprintf("%s", available_fields)))
	} else if !schema_of_parameter_map.IsString("type") {
		errors = append(errors, fmt.Errorf("error: field: %s schema \"type\" attribute value is not a string it's %s", parameter, schema_of_parameter_map.GetType("type")))
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

	validate_parameters_errors := ValidateParameterData(struct_type, schemas_map, schema_type, nil, parameter_type, parameter, object, primary_key_count, auto_increment_count)
	if validate_parameters_errors != nil {
		errors = append(errors, validate_parameters_errors...)
	}

	if len(errors) > 0 {
		return errors
	}

	fields_map.SetObjectForMap(parameter, object)
	validated_true := true
	schema_of_parameter_map.SetObjectForMap("validated", validated_true)
	return nil
}

