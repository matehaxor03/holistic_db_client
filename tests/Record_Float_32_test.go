package tests
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithFloat32(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat32Column(), GetTestSchemaWithFloat32Column())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableFloat32ColumnName():float32(123456789.987654321)})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat32Value(GetTestTableFloat32ColumnName())
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error:  value is nil")
		} else if value != float32(123456789.987654321) {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %f actual: %f", float32(123456789.987654321),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat32(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat32Column(), GetTestSchemaWithFloat32Column())

    record, record_errors := table.CreateRecord(json.Map{GetTestTableFloat32ColumnName():float32(123456789.987654321)})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		set_errors := record.SetFloat32Value(GetTestTableFloat32ColumnName(), float32(987654321.987654321))
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
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat32ColumnNotMandatory(), GetTestSchemaWithFloat32ColumnNotMandatory())

	test_value := float32(987654321.987654321)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableFloat32ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetFloat32(GetTestTableFloat32ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value) {
			t.Errorf("error: value is nil")
		} else if *value != float32(987654321.987654321) {
			t.Errorf(fmt.Sprintf("error:  value not equal expected: %f actual: %f",	float32(987654321.987654321), *value))
		}
	}
}

func TestRecordCanUpdateRecordWithFloat32NotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithFloat32Column(), GetTestSchemaWithFloat32ColumnNotMandatory())

	test_value := float32(123456789.123456789)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableFloat32ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := float32(987654321.987654321)
		set_errors := record.SetFloat32(GetTestTableFloat32ColumnNameNotMandatory(), &update_value)
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


