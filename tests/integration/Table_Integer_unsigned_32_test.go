package integration
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
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


func GetTestSchemaWithIntegerUnsigned32Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned32ColumnName(): json.Map {"type": "uint32"}}
}

func GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableIntegerUnsigned32ColumnNameNotMandatory(): json.Map {"type": "*uint32"}}
}

func TestTableCreateWithIntegerUnsigned32Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned32ColumnName(), GetTestSchemaWithIntegerUnsigned32Column())
}

func TestTableCreateWithIntegerUnsigned32NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory(), GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
}
 
