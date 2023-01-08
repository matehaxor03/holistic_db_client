package integration_test_helpers

import (
    "testing"
	"fmt"
	"sync"
	common "github.com/matehaxor03/holistic_common/common"
	dao "github.com/matehaxor03/holistic_db_client/dao"
)

var database_count uint64 = 0
var lock_get_client_manager = &sync.Mutex{}
var lock_get_database = &sync.Mutex{}
var lock_get_database_name = &sync.Mutex{}
var client_manager *dao.ClientManager

func getTestDatabaseName() string {
	lock_get_database_name.Lock()
	defer lock_get_database_name.Unlock()
	database_count++
	return "holistic" + fmt.Sprintf("%d", database_count)
}

func EnsureDatabaseIsDeleted(t *testing.T, database dao.Database) {
	database_delete_errors := database.DeleteIfExists()
	
	if database_delete_errors != nil {
		t.Error(database_delete_errors)
		t.FailNow()
		return
	}
}

func GetTestClient(t *testing.T) (dao.Client) {
	lock_get_client_manager.Lock()
	defer lock_get_client_manager.Unlock()
	var errors []error
	if common.IsNil(client_manager) {
		temp_client_manager, temp_client_manager_errors := dao.NewClientManager()
		if temp_client_manager_errors != nil {
			errors = append(errors, temp_client_manager_errors...)
		} else if temp_client_manager == nil {
			errors = append(errors, fmt.Errorf("error: client_manager is nil"))
		}

		if len(errors) > 0 {
			t.Error(errors)
			t.FailNow()
			return dao.Client{}
		}
		client_manager = temp_client_manager
	}

	client, client_errors := client_manager.GetClient("holistic_db_config#127.0.0.1#3306#holistic_test#root")
	if client_errors != nil {
		errors = append(errors, client_errors...)
	} else if common.IsNil(client) {
		errors = append(errors, fmt.Errorf("client is nil"))
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return dao.Client{}
	}

	return *client
}

func GetTestDatabase(t *testing.T) (dao.Database) {
	lock_get_database.Lock()
	defer lock_get_database.Unlock()
	var errors []error
	client := GetTestClient(t)

	character_set := dao.GET_CHARACTER_SET_UTF8MB4()
	collate := dao.GET_COLLATE_UTF8MB4_0900_AI_CI()
	test_database_name := getTestDatabaseName()
	database, database_errors := client.GetDatabaseInterface(test_database_name, &character_set, &collate)
	if database_errors != nil {
		errors = append(errors, database_errors...)
	} else if common.IsNil(database) {
		errors = append(errors, fmt.Errorf("database is nil - database interface"))
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return  dao.Database{}
	}

	use_database_errors := client.UseDatabaseByName(test_database_name)
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return dao.Database{}
	}

	return *database
}

func GetTestDatabaseDeleted(t *testing.T) (dao.Database) {
	var errors []error

	database := GetTestDatabase(t)

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return  dao.Database{}
	}

	return database
}

func GetTestDatabaseCreated(t *testing.T) (dao.Database) {
	var errors []error

	database := GetTestDatabase(t)

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return  dao.Database{}
	}

	database_delete_errors := database.DeleteIfExists()
	if database_delete_errors != nil {
		errors = append(errors, database_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return  dao.Database{}
	}

	database_create_errors := database.Create()
	if database_create_errors != nil {
		errors = append(errors, database_create_errors...)	
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return dao.Database{}
	}

	return database
}