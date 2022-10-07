package class

import (
	"fmt"
)

func CloneClient(client *Client) *Client {
	if client == nil {
		return client
	}

	return client.Clone()
}

type Client struct {
	CreateDatabase func(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error)
	UseDatabase func(database *Database) []error
	UseDatabaseUsername func(database_username *string) []error
	CreateUser func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, []error)
	GetDatabaseUsername func() (*string)
	GetHost func() *Host 
	GetDatabase func() *Database 
	Clone func() (*Client)
	Validate func() []error
	Grant func(user *User, grant string, filter string) (*Grant, *string, []error)
}

func NewClient(host *Host, database_username *string, database *Database) (*Client, []error) {
	var this_client *Client
	
	data := Map {
		"host":Map{"value":CloneHost(host),"mandatory":false},
		"database_username":Map{"value":CloneString(database_username),"mandatory":false},
		"database":Map{"value":CloneDatabase(database),"mandatory":false},
	}

	getHost := func() *Host {
		return CloneHost((data.M("host").GetObject("value").(*Host)))
	}

	getDatabaseUsername := func() *string {
		return CloneString(data.M("database_username").S("value"))
	}

	getDatabase := func() *Database {
		return CloneDatabase(data.M("database").GetObject("value").(*Database))
	}

	setDatabase := func(database *Database) {
		(data.M("database"))["value"] = CloneDatabase(database)
	}

	setDatabaseUsername := func(database_username *string) {
		(data.M("database_username"))["value"] = CloneString(database_username)
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "Client")
	}

	setClient := func(client *Client) {
		this_client = client
	}

	getClient := func() *Client {
		return this_client
	}
			    
	x := Client{
		Validate: func() ([]error) {
			return validate()
		},
		Clone: func() (*Client) {
			cloned, _ := NewClient(getHost(), getDatabaseUsername(), getDatabase())
			return cloned
		},
		CreateDatabase: func(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error) {
			database, errs := NewDatabase(getClient(), database_name, database_create_options, options)
			if errs != nil {
				return nil, nil, errs
			}

			stdout, errors := database.Create()
			if errors != nil {
				return nil, stdout, errors
			}
			
			return database, stdout, nil
		},
		CreateUser: func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, []error) {
			credentials, credentail_errors := NewCredentials(username, password)
			if credentail_errors != nil {
				return nil, nil, credentail_errors
			}

			domain, domain_errors := NewDomainName(domain_name)
			if domain_errors != nil {
				return nil, nil, domain_errors
			}

			user, user_errors := NewUser(getClient(), credentials, domain, options)
			if user_errors != nil {
				return nil, nil, user_errors
			}

			sql_output, create_errors := user.Create()
			if create_errors != nil {
				return nil, sql_output, create_errors
			}

			return user, sql_output, nil
		},
		GetHost: func() *Host {
			return getHost()
		},
		GetDatabaseUsername: func() *string {
			return getDatabaseUsername()
		},
		GetDatabase: func() *Database {
			return getDatabase()
		},
		UseDatabase: func(database *Database) []error {
			var errors []error
			if database == nil {
				errors = append(errors,  fmt.Errorf("database is nil"))
				return errors
			}

			database_errors := database.Validate()
			if database_errors != nil {
				return database_errors
			}
			
			setDatabase(database)
			database.SetClient(this_client)
			return nil
		},
		UseDatabaseUsername: func(database_username *string) []error {
			var errors []error
			if database_username == nil {
				errors = append(errors,  fmt.Errorf("database_username is nil"))
				return errors
			}

			setDatabaseUsername(database_username)
			return nil
		},
		Grant: func(user *User, grant string, filter string) (*Grant, *string, []error) {
			client := getClient()
			grant_obj, grant_errors := NewGrant(client, user, &grant, &filter)

			if grant_errors != nil {
				return nil, nil, grant_errors
			}

			stdout, grant_errs := (*grant_obj).Grant()
			if grant_errs != nil {
				return nil, stdout, grant_errs
			}

			return grant_obj, stdout, nil
		},
    }
	setClient(&x)

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
