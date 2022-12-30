package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithFloat32Column() string {
	return "holistic_test_table_with_float32"
}

func GetTestTableNameWithFloat32ColumnNotMandatory() string {
	return "holistic_test_table_with_float32_not_mandatory"
}

func GetTestTableFloat32ColumnName() string {
	return "float32_column"
}

func GetTestTableFloat32ColumnNameNotMandatory() string {
	return "float32_column_not_mandatory"
}

func GetTestSchemaWithFloat32Column() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "float32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableFloat32ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithFloat32ColumnNotMandatory() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*float32")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableFloat32ColumnNameNotMandatory(), column_schema)
	return table_schema
}