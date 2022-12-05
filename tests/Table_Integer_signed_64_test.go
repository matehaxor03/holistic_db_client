package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerSigned64ColumnName() string {
	return "holistic_test_table_with_integer_signed_64"
}

func GetTestTableIntegerSigned64ColumnName() string {
	return "integer_signed_64_column"
}

func GetTestTableInteger64SignedColumnNameNotMandatory() string {
	return "integer_signed_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned64Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned64ColumnName(): class.Map {"type": "int64"}}
}

func GetTestSchemaWithIntegerSigned64NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger64SignedColumnNameNotMandatory(): class.Map {"type": "*int64"}}
}

func TestTableCreateWithIntegerSigned64Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned64ColumnName(), GetTestSchemaWithIntegerSigned64Column())
}

func TestTableCreateWithIntegerSigned64NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned64ColumnName(), GetTestSchemaWithIntegerSigned64NotMandatoryColumn())
}
 
