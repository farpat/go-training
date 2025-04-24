package unique_lines_extractor

import (
	"bufio"
	"os"
)

func ExtractUniqueLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return extractUniqueLinesOfFile(file), nil
}

func extractUniqueLinesOfFile(file *os.File) []string {
	uniqueRows := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		uniqueRows[line] = true
	}

	var uniqueLines []string
	for line := range uniqueRows {
		uniqueLines = append(uniqueLines, line)
	}

	return uniqueLines
}
