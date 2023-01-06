package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
)

type TableSchemaCache struct {
	GetOrSet func(database Database, table_name string, schema json.Map, mode string) (json.Map, []error)
}

func newTableSchemaCache() (*TableSchemaCache) {
	cache := json.NewMapValue()
	
	getOrSet := func(database Database, table_name string, schema json.Map, mode string) (json.Map, []error) {		
		var errors []error
		host, host_errors := database.GetHost()
		if host_errors != nil {
			return json.NewMapValue(), host_errors
		} else if common.IsNil(host) {
			errors = append(errors, fmt.Errorf("host is nil"))
			return json.NewMapValue(), errors
		}

		host_name, host_name_errors := host.GetHostName()
		if host_name_errors != nil {
			return json.NewMapValue(), host_name_errors
		} else if common.IsNil(host_name) {
			errors = append(errors, fmt.Errorf("host name is nil"))
			return json.NewMapValue(), errors
		}

		port_number, port_number_errors := host.GetPortNumber()
		if port_number_errors != nil {
			return json.NewMapValue(), port_number_errors
		} else if common.IsNil(port_number) {
			errors = append(errors, fmt.Errorf("database name is nil"))
			return json.NewMapValue(), errors
		}

		database_name, database_name_errors := database.GetDatabaseName()
		if database_name_errors != nil {
			return json.NewMapValue(), database_name_errors
		} else if common.IsNil(database_name) {
			errors = append(errors, fmt.Errorf("database name is nil"))
			return json.NewMapValue(), errors
		}

		key := host_name + "#" + port_number + "#" + database_name + "#" + table_name


		if mode == "get" {
			 result_from_cache, result_from_cache_errors := cache.GetMap(key)
			 if result_from_cache_errors != nil {
				return json.NewMapValue(), result_from_cache_errors
			 } else if common.IsNil(result_from_cache) {
				errors = append(errors, fmt.Errorf("cache is not there"))
				return json.NewMapValue(), errors
			 } else {
				return *result_from_cache, nil
			 }
		} else if mode == "set" {
			cache.SetMapValue(key, schema)
			return json.NewMapValue(), nil
		} else if mode == "delete" {
			if cache.HasKey(key) {
				_, remove_errors := cache.RemoveKey(key)
				if remove_errors != nil {
					return json.NewMapValue(), remove_errors
				} 
			}
			return json.NewMapValue(), nil
		} else {
			var errors []error
			errors = append(errors, fmt.Errorf("mode not supported please implement: %s", mode))
			return json.NewMapValue(), errors
		}
	}

	return &TableSchemaCache{
		GetOrSet: func(database Database, table_name string, schema json.Map, mode string) (json.Map, []error) {
			return getOrSet(database, table_name, schema, mode)
		},
	}
}
