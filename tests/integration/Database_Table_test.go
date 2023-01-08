package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestDatabaseCanCheckTableExistsTrue(t *testing.T) {
	database := helper.GetTestDatabaseCreated(t)
	table_name := helper.GetTestTableName()
	database.CreateTable(table_name, (helper.GetTestSchema()))
	exists, exists_errors := database.TableExists(table_name)
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if (!exists) {
		t.Errorf("error: exists is 'false' when it should be 'true'")
	}
}

func TestDatabaseCanCheckTableExistsFalse(t *testing.T) {
	database := helper.GetTestDatabaseCreated(t)
	exists, exists_errors := database.TableExists("ThisTableDoesNotExist")
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if (exists) {
		t.Errorf("error: exists is 'true' when it should be 'false'")
	}
}