package task6kyu

// Convert string to camel case
// https://www.codewars.com/kata/517abf86da9663f1d2000003

import (
	"bytes"
	"strings"
)

func ToCamelCase(s string) string {
	var str bytes.Buffer
	var do bool

	for _, ch := range s {
		if ch != '_' && ch != '-' {
			if do {
				ch = rune(strings.ToUpper(string(ch))[0])
			}

			str.WriteRune(ch)
			do = false
		} else {
			do = true
		}
	}

	return str.String()
}
