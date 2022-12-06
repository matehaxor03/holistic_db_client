package tests
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned64ColumnName() string {
	return "holistic_test_table_with_integer_signed_64"
}

func GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_64_not_mandatory"
}


func GetTestTableIntegerSigned64ColumnName() string {
	return "integer_signed_64_column"
}

func GetTestTableInteger64SignedColumnNameNotMandatory() string {
	return "integer_signed_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned64Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned64ColumnName(): json.Map {"type": "int64"}}
}

func GetTestSchemaWithIntegerSigned64ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger64SignedColumnNameNotMandatory(): json.Map {"type": "*int64"}}
}

func TestTableCreateWithIntegerSigned64Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned64ColumnName(), GetTestSchemaWithIntegerSigned64Column())
}

func TestTableCreateWithIntegerSigned64NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
}
 
