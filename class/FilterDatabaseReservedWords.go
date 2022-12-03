package class

type DatabaseReservedWords struct {
	GetDatabaseReservedWords func() (*Map)
}

func newDatabaseReservedWords() (*DatabaseReservedWords) {
	database_reserved_words := GetMySQLKeywordsAndReservedWordsInvalidWords()

	x := DatabaseReservedWords{
		GetDatabaseReservedWords: func() (*Map) {
			return &database_reserved_words
		},
	}

	return &x
}
