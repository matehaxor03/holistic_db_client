package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithIntegerSigned16Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned16ColumnName(), GetTestSchemaWithIntegerSigned16Column())
	test_value := int16(130)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned16ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("errro: %s", record_errors))
	}

	if record == nil {
		t.Errorf("record is nil")
	}

	value, value_errors := record.GetInt16(GetTestTableIntegerSigned16ColumnName())
	if value_errors != nil {
		t.Error(fmt.Sprintf("error: %s", value_errors))
	} else if class.IsNil(value_errors) {
		t.Errorf("error: value is nil")
	} else if *value != int16(130) {
		t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int16(130),  *value))
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned16Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned16ColumnName(), GetTestSchemaWithIntegerSigned16Column())
	test_value := int16(150)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned16ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int16(180)
		set_errors := record.SetInt16(GetTestTableIntegerSigned16ColumnName(), &update_value)
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


