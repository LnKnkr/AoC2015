package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestIfOperationsAreDoneCorrectly(t *testing.T) {
	board := "123 -> x\r\n456 -> y\r\nx AND y -> d\r\nx OR y -> e\r\nx LSHIFT 2 -> f\r\ny RSHIFT 2 -> g\r\nNOT x -> h\r\nNOT y -> i"
	results := createCircuit(board, createMapPart1, 0, "d", "e", "f", "g", "h", "i", "x", "y")
	check.Eq(t, results, []uint16{72, 507, 492, 114, 65412, 65079, 123, 456})
}
