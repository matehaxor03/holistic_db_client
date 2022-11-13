package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseNil(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":null}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "nil" {
			t.Errorf("key is not a nil: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value != nil {
				t.Errorf("GetString is not nil")
			}
		}
	}
}

func TestCanParseMultipleNil(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":null,\"key2\":null}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "nil" {
			t.Errorf("key is not nil: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 
			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value != nil {
				t.Errorf("GetString is not nil")
			}
		}

		if !json.HasKey("key2") {
			t.Errorf("key2 not found")
		} else if json.GetType("key2") != "nil" {
			t.Errorf("key2 is not nil: %s", json.GetType("key2"))
		} else {
			value, value_errors := json.GetString("key2") 
			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value != nil {
				t.Errorf("GetString is not nil")
			}
		}
	}
}
