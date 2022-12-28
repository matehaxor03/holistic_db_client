package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerSigned08Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned08ColumnName(), helper.GetTestSchemaWithIntegerSigned08Column())
}

func TestTableCreateWithIntegerSigned08NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned08ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned08NotMandatoryColumn())
}
 
