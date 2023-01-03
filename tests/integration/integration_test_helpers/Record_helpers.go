package integration_test_helpers

import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	db_client "github.com/matehaxor03/holistic_db_client/db_client"
)

func GetTestTableCreated(t *testing.T) (*db_client.Table) {
	var errors []error

	table := GetTestTableBasicWithCreatedDatabase(t)

	if table == nil {
		t.Error(fmt.Errorf("error: table is nil"))
		t.FailNow()
		return nil
	}

    table_errors := table.Create()
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return table
}

func GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t *testing.T, table_name string, schema json.Map) (*db_client.Table) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, table_name, schema)
	
	table_create_errors := table.Create()
	if table_create_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", table_create_errors))
		t.FailNow()
		return nil
	}

	return table
}
