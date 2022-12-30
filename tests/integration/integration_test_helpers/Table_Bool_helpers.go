package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithBoolColumn() string {
	return "holistic_test_table_with_bool"
}

func GetTestTableNameWithBoolColumnNotMandatory() string {
	return "holistic_test_table_with_bool_not_mandatory"
}

func GetTestTableBoolColumnName() string {
	return "bool_column"
}

func GetTestTableBoolColumnNameNotMandatory() string {
	return "bool_column_not_mandatory"
}

func GetTestSchemaWithBoolColumn() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "bool")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableBoolColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithBoolColumnNotMandatory() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*bool")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableBoolColumnNameNotMandatory(), column_schema)
	return table_schema
}