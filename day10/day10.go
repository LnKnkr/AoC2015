package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sequence := lookAndSay(bytes, 40)
	fmt.Printf("Day 10 - Part 1\nThe sequence is %d long", len(sequence))
}

func lookAndSay(bytes []byte, runs int) string {
	sequence := string(bytes)
	newSequence := ""
	for i := 0; i < runs; i++ {
		temp := ""
		for j, r := range sequence {
			if j == 0 {
				temp += string(r)
			} else if r == rune(temp[0]) {
				temp += string(r)
			} else if r != rune(temp[0]) {
				newSequence += fmt.Sprintf("%d%s", len(temp), string(temp[0]))
				temp = string(r)
			}
		}
		newSequence += fmt.Sprintf("%d%s", len(temp), string(temp[0]))
		sequence = newSequence
		newSequence = ""
	}
	return sequence
}
