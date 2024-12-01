package util

import (
	"os"
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	fullContent := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed auctor ipsum vel lorem suscipit, in ornare lacus lobortis.`

	lineContent := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Sed auctor ipsum vel lorem suscipit, in ornare lacus lobortis.",
	}
	wordsContent := []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit.",
		"Sed", "auctor", "ipsum", "vel", "lorem", "suscipit,", "in", "ornare", "lacus", "lobortis.",
	}

	testFile := "test_file.txt"

	err := os.WriteFile(testFile, []byte(fullContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	full := ReadFileFull(testFile)
	lines := ReadFileLines(testFile)
	words := ReadFileWords(testFile)

	if fullContent != string(full) {
		t.Fatalf("Full content mismatch:\nExpected: %v\nGot: %v", fullContent, full)
	}

	if !reflect.DeepEqual(lineContent, lines) {
		t.Fatalf("Line content mismatch:\nExpected: %v\nGot: %v", lineContent, lines)
	}

	if !reflect.DeepEqual(wordsContent, words) {
		t.Fatalf("Word content mismatch:\nExpected: %v\nGot: %v", wordsContent, words)
	}
}
