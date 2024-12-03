package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* 1. Parse the input into two lists, `left` and `right`.
2. Sort both lists from lowest to highest.
3. Initialize a variable `totalDistance = 0`.
4. For each pair of elements (left[i], right[i]):
    a. Calculate the distance as distance(left[i] - right[i]).
    b. Add this distance to `totalDistance`.
5. Output `totalDistance`. */

// 1. Parse the input into two lists, `left` and `right`.

// Parse the input string into two seperate lists of integers
func ParseInput(input string) ([]int, []int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open file %w", err)
	}
	defer file.Close()

	// Create a left int and a right int
	var left []int
	var right []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into two parts (left and right)
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("Invalid number of lines: %s", line)
		}

		// Convert each part into an int
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left number: %w", err)
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse right number: %w", err)
		}

		// Add to correct lists Ex: Left and Right
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	// Handle any errors during scan
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return left, right, nil
}

// 2. Sort both lists from lowest to highest.
func Sort(left []int, right []int) {
	sort.Ints(left)
	sort.Ints(right)
}

func FindDif(left, right []int) int {
	var totalDif int
	for i := range left {
		dif := left[i] - right[i]
		if dif < 0 {
			dif = -dif
		}
		totalDif += dif
	}
	return totalDif
}
