package concurrent_word_lengths

import (
	"strings"
)

type AverageWordsSize struct {
	Sentence string
	Average  float64
}

func GetAverageWordsSize(sentences []string) map[string]float64 {
	averageWordSizeChannel := make(chan AverageWordsSize)
	averages := make(map[string]float64)

	for _, sentence := range sentences {
		go extractWordsFromSentence(sentence, averageWordSizeChannel)
	}

	for range sentences {
		average := <-averageWordSizeChannel
		averages[average.Sentence] = average.Average
	}

	return averages
}

func extractWordsFromSentence(sentence string, averageChannel chan AverageWordsSize) {
	words := strings.Fields(sentence)
	totalWordsLength := 0.0
	totalWords := 0.0

	for _, word := range words {
		totalWordsLength += float64(len(word))
		totalWords++
	}

	averageChannel <- AverageWordsSize{
		Sentence: sentence,
		Average:  totalWordsLength / totalWords,
	}
}
