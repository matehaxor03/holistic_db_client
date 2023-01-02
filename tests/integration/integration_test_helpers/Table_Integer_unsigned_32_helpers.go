package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned32ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_32"
}

func GetTestTableNameWithIntegerUnsigned32ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_32_not_mandatory"
}

func GetTestTableIntegerUnsigned32ColumnName() string {
	return "integer_unsigned_32_column"
}

func GetTestTableIntegerUnsigned32ColumnNameNotMandatory() string {
	return "integer_unsigned_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned32Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned32ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*uint32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned32ColumnNameNotMandatory(), column_schema)
	return table_schema
}