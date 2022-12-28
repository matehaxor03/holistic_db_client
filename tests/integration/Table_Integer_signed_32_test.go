package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerSigned32Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned32ColumnName(), helper.GetTestSchemaWithIntegerSigned32Column())
}

func TestTableCreateWithIntegerSigned32NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned32ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned32ColumnNotMandatory())
}
 
