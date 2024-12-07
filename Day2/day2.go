package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFile(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %w", err)
	}
	defer file.Close()

	var combinedLines []string
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace and join without spaces
		combined := strings.Join(strings.Fields(line), "")
		combinedLines = append(combinedLines, combined)

		// Check for errors
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("frror reading file %w", err)
		}
	}
	return combinedLines, nil
}
