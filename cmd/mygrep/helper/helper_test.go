package helper

import (
	"fmt"
	"testing"
)

func TestIsWordChar(t *testing.T) {
	pattern := []string{"_", "alo_", "_alo", "haha", "1_h", "123"}
	for _, v := range pattern {
		for _, w := range v {
			if !IsWordChar(v) {
				t.Errorf("Expected alphanumeric or _ but got non alphanumeric: %v", w)
			}
			fmt.Printf("%s is alphanumeric or _\n", string(w))
		}
	}

	falsePattern := []string{"#", "%!", ":/?"}
	for _, v := range falsePattern {
		for _, w := range v {
			if IsWordChar(string(w)) {
				t.Errorf("Expected non alphanumeric but got alphanumeric or _: %v", w)
			}
			fmt.Printf("%s is not alphanumeric or _\n", string(w))
		}
	}
}

func TestProcessSquareBrackets(t *testing.T) {
	patterns := []string{"[abc]", "[^abc]", "[abc[df]]", "[ab[cd[^ef]]]"}

	for _, pattern := range patterns {
		ProcessSquareBrackets(&pattern)
		fmt.Println(pattern)
	}
}
