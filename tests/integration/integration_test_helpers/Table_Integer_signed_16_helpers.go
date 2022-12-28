package integration_test_helpers

import (
    //"testing"
	//"strings"
	//"fmt"
	//"sync"
	json "github.com/matehaxor03/holistic_json/json"
	//common "github.com/matehaxor03/holistic_common/common"
	//class "github.com/matehaxor03/holistic_db_client/class"
)

func GetTestTableNameWithIntegerSigned16ColumnName() string {
	return "holistic_test_table_with_integer_signed_16"
}

func GetTestTableNameWithIntegerSigned16ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_16_not_mandatory"
}

func GetTestTableIntegerSigned16ColumnName() string {
	return "integer_signed_16_column"
}

func GetTestTableInteger16SignedColumnNameNotMandatory() string {
	return "integer_signed_16_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned16Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned16ColumnName(): json.Map {"type": "int16"}}
}

func GetTestSchemaWithIntegerSigned16ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableInteger16SignedColumnNameNotMandatory(): json.Map {"type": "*int16"}}
}