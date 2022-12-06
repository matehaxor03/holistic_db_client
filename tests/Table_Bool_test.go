package tests
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
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

func GetTestSchemaWithBoolColumn() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnName(): json.Map {"type": "bool"}}
}

func GetTestSchemaWithBoolColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnNameNotMandatory(): json.Map {"type": "*bool"}}
}

func TestTableCreateWithBoolColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableBoolColumnName(), GetTestSchemaWithBoolColumn())
}

func TestTableCreateWithBoolColumnNotMandatory(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithBoolColumnNotMandatory(), GetTestSchemaWithBoolColumnNotMandatory())
}
 
