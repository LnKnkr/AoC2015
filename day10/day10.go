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
	fmt.Printf("Day 10 - Part 1\nThe sequence is %d long\n", len(sequence))
	sequence = lookAndSay(bytes, 50)
	fmt.Printf("Day 10 - Part 2\nThe sequence is %d long\n", len(sequence))
}

func lookAndSay(bytes []byte, runs int) string {
	if runs == 0 {
		return string(bytes)
	}
	foo := []byte{}
	repeatCount := 1
	prev := byte(0)
	for j, b := range bytes {
		if j == 0 {
			prev = b
		} else if b == prev {
			prev = b
			repeatCount++
		} else if b != prev {
			foo = append(foo, byte(repeatCount+48), byte(prev))
			repeatCount = 1
			prev = b
		}
	}
	foo = append(foo, byte(repeatCount+48), byte(prev))

	return lookAndSay(foo, runs-1)
}
