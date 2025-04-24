package words_counter

import (
	"testing"
)

func TestCountTopWords(t *testing.T) {
	// ARRANGE
	filePath := "sample.txt"
	maxWords := 2

	// ACT
	result, error := CountTopWords(filePath, maxWords)

	// ASSERT
	if error != nil {
		t.Errorf("Failed to open file: %s", filePath)
	}

	if len(result) != maxWords {
		t.Errorf("Expected %d words, got %d", maxWords, len(result))
	}

	if result["this"] != 2 {
		t.Errorf("Expected 2, got %d", result["this"])
	}

	if result["test"] != 6 {
		t.Errorf("Expected 6, got %d", result["test"])
	}
}
