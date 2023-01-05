package integration_test_helpers

import (
    "testing"
	"fmt"
	"sync"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	dao "github.com/matehaxor03/holistic_db_client/dao"
)

func EnsureTableIsDeleted(t *testing.T, table *dao.Table) {
	table_delete_errors := table.DeleteIfExists()
	
	if table_delete_errors != nil {
		t.Error(table_delete_errors)
		t.FailNow()
		return
	}
}

var table_count uint64 = 0
var lock_get_table_name = &sync.Mutex{}

func GetTestTableName() string {
	lock_get_table_name.Lock()
	defer lock_get_table_name.Unlock()
	table_count++
	return "holistic_test_table" + (common.GenerateRandomLetters(10, false, true)) + fmt.Sprintf("_%d", table_count)
}

func GetTestTablePrimaryKeyName() string {
	return "test_table_id"
}

func GetTestTablePrimaryKeyName2() string {
	return "test_table_id2"
}

func GetTestSchema() json.Map {
	table_schema := json.NewMapValue()
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	return table_schema
}

func GetTestSchemaColumn() json.Map {
	schema := json.NewMapValue()
	schema.SetStringValue("type", "uint64")
	schema.SetBoolValue("auto_increment", true)
	schema.SetBoolValue("primary_key",  true)
	return schema
}

func GetTestSchemaColumnPrimaryKeyAutoIncrement() json.Map {
	schema := json.NewMapValue()
	schema.SetStringValue("type", "uint64")
	schema.SetBoolValue("auto_increment", true)
	schema.SetBoolValue("primary_key", true)
	return schema
}

func GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t *testing.T, schema json.Map) (dao.Table) {
	var errors []error

	table_name := GetTestTableName()
	database := GetTestDatabaseCreated(t)
	database.DeleteTableByTableNameIfExists(table_name, true)

	table, table_errors := database.CreateTable(table_name, schema)
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return dao.Table{}
	}

	table_delete_errors := table.DeleteIfExists()
	if table_delete_errors != nil {
		errors = append(errors, table_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return  dao.Table{}
	}

	return *table
}

func GetTestTableWithTableNameAndSchema(t *testing.T, schema json.Map) (dao.Table) {
	var errors []error

	database := GetTestDatabase(t)
	
	table_name := GetTestTableName()
	table, table_errors := database.GetTableInterface(table_name, schema)
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		t.FailNow()
		return dao.Table{}
	}

	return *table
}

func GetTestTableBasic(t *testing.T) dao.Table {
	return GetTestTableWithTableNameAndSchema(t, GetTestSchema())
}

func GetTestTableBasicWithCreatedDatabase(t *testing.T) dao.Table {
	return GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, GetTestSchema())
}

func CreateTableAndVerifySchema(t *testing.T, expected_schema json.Map) {
	table := GetTestTableWithTableNameAndSchemaWithCreatedDatabase(t, expected_schema)

    table_errors := table.Create()
	if table_errors != nil {
		t.Errorf(fmt.Sprintf("error: %s", table_errors))
	} else {
		expected_schema_column_names := expected_schema.GetKeys()
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
 
