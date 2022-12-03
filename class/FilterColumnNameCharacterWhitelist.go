package class

type ColumnNameCharacterWhitelist struct {
	GetColumnNameCharacterWhitelist func() (*Map)
}

func newColumnNameCharacterWhitelist() (*ColumnNameCharacterWhitelist) {
	column_name_character_whitelist := GetMySQLColumnNameWhitelistCharacters()

	x := ColumnNameCharacterWhitelist {
		GetColumnNameCharacterWhitelist: func() (*Map) {
			return &column_name_character_whitelist
		},
	}

	return &x
}
