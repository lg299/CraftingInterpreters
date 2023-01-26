package scanning

import (
	"reflect"
	"testing"
)

func TestIsAtEnd(t *testing.T) {
	scanner := Scanner{source: "var x = 5;", line: 0, start: 0, current: 6}
	isAtEnd := scanner.isAtEnd()
	if reflect.TypeOf(isAtEnd).Kind() != reflect.Bool {
		t.Fatalf("Incorrect type returned.")
	}
}

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
