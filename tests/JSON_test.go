package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func ParseJSONSuccessfully(t *testing.T, json_string string) (*class.Map) {
	json, json_errors := class.ParseJSON(json_string)

	if json_errors != nil {
		t.Errorf("%s", json_errors)
		t.FailNow()
	} else if json == nil {
		t.Errorf("json is nil")
		t.FailNow()
	} else {
		PrintJSON(t, json)
	}
	return json
}

func PrintJSON(t *testing.T, json *class.Map) {
	if json == nil {
		t.Errorf("json is nil")
		t.FailNow()
	} else {
		json_string, json_string_errors := json.ToJSONString()
		if json_string_errors != nil {
			t.Errorf("%s", json_string_errors)
			t.FailNow()
		} else if json_string == nil {
			t.Errorf("json_string is nil")
			t.FailNow()
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



