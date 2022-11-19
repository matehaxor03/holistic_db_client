package tests
 
import (
    "testing"
)

func TestCanParseNegativeInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-1}")
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int8" {
		t.Errorf("key is not a *int8: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetInt8 has errors")
		} else if value == nil {
			t.Errorf("GetInt8 is nil")
		} else if *value != -1 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseNegativeInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-128}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int8" {
		t.Errorf("key is not a *int8: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetInt8 has errors")
		} else if value == nil {
			t.Errorf("GetInt8 is nil")
		} else if *value != -128 {
			t.Errorf("expected: -128 actual: %d", *value)
		}
	}
}
	