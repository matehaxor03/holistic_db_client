package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableCreated(t *testing.T) (*class.Table) {
	var errors []error

	table := GetTestTable(t)

    table_errors := table.Create()
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return table
}
 
func TestRecordCanCreateRecord(t *testing.T) {
	table := GetTestTableCreated(t)

    record, record_errors := table.CreateRecord(class.Map{})
	if record_errors != nil {
		t.Error(record_errors)
	}

	if record == nil {
		t.Errorf("record is nil")
	}
}
