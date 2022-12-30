package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithFloat32(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat32Column(), helper.GetTestSchemaWithFloat32Column())

	test_value := float32(123456789.987654321)
	map_record := json.Map{}
	map_record.SetFloat32Value(helper.GetTestTableFloat32ColumnName(), test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat32Value(helper.GetTestTableFloat32ColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %f actual: %f", test_value,  value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat32(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat32Column(), helper.GetTestSchemaWithFloat32Column())

	map_record := json.Map{}
	map_record.SetFloat32Value(helper.GetTestTableFloat32ColumnName(), float32(123456789.987654321))
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetFloat32Value(helper.GetTestTableFloat32ColumnName(), float32(987654321.987654321))
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

func TestRecordCanCreateRecordWithFloat32NotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat32ColumnNotMandatory(), helper.GetTestSchemaWithFloat32ColumnNotMandatory())

	test_value := float32(987654321.987654321)
	map_record := json.Map{}
	map_record.SetFloat32(helper.GetTestTableFloat32ColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat32(helper.GetTestTableFloat32ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %f actual: %f", test_value,  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat32NotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat32ColumnNotMandatory(), helper.GetTestSchemaWithFloat32ColumnNotMandatory())

	test_value := float32(123456789.123456789)
	map_record := json.Map{}
	map_record.SetFloat32(helper.GetTestTableFloat32ColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := float32(987654321.987654321)
		set_errors := record.SetFloat32(helper.GetTestTableFloat32ColumnNameNotMandatory(), &update_value)
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


