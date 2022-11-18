package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	amountMemory := 0
	for _, r := range file {
		// fmt.Printf("%v\n", r)
		if r != '\n' && r != '\r' {
			amountMemory++
		}
	}

	amountLiteraler := 0
	amountString := 0
	fileString := string(file)
	fileSplit := strings.Split(fileString, "\r\n")
	for _, l := range fileSplit {
		unliteral, err := strconv.Unquote(l)
		if err != nil {
			panic(err)
		}
		amountString += len(unliteral)
		amountLiteraler += len(strconv.Quote(l))
	}

	fmt.Printf(
		"Day 8 - part 1\n"+
			"The amount of characters as string literals is %d\n"+
			"but without the escape character the amount is %d\n"+
			"So we have a difference of %d charcaters",
		amountMemory,
		amountString,
		amountMemory-amountString,
	)
	fmt.Printf(
		"Day 8 - part 2\n"+
			"The original amount of characters as string literals is %d\n"+
			"but when we get more literal our amount is %d\n"+
			"So we have %d more charcaters",
		amountMemory,
		amountLiteraler,
		amountLiteraler-amountMemory,
	)
}
