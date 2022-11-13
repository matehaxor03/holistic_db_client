package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseBoolTrue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolFalse(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":false}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != false {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseMultipleBools(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true,\"key2\":false}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}

		if !json.HasKey("key2") {
			t.Errorf("key2 not found")
		} else if json.GetType("key2") != "*bool" {
			t.Errorf("key2 is not a *bool: %s", json.GetType("key2"))
		} else {
			value, value_errors := json.GetBool("key2") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != false {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{ \"key\":true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolNewlineBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\n\"key\":true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolDosNewlineBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\r\n\"key\":true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\" :true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolNewlineAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\"\n:true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolDosNewlineAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\"\r\n:true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\": true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceNewlineBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\ntrue}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceDosNewlineBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\r\ntrue}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolSpaceAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true }")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolNewlineAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true\n}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}

func TestCanParseBoolDosNewlineAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true\r\n}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*bool" {
			t.Errorf("key is not a *bool: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetBool("key") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != true {
				t.Errorf("expected: value actual: %t", *value)
			}
		}
	}
}


