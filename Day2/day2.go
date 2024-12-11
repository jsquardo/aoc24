package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(input string) ([][]int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %w", err)
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number %q: %w", part, err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return reports, nil
}

func IsValid(report []int) bool {
	if len(report) < 2 {
		// A single number or empty report is trivially valid
		return true
	}

	// Determine the direction: increasing, decreasing, or invalid
	direction := 0 // 1 for increasing, -1 for decreasing

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// Ensure the difference is within the valid range
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Determine the direction
		if diff > 0 {
			if direction == -1 {
				// Switching from decreasing to increasing
				return false
			}
			direction = 1
		} else if diff < 0 {
			if direction == 1 {
				// Switching from increasing to decreasing
				return false
			}
			direction = -1
		}
	}

	return true
}

func CountValidlines(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if IsValid(report) {
			count++
		}
	}
	return count
}

// Part 2
// This should check if a report is safe as-is or with one level removed
func IsValidWithDampener(report []int) bool {
	if IsValid(report) {
		return true
	}

	// Remove each level and check if the modified report is valid
	for i := 0; i < len(report); i++ {
		// Create new report with index level removed
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check modified report
		if IsValid(modifiedReport) {
			return true
		}
	}
	// If no valid modification was found, the report is unsafe
	return false
}

func CountValidLinesWithDampener(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if IsValidWithDampener(report) {
			count++
		}
	}
	return count
}
