package main

import (
	"errors"
)

func getSquare(n int) string {
	if n < 1 {
		return ""
	}

	if n == 1 {
		line, _ := getFilledLine(1)
		return line
	}

	if n == 2 {
		line, _ := getFilledLine(2)
		return line + "\n" + line
	}

	var square string
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			line, _ := getFilledLine(n)
			square += line + "\n"
		} else {
			line, _ := getHalfFilledLine(n)
			square += line + "\n"
		}
	}

	return square
}

func getFilledLine(n int) (string, error) {
	if n < 1 {
		return "", errors.New("n must be greater than 0")
	}

	var line string
	for i := 0; i < n; i++ {
		line += "#"
	}
	return line, nil
}

func getHalfFilledLine(n int) (string, error) {
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
