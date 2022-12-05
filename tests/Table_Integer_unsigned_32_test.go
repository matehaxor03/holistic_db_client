package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerUnsigned32ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_32"
}

func GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_32_not_mandatory"
}

func GetTestTableIntegerUnsigned32ColumnName() string {
	return "integer_unsigned_32_column"
}

func GetTestTableIntegerUnsigned32ColumnNameNotMandatory() string {
	return "integer_unsigned_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned32Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned32ColumnName(): class.Map {"type": "uint32"}}
}

func GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableIntegerUnsigned32ColumnNameNotMandatory(): class.Map {"type": "*uint32"}}
}

func TestTableCreateWithIntegerUnsigned32Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned32ColumnName(), GetTestSchemaWithIntegerUnsigned32Column())
}

func TestTableCreateWithIntegerUnsigned32NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
}
 
