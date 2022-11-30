package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableCreated(t *testing.T) (*class.Table) {
	var errors []error

	table := GetTestTable(t)

	if table == nil {
		t.Error(fmt.Errorf("table is nil"))
		t.FailNow()
		return nil
	}

    table_errors := table.Create()
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return table
}

func GetTestTableCreatedWithString(t *testing.T) (*class.Table) {
	var errors []error

	table := GetTestTableWithString(t)

	if table == nil {
		t.Error(fmt.Errorf("table is nil"))
		t.FailNow()
		return nil
	}

    table_errors := table.Create()
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
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

	primary_key_id, primary_key_id_errors := record.GetUInt64(GetTestTablePrimaryKeyName())
	if primary_key_id_errors != nil {
		t.Error(fmt.Sprintf("%s", primary_key_id_errors))
	} else if class.IsNil(primary_key_id) {
		t.Errorf("primary key is nil")
	} else if *primary_key_id <= 0 {
		t.Errorf(fmt.Sprintf("primary key zero or less: %d", *primary_key_id))
	}
}

func TestRecordCanCreateRecordWithString(t *testing.T) {
	table := GetTestTableCreatedWithString(t)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnName():"hello world"})
	if record_errors != nil {
		t.Error(record_errors)
	}

	if record == nil {
		t.Errorf("record is nil")
	}

	string_value, string_value_errors := record.GetString(GetTestTableStringColumnName())
	if string_value_errors != nil {
		t.Error(fmt.Sprintf("%s", string_value_errors))
	} else if class.IsNil(string_value) {
		t.Errorf("string value is nil")
	} else if *string_value != "hello world" {
		t.Errorf(fmt.Sprintf("string value not equal expected: %s actual: %s", "hello world",  *string_value))
	}
}

func TestRecordCanUpdateRecordWithString(t *testing.T) {
	table := GetTestTableCreatedWithString(t)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableStringColumnName():"hello world"})
	if record_errors != nil {
		t.Error(record_errors)
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		set_string_errors := record.SetStringValue(GetTestTableStringColumnName(), "hello world2")
		if set_string_errors != nil {
			t.Error(set_string_errors)
		} else {
			update_errors := record.Update()
			if update_errors != nil {
				t.Error(update_errors)
			}
		}
	}
}


