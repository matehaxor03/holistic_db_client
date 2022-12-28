package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithFloat64Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableFloat64ColumnName(), helper.GetTestSchemaWithFloat64Column())
}

func TestTableCreateWithFloat64ColumnNotMandatory(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithFloat64ColumnNotMandatory(), helper.GetTestSchemaWithFloat64ColumnNotMandatory())
}
 
