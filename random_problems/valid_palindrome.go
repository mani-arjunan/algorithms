package main

import "strings"

func isAlphaNumeric(char rune) bool {
	if char <= 'a' && char >= 'z' {
		return true
	}
	if char <= 'A' && char >= 'Z' {
		return true
	}
	if char <= '0' && char >= '9' {
		return true
	}

	return false
}

func isPalindrome(str string) bool {
	leftPtr := 0
	rightPtr := len(str)

	for leftPtr < len(str) && rightPtr > 0 {
		leftChar := str[leftPtr]
		rightChar := str[rightPtr]

		if !isAlphaNumeric(rune(leftChar)) {
			leftPtr++
			continue
		}

		if isAlphaNumeric(rune(leftChar)) {
			if !isAlphaNumeric(rune(rightChar)) {
				rightPtr--
				continue
			}

			if strings.ToLower(string(leftChar)) != strings.ToLower(string(rightChar)) {
				return false
			}

			leftPtr++
			rightPtr--
		}
	}

	return true
}
