package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableIntegerSigned16ColumnName() string {
	return "integer_signed_16_column"
}

func GetTestTableIntegerSigned16ColumnNameNotMandatory() string {
	return "integer_signed_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned16Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "int16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned16ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned16ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*int16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned16ColumnNameNotMandatory(), column_schema)
	return table_schema
}