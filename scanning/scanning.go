package scanning

import (
	"log"
	"os"
)

// Scanner scanner struct
type Scanner struct {
	source  string
	line    int
	start   int
	current int
	tokens  []string
}

// check whether we are at the end of a source line
func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// OpenFile open file function
func OpenFile(path string) error {
	file, openError := os.Open(path)
	if openError != nil {
		log.Println("Error opening file")
		return openError
	}

	closeError := file.Close()
	if closeError != nil {
		log.Println("Error closing file")
		return closeError
	}

	return nil
}
