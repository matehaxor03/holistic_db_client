package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithBoolColumn() string {
	return "holistic_test_table_with_bool"
}

func GetTestTableNameWithBoolColumnNotMandatory() string {
	return "holistic_test_table_with_boool_not_mandatory"
}

func GetTestTableBoolColumnName() string {
	return "bool_column"
}

func GetTestTableBoolColumnNameNotMandatory() string {
	return "bool_column_not_mandatory"
}

func GetTestSchemaWithBoolColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnName(): class.Map {"type": "bool"}}
}

func GetTestSchemaWithBoolColumnNotMandatory() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnNameNotMandatory(): class.Map {"type": "*bool"}}
}

func TestTableCreateWithBoolColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableBoolColumnName(), GetTestSchemaWithBoolColumn())
}

func TestTableCreateWithBoolColumnNotMandatory(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithBoolColumnNotMandatory(), GetTestSchemaWithBoolColumnNotMandatory())
}
 
