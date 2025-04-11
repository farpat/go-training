package main

import "fmt"

func displaySquare(n int) string {
	if n < 1 {
		return ""
	}

	if n == 1 {
		return displayFilledLine(1)
	}

	if n == 2 {
		return displayFilledLine(2) + "\n" + displayFilledLine(2)
	}

	var square string
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			square += displayFilledLine(n) + "\n"
		} else {
			square += displaySemiFilledLine(n) + "\n"
		}
	}

	return square
}

func displayFilledLine(n int) string {
	var line string
	for i := 0; i < n; i++ {
		line += "#"
	}
	return line
}

func displaySemiFilledLine(n int) string {
	var line string
	for i := 0; i < n; i++ {
		if i == 0 || i == (n-1) {
			line += "#"
		} else {
			line += " "
		}
	}
	return line
}

func main() {
	fmt.Println(displaySquare(0))
	fmt.Println("")
	fmt.Println(displaySquare(1))
	fmt.Println("")
	fmt.Println(displaySquare(2))
	fmt.Println("")
	fmt.Println(displaySquare(3))
	fmt.Println("")
	fmt.Println(displaySquare(4))
}
