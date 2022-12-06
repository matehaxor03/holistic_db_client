package class

import (
	"fmt"
	"strings"
	"sync"
)

type TupleCredentials struct {
	host_name *string
	port_number *string
	database_name *string
	database_username *string
 }

type ClientManager struct {
	GetClient func(label string) (*Client, []error)
	GetTupleCredentials func(label string) (*TupleCredentials, []error)
	Validate func() []error
}

func NewClientManager() (*ClientManager, []error) {
	var this_client_manager *ClientManager

	setClientManager := func(client_manager *ClientManager) {
		this_client_manager = client_manager
	}

	getClientManager := func() *ClientManager {
		return this_client_manager
	}
	
	lock_client := &sync.Mutex{}
	lock_tuple := &sync.Mutex{}
	tuple := make(map[string]TupleCredentials)
	database_reserved_words_obj := newDatabaseReservedWords()
	database_name_whitelist_characters_obj := newDatabaseNameCharacterWhitelist()
	table_name_whitelist_characters_obj := newTableNameCharacterWhitelist()
	column_name_whitelist_characters_obj := newColumnNameCharacterWhitelist()

	/*
	data := Map{
		"[fields]":nil,
		"[schema]":nil,
	}*/

	/*
	getData := func() *json.Map {
		return &data
	}*/

	getTupleCredentials := func(label string) (*TupleCredentials, []error) {
		if value, found_value := tuple[label]; found_value {
			return &value, nil
		}

		var errors []error
		parts := strings.Split(label, ":")
		if len(parts) != 5 {
			errors = append(errors, fmt.Errorf("error: database config for %s not in format e.g holistic_db_config:127.0.0.1:3306:holistic_test:root", label))
			return nil, errors
		}
	
		host_name_value := parts[1]
		port_number_value := parts[2]
		database_name_value := parts[3]
		database_username_value := parts[4]

		temp_tuple_creds := TupleCredentials{host_name: &host_name_value, port_number: &port_number_value, database_name: &database_name_value, database_username: &database_username_value}
		tuple[label] = temp_tuple_creds
		return &temp_tuple_creds, nil
	}

	getClient := func(label string) (*Client, []error) {
		var errors []error
		temp_tuple_creds, details_errors := getTupleCredentials(label)
		if details_errors != nil {
			errors = append(errors, details_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		host, host_errors := newHost(*(temp_tuple_creds.host_name), *(temp_tuple_creds.port_number))
		
		if host_errors != nil {
			errors = append(errors, host_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		client, client_errors := newClient(*getClientManager(), host, temp_tuple_creds.database_username, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)

		if client_errors != nil {
			errors = append(errors, client_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		_, use_database_errors := client.UseDatabaseByName(*(temp_tuple_creds.database_name))
		if use_database_errors != nil {
			return nil, use_database_errors
		}

		return client, nil
	}

	validate := func() []error {
		return nil
		//return ValidateData(getData(), "ClientManager")
	}

	x := ClientManager{
		Validate: func() []error {
			return validate()
		},
		GetTupleCredentials: func(label string) (*TupleCredentials, []error) {
			lock_tuple.Lock()
			defer lock_tuple.Unlock()
			return getTupleCredentials(label)
		},
		GetClient: func(label string) (*Client, []error) {
			lock_client.Lock()
			defer lock_client.Unlock()
			return getClient(label)
		},
	}
	setClientManager(&x)

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

