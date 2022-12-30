package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerSigned16Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned16ColumnName(), helper.GetTestSchemaWithIntegerSigned16Column())
	
	test_value := int16(130)
	map_record := json.Map{}
	map_record.SetInt16Value(helper.GetTestTableIntegerSigned16ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("errro: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetInt16Value(helper.GetTestTableIntegerSigned16ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != int16(130) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int16(130),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned16Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned16ColumnName(), helper.GetTestSchemaWithIntegerSigned16Column())
	
	test_value := int16(150)
	map_record := json.Map{}
	map_record.SetInt16Value(helper.GetTestTableIntegerSigned16ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int16(180)
		set_errors := record.SetInt16Value(helper.GetTestTableIntegerSigned16ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerSigned16ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned16ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned16ColumnNotMandatory())
	
	test_value := int16(130)
	map_record := json.Map{}
	map_record.SetInt16(helper.GetTestTableIntegerSigned16ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("errro: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		value, value_errors := record.GetInt16(helper.GetTestTableIntegerSigned16ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != int16(130) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int16(130),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned16ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned16ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned16ColumnNotMandatory())
	
	test_value := int16(130)
	map_record := json.Map{}
	map_record.SetInt16(helper.GetTestTableIntegerSigned16ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)
	
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int16(180)
		set_errors := record.SetInt16(helper.GetTestTableIntegerSigned16ColumnNameNotMandatory(), &update_value)
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


