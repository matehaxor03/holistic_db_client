package tests
 
import (
    "testing"
)

func TestCanParseUInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t,"{\"key\":256}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*uint16" {
		t.Errorf("key is not a *uint16: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetUInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt16 is nil")
		} else if *value != 256 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}	
}

func TestCanParseUInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t,"{\"key\":65535}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*uint16" {
		t.Errorf("key is not a *uint16: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetUInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt16 is nil")
		} else if *value != 65535 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}
