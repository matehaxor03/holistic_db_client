package integration
 
import (
    "testing"
	"time"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)
 
func TestDatabaseCreate(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)

    database_errors := database.Create()
	if database_errors != nil {
		t.Error(database_errors)
	}
}

func TestDatabaseDelete(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)

    database.Create()
	database_errors := database.Delete()
	if database_errors != nil {
		t.Error(database_errors)
	}
}
