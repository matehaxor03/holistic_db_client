package validation_constants

import(
	json "github.com/matehaxor03/holistic_json/json"
)

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil(GET_COLLATE_UTF8_GENERAL_CI())
	valid_chars.SetNil(GET_COLLATE_UTF8MB4_0900_AI_CI())
	return valid_chars
}
