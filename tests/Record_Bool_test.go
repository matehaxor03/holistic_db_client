package tests
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestRecordCanCreateRecordWithBoolTrue(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():true})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBoolValue(GetTestTableBoolColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("value is nil")
		} else if value != true {
			t.Errorf(fmt.Sprintf("value not equal expected: %t actual: %t", true,  value))
		}
	}
}

func TestRecordCanCreateRecordWithBoolFalse(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():false})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetBoolValue(GetTestTableBoolColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if value != false {
			t.Errorf(fmt.Sprintf("value not equal expected: %t actual: %t", false,  value))
		}
	}
}

func TestRecordCanUpdateRecordWithBoolTrue(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():false})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetBoolValue(GetTestTableBoolColumnName(), true)
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

func TestRecordCanUpdateRecordWithBoolFalse(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():true})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetBoolValue(GetTestTableBoolColumnName(), false)
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

func TestRecordCanCreateRecordWithBoolNotMandatoryTrue(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumnNotMandatory(), GetTestSchemaWithBoolColumnNotMandatory())

	test_value := true
    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBool(GetTestTableBoolColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != true {
			t.Errorf(fmt.Sprintf("value not equal expected: %t actual: %t", true,  *value))
		}
	}
}

func TestRecordCanCreateRecordWithBoolNotMandatoryFalse(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumnNotMandatory(), GetTestSchemaWithBoolColumnNotMandatory())

	test_value := false
    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBool(GetTestTableBoolColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if class.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != false {
			t.Errorf(fmt.Sprintf("value not equal expected: %t actual: %t", false,  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithBoolNotMandatoryTrue(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

	test_value := false
    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := true
		set_errors := record.SetBool(GetTestTableBoolColumnName(), &update_value)
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

func TestRecordCanUpdateRecordWithBoolNotMandatoryFalse(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithBoolColumn(), GetTestSchemaWithBoolColumn())

	test_value := true
    record, record_errors := table.CreateRecord(json.Map{GetTestTableBoolColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := false
		set_errors := record.SetBool(GetTestTableBoolColumnName(), &update_value)
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


