package scanning

import (
	"log"
	"os"
)

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
