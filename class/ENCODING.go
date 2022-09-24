package class

func GET_CHARACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHARACTER_SET_UTF8MB4() string {
	return "utf8mb4"
}

func GET_CHARACTER_SETS() Array {
	return Array{GET_CHARACTER_SET_UTF8(), GET_CHARACTER_SET_UTF8MB4()}
}

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() Array {
	return Array{GET_COLLATE_UTF8_GENERAL_CI(), GET_COLLATE_UTF8MB4_0900_AI_CI()}
}
