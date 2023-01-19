package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerUnsigned64Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithIntegerUnsigned64Column())
}

func TestTableCreateWithIntegerUnsigned64NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithIntegerUnsigned64ColumnNotMandatory())
}
 
