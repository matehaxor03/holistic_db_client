package validate

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
	ValidatePortNumber func(port_number string) ([]error)
	ValidateDomainName func(domain_name string) ([]error)


	GetValidateGrantFunc func() (*func(string) []error)
	GetValidateTableNameFilterAllFunc func() (*func(string) []error)
	GetValidateDatabaseNameFilterAllFunc func() (*func(string) []error)
	GetValidateBranchNameFunc func() (*func(string) []error)
	GetValidateCharacterSetFunc func() (*func(string) []error)
	GetValidateCollateFunc func() (*func(string) []error)
	GetValidateColumnNameFunc func() (*func(string) []error)
	GetValidateDatabaseNameFunc func() (*func(string) []error)
	GetValidateDatabaseReservedWordFunc func() (*func(string) []error)
	GetValidateDomainNameFunc func() (*func(string) []error)
	GetValidatePortNumberFunc func() (*func(string) []error)
	GetValidateTableNameFunc func() (*func(string) []error)
	GetValidateUsernameFunc func() (*func(string) []error)
	GetValidateRepositoryNameFunc func() (*func(string) []error)
	GetValidateRepositoryAccountNameFunc func() (*func(string) []error)
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

	valid_grant_words := NewGrantNameWhitelist()

	x := Validator {
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
		ValidatePortNumber: func(port_number string) ([]error) {
			return valid_port_number_characters.ValidatePortNumber(port_number)
		},
		ValidateDomainName: func(domain_name string) ([]error) {
			return valid_domain_name_characters.ValidateDomainName(domain_name)
		},

		GetValidateDatabaseNameFilterAllFunc: func() (*func(string) []error) {
			return valid_grant_words.GetValidateDatabaseNameFilterAllFunc()
		},
		GetValidateTableNameFilterAllFunc: func() (*func(string) []error) {
			return valid_grant_words.GetValidateTableNameFilterAllFunc()
		},
		GetValidateGrantFunc: func() (*func(string) []error) {
			return valid_grant_words.GetValidateGrantFunc()
		},
		GetValidateBranchNameFunc: func() (*func(string) []error) {
			return valid_branch_name_characters.GetValidateBranchNameFunc()
		},
		GetValidateCharacterSetFunc: func() (*func(string) []error) {
			return valid_character_set_words.GetValidateCharacterSetFunc()
		},
		GetValidateCollateFunc: func() (*func(string) []error) {
			return valid_collate_words.GetValidateCollateFunc()
		},
		GetValidateColumnNameFunc: func() (*func(string) []error) {
			return valid_column_name_characters.GetValidateColumnNameFunc()
		},
		GetValidateDatabaseNameFunc: func() (*func(string) []error) {
			return valid_database_name_characters.GetValidateDatabaseNameFunc()
		},
		GetValidateDatabaseReservedWordFunc: func() (*func(string) []error) {
			return database_reserved_words_blacklist.GetValidateDatabaseReservedWordFunc()
		},
		GetValidateDomainNameFunc: func() (*func(string) []error) {
			return valid_domain_name_characters.GetValidateDomainNameFunc()
		},
		GetValidatePortNumberFunc: func() (*func(string) []error) {
			return valid_port_number_characters.GetValidatePortNumberFunc()
		},
		GetValidateRepositoryAccountNameFunc: func() (*func(string) []error) {
			return valid_repository_account_name_characters.GetValidateRepositoryAccountNameFunc()
		},
		GetValidateRepositoryNameFunc: func() (*func(string) []error) {
			return valid_repository_name_characters.GetValidateRepositoryNameFunc()
		},
		GetValidateTableNameFunc: func() (*func(string) []error) {
			return valid_table_name_characters.GetValidateTableNameFunc()
		},
		GetValidateUsernameFunc: func() (*func(string) []error) {
			return valid_username_characters.GetValidateUsernameFunc()
		},
		

	}

	return &x
}
