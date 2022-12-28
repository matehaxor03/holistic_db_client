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
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat32ColumnName(): json.Map {"type": "float32"}}
}

func GetTestSchemaWithFloat32ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat32ColumnNameNotMandatory(): json.Map {"type": "*float32"}}
}