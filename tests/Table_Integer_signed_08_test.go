package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerSigned08ColumnName() string {
	return "holistic_test_table_with_integer_signed_08"
}

func GetTestTableIntegerSigned08ColumnName() string {
	return "integer_signed_08_column"
}

func GetTestTableIntegerSigned08ColumnNameNotMandatory() string {
	return "integer_signed_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned08Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnName(): class.Map {"type": "int8", "mandatory":true}}
}

func GetTestSchemaWithIntegerSigned08NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnNameNotMandatory(): class.Map {"type": "*int8", "mandatory":true}}
}

func TestTableCreateWithIntegerSigned08Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08Column())
}

func TestTableCreateWithIntegerSigned08NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
}
 
