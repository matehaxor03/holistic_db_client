package dao

import (
	common "github.com/matehaxor03/holistic_common/common"
)

type TableReadRecordsCache struct {
	GetOrSetReadRecords func(database Database, sql string, records *[]Record) (*[]Record, []error)
}

func newTableReadRecordsCache() (*TableReadRecordsCache) {
	cache := map[string](*[]Record){}
	
	getOrSetReadRecords := func(database Database, sql string, records *[]Record) (*[]Record, []error) {		
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

		key := host_name + "#" + port_number + "#" + database_name + "#" + sql
		
		if common.IsNil(records) {
			return cache[key], nil
		} else {
			cache[key] = records
			return nil, nil
		}
	}

	return &TableReadRecordsCache{
		GetOrSetReadRecords: func(database Database, sql string, records *[]Record) (*[]Record, []error) {
			return getOrSetReadRecords(database, sql, records)
		},
	}
}
