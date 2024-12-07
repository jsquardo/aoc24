package main

import (
	"fmt"
)

func main() {
	combined, err := ParseFile("../inputs/day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Combined Lines:")
	for _, line := range combined {
		fmt.Println(line)
	}
}
