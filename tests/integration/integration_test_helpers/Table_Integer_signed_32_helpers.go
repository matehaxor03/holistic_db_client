package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableIntegerSigned32ColumnName() string {
	return "integer_signed_32_column"
}

func GetTestTableIntegerSigned32ColumnNameNotMandatory() string {
	return "integer_signed_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned32Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "int32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned32ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned32ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*int32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned32ColumnNameNotMandatory(), column_schema)
	return table_schema
}