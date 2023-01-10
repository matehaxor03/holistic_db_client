package validation_constants

func GRANT_ALL() string {
	return "ALL"
}

func GRANT_INSERT() string {
	return "INSERT"
}

func GRANT_UPDATE() string {
	return "UPDATE"
}

func GRANT_SELECT() string {
	return "SELECT"
}

func GET_ALLOWED_GRANTS() map[string]interface{} {
	valid_chars := make(map[string]interface{})
	valid_chars[GRANT_SELECT()] = nil
	valid_chars[GRANT_UPDATE()] = nil
	valid_chars[GRANT_INSERT()] = nil
	valid_chars[GRANT_ALL()] = nil
	return valid_chars
}