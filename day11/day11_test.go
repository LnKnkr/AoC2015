package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestNewPasswordOfGivenInputIsCorrect(t *testing.T) {
	want := "abcdffaa"
	got := "abcdefgh"
	check.Eq(t, getNewPassword([]byte(got)), want)
}

func TestNewPasswordOfGivenInputIsCorrect2(t *testing.T) {
	want := "ghjaabcc"
	got := "ghijklmn"
	check.Eq(t, getNewPassword([]byte(got)), want)
}

func TestIsValidReturnsTrueOnValidInput(t *testing.T) {
	got := isValidPassword([]byte("ghjaabcc"))
	check.Eq(t, got, true)

	got = isValidPassword([]byte("ghjaabcc"))
	check.Eq(t, got, true)
}
