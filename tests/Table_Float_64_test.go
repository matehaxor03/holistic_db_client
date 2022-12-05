package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithFloat64Column() string {
	return "holistic_test_table_with_float64"
}

func GetTestTableNameWithFloat64ColumnNotMandatory() string {
	return "holistic_test_table_with_float64_not_mandatory"
}

func GetTestTableFloat64ColumnName() string {
	return "float64_column"
}

func GetTestTableFloat64ColumnNameNotMandatory() string {
	return "float64_column_not_mandatory"
}

func GetTestSchemaWithFloat64Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat64ColumnName(): class.Map {"type": "float64"}}
}

func GetTestSchemaWithFloat64ColumnNotMandatory() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat64ColumnNameNotMandatory(): class.Map {"type": "*float64"}}
}

func TestTableCreateWithFloat64Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableFloat64ColumnName(), GetTestSchemaWithFloat64Column())
}

func TestTableCreateWithFloat64ColumnNotMandatory(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithFloat64ColumnNotMandatory(), GetTestSchemaWithFloat64ColumnNotMandatory())
}
 
