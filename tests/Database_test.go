package tests
 
import (
    "testing"
	"strings"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestDatabaseName() string {
	return "holistic_test"
}

func GetTestDatabaseCreateOptions() *class.DatabaseCreateOptions {
	character_set := class.GET_CHARACTER_SET_UTF8MB4()
	collate := class.GET_COLLATE_UTF8MB4_0900_AI_CI()
	return class.NewDatabaseCreateOptions(&character_set, &collate)
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

func getTestDatabase(t *testing.T) (*class.Database) {
	var errors []error

	client := GetTestClient(t)

	database, database_errors := class.NewDatabase(client, GetTestDatabaseName(), GetTestDatabaseCreateOptions())
	if database_errors != nil {
		errors = append(errors, errors...)
	}

	database_exists, database_exists_error := database.Exists()
	if database_exists_error != nil {
		errors = append(errors, database_exists_error...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	if *database_exists {
		database_deleted_errors := database.Delete()
		if database_deleted_errors != nil {
			errors = append(errors, database_deleted_errors...)
		}
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return database
}
 
func TestDatabaseCreate(t *testing.T) {
	database := getTestDatabase(t)

    database_errors := database.Create()
	if database_errors != nil {
		t.Error(database_errors)
	}
}

func TestDatabaseDelete(t *testing.T) {
	database := getTestDatabase(t)

    database.Create()
	database_errors := database.Delete()
	if database_errors != nil {
		t.Error(database_errors)
	}
}

func TestDatabaseExistsTrue(t *testing.T) {
	database := getTestDatabase(t)

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
	database := getTestDatabase(t)

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
	database := getTestDatabase(t)

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
	database := getTestDatabase(t)
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
	database := getTestDatabase(t)
	for blacklist_database_name := range blacklist_map {
		set_database_name_errors := database.SetDatabaseName(blacklist_database_name)
		
		if set_database_name_errors == nil {
			t.Errorf("SetDatabaseName should return error when database_name is blacklisted")
		}

		database_name := database.GetDatabaseName()
		if database_name == blacklist_database_name {
			t.Errorf("database_name was updated to the blacklisted database_name")
		}

		if database_name != GetTestDatabaseName() {
			t.Errorf("database_name is '%s' and should be '%s'", database_name,  GetTestDatabaseName())
		}
	}
}


func TestDatabaseCannotCreateWithBlackListName(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		database, create_database_errors := class.NewDatabase(client, blacklist_database_name, database_create_options)
		
		if create_database_errors == nil {
			t.Errorf("NewDatabase should return error when database_name is blacklisted")
		}

		if database != nil {
			t.Errorf("NewDatabase should be nil when database_name is blacklisted")
		}
	}
}

func TestDatabaseCannotCreateWithBlackListNameUppercase(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		database, create_database_errors := class.NewDatabase(client, strings.ToUpper(blacklist_database_name), database_create_options)
		
		if create_database_errors == nil {
			t.Errorf("NewDatabase should return error when database_name is blacklisted")
		}

		if database != nil {
			t.Errorf("NewDatabase should be nil when database_name is blacklisted")
		}
	}
}


func TestDatabaseCannotCreateWithBlackListNameLowercase(t *testing.T) {
	client := GetTestClient(t)
	database_create_options :=  GetTestDatabaseCreateOptions()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name := range blacklist_map {
		database, create_database_errors := class.NewDatabase(client, strings.ToLower(blacklist_database_name), database_create_options)
		
		if create_database_errors == nil {
			t.Errorf("NewDatabase should return error when database_name is blacklisted")
		}

		if database != nil {
			t.Errorf("NewDatabase should be nil when database_name is blacklisted")
		}
	}
}