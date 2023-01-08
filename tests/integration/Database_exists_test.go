package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)
 
func TestDatabaseExistsTrue(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)

    database.Create()
	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if !(exists) {
		t.Errorf("error: exists is 'false' when it should be 'true'")
	}
}

func TestDatabaseExistsFalse(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)
	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if (exists) {
		t.Errorf("error: exists is 'true' when it should be 'false'")
	} 
}

func TestDatabaseCreateWithExists(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)

	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if (exists) {
		t.Errorf("error: exists is 'true' when it should be 'false'")
	} else {
		create_database_errors := database.Create()
		if create_database_errors != nil {
			t.Error(create_database_errors)
		} else {
			exists_after, exists_after_errors := database.Exists()
			if exists_after_errors != nil {
				t.Error(exists_after_errors)
			} else if !exists_after {
				t.Errorf("error: exists is 'false' when it should be 'true'")
			} 
		}
	}
}

func TestDatabaseDeleteWithExists(t *testing.T) {
	database := helper.GetTestDatabase(t)
	helper.EnsureDatabaseIsDeleted(t, database)
	database.Create()

	exists, exists_errors := database.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} else if !(exists) {
		t.Errorf("error: exists is 'false' when it should be 'true'")
	} else {
		database.Delete()

		exists, exists_errors = database.Exists()
		if exists_errors != nil {
			t.Error(exists_errors)
		} else if (exists) {
			t.Errorf("error:exists is 'true' when it should be 'false'")
		} 
	}
}

