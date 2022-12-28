package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerSigned64Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned64ColumnName(), helper.GetTestSchemaWithIntegerSigned64Column())
}

func TestTableCreateWithIntegerSigned64NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned64ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned64ColumnNotMandatory())
}
 
