package tests
 
import (
    "testing"
)


func TestCanParseArrayContainingSingleUInteger_0(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInteger_0_1(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInteger_0_1_2_3(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInteger_0_1(t *testing.T) {
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