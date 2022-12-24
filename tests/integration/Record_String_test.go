package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithString(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableStringColumnName():"hello world"})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetStringValue(GetTestTableStringColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  value))
		}
	}
}

func TestRecordCanUpdateRecordWithString(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableStringColumnName():"hello world"})
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
    record, record_errors := table.CreateRecord(json.Map{GetTestTableStringColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetString(GetTestTableStringColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if *value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithStringNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

	test_value := "hello world"
    record, record_errors := table.CreateRecord(json.Map{GetTestTableStringColumnName():&test_value})
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


