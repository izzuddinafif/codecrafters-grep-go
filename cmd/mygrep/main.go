package main

import (
	"fmt"
	"io"
	"os"
	"time"

	c "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/checker"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"))
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
	if c.CheckIfPatternIsNotSupported(pattern) {
		fmt.Println("why")
		return false, fmt.Errorf("unsupported pattern: %q", *pattern)
	}

	var ok bool
	if ok = c.CheckPatternMatch(line, pattern); ok {
		fmt.Println("Pattern matches found")
		return ok, nil
	}
	return ok, nil
}
