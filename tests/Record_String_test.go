package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithString(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnName():"hello world"})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		string_value, string_value_errors := record.GetStringValue(GetTestTableStringColumnName())
		if string_value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", string_value_errors))
		} else if class.IsNil(string_value) {
			t.Errorf("error:  value is nil")
		} else if string_value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  string_value))
		}
	}
}

func TestRecordCanUpdateRecordWithString(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnName():"hello world"})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		set_errors := record.SetStringValue(GetTestTableStringColumnName(), "hello world2")
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

func TestRecordCanCreateRecordWithStringNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumnNotMandatory(), GetTestSchemaWithStringColumnNotMandatory())

	test_value := "hello world"
    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		string_value, string_value_errors := record.GetString(GetTestTableStringColumnNameNotMandatory())
		if string_value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", string_value_errors))
		} else if class.IsNil(string_value) {
			t.Errorf("error:  value is nil")
		} else if *string_value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  *string_value))
		}
	}
}

func TestRecordCanUpdateRecordWithStringNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

	test_value := "hello world"
    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := "hello world2"
		set_errors := record.SetString(GetTestTableStringColumnName(), &update_value)
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


