package class

func CloneClient(client *Client) *Client {
	if client == nil {
		return client
	}

	return client.Clone()
}

type Client struct {
	CreateDatabase func(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error)
	UseDatabase func(database_name *string) []error
	CreateUser func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, []error)
	GetCredentials func() (*Credentials)
	GetHost func() *Host 
	Clone func() (*Client)
}

func NewClient(host *Host, credentials *Credentials, database *Database) (*Client, []error) {
	data := Map {
		"host":Map{"type":"*Host","value":CloneHost(host),"mandatory":false},
		"credentials":Map{"type":"*Credentials","value":CloneCredentials(credentials),"mandatory":false},
		"database":Map{"type":"*Database","value":CloneDatabase(database),"mandatory":false},
	}

	getHost := func() *Host {
		return CloneHost((data.M("host").GetObject("value").(*Host)))
	}

	getCredentials := func() *Credentials {
		return CloneCredentials(data.M("credentials").GetObject("value").(*Credentials))
	}

	getDatabase := func() *Database {
		return CloneDatabase(data.M("database").GetObject("value").(*Database))
	}

	setDatabase := func(database *Database) {
		data.M("database")["value"] = database
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "Client")
	}
			    
	x := Client{
		Clone: func() (*Client) {
			cloned, _ := NewClient(host, credentials, database)
			return cloned
		},
		CreateDatabase: func(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error) {
			database, errs := NewDatabase(getHost(), getCredentials(), database_name, database_create_options, options)
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

			user, user_errors := NewUser(getHost(), getCredentials(), getDatabase(), credentials, domain, options)
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
		GetCredentials: func() *Credentials {
			return getCredentials()
		},
		UseDatabase: func(database_name *string) []error {
			database, database_errs := NewDatabase(getHost(), getCredentials(), database_name, nil, nil)
			if database_errs != nil {
				return database_errs
			}
			setDatabase(database)
			return nil
		},
    }

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
