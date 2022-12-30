package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned16ColumnName() string {
	return "holistic_test_table_with_integer_signed_16"
}

func GetTestTableNameWithIntegerSigned16ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_16_not_mandatory"
}

func GetTestTableIntegerSigned16ColumnName() string {
	return "integer_signed_16_column"
}

func GetTestTableIntegerSigned16ColumnNameNotMandatory() string {
	return "integer_signed_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned16Column() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "int16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned16ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned16ColumnNotMandatory() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*int16")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned16ColumnNameNotMandatory(), column_schema)
	return table_schema
}