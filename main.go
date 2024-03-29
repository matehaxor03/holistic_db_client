package main

// mkdir /home/somefolder; mount /dev/dev4 /home/somefolder
//diskutil erasevolume HFS+ "ramdisk" `hdiutil attach -nomount ram://{2048 * {mb}}`

import (
	"fmt"
	dao "github.com/matehaxor03/holistic_db_client/dao"
	db_client "github.com/matehaxor03/holistic_db_client/db_client"
	common "github.com/matehaxor03/holistic_common/common"
	json "github.com/matehaxor03/holistic_json/json"
	"os"
	"sort"
	"strings"
	"bufio"
	"unicode"
)

func main() {
	const CLS_USER string = "username"
	const CLS_HOST string = "host_name"
	const CLS_PORT string = "port_number"
	const CLS_COMMAND string = "command"
	const CLS_CLASS string = "struct"
	//const CLS_IF_EXISTS string = "if_exists"
	//const CLS_IF_NOT_EXISTS string = "if_not_exists"
	const CLS_DATABASE_NAME string = "database_name"
	const CLS_CHARACTER_SET string = "character_set"
	const CLS_COLLATE string = "collate"

	// User
	const CLS_USER_USERNAME string = "user_username"
	const CLS_USER_PASSWORD string = "user_password"
	const CLS_USER_DOMAIN_NAME string = "user_domain_name"

	var COMMAND_CREATE = "CREATE"
	var COMMAND_TEST_DATABASE_NAME_WHITELIST = "TEST_DATABASE_NAME_WHITELIST"
	var COMMAND_TEST_TABLE_NAME_WHITELIST = "TEST_TABLE_NAME_WHITELIST"
	var COMMAND_TEST_COLUMN_NAME_WHITELIST = "TEST_COLUMN_NAME_WHITELIST"
	var COMMAND_GENERATE_KEYWORD_AND_RESERVED_WORDS_BLACKLIST = "GENERATE_KEYWORD_AND_RESERVED_WORDS_BLACKLIST"

	var DATABASE_CLASS = "DATABASE"
	var USER_CLASS = "USER"

	context := db_client.NewContext()
	client_manager, client_manager_errors := dao.NewClientManager()
	if client_manager_errors != nil {
		context.LogErrors(client_manager_errors)
		os.Exit(1)
	}

	params, errors := getParams(os.Args[1:])
	if context.HasErrors() {
		os.Exit(1)
	}

	database_username_value_ptr, database_username_found := params[CLS_USER]
	database_username := ""
	if database_username_value_ptr != nil && database_username_found {
		database_username = *database_username_value_ptr
	}

	database_name_ptr, database_name_found := params[CLS_DATABASE_NAME]
	database_name := ""
	if database_name_ptr != nil && database_name_found {
		database_name = *database_name_ptr
	}
	
	character_set, _ := params[CLS_CHARACTER_SET]
	collate, _ := params[CLS_COLLATE]

	user_username, user_username_found := params[CLS_USER_USERNAME]
	user_password, user_password_found := params[CLS_USER_PASSWORD]
	user_domain_name, user_domain_name_found := params[CLS_USER_DOMAIN_NAME]

	command_pt, command_found := params[CLS_COMMAND]
	command_value := ""
	if command_pt != nil && command_found {
		command_value = strings.ToUpper(*command_pt)
	}

	class_pt, class_found := params[CLS_CLASS]
	class_value := ""
	if class_pt != nil && class_found {
		class_value = strings.ToUpper(*class_pt)
	}

	host_name_pt, host_name_found := params[CLS_HOST]
	host_name_value := ""
	if host_name_pt != nil && host_name_found {
		host_name_value = *host_name_pt
	}

	port_number_pt, port_number_found := params[CLS_PORT]
	port_number_value := ""
	if port_number_pt != nil && port_number_found {
		port_number_value = *port_number_pt
	}

	if !host_name_found {
		context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_HOST, CLS_HOST))
	}

	if !port_number_found {
		context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_PORT, CLS_PORT))
	}


	if !database_name_found {
		context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_DATABASE_NAME, CLS_DATABASE_NAME))
	}

	if !database_username_found {
		context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_USER, CLS_USER))
	}

	if command_pt == nil || !command_found {
		context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_COMMAND, CLS_COMMAND))
	}


	if context.HasErrors() {
		os.Exit(1)
	}

	client, client_errors := client_manager.GetClient(host_name_value, port_number_value, database_name, database_username)
	if client_errors != nil {
		errors = append(errors, client_errors...)
	}

	if len(errors) > 0 {
		context.LogErrors(errors)
		os.Exit(1)
	}

	if command_value == COMMAND_CREATE {
		if class_pt == nil || !class_found {
			context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_CLASS, CLS_CLASS))
		}

		if class_value == DATABASE_CLASS {
			if database_name == "" {
				context.LogError(fmt.Errorf("error: %s is a mandatory field e.g %s=", CLS_DATABASE_NAME, CLS_DATABASE_NAME))
				os.Exit(1)
			}

			database_exists, database_exists_errors := client.DatabaseExists(database_name)
			if database_exists_errors != nil {
				context.LogErrors(database_exists_errors)
				os.Exit(1)
			}

			if database_exists == false {
				_, database_errors := client.CreateDatabase(database_name, character_set, collate)

				if database_errors != nil {
					context.LogErrors(database_errors)
					os.Exit(1)
				}
			} else {
				context.LogError(fmt.Errorf("error: database: %s exists: %t", database_name, database_exists))
				os.Exit(1)
			}
		} else if class_value == USER_CLASS {
			if user_username_found && user_password_found && user_domain_name_found {
				_, user_errors := client.CreateUser(*user_username, *user_password, *user_domain_name)

				if user_errors != nil {
					context.LogErrors(user_errors)
					os.Exit(1)
				}

			} else {
				
				if !user_username_found {
					context.LogError(fmt.Errorf("error: user_username is a mandatory field"))
				}

				if !user_password_found {
					context.LogError(fmt.Errorf("error: user_password is a mandatory field"))
				}

				if !user_domain_name_found {
					context.LogError(fmt.Errorf("error: user_domain_name is a mandatory field"))
				}
			}
		} else {
			fmt.Printf("struct: %s is not supported", class_value)
			os.Exit(1)
		}
	} else if command_value == COMMAND_TEST_DATABASE_NAME_WHITELIST {
		test_database_name_errors := testDatabaseName(client)
		if test_database_name_errors != nil {
			errors = append(errors, test_database_name_errors...)
		}
	} else if command_value == COMMAND_TEST_TABLE_NAME_WHITELIST {
		test_table_name_errors := testTableName(client)
		if test_table_name_errors != nil {
			errors = append(errors, test_table_name_errors...)
		}
	} else if command_value == COMMAND_TEST_COLUMN_NAME_WHITELIST {
		test_column_name_errors := testColumnName(client)
		if test_column_name_errors != nil {
			errors = append(errors, test_column_name_errors...)
		}
	} else if command_value == COMMAND_GENERATE_KEYWORD_AND_RESERVED_WORDS_BLACKLIST {
		generate_errors := generateKeywordAndReservedWordsBlacklist(client)
		if generate_errors != nil {
			errors = append(errors, generate_errors...)
		}
	} else {
		errors = append(errors, fmt.Errorf("error: command: %s is not supported", command_value))
	}

	if context.HasErrors() {
		os.Exit(1)
	}

	os.Exit(0)
}

func generateKeywordAndReservedWordsBlacklist(client *dao.Client) []error {
	var errors []error
	invalid_strings := map[string]bool{}
	
	package_name := "validation_constants"
	filename := fmt.Sprintf("./%s/MySQLKeywordsAndReservedWordsBlacklist.go", package_name)
	raw_text_filename := fmt.Sprintf("./%s/MySQLKeywordsAndReservedWordsBlacklist.txt", package_name)
	method_name := "GetMySQLKeywordsAndReservedWordsInvalidWords()"

	text_file, text_file_error := os.OpenFile(raw_text_filename, os.O_RDWR, 0600)
	if text_file_error != nil {
		errors = append(errors, text_file_error)
	}
	defer text_file.Close()

	if len(errors) > 0 {
		return errors
	}

	scanner := bufio.NewScanner(text_file)

	for scanner.Scan() {
		current_value := strings.ToUpper(scanner.Text())
		current_value = strings.TrimSpace(current_value)
		if current_value != "" {
			parts := strings.Split(current_value, " ")
			invalid_string := parts[0]
			invalid_strings[invalid_string] = true

			if strings.HasSuffix(invalid_string, ";") {
				invalid_strings[invalid_string[:len(invalid_string)-1]] = true
			}
		}
    }

    if scanner_error := scanner.Err(); scanner_error != nil {
        errors = append(errors, scanner_error)
    }

	if len(errors) > 0 {
		return errors
	}

	validation_map_errors := createMapValidationKeysStrings(filename, package_name, method_name, invalid_strings)
	if validation_map_errors != nil {
		return validation_map_errors
	}
	
	return nil
}

func basicFilter(value rune) bool {
	string_value := string(value)
	
	if len(string_value) != 1 {
		return false
	}

	if strings.TrimSpace(string_value) == "" {
		return false
	}

	if unicode.IsControl(value) {
		return false
	}

	if unicode.IsSpace(value) {
		return false
	}


	return true
}

func testDatabaseName(client *dao.Client) []error {
	var errors []error
	
	var percent_completed float64
	var current_value uint64
	var max_value uint64

	current_value = 0
	max_value = 127

	package_name := "db_client"

	valid_runes := map[uint64]bool{}
	filename_whitelist := fmt.Sprintf("./%s/MySQLDatabaseNameWhitelistCharacters.go", package_name)
	method_name_whitelist := "GetMySQLDatabaseNameWhitelistCharacters()"


	for current_value <= max_value {
		percent_completed = (float64(current_value) / float64(max_value)) * 100.0
		percent_completed_string_value := fmt.Sprintf("%.2f", percent_completed) + "%"
		current_rune := rune(current_value)
		string_value := string(current_rune)


		if !basicFilter(current_rune) {
			fmt.Println(fmt.Sprintf("value was filtered by basic filter as invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		// double it due to defect in mysql with database names i or I
		string_value = "aaaaaa" + string_value + "aaaaaaaaa" 

		database_exists, database_exists_errors := client.DatabaseExists(string_value)
		if database_exists_errors != nil {
			fmt.Println(fmt.Sprintf("client.DatabaseExists: invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		if database_exists {
			database_deleted_errors := client.DeleteDatabase(string_value)
			if database_deleted_errors != nil {
				fmt.Println(fmt.Sprintf("client.DeleteDatabase: invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
				errors = append(errors, database_deleted_errors...)
			}
		}

		_, database_errors := client.CreateDatabase(string_value, nil, nil)
		if database_errors != nil {
			fmt.Println(fmt.Sprintf("client.CreateDatabase: invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			fmt.Println(database_errors)
		} else {
			fmt.Println(fmt.Sprintf("valid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			valid_runes[current_value] = true
			database_deleted_errors := client.DeleteDatabase(string_value)
			if database_deleted_errors != nil {
				errors = append(errors, database_deleted_errors...)
			}
		}
		current_value += 1
	}

	if len(errors) > 0 {
		return errors
	}

	validation_whitelist_map_errors := createMapValidationKeys(filename_whitelist, package_name, method_name_whitelist, valid_runes)
	if validation_whitelist_map_errors != nil {
		errors = append(errors, validation_whitelist_map_errors...)
	}

	if len(errors) > 0 {
		return errors
	}
	
	return nil
}

func testTableName(client *dao.Client) []error {
	var errors []error
	valid_runes := map[uint64]bool{}
	var percent_completed float64
	var current_value uint64
	var max_value uint64

	current_value = 0
	max_value = 127

	package_name := "db_client"
	filename := fmt.Sprintf("./%s/MySQLTableNameCharacterWhitelist.go", package_name)
	method_name := "GetMySQLTableNameWhitelistCharacters()"
	database_name := "holistic_test"

	database_exists, database_exists_errors := client.DatabaseExists(database_name)
	if database_exists_errors != nil {
		return database_exists_errors
	}

	if database_exists {
		database_deleted_errors := client.DeleteDatabase(database_name)
		if database_deleted_errors != nil {
			return database_deleted_errors
		}
	}

	database, database_errors := client.CreateDatabase(database_name, nil, nil)
	if database_errors != nil {
		return database_errors
	}

	client.SetDatabase(database)
	
	for current_value <= max_value {
		percent_completed = (float64(current_value) / float64(max_value)) * 100.0
		percent_completed_string_value := fmt.Sprintf("%.2f", percent_completed) + "%"
		current_rune := rune(current_value)
		string_value := string(current_rune)

		if !basicFilter(current_rune) {
			fmt.Println(fmt.Sprintf("value was filtered by basic filter as invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		// double it due to defect in mysql with database names i or I
		string_value = "aaaaaa" + string_value + "aaaaaaaaa" 
		schema := json.NewMapValue()
		schema_column := json.NewMap()
		schema_column.SetStringValue("type", "uint64")
		schema_column.SetBoolValue("primary_key", true)
		schema_column.SetBoolValue("auto_increment", true)
		schema.SetMap("id", schema_column)

		table, table_errors := database.CreateTable(string_value, schema)
		if table_errors != nil {
			fmt.Println(fmt.Sprintf("invalid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			fmt.Println(table_errors)
		} else {
			fmt.Println(fmt.Sprintf("valid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			valid_runes[current_value] = true
			table_deleted_errors := table.Delete()
			if table_deleted_errors != nil {
				errors = append(errors, table_deleted_errors...)
			}
		}
		current_value += 1
	}

	database_deleted_errors := client.DeleteDatabase(database_name)
	if database_deleted_errors != nil {
		errors = append(errors, database_deleted_errors...)
	}

	if len(errors) > 0 {
		return errors
	}


	validation_map_errors := createMapValidationKeys(filename, package_name, method_name, valid_runes)
	if validation_map_errors != nil {
		return validation_map_errors
	}

	return nil
}

func testColumnName(client *dao.Client) []error {
	var errors []error
	valid_runes := map[uint64]bool{}
	var percent_completed float64
	var current_value uint64
	var max_value uint64

	current_value = 0
	max_value = 127

	package_name := "db_client"
	filename := fmt.Sprintf("./%s/MySQLColumnNameWhitelistCharacters.go", package_name)
	method_name := "GetMySQLColumnNameWhitelistCharacters()"
	database_name := "holistic_test"

	database_exists, database_exists_errors := client.DatabaseExists(database_name)
	if database_exists_errors != nil {
		return database_exists_errors
	}

	if database_exists {
		database_deleted_errors := client.DeleteDatabase(database_name)
		if database_deleted_errors != nil {
			return database_deleted_errors
		}
	}

	database, database_errors := client.CreateDatabase(database_name, nil, nil)
	if database_errors != nil {
		return database_errors
	}

	client.SetDatabase(database)

	for current_value <= max_value {
		percent_completed = (float64(current_value) / float64(max_value)) * 100.0
		percent_completed_string_value := fmt.Sprintf("%.2f", percent_completed) + "%"
		current_rune := rune(current_value)
		string_value := string(current_rune)

		if !basicFilter(current_rune) {
			fmt.Println(fmt.Sprintf("value was filtered by basic filter as invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		// double it due to defect in mysql with database names i or I
		string_value = "aaaaaa" + string_value + "aaaaaaaaa" 
		schema := json.NewMapValue()
		schema_column := json.NewMap()
		schema_column.SetStringValue("type", "uint64")
		schema_column.SetBoolValue("primary_key", true)
		schema_column.SetBoolValue("auto_increment", true)
		schema.SetMap(string_value, schema_column)		
		
		table_name := common.GenerateRandomLetters(10, false, true)

		table, table_errors := database.CreateTable(table_name, schema)
		if table_errors != nil {
			fmt.Println(fmt.Sprintf("invalid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			fmt.Println(table_errors)
		} else {
			fmt.Println(fmt.Sprintf("valid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			valid_runes[current_value] = true
			table_deleted_errors := table.Delete()
			if table_deleted_errors != nil {
				errors = append(errors, table_deleted_errors...)
			}
		}
		current_value += 1
	}

	database_deleted_errors := client.DeleteDatabase(database_name)
	if database_deleted_errors != nil {
		errors = append(errors, database_deleted_errors...)
	}

	if len(errors) > 0 {
		return errors
	}


	validation_map_errors := createMapValidationKeys(filename, package_name, method_name, valid_runes)
	if validation_map_errors != nil {
		return validation_map_errors
	}

	return nil
}

func createMapValidationKeys(filename string, package_name string, method_name string, valid_runes map[uint64]bool) []error{
	var errors []error
	valid_rune_file, valid_rune_file_error := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if valid_rune_file_error != nil {
		errors = append(errors, valid_rune_file_error)
	}
	defer valid_rune_file.Close()

	if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("package %s\nfunc %s Map {\nreturn Map{\n", package_name, method_name)); valid_error != nil {
		errors = append(errors, valid_error)
		return errors
	}

	sorted_keys := make([]uint64, 0, len(valid_runes))
	for k := range valid_runes {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })

	length := len(valid_runes)
	for index, key := range sorted_keys {
		key_string := string(rune(key))
		if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("    \"%s\": nil", key_string)); valid_error != nil {
			errors = append(errors, valid_error)
			return errors
		}

		if uint64(index) < uint64(length-1) {
			if _, valid_error := valid_rune_file.WriteString(",\n"); valid_error != nil {
				errors = append(errors, valid_error)
				return errors
			}
		}
	}

	if _, valid_error := valid_rune_file.WriteString("}\n}"); valid_error != nil {
		errors = append(errors, valid_error)
		return errors
	}


	return nil
}

func createMapValidationKeysStrings(filename string, package_name string, method_name string, valid_runes map[string]bool) []error{
	var errors []error
	valid_rune_file, valid_rune_file_error := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if valid_rune_file_error != nil {
		errors = append(errors, valid_rune_file_error)
	}
	defer valid_rune_file.Close()

	if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("package %s\n\nfunc %s map[string]interface{} {\nvalue := make(map[string]interface{})\n", package_name, method_name)); valid_error != nil {
		errors = append(errors, valid_error)
		return errors
	}

	sorted_keys := make([]string, 0, len(valid_runes))
	for k := range valid_runes {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)


	length := len(valid_runes)
	for index, key := range sorted_keys {
		if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("    value[\"%s\"] = nil", key)); valid_error != nil {
			errors = append(errors, valid_error)
			return errors
		}

		if uint64(index) < uint64(length-1) {
			if _, valid_error := valid_rune_file.WriteString("\n"); valid_error != nil {
				errors = append(errors, valid_error)
				return errors
			}
		}
	}

	if _, valid_error := valid_rune_file.WriteString("\nreturn value \n}"); valid_error != nil {
		errors = append(errors, valid_error)
		return errors
	}


	return nil
}

func getParams(params []string) (map[string]*string, []error) {
	errors := []error{}

	m := make(map[string]*string)
	for _, value := range params {
		if !strings.Contains(value, "=") {
			m[value] = nil
			continue
		}

		results := strings.SplitN(value, "=", 2)
		if len(results) != 2 {
			errors = append(errors, fmt.Errorf("error: invalid param found: %s must be in the format {paramName}={paramValue}", value))
			continue
		}
		m[results[0]] = &results[1]
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return m, nil
}
