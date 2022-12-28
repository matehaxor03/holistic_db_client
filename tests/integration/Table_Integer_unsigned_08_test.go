package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerUnsigned08Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnName(), helper.GetTestSchemaWithIntegerUnsigned08Column())
}

func TestTableCreateWithIntegerUnsigned08NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerUnsigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerUnsigned08NotMandatoryColumn())
}
 
