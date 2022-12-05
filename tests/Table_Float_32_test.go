package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithFloat32Column() string {
	return "holistic_test_table_with_float32"
}

func GetTestTableNameWithFloat32ColumnNotMandatory() string {
	return "holistic_test_table_with_float32_not_mandatory"
}

func GetTestTableFloat32ColumnName() string {
	return "float32_column"
}

func GetTestTableFloat32ColumnNameNotMandatory() string {
	return "float32_column_not_mandatory"
}

func GetTestSchemaWithFloat32Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat32ColumnName(): class.Map {"type": "float32"}}
}

func GetTestSchemaWithFloat32ColumnNotMandatory() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat32ColumnNameNotMandatory(): class.Map {"type": "*float32"}}
}

func TestTableCreateWithFloat32Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableFloat32ColumnName(), GetTestSchemaWithFloat32Column())
}

func TestTableCreateWithFloat32ColumnNotMandatory(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithFloat32ColumnNotMandatory(), GetTestSchemaWithFloat32ColumnNotMandatory())
}
 
