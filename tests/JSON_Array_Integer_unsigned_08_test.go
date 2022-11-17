package tests
 
import (
    "testing"
)

// uint8 boundary low
func TestCanParseArrayContainingSingleUInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[0]}")

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
		} else if *((*value)[0].(*uint8)) != 0 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[0],\"key2\":[1]}")

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
		} else if *((*value)[0].(*uint8)) != 0 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
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
		} else if *((*value)[0].(*uint8)) != 1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[0,1],\"key2\":[2,3]}")

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
		} else if *((*value)[0].(*uint8)) != 0 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
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
		} else if *((*value)[0].(*uint8)) != 2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 3 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[0,1]}")

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
		} else if *((*value)[0].(*uint8)) != 0 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
		}
	}
}

// uint8 boundary high
func TestCanParseArrayContainingSingleUInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[255]}")

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
		} else if *((*value)[0].(*uint8)) != 255 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[254],\"key2\":[255]}")

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
		} else if *((*value)[0].(*uint8)) != 254 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
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
		} else if *((*value)[0].(*uint8)) != 255 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[252,253],\"key2\":[254,255]}")

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
		} else if *((*value)[0].(*uint8)) != 252 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 253 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
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
		} else if *((*value)[0].(*uint8)) != 254 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 255 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[254,255]}")

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
		} else if *((*value)[0].(*uint8)) != 254 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint8)))
		} else if *((*value)[1].(*uint8)) != 255 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint8)))
		}
	}
}
