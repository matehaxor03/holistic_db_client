package integration
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned16ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_16"
}

func GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_16_not_mandatory"
}

func GetTestTableIntegerUnsigned16ColumnName() string {
	return "integer_unsigned_16_column"
}

func GetTestTableIntegerUnsigned16ColumnNameNotMandatory() string {
	return "integer_unsigned_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned16Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned16ColumnName(): json.Map {"type": "uint16"}}
}

func GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned16ColumnNameNotMandatory(): json.Map {"type": "*uint16"}}
}

func TestTableCreateWithIntegerUnsigned16Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned16ColumnName(), GetTestSchemaWithIntegerUnsigned16Column())

}

func TestTableCreateWithIntegerUnsigned16NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory())
}
 
