package tests
 
import (
    "testing"
)

// uint32 boundary low
func TestCanParseArrayContainingSingleUInt32_65536(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65536]}")

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
		} else if *((*value)[0].(*uint32)) != 65536 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt32_65536_65537(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65536],\"key2\":[65537]}")

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
		} else if *((*value)[0].(*uint32)) != 65536 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
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
		} else if *((*value)[0].(*uint32)) != 65537 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt32_65536_65537_65538_65539(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65536,65537],\"key2\":[65538,65539]}")

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
		} else if *((*value)[0].(*uint32)) != 65536 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 65537 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
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
		} else if *((*value)[0].(*uint32)) != 65538 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 65539 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt32_65536_65537(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[65536,65537]}")

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
		} else if *((*value)[0].(*uint32)) != 65536 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 65537 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
		}
	}
}

// uint32 boundary high
func TestCanParseArrayContainingSingleUInt32_4294967295(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967295]}")

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
		} else if *((*value)[0].(*uint32)) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt32_4294967294_4294967295(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967294],\"key2\":[4294967295]}")

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
		} else if *((*value)[0].(*uint32)) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
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
		} else if *((*value)[0].(*uint32)) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt32_4294967292_4294967293_4294967294_4294967295(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967292,4294967293],\"key2\":[4294967294,4294967295]}")

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
		} else if *((*value)[0].(*uint32)) != 4294967292 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 4294967293 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
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
		} else if *((*value)[0].(*uint32)) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt32_4294967294_4294967295(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[4294967294,4294967295]}")

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
		} else if *((*value)[0].(*uint32)) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*uint32)))
		} else if *((*value)[1].(*uint32)) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*uint32)))
		}
	}
}
