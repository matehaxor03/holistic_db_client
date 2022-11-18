package tests
 
import (
    "testing"
	"math"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

// low boundary
func TestCanParseFloat64PositiveLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":" + fmt.Sprintf("%f", (math.MaxFloat32 * 10)) +"}")
	
	
	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != (math.MaxFloat32 * 10) {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64NegativeLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":" + fmt.Sprintf("%f", ((-1 * math.MaxFloat32 * 10))) +"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != (-1 * math.MaxFloat32 * 10) {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultiplePositiveLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":341234567890000000000000000000000000000.01,\"key2\":349876543210000000000000000000000000000.01}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 341234567890000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 349876543210000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultipleNegativeLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-341234567890000000000000000000000000000.01,\"key2\":-349876543210000000000000000000000000000.01}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -341234567890000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -349876543210000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultipleLowBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":341234567890000000000000000000000000000.01, \"key2\":-349876543210000000000000000000000000000.01}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 341234567890000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key2"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -349876543210000000000000000000000000000.01 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

// high boundary
func TestCanParseFloat64PositiveHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00}")
	
	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64NegativeHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultiplePositiveHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00,\"key2\":169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.99}")
	
	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat364 is nil")
		} else if *value != 169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.99 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultipleNegativeHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.98,\"key2\":-169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.99}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.98 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -169999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999.99 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat64MultipleHighBoundary(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00, \"key2\":-170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	}
	
	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*float64" {
		t.Errorf("key is not a *float64: %s", json.GetType("key"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != 170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float64" {
		t.Errorf("key2 is not a *float64: %s", json.GetType("key2"))
	} else {
		PrintJSON(t, json)
		value, value_errors := json.GetFloat64("key2") 

		if value_errors != nil {
			t.Errorf("map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("GetFloat64 is nil")
		} else if *value != -170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.00 {
			t.Errorf("expected: value actual: %f", *value)
		}
	}
}
