package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]
	fmt.Println(pattern)
	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, &pattern)
	fmt.Println(ok)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern *string) (bool, error) {
	if utf8.RuneCountInString(*pattern) != 1 && !isSpecialPattern(pattern) {
		return false, fmt.Errorf("unsupported pattern: %q", *pattern)
	}

	var ok bool

	if ok = checkPattern(line, pattern); ok {
		fmt.Println("Pattern matches found")
		return ok, nil
	}
	return ok, nil
}

func checkPattern(line []byte, pattern *string) bool {
	var ok bool
	switch {
	case *pattern == "\\d":
		ok = bytes.ContainsAny(line, "0123456789")
	case *pattern == "\\w":
		ok = isAlphanumeric(string(line))
	case strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]"):
		extractBetweenSquareBrackets(pattern)
		ok = bytes.ContainsAny(line, *pattern)
	case strings.HasPrefix(*pattern, "^"):
		extractAfterCaret(pattern)
		ok = !bytes.ContainsAny(line, *pattern)
	default:
		ok = bytes.ContainsAny(line, *pattern)
	}
	println(*pattern)
	return ok
}
func isSpecialPattern(pattern *string) bool {
	specialPattern := []string{"\\d", "\\w"}
	for _, ptrn := range specialPattern {
		if *pattern == ptrn {
			return true
		}
	}

	if strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]") {
		extractBetweenSquareBrackets(pattern)
		fmt.Println(*pattern)
		return true
	}

	if strings.HasPrefix(*pattern, "^") {
		extractAfterCaret(pattern)
		return true
	}
	return false
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
