package square_drawing

import "testing"

func TestGetSquare(t *testing.T) {
	tests := [4]struct {
		name     string
		input    uint
		expected string
	}{
		{name: "Square of size 0x0", input: 0, expected: ""},
		{name: "Square of size 1x1", input: 1, expected: "*\n"},
		{name: "Square of size 2x2", input: 2, expected: "**\n**\n"},
		{name: "Square of size 3x3", input: 3, expected: "***\n* *\n***\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ACT
			result := GetSquare(test.input)

			// ASSERT
			if result != test.expected {
				t.Errorf("GetSquare(%d) = %s; expected %s", test.input, result, test.expected)
			}
		})
	}
}
