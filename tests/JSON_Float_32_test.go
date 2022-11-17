package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

// low boundary
func TestCanParseFloat32Zero(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.00}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 0.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32PositiveLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.1234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32NegativeLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-0.1234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultiplePositiveLowBoundary(t *testing.T) {
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
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleNegativeLowBoundary(t *testing.T) {
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
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0.1234567890, \"key2\":-0.2234567890}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 0.1234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key2"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat3t2 is nil")
		} else if *value != -0.2234567890 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

// high boundary
func TestCanParseFloat32PositiveHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":340000000000000000000000000000000000000.00}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32NegativeHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-340000000000000000000000000000000000000.00}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -340000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultiplePositiveHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":340000000000000000000000000000000000000.00,\"key2\":339999999999999999999999999999999999999.99}")

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
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 339999999999999999999999999999999999999.99 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleNegativeHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-340000000000000000000000000000000000000.00,\"key2\":-339999999999999999999999999999999999999.99}")

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
		} else if *value != -340000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != -339999999999999999999999999999999999999.99 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":340000000000000000000000000000000000000.00, \"key2\":-339999999999999999999999999999999999999.99}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("key is not a *float32: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat32 is nil")
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key2"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("GetFloat3t2 is nil")
		} else if *value != -339999999999999999999999999999999999999.99 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}