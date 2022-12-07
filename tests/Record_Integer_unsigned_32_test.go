package tests
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithIntegerUnsigned32Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned32ColumnName(), GetTestSchemaWithIntegerUnsigned32Column())
	test_value := uint32(188)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned32ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt32Value(GetTestTableIntegerUnsigned32ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != uint32(188) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint32(188),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned32Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned32ColumnName(), GetTestSchemaWithIntegerUnsigned32Column())
	test_value := uint32(5556)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned32ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint32(84534)
		set_errors := record.SetUInt32Value(GetTestTableIntegerUnsigned32ColumnName(), update_value)
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
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
	test_value := uint32(188)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned32ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt32(GetTestTableIntegerUnsigned32ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != uint32(188) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint32(188),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned32ColumNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
	test_value := uint32(5556)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned32ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint32(84534)
		set_errors := record.SetUInt32(GetTestTableIntegerUnsigned32ColumnNameNotMandatory(), &update_value)
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


