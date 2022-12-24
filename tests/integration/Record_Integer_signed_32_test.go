package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithIntegerSigned32Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32Column())
	test_value := int32(188)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerSigned32ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt32Value(GetTestTableIntegerSigned32ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != int32(188) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int32(188),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned32Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32Column())
	test_value := int32(199)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerSigned32ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int32(834)
		set_errors := record.SetInt32Value(GetTestTableIntegerSigned32ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerSigned32ColumnNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned32ColumnNotMandatory())
	test_value := int32(188)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableInteger32SignedColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt32(GetTestTableInteger32SignedColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != int32(188) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", int32(188),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned32ColumNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned32ColumnNotMandatory())
	test_value := int32(199)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableInteger32SignedColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int32(834)
		set_errors := record.SetInt32(GetTestTableInteger32SignedColumnNameNotMandatory(), &update_value)
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


