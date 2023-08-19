package iteration

import (
	"strings"
)

// Repeat return character repeated count times
func Repeat(character string, count int) string {
	return strings.Repeat(character, count)
}
