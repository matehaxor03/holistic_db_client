package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func main() {
	errors := []error{}
	const CLS_USER string = "username"
	const CLS_HOST string = "host_name"
	const CLS_PORT string = "port_number"
	const CLS_COMMAND string = "command"
	const CLS_CLASS string = "class"
	const CLS_IF_EXISTS string = "if_exists"
	const CLS_IF_NOT_EXISTS string = "if_not_exists"
	const CLS_DATABASE_NAME string = "database_name"
	const CLS_CHARACTER_SET string = "character_set"
	const CLS_COLLATE string = "collate"

	// User 
	const CLS_USER_USERNAME string = "user_username"
	const CLS_USER_PASSWORD string = "user_password"
	const CLS_USER_DOMAIN_NAME string = "user_domain_name"



	var CREATE_COMMAND = "CREATE"
	var TEST_DATABASE_NAME_WHITELIST_COMMAND = "TEST_DATABASE_NAME_WHITELIST"
	var TEST_TABLE_NAME_WHITELIST_COMMAND = "TEST_TABLE_NAME_WHITELIST"

	
	var DATABASE_CLASS = "DATABASE"
	var USER_CLASS = "USER"

	//var IF_EXISTS string = "IF EXISTS"
	//var IF_NOT_EXISTS string = "IF NOT EXISTS"
	
	context := class.NewContext()
    params, errors := getParams(os.Args[1:])
	if errors != nil || len((errors)) > 0 {
		context.LogErrors(errors)
		os.Exit(1)
	}

	host_value, _ := params[CLS_HOST] 
	port_value, _ := params[CLS_PORT] 

	user_value, _ := params[CLS_USER] 

	database_name, _ := params[CLS_DATABASE_NAME]
	character_set, _ := params[CLS_CHARACTER_SET]
	collate, _ := params[CLS_COLLATE]

	// User
	user_username, _ := params[CLS_USER_USERNAME]
	user_password, _ := params[CLS_USER_PASSWORD]
	user_domain_name, _ := params[CLS_USER_DOMAIN_NAME]

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

	_, if_exists := params[CLS_IF_EXISTS]
	_, if_not_exists := params[CLS_IF_NOT_EXISTS]
	
	if if_exists && if_not_exists {
		context.LogError(fmt.Errorf("%s and %s cannot be used together", CLS_IF_EXISTS, CLS_IF_NOT_EXISTS))
	}

	if command_pt == nil || !command_found {
		context.LogError(fmt.Errorf("%s is a mandatory field e.g %s=", CLS_COMMAND, CLS_COMMAND))
	}

	if class_pt == nil || !class_found {
		context.LogError(fmt.Errorf("%s is a mandatory field e.g %s=", CLS_CLASS, CLS_CLASS))
	}

	if database_name == nil || *database_name == "" {
		context.LogError(fmt.Errorf("%s is a mandatory field e.g %s=", CLS_DATABASE_NAME, CLS_DATABASE_NAME))
	}

	if len(errors) > 0 {
	    context.LogErrors(errors)
		os.Exit(1)
	}

	options := make(map[string]map[string][][]string)
	if if_not_exists {
		logic_options := make(map[string][][]string)
		logic_options[command_value] = append(logic_options[command_value], class.GET_LOGIC_STATEMENT_IF_NOT_EXISTS())
		options[class.GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options
	}

	if if_exists {
		logic_options := make(map[string][][]string)
		logic_options[command_value] = append(logic_options[command_value], class.GET_LOGIC_STATEMENT_IF_EXISTS())
		options[class.GET_LOGIC_STATEMENT_FIELD_NAME()] = logic_options
	}

	host, host_errors := class.NewHost(host_value, port_value)
	if host_errors != nil {
		errors = append(errors, host_errors...)
	}
	
	client, client_errors := class.NewClient(host, user_value, nil)
	if client_errors != nil {
		errors = append(errors, client_errors...)
	}

	if len(errors) > 0 {
	    context.LogErrors(errors)
		os.Exit(1)
	}	

	if command_value == CREATE_COMMAND {
		if class_value == DATABASE_CLASS {
			database_exists, database_exists_errors := client.DatabaseExists(*database_name)
			if database_exists_errors != nil {
				context.LogErrors(database_exists_errors)
				os.Exit(1)
			}

			if if_not_exists && *database_exists == false {
				database_create_options := class.NewDatabaseCreateOptions(character_set, collate)
				_, database_errors := client.CreateDatabase(*database_name, database_create_options)
				
				if database_errors != nil {
					context.LogErrors(database_errors)
					os.Exit(1)
				}
			} else {
				context.LogError(fmt.Errorf("database: %s exists: %b", *database_name, database_exists))
				os.Exit(1)
			}
		} else if class_value == USER_CLASS {
			_, user_errors := client.CreateUser(user_username, user_password, user_domain_name, options)

			if user_errors != nil {
				context.LogErrors(user_errors)
				os.Exit(1)
			}
		} else {
			fmt.Printf("class: %s is not supported", class_value)
			os.Exit(1)
		}
	} else if command_value == TEST_DATABASE_NAME_WHITELIST_COMMAND {
		test_database_name_errors := testDatabaseName(client)
		if test_database_name_errors != nil {
			errors = append(errors, test_database_name_errors...)
		}
	} else if command_value == TEST_TABLE_NAME_WHITELIST_COMMAND {
		test_table_name_errors := testTableName(client)
		if test_table_name_errors != nil {
			errors = append(errors, test_table_name_errors...)
		}
	} else {
		fmt.Printf("command: %s is not supported", command_value)
		os.Exit(1)
	}

	if len(errors) > 0 {
	    context.LogErrors(errors)
		os.Exit(1)
	}	
	
	os.Exit(0)
}

func testTableName(client *class.Client) []error {
	var errors []error
	valid_runes := map[uint64]bool{}
	var percent_completed float64
	var current_value uint64 
	var max_value uint64

	current_value = 0
	max_value = 127

	filename := "./class/TableNameWhitelist.go"
	database_name := "holistic_test"

	database_exists, database_exists_errors := client.DatabaseExists(database_name)
	if database_exists_errors != nil {
		return database_exists_errors
	}
		
	if *database_exists {
		database_deleted_errors := client.DeleteDatabase(database_name)
		if database_deleted_errors != nil {
			return database_deleted_errors
		}
	}

	_, database_errors := client.CreateDatabase(database_name, nil)
	if database_errors != nil {
		return database_errors
	}

	database, use_database_errors := client.UseDatabaseByName(database_name)
	if use_database_errors != nil {
		return use_database_errors
	}

	for current_value <= max_value {
		percent_completed = (float64(current_value) / float64(max_value)) * 100.0
		percent_completed_string_value := fmt.Sprintf("%.2f", percent_completed) + "%"
		current_rune := rune(current_value)
		string_value := string(current_rune)

		if len(string_value) != 1 {
			fmt.Println(fmt.Sprintf("value has length != 1 invalid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		if strings.TrimSpace(string_value) == "" {
			fmt.Println(fmt.Sprintf("value is empty invalid rune for table_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		// double it due to defect in mysql with database names i or I
		string_value += string_value
		schema := class.Map{"id":class.Map{"type":"uint64", "primary_key":true, "auto_increment":true}}
		
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

	valid_rune_file, valid_rune_file_error := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if valid_rune_file_error != nil {
		errors = append(errors, valid_rune_file_error)
	}
	defer valid_rune_file.Close()


	if _, valid_error := valid_rune_file.WriteString("package class\nfunc GetTableNameValidCharacters() Map {\nreturn Map{\n"); valid_error != nil {
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
		if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("\t\"%s\":nil", string(key))); valid_error != nil {
			errors = append(errors, valid_error)
			return errors
		}

		if uint64(index) < uint64(length - 1) {
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

func testDatabaseName(client *class.Client) []error {
	var errors []error
	valid_runes := map[uint64]bool{}
	var percent_completed float64
	var current_value uint64 
	var max_value uint64

	current_value = 0
	max_value = 127

	filename := "./class/DatabaseNameWhitelist.go"

	for current_value <= max_value {
		percent_completed = (float64(current_value) / float64(max_value)) * 100.0
		percent_completed_string_value := fmt.Sprintf("%.2f", percent_completed) + "%"
		current_rune := rune(current_value)
		string_value := string(current_rune)

		if len(string_value) != 1 {
			fmt.Println(fmt.Sprintf("value has length != 1 invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		if strings.TrimSpace(string_value) == "" {
			fmt.Println(fmt.Sprintf("value is empty invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}

		// double it due to defect in mysql with database names i or I
		string_value += string_value

		database_exists, database_exists_errors := client.DatabaseExists(string_value)
		if database_exists_errors != nil {
			fmt.Println(fmt.Sprintf("client.DatabaseExists: invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
			current_value += 1
			continue
		}
		
		if *database_exists {
			database_deleted_errors := client.DeleteDatabase(string_value)
			if database_deleted_errors != nil {
				fmt.Println(fmt.Sprintf("client.DeleteDatabase: invalid rune for database_name string_value: %s rune_count: %d precent_completed: %s", string_value, current_value, percent_completed_string_value))
				errors = append(errors, database_deleted_errors...)
			}
		}

		_, database_errors := client.CreateDatabase(string_value, nil)
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

	valid_rune_file, valid_rune_file_error := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if valid_rune_file_error != nil {
		errors = append(errors, valid_rune_file_error)
	}
	defer valid_rune_file.Close()


	if _, valid_error := valid_rune_file.WriteString("package class\nfunc GetDatabaseNameValidCharacters() Map {\nreturn Map{\n"); valid_error != nil {
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
		if _, valid_error := valid_rune_file.WriteString(fmt.Sprintf("\t\"%s\":nil", string(key))); valid_error != nil {
			errors = append(errors, valid_error)
			return errors
		}

		if uint64(index) < uint64(length - 1) {
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
			errors = append(errors, fmt.Errorf("invalid param found: %s must be in the format {paramName}={paramValue}", value))
			continue
		}
		m[results[0]] = &results[1]
	}

	if len(errors) > 0 {
		return nil, errors
	}
 
	return m, nil
}
