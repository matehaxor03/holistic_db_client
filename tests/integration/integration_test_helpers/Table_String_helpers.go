package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithStringColumn() string {
	return "holistic_test_table_with_string"
}

func GetTestTableNameWithStringColumnNotMandatory() string {
	return "holistic_test_table_with_string_not_mandatory"
}

func GetTestTableStringColumnName() string {
	return "string_column"
}

func GetTestTableStringColumnNameNotMandatory() string {
	return "string_column_not_mandatory"
}

func GetTestSchemaWithStringColumn() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "string")
	column_schema.SetIntValue("max_length", 100)
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableStringColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithStringColumnNotMandatory() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*string")
	column_schema.SetIntValue("max_length", 100)
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableStringColumnNameNotMandatory(), column_schema)
	return table_schema
}