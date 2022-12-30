package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned32ColumnName() string {
	return "holistic_test_table_with_integer_signed_32"
}

func GetTestTableNameWithIntegerSigned32ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_32_not_mandatory"
}

func GetTestTableIntegerSigned32ColumnName() string {
	return "integer_signed_32_column"
}

func GetTestTableIntegerSigned32ColumnNameNotMandatory() string {
	return "integer_signed_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned32Column() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "int32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned32ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned32ColumnNotMandatory() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*int32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned32ColumnNameNotMandatory(), column_schema)
	return table_schema
}