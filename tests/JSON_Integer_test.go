package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseUInt16_256(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":256}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
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
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt16_65535(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":65535}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
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
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt32_65536(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":65536}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint32" {
			t.Errorf("key is not a *uint32: %s", json.GetType("key"))
		} else {
			PrintJSON(t, json)
			value, value_errors := json.GetUInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 65536 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt32_4294967295(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":4294967295}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint32" {
			t.Errorf("key is not a *uint32: %s", json.GetType("key"))
		} else {
			PrintJSON(t, json)
			value, value_errors := json.GetUInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 4294967295 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt64_4294967296(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":4294967296}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint64" {
			t.Errorf("key is not a *uint64: %s", json.GetType("key"))
		} else {
			PrintJSON(t, json)
			value, value_errors := json.GetUInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 4294967296 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt64_18446744073709551615(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":18446744073709551615}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint64" {
			t.Errorf("key is not a *uint64: %s", json.GetType("key"))
		} else {
			PrintJSON(t, json)
			value, value_errors := json.GetUInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 18446744073709551615 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCannotParseUInt64_18446744073709551616(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":18446744073709551616}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}

func TestCanParseUInt64PositiveMuitple(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":4294967296,\"key2\":4294967297}")

	if json_errors != nil {
		t.Errorf("%s", json_errors) 
	} else if json == nil {
		t.Errorf("json is nil") 
	} else {
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
			PrintJSON(t, json)
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
	
}

	
	
	