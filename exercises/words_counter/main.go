package words_counter

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
)

func CountTopWords(filePath string, maxWords int) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open file: " + filePath)
	}
	defer file.Close()

	wordsCount := countWordsInFile(file)
	sortedWords := sortElements(wordsCount)
	topWords := pickElements(sortedWords[:maxWords], wordsCount)

	return topWords, nil
}

func countWordsInFile(file *os.File) map[string]int {
	wordsCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		updateWordsCount(line, &wordsCount)
	}

	return wordsCount
}

func sortElements(wordsCount map[string]int) []string {
	var sortedWords []string
	for word := range wordsCount {
		sortedWords = append(sortedWords, word)
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		return wordsCount[sortedWords[i]] > wordsCount[sortedWords[j]]
	})

	return sortedWords
}

func pickElements(keys []string, elements map[string]int) map[string]int {
	pickedElements := make(map[string]int)

	for _, key := range keys {
		pickedElements[key] = elements[key]
	}

	return pickedElements
}

func updateWordsCount(line string, wordsCount *map[string]int) {
	words := strings.Fields(line)
	for _, word := range words {
		if word == "" {
			continue
		}

		(*wordsCount)[word] = (*wordsCount)[word] + 1
	}
}
