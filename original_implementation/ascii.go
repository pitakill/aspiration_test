package main

import (
	"fmt"
)

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// c handles the current position candidate for transformation
	// m the ith position to transform
	var c, m int8 = 1, 3
	t := make([]rune, len(s))

	for i, r := range s {
		if !isAlphanumeric(r) {
			t[i] = r
			continue
		}

		if isDigit(r) {
			t[i] = r
			c++
			continue
		}

		if c%m == 0 && isLowercase(r) || c%m != 0 && isUppercase(r) {
			t[i] = changeCase(r)
			c++
			continue
		}

		t[i] = r
		c++
	}

	return string(t)
}

// changeCase changes the rune r from case if is between the ASCII range of [a-z][A-Z]
// otherwise returns the same rune
func changeCase(r rune) rune {
	// There is 32 positions between a and A in ASCII range
	const STEP = 32

	if isUppercase(r) {
		return r + STEP
	}
	if isLowercase(r) {
		return r - STEP
	}

	return r
}

// isAlphanumeric returns a bool if the rune r is between the correct ASCII ranges
func isAlphanumeric(r rune) bool {
	if isLetter(r) || isDigit(r) {
		return true
	}

	return false
}

// isUppercase verifies if the rune is between the [A-Z] ASCII range
func isUppercase(r rune) bool {
	if 'A' <= r && r <= 'Z' {
		return true
	}

	return false
}

// isLowercase verifies if the rune is between the [a-z] ASCII range
func isLowercase(r rune) bool {
	if 'a' <= r && r <= 'z' {
		return true
	}

	return false
}

// isLetter verifies if the rune is between the [A-Z] or [a-z] ASCII ranges
func isLetter(r rune) bool {
	if isUppercase(r) || isLowercase(r) {
		return true
	}

	return false
}

// isDigit verifies if the rune is between the [0-9] ASCII range
func isDigit(r rune) bool {
	if '0' <= r && r <= '9' {
		return true
	}

	return false
}
