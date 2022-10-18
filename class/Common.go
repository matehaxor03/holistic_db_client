package class

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
	"math/rand"
)

func EscapeString(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "'", "\\'")
	return value
}

func CloneString(value *string) *string {
	if value == nil {
		return nil
	}

	temp := strings.Clone(*value)
	return &temp
}

func FIELD_NAME_VALIDATION_FUNCTIONS() string {
	return "validation_functions"
}

func FIELD_NAME_VALIDATION_FUNCTIONS_PARAMETERS() string {
	return "validation_functions_parameters"
}

func getWhitelistStringFunc() *func(m Map) []error {
	function := WhiteListString
	return &function
}

func WhiteListString(m Map) []error {
	var errors []error
	map_values := m.M("values")
	str := m.S("value")
	label := m.S("label")
	data_type := m.S("data_type")

	if map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	_, found := (*map_values)[*str]

	if !found {
		errors = append(errors, fmt.Errorf("%s: %s: WhiteListString: did not find value", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func ArraysContainsArraysOrdered(a [][]string, b [][]string, label string, typeof string) []error {
	var errors []error

	for _, b_value := range b {
		var current = strings.Join(b_value, "")
		var found = false
		for _, a_value := range a {
			var compare = strings.Join(a_value, "")

			if current == compare {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("%s %s has value '%s' expected to have value in '%s", typeof, label, b_value, a))
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func getWhitelistCharactersFunc() *func(m Map) []error {
	funcValue := WhitelistCharacters
	return &funcValue
}

func WhitelistCharacters(m Map) []error {
	var errors []error
	map_values := m.M("values")
	str := m.S("value")
	label := m.S("label")
	data_type := m.S("data_type")

	if map_values == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has nil map", *data_type, *label))
	} else if len(*map_values) == 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has empty array", *data_type, *label))
	}

	if str == nil {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: compare value is nil", *data_type, *label))
	} else if *str == "" {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: compare value is empty", *data_type, *label))
	}

	if len(errors) > 0 {
		return errors
	}

	var invalid_letters []string
	for _, letter_rune := range *str {
		letter_string := string(letter_rune)
		_, found := (*map_values)[letter_string]

		if !found {
			invalid_letters = append(invalid_letters, letter_string)
		}
	}

	if len(invalid_letters) > 0 {
		errors = append(errors, fmt.Errorf("%s: %s: WhitelistCharacters: has invalid character(s): %s", *data_type, *label, invalid_letters))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func GetConstantValueAllowedCharacters() Map {
	temp := Map{"a": nil,
		"b": nil,
		"c": nil,
		"d": nil,
		"e": nil,
		"f": nil,
		"g": nil,
		"h": nil,
		"i": nil,
		"j": nil,
		"k": nil,
		"l": nil,
		"m": nil,
		"n": nil,
		"o": nil,
		"p": nil,
		"q": nil,
		"r": nil,
		"s": nil,
		"t": nil,
		"u": nil,
		"v": nil,
		"w": nil,
		"x": nil,
		"y": nil,
		"z": nil,
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
		"0": nil,
		"1": nil,
		"2": nil,
		"3": nil,
		"4": nil,
		"5": nil,
		"6": nil,
		"7": nil,
		"8": nil,
		"9": nil,
		"-": nil,
		"_": nil,
		".": nil}
	return temp
}

func GetAllowedStringValues() Map {
	temp := Map{"a": nil,
		"b": nil,
		"c": nil,
		"d": nil,
		"e": nil,
		"f": nil,
		"g": nil,
		"h": nil,
		"i": nil,
		"j": nil,
		"k": nil,
		"l": nil,
		"m": nil,
		"n": nil,
		"o": nil,
		"p": nil,
		"q": nil,
		"r": nil,
		"s": nil,
		"t": nil,
		"u": nil,
		"v": nil,
		"w": nil,
		"x": nil,
		"y": nil,
		"z": nil,
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
		"0": nil,
		"1": nil,
		"2": nil,
		"3": nil,
		"4": nil,
		"5": nil,
		"6": nil,
		"7": nil,
		"8": nil,
		"9": nil,
		"-": nil,
		"_": nil,
		".": nil,
		"*": nil}
	return temp
}

func ValidateOptions(extra_options map[string]map[string][][]string, reflect_value reflect.Value) []error {
	var errors []error
	var VALID_CHARACTERS = GetConstantValueAllowedCharacters()
	for key, value := range extra_options {
		filters1 := Map{"values": VALID_CHARACTERS, "value": &key, "label": fmt.Sprintf("extra_options root key %s", key), "typeOf": fmt.Sprintf("%T", reflect_value)}

		key_extra_options_errors := WhitelistCharacters(filters1)
		if key_extra_options_errors != nil {
			errors = append(errors, key_extra_options_errors...)
		}

		for key2, value2 := range value {
			var combined = ""
			for _, array_value := range value2 {
				combined += strings.Join(array_value, "")
			}

			filters2 := Map{"values": VALID_CHARACTERS, "value": &combined, "label": fmt.Sprintf("extra_options sub key: %s value %s", key2, value2), "typeOf": fmt.Sprintf("%T", reflect_value)}

			value_extra_options_errors := WhitelistCharacters(filters2)
			if value_extra_options_errors != nil {
				errors = append(errors, value_extra_options_errors...)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func GetLogicCommand(command string, field_name string, allowed_options map[string]map[string][][]string, options map[string]map[string][][]string, typeof string) (*string, []error) {
	var errors []error
	logic_option := ""

	if options == nil {
		return &logic_option, nil
	}

	logic_option_value, logic_option_exists := options[field_name]
	if !logic_option_exists {
		return &logic_option, nil
	}

	allowed_logic_option_value, allowed_logic_option_exists := allowed_options[field_name]

	if !allowed_logic_option_exists {
		errors = append(errors, fmt.Errorf("%s field: %s is not supported but was provided", typeof, field_name))
		return nil, errors
	}

	logic_option_command_value, logic_option_command_exists := logic_option_value[command]
	if !logic_option_command_exists {
		return &logic_option, nil
	}

	allowed_logic_option_command_value, allowed_logic_option_command_exists := allowed_logic_option_value[command]

	if !allowed_logic_option_command_exists {
		errors = append(errors, fmt.Errorf("%s field: %s is not supported for command: %s but was provided", typeof, field_name, command))
		return nil, errors
	}

	logic_option_errors := ArraysContainsArraysOrdered(allowed_logic_option_command_value, logic_option_command_value, field_name+"->"+command, typeof)
	if logic_option_errors != nil {
		errors = append(errors, logic_option_errors...)
		return nil, errors
	}

	for _, statements := range logic_option_command_value {
		logic_option += strings.Join(statements, " ") + " "
	}

	return &logic_option, nil
}

func ValidateData(fields Map, structType string) []error {	
	var errors []error
	var parameters = fields.Keys()

	for _, parameter := range parameters {

		if fields.GetType(parameter) != "class.Map" {
			errors = append(errors, fmt.Errorf("table: %s column: %s is not of type class.Map", structType, parameter))
		}

		if len(errors) > 0 {
			continue
		}

		parameter_fields := fields.M(parameter)

		value_is_mandatory := true
		value_is_null := parameter_fields.IsNil("value")
		mandatory_field, mandatory_field_errors := parameter_fields.GetBool("mandatory")

		if mandatory_field_errors != nil {
			errors = append(errors, mandatory_field_errors...)
		}

		if mandatory_field != nil {
			value_is_mandatory = *mandatory_field
		}

		attribute_to_validate := "value"
		if value_is_null && value_is_mandatory && parameter_fields.IsNil("default") {
			if parameter_fields.IsBoolFalse("primary_key") {
				fmt.Println(fields.ToJSONString())

				errors = append(errors, fmt.Errorf("table: %s parameter: %s is mandatory but primary key is nil and default is nil", structType, parameter))
				continue
			} else if parameter_fields.IsBoolTrue("primary_key") {
				primary_key_value, primary_key_value_errors := parameter_fields.GetBool("primary_key")
				if primary_key_value_errors != nil {
					errors = append(errors, primary_key_value_errors...)
					continue
				} else if *primary_key_value == false {
					errors = append(errors, fmt.Errorf("table: %s parameter: %s is mandatory has primary key but it was false default is nil", structType, parameter))
					continue
				}
			}

			if parameter_fields.IsBoolFalse("auto_increment") {
				errors = append(errors, fmt.Errorf("table: %s parameter: %s is mandatory but auto_increment was nil and default is nil", structType, parameter))
			} else if parameter_fields.IsBoolTrue("auto_increment") {
				auto_increment_value, auto_increment_value_errors := parameter_fields.GetBool("auto_increment")
				if auto_increment_value_errors != nil {
					errors = append(errors, auto_increment_value_errors...)
					continue
				} else if *auto_increment_value == false {
					errors = append(errors, fmt.Errorf("table: %s parameter: %s is mandatory but auto_increment was false and default is nil",structType, parameter))
					continue
				}
			}

			continue
		} else if value_is_null && value_is_mandatory && !parameter_fields.IsNil("default") {
			attribute_to_validate = "default"
		} else if value_is_null && !value_is_mandatory && parameter_fields.IsNil("default") {
			continue
		}

		typeOf := fmt.Sprintf("%T", (*parameter_fields)[attribute_to_validate])

		switch typeOf {
		case "map[string]map[string][][]string":
			// todo: convert these to objects for validations
			break
		case "string":
		case "*string":
			string_value := parameter_fields.S(attribute_to_validate)
			if parameter_fields.GetType(FILTERS()) != "class.Array" {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s is not an array", structType, parameter, FILTERS()))
				continue
			}

			filters := parameter_fields.A(FILTERS())

			if len(filters) == 0 {
				errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s has no filters", structType, parameter, FILTERS()))
				continue
			}

			for filter_index, filter := range filters {
				filter_map := filter.(Map)

				if !filter_map.HasKey("function") {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is empty", structType, parameter, FILTERS(), filter_index))
					continue
				}

				function := filter_map.Func("function")
				if function == nil {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d function is nil", structType, parameter, FILTERS(), filter_index))
					continue
				}

				if filter_map.GetType("values") == "nil" || filter_map.GetType("values") == "<nil>" {
					errors = append(errors, fmt.Errorf("table: %s column: %s attribute: %s at index: %d values is nil", structType, parameter, FILTERS(), filter_index))
					continue
				}

				filter_map.SetString("value", string_value)
				filter_map.SetString("data_type", &structType)
				filter_map.SetString("label", &parameter)

				function_errors := function(filter.(Map))
				if function_errors != nil {
					errors = append(errors, function_errors...)
				}
				/*
					var vargsConvert = []reflect.Value{reflect.ValueOf(filter)}

					var output_array_map_result = reflect.ValueOf(function).Call(vargsConvert)

					validation_errors := ConvertPrimitiveReflectValueArrayToArray(output_array_map_result)
					outer_array_length := len(validation_errors)
					for i := 0; i < outer_array_length; i++ {
						validation_error := validation_errors[i]
						error_value := fmt.Sprintf("%s", reflect.ValueOf(validation_error).Interface())
						if error_value == "[]" {
							continue
						}
						errors = append(errors, fmt.Errorf(error_value))
					}*/
			}

			break
		case "*int", "int":
			_, value_of_errors := parameter_fields.GetInt(attribute_to_validate)
			if value_of_errors != nil {
				errors = append(errors, value_of_errors...)
				continue
			}
		case "*bool", "bool":
			_, value_of_errors := parameter_fields.GetBool(attribute_to_validate)
			if value_of_errors != nil {
				errors = append(errors, value_of_errors...)
				continue
			}
		case "*int64", "int64":
			_, value_of_errors := parameter_fields.GetInt64(attribute_to_validate)
			if value_of_errors != nil {
				errors = append(errors, value_of_errors...)
				continue
			}
		case "*uint64", "uint64":
			_, value_of_errors := parameter_fields.GetUInt64(attribute_to_validate)
			if value_of_errors != nil {
				errors = append(errors, value_of_errors...)
				continue
			}
		case "*time.Time":
			_, value_of_errors := parameter_fields.GetTime(attribute_to_validate)
			if value_of_errors != nil {
				errors = append(errors, value_of_errors...)
				continue
			}
		case "*class.Database":
			database := parameter_fields.GetObject(attribute_to_validate).(*Database)

			errors_for_database := database.Validate()
			if errors_for_database != nil {
				errors = append(errors, errors_for_database...)
			}
			break
		case "*class.DomainName":
			domain_name := parameter_fields.GetObject(attribute_to_validate).(*DomainName)

			errors_for_domain_name := domain_name.Validate()
			if errors_for_domain_name != nil {
				errors = append(errors, errors_for_domain_name...)
			}
			break
		case "*class.Host":
			host := parameter_fields.GetObject(attribute_to_validate).(*Host)

			errors_for_host := host.Validate()
			if errors_for_host != nil {
				errors = append(errors, errors_for_host...)
			}

			break
		case "*class.Credentials":
			credentials := parameter_fields.GetObject(attribute_to_validate).(*Credentials)

			errors_for_credentaials := credentials.Validate()
			if errors_for_credentaials != nil {
				errors = append(errors, errors_for_credentaials...)
			}

			break
		case "*class.DatabaseCreateOptions":
			database_create_options := parameter_fields.GetObject(attribute_to_validate).(*DatabaseCreateOptions)

			errors_for_database_create_options := database_create_options.Validate()
			if errors_for_database_create_options != nil {
				errors = append(errors, errors_for_database_create_options...)
			}

			break
		case "*class.Client":
			client := parameter_fields.GetObject(attribute_to_validate).(*Client)

			errors_for_client := client.Validate()
			if errors_for_client != nil {
				errors = append(errors, errors_for_client...)
			}

			break
		case "*class.Grant":
			grant := parameter_fields.GetObject(attribute_to_validate).(*Grant)

			errors_for_grant := grant.Validate()
			if errors_for_grant != nil {
				errors = append(errors, errors_for_grant...)
			}

			break
		case "*class.User":
			user := parameter_fields.GetObject(attribute_to_validate).(*User)

			errors_for_user := user.Validate()
			if errors_for_user != nil {
				errors = append(errors, errors_for_user...)
			}

			break
		case "*class.Table":
			table := parameter_fields.GetObject(attribute_to_validate).(*Table)

			errors_for_table := table.Validate()
			if errors_for_table != nil {
				errors = append(errors, errors_for_table...)
			}

			break
		/*
			case "<nil>":
				if !parameter_fields.HasKey("default") {
					if value_is_null && value_is_mandatory {
						errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil", parameter))
					}
				} else if !parameter_fields.HasKey("type") {
					if value_is_null && value_is_mandatory {
						errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil and did not have type", parameter))
					}
				} else {
					typeOf := parameter_fields.S("type")
					switch *typeOf {
					case "*time.Time":
						type_of_default := parameter_fields.GetType("default")
						switch type_of_default {
						case "*string":
							type_of_default_value := parameter_fields.S("default")
							if *type_of_default_value != "now" {
								errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s and had default: %s please implement default value: %s", parameter, *typeOf, type_of_default, *type_of_default_value))
							}
						default:
							errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s please implement default for %s", parameter, *typeOf, type_of_default))
						}
					case "*int64", "int64":
						type_of_default := parameter_fields.GetType("default")
						switch type_of_default {
						case "*int", "int", "*int64", "int64":
							type_of_default_value, type_of_default_value_errors := parameter_fields.GetInt64("default")
							if type_of_default_value_errors != nil {
								errors = append(errors, type_of_default_value_errors...)
							} else if type_of_default_value == nil {
								errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s and had default: %s please implement default value: %s", parameter, *typeOf, type_of_default, "nil"))
							}
						default:
							errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil with type: %s please implement default for %s", parameter, *typeOf, type_of_default))
						}
					default:
						errors = append(errors, fmt.Errorf("parameter: %s is mandatory but was nil please implement for type: %s", parameter, *typeOf))
					}
				}
		*/
		default:
			fmt.Println(fields.ToJSONString())
			panic(fmt.Sprintf("please implement type: %s for parameter: %s", typeOf, parameter))
		}

	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func GetTimeNow() *time.Time {
	now := time.Now()
	return &now
}

func FormatTime(value time.Time) string {
	return value.Format("2006-01-02 15:04:05.000000")
}

func GetTimeNowString() string {
	return (*GetTimeNow()).Format("2006-01-02 15:04:05.000000")
}

func GenerateRandomLetters(length uint64, upper_case *bool) (*string) {
	rand.Seed(time.Now().UnixNano())
	
	var letters_to_use string
	uppercase_letters :=  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase_letters := "abcdefghijklmnopqrstuvwxyz"

	if upper_case == nil {
		letters_to_use = uppercase_letters + lowercase_letters
	} else if *upper_case {
		letters_to_use = uppercase_letters
	} else {
		letters_to_use = lowercase_letters
	}

	var sb strings.Builder

	l := len(letters_to_use)

	for i := uint64(0); i < length; i++ {
		c := letters_to_use[rand.Intn(l)]
		sb.WriteByte(c)
	}

	value := sb.String()
	return &value
}
