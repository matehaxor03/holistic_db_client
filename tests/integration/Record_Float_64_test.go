package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithFloat64(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat64Column(), helper.GetTestSchemaWithFloat64Column())

	test_value := float64(123456789.987654321)
	map_record := json.Map{}
	map_record.SetFloat64Value(helper.GetTestTableFloat64ColumnName(), test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat64Value(helper.GetTestTableFloat64ColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %f actual: %f", test_value,  value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat64(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat64Column(), helper.GetTestSchemaWithFloat64Column())

	map_record := json.Map{}
	map_record.SetFloat64Value(helper.GetTestTableFloat64ColumnName(), float64(123456789.987654321))
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetFloat64Value(helper.GetTestTableFloat64ColumnName(), float64(987654321.987654321))
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

func TestRecordCanCreateRecordWithFloat64NotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat64ColumnNotMandatory(), helper.GetTestSchemaWithFloat64ColumnNotMandatory())

	test_value := float64(987654321.987654321)
	map_record := json.Map{}
	map_record.SetFloat64(helper.GetTestTableFloat64ColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat64(helper.GetTestTableFloat64ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != test_value {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %f actual: %f", test_value,  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat64NotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithFloat64ColumnNotMandatory(), helper.GetTestSchemaWithFloat64ColumnNotMandatory())

	test_value := float64(123456789.123456789)
	map_record := json.Map{}
	map_record.SetFloat64(helper.GetTestTableFloat64ColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := float64(987654321.987654321)
		set_errors := record.SetFloat64(helper.GetTestTableFloat64ColumnNameNotMandatory(), &update_value)
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


