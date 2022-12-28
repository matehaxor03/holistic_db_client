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

func GetTestTableNameWithBoolColumn() string {
	return "holistic_test_table_with_bool"
}

func GetTestTableNameWithBoolColumnNotMandatory() string {
	return "holistic_test_table_with_boool_not_mandatory"
}

func GetTestTableBoolColumnName() string {
	return "bool_column"
}

func GetTestTableBoolColumnNameNotMandatory() string {
	return "bool_column_not_mandatory"
}

func GetTestSchemaWithBoolColumn() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnName(): json.Map {"type": "bool"}}
}

func GetTestSchemaWithBoolColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableBoolColumnNameNotMandatory(): json.Map {"type": "*bool"}}
}