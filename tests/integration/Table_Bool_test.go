package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithBoolColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithBoolColumn())
}

func TestTableCreateWithBoolColumnNotMandatory(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithBoolColumnNotMandatory())
}
 
