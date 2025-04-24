package unique_lines_extractor

import (
	"path"
	"runtime"
	"testing"
)

func TestExtractUniqueLines(t *testing.T) {
	// ARRANGE
	_, filename, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(filename), "sample.txt")

	// ACT
	result, err := ExtractUniqueLines(filePath)

	// ASSERT
	if err != nil {
		t.Errorf("Failed to open file: %s", filePath)
		return
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 unique lines, got %d", len(result))
		return
	}

	if result[0] != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", result[0])
		return
	}

	if result[1] != "go is great" {
		t.Errorf("Expected 'go is great', got '%s'", result[1])
		return
	}

	if result[2] != "learning go" {
		t.Errorf("Expected 'learning go', got '%s'", result[2])
		return
	}
}
