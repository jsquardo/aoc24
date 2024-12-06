package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parse file and read Line by Line
func ParseFile(input string) ([]int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %w", err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into individual number strings
		parts := strings.Fields(line)

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number %q: %w", part, err)
			}
			numbers = append(numbers, num)
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("frror reading file %w", err)
	}

	return numbers, nil
}
