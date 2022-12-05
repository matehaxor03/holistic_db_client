package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerUnsigned16ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_16"
}

func GetTestTableIntegerUnsigned16ColumnName() string {
	return "integer_unsigned_16_column"
}

func GetTestTableInteger16UnsignedColumnNameNotMandatory() string {
	return "integer_unsigned_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned16Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned16ColumnName(): class.Map {"type": "uint16"}}
}

func GetTestSchemaWithIntegerUnsigned16NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger16UnsignedColumnNameNotMandatory(): class.Map {"type": "*uint16"}}
}

func TestTableCreateWithIntegerUnsigned16Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned16ColumnName(), GetTestSchemaWithIntegerUnsigned16Column())

}

func TestTableCreateWithIntegerUnsigned16NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned16ColumnName(), GetTestSchemaWithIntegerUnsigned16NotMandatoryColumn())
}
 
