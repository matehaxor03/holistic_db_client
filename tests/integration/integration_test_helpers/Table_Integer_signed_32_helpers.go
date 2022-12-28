package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerSigned32ColumnName() string {
	return "holistic_test_table_with_integer_signed_32"
}

func GetTestTableNameWithIntegerSigned32ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_32_not_mandatory"
}

func GetTestTableIntegerSigned32ColumnName() string {
	return "integer_signed_32_column"
}

func GetTestTableInteger32SignedColumnNameNotMandatory() string {
	return "integer_signed_32_column_not_mandatory"
}


func GetTestSchemaWithIntegerSigned32Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned32ColumnName(): json.Map {"type": "int32"}}
}

func GetTestSchemaWithIntegerSigned32ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger32SignedColumnNameNotMandatory(): json.Map {"type": "*int32"}}
}