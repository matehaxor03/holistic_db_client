package class

import (
	"fmt"
	"reflect"
)



type Client struct {
	host *Host
	credentials *Credentials
	database *Database
}

func NewClient(host *Host, credentials *Credentials, database *Database) (*Client) {
	x := Client{host: host,
				credentials: credentials,
				database: database}
			    
	return &x
}

func (this *Client) CreateDatabase(database_name *string, database_create_options *DatabaseCreateOptions, options map[string]map[string][][]string) (*Database, *string, []error) {
	database, errs := NewDatabase((*this).host, (*this).credentials, database_name, database_create_options, options)
	if errs != nil {
		return nil, nil, errs
	}
	
	return database.Create()
}

func (this *Client) CreateUser(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, *string, []error) {
	//var errors []error
	
	//todo check validations
	//credentials := NewCredentials(username, password)
	// todo check validations
	//domain := NewDomainName(domain_name)
	return nil , nil, nil//user.Create()
}
