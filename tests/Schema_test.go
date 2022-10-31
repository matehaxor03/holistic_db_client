package tests
 
import (
    "testing"
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

func TestSchemaCanNewTable(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), GetTestSchema())
	if table_errors != nil {
		t.Errorf("expect table_errors to be nil")
	}

	if table == nil {
		t.Errorf("expect table to be not nil")
	}
}
 
func TestSchemaCanNewTableIfNil(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), nil)
	if table_errors != nil {
		t.Errorf("expect table_errors to be nil")
	}

	if table == nil {
		t.Errorf("expect table to be not nil")
	}
}

func TestSchemaCannotNewTableIfNoColumns(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), class.Map{})
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotNewTableIfNoPrimaryKey(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), GetTestTableSchemaNoPrimaryKey())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotNewTableIfMoreThanOneAutoIncrementPrimaryKey(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotNewTableIfNoTypeAttribute(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), GetTestTableSchemaNoType())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

func TestSchemaCannotNewTableIfHasValueAttribute(t *testing.T) {
	table, table_errors := class.NewTable(GetTestDatabaseCreated(t), GetTestTableName(), GetTestColumnSchemaWithValue())
	if table_errors == nil {
		t.Errorf("expect table_errors to be not nil")
	}

	if table != nil {
		t.Errorf("expect table to be nil")
	}
}

