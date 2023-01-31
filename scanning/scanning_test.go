package scanning

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScanTokens(t *testing.T) {
	scanner := Scanner{source: "var x = 5;", line: 1, start: 0, current: 6}
	tokens := scanner.ScanTokens(scanner.source)
	fmt.Printf("tokens: %v\n", reflect.TypeOf(tokens))
	if len(tokens) < 1 {
		t.Fatalf("No tokens added")
	}
}

func TestIsAtEnd(t *testing.T) {
	scanner := Scanner{source: "var x = 5;", line: 1, start: 0, current: 6}
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
