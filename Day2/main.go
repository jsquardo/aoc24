package main

import "fmt"

func main() {
	file := "inputs/day2.txt"

	// Call DayTwoPartOne func
	numbers, err := DayTwoPartOne(file)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Print parsed numbers
	fmt.Println("Numbers:", numbers)
}
