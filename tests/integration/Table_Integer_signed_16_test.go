package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithIntegerSigned16Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned16ColumnName(), helper.GetTestSchemaWithIntegerSigned16Column())
}

func TestTableCreateWithIntegerSigned16NotMandatoryColumn(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithIntegerSigned16ColumnNameNotMandatory(), helper.GetTestSchemaWithIntegerSigned16ColumnNotMandatory())
}
 
