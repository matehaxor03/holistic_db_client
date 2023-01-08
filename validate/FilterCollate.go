package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type CollateWordWhitelist struct {
	GetCollateWordWhitelist func() (*json.Map)
}

func NewCollateWordWhitelist() (*CollateWordWhitelist) {
	valid_words := validation_constants.GET_COLLATES()

	x := CollateWordWhitelist {
		GetCollateWordWhitelist: func() (*json.Map) {
			return &valid_words
		},
	}
	return &x
}
