package helper

import (
	//"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	//common "github.com/matehaxor03/holistic_common/common"
)

func GetRecordColumns(caller string, data *json.Map) (*[]string, []error) {
	fields_map, fields_map_errors := GetFields(caller, data, "[fields]")
	if fields_map_errors != nil {
		return nil, fields_map_errors
	}
	columns := fields_map.GetKeys()
	return &columns, nil
}