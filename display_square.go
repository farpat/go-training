package main

import (
	"errors"
)

func displaySquare(n int) string {
	if n < 1 {
		return ""
	}

	if n == 1 {
		line, _ := displayFilledLine(1)
		return line
	}

	if n == 2 {
		line, _ := displayFilledLine(2)
		return line + "\n" + line
	}

	var square string
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			line, _ := displayFilledLine(n)
			square += line + "\n"
		} else {
			line, _ := displaySemiFilledLine(n)
			square += line + "\n"
		}
	}

	return square
}

func displayFilledLine(n int) (string, error) {
	if n < 1 {
		return "", errors.New("n must be greater than 0")
	}

	var line string
	for i := 0; i < n; i++ {
		line += "#"
	}
	return line, nil
}

func displaySemiFilledLine(n int) (string, error) {
	if n < 1 {
		return "", errors.New("n must be greater than 0")
	}

	var line string
	for i := 0; i < n; i++ {
		if i == 0 || i == (n-1) {
			line += "#"
		} else {
			line += " "
		}
	}
	return line, nil
}
