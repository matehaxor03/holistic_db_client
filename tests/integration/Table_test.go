package integration
 
import (
    "testing"
	"strings"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreate(t *testing.T) {
	table := helper.GetTestTableBasicWithCreatedDatabase(t)

    table_errors := table.Create()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableDelete(t *testing.T) {
	table := helper.GetTestTableBasicWithCreatedDatabase(t)

    table.Create()
	table_errors := table.Delete()
	if table_errors != nil {
		t.Error(table_errors)
	}
}

func TestTableExistsTrue(t *testing.T) {
	table := helper.GetTestTableBasicWithCreatedDatabase(t)

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
	table := helper.GetTestTableBasicWithCreatedDatabase(t)

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
	table := helper.GetTestTableBasicWithCreatedDatabase(t)

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
	table := helper.GetTestTableBasicWithCreatedDatabase(t)
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
	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()
	table := helper.GetTestTableBasic(t)

	for blacklist_table_name, _  := range (blacklist_map) {
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
	database := helper.GetTestDatabase(t)
	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name, _  := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(blacklist_table_name, helper.GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("error: database.GetTableInterface sohuld not have errors")
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface table should not be nil")
		}
	}
}

func TestTableCanCreateWithBlackListNameUppercase(t *testing.T) {
	database := helper.GetTestDatabase(t)
	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name, _  := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(strings.ToUpper(blacklist_table_name), helper.GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("error: database.GetTableInterface should not return error was expected to have errors")
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface should not be nil but was nil")
		}
	}
}


func TestTableCanCreateWithBlackListNameLowercase(t *testing.T) {
	database := helper.GetTestDatabase(t)
	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_table_name, _  := range blacklist_map {
		if len(blacklist_table_name) == 1 || strings.Contains(blacklist_table_name, ";") {
			continue
		}
		
		table, get_table_interface_errors := database.GetTableInterface(strings.ToLower(blacklist_table_name), helper.GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf(fmt.Sprintf("error: database.GetTableInterface should not have errors was expected to have errors %s", get_table_interface_errors))
		}

		if table == nil {
			t.Errorf("error: database.GetTableInterface table should not be was not nil")
		}
	}
}

func TestTableCanCreateWithWhiteListCharacters(t *testing.T) {
	database := helper.GetTestDatabaseCreated(t)
	valid_characters_map := validation_constants.GetMySQLTableNameWhitelistCharacters()

	for valid_character, _ := range (valid_characters_map) {
		table, get_table_interface_errors := database.GetTableInterface("a" + valid_character + "a", helper.GetTestSchema())

		if get_table_interface_errors != nil {
			t.Errorf("database.GetTableInterface had errors %s", get_table_interface_errors)
		} else {
			helper.EnsureTableIsDeleted(t, table)
			create_table_errors := table.Create()
			
			if create_table_errors != nil {
				t.Errorf("table.Create should not return error when table_name is whitelisted but has length 2. errors: %s", create_table_errors)
			}
		}
	}
}

func TestTableCannotCreateWithNonWhiteListCharacters(t *testing.T) {
	database := helper.GetTestDatabase(t)
	non_whitelist_map := json.NewMapValue()
	non_whitelist_map.SetNil("(")
	non_whitelist_map.SetNil(")")

	for _, invalid_character := range (non_whitelist_map.GetKeys()) {
		table, get_table_interface_errors := database.GetTableInterface(invalid_character + invalid_character, helper.GetTestSchema())

		if get_table_interface_errors == nil {
			t.Errorf("database.GetTableInterface was expected to have errors")
		}

		if table != nil {
			t.Errorf("database.GetTableInterface table was not nil")
		}
	}
}

func TestTableCannotCreateWithWhiteListCharactersIfTableNameLength1(t *testing.T) {
	database := helper.GetTestDatabase(t)
	valid_characters_map := validation_constants.GetMySQLTableNameWhitelistCharacters()

	for valid_character, _  := range valid_characters_map {
		table, get_table_interface_errors := database.GetTableInterface(valid_character, helper.GetTestSchema())

		if get_table_interface_errors == nil {
			t.Errorf("database.GetTableInterface was expected to have errors")
		}

		if table != nil {
			t.Errorf("database.GetTableInterface table was not nil")
		}
	}
}