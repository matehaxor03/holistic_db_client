package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseNegativeInt8_1(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-1}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int8" {
			t.Errorf("key is not a *int8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt8 has errors")
			} else if value == nil {
				t.Errorf("GetInt8 is nil")
			} else if *value != -1 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt8_128(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-128}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int8" {
			t.Errorf("key is not a *int8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt8 has errors")
			} else if value == nil {
				t.Errorf("GetInt8 is nil")
			} else if *value != -128 {
				t.Errorf("expected: -128 actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt16_129(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-129}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
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

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}