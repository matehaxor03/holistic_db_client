package class

import (
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

type TableStatusCache struct {
	GetOrSetTableStatus func(database Database, table_name string, table_status *json.Map) (*json.Map, []error)
}

func newTableStatusCache() (*TableStatusCache) {
	cache := json.NewMapValue()
	
	getOrSetTableStatus := func(database Database, table_name string, table_status *json.Map) (*json.Map, []error) {		
		client, client_errors := database.GetClient()
		if client_errors != nil {
			return nil, client_errors
		} else if common.IsNil(client) {
			return nil, nil
		}

		host, host_errors := client.GetHost()
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
		
		if common.IsNil(table_status) {
			return cache.GetMap(key)
		} else {
			cache.SetMap(key, table_status)
			return nil, nil
		}
	}

	return &TableStatusCache{
		GetOrSetTableStatus: func(database Database, table_name string, table_status *json.Map) (*json.Map, []error) {
			return getOrSetTableStatus(database, table_name, table_status)
		},
	}
}
