package main

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func checkIfPatternIsSupported(pattern *string) bool {
	return utf8.RuneCountInString(*pattern) != 1 && !isSpecialPattern(pattern) && !isSpecialChar(pattern)
}

func checkPattern(line []byte, pattern *string) bool {
	var ok bool

	for checkIfPatternIsSupported(pattern) {
		switch {
		case *pattern == "\\d":
			ok = bytes.ContainsAny(line, "0123456789")
		case *pattern == "\\w":
			ok = isAlphanumeric(string(line))
		case isSpecialChar(pattern):
			//fmt.Println("its a special char")
			switch {
			case strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]"):
				extractBetweenSquareBrackets(pattern)
				ok = bytes.ContainsAny(line, *pattern)
			case strings.HasPrefix(*pattern, "^"):
				//fmt.Println("here")
				extractAfterCaret(pattern)
				//fmt.Println(*pattern)
				ok = !bytes.ContainsAny(line, *pattern)
			}
		default:
			ok = bytes.ContainsAny(line, *pattern)
		}
		println(*pattern)
	}
	return ok
}
