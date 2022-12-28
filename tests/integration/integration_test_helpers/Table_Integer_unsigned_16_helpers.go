package integration_test_helpers

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func GetTestTableNameWithIntegerUnsigned16ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_16"
}

func GetTestTableNameWithIntegerUnsigned16ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_16_not_mandatory"
}

func GetTestTableIntegerUnsigned16ColumnName() string {
	return "integer_unsigned_16_column"
}

func GetTestTableIntegerUnsigned16ColumnNameNotMandatory() string {
	return "integer_unsigned_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned16Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned16ColumnName(): json.Map {"type": "uint16"}}
}

func GetTestSchemaWithIntegerUnsigned16ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned16ColumnNameNotMandatory(): json.Map {"type": "*uint16"}}
}