package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithIntegerSigned32Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32Column())
	test_value := int32(188)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned32ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	}

	if record == nil {
		t.Errorf("error: record is nil")
	}

	value, value_errors := record.GetInt32(GetTestTableIntegerSigned32ColumnName())
	if value_errors != nil {
		t.Error(fmt.Sprintf("error: %s", value_errors))
	} else if class.IsNil(value_errors) {
		t.Errorf("error: value is nil")
	} else if *value != int32(188) {
		t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int32(188),  *value))
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned32Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32Column())
	test_value := int32(199)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned32ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int32(834)
		set_errors := record.SetInt32(GetTestTableIntegerSigned32ColumnName(), &update_value)
		if set_errors != nil {
			t.Error(set_errors)
		} else {
			update_errors := record.Update()
			if update_errors != nil {
				t.Error(update_errors)
			}
		}
	}
}


