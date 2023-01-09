package validation_constants

import(
	json "github.com/matehaxor03/holistic_json/json"
)

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

func GET_ALLOWED_GRANTS() json.Map {
	valid := json.NewMapValue()
	valid.SetNil(GRANT_ALL())
	valid.SetNil(GRANT_INSERT())
	valid.SetNil(GRANT_UPDATE())
	valid.SetNil(GRANT_SELECT())
	return valid
}