package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseNegativeInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-2147483649}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int64" {
		t.Errorf("key is not a *int64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors")
		} else if value == nil {
			t.Errorf("GetInt64 is nil")
		} else if *value != -2147483649 {
			t.Errorf("expected: -2147483649  actual: %d", *value)
		}
	}	
}

func TestCanParseNegativeInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-9223372036854775808}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int64" {
		t.Errorf("key is not a *int64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors")
		} else if value == nil {
			t.Errorf("GetInt64 is nil")
		} else if *value != -9223372036854775808 {
			t.Errorf("expected: -9223372036854775808  actual: %d", *value)
		}
	}	
}

func TestCannotParseNegativeInt64Overflow(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-9223372036854775809}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}