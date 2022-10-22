package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)
 
func TestSchemaCannotNewTableIfNil(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), nil)
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotNewTableIfNoColumns(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), class.Map{})
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

