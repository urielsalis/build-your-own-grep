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

// Usage: echo <input_text> | your_grep.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	switch {
	case pattern == `\d`:
		for _, char := range string(line) { // ranging over string handles utf-8 multibyte characters correctly
			if unicode.IsDigit(char) {
				return true, nil
			}
		}

		return false, nil
	case pattern == `\w`:
		for _, char := range string(line) {
			if unicode.IsLetter(char) {
				return true, nil
			}
		}

		return false, nil
	case strings.HasPrefix(pattern, "[^"):
		end := strings.IndexByte(pattern, ']')
		charset := pattern[2:end]

		return !bytes.ContainsAny(line, charset), nil
	case strings.HasPrefix(pattern, "["):
		end := strings.IndexByte(pattern, ']')
		charset := pattern[1:end]

		return bytes.ContainsAny(line, charset), nil
	case utf8.RuneCountInString(pattern) == 1:
		return bytes.ContainsAny(line, pattern), nil
	}

	return false, fmt.Errorf("unsupported pattern: %q", pattern)
}
