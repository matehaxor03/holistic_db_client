package integration_test_helpers

import (
    "testing"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func EnsureTableIsDeleted(t *testing.T, table *class.Table) {
	table_delete_errors := table.DeleteIfExists()
	
	if table_delete_errors != nil {
		t.Error(table_delete_errors)
		t.FailNow()
		return
	}
}

func GetTestTableName() string {
	return "holistic_test_table"
}
func GetTestTablePrimaryKeyName() string {
	return "test_table_id"
}

func GetTestTablePrimaryKeyName2() string {
	return "test_table_id2"
}

func GetTestSchema() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true}}
}

func GetTestSchemaColumn() json.Map {
	return json.Map {"type": "uint64", "auto_increment": true, "primary_key": true}
}

func GetTestSchemaColumnPrimaryKeyAutoIncrement() json.Map {
	return json.Map {"type": "uint64", "auto_increment": true, "primary_key": true}
}

func GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t *testing.T, table_name string, schema json.Map) (*class.Table) {
	var errors []error

	database := GetTestDatabaseCreated(t)

	if database == nil {
		t.Error(fmt.Errorf("error: database is nil"))
		t.FailNow()
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	table, table_errors := database.CreateTable(table_name, schema)
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

func GetTestTableWithTableNameAndSchema(t *testing.T, table_name string, schema json.Map) (*class.Table) {
	var errors []error

	database := GetTestDatabase(t)

	if database == nil {
		t.Error(fmt.Errorf("error: database is nil"))
		t.FailNow()
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	table, table_errors := database.GetTableInterface(table_name, schema)
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return nil
	}

	return table
}

func GetTestTableBasic(t *testing.T) *class.Table {
	return GetTestTableWithTableNameAndSchema(t, GetTestTableName(), GetTestSchema())
}

func GetTestTableBasicWithCreatedDatabase(t *testing.T) *class.Table {
	return GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, GetTestTableName(), GetTestSchema())
}

func CreateTableAndVerifySchema(t *testing.T, table_name string, expected_schema json.Map) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, table_name, expected_schema)

    table_errors := table.Create()
	if table_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", table_errors))
	} else {
		read_errors := table.Read()
		if read_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", read_errors))
		} else {
			expected_schema_column_names := expected_schema.Keys()
			actual_schema, actual_schema_errors := table.GetSchema()
			if actual_schema_errors != nil {
				t.Errorf(fmt.Sprintf("error: %s", actual_schema_errors))
			} else if common.IsNil(actual_schema) {
				t.Errorf("error: actual schema is nil")
			} else {
				for _, expected_schema_column_name := range expected_schema_column_names {

					expected_schema_field, expected_schema_field_errors := expected_schema.GetMap(expected_schema_column_name)
					if expected_schema_field_errors != nil {
						t.Error(expected_schema_field_errors)
						continue
					} else if !expected_schema.IsMap(expected_schema_column_name) {
						t.Errorf("error: Table_test.CreateTableAndVerifySchema: %s expected schema is not a map: %s", expected_schema_column_name, expected_schema.GetType(expected_schema_column_name))
						continue
					}

					expected_schema_type, expected_schema_type_errors := expected_schema_field.GetString("type")
					if expected_schema_type_errors != nil {
						t.Errorf(fmt.Sprintf("error: %s", expected_schema_type_errors))
						continue
					} else if common.IsNil(expected_schema_type) {
						t.Errorf("error: field: %s expected_schem type is nil", expected_schema_column_name)
						continue
					}

					actual_schema_field_map, actual_schema_field_map_errors := actual_schema.GetMap(expected_schema_column_name)
					if actual_schema_field_map_errors != nil {
						t.Error(actual_schema_field_map_errors)
						continue
					} else if !actual_schema.IsMap(expected_schema_column_name) {
						t.Errorf("error: field: %s actual schema is not a map: %s", expected_schema_column_name, actual_schema.GetType(expected_schema_column_name))
						continue
					}

					actual_schema_type, actual_schema_type_errors := actual_schema_field_map.GetString("type")
					if actual_schema_type_errors != nil {
						t.Error(actual_schema_type_errors)
						continue
					} else if common.IsNil(actual_schema_type) {
						t.Errorf("error: field: %s actual_schema is nil", expected_schema_column_name)
						continue
					}

					if *expected_schema_type != *actual_schema_type {
						t.Errorf("error: schema types do not match expected: %s actual: %s", *expected_schema_type, *actual_schema_type)
					}
				}
			}
		}
		
		
	}
}
 
