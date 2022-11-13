package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseBoolTrue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	}

	type_of := json.GetType("key")
	if type_of != "*bool" {
		t.Errorf("key is not a bool: %s", type_of)
	}

	value, value_errors := json.GetBool("key") 

	if value_errors != nil {
		t.Errorf("map GetBool has errors")
	} else if value == nil {
		t.Errorf("GetBool is nil")
	} else if *value != true {
		t.Errorf("expected: value actual: %t", *value)
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCanParseBoolFalse(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":false}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	}

	type_of := json.GetType("key")
	if type_of != "*bool" {
		t.Errorf("key is not a bool: %s", type_of)
	}

	value, value_errors := json.GetBool("key") 

	if value_errors != nil {
		t.Errorf("map GetBool has errors")
	} else if value == nil {
		t.Errorf("GetBool is nil")
	} else if *value != false {
		t.Errorf("expected: value actual: %t", *value)
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCanParseMultipleBools(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":true,\"key2\":false}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	} else {
		type_of := json.GetType("key")
		if type_of != "*bool" {
			t.Errorf("key is not a bool: %s", type_of)
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

	has_key2 := json.HasKey("key2")
	if !has_key2 {
		t.Errorf("key2 not found")
	} else {
		type_of := json.GetType("key2")
		if type_of != "*bool" {
			t.Errorf("key2 is not a bool: %s", type_of)
		} else {
			value, value_errors := json.GetBool("key2") 

			if value_errors != nil {
				t.Errorf("map GetBool has errors")
			} else if value == nil {
				t.Errorf("GetBool is nil")
			} else if *value != false {
				t.Errorf("expected: value2 actual: %t", *value)
			}
		}
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}


