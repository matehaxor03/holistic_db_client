package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)


func TestCanParseUInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t,"{\"key\":4294967296}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*uint64" {
		t.Errorf("key is not a *uint64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetUInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt32 is nil")
		} else if *value != 4294967296 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t,"{\"key\":18446744073709551615}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*uint64" {
		t.Errorf("key is not a *uint64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetUInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt32 is nil")
		} else if *value != 18446744073709551615 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCannotParseUInt64Overflow(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":18446744073709551616}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}

func TestCanParseUInt64PositiveMuitpleHighBondary(t *testing.T) {
	json := ParseJSONSuccessfully(t,"{\"key\":4294967296,\"key2\":4294967297}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*uint64" {
		t.Errorf("key is not a *uint64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetUInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors")
		} else if value == nil {
			t.Errorf("GetUInt64 is nil")
		} else if *value != 4294967296 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key") != "*uint64" {
		t.Errorf("key2 is not a *uint64: %s", json.GetType("key2"))
	} else {
		value, value_errors := json.GetUInt64("key2") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors")
		} else if value == nil {
			t.Errorf("GetUInt64 is nil")
		} else if *value != 4294967297 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}
	