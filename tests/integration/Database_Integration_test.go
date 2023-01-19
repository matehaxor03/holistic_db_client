package integration
 
import (
    "testing"
	"strings"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
	common "github.com/matehaxor03/holistic_common/common"
	validation_constants "github.com/matehaxor03/holistic_validator/validation_constants"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestDatabaseCanSetDatabaseNameWithBlackListName(t *testing.T) {
	database := helper.GetTestDatabase(t)
	previous_database_name := database.GetDatabaseName()

	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()
	for blacklist_database_name, _ := range blacklist_map {
		
		if len(blacklist_database_name) == 1 || strings.Contains(blacklist_database_name, ";") {
			continue
		}
		
		set_database_name_errors := database.SetDatabaseName(blacklist_database_name)
		
		if set_database_name_errors != nil {
			t.Errorf(fmt.Sprintf("error: SetDatabaseName should not return error when database_name is blacklisted, %s", set_database_name_errors))
		}

		database_name := database.GetDatabaseName()

		if database_name != blacklist_database_name {
			t.Errorf("error: database_name was not updated to the blacklisted database_name")
		}

		if database_name != blacklist_database_name {
			t.Errorf("error: database_name is '%s' and should be '%s'", database_name,  previous_database_name)
		}
	}
}


func TestDatabaseCanCreateWithBlackListName(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()
	
	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name, _ := range blacklist_map {
		
		if len(blacklist_database_name) == 1 || strings.Contains(blacklist_database_name, ";") {
			continue
		}
		
		database, create_database_errors := client.GetDatabaseInterface(blacklist_database_name, &character_set, &collate)
		
		if create_database_errors != nil {
			t.Errorf(fmt.Sprintf("error: NewDatabase should not return error when database_name is blacklisted %s", create_database_errors))
		}

		if database == nil {
			t.Errorf("error: NewDatabase should not be nil when database_name is blacklisted")
		}
	}
}

func TestDatabaseCanCreateWithBlackListNameUppercase(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()

	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name, _ := range blacklist_map {
		
		if len(blacklist_database_name) == 1 || strings.Contains(blacklist_database_name, ";") {
			continue
		}

		database, create_database_errors := client.GetDatabaseInterface(strings.ToUpper(blacklist_database_name), &character_set, &collate)
		
		if create_database_errors != nil {
			t.Errorf(fmt.Sprintf("error: NewDatabase should not return error when database_name is blacklisted %s", create_database_errors))
		}

		if database == nil {
			t.Errorf("error: NewDatabase should not be nil when database_name is blacklisted")
		}
	}
}


func TestDatabaseCanCreateWithBlackListNameLowercase(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()

	blacklist_map := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	for blacklist_database_name, _  := range blacklist_map {
		
		if len(blacklist_database_name) == 1 || strings.Contains(blacklist_database_name, ";") {
			continue
		}

		database, create_database_errors :=  client.GetDatabaseInterface(strings.ToLower(blacklist_database_name), &character_set, &collate)
		
		if create_database_errors != nil {
			t.Errorf(fmt.Sprintf("error: NewDatabase should not return error when database_name is blacklisted %s", create_database_errors))
		}

		if database == nil {
			t.Errorf("error: NewDatabase not should be nil when database_name is blacklisted")
		}
	}
}

func TestDatabaseCanCreateWithWhiteListCharacters(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()

	whitelist_map := validation_constants.GetMySQLDatabaseNameWhitelistCharacters()

	for whitelist_database_character, _  := range whitelist_map {
		database, new_database_errors := client.GetDatabaseInterface("a" + whitelist_database_character + "a", &character_set, &collate)
			
		if new_database_errors != nil {
			t.Errorf(fmt.Sprintf("error: %s", new_database_errors))
		} else if database == nil {
			t.Errorf("error: NewDatabase should not be nil when database_name is whitelisted: %s", whitelist_database_character)
		} else {
			database_delete_errors := database.DeleteIfExists()
			if database_delete_errors != nil {
				t.Error(database_delete_errors)
			} else {
				create_database_errors := database.Create()
				if create_database_errors != nil {
					t.Error(create_database_errors)
				}
			}
		}
	}
}

func TestDatabaseCannotCreateWithNonWhiteListCharacters(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()

	non_whitelist_map := json.NewMap()
	non_whitelist_map.SetNil("(")
	non_whitelist_map.SetNil(")")

	for _, non_whitelist_characters := range non_whitelist_map.GetKeys() {
		database, new_database_errors := client.GetDatabaseInterface("a" + non_whitelist_characters + "a", &character_set, &collate)
		
		if new_database_errors == nil {
			t.Errorf("error: NewDatabase should return error when database_name character is non-whitelisted: %s", non_whitelist_characters)
		} else if database != nil {
			t.Errorf("error: NewDatabase should be nil when database_name is non-whitelisted: %s", non_whitelist_characters)
		}
	}
}

func TestDatabaseCannotCreateWithWhiteListCharactersIfDatabaseNameLength1(t *testing.T) {
	client := helper.GetTestClient(t)
	character_set := validation_constants.GET_CHARACTER_SET_UTF8MB4()
	collate := validation_constants.GET_COLLATE_UTF8MB4_0900_AI_CI()

	whitelist_map := validation_constants.GetMySQLDatabaseNameWhitelistCharacters()

	for whitelist_database_character, _  := range whitelist_map {
		database, new_database_errors := client.GetDatabaseInterface(whitelist_database_character, &character_set, &collate)
		
		if new_database_errors == nil {
			t.Errorf("error: NewDatabase should return error when database_name character is whitelisted and database_name is only one character long: %s", whitelist_database_character)
		} else if database != nil {
			t.Errorf("error: NewDatabase should be nil when database_name is whitelisted and is only one character long: %s", whitelist_database_character)
		}
	}
}

func TestDatabaseCanGetTableNames(t *testing.T) {
	database := helper.GetTestDatabaseCreated(t)
	table_name := helper.GetTestTableName()
	database.CreateTable(table_name, (helper.GetTestSchema()))

	table_names, tables_name_errors := database.GetTableNames()
	if tables_name_errors != nil {
		t.Error(tables_name_errors)
	} else if len(table_names) == 0 {
		t.Errorf("error: database.GetTables should return at least one table name")
	} else if !common.Contains(table_names, table_name) {
		t.Errorf("error: table: %s not found in table_names: %s", table_name, table_names)
	}
}

func TestDatabaseCanGetTables(t *testing.T) {
	database := helper.GetTestDatabaseCreated(t)
	table_name := helper.GetTestTableName()
	_, created_table_errors := database.CreateTable(table_name,  (helper.GetTestSchema()))
	if created_table_errors != nil {
		t.Errorf("error: created tables should succeed: %s", fmt.Sprintf("%s", created_table_errors))
	} else {
		tables, tables_errors := database.GetTables()
		if tables_errors != nil {
			t.Errorf("error: created tables should succeed: %s",  fmt.Sprintf("%s", tables_errors))
		} else if common.IsNil(tables) {
			t.Errorf("error: tables should not be nil")
		} else if !(len(tables) >= 0) {
			t.Errorf("error: database.GetTables should return at least one table")
		}
	}
}