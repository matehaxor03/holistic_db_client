package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
)

func TestRecordCanCreateRecordWithIntegerSigned64Column(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned64ColumnName(), GetTestSchemaWithIntegerSigned64Column())
	test_value := int64(555)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerSigned64ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt64Value(GetTestTableIntegerSigned64ColumnName())
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
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned64ColumnName(), GetTestSchemaWithIntegerSigned64Column())
	test_value := int64(66775)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableIntegerSigned64ColumnName():test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int64(45452)
		set_errors := record.SetInt64Value(GetTestTableIntegerSigned64ColumnName(), update_value)
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
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
	test_value := int64(555)

    record, record_errors := table.CreateRecord(json.Map{GetTestTableInteger64SignedColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		value, value_errors := record.GetInt64(GetTestTableInteger64SignedColumnNameNotMandatory())
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
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t, GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
	test_value := int64(66775)
    record, record_errors := table.CreateRecord(json.Map{GetTestTableInteger64SignedColumnNameNotMandatory():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", record_errors))
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		update_value := int64(45452)
		set_errors := record.SetInt64(GetTestTableInteger64SignedColumnNameNotMandatory(), &update_value)
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


