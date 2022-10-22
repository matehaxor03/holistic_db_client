package tests
 
import (
    "testing"
	"strings"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableName() string {
	return "holistic_test_table"
}

func GetTestTablePrimaryKeyName() string {
	return "test_table_id"
}

func GetTestTablePrimaryKeyName2() string {
	return "test_table_id2"
}

func GetTestSchema() class.Map {
	return class.Map {GetTestTablePrimaryKeyName(): class.Map {"type": "uint64", "auto_increment": true, "primary_key": true}}
}

func GetTestSchemaColumn() class.Map {
	return class.Map {"type": "uint64", "auto_increment": true, "primary_key": true}
}

func GetTestSchemaColumnPrimaryKeyAutoIncrement() class.Map {
	return class.Map {"type": "uint64", "auto_increment": true, "primary_key": true}
}

func GetTestTable(t *testing.T) (*class.Table) {
	var errors []error

	database := GetTestDatabase(t)
	database_create_errors := database.Create()
	if database_create_errors != nil {
		t.Error(errors)
		return nil
	}

	use_database_errors := database.UseDatabase() 
	if use_database_errors != nil {
		errors = append(errors, use_database_errors...)
	}

	table, table_errors := class.NewTable(database, GetTestTableName(), GetTestSchema())
	if table_errors != nil {
		errors = append(errors, table_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	table_delete_errors := table.DeleteIfExists()
	if table_delete_errors != nil {
		errors = append(errors, table_delete_errors...)
	}

	if len(errors) > 0 {
		t.Error(errors)
		return nil
	}

	return table
}
 
func TestTableCreate(t *testing.T) {
	table := GetTestTable(t)

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableDelete(t *testing.T) {
	table := GetTestTable(t)

    table.Create()
	table_errors := table.Delete()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableExistsTrue(t *testing.T) {
	table := GetTestTable(t)

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
	table := GetTestTable(t)

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
	table := GetTestTable(t)

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
	table := GetTestTable(t)
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

func TestTableCannotSetTableNameWithBlackListName(t *testing.T) {
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()
	table := GetTestTable(t)

	for blacklist_database_name := range blacklist_map {
		set_table_name_errors := table.SetTableName(blacklist_database_name)
		
		if set_table_name_errors == nil {
			t.Errorf("SetTableName should return error when table_name is blacklisted")
		}

		table_name := table.GetTableName()
		if table_name == blacklist_database_name {
			t.Errorf("table_name was updated to the blacklisted table_name")
		}

		if table_name != GetTestTableName() {
			t.Errorf("table_name is '%s' and should be '%s'", table_name,  GetTestTableName())
		}
	}
}


func TestTableCannotCreateWithBlackListName(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		table, create_table_errors := class.NewTable(database, blacklist_table_name, GetTestSchema())
		
		if create_table_errors == nil {
			t.Errorf("NewTable should return error when table_name is blacklisted")
		}

		if table != nil {
			t.Errorf("NewTable should be nil when table_name is blacklisted")
		}
	}
}

func TestTableCannotCreateWithBlackListNameUppercase(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		table, create_table_errors := class.NewTable(database, strings.ToUpper(blacklist_table_name), GetTestSchema())
		
		if create_table_errors == nil {
			t.Errorf("NewTable should return error when table_name is blacklisted")
		}

		if table != nil {
			t.Errorf("NewTable should be nil when table_name is blacklisted")
		}
	}
}


func TestTableCannotCreateWithBlackListNameLowercase(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	blacklist_map := class.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name := range blacklist_map {
		table, create_table_errors := class.NewTable(database, strings.ToLower(blacklist_table_name), GetTestSchema())
		
		if create_table_errors == nil {
			t.Errorf("NewTable should return error when table_name is blacklisted")
		}

		if table != nil {
			t.Errorf("NewTable should be nil when table_name is blacklisted")
		}
	}
}

func TestTableCanCreateWithWhiteListCharacters(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	valid_characters_map := class.GetTableNameValidCharacters()

	for valid_character := range valid_characters_map {
		table, create_table_errors := class.NewTable(database, valid_character + valid_character, GetTestSchema())
		
		if create_table_errors != nil {
			t.Errorf("NewTable should not return error when table_name is whitelisted")
		}

		if table == nil {
			t.Errorf("NewTable should be not be nil when table_name is whitelisted")
		}
	}
}

func TestTableCannotCreateWithNonWhiteListCharacters(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	non_whitelist_map := class.Map{"(":nil, ")":nil}

	for invalid_character := range non_whitelist_map {
		table, create_table_errors := class.NewTable(database, invalid_character + invalid_character, GetTestSchema())
		
		if create_table_errors == nil {
			t.Errorf("NewTable should return error when table_name is non-whitelisted")
		}

		if table != nil {
			t.Errorf("NewTable should be nil when table_name is non-whitelisted")
		}
	}
}

func TestTableCannotCreateWithWhiteListCharactersIfTableNameLength1(t *testing.T) {
	database := GetTestDatabase(t)
	database.Create()
	
	valid_characters_map := class.GetTableNameValidCharacters()

	for valid_character := range valid_characters_map {
		table, create_table_errors := class.NewTable(database, valid_character, GetTestSchema())
		
		if create_table_errors == nil {
			t.Errorf("NewTable should return error when table_name is whitelisted but has length 1")
		}

		if table != nil {
			t.Errorf("NewTable should be nil when table_name is whitelisted but has length 1")
		}
	}
}