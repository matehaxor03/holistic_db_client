package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseFloat64Positive(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.1234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}

		if json != nil {
			json_string, _ := json.ToJSONString()
			fmt.Println(*json_string)
		}
	}
}

func TestCanParseFloat32Negative(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-0.1234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}

		if json != nil {
			json_string, _ := json.ToJSONString()
			fmt.Println(*json_string)
		}
	}
}

func TestCanParseFloat32MultiplePositive(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.1234567890,\"key2\":0.2234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCanParseFloat32MultipleNegative(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-0.1234567890,\"key2\":-0.2234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if json != nil {
		json_string, _ := json.ToJSONString()
		fmt.Println(*json_string)
	}
}

func TestCanParseFloat32Multiple(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.1234567890, \"key2\":-0.2234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}

		if json != nil {
			json_string, _ := json.ToJSONString()
			fmt.Println(*json_string)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key2"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat3t2 is nil")
		} else if *value != -0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}

		if json != nil {
			json_string, _ := json.ToJSONString()
			fmt.Println(*json_string)
		}
	}
}