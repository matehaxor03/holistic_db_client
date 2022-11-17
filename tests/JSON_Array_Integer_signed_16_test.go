package tests
 
import (
    "testing"
)

// int16 boundary low boundary
func TestCanParseArrayContainingSingleInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-129]}")

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
		} else if *((*value)[0].(*int16)) != -129 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleLowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-129],\"key2\":[-130]}")

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
		} else if *((*value)[0].(*int16)) != -129 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
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
		} else if *((*value)[0].(*int16)) != -130 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-129,-130],\"key2\":[-131,-132]}")

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
		} else if *((*value)[0].(*int16)) != -129 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -130 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
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
		} else if *((*value)[0].(*int16)) != -131 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -132 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt16LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-129,-130]}")

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
		} else if *((*value)[0].(*int16)) != -129 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -130 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
		}
	}
}

// int16 boundary high
func TestCanParseArrayContainingSingleInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32768]}")

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
		} else if *((*value)[0].(*int16)) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32768],\"key2\":[-32767]}")

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
		} else if *((*value)[0].(*int16)) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
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
		} else if *((*value)[0].(*int16)) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32768,-32767],\"key2\":[-32766,-32765]}")

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
		} else if *((*value)[0].(*int16)) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
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
		} else if *((*value)[0].(*int16)) != -32766 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -32765 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt16HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32768,-32767]}")

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
		} else if *((*value)[0].(*int16)) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int16)))
		} else if *((*value)[1].(*int16)) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int16)))
		}
	}
}
