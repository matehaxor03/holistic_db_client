package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned16ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_16"
}

func GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_16_not_mandatory"
}

func GetTestTableIntegerUnsigned16ColumnName() string {
	return "integer_unsigned_16_column"
}

func GetTestTableIntegerUnsigned16ColumnNameNotMandatory() string {
	return "integer_unsigned_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned16Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned16ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*uint16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned16ColumnNameNotMandatory(), column_schema)
	return table_schema
}