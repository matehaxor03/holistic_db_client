package helper

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTableColumns(data json.Map) (*map[string]bool, []error) {
	temp_schemas, temp_schemas_error := GetSchemas(data, "[schema]")
	if temp_schemas_error != nil {
		return nil, temp_schemas_error
	}
	columns_from_schema := temp_schemas.GetKeys()
	columns := make(map[string]bool)
	for _, column := range columns_from_schema {
		columns[column] = true
	}
	return &columns, nil
}

func GetTablePrimaryKeyColumnsForSchema(schema_map json.Map) (*map[string]bool, []error) {
	var errors []error
	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
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

func GetTablePrimaryKeyColumns(data json.Map) (*map[string]bool, []error) {
	var errors []error
	
	schema_map, schema_map_errors := GetSchemas(data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
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

func GetTableForeignKeyColumnsForSchema(schema_map json.Map) (*map[string]bool, []error) {
	var errors []error

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
			continue
		}

		if column_schema.IsBoolTrue("foreign_key") {
			columns[column] = true
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableForeignKeyColumns(data json.Map) (*map[string]bool, []error) {
	var errors []error
	
	schema_map, schema_map_errors := GetSchemas(data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
			continue
		}

		if column_schema.IsBoolTrue("foreign_key") {
			columns[column] = true
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableIdentityColumns(data json.Map) (*map[string]bool, []error) {
	var errors []error

	schema_map, schema_map_errors := GetSchemas(data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
			continue
		}

		if column_schema.IsBoolTrue("primary_key") || column_schema.IsBoolTrue("foreign_key") {
			columns[column] = true
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetTableNonPrimaryKeyColumns(data json.Map) (*map[string]bool, []error) {
	var errors []error

	schema_map, schema_map_errors := GetSchemas(data, "[schema]")
	if schema_map_errors != nil {
		errors = append(errors, schema_map_errors...)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
			continue
		}

		if !(column_schema.IsBoolTrue("primary_key")) {
			columns[column] = true
		}
	}
	return &columns, nil
}

func GetTableNonPrimaryKeyColumnsForSchema(schema_map json.Map) (*map[string]bool, []error) {
	var errors []error

	columns := make(map[string]bool)
	for _, column := range schema_map.GetKeys() {
		column_schema, column_schema_errors := schema_map.GetMap(column)
		if column_schema_errors != nil {
			errors = append(errors, column_schema_errors...)
			continue
		} else if column_schema == nil {
			errors = append(errors, fmt.Errorf("error: schema: %s is nill", column))
			continue
		}

		if !(column_schema.IsBoolTrue("primary_key")) {
			columns[column] = true
		}
	}
	return &columns, nil
}


