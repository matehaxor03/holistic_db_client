package consts

func GET_LOGIC_STATEMENT_FIELD_NAME() (string) {
	return "LOGIC"
}

func GET_LOGIC_STATEMENT_IF() (string) {
	return "IF"
}

func GET_LOGIC_STATEMENT_NOT() (string) {
	return "NOT"
}

func GET_LOGIC_STATEMENT_EXISTS() (string) {
	return "EXISTS"
}

func GET_LOGIC_STATEMENT_IF_NOT_EXISTS() ([]string) {
	return []string{GET_LOGIC_STATEMENT_IF(), GET_LOGIC_STATEMENT_NOT(), GET_LOGIC_STATEMENT_EXISTS()}
}

func GET_LOGIC_STATEMENT_IF_EXISTS() ([]string) {
	return []string{GET_LOGIC_STATEMENT_IF(), GET_LOGIC_STATEMENT_EXISTS()}
}
