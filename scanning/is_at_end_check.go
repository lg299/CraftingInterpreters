package scanning

// check whether we are at the end of a source line
func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
