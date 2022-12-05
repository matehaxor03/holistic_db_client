package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithFloat64(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat64Column(), GetTestSchemaWithFloat64Column())

    record, record_errors := table.CreateRecord(class.Map{GetTestTableFloat64ColumnName():float64(123456789.987654321)})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat64Value(GetTestTableFloat64ColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if value != float64(123456789.987654321) {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %f actual: %f", float64(123456789.987654321),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat64(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat64Column(), GetTestSchemaWithFloat64Column())

    record, record_errors := table.CreateRecord(class.Map{GetTestTableFloat64ColumnName():float64(123456789.987654321)})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetFloat64Value(GetTestTableFloat64ColumnName(), float64(987654321.987654321))
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

func TestRecordCanCreateRecordWithFloat64NotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat64ColumnNotMandatory(), GetTestSchemaWithFloat64ColumnNotMandatory())

	test_value := float64(987654321.987654321)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableFloat64ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat64(GetTestTableFloat64ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != float64(987654321.987654321) {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %f actual: %f",	float64(987654321.987654321), *value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat64NotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat64Column(), GetTestSchemaWithFloat64ColumnNotMandatory())

	test_value := float64(123456789.123456789)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableFloat64ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := float64(987654321.987654321)
		set_errors := record.SetFloat64(GetTestTableFloat64ColumnNameNotMandatory(), &update_value)
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


