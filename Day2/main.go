package main

import (
	"fmt"
)

func main() {
	lines, err := ParseFile("../inputs/day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	validCount := CountValidLinesWithDampener(lines)
	fmt.Println("Safe Reports(with dampener):", validCount)
}
