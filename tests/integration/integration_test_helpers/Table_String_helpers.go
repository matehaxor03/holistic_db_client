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
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnName(): json.Map {"type": "string", "max_length":100}}
}

func GetTestSchemaWithStringColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableStringColumnNameNotMandatory(): json.Map {"type": "*string", "max_length":100}}
}