package tests
 
import (
    "testing"
)

func TestCanParseNegativeInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-129}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int16" {
		t.Errorf("key is not a *int16: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetInt16 has errors")
		} else if value == nil {
			t.Errorf("GetInt16 is nil")
		} else if *value != -129 {
			t.Errorf("expected: -129  actual: %d", *value)
		}
	}	
}

func TestCanParseNegativeInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-32768}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*int16" {
		t.Errorf("key is not a *int16: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetInt16 has errors")
		} else if value == nil {
			t.Errorf("GetInt16 is nil")
		} else if *value != -32768 {
			t.Errorf("expected: -32768  actual: %d", *value)
		}
	}
}


	