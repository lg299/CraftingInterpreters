package scanning

type TokenTypesStruct struct {
	// Single-character tokens.
	LeftParen  string
	RightParen string
	LeftBrace  string
	RightBrace string
	COMMA      string
	DOT        string
	MINUS      string
	PLUS       string
	SEMICOLON  string
	SLASH      string
	STAR       string

	// One or two character tokens.
	BANG         string
	BangEqual    string
	EQUAL        string
	EqualEqual   string
	GREATER      string
	GreaterEqual string
	LESS         string
	LessEqual    string

	// Literals.
	IDENTIFIER string
	STRING     string
	NUMBER     string

	// Keywords.
	AND    string
	CLASS  string
	ELSE   string
	FALSE  string
	FUN    string
	FOR    string
	IF     string
	NIL    string
	OR     string
	PRINT  string
	RETURN string
	SUPER  string
	THIS   string
	TRUE   string
	VAR    string
	WHILE  string

	EOF string
}

func TokenTypes() map[string]string {
	tokenTypes := make(map[string]string)
	tokenTypes["("] = "LEFT_PAREN"
	tokenTypes[")"] = "RIGHT_PAREN"
	tokenTypes["["] = "LEFT_BRACE"
	tokenTypes["]"] = "RIGHT_BRACE"
	tokenTypes[","] = "COMMA"
	tokenTypes["."] = "DOT"
	tokenTypes["="] = "EQUAL"
	tokenTypes["-"] = "MINUS"
	tokenTypes["+"] = "PLUS"
	tokenTypes[";"] = "SEMICOLON"
	tokenTypes["/"] = "SLASH"
	tokenTypes["*"] = "STAR"
	tokenTypes["while"] = "WHILE"
	return tokenTypes
}
