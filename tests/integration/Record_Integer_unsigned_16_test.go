package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerUnsigned16Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned16ColumnName(), helper.GetTestSchemaWithIntegerUnsigned16Column())
	
	test_value := uint16(130)
	map_record := json.Map{}
	map_record.SetUInt16Value(helper.GetTestTableIntegerUnsigned16ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("errro: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt16Value(helper.GetTestTableIntegerUnsigned16ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != uint16(130) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint16(130),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned16Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned16ColumnName(), helper.GetTestSchemaWithIntegerUnsigned16Column())
	
	test_value := uint16(150)
	map_record := json.Map{}
	map_record.SetUInt16Value(helper.GetTestTableIntegerUnsigned16ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint16(180)
		set_errors := record.SetUInt16Value(helper.GetTestTableIntegerUnsigned16ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerUnsigned16ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory())
	
	test_value := uint16(188)
	map_record := json.Map{}
	map_record.SetUInt16(helper.GetTestTableIntegerUnsigned16ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt16(helper.GetTestTableIntegerUnsigned16ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != uint16(130) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint16(188),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned16ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory())
	
	test_value := uint16(255)
	map_record := json.Map{}
	map_record.SetUInt16(helper.GetTestTableIntegerUnsigned16ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint16(180)
		set_errors := record.SetUInt16(helper.GetTestTableIntegerUnsigned16ColumnNameNotMandatory(), &update_value)
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





