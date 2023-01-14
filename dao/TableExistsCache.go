package dao

import (
	"sync"
	"fmt"
)

type TableExistsCache struct {
	GetOrSet func(database Database, table_name string, mode string) (bool, []error)
}

func newTableExistsCache() (*TableExistsCache) {
	cache := make(map[string]interface{})
	lock_table_exists_cache := &sync.RWMutex{}

	
	getOrSet := func(database Database, table_name string, mode string) (bool, []error) {		
		lock_table_exists_cache.Lock()
		defer lock_table_exists_cache.Unlock()
		var errors []error

		if table_name == "" {
			errors = append(errors, fmt.Errorf("table_name is empty"))
		}

		if mode == "" {
			errors = append(errors, fmt.Errorf("mode is empty"))
		}

		if len(errors) > 0 {
			return false, errors
		}

		database_name := database.GetDatabaseName()
		host := database.GetHost()
		host_name := host.GetHostName()
		port_number := host.GetPortNumber()
		key := host_name + "#" + port_number + "#" + database_name + "#" + table_name

		if mode == "get" {
			 _, found := cache[key]
			 if found {
				return true, nil
			 }
			 return false, nil
		} else if mode == "set" {
			cache[key] = nil
			return false, nil
		} else if mode == "delete" {
			_, found := cache[key]
			if found {
				delete(cache, key)
			}
			return false, nil
		} else {
			errors = append(errors, fmt.Errorf("mode not supported please implement: %s", mode))
			return false, errors
		}
	}

	return &TableExistsCache{
		GetOrSet: func(database Database, table_name string, mode string) (bool, []error) {
			return getOrSet(database, table_name, mode)
		},
	}
}
