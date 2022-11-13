package tests
 
import (
    "testing"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func TestCanParseString(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {			
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithSpaceBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{ \"key\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithNewlineBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\n\"key\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithDosNewlineBeforeKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\r\n\"key\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithSpaceAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\" :\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithNewlineAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\"\n:\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithDosNewlineAfterKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\"\r\n:\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithSpaceBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\": \"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithNewlineBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\n\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithDosNewlineBeforeValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\r\n\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithSpaceAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\" }")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithNewlineAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\"\n}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithDosNewlineAfterValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\"\r\n}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithQuoteKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke\\\"y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke\"y") {
			t.Errorf("key not found")
		} else if json.GetType("ke\"y") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("ke\"y"))
		} else {
			value, value_errors := json.GetString("ke\"y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithOpenBracketKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke{y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke{y") {
			t.Errorf("key not found")
		} else if json.GetType("ke{y") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("ke{y"))
		} else {
			value, value_errors := json.GetString("ke{y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCloseBracketKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke}y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke}y") {
			t.Errorf("key not found")
		} else if json.GetType("ke}y") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("ke}y"))
		} else {
			value, value_errors := json.GetString("ke}y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithOpenSquareBracketKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke[y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke[y") {
			t.Errorf("key not found")
		} else if json.GetType("ke[y") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("ke[y"))
		} else {
			value, value_errors := json.GetString("ke[y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCloseSquareBracketKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke]y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke]y") {
			t.Errorf("key not found")
		} else if json.GetType("ke]y") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("ke]y"))
		} else {
			value, value_errors := json.GetString("ke]y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCommaKey(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke,y\":\"value\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("ke,y") {
			t.Errorf("ke,y not found")
		} else if json.GetType("ke,y") != "*string" {
			t.Errorf("ke,y is not a string: %s", json.GetType("ke,y"))
		} else {
			value, value_errors := json.GetString("ke,y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}


func TestCanParseStringWithQuoteValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val\\\"ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val\"ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithOpenBracketValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val{ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val{ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCloseBracketValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val}ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val}ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithOpenSquareBracketValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val[ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)

		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val[ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCloseSquareBracketValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val]ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val]ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithCommaValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"val,ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string" {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val,ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}

func TestCanParseStringWithQuoteKeyAndValue(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"ke\\\"y\":\"val\\\"ue\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("ke\"y") {
			t.Errorf("ke\"y not found")
		} else if json.GetType("ke\"y") != "*string" {
			t.Errorf("ke\"y is not a string: %s", json.GetType("ke\"y"))
		} else {
			value, value_errors := json.GetString("ke\"y") 

			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "val\"ue" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}	
}


func TestCannotParseStringWithoutQuotePrefix(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":value\"}")

	if json_errors == nil {
		t.Errorf("expected errors for ParseJSON")
	}

	if json != nil {
		t.Errorf("expected nil json")
	}
}

func TestCannotParseStringWithoutQuoteSuffix(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value}")

	if json_errors == nil {
		t.Errorf("expected errors for ParseJSON")
	}

	if json != nil {
		t.Errorf("expected nil json")
	}
}

func TestCanParseMultipleStrings(t *testing.T) {
	json, json_errors := class.ParseJSON("{\"key\":\"value\",\"key2\":\"value2\"}")

	if json_errors != nil {
		t.Errorf("%s", json_errors)
	} else if json == nil {
		t.Errorf("json is nil")
	} else {
		PrintJSON(t, json)
		if !json.HasKey("key") {
			t.Errorf("key not found")
		} else if json.GetType("key") != "*string"  {
			t.Errorf("key is not a string: %s", json.GetType("key"))
		} else {
			value, value_errors := json.GetString("key") 
			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}

		if !json.HasKey("key2") {
			t.Errorf("key2 not found")
		} else if json.GetType("key2") != "*string"  {
			t.Errorf("key2 is not a string: %s", json.GetType("key2"))
		} else {
			value, value_errors := json.GetString("key2") 
			if value_errors != nil {
				t.Errorf("map GetString has errors")
			} else if value == nil {
				t.Errorf("GetString is nil")
			} else if *value != "value2" {
				t.Errorf("expected: value actual: %s", *value)
			}
		}
	}
}
