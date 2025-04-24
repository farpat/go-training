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
	uniqueSeen := make(map[string]bool)
	var uniqueLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, isExisting := uniqueSeen[line]
		if line == "" || isExisting {
			continue
		}

		uniqueSeen[line] = true
		uniqueLines = append(uniqueLines, line)
	}

	return uniqueLines
}
