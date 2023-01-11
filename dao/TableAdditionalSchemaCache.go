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

		database_name := database.GetDatabaseName()
		host := database.GetHost()

		if len(errors) > 0 {
			return nil, errors
		}

		host_name := host.GetHostName()
		port_number := host.GetPortNumber()
		
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
