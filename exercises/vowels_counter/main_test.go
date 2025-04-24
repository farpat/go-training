package vowel_counter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetVowelsCount(t *testing.T) {
	// ARRANGE
	input := []string{"Hello", "World", "A big apple", "Z", "", "aeiouyAEIOUY"}

	// ACT
	result := GetVowelsCount(input)

	// ASSERT
	expectedResult := map[string]int{"Hello": 2, "World": 1, "A big apple": 4, "Z": 0, "": 0, "aeiouyAEIOUY": 12}
	if diff := cmp.Diff(result, expectedResult); diff != "" {
		t.Errorf("GetVowelsCount() = %v diff %v", result, diff)
	}
}
