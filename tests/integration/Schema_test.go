package integration
 
import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestSchemaCanCreateTable(t *testing.T) {
	table, table_errors := helper.GetTestDatabaseCreated(t).CreateTable(helper.GetTestTableName(), helper.GetTestSchema())
	if table_errors != nil {
		t.Errorf("expect table_errors to be nil: %s", fmt.Sprintf("%s", table_errors))
	}

	if table == nil {
		t.Errorf("expect table to be not nil")
	}
}
 
func TestSchemaCannotCreateTableIfNil(t *testing.T) {
	table, table_errors := helper.GetTestDatabaseCreated(t).CreateTable(helper.GetTestTableName(), nil)
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfNoColumns(t *testing.T) {
	table, table_errors := helper.GetTestDatabase(t).CreateTable(helper.GetTestTableName(), json.Map{})
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableNoPrimaryKey(t *testing.T) {
	table, table_errors := helper.GetTestDatabase(t).CreateTable(helper.GetTestTableName(), helper.GetTestTableSchemaNoPrimaryKey())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfMoreThanOneAutoIncrementPrimaryKey(t *testing.T) {
	table, table_errors := helper.GetTestDatabase(t).CreateTable(helper.GetTestTableName(), helper.GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfNoTypeAttribute(t *testing.T) {
	table, table_errors := helper.GetTestDatabase(t).CreateTable(helper.GetTestTableName(), helper.GetTestTableSchemaNoType())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfHasValueAttribute(t *testing.T) {
	table, table_errors := helper.GetTestDatabase(t).CreateTable(helper.GetTestTableName(), helper.GetTestColumnSchemaWithValue())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

