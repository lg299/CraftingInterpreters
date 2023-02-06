package scanning

import "testing"

func scanLineReturnsEmptyScanner(t *testing.T){
    file, err := OpenFile("test_data/empty_file.txt")
    scanner := scanLines(file)


}