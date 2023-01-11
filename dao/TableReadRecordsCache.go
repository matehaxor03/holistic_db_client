package dao

import (
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
)

type TableReadRecordsCache struct {
	GetOrSetReadRecords func(table Table, sql string, records *[]Record) (*[]Record, []error)
}

func newTableReadRecordsCache() (*TableReadRecordsCache) {
	cache := map[string](*[]Record){}
	
	getOrSetReadRecords := func(table Table, sql string, records *[]Record) (*[]Record, []error) {		
		var errors []error
	
		if common.IsNil(records) {
			errors = append(errors, fmt.Errorf("records is nil"))
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database, database_errors := table.GetDatabase()
		if database_errors != nil {
			errors = append(errors, database_errors...)
		} else if common.IsNil(database) {
			errors = append(errors, fmt.Errorf("database is nil"))
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
