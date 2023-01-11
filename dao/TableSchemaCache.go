package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"sync"
)

type TableSchemaCache struct {
	GetOrSet func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error)
}

func newTableSchemaCache() (*TableSchemaCache) {
	cache := json.NewMap()
	lock_table_schema_cache := &sync.Mutex{}

	
	getOrSet := func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error) {		
		lock_table_schema_cache.Lock()
		defer lock_table_schema_cache.Unlock()
		var errors []error

		if table_name == "" {
			errors = append(errors, fmt.Errorf("table_name is empty"))
		}

		if mode == "" {
			errors = append(errors, fmt.Errorf("mode is empty"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database_name := database.GetDatabaseName()


		host := database.GetHost()

		if len(errors) > 0 {
			return nil, errors
		}

		host_name, host_name_errors := host.GetHostName()
		if host_name_errors != nil {
			errors = append(errors, host_name_errors...)
		} else if common.IsNil(host_name) {
			errors = append(errors, fmt.Errorf("host_name is nil"))
		}

		port_number, port_number_errors := host.GetPortNumber()
		if port_number_errors != nil {
			errors = append(errors, port_number_errors...)
		} else if common.IsNil(port_number) {
			errors = append(errors, fmt.Errorf("port_number is nil"))
		}

		if len(errors) > 0 {
			return nil, errors
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
			if common.IsNil(schema) {
				errors = append(errors, fmt.Errorf("schema is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			cache.SetMap(key, schema)
			return nil, nil
		} else if mode == "delete" {
			if cache.HasKey(key) {
				_, remove_error := cache.RemoveKey(key)
				if remove_error != nil {
					errors = append(errors, remove_error)
					return nil, errors
				} 
			}
			return nil, nil
		} else {
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
