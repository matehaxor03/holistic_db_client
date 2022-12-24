package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableCreated(t *testing.T) (*class.Table) {
	var errors []error

	table := GetTestTableBasicWithCreatedDatabase(t)

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

func GetTestTableWithTableNameAndSchemaWithCreatedDatabaseAndTable(t *testing.T, table_name string, schema json.Map) (*class.Table) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, table_name, schema)
	
	table_create_errors := table.Create()
	if table_create_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", table_create_errors))
		t.FailNow()
		return nil
	}

	return table
}

 
func TestRecordCanCreateRecord(t *testing.T) {
	table := GetTestTableCreated(t)

    record, record_errors := table.CreateRecord(json.Map{})
	if record_errors != nil {
		t.Error(record_errors)
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		primary_key_id, primary_key_id_errors := record.GetUInt64(GetTestTablePrimaryKeyName())
		if primary_key_id_errors != nil {
			t.Error(fmt.Sprintf("%s", primary_key_id_errors))
		} else if common.IsNil(primary_key_id) {
			t.Errorf("primary key is nil")
		} else if *primary_key_id <= 0 {
			t.Errorf(fmt.Sprintf("primary key zero or less: %d", *primary_key_id))
		}
	}
}
