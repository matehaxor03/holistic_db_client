package tests
 
import (
    "testing"
	"fmt"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseNegativeInt8_1(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-1}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int8" {
			t.Errorf("key is not a *int8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt8 has errors")
			} else if value == nil {
				t.Errorf("GetInt8 is nil")
			} else if *value != -1 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt8_128(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-128}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int8" {
			t.Errorf("key is not a *int8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt8 has errors")
			} else if value == nil {
				t.Errorf("GetInt8 is nil")
			} else if *value != -128 {
				t.Errorf("expected: -128 actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt16_129(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-129}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int16" {
			t.Errorf("key is not a *int16: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt16("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt16 has errors")
			} else if value == nil {
				t.Errorf("GetInt16 is nil")
			} else if *value != -129 {
				t.Errorf("expected: -129  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt16_32768(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-32768}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int16" {
			t.Errorf("key is not a *int16: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt16("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt16 has errors")
			} else if value == nil {
				t.Errorf("GetInt16 is nil")
			} else if *value != -32768 {
				t.Errorf("expected: -32768  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt32_32769(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-32769}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int32" {
			t.Errorf("key is not a *int32: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt32 has errors")
			} else if value == nil {
				t.Errorf("GetInt32 is nil")
			} else if *value != -32769 {
				t.Errorf("expected: -32769  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt32_2147483648(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-2147483648}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int32" {
			t.Errorf("key is not a *int32: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt32 has errors")
			} else if value == nil {
				t.Errorf("GetInt32 is nil")
			} else if *value != -2147483648 {
				t.Errorf("expected: -2147483648  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt64_2147483649(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-2147483649}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int64" {
			t.Errorf("key is not a *int64: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt64 has errors")
			} else if value == nil {
				t.Errorf("GetInt64 is nil")
			} else if *value != -2147483649 {
				t.Errorf("expected: -2147483649  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseNegativeInt64_9223372036854775808(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-9223372036854775808}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*int64" {
			t.Errorf("key is not a *int64: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetInt64 has errors")
			} else if value == nil {
				t.Errorf("GetInt64 is nil")
			} else if *value != -9223372036854775808 {
				t.Errorf("expected: -9223372036854775808  actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCannotParseNegativeInt64_9223372036854775809(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":-9223372036854775809}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}

func TestCanParseUInt8_0(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":0}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint8" {
			t.Errorf("key is not a *uint8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt8 is nil")
			} else if *value != 0 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt8_255(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":255}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint8" {
			t.Errorf("key is not a *uint8: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt8("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt8 is nil")
			} else if *value != 255 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt16_256(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":256}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint16" {
			t.Errorf("key is not a *uint16: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt16("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt16 is nil")
			} else if *value != 256 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt16_65535(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":65535}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint16" {
			t.Errorf("key is not a *uint16: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt16("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt16 is nil")
			} else if *value != 65535 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt32_65536(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":65536}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint32" {
			t.Errorf("key is not a *uint32: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 65536 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt32_4294967295(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":4294967295}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint32" {
			t.Errorf("key is not a *uint32: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt32("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 4294967295 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt64_4294967296(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":4294967296}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint64" {
			t.Errorf("key is not a *uint64: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 4294967296 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCanParseUInt64_18446744073709551615(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":18446744073709551615}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json != nil {
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*uint64" {
			t.Errorf("key is not a *uint64: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetUInt64("key") 
	
			if value_errors != nil {
				t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
			} else if value == nil {
				t.Errorf("GetUInt32 is nil")
			} else if *value != 18446744073709551615 {
				t.Errorf("expected: value actual: %d", *value)
			}
		}

		if json != nil {
			json_string, json_string_errors := json.ToJSONString()
			if json_string_errors != nil {
				fmt.Println(json_string_errors)
			} else {
				fmt.Println(*json_string)
			}
		}
	} else {
		t.Errorf("json is nil")
	}	
}

func TestCannotParseUInt64_18446744073709551616(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":18446744073709551616}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}