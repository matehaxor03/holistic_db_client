package dao

import (
	common "github.com/matehaxor03/holistic_common/common"
	"sync"
)

type TableReadRecordsCache struct {
	GetOrSetReadRecords func(table Table, sql string, records *[]Record) (*[]Record, []error)
}

func newTableReadRecordsCache() (*TableReadRecordsCache) {
	cache := map[string](*[]Record){}
	lock_read_records_cache := &sync.RWMutex{}
	
	getOrSetReadRecords := func(table Table, sql string, records *[]Record) (*[]Record, []error) {			
		lock_read_records_cache.Lock()
		defer lock_read_records_cache.Unlock()
		
		database := table.GetDatabase()
		database_name := database.GetDatabaseName()
		host := database.GetHost()
		host_name := host.GetHostName()
		port_number := host.GetPortNumber()
	
		key := host_name + "#" + port_number + "#" + database_name + "#" + sql
		
		if common.IsNil(records) {
			return cache[key], nil
		} else {
			cache[key] = records
			return nil, nil
		}
	}

	return &TableReadRecordsCache{
		GetOrSetReadRecords: func(table Table, sql string, records *[]Record) (*[]Record, []error) {
			return getOrSetReadRecords(table, sql, records)
		},
	}
}
