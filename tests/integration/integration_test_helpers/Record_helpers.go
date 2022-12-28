package integration_test_helpers

import (
    "testing"
	//"strings"
	"fmt"
	//"sync"
	json "github.com/matehaxor03/holistic_json/json"
	//common "github.com/matehaxor03/holistic_common/common"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableCreated(t *testing.T) (*class.Table) {
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

func GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t *testing.T, table_name string, schema json.Map) (*class.Table) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, table_name, schema)
	
	table_create_errors := table.Create()
	if table_create_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", table_create_errors))
		t.FailNow()
		return nil
	}

	return table
}
