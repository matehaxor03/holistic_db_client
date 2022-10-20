package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetDatabaseName() string {
	return "holistic_test"
}

func GetDatabaseCreateOptions() *class.DatabaseCreateOptions {
	character_set := class.GET_CHARACTER_SET_UTF8MB4()
	collate := class.GET_COLLATE_UTF8MB4_0900_AI_CI()
	return class.NewDatabaseCreateOptions(&character_set, &collate)
}

func getTestDatabase(t *testing.T) (*class.Database) {
	var errors []error
	host_value := "127.0.0.1"
	port_value := "3306"
	user_value := "root"

	host, host_errors := class.NewHost(host_value, port_value)
	if host_errors != nil {
		errors = append(errors, host_errors...)
	}

	client, client_errors := class.NewClient(host, &user_value, nil)
	if client_errors != nil {
		errors = append(errors, client_errors...)
	}

	database, database_errors := class.NewDatabase(client, GetDatabaseName(), GetDatabaseCreateOptions())
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
 
func TestCanCreateDatabase(t *testing.T) {
	database := getTestDatabase(t)
	//fmt.Println(client.ToJSONString())

    database_errors := database.Create()
	if database_errors != nil {
		t.Error(database_errors)
	}
}