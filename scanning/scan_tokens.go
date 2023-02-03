package scanning

// ScanTokens function to scan the tokens
func (s *Scanner) ScanTokens() {
	s.current = 0
	s.start = 0
	for !s.isAtEnd() {
		// call scanToken function
		s.ScanToken()
		// update indexes - next lexeme
		s.start = s.current
	}
}
