package scanning

import (
	"testing"
)

func TestOpenFileWithPath(t *testing.T) {
	expected := OpenFile("example_input.txt")
	if expected != nil {
		t.Fatalf("Error opening file")
	}
}

func TestOpenFileNoPath(t *testing.T) {
	expectedFail := OpenFile("")
	if expectedFail.Error() != "open : no such file or directory" {
		t.Fatalf("Error should be thrown - empty path.")
	}
}
