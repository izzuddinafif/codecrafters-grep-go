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
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		fmt.Println("Pattern matches not found")
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern *string) (bool, error) {
	if utf8.RuneCountInString(*pattern) != 1 && !isSpecialPattern(pattern) {
		return false, fmt.Errorf("unsupported pattern: %q", *pattern)
	}

	var ok bool

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	switch *pattern {
	case "\\d":
		ok = bytes.ContainsAny(line, "0123456789")
	case "\\w":
		ok = isAlphanumeric(string(line))
	default:
		ok = bytes.ContainsAny(line, *pattern)
	}

	if ok {
		fmt.Println("Pattern matches found")
	}
	return ok, nil
}

func isSpecialPattern(pattern *string) bool {
	specialPattern := []string{"\\d", "\\w"}
	for _, ptrn := range specialPattern {
		if *pattern == ptrn {
			return true
		}
	}

	if strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]") {
		*pattern = extractBetweenBrackets(*pattern)
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

func extractBetweenBrackets(pattern string) string {
	start := strings.Index(pattern, "[")
	end := strings.Index(pattern, "]")

	return pattern[start+1 : end]
}
