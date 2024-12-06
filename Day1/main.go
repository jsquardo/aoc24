package main

import "fmt"

func main() {
	// Part One
	left, right, err := ParseInput("inputs/day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	Sort(left, right)

	dif := FindDif(left, right)

	// Print results
	fmt.Println("Total Dif:", dif)

	// Part Two
	PartTwo(left, right)
}
