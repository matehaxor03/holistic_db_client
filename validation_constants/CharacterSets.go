package validation_constants

func GET_CHARACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHARACTER_SET_UTF8MB4() string {
	return "utf8mb4"
}

func GET_CHARACTER_SETS() map[string]interface{} {
	valid_chars := make(map[string]interface{})
	valid_chars[GET_CHARACTER_SET_UTF8()] = nil
	valid_chars[GET_CHARACTER_SET_UTF8MB4()] = nil
	return valid_chars
}
