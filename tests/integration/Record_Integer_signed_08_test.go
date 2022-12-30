package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerSigned08Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned08ColumnName(), helper.GetTestSchemaWithIntegerSigned08Column())
	
	test_value := int8(100)
	map_record := json.Map{}
	map_record.SetInt8Value(helper.GetTestTableIntegerSigned08ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt8Value(helper.GetTestTableIntegerSigned08ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != int8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int8(100),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned08Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned08ColumnName(), helper.GetTestSchemaWithIntegerSigned08Column())
	
	test_value := int8(101)
	map_record := json.Map{}
	map_record.SetInt8Value(helper.GetTestTableIntegerSigned08ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int8(120)
		set_errors := record.SetInt8Value(helper.GetTestTableIntegerSigned08ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerSigned08ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
	
	test_value := int8(101)
	map_record := json.Map{}
	map_record.SetInt8(helper.GetTestTableIntegerSigned08ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt8(helper.GetTestTableIntegerSigned08ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != int8(100) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int8(100),  *value))
		}
	}
}


func TestRecordCanUpdateRecordWithIntegerSigned08ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
	
	test_value := int8(101)
	map_record := json.Map{}
	map_record.SetInt8(helper.GetTestTableIntegerSigned08ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int8(120)
		set_errors := record.SetInt8(helper.GetTestTableIntegerSigned08ColumnNameNotMandatory(), &update_value)
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


