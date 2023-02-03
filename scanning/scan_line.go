package scanning

import (
	"bufio"
	"os"
)

func scanLines(file *os.File) Scanner {
	s := Scanner{
		current: 0,
		start:   0,
		line:    0,
		tokens:  []Token{},
		source:  "",
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.line = s.line + 1
		line := scanner.Text()
		println("Line: ", line)
		s.source = line
		s.ScanTokens()
	}
	s.tokens = append(s.tokens, Token{tokenType: "EOF", line: s.line + 1})
	return s
}
