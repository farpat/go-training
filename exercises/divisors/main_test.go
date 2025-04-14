package divisors

import "testing"

func TestGetDivisors(t *testing.T) {
	// ARRANGE
	input := 12

	// ACT
	result := GetDivisors(input)

	// ASSERT
	expected := []int{1, 2, 3, 4, 6, 12}

	if len(result) != len(expected) {
		t.Errorf("GetDivisors(12) = %v; expected %v", result, expected)
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("GetDivisors(12) = %v; expected %v", result, expected)
		}
	}
}
