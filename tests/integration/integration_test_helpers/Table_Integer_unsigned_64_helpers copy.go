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

func GetTestTableNameWithIntegerUnsigned64ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_64"
}

func GetTestTableNameWithIntegerUnsigned64ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_64_not_mandatory"
}

func GetTestTableIntegerUnsigned64ColumnName() string {
	return "integer_unsigned_64_column"
}

func GetTestTableIntegerUnsigned64ColumnNameNotMandatory() string {
	return "integer_unsigned_64_column_not_mandatory"
}


func GetTestSchemaWithIntegerUnsigned64Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned64ColumnName(): json.Map {"type": "uint64"}}
}

func GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned64ColumnNameNotMandatory(): json.Map {"type": "*uint64"}}
}