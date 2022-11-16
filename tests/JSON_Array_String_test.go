package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseEmptyArray(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":[]}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		//PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*class.Array" {
			t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
		} else {			
			value, value_errors := json.GetArray("key") 

			if value_errors != nil {
				t.Errorf("map GetArray has errors")
			} else if value == nil {
				t.Errorf("GetArray is nil")
			} else if len(*value) != 0 {
				t.Errorf("expected: length=0 actual: length=%d", len(*value))
			}
		}
	}	
}

func TestCanParseArrayContainingSingleString(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":[\"value\"]}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		//PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*class.Array" {
			t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
		} else {			
			value, value_errors := json.GetArray("key") 

			if value_errors != nil {
				t.Errorf("map GetArray has errors")
			} else if value == nil {
				t.Errorf("GetArray is nil")
			} else if len(*value) != 1 {
				t.Errorf("expected: length=1 actual: length=%d", len(*value))
			} else if (*value)[0].(string) != "value" {
				t.Errorf("expected \"value\" actual: \"%s\"", (*value)[0].(string))
			}
		}
	}	
}

func TestCanParseMultipleArraysContainingSingleString(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":[\"value\"],\"key2\":[\"value2\"]}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		//PrintJSON(t, json)

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


		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*class.Array" {
			t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
		} else {			
			value, value_errors := json.GetArray("key") 

			if value_errors != nil {
				t.Errorf("map GetArray has errors")
			} else if value == nil {
				t.Errorf("GetArray is nil")
			} else if len(*value) != 1 {
				t.Errorf("expected: length=1 actual: length=%d", len(*value))
			} else if (*value)[0].(string) != "value" {
				t.Errorf("expected \"value\" actual: \"%s\"", (*value)[0].(string))
			}
		}

		if !json.HasKey("key2") {
			t.Errorf("key2 not found")
		} else if json.GetType("key2") != "*class.Array" {
			t.Errorf("key2 is not a *class.Array: %s", json.GetType("key2"))
		} else {			
			value, value_errors := json.GetArray("key2") 

			if value_errors != nil {
				t.Errorf("map GetArray has errors")
			} else if value == nil {
				t.Errorf("GetArray is nil")
			} else if len(*value) != 1 {
				t.Errorf("expected: length=1 actual: length=%d", len(*value))
			} else if (*value)[0].(string) != "value2" {
				t.Errorf("expected \"value2\" actual: \"%s\"", (*value)[0].(string))
			}
		}
	}	
}

func TestCanParseArrayContainingMultipleStrings(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":[\"value\",\"value2\"]}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		//PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*class.Array" {
			t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
		} else {			
			value, value_errors := json.GetArray("key") 

			if value_errors != nil {
				t.Errorf("map GetArray has errors")
			} else if value == nil {
				t.Errorf("GetArray is nil")
			} else if len(*value) != 2 {
				t.Errorf("expected: length=2 actual: length=%d", len(*value))
			} else if (*value)[0].(string) != "value" {
				t.Errorf("expected \"value\" actual: \"%s\"", (*value)[0].(string))
			} else if (*value)[1].(string) != "value2" {
				t.Errorf("expected \"value2\" actual: \"%s\"", (*value)[0].(string))
			}
		}
	}	
}