package scanning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanReturnsEmptyListOfTokens(t *testing.T) {
	tokens, err := Scan("./test_data/empty_file.txt")
	assert.Nil(t, err)
	expectedTokens := []Token{
		{
			tokenType: EOF,
			line:      1,
			lexeme:    "",
		},
	}
	assert.Equal(t, expectedTokens, tokens)
}

func TestScanReturnsPlusToken(t *testing.T) {
	tokens, err := Scan("./test_data/one_token.txt")
	assert.Nil(t, err)
	expectedTokens := []Token{
		{
			tokenType: PLUS,
			line:      1,
			lexeme:    "+",
		},
		{
			tokenType: EOF,
			line:      2,
			lexeme:    "",
		},
	}
	assert.Equal(t, expectedTokens, tokens)
}

func TestScanErrorsWhenFileDoesntExist(t *testing.T) {
	tokens, err := Scan("./test_data/missing_file.txt")
	assert.NotNilf(t, err, "Expected error when file doesn't exist")
	assert.Nil(t, tokens)
}

func TestScanReturnsMultiLineTokens(t *testing.T) {
	tokens, err := Scan("./test_data/multi_line_tokens.txt")
	assert.Nil(t, err)
	expectedTokens := []Token{
		{
			tokenType: PLUS,
			line:      1,
			lexeme:    "+",
		},
		{
			tokenType: MINUS,
			line:      2,
			lexeme:    "-",
		},
		{
			tokenType: STAR,
			line:      3,
			lexeme:    "*",
		},
		{
			tokenType: EOF,
			line:      4,
			lexeme:    "",
		},
	}
	assert.Equal(t, expectedTokens, tokens)
}
