package class

import (
	"fmt"
	"strings"
	"io/ioutil"
	"bufio"
	"os"
)

func CloneClient(client *Client) *Client {
	if client == nil {
		return client
	}

	return client.Clone()
}

type Client struct {
	CreateDatabase func(database_name *string, database_create_options *DatabaseCreateOptions) (*Database, []error)
	DatabaseExists func(database_name *string) (*bool, []error)
	UseDatabase func(database *Database) []error
	UseDatabaseByName func(database_name string) (*Database, []error)
	UseDatabaseUsername func(database_username *string) []error
	CreateUser func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, []error)
	GetDatabaseUsername func() (*string)
	GetHost func() *Host 
	GetDatabase func() *Database 
	Clone func() (*Client)
	Validate func() []error
	Grant func(user *User, grant string, filter string) (*Grant, []error)
	ToJSONString func() string
}

func NewClient(host *Host, database_username *string, database *Database) (*Client, []error) {
	var this_client *Client
	
	data := Map {
		"[host]":Map{"value":CloneHost(host),"mandatory":false},
		"[database_username]":Map{"value":CloneString(database_username),"mandatory":false, 
		FILTERS(): Array{ Map {"values":GetCredentialsUsernameValidCharacters(),"function":getWhitelistCharactersFunc()}}},
		"[database]":Map{"value":CloneDatabase(database),"mandatory":false},
	}

	getHost := func() *Host {
		return CloneHost((data.M("[host]").GetObject("value").(*Host)))
	}

	getDatabaseUsername := func() *string {
		return CloneString(data.M("[database_username]").S("value"))
	}

	getDatabase := func() *Database {
		return CloneDatabase(data.M("[database]").GetObject("value").(*Database))
	}

	setDatabase := func(database *Database) {
		(data.M("[database]"))["value"] = CloneDatabase(database)
	}

	setDatabaseUsername := func(database_username *string) {
		(data.M("[database_username]"))["value"] = CloneString(database_username)
	}

	validate := func() ([]error) {
		return ValidateData(data, "Client")
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
		CreateDatabase: func(database_name *string, database_create_options *DatabaseCreateOptions) (*Database, []error) {
			database, errs := NewDatabase(getClient(), database_name, database_create_options)
			if errs != nil {
				return nil, errs
			}

			errors := database.Create()
			if errors != nil {
				return nil, errors
			}
			
			return database, nil
		},
		DatabaseExists: func(database_name *string) (*bool, []error) {
			database, database_errors := NewDatabase(getClient(), database_name, nil)
			if database_errors != nil {
				return nil, database_errors
			}

			exists, exists_errors := database.Exists()

			return exists, exists_errors
		},
		CreateUser: func(username *string, password *string, domain_name *string, options map[string]map[string][][]string) (*User, []error) {
			credentials, credentail_errors := NewCredentials(username, password)
			if credentail_errors != nil {
				return nil, credentail_errors
			}

			domain, domain_errors := NewDomainName(domain_name)
			if domain_errors != nil {
				return nil, domain_errors
			}

			user, user_errors := NewUser(getClient(), credentials, domain, options)
			if user_errors != nil {
				return nil, user_errors
			}

			create_errors := user.Create()
			if create_errors != nil {
				return nil, create_errors
			}

			return user, nil
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
		UseDatabaseByName: func(database_name string) (*Database, []error) {			
			database, database_errors := NewDatabase(getClient(), &database_name, nil)
			if database_errors != nil {
				return nil, database_errors
			}

			setDatabase(database)
			database.SetClient(this_client)
			return database, nil
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
		Grant: func(user *User, grant string, filter string) (*Grant, []error) {
			client := getClient()
			grant_obj, grant_errors := NewGrant(client, user, &grant, &filter)

			if grant_errors != nil {
				return nil, grant_errors
			}

			grant_errs := (*grant_obj).Grant()
			if grant_errs != nil {
				return nil, grant_errs
			}

			return grant_obj, nil
		},
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
    }
	setClient(&x)

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

func GetCredentialDetails(label string) (string, string, string, string, string, []error) {
	var errors []error

	files, err := ioutil.ReadDir("./")
    if err != nil {
		errors = append(errors, err)
		return "", "", "", "", "", errors
    }

	filename := ""
    for _, file := range files {
		if file.IsDir() {
			continue
		}

		currentFileName := file.Name()

		if !strings.HasPrefix(currentFileName, "holistic_db_config:") {
			continue
		}

		if !strings.HasSuffix(currentFileName, label + ".config") {
			continue
		}		
		filename = currentFileName
    }

	if filename == "" {
		errors = append(errors, fmt.Errorf("database config for %s not found ust be in the format: holistic_db_config|{database_ip_address}|{database_port_number}|{database_name}|{database_username}.config e.g holistic_db_config|127.0.0.1|3306|holistic|root.config", label))
		return "", "", "", "", "", errors
	}

	parts := strings.Split(filename, ":")
	if len(parts) != 5 {
		errors = append(errors, fmt.Errorf("database config for %s not found ust be in the format: holistic_db_config|{database_ip_address}|{database_port_number}|{database_name}|{database_username}.config e.g holistic_db_config|127.0.0.1|3306|holistic|root.config", label))
		return "", "", "", "", "", errors
	}

	ip_address := parts[1]
	port_number := parts[2]
	database_name := parts[3]

	password := ""
	username := ""
	
	file, err_file := os.Open(filename)

    if err_file != nil {
        errors = append(errors, err_file)
		return "", "", "", "", "", errors
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		currentText := scanner.Text()
		if strings.HasPrefix(currentText, "password=") {
			password = currentText[9:len(currentText)]
		}

		if strings.HasPrefix(currentText, "user=") {
			username = currentText[5:len(currentText)]
		}
    }

    if file_errs := scanner.Err(); err != nil {
        errors = append(errors, file_errs)
    }

	if password == "" {
		errors = append(errors, fmt.Errorf("password not found for file: %s", filename))
	}

	if username == "" {
		errors = append(errors, fmt.Errorf("user not found for file: %s", filename))
	}

	if len(errors) > 0 {
		return "", "", "", "", "", errors
	}

	return ip_address, port_number, database_name, username, password, errors
}

