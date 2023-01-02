package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned64ColumnName() string {
	return "holistic_test_table_with_integer_signed_64"
}

func GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_64_not_mandatory"
}


func GetTestTableIntegerSigned64ColumnName() string {
	return "integer_signed_64_column"
}

func GetTestTableIntegerSigned64ColumnNameNotMandatory() string {
	return "integer_signed_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned64Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "int64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned64ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned64ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*int64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned64ColumnNameNotMandatory(), column_schema)
	return table_schema
}