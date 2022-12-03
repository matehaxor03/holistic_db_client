package class

type DatabaseNameCharacterWhitelist struct {
	GetDatabaseNameCharacterWhitelist func() (*Map)
}

func newDatabaseNameCharacterWhitelist() (*DatabaseNameCharacterWhitelist) {
	database_name_character_whitelist := GetMySQLDatabaseNameWhitelistCharacters()

	x := DatabaseNameCharacterWhitelist {
		GetDatabaseNameCharacterWhitelist: func() (*Map) {
			return &database_name_character_whitelist
		},
	}

	return &x
}
