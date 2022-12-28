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

func GetTestTableNameWithFloat64Column() string {
	return "holistic_test_table_with_float64"
}

func GetTestTableNameWithFloat64ColumnNotMandatory() string {
	return "holistic_test_table_with_float64_not_mandatory"
}

func GetTestTableFloat64ColumnName() string {
	return "float64_column"
}

func GetTestTableFloat64ColumnNameNotMandatory() string {
	return "float64_column_not_mandatory"
}

func GetTestSchemaWithFloat64Column() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat64ColumnName(): json.Map {"type": "float64"}}
}

func GetTestSchemaWithFloat64ColumnNotMandatory() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "auto_increment": true, "primary_key": true},
	                  GetTestTableFloat64ColumnNameNotMandatory(): json.Map {"type": "*float64"}}
}