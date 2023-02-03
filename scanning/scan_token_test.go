package scanning

import (
	"testing"
)

func TestAddTokensAppendsEmptyToken(t *testing.T) {
	s := Scanner{
		current: 0,
		start:   0,
		line:    1,
		tokens:  []Token{},
		source:  " ",
	}
	expectedTokens := []Token{
		{line: 1},
	}
	s.AddTokens(" ")
	assert.Equal(t, expectedTokens, s.tokens)
}

func TestAddTokensAppendsSingleCharToken() {
	s := Scanner{
		current: 0,
		start:   0,
		line:    1,
		tokens:  []Token{},
		source:  "+",
	}
	expectedTokens := []Token{
		{
			line:      1,
			lexeme:    "+",
			tokenType: "PLUS",
		},
	}
	s.AddTokens(" ")
	assert.Equal(t, expectedTokens, s.tokens)
}

