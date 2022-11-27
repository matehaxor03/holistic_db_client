package tests
 
import (
    "testing"
	"strings"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestDatabaseName() string {
	return "holistic_test"
}

func GetTestDatabaseCreateOptions() (*class.DatabaseCreateOptions) {
	character_set := class.GET_CHARACTER_SET_UTF8MB4()
	collate := class.GET_COLLATE_UTF8MB4_0900_AI_CI()
	database_create_options, _ :=  class.NewDatabaseCreateOptions(&character_set, &collate)
	return database_create_options
}

func GetTestHost(t *testing.T) (*class.Host) {
	var errors []error
	host_value := "127.0.0.1"
	port_value := "3306"

	host, host_errors := class.NewHost(host_value, port_value)
	if host_errors != nil {
		errors = append(errors, host_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return host
}

func GetTestClient(t *testing.T) (*class.Client) {
	var errors []error
	
	user_value := "root"
	host := GetTestHost(t)

	client, client_errors := class.NewClient(host, &user_value, nil)
	if client_errors != nil {
		errors = append(errors, client_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return client
}

func GetTestDatabase(t *testing.T) (*class.Database) {
	var errors []error

	client := GetTestClient(t)

	database, database_errors := class.NewDatabase(client, GetTestDatabaseName(), GetTestDatabaseCreateOptions())
	if database_errors != nil {
		errors = append(errors, database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return database
}

func GetTestDatabaseCreated(t *testing.T) (*class.Database) {
	var errors []error

	client := GetTestClient(t)

	database, database_errors := class.NewDatabase(client, GetTestDatabaseName(), GetTestDatabaseCreateOptions())
	if database_errors != nil {
		errors = append(errors, database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	database_create_errors := database.Create()
	if database_create_errors != nil {
		errors = append(errors, database_create_errors...)	
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return database
}
 
func TestDatabaseCreate(t *testing.T) {
	database := GetTestDatabase(t)

    database_errors := database.Create()
	if database_errors != nil {
		t.Error(database_errors)
	}
}

func TestDatabaseDelete(t *testing.T) {
	database := GetTestDatabase(t)

    database.Create()
	database_errors := database.Delete()
	if database_errors != nil {
		t.Error(database_errors)
	}
}

func TestDatabaseExistsTrue(t *testing.T) {
	database := GetTestDatabase(t)

    database.Create()
	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 
}

func TestDatabaseExistsFalse(t *testing.T) {
	database := GetTestDatabase(t)

	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 
}

func TestDatabaseCreateWithExists(t *testing.T) {
	database := GetTestDatabase(t)

	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 

    database_errors := database.Create()
	if database_errors != nil {
		t.Error(database_errors)
	}

	exists, exists_errors = database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 
}

func TestDatabaseDeleteWithExists(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()

	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 

    database.Delete()

	exists, exists_errors = database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 
}

func TestDatabaseCannotSetDatabaseNameWithBlackListName(t *testing.T) {
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()
	database := GetTestDatabase(t)
	for blacklist_database_name := range blacklist_map {
		t.Run(blacklist_database_name, func(t *testing.T) {
			t.Parallel()
			set_database_name_errors := database.SetDatabaseName(blacklist_database_name)
			
			if set_database_name_errors == nil {
				t.Errorf("SetDatabaseName should return error when database_name is blacklisted")
			}

			database_name, database_name_errors := database.GetDatabaseName()
			if database_name_errors != nil {
				t.Errorf(fmt.Sprintf("%s", database_name_errors))
			}

			if database_name == blacklist_database_name {
				t.Errorf("database_name was updated to the blacklisted database_name")
			}

			if database_name != GetTestDatabaseName() {
				t.Errorf("database_name is '%s' and should be '%s'", database_name,  GetTestDatabaseName())
			}
		})
	}
}


func TestDatabaseCannotCreateWithBlackListName(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		t.Run(blacklist_database_name, func(t *testing.T) {
			t.Parallel()
			database, create_database_errors := class.NewDatabase(client, blacklist_database_name, database_create_options)
			
			if create_database_errors == nil {
				t.Errorf("NewDatabase should return error when database_name is blacklisted")
			}

			if database != nil {
				t.Errorf("NewDatabase should be nil when database_name is blacklisted")
			}
		})
	}
}

func TestDatabaseCannotCreateWithBlackListNameUppercase(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		t.Run(blacklist_database_name, func(t *testing.T) {
			t.Parallel()
			database, create_database_errors := class.NewDatabase(client, strings.ToUpper(blacklist_database_name), database_create_options)
			
			if create_database_errors == nil {
				t.Errorf("NewDatabase should return error when database_name is blacklisted")
			}

			if database != nil {
				t.Errorf("NewDatabase should be nil when database_name is blacklisted")
			}
		})
	}
}


func TestDatabaseCannotCreateWithBlackListNameLowercase(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		t.Run(blacklist_database_name, func(t *testing.T) {
			t.Parallel()
			database, create_database_errors := class.NewDatabase(client, strings.ToLower(blacklist_database_name), database_create_options)
			
			if create_database_errors == nil {
				t.Errorf("NewDatabase should return error when database_name is blacklisted")
			}

			if database != nil {
				t.Errorf("NewDatabase should be nil when database_name is blacklisted")
			}
		})
	}
}

func TestDatabaseCanCreateWithWhiteListCharacters(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	whitelist_map := class.GetDatabaseNameWhitelistCharacters()

	for whitelist_database_character := range whitelist_map {
		database, new_database_errors := class.NewDatabase(client, "a" + whitelist_database_character + "a", database_create_options)
			
		if new_database_errors != nil {
			t.Errorf("NewDatabase should not return error when database_name character is whitelisted: %s", whitelist_database_character)
		} else if database == nil {
			t.Errorf("NewDatabase should not be nil when database_name is whitelisted: %s", whitelist_database_character)
		} else {
			database_delete_errors := database.DeleteIfExists()
			if database_delete_errors != nil {
				t.Error(database_delete_errors)
			} else {
				create_database_errors := database.Create()
				if create_database_errors != nil {
					t.Error(create_database_errors)
				}
			}
		}
	}
}

func TestDatabaseCannotCreateWithNonWhiteListCharacters(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	non_whitelist_map := class.Map{"(":nil, ")":nil}

	for non_whitelist_characters := range non_whitelist_map {
		t.Run(non_whitelist_characters, func(t *testing.T) {
			t.Parallel()
			database, new_database_errors := class.NewDatabase(client, "a" + non_whitelist_characters + "a", database_create_options)
			
			if new_database_errors == nil {
				t.Errorf("NewDatabase should return error when database_name character is non-whitelisted: %s", non_whitelist_characters)
			} else if database != nil {
				t.Errorf("NewDatabase should be nil when database_name is non-whitelisted: %s", non_whitelist_characters)
			}
		})
	}
}

func TestDatabaseCannotCreateWithWhiteListCharactersIfDatabaseNameLength1(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	whitelist_map := class.GetDatabaseNameWhitelistCharacters()

	for whitelist_database_character := range whitelist_map {
		t.Run(whitelist_database_character, func(t *testing.T) {
			t.Parallel()
			database, new_database_errors := class.NewDatabase(client, whitelist_database_character, database_create_options)
			
			if new_database_errors == nil {
				t.Errorf("NewDatabase should return error when database_name character is whitelisted and database_name is only one character long: %s", whitelist_database_character)
			} else if database != nil {
				t.Errorf("NewDatabase should be nil when database_name is whitelisted and is only one character long: %s", whitelist_database_character)
			}
		})

	}
}

func TestDatabaseCanGetTableNames(t *testing.T) {
	database := GetTestDatabaseCreated(t)
	database.CreateTable("some_table", GetTestSchema())

	table_names, tables_name_errors := database.GetTableNames()
	if tables_name_errors != nil {
		t.Error(tables_name_errors)
	}

	if table_names == nil {
		t.Errorf("table_names should not be nil")
	} else if !(len(*table_names) >= 0) {
		t.Errorf("database.GetTables should return at least one table name")

		if !class.Contains(*table_names, "some_table") {
			t.Errorf("some_table not found in table_names")
		}
	}
}

func TestDatabaseCanGetTables(t *testing.T) {
	database := GetTestDatabaseCreated(t)
	database.CreateTable("some_table", GetTestSchema())

	tables, tables_errors := database.GetTables()
	if tables_errors != nil {
		t.Error(tables_errors)
	}

	if tables == nil {
		t.Errorf("tables should not be nil")
	} else if !(len(*tables) >= 0) {
		t.Errorf("database.GetTables should return at least one table")
	}
}