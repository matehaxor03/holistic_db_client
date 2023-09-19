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
	schema_fields.SetNil("foreign_keys")
	schema_fields.SetNil("foreign_key_table_name")
	schema_fields.SetNil("foreign_key_column_name")
	schema_fields.SetNil("foreign_key_type")
	schema_fields.SetNil("unique")
	schema_fields.SetNil("unsigned")
	schema_fields.SetNil("auto_increment")
	schema_fields.SetNil("default")
	schema_fields.SetNil("validated")
	schema_fields.SetNil("filters")
	schema_fields.SetNil("min_length")
	schema_fields.SetNil("max_length")
	schema_fields.SetNil("decimal_places")
	return schema_fields
}

func GetSchemas(m json.Map, schema_type string) (*json.Map, []error) {
	var errors []error
	if !(schema_type == "[schema]" || schema_type == "[system_schema]") {
		available_fields := m.GetKeys()
		errors = append(errors, fmt.Errorf("error: %s is not a valid root system schema, available root fields: %s", schema_type, available_fields))
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	schemas_map, schemas_map_errors := m.GetMap(schema_type)
	if schemas_map_errors != nil {
		errors = append(errors, schemas_map_errors...)
	} else if common.IsNil(schemas_map) {
		errors = append(errors, fmt.Errorf("error: %s is nil", schema_type))
	} else {
		schema_paramters := schemas_map.GetKeys()
		for _, schema_paramter := range schema_paramters {
			if !schemas_map.IsMap(schema_paramter) {
				errors = append(errors, fmt.Errorf("error: %s %s is not a map", schema_type, schema_paramter))
			} else {
				schema_paramter_map, schema_paramter_map_errors := schemas_map.GetMap(schema_paramter) 
				if schema_paramter_map_errors != nil {
					errors = append(errors, fmt.Errorf("error: %s %s had errors getting map: %s", schema_type, schema_paramter, fmt.Sprintf("%s",schema_paramter_map_errors))) 
				} else {
					attributes := schema_paramter_map.GetKeys()
					valid_attributes_map := GetValidSchemaFields()
					for _, attribute := range attributes {
						if !valid_attributes_map.HasKey(attribute) {
							errors = append(errors, fmt.Errorf("error: %s %s has an invalid attribute: %s valid attributes are: %s", schema_type, schema_paramter, attribute, valid_attributes_map.GetKeys()))
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

