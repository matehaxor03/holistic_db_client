package helper

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func GetRecordColumns(data json.Map) (*map[string]bool, []error) {
	fields_map, fields_map_errors := GetFields(data, "[fields]")
	if fields_map_errors != nil {
		return nil, fields_map_errors
	}
	columns_from_schema := fields_map.GetKeys()
	columns := make(map[string]bool)
	for _, column := range columns_from_schema {
		columns[column] = true
	}
	return &columns, nil
}


func GetRecordNonPrimaryKeyColumnsUpdate(data json.Map, table_non_primary_key_columns *map[string]bool) (*map[string]bool, []error) {
	var errors []error
	if common.IsNil(table_non_primary_key_columns) {
		errors = append(errors, fmt.Errorf("table_non_primary_key_columns is nil. GetRecordPrimaryKeyColumns()"))
	}
	
	record_columns, record_columns_errors := GetRecordColumns(data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordPrimaryKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for record_column, _  := range *record_columns {
		if record_column == "created_date" {
			continue
		}

		if _, found := (*table_non_primary_key_columns)[record_column]; found {
			columns[record_column] = true
		} 
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetRecordPrimaryKeyColumns(data json.Map, table_primary_key_columns *map[string]bool) (*map[string]bool, []error) {
	var errors []error
	if common.IsNil(table_primary_key_columns) {
		errors = append(errors, fmt.Errorf("table_primary_key_columns is nil. GetRecordPrimaryKeyColumns()"))
	}
	
	record_columns, record_columns_errors := GetRecordColumns(data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordPrimaryKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for record_column, _  := range *record_columns {
		if _, found := (*table_primary_key_columns)[record_column]; found {
			columns[record_column] = true
		}

	}
	return &columns, nil
}

func GetRecordForeignKeyColumns(data json.Map, table_foreign_key_columns *map[string]bool) (*map[string]bool, []error) {
	var errors []error
	if common.IsNil(table_foreign_key_columns) {
		errors = append(errors, fmt.Errorf("table_foreign_key_columns is nil. GetRecordForeignKeyColumns()"))
	}

	record_columns, record_columns_errors := GetRecordColumns(data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordForeignKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for record_column, _  := range *record_columns {
		if _, found := (*table_foreign_key_columns)[record_column]; found {
			columns[record_column] = true
		}
	}
	return &columns, nil
}