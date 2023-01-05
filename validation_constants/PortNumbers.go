package validation_constants

import(
	json "github.com/matehaxor03/holistic_json/json"
)

func GetValidPortNumberCharacters() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil("0")
	valid_chars.SetNil("1")
	valid_chars.SetNil("2")
	valid_chars.SetNil("3")
	valid_chars.SetNil("4")
	valid_chars.SetNil("5")
	valid_chars.SetNil("6")
	valid_chars.SetNil("7")
	valid_chars.SetNil("8")
	valid_chars.SetNil("9")
	return valid_chars
}