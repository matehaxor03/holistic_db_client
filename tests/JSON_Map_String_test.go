package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

//json := ParseJSONSuccessfully(t, "{\"key\":{\"key2\":\"value2\"},\"key3\":\"value3\"}")


func TestCanParseEmptyMap(t *testing.T) {
	json, _ := class.ParseJSON("{\"key\":{}}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Map" {
		t.Errorf("key is not a *class.Map: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetMap("key") 

		if value_errors != nil {
			t.Errorf("map GetMap has errors")
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len(value.Keys()) != 0 {
			t.Errorf("expected key length: length=0 actual: length=%d", len(value.Keys()))
		}
	}
}

func TestCanParseNestedMapWithStringValue(t *testing.T) {
	json, _ := class.ParseJSON("{\"key\":{\"key1\":\"value1\"}}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Map" {
		t.Errorf("key is not a *class.Map: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len(value.Keys()) != 1 {
			t.Errorf("expected key length: length=1 actual: length=%d", len(value.Keys()))
		} else {
			inner_value, inner_value_errors := value.GetString("key1")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key1 has nil value")
			} else if *inner_value != "value1" {
				t.Errorf("expected key1:\"value1\" actual:\"%s\"", *inner_value)
			}
		}
	}
}

func TestCanParseNestedMapWithMultipleStringValue(t *testing.T) {
	json, _ := class.ParseJSON("{\"key\":{\"key1\":\"value1\",\"key2\":\"value2\"}}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Map" {
		t.Errorf("key is not a *class.Map: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len(value.Keys()) != 2 {
			t.Errorf("expected key length: length=1 actual: length=%d", len(value.Keys()))
		} else {
			inner_value, inner_value_errors := value.GetString("key1")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key1 has nil value")
			} else if *inner_value != "value1" {
				t.Errorf("expected key1:\"value1\" actual:\"%s\"", *inner_value)
			}

			inner_value2, inner_value2_errors := value.GetString("key2")
			if inner_value2_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
			} else if inner_value2 == nil {
				t.Errorf("key2 has nil value")
			} else if *inner_value2 != "value2" {
				t.Errorf("expected key1:\"value2\" actual:\"%s\"", *inner_value2)
			}
		}
	}
}

func TestCanParseDoubleNestedMapWithStringValue(t *testing.T) {
	json, _ := class.ParseJSON("{\"key\":{\"key2\":{\"key3\":\"value3\"}}}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Map" {
		t.Errorf("key is not a *class.Map: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len(value.Keys()) != 1 {
			t.Errorf("expected key length: length=1 actual: length=%d", len(value.Keys()))
		} else {
			inner_value, inner_value_errors := value.GetMap("key2")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key2 has nil value")
			} else if value.GetType("key2") != "*class.Map" {
				t.Errorf("key2 is not a *class.Map: %s", inner_value.GetType("key2"))
			} else {
				inner_value2, inner_value2_errors := inner_value.GetString("key3")
				if inner_value2_errors != nil {
					t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
				} else if inner_value2 == nil {
					t.Errorf("key3 has nil value")
				} else if *inner_value2 != "value3" {
					t.Errorf("expected key3:\"value3\" actual:\"%s\"", *inner_value2)
				}
			}
		}
	}
}

func TestCanParseDoubleNestedMapWithStringValueAndStringValueAtRootLevelAfter(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":{\"key2\":{\"key3\":\"value3\"}},\"key4\":\"value4\"}")

	if json_errors != nil {
		t.Errorf(fmt.Sprintf("%s", json_errors))
	} else if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Map" {
		t.Errorf("key is not a *class.Map: %s", json.GetType("key"))
	} else {		
		value, value_errors := json.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len(value.Keys()) != 1 {
			t.Errorf("expected keys length=1 actual: length=%d", len(value.Keys()))
		} else {
			inner_value, inner_value_errors := value.GetMap("key2")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key2 has nil value")
			} else if value.GetType("key2") != "*class.Map" {
				t.Errorf("key2 is not a *class.Map: %s", inner_value.GetType("key2"))
			} else {
				inner_value2, inner_value2_errors := inner_value.GetString("key3")
				if inner_value2_errors != nil {
					t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
				} else if inner_value2 == nil {
					t.Errorf("key3 has nil value")
				} else if *inner_value2 != "value3" {
					t.Errorf("expected key3:\"value3\" actual:\"%s\"", *inner_value2)
				}
			}
		}

		value4, inner_value4_errors := json.GetString("key4")
		if inner_value4_errors != nil {
			t.Errorf(fmt.Sprintf("%s", inner_value4_errors))
		} else if value4 == nil {
			t.Errorf("key4 has nil value")
		} else if *value4 != "value4" {
			t.Errorf("expected key4:\"value4\" actual:\"%s\"", *value4)
		}
	}
}
