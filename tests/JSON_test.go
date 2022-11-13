package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseMinimal(t *testing.T) {
	value, value_errors := class.ParseJSON("{}")

	if value_errors != nil {
		t.Errorf("%s", value_errors)
	} 

	if value == nil {
		t.Errorf("map is nil")
	} else if len(value.Keys()) != 0  {
		t.Errorf("keys is not zero")
	}
}

func TestCannotParseEmptyString(t *testing.T) {
	value, value_errors := class.ParseJSON("")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}

func TestCannotParseMalformedBrackets1(t *testing.T) {
	value, value_errors := class.ParseJSON("{")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}

func TestCannotParseMalformedBrackets2(t *testing.T) {
	value, value_errors := class.ParseJSON("}")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}

func TestCanParseString(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	}

	type_of := json.GetType("key")
	if type_of != "*string" {
		t.Errorf("key is not a string: %s", type_of)
	}

	value, value_errors := json.GetString("key") 

	if value_errors != nil {
		t.Errorf("map GetString has errors")
	} else if value == nil {
		t.Errorf("GetString is nil")
	} else if *value != "value" {
		t.Errorf("expected: value actual: %s", *value)
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCannotParseStringWithoutQuotePrefix(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":value\"}")

	if json_errors == nil {
		t.Errorf("expected errors for ParseJSON")
	}

	if json != nil {
		t.Errorf("expected nil json")
	}
}

func TestCannotParseStringWithoutQuoteSuffix(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value}")

	if json_errors == nil {
		t.Errorf("expected errors for ParseJSON")
	}

	if json != nil {
		t.Errorf("expected nil json")
	}
}

func TestCanParseMultipleStrings(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\",\"key2\":\"value2\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	} else {
		type_of := json.GetType("key")
		if type_of != "*string" {
			t.Errorf("key is not a string: %s", type_of)
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}

	has_key2 := json.HasKey("key2")
	if !has_key2 {
		t.Errorf("key2 not found")
	} else {
		type_of := json.GetType("key2")
		if type_of != "*string" {
			t.Errorf("key2 is not a string: %s", type_of)
		} else {
			value, value_errors := json.GetString("key2") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value2" {
				t.Errorf("expected: value2 actual: %s", *value)
			}
		}
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

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

func TestCanParseNil(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":null}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	has_key := json.HasKey("key")
	if !has_key {
		t.Errorf("key not found")
	}

	type_of := json.GetType("key")
	if type_of != "nil" {
		t.Errorf("key is not nil: %s", type_of)
	}

	value, value_errors := json.GetString("key") 

	if value_errors != nil {
		t.Errorf("map GetString has errors")
	} else if value != nil {
		t.Errorf("GetString is not nil")
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCanParseMultipleNil(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":null,\"key2\":null}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
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

		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}
