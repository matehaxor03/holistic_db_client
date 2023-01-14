package dao

import (
	"sync"
	"fmt"
)

type TableCache struct {
	GetOrSet func(database Database, table_name string, table *Table, mode string) (*Table, []error)
}

func newTableCache() (*TableCache) {
	cache := make(map[string]interface{})
	lock := &sync.RWMutex{}

	
	getOrSet := func(database Database, table_name string, table *Table, mode string) (*Table, []error) {		
		lock.Lock()
		defer lock.Unlock()
		var errors []error

		if table_name == "" {
			errors = append(errors, fmt.Errorf("table_name is empty"))
		}

		if mode == "" {
			errors = append(errors, fmt.Errorf("mode is empty"))
		} else if mode == "set" {
			if table == nil {
				errors = append(errors, fmt.Errorf("table is nil"))
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
			return result_from_cache.(*Table), nil
		} else if mode == "set" {
			cache[key] = table
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

	return &TableCache{
		GetOrSet: func(database Database, table_name string, table *Table, mode string) (*Table, []error) {
			return getOrSet(database, table_name, table, mode)
		},
	}
}
