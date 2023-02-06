package scanning

// Scanner scanner struct
type Scanner struct {
	source  string
	line    int
	start   int
	current int
	tokens  []Token
}

func Scanner(path string) ([]Token, error, error) {
	file, openError := OpenFile(path)
	if openError != nil {
		return nil, openError, nil
	}
	scanner := scanLines(file)
	closeError := closeFile(file)
	if closeError != nil {
		return scanner.tokens, nil, closeError
	}
	return scanner.tokens, nil, nil
}
