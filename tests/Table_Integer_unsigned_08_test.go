package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerUnsigned08ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_08"
}

func GetTestTableIntegerUnsigned08ColumnName() string {
	return "integer_unsigned_08_column"
}

func GetTestTableIntegerUnsigned08ColumnNameNotMandatory() string {
	return "integer_unsigned_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned08Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned08ColumnName(): class.Map {"type": "uint8", "mandatory":true}}
}

func GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned08ColumnNameNotMandatory(): class.Map {"type": "*uint8", "mandatory":true}}
}

func TestTableCreateWithIntegerUnsigned08Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned08ColumnName(), GetTestSchemaWithIntegerUnsigned08Column())
}

func TestTableCreateWithIntegerUnsigned08NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned08ColumnName(), GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn())
}
 
