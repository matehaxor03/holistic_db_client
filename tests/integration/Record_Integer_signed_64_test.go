package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithIntegerSigned64Column(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned64ColumnName(), helper.GetTestSchemaWithIntegerSigned64Column())
	
	test_value := int64(555)
	map_record := json.Map{}
	map_record.SetInt64Value(helper.GetTestTableIntegerSigned64ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt64Value(helper.GetTestTableIntegerSigned64ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != int64(555) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int64(555),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned64Colum(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned64ColumnName(), helper.GetTestSchemaWithIntegerSigned64Column())
	
	test_value := int64(234234)
	map_record := json.Map{}
	map_record.SetInt64Value(helper.GetTestTableIntegerSigned64ColumnName(), test_value)
    record, record_errors := table.CreateRecord(map_record)
	
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int64(45452)
		set_errors := record.SetInt64Value(helper.GetTestTableIntegerSigned64ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerSigned64ColumnNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
	
	test_value := int64(555)
	map_record := json.Map{}
	map_record.SetInt64(helper.GetTestTableIntegerSigned64ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt64(helper.GetTestTableIntegerSigned64ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != int64(555) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int64(555),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned64ColumNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
	
	test_value := int64(435345)
	map_record := json.Map{}
	map_record.SetInt64(helper.GetTestTableIntegerSigned64ColumnNameNotMandatory(), &test_value)
    record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int64(45452)
		set_errors := record.SetInt64(helper.GetTestTableIntegerSigned64ColumnNameNotMandatory(), &update_value)
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


