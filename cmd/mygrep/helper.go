package main

import (
	"strings"
	"unicode"
)

func isSpecialPattern(pattern *string) bool {
	specialPattern := []string{"\\d", "\\w"}
	for _, ptrn := range specialPattern {
		if *pattern == ptrn {
			return true
		}
	}
	return false
}

func isSpecialChar(pattern *string) bool {
	return (strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]")) || strings.HasPrefix(*pattern, "^")
}

func isAlphanumeric(line string) bool {
	for _, r := range line {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func extractBetweenSquareBrackets(pattern *string) {
	*pattern = (*pattern)[strings.Index(*pattern, "[")+1 : strings.Index(*pattern, "]")]
}

func extractAfterCaret(pattern *string) {
	*pattern = (*pattern)[strings.Index(*pattern, "^")+1:]
}

//func matchDigit(pattern *string) {}
