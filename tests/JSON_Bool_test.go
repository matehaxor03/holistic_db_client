package tests
 
import (
    "testing"
)

func TestCanParseBoolTrue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":true}")

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

func TestCanParseBoolFalse(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":false}")

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

func TestCanParseMultipleBools(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":true,\"key2\":false}")

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

func TestCanParseBoolSpaceBeforeKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{ \"key\":true}")

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

func TestCanParseBoolNewlineBeforeKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\n\"key\":true}")

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

func TestCanParseBoolDosNewlineBeforeKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\r\n\"key\":true}")

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

func TestCanParseBoolSpaceAfterKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\" :true}")

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

func TestCanParseBoolNewlineAfterKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\"\n:true}")

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

func TestCanParseBoolDosNewlineAfterKey(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\"\r\n:true}")

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

func TestCanParseBoolSpaceBeforeValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\": true}")

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

func TestCanParseBoolSpaceNewlineBeforeValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":\ntrue}")

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

func TestCanParseBoolSpaceDosNewlineBeforeValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":\r\ntrue}")

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

func TestCanParseBoolSpaceAfterValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":true }")

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

func TestCanParseBoolNewlineAfterValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":true\n}")

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

func TestCanParseBoolDosNewlineAfterValue(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":true\r\n}")

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
