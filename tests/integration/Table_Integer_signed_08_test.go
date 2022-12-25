package integration
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned08ColumnName() string {
	return "holistic_test_table_with_integer_signed_08"
}

func GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_08_not_mandatory"
}

func GetTestTableIntegerSigned08ColumnName() string {
	return "integer_signed_08_column"
}

func GetTestTableIntegerSigned08ColumnNameNotMandatory() string {
	return "integer_signed_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned08Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnName(): json.Map {"type": "int8"}}
}

func GetTestSchemaWithIntegerSigned08NotMandatoryColumn() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnNameNotMandatory(): json.Map {"type": "*int8"}}
}

func TestTableCreateWithIntegerSigned08Column(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned08ColumnName(), GetTestSchemaWithIntegerSigned08Column())
}

func TestTableCreateWithIntegerSigned08NotMandatoryColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
}
 