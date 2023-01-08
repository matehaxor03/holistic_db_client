package validation_constants

import(
	json "github.com/matehaxor03/holistic_json/json"
)

func GET_CHARACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHARACTER_SET_UTF8MB4() string {
	return "utf8mb4"
}

func GET_CHARACTER_SETS() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil(GET_CHARACTER_SET_UTF8())
	valid_chars.SetNil(GET_CHARACTER_SET_UTF8MB4())
	return valid_chars
}
