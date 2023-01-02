package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithBoolTrue(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumn(), helper.GetTestSchemaWithBoolColumn())

	test_value := true
	map_record := json.NewMapValue()
	map_record.SetBoolValue(helper.GetTestTableBoolColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBoolValue(helper.GetTestTableBoolColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("value is nil")
		} else if value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %t actual: %t", test_value,  value))
		}
	}
}

func TestRecordCanCreateRecordWithBoolFalse(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumn(), helper.GetTestSchemaWithBoolColumn())

	test_value := false
	map_record := json.NewMapValue()
	map_record.SetBoolValue(helper.GetTestTableBoolColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetBoolValue(helper.GetTestTableBoolColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %t actual: %t", test_value,  value))
		}
	}
}

func TestRecordCanUpdateRecordWithBoolTrue(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumn(), helper.GetTestSchemaWithBoolColumn())

	map_record := json.NewMapValue()
	map_record.SetBoolValue(helper.GetTestTableBoolColumnName(), false)
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetBoolValue(helper.GetTestTableBoolColumnName(), true)
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
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumn(), helper.GetTestSchemaWithBoolColumn())

	map_record := json.NewMapValue()
	map_record.SetBoolValue(helper.GetTestTableBoolColumnName(), true)
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetBoolValue(helper.GetTestTableBoolColumnName(), false)
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
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumnNotMandatory(), helper.GetTestSchemaWithBoolColumnNotMandatory())

	test_value := true
	map_record := json.NewMapValue()
	map_record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &test_value)
	
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBool(helper.GetTestTableBoolColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %t actual: %t", test_value,  *value))
		}
	}
}

func TestRecordCanCreateRecordWithBoolNotMandatoryFalse(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumnNotMandatory(), helper.GetTestSchemaWithBoolColumnNotMandatory())

	test_value := false
	map_record := json.NewMapValue()
	map_record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetBool(helper.GetTestTableBoolColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %t actual: %t", test_value,  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithBoolNotMandatoryTrue(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumnNotMandatory(), helper.GetTestSchemaWithBoolColumnNotMandatory())

	test_value := true
	map_record := json.NewMapValue()
	map_record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := false
		set_errors := record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &update_value)
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
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithBoolColumnNotMandatory(), helper.GetTestSchemaWithBoolColumnNotMandatory())

	test_value := true
	map_record := json.NewMapValue()
	map_record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := false
		set_errors := record.SetBool(helper.GetTestTableBoolColumnNameNotMandatory(), &update_value)
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