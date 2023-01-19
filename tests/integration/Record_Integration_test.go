package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)
 
func TestRecordCanCreateRecord(t *testing.T) {
	table := helper.GetTestTableCreated(t)

    record, record_errors := table.CreateRecord(json.NewMapValue())
	if record_errors != nil {
		t.Error(record_errors)
	} else if record == nil {
		t.Errorf("error: record is nil")
	} else {
		primary_key_id, primary_key_id_errors := record.GetUInt64(helper.GetTestTablePrimaryKeyName())
		if primary_key_id_errors != nil {
			t.Error(fmt.Sprintf("%s", primary_key_id_errors))
		} else if common.IsNil(primary_key_id) {
			t.Errorf("primary key is nil")
		} else if *primary_key_id <= 0 {
			t.Errorf(fmt.Sprintf("primary key zero or less: %d", *primary_key_id))
		}
	}
}
