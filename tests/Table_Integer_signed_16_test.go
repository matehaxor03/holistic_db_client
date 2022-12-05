package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerSigned16ColumnName() string {
	return "holistic_test_table_with_integer_signed_16"
}

func GetTestTableIntegerSigned16ColumnName() string {
	return "integer_signed_16_column"
}

func GetTestTableInteger16SignedColumnNameNotMandatory() string {
	return "integer_signed_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned16Column() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned16ColumnName(): class.Map {"type": "int16"}}
}

func GetTestSchemaWithIntegerSigned16NotMandatoryColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger16SignedColumnNameNotMandatory(): class.Map {"type": "*int16"}}
}

func TestTableCreateWithIntegerSigned16Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned16ColumnName(), GetTestSchemaWithIntegerSigned16Column())

}

func TestTableCreateWithIntegerSigned16NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned16ColumnName(), GetTestSchemaWithIntegerSigned16NotMandatoryColumn())

}
 
