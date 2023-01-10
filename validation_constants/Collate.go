package validation_constants

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() map[string]interface{} {
	valid_chars := make(map[string]interface{})
	valid_chars[GET_COLLATE_UTF8_GENERAL_CI()] = nil
	valid_chars[GET_COLLATE_UTF8MB4_0900_AI_CI()] = nil
	return valid_chars
}
