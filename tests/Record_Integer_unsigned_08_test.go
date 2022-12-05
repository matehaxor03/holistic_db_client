package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithIntegerUnsigned08Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned08ColumnName(), GetTestSchemaWithIntegerUnsigned08Column())
	test_value := uint8(100)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerUnsigned08ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	}

	if record == nil {
		t.Errorf("error: record is nil")
	}

	value, value_errors := record.GetUInt8(GetTestTableIntegerUnsigned08ColumnName())
	if value_errors != nil {
		t.Error(fmt.Sprintf("error: %s", value_errors))
	} else if class.IsNil(value_errors) {
		t.Errorf("error: value is nil")
	} else if *value != uint8(100) {
		t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint8(100),  *value))
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned08Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned08ColumnName(), GetTestSchemaWithIntegerUnsigned08Column())
	test_value := uint8(101)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerUnsigned08ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint8(120)
		set_errors := record.SetUInt8(GetTestTableIntegerUnsigned08ColumnName(), &update_value)
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


