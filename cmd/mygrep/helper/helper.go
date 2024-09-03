package helper

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func IsSpecialPattern(pattern *string) bool {
	return slices.Contains([]string{"\\d", "\\w"}, *pattern)
}

func IsSpecialChar(pattern *string) bool {
	return strings.Contains(*pattern, "[") || strings.Contains(*pattern, "^")
}

func IsWordChar(line string) bool {
	return strings.IndexFunc(line, func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '_'
	}) == -1
}
func IsWordCharacter(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'
}

func IsDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func WordChars() string {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
}

func Digits() string {
	return "0123456789"
}

func ExtractBetweenSquareBrackets(pattern string) string {
	return pattern[strings.Index(pattern, "[")+1 : strings.LastIndex(pattern, "]")]
}

func ExtractAfterCaret(pattern string) string {
	return pattern[strings.Index(pattern, "^")+1:]
}

func ProcessSquareBrackets(pattern *string) {
	if !strings.ContainsAny(*pattern, "[]") {
		return
	}

	if strings.Contains(*pattern, "[") {
		if !strings.Contains(*pattern, "]") {
			fmt.Println("invalid pattern: no enclosing square bracket")
			os.Exit(2)
		}
		openIndex := strings.Index(*pattern, "[")
		closeIndex := strings.LastIndex(*pattern, "]")
		before := (*pattern)[:openIndex]
		inside := ExtractBetweenSquareBrackets(*pattern)
		after := (*pattern)[closeIndex+1:]
		ProcessSquareBrackets(&inside)
		*pattern = before + inside + after
		return
	}
}
