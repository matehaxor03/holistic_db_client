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

func GetTestTableNameWithIntegerSigned08ColumnName() string {
	return "holistic_test_table_with_integer_signed_08"
}

func GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory() string {
	return "holistic_test_table_with_integer_signed_08_not_mandatory"
}

func GetTestTableIntegerSigned08ColumnName() string {
	return "integer_signed_08_column"
}

func GetTestTableIntegerSigned08ColumnNameNotMandatory() string {
	return "integer_signed_08_column_not_mandatory"
}

func GetTestSchemaWithIntegerSigned08Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnName(): json.Map {"type": "int8"}}
}

func GetTestSchemaWithIntegerSigned08NotMandatoryColumn() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
					  GetTestTableIntegerSigned08ColumnNameNotMandatory(): json.Map {"type": "*int8"}}
}