package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type TableAdditionalSchemaCache struct {
	GetOrSet func(database Database, table_name string, additional_schema *json.Map) (*json.Map, []error)
}

func newTableAdditionalSchemaCache() (*TableAdditionalSchemaCache) {
	cache := json.NewMapValue()
	
	getOrSet := func(database Database, table_name string, additional_schema *json.Map) (*json.Map, []error) {		
		host, host_errors := database.GetHost()
		if host_errors != nil {
			return nil, host_errors
		} else if common.IsNil(host) {
			return nil, nil
		}

		host_name, host_name_errors := host.GetHostName()
		if host_name_errors != nil {
			return nil, host_name_errors
		} else if common.IsNil(host_name) {
			return nil, nil
		}

		port_number, port_number_errors := host.GetPortNumber()
		if port_number_errors != nil {
			return nil, port_number_errors
		} else if common.IsNil(port_number) {
			return nil, nil
		}

		database_name, database_name_errors := database.GetDatabaseName()
		if database_name_errors != nil {
			return nil, database_name_errors
		} else if common.IsNil(database_name) {
			return nil, nil
		}

		key := host_name + "#" + port_number + "#" + database_name + "#" + table_name
		
		if common.IsNil(additional_schema) {
			return cache.GetMap(key)
		} else {
			cache.SetMap(key, additional_schema)
			return nil, nil
		}
	}

	return &TableAdditionalSchemaCache{
		GetOrSet: func(database Database, table_name string, additional_schema *json.Map) (*json.Map, []error) {
			return getOrSet(database, table_name, additional_schema)
		},
	}
}
