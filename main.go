package main

import "fmt"

func main() {
	left, right, err := ParseInput("inputs/day1/input1.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	Sort(left, right)

	dif := FindDif(left, right)

	// Print results
	fmt.Println("Total Dif:", dif)
}
