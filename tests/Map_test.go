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