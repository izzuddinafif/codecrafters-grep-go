package checker

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	h "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/helper"
)

func CheckIfPatternIsNotSupported(pattern *string) bool {
	return utf8.RuneCountInString(*pattern) < 1 && !h.IsSpecialPattern(pattern) && !h.IsSpecialChar(pattern)
}

// func CheckPattern(line []byte, pattern *string) bool {
// 	var ok bool
// 	//lineIndex := 0
// 	for h.IsSpecialChar(pattern) {
// 		fmt.Println(*pattern)
// 		switch {
// 		case strings.HasPrefix(*pattern, "[") && strings.HasSuffix(*pattern, "]"):
// 			h.ProcessPattern(pattern)
// 			fmt.Println(*pattern)
// 		case strings.HasPrefix(*pattern, "^"):
// 			*pattern = h.ExtractAfterCaret(*pattern)
// 			if bytes.ContainsAny(line, *pattern) {
// 				return false
// 			}
// 			*pattern = (*pattern)[len(*pattern):]
// 			return true
// 		}
// 	}
// 	switch {
// 	case *pattern == "\\d":
// 		ok = bytes.ContainsAny(line, h.Digits())
// 	case *pattern == "\\w":
// 		ok = h.IsWordChar(string(line))
// 	default:
// 		ok = bytes.ContainsAny(line, *pattern)
// 	}
// 	println(*pattern)
// 	return ok
// }

func CheckPatternMatch(line []byte, pattern *string) bool {
	if len(*pattern) == 0 {
		return true
	}
	if len(line) == 0 {
		return false
	}
	if (*pattern)[0] == '^' {
		i := bytes.Index([]byte(*pattern), []byte{byte('^')})
		fmt.Println(i)
		if i+1 >= len(*pattern) {
			fmt.Println("here")
			return false
		}
		str := (*pattern)[i+1:]
		fmt.Println(str, string(line))
		return bytes.HasPrefix(line, []byte(str))
	} else if (*pattern)[len(*pattern)-1] == '$' {
		i := bytes.Index([]byte(*pattern), []byte{byte('$')})
		fmt.Println(i)
		if i >= len(*pattern) {
			fmt.Println("here")
			return false
		}
		str := (*pattern)[:len(*pattern)-1]
		return bytes.HasSuffix(line, []byte(str))
	}
	// Try to match the pattern starting from each position in the line
	for j := 0; j < len(line); j++ {
		fmt.Println("iteration", j)
		if matchFrom(line[j:], pattern) {
			return true
		}
	}
	fmt.Println("no match")
	return false
}

func matchFrom(line []byte, pattern *string) bool {
	var i, j int
	for i < len(*pattern) && j < len(line) {
		pt := rune((*pattern)[i])
		li := rune(line[j])
		fmt.Println("pt = ", string(pt))
		if pt == '\\' {
			i++
			if i >= len(*pattern) {
				return false
			}
			pt = rune((*pattern)[i])
			fmt.Println(string(pt), string(li))
			if !matchSpecialCharacter(pt, li) {
				return false
			}
		} else if pt == '[' {
			end := findClosingBracket(*pattern, i)
			if end == -1 {
				return false
			}
			if i >= len(*pattern) {
				return false
			}
			str := (*pattern)[i+1 : end]
			if str[0] == '^' {
				str = str[1:]
				fmt.Println(string(li), str)
				if bytes.ContainsAny([]byte(string(li)), str) {
					return false
				}
			} else {
				fmt.Println(string(li), str)
				if !bytes.ContainsAny([]byte(string(li)), str) {
					return false
				}
				fmt.Println(i)
			}
			i = end
		} else if pt == '.' {
			i++
			j++
			continue
		} else if pt == '(' {
			end := strings.LastIndex(*pattern, ")")
			if end == -1 {
				return false
			}
			if i >= len(*pattern) {
				return false
			}
			str := (*pattern)[i+1 : end]
			orIndex := strings.Index(str, "|")
			str2 := str[:orIndex]
			str1 := str[orIndex+1:]
			ln := line[j:]
			fmt.Println("ln: ", string(ln))
			fmt.Println(str1, str2)
			if !matchFrom(ln, &str1) && !matchFrom(ln, &str2) {
				return false
			}
			fmt.Println("this should print")
			i = end + 1
			continue
		} else {
			if i < len(*pattern)-1 && rune((*pattern)[i+1]) == '+' {
				fmt.Println(string(pt), string(rune((*pattern)[i+1])))
				if pt != li {
					return false
				}
				// match as many as possible
				for j < len(line) && line[j] == byte(pt) {
					j++
				}
				i += 2 // skip the character and '+'
				continue
			} else if i < len(*pattern)-1 && rune((*pattern)[i+1]) == '?' {
				fmt.Println(string(pt), string(rune((*pattern)[i+1])))
				for j < len(line) && line[j] == byte(pt) {
					j++
				}
				i += 2 // skip the character and '?'
				continue
			} else if pt != li { // normal pattern
				return false
			}
		}
		i++
		j++
		fmt.Println(i, j)
		fmt.Println("happy path")
	}
	fmt.Println(i, len(*pattern))
	return i == len(*pattern)
}

func findClosingBracket(pattern string, startIdx int) int {
	for i := startIdx + 1; i < len(pattern); i++ {
		if pattern[i] == ']' {
			return i
		}
	}
	return -1
}

func matchSpecialCharacter(patternChar rune, lineChar rune) bool {
	switch patternChar {
	case 'w':
		return h.IsWordCharacter(lineChar)
	case 'd':
		return h.IsDigit(lineChar)
	default:
		return false
	}
}
