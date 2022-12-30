package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestRecordCanCreateRecordWithString(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithStringColumn(), helper.GetTestSchemaWithStringColumn())

	map_record := json.Map{}
	map_record.SetStringValue(helper.GetTestTableStringColumnName(), "hello world")
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetStringValue(helper.GetTestTableStringColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  value))
		}
	}
}

func TestRecordCanUpdateRecordWithString(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithStringColumn(), helper.GetTestSchemaWithStringColumn())

	map_record := json.Map{}
	map_record.SetStringValue(helper.GetTestTableStringColumnName(), "hello world")
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		set_errors := record.SetStringValue(helper.GetTestTableStringColumnName(), "hello world2")
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

func TestRecordCanCreateRecordWithStringNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithStringColumnNotMandatory(), helper.GetTestSchemaWithStringColumnNotMandatory())

	test_value := "hello world"
	map_record := json.Map{}
	map_record.SetString(helper.GetTestTableStringColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)

	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetString(helper.GetTestTableStringColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if *value != "hello world" {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %s actual: %s", "hello world",  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithStringNotMandatory(t *testing.T) {
	table := helper.GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, helper.GetTestTableNameWithStringColumnNotMandatory(), helper.GetTestSchemaWithStringColumnNotMandatory())

	test_value := "hello world"
	map_record := json.Map{}
	map_record.SetString(helper.GetTestTableStringColumnNameNotMandatory(), &test_value)
	record, record_errors := table.CreateRecord(map_record)


	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := "hello world2"
		set_errors := record.SetString(helper.GetTestTableStringColumnNameNotMandatory(), &update_value)
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


