package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func PrintJSON(t *testing.T, json *class.Map) {
	if json == nil {
		t.Errorf("json is nil")
	} else {
		json_string, json_string_errors := json.ToJSONString()
		if json_string_errors != nil {
			t.Errorf("%s", json_string_errors)
		} else if json_string == nil {
			t.Errorf("json_string is nil")
		} else {
			fmt.Println(*json_string)
		}
	}
}

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
