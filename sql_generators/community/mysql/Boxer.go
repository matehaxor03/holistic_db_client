package mysql

import (
	"strings"
)

func Box(builder *strings.Builder, thing string, opening_symbol string, closing_symbol string) {
	runes := []rune(thing)
	slash := []byte("\\")

	builder.WriteString(opening_symbol)
	for _, rune_string := range runes {
		if string(rune_string) == opening_symbol || string(rune_string) == closing_symbol {
			builder.WriteByte(slash[0])
			builder.WriteRune(rune_string)
		} else {
			builder.WriteRune(rune_string)
		}
	}
	builder.WriteString(closing_symbol)
}