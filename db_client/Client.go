package db_client

import (
	"fmt"
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	dao "github.com/matehaxor03/holistic_db_client/dao"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

type Client struct {
	CreateDatabase      func(database_name string, character_set *string, collate *string) (*dao.Database, []error)
	GetDatabaseInterface func(database_name string, character_set *string, collate *string) (*dao.Database, []error)
	DeleteDatabase      func(database_name string) []error
	DatabaseExists      func(database_name string) (*bool, []error)
	UseDatabase         func(database dao.Database) []error
	UseDatabaseByName   func(database_name string) ([]error)
	UseDatabaseUsername func(database_username string) []error
	GetUser             func(username string) (*dao.User, []error)
	CreateUser          func(username string, password string, domain_name string) (*dao.User, []error)
	UserExists          func(username string) (*bool, []error)
	GetDatabaseUsername func() (*string, []error)
	GetHost             func() (*dao.Host, []error)
	GetDatabase         func() (*dao.Database, []error)
	GetClientManager    func() (*ClientManager, []error)
	Validate            func() []error
	ToJSONString        func(json *strings.Builder) []error
	Grant            func(user dao.User, grant string, database_filter *string, table_filter *string) (*dao.Grant, []error)
}

func newClient(client_manager ClientManager, host *dao.Host, database_username *string, database *dao.Database, database_reserved_words_obj *validation_constants.DatabaseReservedWords, database_name_whitelist_characters_obj *validation_constants.DatabaseNameCharacterWhitelist, table_name_whitelist_characters_obj *validation_constants.TableNameCharacterWhitelist, column_name_whitelist_characters_obj *validation_constants.ColumnNameCharacterWhitelist) (*Client, []error) {
	var errors []error
	var this_client *Client
	struct_type := "*db_client.Client"

	setClient := func(client *Client) {
		this_client = client
	}

	getClient := func() *Client {
		return this_client
	}

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[client_manager]", client_manager)
	map_system_fields.SetObjectForMap("[host]", host)
	map_system_fields.SetObjectForMap("[database]", database)
	
	map_system_fields.SetObjectForMap("[database_username]", database_username)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()
	
	map_client_manager := json.NewMapValue()
	map_client_manager.SetStringValue("type", "db_client.ClientManager")
	map_system_schema.SetMapValue("[client_manager]", map_client_manager)

	map_host := json.NewMapValue()
	map_host.SetStringValue("type", "*dao.Host")
	map_system_schema.SetMapValue("[host]", map_host)

	map_database := json.NewMapValue()
	map_database.SetStringValue("type", "*dao.Database")
	map_system_schema.SetMapValue("[database]", map_database)

	map_database_username := json.NewMapValue()
	map_database_username.SetStringValue("type", "*string")
	array_database_username_filters := json.NewArrayValue()
	map_database_username_filter := json.NewMapValue()
	map_database_username_filter.SetObjectForMap("values", validation_constants.GetValidUsernameCharacters())
	map_database_username_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
	array_database_username_filters.AppendMapValue(map_database_username_filter)
	map_database_username.SetArrayValue("filters", array_database_username_filters)
	map_system_schema.SetMapValue("[database_username]", map_database_username)
	
	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return &data
	}

	getHost := func() (*dao.Host, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[host]", "*dao.Host")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*dao.Host), nil
	}

	getDatabaseUsername := func() (*string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*string), nil
	}

	getDatabaseInterface := func(database_name string, character_set *string, collate *string) (*dao.Database, []error) {
		var errors []error
		temp_database_create_options, temp_database_create_options_errors := dao.NewDatabaseCreateOptions(character_set, collate)
		if temp_database_create_options_errors != nil {
			return nil, temp_database_create_options_errors
		}

		temp_host, temp_host_errors := getHost()
		if temp_host_errors != nil {
			errors = append(errors, temp_host_errors...)
		} else if common.IsNil(temp_host) {
			errors = append(errors, fmt.Errorf("host is nil"))
		}

		temp_database_username, temp_database_username_errors := getDatabaseUsername()
		if temp_database_username_errors != nil {
			errors = append(errors, temp_database_username_errors...)
		} else if common.IsNil(temp_database_username) {
			errors = append(errors, fmt.Errorf("database username is nil"))
		} else if *temp_database_username == "" {
			errors = append(errors, fmt.Errorf("database username is an empty string"))
		}

		if len(errors) > 0 {
			return nil, errors
		}
		
		database, errs := dao.NewDatabase(*temp_host, *temp_database_username, database_name, temp_database_create_options, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if errs != nil {
			return nil, errs
		}

		return database, nil
	}

	getDatabase := func() (*dao.Database, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[database]", "*dao.Database")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*dao.Database), nil
	}

	getClientManager := func() (*ClientManager, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[client_manager]", "db_client.ClientManager")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if common.IsNil(temp_value) {
			return nil, nil
		}
		return temp_value.(*ClientManager), nil
	}

	setDatabase := func(database *dao.Database) []error {
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database]", database)
	}

	setDatabaseUsername := func(database_username string) []error {
		return helper.SetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[database_username]", database_username)
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	useDatabaseByName := func(database_name string) ([]error) {
		var errors [] error
		temp_host, temp_host_errors := getHost()
		if temp_host_errors != nil {
			errors = append(errors, temp_host_errors...)
		} else if common.IsNil(temp_host) {
			errors = append(errors, fmt.Errorf("host is nil"))
		}

		temp_database_username, temp_database_username_errors := getDatabaseUsername()
		if temp_database_username_errors != nil {
			errors = append(errors, temp_database_username_errors...)
		} else if common.IsNil(temp_database_username) {
			errors = append(errors, fmt.Errorf("database username is nil"))
		} else if *temp_database_username == "" {
			errors = append(errors, fmt.Errorf("database username is an empty string"))
		}

		if len(errors) > 0 {
			return errors
		}

		database, database_errors := dao.NewDatabase(*temp_host, *temp_database_username, database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if database_errors != nil {
			return database_errors
		}

		setDatabase(database)
		return nil
	}


	getUser := func(username string) (*dao.User, []error) {
		var errors []error

		temp_client_manager, temp_client_manager_errors := getClientManager()
		if temp_client_manager_errors != nil {
			return nil, temp_client_manager_errors
		}

		temp_host, temp_host_errors := getHost()
		if temp_host_errors != nil {
			return nil, temp_host_errors
		}

		if temp_host == nil {
			errors = append(errors, fmt.Errorf("error: Client.getHost returned nil host"))
			return nil, errors
		}

		temp_host_name, temp_host_name_errors := temp_host.GetHostName()
		if temp_host_name_errors != nil {
			return nil, temp_host_name_errors
		}

		temp_port_number, temp_port_number_errors := temp_host.GetPortNumber()
		if temp_port_number_errors != nil {
			return nil, temp_port_number_errors
		}

		temp_database, temp_database_errors := getDatabase()
		if temp_database_errors != nil {
			return nil, temp_database_errors
		}

		temp_database_name := ""
		if temp_database != nil {
			temp_database_name_value, temp_database_name_value_errors := temp_database.GetDatabaseName()
			if temp_database_name_value_errors != nil {
				return nil, temp_database_name_value_errors
			}
			temp_database_name = temp_database_name_value
		} 

		connection_string := "holistic_db_config:" + temp_host_name + ":" + temp_port_number + ":" + temp_database_name + ":" + username
		tuple_credentials, tuple_credentials_errors := temp_client_manager.GetTupleCredentials(connection_string)
		if tuple_credentials_errors != nil {
			return nil, tuple_credentials_errors
		}

		host, host_errors := dao.NewHost(*(tuple_credentials.host_name), *(tuple_credentials.port_number))
		if host_errors != nil {
			return nil, host_errors
		}

		client, client_errors := newClient(*temp_client_manager, host, tuple_credentials.database_username, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
		if client_errors != nil {
			return nil, client_errors
		}

		credentials, credentials_errors := dao.NewCredentials(*(tuple_credentials.database_username), nil)
		if credentials_errors != nil {
			return nil, credentials_errors
		}

		use_database_errors := client.UseDatabaseByName(*(tuple_credentials.database_name))
		if use_database_errors != nil {
			return nil, use_database_errors
		}

		domain_name, domain_name_errors := dao.NewDomainName(*(tuple_credentials.host_name))
		if domain_name_errors != nil {
			return nil, domain_name_errors
		}

		user, user_errors := dao.NewUser(*temp_database, *credentials, *domain_name)
		if user_errors != nil {
			return nil, user_errors
		}

		return user, nil
	}

	x := Client{
		Validate: func() []error {
			return validate()
		},
		GetDatabaseInterface: func(database_name string, character_set *string, collate *string) (*dao.Database, []error) {
			return getDatabaseInterface(database_name, character_set, collate)
		},
		CreateDatabase: func(database_name string, character_set *string, collate *string) (*dao.Database, []error) {
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
			var errors [] error
			temp_host, temp_host_errors := getHost()
			if temp_host_errors != nil {
				errors = append(errors, temp_host_errors...)
			} else if common.IsNil(temp_host) {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			temp_database_username, temp_database_username_errors := getDatabaseUsername()
			if temp_database_username_errors != nil {
				errors = append(errors, temp_database_username_errors...)
			} else if common.IsNil(temp_database_username) {
				errors = append(errors, fmt.Errorf("database username is nil"))
			} else if *temp_database_username == "" {
				errors = append(errors, fmt.Errorf("database username is an empty string"))
			}

			if len(errors) > 0 {
				return errors
			}
			
			database, database_errors := dao.NewDatabase(*temp_host, *temp_database_username, database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
			if database_errors != nil {
				return database_errors
			}

			validateion_erors := database.Delete()
			if validateion_erors != nil {
				return validateion_erors
			}

			return nil
		},
		DatabaseExists: func(database_name string) (*bool, []error) {
			var errors [] error
			temp_host, temp_host_errors := getHost()
			if temp_host_errors != nil {
				errors = append(errors, temp_host_errors...)
			} else if common.IsNil(temp_host) {
				errors = append(errors, fmt.Errorf("host is nil"))
			}

			temp_database_username, temp_database_username_errors := getDatabaseUsername()
			if temp_database_username_errors != nil {
				errors = append(errors, temp_database_username_errors...)
			} else if common.IsNil(temp_database_username) {
				errors = append(errors, fmt.Errorf("database username is nil"))
			} else if *temp_database_username == "" {
				errors = append(errors, fmt.Errorf("database username is an empty string"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			database, database_errors := dao.NewDatabase(*temp_host, *temp_database_username, database_name, nil, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj, column_name_whitelist_characters_obj)
			if database_errors != nil {
				return nil, database_errors
			}

			return database.Exists()
		},
		GetUser: func(username string) (*dao.User, []error) {
			return getUser(username)
		},
		CreateUser: func(username string, password string, domain_name string) (*dao.User, []error) {
			var errors []error
			credentials, credentail_errors := dao.NewCredentials(username, &password)
			if credentail_errors != nil {
				return nil, credentail_errors
			}

			domain, domain_errors := dao.NewDomainName(domain_name)
			if domain_errors != nil {
				return nil, domain_errors
			}

			database, database_errors := getDatabase()
			if database_errors != nil {
				return nil, database_errors
			} else if common.IsNil(database) {
				errors = append(errors, fmt.Errorf("database is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			user, user_errors := dao.NewUser(*database, *credentials, *domain)
			if user_errors != nil {
				return nil, user_errors
			}

			create_errors := user.Create()
			if create_errors != nil {
				return nil, create_errors
			}

			return user, nil
		},
		GetHost: func() (*dao.Host, []error) {
			return getHost()
		},
		GetDatabaseUsername: func() (*string, []error) {
			return getDatabaseUsername()
		},
		GetDatabase: func() (*dao.Database, []error) {
			return getDatabase()
		},
		GetClientManager: func() (*ClientManager, []error) {
			return getClientManager()
		},
		UseDatabase: func(database dao.Database) []error {
			database_errors := database.Validate()
			if database_errors != nil {
				return database_errors
			}

			setDatabase(&database)
			return nil
		},
		UseDatabaseByName: func(database_name string) ([]error) {
			return useDatabaseByName(database_name)
		},
		UseDatabaseUsername: func(database_username string) []error {
			setDatabaseUsername(database_username)
			return nil
		},
		Grant: func(user dao.User, grant string, database_filter *string, table_filter *string) (*dao.Grant, []error) {
			var errors []error
			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				errors = append(errors, temp_database_errors...)
			}

			if temp_database == nil {
				temp_database_username, temp_database_username_errors := getDatabaseUsername()
				if temp_database_username_errors != nil {
					return nil, temp_database_username_errors
				} else if common.IsNil(temp_database_username) {
					errors = append(errors, fmt.Errorf("database username is nil"))
				} else if *temp_database_username == "" {
					errors = append(errors, fmt.Errorf("database username is an empty string"))
				} else {
					useDatabaseByName(*temp_database_username)
					temp_database, temp_database_errors = getDatabase()
					if temp_database_errors != nil {
						errors = append(errors, temp_database_errors...)
					} else if common.IsNil(temp_database) {
						errors = append(errors, fmt.Errorf("database is nil"))
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}
			
			grant_obj, grant_errors := dao.NewGrant(*temp_database, user, grant, database_filter, table_filter, database_reserved_words_obj, database_name_whitelist_characters_obj, table_name_whitelist_characters_obj)

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

			_, user_errors := getUser("root") 
			if user_errors != nil {
				errors = append(errors, user_errors...)
				return nil, errors
			}

			client := getClient()

			use_database_errors := client.UseDatabaseByName("mysql")
			if use_database_errors != nil {
				errors = append(errors, use_database_errors...)
				return nil, errors
			}

			temp_database, temp_database_errors := getDatabase()
			if temp_database_errors != nil {
				errors = append(errors, temp_database_errors...)
			} else if common.IsNil(temp_database) {
				errors = append(errors, fmt.Errorf("database is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
			}

			table, table_errors := database.GetTable("user")
			if table_errors != nil {
				errors = append(errors, table_errors...)
				return nil, errors
			}

			username_escaped, username_escaped_error := common.EscapeString(username, "'")
			if username_escaped_error != nil {
				errors = append(errors, username_escaped_error)
				return nil, errors
			}

			select_fields := json.NewArray()
			select_fields.AppendStringValue("User")

			filter_fields := json.NewMap()
			filter_fields.SetStringValue("User", username_escaped)

			records, records_errors := table.ReadRecords(select_fields, filter_fields, nil, nil, nil, nil)

			if records_errors != nil {
				return nil, records_errors
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
				return nil, errors
			}

			return &exists, nil
		},
	}
	setClient(&x)

	validation_errors := validate()
	if validation_errors != nil {
		errors = append(errors, validation_errors...)
	}

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
