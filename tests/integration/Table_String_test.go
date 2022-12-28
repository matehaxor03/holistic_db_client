package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithStringColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableStringColumnName(), helper.GetTestSchemaWithStringColumn())
}

func TestTableCreateWithStringColumnNotMandatory(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithStringColumnNotMandatory(), helper.GetTestSchemaWithStringColumnNotMandatory())
}
 
