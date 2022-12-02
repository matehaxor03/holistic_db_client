package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestColumnSchemaNoPrimaryKey() class.Map {
	return class.Map {"type": "uint64" }
}

func GetTestColumnSchemaNoType() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"primary_key": true}}
}

func GetTestColumnSchemaWithValue() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "value":"something", "auto_increment": true, "primary_key": true}}
}

func GetTestTableSchemaNoPrimaryKey() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoPrimaryKey()}
}

func GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): GetTestSchemaColumnPrimaryKeyAutoIncrement(),
	                  GetTestTablePrimaryKeyName2(): GetTestSchemaColumnPrimaryKeyAutoIncrement()}
}

func GetTestTableSchemaNoType() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoType()}
}

func TestSchemaCanCreateTable(t *testing.T) {
	table, table_errors := GetTestDatabaseCreated(t).CreateTable(GetTestTableName(), GetTestSchema())
	if table_errors != nil {
		t.Errorf("expect table_errors to be nil: %s", fmt.Sprintf("%s", table_errors))
	}

	if table == nil {
		t.Errorf("expect table to be not nil")
	}
}
 
func TestSchemaCannotCreateTableIfNil(t *testing.T) {
	table, table_errors := GetTestDatabaseCreated(t).CreateTable(GetTestTableName(), nil)
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfNoColumns(t *testing.T) {
	table, table_errors := GetTestDatabase(t).CreateTable(GetTestTableName(), class.Map{})
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableNoPrimaryKey(t *testing.T) {
	table, table_errors := GetTestDatabase(t).CreateTable(GetTestTableName(), GetTestTableSchemaNoPrimaryKey())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfMoreThanOneAutoIncrementPrimaryKey(t *testing.T) {
	table, table_errors := GetTestDatabase(t).CreateTable(GetTestTableName(), GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfNoTypeAttribute(t *testing.T) {
	table, table_errors := GetTestDatabase(t).CreateTable(GetTestTableName(), GetTestTableSchemaNoType())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotCreateTableIfHasValueAttribute(t *testing.T) {
	table, table_errors := GetTestDatabase(t).CreateTable(GetTestTableName(), GetTestColumnSchemaWithValue())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

