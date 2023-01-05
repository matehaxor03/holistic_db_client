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
			 result_from_cache, result_from_cache_errors := cache.GetMap(key)
			 if result_from_cache_errors != nil {
				return nil, result_from_cache_errors
			 } else if common.IsNil(result_from_cache) {
				return nil, nil
			 } else {
				return result_from_cache, nil
			 }
		} else if mode == "set" {
			clone_schema, clone_schema_errors := schema.Clone()
			if clone_schema_errors != nil {
				return nil, clone_schema_errors
			} else if common.IsNil(clone_schema) {
				var errors []error
				errors = append(errors, fmt.Errorf("cloned schema is nil"))
				return nil, clone_schema_errors
			} else {
				cache.SetMap(key, clone_schema)
				return nil, nil
			}
		} else if mode == "delete" {
			if cache.HasKey(key) {
				_, remove_errors := cache.RemoveKey(key)
				if remove_errors != nil {
					return nil, remove_errors
				} 
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
