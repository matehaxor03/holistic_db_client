package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerSigned32ColumnName() string {
	return "holistic_test_table_with_integer_signed_32"
}

func GetTestTableIntegerSigned32ColumnName() string {
	return "integer_signed_32_column"
}

func GetTestTableInteger32SignedColumnNameNotMandatory() string {
	return "integer_signed_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned32Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned32ColumnName(): class.Map {"type": "int32"}}
}

func GetTestSchemaWithIntegerSigned32NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger32SignedColumnNameNotMandatory(): class.Map {"type": "*int32"}}
}

func TestTableCreateWithIntegerSigned32Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32Column())
}

func TestTableCreateWithIntegerSigned32NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned32ColumnName(), GetTestSchemaWithIntegerSigned32NotMandatoryColumn())
}
 
