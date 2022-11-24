package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	newPw := getNewPassword(bytes)
	fmt.Printf("Day 11 - part 1\nSanta's new password should be %s\n", newPw)
	temp := []byte(newPw)
	temp[7]++
	newerPw := getNewPassword(temp)
	fmt.Printf("Day 11 - part 2\nSanta's very new password is %s\n", newerPw)
}

func getNewPassword(bytes []byte) string {
	if isValidPassword(bytes) {
		return string(bytes)
	} else {
		next := addToPassword(bytes)
		return getNewPassword(next)
	}
}

func isValidPassword(bytes []byte) bool {
	containsNoIllegalRune := true
	containsAscendingRunes := false
	containsTwoPairs := false
	pairOne := byte('0')
	pairTwo := byte('0')
	for i, b := range bytes {
		if b == 'i' || b == 'o' || b == 'l' {
			containsNoIllegalRune = false
			break
		}
		if i+2 < len(bytes)-1 && b+1 == bytes[i+1] && b+2 == bytes[i+2] {
			containsAscendingRunes = true
		}
		if i+1 <= len(bytes)-1 && b == bytes[i+1] {
			if pairOne == '0' {
				pairOne = b
			}
			if pairTwo == '0' && b != pairOne {
				pairTwo = b
			}
		}
	}
	if pairOne != '0' && pairTwo != '0' && pairOne != pairTwo {
		containsTwoPairs = true
	}
	return containsAscendingRunes && containsNoIllegalRune && containsTwoPairs
}

func addToPassword(bytes []byte) []byte {
	bytes[7]++
	if bytes[7] > 'z' {
		bytes[7] = 'a'
		bytes[6]++
	}
	if bytes[6] > 'z' {
		bytes[6] = 'a'
		bytes[5]++
	}
	if bytes[5] > 'z' {
		bytes[5] = 'a'
		bytes[4]++
	}
	if bytes[4] > 'z' {
		bytes[4] = 'a'
		bytes[3]++
	}
	if bytes[3] > 'z' {
		bytes[3] = 'a'
		bytes[2]++
	}
	if bytes[2] > 'z' {
		bytes[2] = 'a'
		bytes[1]++
	}
	if bytes[1] > 'z' {
		bytes[1] = 'a'
		bytes[0]++
	}
	if bytes[0] > 'z' {
		panic("Size of bytes[0] exceeds limit")
	}
	return bytes
}
