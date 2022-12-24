package integration
 
import (
    "testing"
	"strings"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func ensureTableIsDeleted(t *testing.T, table *class.Table) {
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
 
func TestTableCreate(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableDelete(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)

    table.Create()
	table_errors := table.Delete()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableExistsTrue(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)

    table.Create()
	exists, exists_errors := table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 
}

func TestTableExistsFalse(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)

	exists, exists_errors := table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 
}

func TestTableCreateWithExists(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)

	exists, exists_errors := table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}

	exists, exists_errors = table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 
}

func TestTableDeleteWithExists(t *testing.T) {
	table := GetTestTableBasicWithCreatedDatabase(t)
	table.Create()

	exists, exists_errors := table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if !(*exists) {
		t.Errorf("exists is 'false' when it should be 'true'")
	} 

    table.Delete()

	exists, exists_errors = table.Exists()
	if exists_errors != nil {
		t.Error(exists_errors)
	} 

	if exists == nil {
		t.Errorf("exists is nil")
	} 

	if (*exists) {
		t.Errorf("exists is 'true' when it should be 'false'")
	} 
}

func TestTableCanSetTableNameWithBlackListName(t *testing.T) {
	t.Parallel()
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()
	table := GetTestTableBasic(t)

	for blacklist_table_name := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		set_table_name_errors := table.SetTableName(blacklist_table_name)
		
		if set_table_name_errors != nil {
			t.Errorf(fmt.Sprintf("SetTableName should not return error when table_name is blacklisted %s", set_table_name_errors))
		}

		table_name, table_name_errors := table.GetTableName()
		if table_name_errors != nil {
			t.Errorf(fmt.Sprintf("%s", table_name_errors))
		}
		

		if table_name != blacklist_table_name {
			t.Errorf("table_name was not updated to the blacklisted table_name")
		}

		if table_name != blacklist_table_name {
			t.Errorf("table_name is '%s' and should be '%s'", table_name, blacklist_table_name)
		}
	}
}


func TestTableCanCreateWithBlackListName(t *testing.T) {
	t.Parallel()
	database := GetTestDatabase(t)
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(blacklist_table_name, GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("error: database.GetTableInterface sohuld not have errors")
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface table should not be nil")
		}
	}
}

func TestTableCanCreateWithBlackListNameUppercase(t *testing.T) {
	t.Parallel()
	database := GetTestDatabase(t)
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(strings.ToUpper(blacklist_table_name), GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("error: database.GetTableInterface should not return error was expected to have errors")
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface should not be nil table was not nil")
		}
	}
}


func TestTableCanCreateWithBlackListNameLowercase(t *testing.T) {
	t.Parallel()
	database := GetTestDatabase(t)
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(strings.ToLower(blacklist_table_name), GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf(fmt.Sprintf("error: database.GetTableInterface should not have errors was expected to have errors %s", get_table_interface_errors))
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface table should not be was not nil")
		}
	}
}

func TestTableCanCreateWithWhiteListCharacters(t *testing.T) {
	database := GetTestDatabaseCreated(t)
	valid_characters_map := class.GetMySQLTableNameWhitelistCharacters()

	for valid_character := range valid_characters_map {
		table, get_table_interface_errors := database.GetTableInterface("a" + valid_character + "a", GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("database.GetTableInterface had errors %s", get_table_interface_errors)
		} else {
			ensureTableIsDeleted(t, table)
			create_table_errors := table.Create()
			
			if create_table_errors != nil {
				t.Errorf("table.Create should not return error when table_name is whitelisted but has length 2. errors: %s", create_table_errors)
			}
		}
	}
}

func TestTableCannotCreateWithNonWhiteListCharacters(t *testing.T) {
	t.Parallel()
	database := GetTestDatabase(t)
	non_whitelist_map := json.Map{"(":nil, ")":nil}

	for invalid_character := range non_whitelist_map {
		table, get_table_interface_errors := database.GetTableInterface(invalid_character + invalid_character, GetTestSchema())

		if get_table_interface_errors == nil {
			t.Errorf("database.GetTableInterface was expected to have errors")
		}

		if table != nil {
			t.Errorf("database.GetTableInterface table was not nil")
		}
	}
}

func TestTableCannotCreateWithWhiteListCharactersIfTableNameLength1(t *testing.T) {
	t.Parallel()
	database := GetTestDatabase(t)
	valid_characters_map := class.GetMySQLTableNameWhitelistCharacters()

	for valid_character := range valid_characters_map {
		table, get_table_interface_errors := database.GetTableInterface(valid_character, GetTestSchema())

		if get_table_interface_errors == nil {
			t.Errorf("database.GetTableInterface was expected to have errors")
		}

		if table != nil {
			t.Errorf("database.GetTableInterface table was not nil")
		}
	}
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