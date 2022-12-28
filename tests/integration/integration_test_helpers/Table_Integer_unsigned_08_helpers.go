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

func GetTestTableNameWithIntegerUnsigned08ColumnName() string {
	return "holistic_test_table_with_integer_unsigned_08"
}

func GetTestTableNameWithIntegerUnsigned08ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_unsigned_08_not_mandatory"
}


func GetTestTableIntegerUnsigned08ColumnName() string {
	return "integer_unsigned_08_column"
}

func GetTestTableIntegerUnsigned08ColumnNameNotMandatory() string {
	return "integer_unsigned_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerUnsigned08Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned08ColumnName(): json.Map {"type": "uint8"}}
}

func GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerUnsigned08ColumnNameNotMandatory(): json.Map {"type": "*uint8"}}
}