package scanning

import "fmt"

// Scanner scanner struct
type Scanner struct {
	source  string
	line    int
	start   int
	current int
	tokens  []Token
}

// Token token struct
type Token struct {
	lexeme    string
	line      int
	tokenType string
}

func Scan() {
	file, _ := OpenFile("../scanning/example_input.txt")
	scanner := scanLine(file)
	fmt.Printf("scanner: %v\n", scanner)
}
