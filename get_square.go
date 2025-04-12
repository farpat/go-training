package main

func getSquare(n uint) string {
	if n == 0 {
		return ""
	}

	if n == 1 {
		return getFilledLine(uint(n))
	}

	if n == 2 {
		return getFilledLine(uint(n)) + "\n" + getFilledLine(uint(n))
	}

	var square string
	for i := uint(0); i < n; i++ {
		if i == 0 || i == (n-1) {
			line := getFilledLine(uint(n))
			square += line + "\n"
		} else {
			line := getHalfFilledLine(uint(n))
			square += line + "\n"
		}
	}

	return square
}

func getFilledLine(n uint) string {
	if n < 1 {
		return ""
	}

	var line string
	for i := uint(0); i < n; i++ {
		line += "#"
	}
	return line
}

func getHalfFilledLine(n uint) string {
	if n < 1 {
		return ""
	}

	var line string
	for i := uint(0); i < n; i++ {
		if i == 0 || i == (n-1) {
			line += "#"
		} else {
			line += " "
		}
	}
	return line
}
