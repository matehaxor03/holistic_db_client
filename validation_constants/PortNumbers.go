package validation_constants

func GetValidPortNumberCharacters() map[string]interface{} {
	valid_chars := make(map[string]interface{})
	valid_chars["0"] = nil
	valid_chars["1"] = nil
	valid_chars["2"] = nil
	valid_chars["3"] = nil
	valid_chars["4"] = nil
	valid_chars["5"] = nil
	valid_chars["6"] = nil
	valid_chars["7"] = nil
	valid_chars["8"] = nil
	valid_chars["9"] = nil
	return valid_chars
}