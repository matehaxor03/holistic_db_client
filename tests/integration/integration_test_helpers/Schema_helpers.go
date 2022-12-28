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

func GetTestColumnSchemaNoPrimaryKey() json.Map {
	return json.Map {"type": "uint64" }
}

func GetTestColumnSchemaNoType() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"primary_key": true}}
}

func GetTestColumnSchemaWithValue() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): json.Map {"type": "uint64", "value":"something", "auto_increment": true, "primary_key": true}}
}

func GetTestTableSchemaNoPrimaryKey() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoPrimaryKey()}
}

func GetTestTableSchemaMoreThanOnePrimaryKeyAutoIncrement() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestSchemaColumnPrimaryKeyAutoIncrement(),
	                  GetTestTablePrimaryKeyName2(): GetTestSchemaColumnPrimaryKeyAutoIncrement()}
}

func GetTestTableSchemaNoType() json.Map {
	return json.Map {GetTestTablePrimaryKeyName(): GetTestColumnSchemaNoType()}
}