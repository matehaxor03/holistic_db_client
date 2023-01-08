package helper

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTableName(struct_type string, m *json.Map) (string, []error) {
	var errors []error
	temp_value, temp_value_errors := GetField(struct_type, m, "[system_schema]", "[system_fields]", "[table_name]", "string")
	if temp_value_errors != nil {
		errors = append(errors, temp_value_errors...)
	} 
	
	if len(errors) > 0 {
		return "", errors
	}
	
	return temp_value.(string), nil
}

func GetTableColumns(caller string, data *json.Map) (*[]string, []error) {
	temp_schemas, temp_schemas_error := GetSchemas(caller, data, "[schema]")
	if temp_schemas_error != nil {
		return nil, temp_schemas_error
	}
	columns := temp_schemas.GetKeys()
	return &columns, nil
}

func GetTablePrimaryKeyColumns(caller string, data *json.Map) (*map[string]bool, []error) {
	var errors []error
	columns := make(map[string]bool)

	schema_map, schema_map_errors := GetSchemas(caller, data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", caller, column))
			continue
		}

		if column_schema.IsBoolTrue("primary_key") {
			columns[column] = true
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableForeignKeyColumns(caller string, data *json.Map) (*[]string, []error) {
	var errors []error
	var columns []string

	schema_map, schema_map_errors := GetSchemas(caller, data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", caller, column))
			continue
		}

		if column_schema.IsBoolTrue("foreign_key") {
			columns = append(columns, column)
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableIdentityColumns(caller string, data *json.Map) (*[]string, []error) {
	var errors []error
	var columns []string

	schema_map, schema_map_errors := GetSchemas(caller, data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", caller, column))
			continue
		}

		if column_schema.IsBoolTrue("primary_key") || column_schema.IsBoolTrue("foreign_key") {
			columns = append(columns, column)
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableNonPrimaryKeyColumns(caller string, data *json.Map) (*[]string, []error) {
	var errors []error
	var columns []string

	schema_map, schema_map_errors := GetSchemas(caller, data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: %s schema: %s is nill", caller, column))
			continue
		}

		if !(column_schema.IsBoolTrue("primary_key")) {
			columns = append(columns, column)
		}
	}
	return &columns, nil
}

