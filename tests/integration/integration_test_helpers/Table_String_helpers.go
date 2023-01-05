package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableStringColumnName() string {
	return "string_column"
}

func GetTestTableStringColumnNameNotMandatory() string {
	return "string_column_not_mandatory"
}

func GetTestSchemaWithStringColumn() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "string")
	column_schema.SetIntValue("max_length", 100)
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableStringColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithStringColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*string")
	column_schema.SetIntValue("max_length", 100)
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableStringColumnNameNotMandatory(), column_schema)
	return table_schema
}