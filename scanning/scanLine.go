package scanning

import (
	"bufio"
	"os"
)

func scanLine(file *os.File) Scanner {
	s := Scanner{current: 0, start: 0, line: 1, tokens: []Token{}, source: ""}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		s.source = line
		s.ScanTokens()

		s.line = s.line + 1
	}
	return s
}
