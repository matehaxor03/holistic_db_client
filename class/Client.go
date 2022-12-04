package class

import (
	//"bufio"
	"fmt"
	//"io/ioutil"
	//"os"
	"strings"
)

type Client struct {
	CreateDatabase      func(database_name string, character_set *string, collate *string) (*Database, []error)
	GetDatabaseInterface func(database_name string, character_set *string, collate *string) (*Database, []error)
	DeleteDatabase      func(database_name string) []error
	DatabaseExists      func(database_name string) (*bool, []error)
	UseDatabase         func(database *Database) []error
	UseDatabaseByName   func(database_name string) (*Database, []error)
	UseDatabaseUsername func(database_username string) []error
	GetUser             func(username string) (*User, []error)
	CreateUser          func(username string, password string, domain_name string) (*User, []error)
	UserExists          func(username string) (*bool, []error)
	GetDatabaseUsername func() (*string, []error)
	GetHost             func() (*Host, []error)
	GetDatabase         func() (*Database, []error)
	Validate            func() []error
	ToJSONString        func(json *strings.Builder) []error
	Grant               func(user *User, grant string, database_filter *string, table_filter *string) (*Grant, []error)
}

func newClient(client_manager *ClientManager, host *Host, database_username *string, database *Database, database_reserved_words_obj *DatabaseReservedWords, database_name_whitelist_characters_obj *DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *TableNameCharacterWhitelist, column_name_whitelist_characters_obj *ColumnNameCharacterWhitelist) (*Client, []error) {
	var this_client *Client
	struct_type := "*class.Client"

	setClient := func(client *Client) {
		this_client = client
	}

	getClient := func() *Client {
		return this_client
	}

	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]":Map{
			"[client_manager]": client_manager, "[host]": host, "[database]": database, "[database_username]": database_username },
		"[system_schema]":Map{
			"[client_manager]": Map{"type":"*class.ClientManager", "mandatory": true},
			"[host]": Map{"type":"*class.Host", "mandatory": false},
			"[database]": Map{"type":"*class.Database", "mandatory": false},
			"[database_username]": Map{"type":"*string", "mandatory":false,
				FILTERS(): Array{Map{"values": GetCredentialsUsernameValidCharacters(), "function": getWhitelistCharactersFunc()}}}},
	}

	getData := func() *Map {
		return &data
	}

	getDatabaseInterface := func(database_name string, character_set *string, collate *string) (*Database, []error) {
		temp_database_create_options, temp_database_create_options_errors := newDatabaseCreateOptions(character_set, collate)
		if temp_database_create_options_errors != nil {
			return nil, temp_database_create_options_errors
		}
		
		database, errs := newDatabase(getClient(), database_name, temp_database_create_options, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if errs != nil {
			return nil, errs
		}

		return database, nil
	}

	getHost := func() (*Host, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[host]", "*class.Host")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*Host), nil
	}

	getDatabaseUsername := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*string), nil
	}

	getDatabase := func() (*Database, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[database]", "*class.Database")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*Database), nil
	}

	getClientManager := func() (*ClientManager, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client_manager]", "*class.ClientManager")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*ClientManager), nil
	}

	setDatabase := func(database *Database) []error {
		return SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database]", database)
	}

	setDatabaseUsername := func(database_username string) []error {
		temp_database_map, temp_database_map_errors := getData().GetMap("[database_username]")
		if temp_database_map_errors != nil {
			return temp_database_map_errors
		}
		temp_database_map.SetString("value", &database_username)
		return nil
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getUser := func(username string) (*User, []error) {
		temp_client_manager, temp_client_manager_errors := getClientManager()
		if temp_client_manager_errors != nil {
			return nil, temp_client_manager_errors
		}

		tuple_credentials, tuple_credentials_errors := temp_client_manager.GetTupleCredentials(username)
		if tuple_credentials_errors != nil {
			return nil, tuple_credentials_errors
		}

		host, host_errors := newHost(*(tuple_credentials.host_name), *(tuple_credentials.port_number))
		if host_errors != nil {
			return nil, host_errors
		}

		client, client_errors := newClient(temp_client_manager, host, tuple_credentials.database_username, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if client_errors != nil {
			return nil, client_errors
		}

		credentials, credentials_errors := newCredentials(*(tuple_credentials.database_username), nil)
		if credentials_errors != nil {
			return nil, credentials_errors
		}

		_, use_database_errors := client.UseDatabaseByName(*(tuple_credentials.database_name))
		if use_database_errors != nil {
			return nil, use_database_errors
		}

		domain_name, domain_name_errors := NewDomainName(*(tuple_credentials.host_name))
		if domain_name_errors != nil {
			return nil, domain_name_errors
		}

		user, user_errors := newUser(client, credentials, domain_name)
		if user_errors != nil {
			return nil, user_errors
		}

		return user, nil
	}

	x := Client{
		Validate: func() []error {
			return validate()
		},
		GetDatabaseInterface: func(database_name string, character_set *string, collate *string) (*Database, []error) {
			return getDatabaseInterface(database_name, character_set, collate)
		},
		CreateDatabase: func(database_name string, character_set *string, collate *string) (*Database, []error) {
			database, errs := getDatabaseInterface(database_name, character_set, collate)
			if errs != nil {
				return nil, errs
			}

			errors := database.Create()
			if errors != nil {
				return nil, errors
			}

			return database, nil
		},
		DeleteDatabase: func(database_name string) []error {
			database, database_errors := newDatabase(getClient(), database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
			if database_errors != nil {
				return database_errors
			}

			errors := database.Delete()
			if errors != nil {
				return errors
			}

			return nil
		},
		DatabaseExists: func(database_name string) (*bool, []error) {
			database, database_errors := newDatabase(getClient(), database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
			if database_errors != nil {
				return nil, database_errors
			}

			exists, exists_errors := database.Exists()

			return exists, exists_errors
		},
		GetUser: func(username string) (*User, []error) {
			return getUser(username)
		},
		CreateUser: func(username string, password string, domain_name string) (*User, []error) {
			credentials, credentail_errors := newCredentials(username, &password)
			if credentail_errors != nil {
				return nil, credentail_errors
			}

			domain, domain_errors := NewDomainName(domain_name)
			if domain_errors != nil {
				return nil, domain_errors
			}

			user, user_errors := newUser(getClient(), credentials, domain)
			if user_errors != nil {
				return nil, user_errors
			}

			create_errors := user.Create()
			if create_errors != nil {
				return nil, create_errors
			}

			return user, nil
		},
		GetHost: func() (*Host, []error) {
			return getHost()
		},
		GetDatabaseUsername: func() (*string, []error) {
			return getDatabaseUsername()
		},
		GetDatabase: func() (*Database, []error) {
			return getDatabase()
		},
		UseDatabase: func(database *Database) []error {
			var errors []error
			if database == nil {
				errors = append(errors, fmt.Errorf("database is nil"))
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
			database, database_errors := newDatabase(getClient(), database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
			if database_errors != nil {
				return nil, database_errors
			}

			setDatabase(database)
			database.SetClient(this_client)
			return database, nil
		},
		UseDatabaseUsername: func(database_username string) []error {
			setDatabaseUsername(database_username)
			return nil
		},
		Grant: func(user *User, grant string, database_filter *string, table_filter *string) (*Grant, []error) {
			client := getClient()
			grant_obj, grant_errors := newGrant(client, user, grant, database_filter, table_filter, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj)

			if grant_errors != nil {
				return nil, grant_errors
			}

			grant_errs := (*grant_obj).Grant()
			if grant_errs != nil {
				return nil, grant_errs
			}

			return grant_obj, nil
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return data.ToJSONString(json)
		},
		UserExists: func(username string) (*bool, []error) {
			errors := validate()
			if len(errors) > 0 {
				return nil, errors
			}

			user, user_errors := getUser("root") 
			if user_errors != nil {
				errors = append(errors, user_errors...)
				return nil, errors
			}

			client, client_errors := user.GetClient()
			if client_errors != nil {
				return nil, client_errors
			}

			database, use_database_errors := client.UseDatabaseByName("mysql")
			if use_database_errors != nil {
				errors = append(errors, use_database_errors...)
				return nil, errors
			}

			table, table_errors := database.GetTable("user")
			if table_errors != nil {
				errors = append(errors, table_errors...)
				return nil, errors
			}

			records, records_errors := table.ReadRecords(Map{"User": EscapeString(username)}, nil, nil)

			if records_errors != nil {
				return nil, records_errors
			}

			var exists bool
			if len(*records) == 0 {
				exists = false
			} else if (len(*records) == 1) {
				exists = true
			} else {
				errors = append(errors, fmt.Errorf("User: Exists: %d records found with username %s", len(*records), EscapeString(username)))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			return &exists, nil
		},
	}
	setClient(&x)

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
