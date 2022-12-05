package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithIntegerSigned08Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08Column())
	test_value := int8(100)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt8Value(GetTestTableIntegerSigned08ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != int8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int8(100),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned08Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08Column())
	test_value := int8(101)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int8(120)
		set_errors := record.SetInt8Value(GetTestTableIntegerSigned08ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerSigned08ColumnNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
	test_value := int8(100)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt8(GetTestTableIntegerSigned08ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != int8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int8(100),  *value))
		}
	}
}


func TestRecordCanUpdateRecordWithIntegerSigned08ColumNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
	test_value := int8(101)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int8(120)
		set_errors := record.SetInt8(GetTestTableIntegerSigned08ColumnNameNotMandatory(), &update_value)
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


