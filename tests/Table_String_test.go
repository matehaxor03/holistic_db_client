package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithString() string {
	return "holistic_test_table_with_string"
}

func GetTestTableStringColumnName() string {
	return "string_column"
}

func GetTestSchemaStringColumn() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnName(): class.Map {"type": "string", "mandatory":true, "max_length":100}}
}

func GetTestTableWithStringColumn(t *testing.T) (*class.Table) {
	var errors []error

	database := GetTestDatabase(t)

	if database == nil {
		t.Error(fmt.Errorf("database is nil"))
		t.FailNow()
		return nil
	}

	database_create_errors := database.Create()
	if database_create_errors != nil {
		errors = append(errors, database_create_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	table, table_errors := class.NewTable(database, GetTestTableNameWithString(), GetTestSchemaStringColumn())
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	table_delete_errors := table.DeleteIfExists()
	if table_delete_errors != nil {
		errors = append(errors, table_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return table
}

func TestTableCreateWithStringColumn(t *testing.T) {
	table := GetTestTableWithStringColumn(t)

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}
}
 
