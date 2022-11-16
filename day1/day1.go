package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	result := 0
	firstBasement := -1
	for i, r := range file {
		result += getValue(r)
		if firstBasement == -1 && result == -1 {
			firstBasement = i + 1
		}
	}
	fmt.Println("Day 1 - Part 1")
	fmt.Printf("Santa will end on floor %d\n", result)
	fmt.Println("Day 1 - Part 2")
	fmt.Printf("Santa enters the basement at position %d\n", firstBasement)
}

func getValue(r byte) int {
	if r == ')' {
		return -1
	} else {
		return 1
	}
}
