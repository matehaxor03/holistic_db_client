package tests
 
import (
    "testing"
)

// int8 boundary low boundary
func TestCanParseArrayContainingSingleInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-1]}")

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
		} else if *((*value)[0].(*int8)) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-1],\"key2\":[-2]}")

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
		} else if *((*value)[0].(*int8)) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
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
		} else if *((*value)[0].(*int8)) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-1,-2],\"key2\":[-3,-4]}")

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
		} else if *((*value)[0].(*int8)) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
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
		} else if *((*value)[0].(*int8)) != -3 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -4 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt8LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-1,-2]}")

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
		} else if *((*value)[0].(*int8)) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
		}
	}
}

// int8 boundary high
func TestCanParseArrayContainingSingleInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-128]}")

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
		} else if *((*value)[0].(*int8)) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-128],\"key2\":[-127]}")

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
		} else if *((*value)[0].(*int8)) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
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
		} else if *((*value)[0].(*int8)) != -127 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleHighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-128,-127],\"key2\":[-126,-125]}")

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
		} else if *((*value)[0].(*int8)) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -127 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
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
		} else if *((*value)[0].(*int8)) != -126 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -125 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt8HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-128,-127]}")

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
		} else if *((*value)[0].(*int8)) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int8)))
		} else if *((*value)[1].(*int8)) != -127 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int8)))
		}
	}
}
