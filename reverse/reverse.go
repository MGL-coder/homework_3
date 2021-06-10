package reverse

import "strings"

func Reverse(s string) string {
	var str strings.Builder

	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		str.WriteRune(runes[i])
	}

	return str.String()
}
