package scanning

type TokenType string

const (
	// Single-character tokens.
	LeftParen  TokenType = "LEFT_PAREN"
	RightParen           = "RIGHT_PAREN"
	LeftBrace            = "LEFT_BRACE"
	RightBrace           = "RIGHT_BRACE"
	COMMA                = "COMMA"
	DOT                  = "DOT"
	MINUS                = "MINUS"
	PLUS                 = "PLUS"
	SEMICOLON            = "SEMICOLON"
	SLASH                = "SLASH"
	STAR                 = "STAR"

	// One or two character tokens.
	BANG         = "BANG"
	BangEqual    = "BANG_EQUAL"
	EQUAL        = "EQUAL"
	EqualEqual   = "EQUAL_EQUAL"
	GREATER      = "GREATER"
	GreaterEqual = "GREATER_EQUAL"
	LESS         = "LESS"
	LessEqual    = "LESS_EQUAL"

	// Literals.
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords.
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NIL    = "NIL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"

	EOF = "EOF"
)

func TokenTypes() map[string]TokenType {
	tokenTypes := make(map[string]TokenType)
	tokenTypes["("] = LeftParen
	tokenTypes[")"] = RightParen
	tokenTypes["["] = LeftBrace
	tokenTypes["]"] = RightBrace
	tokenTypes[","] = COMMA
	tokenTypes["."] = DOT
	tokenTypes["="] = EQUAL
	tokenTypes["-"] = MINUS
	tokenTypes["+"] = PLUS
	tokenTypes[";"] = SEMICOLON
	tokenTypes["/"] = SLASH
	tokenTypes["*"] = STAR
	tokenTypes["while"] = WHILE
	return tokenTypes
}
