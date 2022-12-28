package integration
 
import (
    "testing"
	helper "github.com/matehaxor03/holistic_db_client/tests/integration/integration_test_helpers"
)

func TestTableCreateWithFloat32Column(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableFloat32ColumnName(), helper.GetTestSchemaWithFloat32Column())
}

func TestTableCreateWithFloat32ColumnNotMandatory(t *testing.T) {
	helper.CreateTableAndVerifySchema(t, helper.GetTestTableNameWithFloat32ColumnNotMandatory(), helper.GetTestSchemaWithFloat32ColumnNotMandatory())
}
 
