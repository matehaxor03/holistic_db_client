package validation_constants

func GetValidDomainNameCharacters() map[string]interface{} {
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
	valid_chars["a"] = nil
	valid_chars["b"] = nil
	valid_chars["c"] = nil
	valid_chars["d"] = nil
	valid_chars["e"] = nil
	valid_chars["f"] = nil
	valid_chars["g"] = nil
	valid_chars["h"] = nil
	valid_chars["i"] = nil
	valid_chars["j"] = nil
	valid_chars["k"] = nil
	valid_chars["l"] = nil
	valid_chars["m"] = nil
	valid_chars["n"] = nil
	valid_chars["o"] = nil
	valid_chars["p"] = nil
	valid_chars["q"] = nil
	valid_chars["r"] = nil
	valid_chars["s"] = nil
	valid_chars["t"] = nil
	valid_chars["u"] = nil
	valid_chars["v"] = nil
	valid_chars["w"] = nil
	valid_chars["x"] = nil
	valid_chars["y"] = nil
	valid_chars["z"] = nil
	valid_chars["A"] = nil
	valid_chars["B"] = nil
	valid_chars["C"] = nil
	valid_chars["D"] = nil
	valid_chars["E"] = nil
	valid_chars["F"] = nil
	valid_chars["G"] = nil
	valid_chars["H"] = nil
	valid_chars["I"] = nil
	valid_chars["J"] = nil
	valid_chars["K"] = nil
	valid_chars["L"] = nil
	valid_chars["M"] = nil
	valid_chars["N"] = nil
	valid_chars["O"] = nil
	valid_chars["P"] = nil
	valid_chars["Q"] = nil
	valid_chars["R"] = nil
	valid_chars["S"] = nil
	valid_chars["T"] = nil
	valid_chars["U"] = nil
	valid_chars["V"] = nil
	valid_chars["W"] = nil
	valid_chars["X"] = nil
	valid_chars["Y"] = nil
	valid_chars["Z"] = nil
	valid_chars["-"] = nil
	valid_chars["_"] = nil
	valid_chars["."] = nil
	return valid_chars
}

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() map[string]interface{} {
	valid := make(map[string]interface{})
	valid[LOCALHOST_IP()] = nil
	return valid
}