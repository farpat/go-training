package unique_lines_extractor

import "testing"

func TestExtractUniqueLines(t *testing.T) {
	// ARRANGE
	filePath := "sample.txt"

	// ACT
	result, err := ExtractUniqueLines(filePath)

	// ASSERT
	if err != nil {
		t.Errorf("Failed to open file: %s", filePath)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 unique lines, got %d", len(result))
	}

	if result[0] != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", result[0])
	}

	if result[1] != "go is great" {
		t.Errorf("Expected 'go is great', got '%s'", result[1])
	}

	if result[2] != "learning go" {
		t.Errorf("Expected 'learning go', got '%s'", result[2])
	}
}
