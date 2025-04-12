package main

import (
	"strings"
)

func reverseWords(sentence string) string {
	reverseWords := ""
	words := strings.Split(sentence, " ")

	for key, word := range words {
		if key == 0 {
			reverseWords += reserveWord(word)
		} else {
			reverseWords += " " + reserveWord(word)
		}
	}

	return reverseWords
}

func reserveWord(word string) string {
	chars := strings.Split(word, "")
	reversedWord := ""

	for _, char := range chars {
		reversedWord = char + reversedWord
	}

	return reversedWord
}
