package scanning

func (s *Scanner) ScanToken() {
	tokenTypes := TokenTypes()
	switch string(s.source[s.current]) {
	case "(":
		s.AddTokens(tokenTypes["("])
	case ")":
		s.AddTokens(tokenTypes[")"])
	case "{":
		s.AddTokens(tokenTypes["["])
	case "}":
		s.AddTokens(tokenTypes["]"])
	case ",":
		s.AddTokens(tokenTypes[","])
	case ".":
		s.AddTokens(tokenTypes["."])
	case "-":
		s.AddTokens(tokenTypes["-"])
	case "+":
		s.AddTokens(tokenTypes["+"])
	case ";":
		s.AddTokens(tokenTypes[";"])
	case "*":
		s.AddTokens(tokenTypes["*"])
	case "=":
		s.AddTokens(tokenTypes["="])
	default:
		break
	}

	s.current = s.current + 1
}

// AddTokens need to add the token to s.tokens
func (s *Scanner) AddTokens(token string) {
	newToken := Token{
		tokenType: token,
		lexeme:    s.source[s.start : s.current+1],
		line:      s.line,
	}
	s.tokens = append(s.tokens, newToken)
}
