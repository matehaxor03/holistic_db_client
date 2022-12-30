package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerUnsigned64Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned64ColumnName(), helper.GetTestSchemaWithIntegerUnsigned64Column())
	
	test_value := uint64(555)
	map_record := json.Map{}
	map_record.SetUInt64Value(helper.GetTestTableIntegerUnsigned64ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt64Value(helper.GetTestTableIntegerUnsigned64ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != uint64(555) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint64(555),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned64Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned64ColumnName(), helper.GetTestSchemaWithIntegerUnsigned64Column())
	
	test_value := uint64(555)
	map_record := json.Map{}
	map_record.SetUInt64Value(helper.GetTestTableIntegerUnsigned64ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint64(453535)
		set_errors := record.SetUInt64Value(helper.GetTestTableIntegerUnsigned64ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerUnsigned64ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
	
	test_value := uint64(349734)
	map_record := json.Map{}
	map_record.SetUInt64(helper.GetTestTableIntegerUnsigned64ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt64(helper.GetTestTableIntegerUnsigned64ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != uint64(349734) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint64(349734),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned64ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
	
	test_value := uint64(349734)
	map_record := json.Map{}
	map_record.SetUInt64(helper.GetTestTableIntegerUnsigned64ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)
	
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint64(453535)
		set_errors := record.SetUInt64(helper.GetTestTableIntegerUnsigned64ColumnNameNotMandatory(), &update_value)
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



