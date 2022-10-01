package class

/*
import (
	"fmt"
	"reflect"
)*/


func CloneClient(client *Client) *Client {
	if client == nil {
		return client
	}

	return client.Clone()
}

type Client struct {
	CreateDatabase func(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error)
	CreateUser func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, *string, []error)
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
			
			return database.Create()
		},
		CreateUser: func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, *string, []error) {
			credentials, credentail_errors := NewCredentials(username, password)
			if credentail_errors != nil {
				return nil, nil, nil, credentail_errors
			}

			domain, domain_errors := NewDomainName(domain_name)
			if domain_errors != nil {
				return nil, nil, nil, domain_errors
			}

			user, user_errors := NewUser(getHost(), getCredentials(), credentials, domain, options)
			if user_errors != nil {
				return nil, nil, nil, user_errors
			}

			sql_output, sql_output_errs, create_errors := user.Create()
			if create_errors != nil || *sql_output_errs != "" {
				return nil, sql_output, sql_output_errs, create_errors
			}

			return user, sql_output, nil, nil
		},
		GetHost: func() *Host {
			return getHost()
		},
		GetCredentials: func() *Credentials {
			return getCredentials()
		},
    }

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

/*
func (this *Client) CreateDatabase(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error) {
	database, errs := NewDatabase((*this).host, (*this).credentials, database_name, database_create_options, options)
	if errs != nil {
		return nil, nil, errs
	}
	
	return database.Create()
}
*/

/*
func (this *Client) CreateUser(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, []error) {
	//var errors []error
	
	//todo check validations
	//credentials := NewCredentials(username, password)
	// todo check validations
	//domain := NewDomainName(domain_name)
	return nil , nil, nil//user.Create()
}


func (this *Client) GetHost() (*Host) {
	return (*this).host
}

func (this *Client) GetCredentials() (*Credentials) {
	return (*this).credentials
}*/
