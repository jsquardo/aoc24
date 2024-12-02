package main

import "fmt"

func main() {
	left, right, err := ParseInput("inputs/day1/input1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print results
	fmt.Println("Left List:", left)
	fmt.Println("Right List:", right)
}
