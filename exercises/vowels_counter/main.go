package vowel_counter

import "unicode"

type VowelCount struct {
	Sentence string
	Count    int
}

func GetVowelsCount(sentences []string) map[string]int {
	vowelCount := make(chan VowelCount)
	for _, sentence := range sentences {
		go countVowelsInSentence(sentence, vowelCount)
	}

	vowels := make(map[string]int)
	for range sentences {
		vowel := <-vowelCount
		vowels[vowel.Sentence] = vowel.Count
	}

	return vowels
}

func countVowelsInSentence(sentence string, vowelCount chan VowelCount) {
	count := 0

	for _, char := range sentence {
		if isVowel(char) {
			count++
		}
	}

	vowelCount <- VowelCount{
		Sentence: sentence,
		Count:    count,
	}
}

func isVowel(char rune) bool {
	switch unicode.ToLower(char) {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	default:
		return false
	}
}
