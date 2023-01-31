package scanning

// ScanTokens function to scan the tokens
func (s *Scanner) ScanTokens() {
	for s.isAtEnd() != true {
		// call scanToken function
		s.ScanToken()
		// update indexes - next lexeme
		s.start = s.current
	}

	s.tokens = append(s.tokens, Token{tokenType: "EOF", line: s.line + 1})
}
