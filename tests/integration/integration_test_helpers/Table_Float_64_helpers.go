package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableFloat64ColumnName() string {
	return "float64_column"
}

func GetTestTableFloat64ColumnNameNotMandatory() string {
	return "float64_column_not_mandatory"
}

func GetTestSchemaWithFloat64Column() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "float64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableFloat64ColumnName(), column_schema)
	return table_schema
}


func GetTestSchemaWithFloat64ColumnNotMandatory() json.Map {
	table_schema := json.NewMapValue()
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "*float64")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableFloat64ColumnNameNotMandatory(), column_schema)
	return table_schema
}