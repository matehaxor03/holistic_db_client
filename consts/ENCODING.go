package consts

func GET_CHRACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHRACTER_SETS() []string {
	return []string{GET_CHRACTER_SET_UTF8()}
}

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATES() []string {
	return []string{GET_COLLATE_UTF8_GENERAL_CI()}
}
