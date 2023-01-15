package dao

import (
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type Client struct {
	CreateDatabase      func(database_name string, character_set *string, collate *string) (*Database, []error)
	GetDatabaseInterface func(database_name string, character_set *string, collate *string) (*Database, []error)
	DeleteDatabase      func(database_name string) []error
	DatabaseExists      func(database_name string) (bool, []error)
	UseDatabase         func(database Database) []error
	SetDatabaseUsername func(database_username string) []error
	GetUser             func(username string) (*User, []error)
	CreateUser          func(username string, password string, domain_name string) (*User, []error)
	UserExists          func(username string) (bool, []error)
	GetDatabaseUsername func() (*string)
	GetHost             func() (*Host)
	GetDatabaseByName   func(database_name string) (*Database, []error)
	GetDatabase         func() (*Database)
	SetDatabase         func(database *Database)
	GetClientManager    func() (ClientManager)
	Validate            func() []error
	Grant            func(user User, grant string, database_filter *string, table_filter *string) (*Grant, []error)
	ValidateTableName func(table_name string) []error
}

func newClient(verify *validate.Validator, client_manager ClientManager, host *Host, database_username *string, database *Database) (*Client, []error) {
	var this_client *Client

	sql_command, sql_command_errors := newSQLCommand()
	if sql_command_errors != nil {
		return nil, sql_command_errors
	}

	setClient := func(client *Client) {
		this_client = client
	}

	getClient := func() *Client {
		return this_client
	}

	getHost := func() (*Host) {
		return host
	}

	getDatabaseUsername := func() (*string) {
		return database_username
	}

	getDatabaseInterface := func(database_name string, character_set *string, collate *string) (*Database, []error) {
		temp_database_create_options, temp_database_create_options_errors := newDatabaseCreateOptions(verify, character_set, collate)
		if temp_database_create_options_errors != nil {
			return nil, temp_database_create_options_errors
		}
		
		database_interface, database_interface_errors := newDatabase(verify, *this_client, *host, database_username, database_name, temp_database_create_options, *sql_command)
		if database_interface_errors != nil {
			return nil, database_interface_errors
		}

		return database_interface, nil
	}

	getDatabase := func() (*Database) {
		return database
	}

	getClientManager := func() (ClientManager) {
		return client_manager
	}

	setDatabase := func(new_database *Database) []error {
		if database_validation_errors := new_database.Validate(); database_validation_errors != nil {
			return database_validation_errors
		}
		database = new_database
		return nil
	}

	setDatabaseUsername := func(new_database_username string) []error {
		if new_database_username_errors := verify.ValidateUsername(new_database_username); new_database_username_errors != nil {
			return new_database_username_errors
		}
		database_username = &new_database_username
		return nil
	}

	validate := func() []error {
		var errors []error
		if client_manager_errors := client_manager.Validate(); client_manager_errors != nil {
			errors = append(errors, client_manager_errors...)
		}

		if host != nil {
			if host_errors := host.Validate(); host_errors != nil {
				errors = append(errors, host_errors...)
			}
		}

		if database_username != nil {
			if database_username_errors := verify.ValidateUsername(*database_username); database_username_errors != nil {
				errors = append(errors, database_username_errors...)
			}
		}

		if database != nil {
			if database_errors := database.Validate(); database_errors != nil {
				errors = append(errors, database_errors...)
			}
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getUser := func(username string) (*User, []error) {
		var errors []error

		if host == nil {
			errors = append(errors, fmt.Errorf("error: Client.getHost returned nil host"))
			return nil, errors
		}

		temp_host_name := host.GetHostName()
		temp_port_number := host.GetPortNumber()
	
		temp_database_name := ""
		if database != nil {
			temp_database_name = database.GetDatabaseName()
		} 

		new_temp_host, new_temp_host_errors := newHost(verify, temp_host_name, temp_port_number)
		if new_temp_host_errors != nil {
			return nil, new_temp_host_errors
		}

		client, client_errors := newClient(verify, getClientManager(), new_temp_host, &username, nil)
		if client_errors != nil {
			return nil, client_errors
		}

		credentials, credentials_errors := newCredentials(verify, username, "")
		if credentials_errors != nil {
			return nil, credentials_errors
		}

		get_database_by_name, get_database_by_name_errors := client.GetDatabaseByName(temp_database_name)
		if get_database_by_name_errors != nil {
			return nil, get_database_by_name_errors
		}

		domain_name, domain_name_errors := NewDomainName(verify, temp_host_name)
		if domain_name_errors != nil {
			return nil, domain_name_errors
		}

		user, user_errors := newUser(*get_database_by_name, *credentials, *domain_name)
		if user_errors != nil {
			return nil, user_errors
		}

		return user, nil
	}

	validation_errors := validate()
	if validation_errors != nil {
		return nil, validation_errors
	}

	x := Client{
		Validate: func() []error {
			return validate()
		},
		GetDatabaseInterface: func(database_name string, character_set *string, collate *string) (*Database, []error) {
			return getDatabaseInterface(database_name, character_set, collate)
		},
		CreateDatabase: func(database_name string, character_set *string, collate *string) (*Database, []error) {
			database_interface, database_interface_errors := getDatabaseInterface(database_name, character_set, collate)
			if database_interface_errors != nil {
				return nil, database_interface_errors
			}

			create_errors := database_interface.Create()
			if create_errors != nil {
				return nil, create_errors
			}

			return database_interface, nil
		},
		DeleteDatabase: func(database_name string) []error {
			var errors [] error
			if host == nil {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			if database_username == nil {
				errors = append(errors, fmt.Errorf("database_username is nil"))
			}

			if len(errors) > 0 {
				return errors
			}
			
			database, database_errors := newDatabase(verify, *this_client, *host, database_username, database_name, nil, *sql_command)
			if database_errors != nil {
				return database_errors
			}

			validateion_erors := database.Delete()
			if validateion_erors != nil {
				return validateion_erors
			}

			return nil
		},
		DatabaseExists: func(database_name string) (bool, []error) {
			var errors [] error
			if host == nil {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			if database_username == nil {
				errors = append(errors, fmt.Errorf("database_username is nil"))
			}

			if len(errors) > 0 {
				return false, errors
			}

			database, database_errors := newDatabase(verify, *this_client, *host, database_username, database_name, nil, *sql_command)
			if database_errors != nil {
				return false, database_errors
			}

			return database.Exists()
		},
		GetUser: func(username string) (*User, []error) {
			return getUser(username)
		},
		CreateUser: func(username string, password string, domain_name string) (*User, []error) {
			var errors []error
			credentials, credentail_errors := newCredentials(verify, username, password)
			if credentail_errors != nil {
				return nil, credentail_errors
			}

			domain, domain_errors := NewDomainName(verify, domain_name)
			if domain_errors != nil {
				return nil, domain_errors
			}

			if database == nil {
				errors = append(errors, fmt.Errorf("database is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			user, user_errors := newUser(*database, *credentials, *domain)
			if user_errors != nil {
				return nil, user_errors
			}

			create_errors := user.Create()
			if create_errors != nil {
				return nil, create_errors
			}

			return user, nil
		},
		GetHost: func() (*Host) {
			return getHost()
		},
		GetDatabaseUsername: func() (*string) {
			return getDatabaseUsername()
		},
		GetDatabase: func() (*Database) {
			return getDatabase()
		},
		GetClientManager: func() (ClientManager) {
			return getClientManager()
		},
		GetDatabaseByName: func(get_database_name string) (*Database, []error)  {
			var errors []error
			host := getHost()
			if host == nil {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			database_username := getDatabaseUsername()
			if database_username == nil {
				errors = append(errors, fmt.Errorf("database username is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return newDatabase(verify, *this_client, *getHost(), getDatabaseUsername(), get_database_name, nil, *sql_command)
		},
		UseDatabase: func(new_database Database) []error {
			database_errors := new_database.Validate()
			if database_errors != nil {
				return database_errors
			}

			setDatabase(&new_database)
			return nil
		},
		SetDatabaseUsername: func(database_username string) []error {
			return setDatabaseUsername(database_username)
		},
		Grant: func(user User, grant string, database_filter *string, table_filter *string) (*Grant, []error) {
			var errors []error
			temp_database := database
			if database == nil {
				errors = append(errors, fmt.Errorf("database is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}
			
			grant_obj, grant_errors := newGrant(verify, *temp_database, user, grant, database_filter, table_filter)

			if grant_errors != nil {
				return nil, grant_errors
			}

			grant_errs := (*grant_obj).Grant()
			if grant_errs != nil {
				return nil, grant_errs
			}

			return grant_obj, nil
		},
		UserExists: func(username string) (bool, []error) {
			var errors []error
			_, user_errors := getUser("root") 
			if user_errors != nil {
				fmt.Println("get root user errors")
				errors = append(errors, user_errors...)
				return false, errors
			}

			client := getClient()

			mysql_database, mysql_database_errors := client.GetDatabaseByName("mysql")
			if mysql_database_errors != nil {
				fmt.Println("GetDatabaseByName errors")
				errors = append(errors, mysql_database_errors...)
				return false, errors
			}

			table, table_errors := mysql_database.GetTable("user")
			if table_errors != nil {
				fmt.Println("GetTable errors")
				errors = append(errors, table_errors...)
				return false, errors
			}

			username_escaped, username_escaped_error := common.EscapeString(username, "'")
			if username_escaped_error != nil {
				fmt.Println("username_escaped_error errors")
				errors = append(errors, username_escaped_error)
				return false, errors
			}

			select_fields := json.NewArray()
			select_fields.AppendStringValue("User")

			filter_fields := json.NewMap()
			filter_fields.SetStringValue("User", username_escaped)

			records, records_errors := table.ReadRecords(select_fields, filter_fields, nil, nil, nil, nil)

			if records_errors != nil {
				fmt.Println("records_errors errors")
				return false, records_errors
			}

			var exists bool
			if len(*records) == 0 {
				exists = false
			} else if (len(*records) == 1) {
				exists = true
			} else {
				errors = append(errors, fmt.Errorf("error: User: Exists: %d records found with username %s", len(*records), username_escaped))
			}

			if len(errors) > 0 {
				return false, errors
			}

			return exists, nil
		},
		ValidateTableName: func(table_name string) []error {
			return verify.ValidateTableName(table_name)
		},
		SetDatabase: func(set_database *Database) {
			database = set_database
		},
	}
	setClient(&x)

	return &x, nil
}
