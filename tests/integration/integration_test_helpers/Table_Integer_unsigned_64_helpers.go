package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned64ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_64"
}

func GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_64_not_mandatory"
}

func GetTestTableIntegerUnsigned64ColumnName() string {
	return "integer_unsigned_64_column"
}

func GetTestTableIntegerUnsigned64ColumnNameNotMandatory() string {
	return "integer_unsigned_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned64Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned64ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*uint64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerUnsigned64ColumnNameNotMandatory(), column_schema)
	return table_schema
}