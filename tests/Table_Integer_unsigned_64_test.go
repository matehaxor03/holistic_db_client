package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerUnsigned64ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_64"
}

func GetTestTableIntegerUnsigned64ColumnName() string {
	return "integer_unsigned_64_column"
}

func GetTestTableInteger64UnsignedColumnNameNotMandatory() string {
	return "integer_unsigned_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned64Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned64ColumnName(): class.Map {"type": "uint64"}}
}

func GetTestSchemaWithIntegerUnsigned64NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger64UnsignedColumnNameNotMandatory(): class.Map {"type": "*uint64"}}
}

func TestTableCreateWithIntegerUnsigned64Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
}

func TestTableCreateWithIntegerUnsigned64NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64NotMandatoryColumn())
}
 
