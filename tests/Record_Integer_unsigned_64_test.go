package tests
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithIntegerUnsigned64Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
	test_value := uint64(555)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned64ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt64Value(GetTestTableIntegerUnsigned64ColumnName())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if value != uint64(555) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint64(555),  value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned64Colum(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
	test_value := uint64(343443)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned64ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint64(453535)
		set_errors := record.SetUInt64Value(GetTestTableIntegerUnsigned64ColumnName(), update_value)
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

func TestRecordCanCreateRecordWithIntegerUnsigned64ColumnNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
	test_value := uint64(349734)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned64ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetUInt64(GetTestTableIntegerUnsigned64ColumnNameNotMandatory())
		if value_errors != nil {
			t.Error(fmt.Sprintf("error: %s", value_errors))
		} else if common.IsNil(value_errors) {
			t.Errorf("error: value is nil")
		} else if *value != uint64(349734) {
			t.Errorf(fmt.Sprintf("error: value not equal expected: %d actual: %d", uint64(349734),  *value))
		}
	}
}

func TestRecordCanUpdateRecordWithIntegerUnsigned64ColumNotMandatory(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
	test_value := uint64(343443)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerUnsigned64ColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := uint64(453535)
		set_errors := record.SetUInt64(GetTestTableIntegerUnsigned64ColumnNameNotMandatory(), &update_value)
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



