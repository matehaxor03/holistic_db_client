package tests
 
import (
    "testing"
)

func TestCanParseNegativeInt32LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-32769}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int32" {
		t.Errorf("key is not a *int32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetInt32 has errors")
		} else if value == nil {
			t.Errorf("GetInt32 is nil")
		} else if *value != -32769 {
			t.Errorf("expected: -32769  actual: %d", *value)
		}
	}
}

func TestCanParseNegativeInt32HighBoundarfy(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-2147483648}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int32" {
		t.Errorf("key is not a *int32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetInt32 has errors")
		} else if value == nil {
			t.Errorf("GetInt32 is nil")
		} else if *value != -2147483648 {
			t.Errorf("expected: -2147483648  actual: %d", *value)
		}
	}	
}
