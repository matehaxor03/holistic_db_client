package tests
 
import (
    "testing"
)

// uint8 boundary low
func TestCanParseArrayContainingSingleUInt8_0(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt8_0_1(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt8_0_1_2_3(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt8_0_1(t *testing.T) {
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
func TestCanParseArrayContainingSingleUInt8_255(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt8_254_255(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt8_252_253_254_255(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt8_254_255(t *testing.T) {
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

// uint16 boundary low
func TestCanParseArrayContainingSingleUInt16_256(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt16_256_257(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt16_256_257_258_259(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt16_256_257(t *testing.T) {
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
func TestCanParseArrayContainingSingleUInt16_65535(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt16_65534_65535(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt16_65532_65533_65534_65535(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt16_65534_65535(t *testing.T) {
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

// uint64 boundary low
func TestCanParseArrayContainingSingleUInt64_4294967296(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt64_4294967296_4294967297(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt64_4294967296_4294967297_4294967298_4294967299(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt64_4294967296_4294967297(t *testing.T) {
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
func TestCanParseArrayContainingSingleUInt64_18446744073709551615(t *testing.T) {
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

func TestCanParseMultipleArraysContainingSingleUInt64_18446744073709551614_18446744073709551615(t *testing.T) {
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

func TestCanParseMultipleArraysContainingMultipleUInt64_18446744073709551612_18446744073709551613_18446744073709551614_18446744073709551615(t *testing.T) {
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

func TestCanParseArrayContainingMultipleUInt64_18446744073709551614_18446744073709551615(t *testing.T) {
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
