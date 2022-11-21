package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 4 - Part 1")
	hash, index := getHash(file, getFiveZeros)
	fmt.Printf("%x\t%d\n", hash, index)
	fmt.Println("Day 4 - Part 2")
	hash, index = getHash(file, getSixZeros)
	fmt.Printf("%x\t%d\n", hash, index)

	// hash, index = getHash(file, getSevenZeros)
	// fmt.Printf("%x\t%d\n", hash, index)
}

// Since go had this already implemented i just needed to use it.
func getHash(input []byte, condition func(hash string) bool) (hash [16]byte, index int) {
	for i := 0; ; i++ {
		current := fmt.Sprintf(string(input)+"%d", i)
		value := md5.Sum([]byte(current))
		hashString := fmt.Sprintf("%x", value)
		if condition(hashString) {
			index = i
			hash = value
			return
		}
	}
}

func getFiveZeros(hashString string) bool {
	return hashString[0] == '0' &&
		hashString[1] == '0' &&
		hashString[2] == '0' &&
		hashString[3] == '0' &&
		hashString[4] == '0'
}

func getSixZeros(hashString string) bool {
	return hashString[0] == '0' &&
		hashString[1] == '0' &&
		hashString[2] == '0' &&
		hashString[3] == '0' &&
		hashString[4] == '0' &&
		hashString[5] == '0'
}

func getSevenZeros(hashString string) bool {
	return hashString[0] == '0' &&
		hashString[1] == '0' &&
		hashString[2] == '0' &&
		hashString[3] == '0' &&
		hashString[4] == '0' &&
		hashString[5] == '0' &&
		hashString[6] == '0'
}
