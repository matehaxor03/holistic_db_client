package tests
 
import (
    "testing"
)

// uint16 boundary low
func TestCanParseArrayContainingSingleUInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[256]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 256 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[256],\"key2\":[257]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 256 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*class.Array" {
		t.Errorf("key2 is not a *class.Array: %s", json.GetType("key2"))
	} else {			
		value, value_errors := json.GetArray("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 257 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[256,257],\"key2\":[258,259]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 256 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 257 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*class.Array" {
		t.Errorf("key2 is not a *class.Array: %s", json.GetType("key2"))
	} else {			
		value, value_errors := json.GetArray("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 258 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 259 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[256,257]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 256 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 257 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}
}

// uint16 boundary high
func TestCanParseArrayContainingSingleUInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65535]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65535 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65534],\"key2\":[65535]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65534 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*class.Array" {
		t.Errorf("key2 is not a *class.Array: %s", json.GetType("key2"))
	} else {			
		value, value_errors := json.GetArray("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65535 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65532,65533],\"key2\":[65534,65535]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65532 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 65533 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}

	if !json.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json.GetType("key2") != "*class.Array" {
		t.Errorf("key2 is not a *class.Array: %s", json.GetType("key2"))
	} else {			
		value, value_errors := json.GetArray("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65534 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 65535 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65534,65535]}")

	if !json.HasKey("key") {
		t.Errorf("key not found")
	} else if json.GetType("key") != "*class.Array" {
		t.Errorf("key is not a *class.Array: %s", json.GetType("key"))
	} else {			
		value, value_errors := json.GetArray("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if *((*value)[0].(*uint16)) != 65534 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint16)))
		} else if *((*value)[1].(*uint16)) != 65535 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint16)))
		}
	}
}
