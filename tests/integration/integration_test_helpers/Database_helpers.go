package integration_test_helpers

import (
    "testing"
	//"strings"
	"fmt"
	"sync"
	//json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	class "github.com/matehaxor03/holistic_db_client/class"
)

var lock_get_client_manager = &sync.Mutex{}
var client_manager *class.ClientManager

func GetTestDatabaseName() string {
	return "holistic_test"
}

func EnsureDatabaseIsDeleted(t *testing.T, database *class.Database) {
	database_delete_errors := database.DeleteIfExists()
	
	if database_delete_errors != nil {
		t.Error(database_delete_errors)
		t.FailNow()
		return
	}
}

func GetTestClient(t *testing.T) (*class.Client) {
	lock_get_client_manager.Lock()
	defer lock_get_client_manager.Unlock()
	var errors []error
	if common.IsNil(client_manager) {
		temp_client_manager, temp_client_manager_errors := class.NewClientManager()
		if temp_client_manager_errors != nil {
			errors = append(errors, temp_client_manager_errors...)
		} else if temp_client_manager == nil {
			errors = append(errors, fmt.Errorf("error: client_manager is nil"))
		}

		if len(errors) > 0 {
			t.Error(errors)
			t.FailNow()
			return nil
		}
		client_manager = temp_client_manager
	}

	client, client_errors := client_manager.GetClient("holistic_db_config#127.0.0.1#3306#holistic_test#root")
	if client_errors != nil {
		errors = append(errors, client_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return client
}

func GetTestDatabase(t *testing.T) (*class.Database) {
	var errors []error

	client := GetTestClient(t)

	if client == nil {
		t.Error(fmt.Errorf("error: test client is nil"))
		t.FailNow()
		return nil
	}

	character_set := class.GET_CHARACTER_SET_UTF8MB4()
	collate := class.GET_COLLATE_UTF8MB4_0900_AI_CI()
	database, database_errors := client.GetDatabaseInterface(GetTestDatabaseName(), &character_set, &collate)
	if database_errors != nil {
		errors = append(errors, database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	user_database_errors := database.UseDatabase()
	if user_database_errors != nil {
		errors = append(errors, user_database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return database
}

func GetTestDatabaseDeleted(t *testing.T) (*class.Database) {
	var errors []error

	database := GetTestDatabase(t)

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return database
}

func GetTestDatabaseCreated(t *testing.T) (*class.Database) {
	var errors []error

	database := GetTestDatabase(t)

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	database_create_errors := database.Create()
	if database_create_errors != nil {
		errors = append(errors, database_create_errors...)	
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return database
}