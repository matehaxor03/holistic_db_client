package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerUnsigned32Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestSchemaWithIntegerUnsigned32Column())
	
	test_value := uint32(188)
	map_record := json.NewMapValue()
	map_record.SetUInt32Value(helper.GetTestTableIntegerUnsigned32ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt32Value(helper.GetTestTableIntegerUnsigned32ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", test_value, value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned32Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestSchemaWithIntegerUnsigned32Column())
	
	test_value := uint32(5546)
	map_record := json.NewMapValue()
	map_record.SetUInt32Value(helper.GetTestTableIntegerUnsigned32ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint32(84534)
		set_errors := record.SetUInt32Value(helper.GetTestTableIntegerUnsigned32ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerUnsigned32ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
	
	test_value := uint32(188)
	map_record := json.NewMapValue()
	map_record.SetUInt32(helper.GetTestTableIntegerUnsigned32ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt32(helper.GetTestTableIntegerUnsigned32ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", test_value,  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned32ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
	
	test_value := uint32(4345)
	map_record := json.NewMapValue()
	map_record.SetUInt32(helper.GetTestTableIntegerUnsigned32ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint32(84534)
		set_errors := record.SetUInt32(helper.GetTestTableIntegerUnsigned32ColumnNameNotMandatory(), &update_value)
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


