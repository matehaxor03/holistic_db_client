package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
)

type TableAdditionalSchemaCache struct {
	GetOrSet func(database Database, table_name string, additional_schema *json.Map) (*json.Map, []error)
}

func newTableAdditionalSchemaCache() (*TableAdditionalSchemaCache) {
	cache := json.NewMap()
	
	getOrSet := func(database Database, table_name string, additional_schema *json.Map) (*json.Map, []error) {		
		var errors []error
	
		if table_name == "" {
			errors = append(errors, fmt.Errorf("table_name is empty string"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database_name, database_name_errors := database.GetDatabaseName()
		if database_name_errors != nil {
			errors = append(errors, database_name_errors...)
		} else if common.IsNil(database_name) {
			errors = append(errors, fmt.Errorf("database_name is nil"))
		}

		host, host_errors := database.GetHost()
		if host_errors != nil {
			errors = append(errors, host_errors...)
		} else if common.IsNil(host) {
			errors = append(errors, fmt.Errorf("host is nil"))
		}

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
