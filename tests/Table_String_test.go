package tests
 
import (
    "testing"
	//"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithStringColumn() string {
	return "holistic_test_table_with_string"
}

func GetTestTableStringColumnName() string {
	return "string_column"
}

func GetTestSchemaWithStringColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnName(): class.Map {"type": "string", "max_length":100}}
}


func TestTableCreateWithStringColumn(t *testing.T) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t,  GetTestTableNameWithStringColumn(), GetTestSchemaWithStringColumn())

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}
}
 
