package tests
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned64ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_64"
}

func GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_64_not_mandatory"
}

func GetTestTableIntegerUnsigned64ColumnName() string {
	return "integer_unsigned_64_column"
}

func GetTestTableIntegerUnsigned64ColumnNameNotMandatory() string {
	return "integer_unsigned_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned64Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned64ColumnName(): json.Map {"type": "uint64"}}
}

func GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned64ColumnNameNotMandatory(): json.Map {"type": "*uint64"}}
}

func TestTableCreateWithIntegerUnsigned64Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned64ColumnName(), GetTestSchemaWithIntegerUnsigned64Column())
}

func TestTableCreateWithIntegerUnsigned64NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
}
 
