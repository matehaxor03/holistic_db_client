package tests
 
import (
    "testing"
)

// int64 boundary low
func TestCanParseArrayContainingSingleInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483649]}")

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
		} else if *((*value)[0].(*int64)) != -2147483649 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483649],\"key2\":[-2147483650]}")

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
		} else if *((*value)[0].(*int64)) != -2147483649 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
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
		} else if *((*value)[0].(*int64)) != -2147483650 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483649,-2147483650],\"key2\":[-2147483651,-2147483652]}")

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
		} else if *((*value)[0].(*int64)) != -2147483649 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -2147483650 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
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
		} else if *((*value)[0].(*int64)) != -2147483651 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -2147483652 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt64LowBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483649,-2147483650]}")

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
		} else if *((*value)[0].(*int64)) != -2147483649 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -2147483650 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
		}
	}
}

// int64 boundary high
func TestCanParseArrayContainingSingleInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-9223372036854775808]}")

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
		} else if *((*value)[0].(*int64)) != -9223372036854775808 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-9223372036854775808],\"key2\":[-9223372036854775807]}")

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
		} else if *((*value)[0].(*int64)) != -9223372036854775808 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
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
		} else if *((*value)[0].(*int64)) != -9223372036854775807 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt64HighBoundary(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-9223372036854775808,-9223372036854775807],\"key2\":[-9223372036854775806,-9223372036854775805]}")

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
		} else if *((*value)[0].(*int64)) != -9223372036854775808 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -9223372036854775807 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
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
		} else if *((*value)[0].(*int64)) != -9223372036854775806 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -9223372036854775805 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt64HighBoundaryh(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-9223372036854775808,-9223372036854775807]}")

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
		} else if *((*value)[0].(*int64)) != -9223372036854775808 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int64)))
		} else if *((*value)[1].(*int64)) != -9223372036854775807 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int64)))
		}
	}
}
