package integration_test_helpers

import( 
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestColumnSchemaNoPrimaryKey() json.Map {
	return json.Map {"type": "uint64" }
}

func GetTestColumnSchemaNoType() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"primary_key": true}}
}

func GetTestColumnSchemaWithValue() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "value":"something", "auto_increment": true, "primary_key": true}}
}

func GetTestTableSchemaNoPrimaryKey() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoPrimaryKey()}
}

func GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestSchemaColumnPrimaryKeyAutoIncrement(),
	                  GetTestTablePrimaryKeyName2(): GetTestSchemaColumnPrimaryKeyAutoIncrement()}
}

func GetTestTableSchemaNoType() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoType()}
}