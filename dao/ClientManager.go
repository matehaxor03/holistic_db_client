package dao

import (
	"sync"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type ClientManager struct {
	GetClient func(host_name string, port_number string, database_name string, database_username string) (*Client, []error)
	Validate func() []error
	GetNextUserCount func() int
}

func NewClientManager() (*ClientManager, []error) {
	verify := validate.NewValidator()
	var this_client_manager *ClientManager

	setClientManager := func(client_manager *ClientManager) {
		this_client_manager = client_manager
	}

	getClientManager := func() *ClientManager {
		return this_client_manager
	}
	
	user_count := 0
	lock_user_count := &sync.Mutex{}
	lock_client := &sync.Mutex{}
	lock_table_schema := &sync.Mutex{}
	lock_table_additional_schema := &sync.Mutex{}
	
	getClient := func(host_name string, port_number string, database_name string, database_username string) (*Client, []error) {
		var errors []error
		host, host_errors := newHost(verify, host_name, port_number)
		
		if host_errors != nil {
			errors = append(errors, host_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		client, client_errors := newClient(verify, *getClientManager(), host, &database_username, nil, lock_table_schema, lock_table_additional_schema)

		if client_errors != nil {
			errors = append(errors, client_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}

		database, database_errors := client.GetDatabaseByName(database_name)
		if database_errors != nil {
			return nil, database_errors
		}
		client.SetDatabase(database)
		database.SetDatabaseUsername(database_username)

		return client, nil
	}

	validate := func() []error {
		return nil
	}

	x := ClientManager{
		Validate: func() []error {
			return validate()
		},
		GetClient: func(host_name string, port_number string, database_name string, database_username string) (*Client, []error) {
			lock_client.Lock()
			defer lock_client.Unlock()
			return getClient(host_name, port_number, database_name, database_username)
		},
		GetNextUserCount: func() int {
			lock_user_count.Lock()
			defer lock_user_count.Unlock()
			user_count++
			if user_count >= 100 {
				user_count = 0
			}
			return user_count
		},
	}
	setClientManager(&x)

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

