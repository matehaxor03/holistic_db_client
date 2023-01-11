package mysql

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
)

func box(options *json.Map, builder *strings.Builder, thing string, opening_symbol string, closing_symbol string) {
	runes := []rune(thing)
	slash := []byte("\\")

	if options.IsBoolFalse("use_file") {
		builder.WriteByte(slash[0])
	} 
	builder.WriteString(opening_symbol)
	for _, rune_string := range runes {
		if string(rune_string) == opening_symbol || string(rune_string) == closing_symbol {
			builder.WriteByte(slash[0])
			builder.WriteRune(rune_string)
		} else {
			builder.WriteRune(rune_string)
		}
	}
	
	if options.IsBoolFalse("use_file") {
		builder.WriteByte(slash[0])
	} 
	builder.WriteString(closing_symbol)
}