package class

type TableNameCharacterWhitelist struct {
	GetTableNameCharacterWhitelist func() (*Map)
}

func newTableNameCharacterWhitelist() (*TableNameCharacterWhitelist) {
	table_name_character_whitelist := GetMySQLTableNameWhitelistCharacters()

	x := TableNameCharacterWhitelist {
		GetTableNameCharacterWhitelist: func() (*Map) {
			return &table_name_character_whitelist
		},
	}

	return &x
}
