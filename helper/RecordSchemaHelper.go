package helper

import (
	//"fmt"
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


func GetRecordNonPrimaryKeyColumnsUpdate(caller string, data *json.Map, table_non_primary_key_columns []string) (*[]string, []error) {
	record_columns, record_columns_errors := GetRecordColumns(caller, data)
	if record_columns_errors != nil {
		return nil, record_columns_errors
	}

	var errors []error
	var columns []string
	for _, record_column := range *record_columns {
		if record_column == "created_date" {
			continue
		}

		if common.Contains(table_non_primary_key_columns, record_column) {
			columns = append(columns, record_column)
		} 
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &columns, nil
}