package concurrent_word_lengths

import (
	"testing"
)

func TestGetAverageWordsSize(t *testing.T) {
	// ARRANGE
	sentences := []string{
		"Go is fun",
		"Concurrency is not parallelism",
		"Channels orchestrate goroutines",
	}

	// ACT
	result := GetAverageWordsSize(sentences)

	// ASSERT
	if result["Go is fun"] != (7.0 / 3) {
		t.Errorf("GetAverageWordsSize() = %v; expected %v", result["Go is fun"], (7.0 / 3))
	}

	if result["Concurrency is not parallelism"] != (27.0 / 4) {
		t.Errorf("GetAverageWordsSize() = %v; expected %v", result["Concurrency is not parallelism"], (27.0 / 4))
	}

	if result["Channels orchestrate goroutines"] != (29.0 / 3) {
		t.Errorf("GetAverageWordsSize() = %v; expected %v", result["Channels orchestrate goroutines"], (29.0 / 3))
	}
}
