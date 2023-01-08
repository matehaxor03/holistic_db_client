package validate

import(
	json "github.com/matehaxor03/holistic_json/json"
)

type Validator struct {
	ValidateTableName func(table_name string) ([]error)
	ValidateDatabaseName func(database_name string) ([]error)
	ValidateColumnName  func(column_name string) ([]error)
	ValidateCollate  func(collate string) ([]error)
	ValidateCharacterSet  func(character_set string) ([]error)
	ValidateUsername func(username string) ([]error)
	ValidateBranchName func(branch_name string) ([]error)
	ValidateRepositoryName func(repository_name string) ([]error)
	ValidateRepositoryAccountName func(repository_account_name string) ([]error)



	GetDatabaseReservedWordsBlackList func() *json.Map 
	GetDatabaseNameWhitelistCharacters func() *json.Map
	GetColumnNameCharacterWhitelist func() *json.Map
	GetTableNameCharacterWhitelist func() *json.Map
	
	GetUsernameCharacterWhitelist func() *json.Map 
	GetBranchNameCharacterWhitelist func() *json.Map 
	GetRepositoryNameCharacterWhitelist func() *json.Map 
	GetRepositoryAccountNameCharacterWhitelist func() *json.Map 

	GetDomainNameCharacterWhitelist func() *json.Map 
	GetPortNumberCharacterWhitelist func() *json.Map 

	GetCharacterSetWordWhitelist func() *json.Map
	GetCollateWordWhitelist func() *json.Map
}

func NewValidator() (*Validator) {
	database_reserved_words_blacklist := NewDatabaseReservedWordsBlackList()
	valid_database_name_characters := NewDatabaseNameCharacterWhitelist()
	valid_table_name_characters := NewTableNameCharacterWhitelist()
	valid_column_name_characters := NewColumnNameCharacterWhitelist()
	
	valid_username_characters := NewUsernameCharacterWhitelist()
	valid_branch_name_characters := NewBranchNameCharacterWhitelist()
	valid_repository_name_characters := NewRepositoryNameCharacterWhitelist()
	valid_repository_account_name_characters := NewRepositoryAccountNameCharacterWhitelist()
	valid_domain_name_characters := NewDomainNameCharacterWhitelist()
	valid_port_number_characters := NewPortNumberCharacterWhitelist()

	valid_character_set_words := NewCharacterSetWordWhitelist()
	valid_collate_words := NewCollateWordWhitelist()

	

	x := Validator {
		GetDatabaseNameWhitelistCharacters: func() *json.Map {
			return valid_database_name_characters.GetDatabaseNameCharacterWhitelist()
		},
		GetTableNameCharacterWhitelist: func() *json.Map {
			return valid_table_name_characters.GetTableNameCharacterWhitelist()
		},
		GetColumnNameCharacterWhitelist: func() *json.Map {
			return valid_column_name_characters.GetColumnNameCharacterWhitelist()
		},
		GetUsernameCharacterWhitelist: func() *json.Map {
			return valid_username_characters.GetUsernameCharacterWhitelist()
		},
		GetDatabaseReservedWordsBlackList: func() *json.Map {
			return database_reserved_words_blacklist.GetDatabaseReservedWordsBlackList()
		},
		GetBranchNameCharacterWhitelist: func() *json.Map {
			return valid_branch_name_characters.GetBranchNameCharacterWhitelist()
		},
		GetRepositoryNameCharacterWhitelist: func() *json.Map {
			return valid_repository_name_characters.GetRepositoryNameCharacterWhitelist()
		},
		GetRepositoryAccountNameCharacterWhitelist: func() *json.Map {
			return valid_repository_account_name_characters.GetRepositoryAccountNameCharacterWhitelist()
		},
		GetDomainNameCharacterWhitelist: func() *json.Map {
			return valid_domain_name_characters.GetDomainNameCharacterWhitelist()
		},
		GetPortNumberCharacterWhitelist: func() *json.Map {
			return valid_port_number_characters.GetPortNumberCharacterWhitelist()
		},
		GetCharacterSetWordWhitelist: func() *json.Map {
			return valid_character_set_words.GetCharacterSetWordWhitelist()
		},
		GetCollateWordWhitelist: func() *json.Map {
			return valid_collate_words.GetCollateWordWhitelist()
		},
		ValidateTableName: func(table_name string) ([]error) {
			return valid_table_name_characters.ValidateTableName(table_name)
		},
		ValidateCollate: func(collate string) ([]error) {
			return valid_collate_words.ValidateCollate(collate)
		},
		ValidateCharacterSet: func(character_set string) ([]error) {
			return valid_character_set_words.ValidateCharacterSet(character_set)
		},
		ValidateColumnName: func(column_name string) ([]error) {
			return valid_column_name_characters.ValidateColumnName(column_name)
		},
		ValidateDatabaseName: func(database_name string) ([]error) {
			return valid_database_name_characters.ValidateDatabaseName(database_name)
		},
		ValidateUsername: func(username string) ([]error) {
			return valid_username_characters.ValidateUsername(username)
		},
		ValidateBranchName: func(branch_name string) ([]error) {
			return valid_branch_name_characters.ValidateBranchName(branch_name)
		},
		ValidateRepositoryName: func(repository_name string) ([]error) {
			return valid_repository_name_characters.ValidateRepositoryName(repository_name)
		},
		ValidateRepositoryAccountName: func(repository_account_name string) ([]error) {
			return valid_repository_account_name_characters.ValidateRepositoryAccountName(repository_account_name)
		},

	}

	return &x
}
