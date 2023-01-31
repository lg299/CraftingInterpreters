package scanning

import (
	"log"
	"os"
)

// OpenFile open file function
func OpenFile(path string) (*os.File, error) {
	file, openError := os.Open(path)
	if openError != nil {
		log.Println("Error opening file")
		return file, openError
	}
	return file, nil
}

func closeFile(file *os.File) error {
	closeError := file.Close()
	if closeError != nil {
		log.Println("Error closing file")
		return closeError
	}
	return nil
}
