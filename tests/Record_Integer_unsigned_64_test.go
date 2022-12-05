package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithIntegerUnsigned64Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
	test_value := uint64(555)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerUnsigned64ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	}

	if record == nil {
		t.Errorf("error: record is nil")
	}

	value, value_errors := record.GetUInt64(GetTestTableIntegerUnsigned64ColumnName())
	if value_errors != nil {
		t.Error(fmt.Sprintf("error: %s", value_errors))
	} else if class.IsNil(value_errors) {
		t.Errorf("error: value is nil")
	} else if *value != uint64(555) {
		t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint64(555),  *value))
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned64Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
	test_value := uint64(343443)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerUnsigned64ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint64(453535)
		set_errors := record.SetUInt64(GetTestTableIntegerUnsigned64ColumnName(), &update_value)
		if set_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", set_errors))
		} else {
			update_errors := record.Update()
			if update_errors != nil {
				t.Errorf(fmt.Sprintf("error: %s", update_errors))
			}
		}
	}
}


