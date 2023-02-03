package scanning

// Scanner scanner struct
type Scanner struct {
	source  string
	line    int
	start   int
	current int
	tokens  []Token
}

func Scan(path string) ([]Token, error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}
	scanner := scanLine(file)
	return scanner.tokens, nil
}
