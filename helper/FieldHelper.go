package helper

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func GetRecordColumns(struct_type string, m *json.Map) (*[]string, []error) {
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
	if fields_map.IsNull(field) {
		if schema_map.HasKey("default") && !schema_map.IsNull("default") {
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

	type_of := common.GetType(result)
	if type_of == "*json.Value" {
		value_to_validate_unboxed := result.(*json.Value).GetObject()
		result = value_to_validate_unboxed
	} else if type_of == "json.Value" {
		value_to_validate_unboxed := result.(json.Value).GetObject()
		result = value_to_validate_unboxed
	} 
	type_of = common.GetType(result)

	if strings.ReplaceAll(type_of, "*", "") != strings.ReplaceAll(*schema_type_value, "*", "") {
		object_type_simple := strings.ReplaceAll(type_of, "*", "")
		schema_type_value_simple := strings.ReplaceAll(*schema_type_value, "*", "") 
		
		if strings.Contains(object_type_simple, "int") && strings.Contains(schema_type_value_simple, "int") {

		} else if strings.Contains(object_type_simple, "float") && strings.Contains(schema_type_value_simple, "float"){

		} else if strings.Contains(object_type_simple, "string") && strings.Contains(schema_type_value_simple, "time.Time") {
			var convert_default_time_string string
			if type_of == "*string" {
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
			errors = append(errors, fmt.Errorf("error: field: %s schema type: %s actual: %s are not a match", field, *schema_type_value, type_of))
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if desired_type == "self" {
		return result, nil
	}

	if desired_type != type_of {
		errors = append(errors, fmt.Errorf("error: field: %s schema type: %s actual: %s is not the desired type: %s", field, *schema_type_value, type_of, desired_type))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
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

	fields_map.SetObjectForMap(parameter, object)
	validated_true := false
	schema_of_parameter_map.SetObjectForMap("validated", validated_true)
	return nil
}

