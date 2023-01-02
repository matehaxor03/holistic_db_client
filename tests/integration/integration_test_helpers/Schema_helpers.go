package integration_test_helpers

import( 
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestColumnSchemaNoPrimaryKey() json.Map {
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint64")
	column := json.NewMapValue()
	column.SetMapValue(GetTestTablePrimaryKeyName(), column_schema)
	return column
}

func GetTestColumnSchemaNoType() json.Map {
	column_schema := json.NewMapValue()
	column_schema.SetBoolValue("primary_key", true)
	column := json.NewMapValue()
	column.SetMapValue(GetTestTablePrimaryKeyName(), column_schema)
	return column
}

func GetTestColumnSchemaWithValue() json.Map {
	column_schema := json.NewMapValue()
	column_schema.SetStringValue("type", "uint64")
	column_schema.SetStringValue("value", "something")
	column_schema.SetBoolValue("auto_increment", true)
	column_schema.SetBoolValue("primary_key", true)
	column := json.NewMapValue()
	column.SetMapValue(GetTestTablePrimaryKeyName(), column_schema)
	return column
}

func GetTestTableSchemaNoPrimaryKey() json.Map {
	table_schema := json.NewMapValue()
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestColumnSchemaNoPrimaryKey())
	return table_schema
}

func GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement() json.Map {
	table_schema := json.NewMapValue()
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	table_schema.SetMapValue(GetTestTablePrimaryKeyName2(), GetTestSchemaColumnPrimaryKeyAutoIncrement())
	return table_schema
}

func GetTestTableSchemaNoType() json.Map {
	table_schema := json.NewMapValue()
	table_schema.SetMapValue(GetTestTablePrimaryKeyName(), GetTestColumnSchemaNoType())
	return table_schema
}