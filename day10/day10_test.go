package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestSingleOneBecomesTwoOnes(t *testing.T) {
	a := []byte("1")
	have := lookAndSay(a, 1)
	check.Eq(t, have, "11")
}

func TestMultipleDigitsGetReducedToTwoDigits(t *testing.T) {
	a := []byte("11")
	have := lookAndSay(a, 1)
	check.Eq(t, have, "21")
	have = lookAndSay([]byte("111"), 1)
	check.Eq(t, have, "31")
}

func TestOneGetsExtendedByMultipleRuns(t *testing.T) {
	a := []byte("1")
	have := lookAndSay(a, 4)
	check.Eq(t, have, "111221")
	have = lookAndSay([]byte(have), 1)
	check.Eq(t, have, "312211")
}

func BenchmarkLookAndSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lookAndSay([]byte("1"), 40)
	}
}
