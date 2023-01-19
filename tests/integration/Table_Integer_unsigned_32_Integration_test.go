package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerUnsigned32Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithIntegerUnsigned32Column())
}

func TestTableCreateWithIntegerUnsigned32NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestSchemaWithIntegerUnsigned32ColumnNotMandatory())
}
 
