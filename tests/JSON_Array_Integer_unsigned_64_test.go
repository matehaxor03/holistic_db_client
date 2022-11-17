package tests
 
import (
    "testing"
)

// uint64 boundary low
func TestCanParseArrayContainingSingleUInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967296]}")

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
		} else if *((*value)[0].(*uint64)) != 4294967296 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967296],\"key2\":[4294967297]}")

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
		} else if *((*value)[0].(*uint64)) != 4294967296 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
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
		} else if *((*value)[0].(*uint64)) != 4294967297 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967296,4294967297],\"key2\":[4294967298,4294967299]}")

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
		} else if *((*value)[0].(*uint64)) != 4294967296 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 4294967297 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
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
		} else if *((*value)[0].(*uint64)) != 4294967298 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 4294967299 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967296,4294967297]}")

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
		} else if *((*value)[0].(*uint64)) != 4294967296 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 4294967297 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
		}
	}
}

// uint64 boundary high
func TestCanParseArrayContainingSingleUInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[18446744073709551615]}")

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
		} else if *((*value)[0].(*uint64)) != 18446744073709551615 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[18446744073709551614],\"key2\":[18446744073709551615]}")

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
		} else if *((*value)[0].(*uint64)) != 18446744073709551614 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
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
		} else if *((*value)[0].(*uint64)) != 18446744073709551615 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[18446744073709551612,18446744073709551613],\"key2\":[18446744073709551614,18446744073709551615]}")

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
		} else if *((*value)[0].(*uint64)) != 18446744073709551612 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 18446744073709551613 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
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
		} else if *((*value)[0].(*uint64)) != 18446744073709551614 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 18446744073709551615 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[18446744073709551614,18446744073709551615]}")

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
		} else if *((*value)[0].(*uint64)) != 18446744073709551614 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint64)))
		} else if *((*value)[1].(*uint64)) != 18446744073709551615 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint64)))
		}
	}
}
