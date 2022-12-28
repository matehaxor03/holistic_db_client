package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerUnsigned08Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnName(), helper.GetTestSchemaWithIntegerUnsigned08Column())
	test_value := uint8(100)

    record, record_errors := table.CreateRecord(json.Map{helper.GetTestTableIntegerUnsigned08ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt8Value(helper.GetTestTableIntegerUnsigned08ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != uint8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint8(100),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned08Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnName(), helper.GetTestSchemaWithIntegerUnsigned08Column())
	test_value := uint8(101)
    record, record_errors := table.CreateRecord(json.Map{helper.GetTestTableIntegerUnsigned08ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint8(120)
		set_errors := record.SetUInt8Value(helper.GetTestTableIntegerUnsigned08ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerUnsigned08ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn())
	test_value := uint8(100)

    record, record_errors := table.CreateRecord(json.Map{helper.GetTestTableIntegerUnsigned08ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt8(helper.GetTestTableIntegerUnsigned08ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != uint8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint8(100),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned08ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn())
	test_value := uint8(101)
    record, record_errors := table.CreateRecord(json.Map{helper.GetTestTableIntegerUnsigned08ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint8(120)
		set_errors := record.SetUInt8(helper.GetTestTableIntegerUnsigned08ColumnNameNotMandatory(), &update_value)
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





