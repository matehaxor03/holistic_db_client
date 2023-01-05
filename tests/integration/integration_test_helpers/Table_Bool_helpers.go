package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableBoolColumnName() string {
	return "bool_column"
}

func GetTestTableBoolColumnNameNotMandatory() string {
	return "bool_column_not_mandatory"
}

func GetTestSchemaWithBoolColumn() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "bool")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableBoolColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithBoolColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*bool")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableBoolColumnNameNotMandatory(), column_schema)
	return table_schema
}