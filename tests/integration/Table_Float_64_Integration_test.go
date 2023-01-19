package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithFloat64Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithFloat64Column())
}

func TestTableCreateWithFloat64ColumnNotMandatory(t *testing.T) {
	helper.CreateTableAndVerifySchema(t,  helper.GetTestSchemaWithFloat64ColumnNotMandatory())
}
 
