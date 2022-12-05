package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableCreatedWithIntegerSigned08Column(t *testing.T) (*class.Table) {
	var errors []error

	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08Column())

	if table == nil {
		t.Error(fmt.Errorf("error: table is nil"))
		t.FailNow()
		return nil
	}

    table_errors := table.Create()
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return table
}

func TestRecordCanCreateRecordWithIntegerSigned08Column(t *testing.T) {
	table := GetTestTableCreatedWithIntegerSigned08Column(t)
	test_value := int8(100)

    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("%s", record_errors))
	}

	if record == nil {
		t.Errorf("record is nil")
	}

	value, value_errors := record.GetInt8(GetTestTableIntegerSigned08ColumnName())
	if value_errors != nil {
		t.Error(fmt.Sprintf("%s", value_errors))
	} else if class.IsNil(value_errors) {
		t.Errorf("value is nil")
	} else if *value != int8(100) {
		t.Errorf(fmt.Sprintf("value not equal expected: %d actual: %d", int8(100),  *value))
	}
}

func TestRecordCanUpdateRecordWithIntegerSigned08Colum(t *testing.T) {
	table := GetTestTableCreatedWithIntegerSigned08Column(t)
	test_value := int8(101)
    record, record_errors := table.CreateRecord(class.Map{GetTestTableIntegerSigned08ColumnName():&test_value})
	if record_errors != nil {
		t.Errorf(fmt.Sprintf("%s", record_errors))
	} else if record == nil {
		t.Errorf("record is nil")
	} else {
		update_value := int8(120)
		set_errors := record.SetInt8(GetTestTableIntegerSigned08ColumnName(), &update_value)
		if set_errors != nil {
			t.Error(set_errors)
		} else {
			update_errors := record.Update()
			if update_errors != nil {
				t.Error(update_errors)
			}
		}
	}
}


