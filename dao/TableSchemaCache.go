package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	"sync"
	"fmt"
)

type TableSchemaCache struct {
	GetOrSet func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error)
}

func newTableSchemaCache() (*TableSchemaCache) {
	cache := make(map[string]interface{})
	lock_table_schema_cache := &sync.RWMutex{}

	
	getOrSet := func(database Database, table_name string, schema *json.Map, mode string) (*json.Map, []error) {		
		lock_table_schema_cache.Lock()
		defer lock_table_schema_cache.Unlock()
		var errors []error

		if table_name == "" {
			errors = append(errors, fmt.Errorf("table_name is empty"))
		}

		if mode == "" {
			errors = append(errors, fmt.Errorf("mode is empty"))
		} else if mode == "set" {
			if schema == nil {
				errors = append(errors, fmt.Errorf("schema is nil"))
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database_name := database.GetDatabaseName()


		host := database.GetHost()

		if len(errors) > 0 {
			return nil, errors
		}

		host_name := host.GetHostName()
		port_number := host.GetPortNumber()

		key := host_name + "#" + port_number + "#" + database_name + "#" + table_name


		if mode == "get" {
			 result_from_cache, found := cache[key]
			 if !found {
				return nil, nil
			} 
			return result_from_cache.(*json.Map), nil
		} else if mode == "set" {
			cache[key] = schema
			return nil, nil
		} else if mode == "delete" {
			_, found := cache[key]
			if found {
				delete(cache, key)
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
