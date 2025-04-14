package square_drawing

import "strings"

func GetSquare(size uint) string {
	var result string

	if size < 3 {
		for i := uint(0); i < size; i++ {
			result += displayFullLine(size) + "\n"
		}

		return result
	}

	for i := uint(0); i < size; i++ {
		if i == 0 || i == size-1 {
			result += displayFullLine(size) + "\n"
		} else {
			result += displayHalfLine(size) + "\n"
		}
	}

	return result
}

func displayHalfLine(size uint) string {
	return "*" + strings.Repeat(" ", int(size)-2) + "*"
}

func displayFullLine(size uint) string {
	return strings.Repeat("*", int(size))
}
