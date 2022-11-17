package tests
 
import (
    "testing"
)

// int32 boundary low
func TestCanParseArrayContainingSingleInt32_neg32769(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32769]}")

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
		} else if *((*value)[0].(*int32)) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt32_neg32769_neg32770(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32769],\"key2\":[-32770]}")

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
		} else if *((*value)[0].(*int32)) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
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
		} else if *((*value)[0].(*int32)) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt32_neg32769_neg32770_neg32771_neg32772(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32769,-32770],\"key2\":[-32771,-32772]}")

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
		} else if *((*value)[0].(*int32)) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
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
		} else if *((*value)[0].(*int32)) != -32771 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -32772 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt32_neg32769_neg32770(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-32769,-32770]}")

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
		} else if *((*value)[0].(*int32)) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
		}
	}
}

// int32 boundary high
func TestCanParseArrayContainingSingleInt32_neg2147483648(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483648]}")

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
		} else if *((*value)[0].(*int32)) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt32_neg2147483648_neg2147483647(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483648],\"key2\":[-2147483647]}")

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
		} else if *((*value)[0].(*int32)) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
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
		} else if *((*value)[0].(*int32)) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt32_neg2147483648_neg2147483647_neg2147483646_neg2147483645(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483648,-2147483647],\"key2\":[-2147483646,-2147483645]}")

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
		} else if *((*value)[0].(*int32)) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
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
		} else if *((*value)[0].(*int32)) != -2147483646 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -2147483645 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
		}
	}
}

func TestCanParseArrayContainingMultipleInt32_neg2147483648_neg2147483647(t *testing.T) {
	json := ParseJSONSuccessfully(t, "{\"key\":[-2147483648,-2147483647]}")

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
		} else if *((*value)[0].(*int32)) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0].(*int32)))
		} else if *((*value)[1].(*int32)) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1].(*int32)))
		}
	}
}
