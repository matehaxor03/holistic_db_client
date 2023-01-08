package helper

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	//"strings"
	common "github.com/matehaxor03/holistic_common/common"
)

func GetRecordColumns(caller string, data *json.Map) (*[]string, []error) {
	fields_map, fields_map_errors := GetFields(caller, data, "[fields]")
	if fields_map_errors != nil {
		return nil, fields_map_errors
	}
	columns := fields_map.GetKeys()
	return &columns, nil
}


func GetRecordNonPrimaryKeyColumnsUpdate(caller string, data *json.Map, table_non_primary_key_columns *[]string) (*[]string, []error) {
	var errors []error
	if common.IsNil(table_non_primary_key_columns) {
		errors = append(errors, fmt.Errorf("table_non_primary_key_columns is nil. GetRecordPrimaryKeyColumns()"))
	}
	
	record_columns, record_columns_errors := GetRecordColumns(caller, data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordPrimaryKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var columns []string
	for _, record_column := range *record_columns {
		if record_column == "created_date" {
			continue
		}

		if common.Contains(*table_non_primary_key_columns, record_column) {
			columns = append(columns, record_column)
		} 
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}

func GetRecordPrimaryKeyColumns(caller string, data *json.Map, table_primary_key_columns *map[string]bool) (*[]string, []error) {
	var errors []error
	if common.IsNil(table_primary_key_columns) {
		errors = append(errors, fmt.Errorf("table_primary_key_columns is nil. GetRecordPrimaryKeyColumns()"))
	}
	
	record_columns, record_columns_errors := GetRecordColumns(caller, data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordPrimaryKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var columns []string
	for _, record_column := range *record_columns {
		if _, found := (*table_primary_key_columns)[record_column]; found {
			columns = append(columns, record_column)
		}
		/*if common.Contains(*table_primary_key_columns, record_column) {
			columns = append(columns, record_column)
		} */
	}
	return &columns, nil
}

func GetRecordForeignKeyColumns(caller string, data *json.Map, table_foreign_key_columns *map[string]bool) (*map[string]bool, []error) {
	var errors []error
	if common.IsNil(table_foreign_key_columns) {
		errors = append(errors, fmt.Errorf("table_foreign_key_columns is nil. GetRecordForeignKeyColumns()"))
	}

	record_columns, record_columns_errors := GetRecordColumns(caller, data)
	if record_columns_errors != nil {
		errors = append(errors, record_columns_errors...)
	} else if common.IsNil(record_columns) {
		errors = append(errors, fmt.Errorf("record_columns is nil. GetRecordForeignKeyColumns()"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	columns := make(map[string]bool)
	for _, record_column := range *record_columns {
		if _, found := (*table_foreign_key_columns)[record_column]; found {
			columns[record_column] = true
		}
		
		/*if common.Contains(*table_foreign_key_columns, record_column) {
			columns = append(columns, record_column)
		} */
	}
	return &columns, nil
}