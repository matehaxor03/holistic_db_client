package helper

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func GetValidSchemaFields() json.Map {
	schema_fields := json.NewMapValue()
	schema_fields.SetNil("type")
	schema_fields.SetNil("primary_key")
	schema_fields.SetNil("foreign_key")
	schema_fields.SetNil("foreign_key_table_name")
	schema_fields.SetNil("foreign_key_column_name")
	schema_fields.SetNil("foreign_key_type")
	schema_fields.SetNil("unique")
	schema_fields.SetNil("unsigned")
	schema_fields.SetNil("auto_increment")
	schema_fields.SetNil("default")
	schema_fields.SetNil("validated")
	schema_fields.SetNil("filters")
	schema_fields.SetNil("not_empty_string_value")
	schema_fields.SetNil("min_length")
	schema_fields.SetNil("max_length")
	schema_fields.SetNil("decimal_places")
	return schema_fields
}

func getTableName(struct_type string, m *json.Map) (string, []error) {
	temp_value, temp_value_errors := GetField(struct_type, m, "[system_schema]", "[system_fields]", "[table_name]", "string")
	if temp_value_errors != nil {
		return "", temp_value_errors
	} else if temp_value == nil {
		return "", nil
	}
	
	return temp_value.(string), temp_value_errors
}

func GetSchemas(struct_type string, m *json.Map, schema_type string) (*json.Map, []error) {
	var errors []error
	if !(schema_type == "[schema]" || schema_type == "[system_schema]") {
		available_fields := m.GetKeys()
		errors = append(errors, fmt.Errorf("error: %s, %s is not a valid root system schema, available root fields: %s", struct_type, schema_type, available_fields))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	schemas_map, schemas_map_errors := m.GetMap(schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if common.IsNil(schemas_map) {
		errors = append(errors, fmt.Errorf("error: %s %s is nil", struct_type, schema_type))
	} else {
		schema_paramters := schemas_map.GetKeys()
		for _, schema_paramter := range schema_paramters {
			if !schemas_map.IsMap(schema_paramter) {
				errors = append(errors, fmt.Errorf("error: %s %s %s is not a map", struct_type, schema_type, schema_paramter))
			} else {
				schema_paramter_map, schema_paramter_map_errors := schemas_map.GetMap(schema_paramter) 
				if schema_paramter_map_errors != nil {
					errors = append(errors, fmt.Errorf("error: %s %s %s had errors getting map: %s", struct_type, schema_type, schema_paramter, fmt.Sprintf("%s",schema_paramter_map_errors))) 
				} else {
					attributes := schema_paramter_map.GetKeys()
					valid_attributes_map := GetValidSchemaFields()
					for _, attribute := range attributes {
						if !valid_attributes_map.HasKey(attribute) {
							errors = append(errors, fmt.Errorf("error: %s %s %s has an invalid attribute: %s valid attributes are: %s", struct_type, schema_type, schema_paramter, attribute, valid_attributes_map.GetKeys()))
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

func getTableColumns(caller string, data *json.Map) (*[]string, []error) {
	temp_schemas, temp_schemas_error := GetSchemas(caller, data, "[schema]")
	if temp_schemas_error != nil {
		return nil, temp_schemas_error
	}
	columns := temp_schemas.GetKeys()
	return &columns, nil
}

