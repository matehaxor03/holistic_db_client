package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned08ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_08"
}

func GetTestTableNameWithIntegerUnsigned08ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_08_not_mandatory"
}


func GetTestTableIntegerUnsigned08ColumnName() string {
	return "integer_unsigned_08_column"
}

func GetTestTableIntegerUnsigned08ColumnNameNotMandatory() string {
	return "integer_unsigned_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned08Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint8")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned08ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*uint8")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned08ColumnNameNotMandatory(), column_schema)
	return table_schema
}