package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithStringColumn() string {
	return "holistic_test_table_with_string"
}

func GetTestTableNameWithStringColumnNotMandatory() string {
	return "holistic_test_table_with_string_not_mandatory"
}

func GetTestTableStringColumnName() string {
	return "string_column"
}

func GetTestTableStringColumnNameNotMandatory() string {
	return "string_column_not_mandatory"
}

func GetTestSchemaWithStringColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnName(): class.Map {"type": "string", "max_length":100}}
}

func GetTestSchemaWithStringColumnNotMandatory() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnNameNotMandatory(): class.Map {"type": "*string", "max_length":100}}
}

func TestTableCreateWithStringColumn(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableStringColumnName(), GetTestSchemaWithStringColumn())
}

func TestTableCreateWithStringColumnNotMandatory(t *testing.T) {
	CreateTableAndVerifySchema(t, GetTestTableNameWithStringColumnNotMandatory(), GetTestSchemaWithStringColumnNotMandatory())
}
 
