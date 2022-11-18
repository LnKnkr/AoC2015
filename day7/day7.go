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

	fileString := string(file)
	result := createCircuit(fileString, createMapPart1, 0, "a")
	fmt.Printf("Day 7 - part 1\nA will be at '%d' in the end\n", result[0])
	result = createCircuit(fileString, createMapPart2, result[0], "a")
	fmt.Printf("Day 7 - part 2\nA will be at '%d' in the end\n", result[0])
}

func createCircuit(input string, mapper func(string, uint16) map[string]string, override uint16, wantedResults ...string) (results []uint16) {
	ops := mapper(input, override)

	results = make([]uint16, 0, len(wantedResults))
	for _, s := range wantedResults {
		r := getResult(s, ops)
		results = append(results, r)
		fmt.Printf("%d\n", r)
	}
	return results
}

func createMapPart1(input string, notused uint16) map[string]string {
	lines := strings.Split(input, "\r\n")
	operations := make(map[string]string, len(lines))

	for _, l := range lines {
		splitLine := strings.Split(l, "->")
		operations[strings.TrimSpace(splitLine[1])] = strings.TrimSpace(splitLine[0])
	}
	return operations
}

func createMapPart2(input string, newB uint16) map[string]string {
	lines := strings.Split(input, "\r\n")
	operations := make(map[string]string, len(lines))

	for _, l := range lines {
		splitLine := strings.Split(l, "->")
		operations[strings.TrimSpace(splitLine[1])] = strings.TrimSpace(splitLine[0])
	}
	operations["b"] = fmt.Sprintf("%d", newB)
	return operations
}

func getResult(input string, op map[string]string) uint16 {
	operation := op[input]
	if operation == "" {
		result, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			panic(err)
		}
		return uint16(result)
	}

	tasks := strings.Split(operation, " ")
	if len(tasks) == 1 {
		return getResult(strings.TrimSpace(tasks[0]), op)
	}
	if len(tasks) == 2 {
		op[input] = fmt.Sprintf("%d", ^getResult(strings.TrimSpace(tasks[1]), op))
		return ^getResult(strings.TrimSpace(tasks[1]), op)
	}
	if len(tasks) == 3 {
		left := getResult(strings.TrimSpace(tasks[0]), op)
		right := getResult(strings.TrimSpace(tasks[2]), op)

		switch tasks[1] {
		case "AND":
			op[input] = fmt.Sprintf("%d", left&right)
			return left & right
		case "OR":
			op[input] = fmt.Sprintf("%d", left|right)
			return left | right
		case "LSHIFT":
			op[input] = fmt.Sprintf("%d", left<<right)
			return left << right
		case "RSHIFT":
			op[input] = fmt.Sprintf("%d", left>>right)
			return left >> right
		}
	}
	panic("should not be reached")
}
