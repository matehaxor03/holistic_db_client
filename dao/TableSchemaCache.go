package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
)

type TableSchemaCache struct {
	GetOrSet func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error)
}

func newTableSchemaCache() (*TableSchemaCache) {
	cache := json.NewMapValue()
	
	getOrSet := func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error) {		
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


		if mode == "get" {
			return cache.GetMap(key)
		} else if mode == "set" {
			cache.SetMap(key, schema)
			return nil, nil
		} else if mode == "delete" {
			_, remove_errors := cache.RemoveKey(key)
			if remove_errors != nil {
				return nil, remove_errors
			}
			return nil, nil
		} else {
			var errors []error
			errors = append(errors, fmt.Errorf("mode not supported please implement: %s", mode))
			return nil, errors
		}
	}

	return &TableSchemaCache{
		GetOrSet: func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error) {
			return getOrSet(database, table_name, schema, mode)
		},
	}
}
