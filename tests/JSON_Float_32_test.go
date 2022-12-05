package tests
 
import (
    "testing"
)

// low boundary
func TestCanParseFloat32Zero(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":0.0}")

	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 0.0 {
			t.Errorf("errror: expected: 0.0 actual: %f", *value)
		}
	}
}

func TestCanParseFloat32PositiveLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":0.123457}")

	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 0.123457 {
			t.Errorf("error: expected: 0.123457 actual: %f", *value)
		}
	}
}

func TestCanParseFloat32NegativeLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-0.123457}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -0.123457 {
			t.Errorf("error: expected: -0.123457 actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultiplePositiveLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":0.123457,\"key2\":0.223457}")

	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 0.123457 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("error: key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("error: key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 0.223457 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleNegativeLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-0.123457,\"key2\":-0.223457}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -0.123457 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("error: key2 not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -0.223457 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":0.123457, \"key2\":-0.223457}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat64 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat64 is nil")
		} else if *value != 0.123457 {
			t.Errorf("error: expected: 123457 actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("error: key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("error: key2 is not a *float32: %s", json.GetType("key2"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat3t2 is nil")
		} else if *value != -0.223457 {
			t.Errorf("error: expected: 0.223457 actual: %f", *value)
		}
	}
}

// high boundary
func TestCanParseFloat32PositiveHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":340000000000000000000000000000000000000.00}")

	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32NegativeHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-340000000000000000000000000000000000000.00}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -340000000000000000000000000000000000000.00 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultiplePositiveHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":340000000000000000000000000000000000000.00,\"key2\":339999999999999999999999999999999999999.99}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("error: key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("error: key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 339999999999999999999999999999999999999.99 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleNegativeHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":-340000000000000000000000000000000000000.00,\"key2\":-339999999999999999999999999999999999999.99}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -340000000000000000000000000000000000000.00 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("error: key2 not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key2 is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != -339999999999999999999999999999999999999.99 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}

func TestCanParseFloat32MultipleHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":340000000000000000000000000000000000000.00, \"key2\":-339999999999999999999999999999999999999.99}")
	
	if !json.HasKey("key") {
		t.Errorf("error: key not found")
	} else if json.GetType("key") != "*float32" {
		t.Errorf("error: key is not a *float32: %s", json.GetType("key"))
	} else {
		value, value_errors := json.GetFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat32 is nil")
		} else if *value != 340000000000000000000000000000000000000.00 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*float32" {
		t.Errorf("key2 is not a *float32: %s", json.GetType("key2"))
	} else {
		value, value_errors := json.GetFloat32("key2") 

		if value_errors != nil {
			t.Errorf("error: map GetFloat32 has errors")
		} else if value == nil {
			t.Errorf("error: GetFloat3t2 is nil")
		} else if *value != -339999999999999999999999999999999999999.99 {
			t.Errorf("error: expected: value actual: %f", *value)
		}
	}
}