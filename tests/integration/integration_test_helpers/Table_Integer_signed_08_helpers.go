package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned08ColumnName() string {
	return "holistic_test_table_with_integer_signed_08"
}

func GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_08_not_mandatory"
}

func GetTestTableIntegerSigned08ColumnName() string {
	return "integer_signed_08_column"
}

func GetTestTableIntegerSigned08ColumnNameNotMandatory() string {
	return "integer_signed_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned08Column() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "int8")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned08ColumnName(), column_schema)
	return table_schema
}

func GetTestSchemaWithIntegerSigned08NotMandatoryColumn() json.Map {
	table_schema := json.Map{}
	column_schema := json.Map{}
	column_schema.SetStringValue("type", "*int8")
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTableIntegerSigned08ColumnNameNotMandatory(), column_schema)
	return table_schema
}