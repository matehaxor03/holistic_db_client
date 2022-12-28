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