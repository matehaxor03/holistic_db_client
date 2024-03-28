package dao

import (
	"sync"
	validate "github.com/matehaxor03/holistic_validator/validate"
	host_client "github.com/matehaxor03/holistic_host_client/host_client"
)

type ClientManager struct {
	GetClient func(host_name string, port_number string, database_name string, database_username string) (*Client, []error)
	Validate func() []error
	GetNextUserCount func() int
	GetHostClientUser func() host_client.User
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
	lock_user_count := &sync.RWMutex{}
	lock_client := &sync.RWMutex{}
	
	getClient := func(host_name string, port_number string, database_name string, database_username string) (*Client, []error) {
		var errors []error
		host_client_instance, host_client_errors := host_client.NewHostClient()
		if host_client_errors != nil {
			return nil, host_client_errors
		}

		host_client_user, host_client_user_errors := host_client_instance.Whoami()
		if host_client_user_errors != nil {
			return nil, host_client_user_errors
		}

		host, host_errors := newHost(verify, host_name, port_number)
		
		if host_errors != nil {
			errors = append(errors, host_errors...)
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		client, client_errors := newClient(*host_client_user, verify, *getClientManager(), host, &database_username, nil)

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

