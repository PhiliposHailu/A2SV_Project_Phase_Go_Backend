package main

import (
	"strings"
	"unicode"
)

func dictionary(s string) map[string]int {
	myRune := []rune(s)
	wordMap := map[string]int{}
	L := 0
	for r := range myRune {
		char := myRune[r]
		if !unicode.IsLetter(char) {
			word := strings.ToLower(string(myRune[L:r]))
			if len(word) > 0 {
				wordMap[word] += 1
			}
			L = r + 1
		}

	}
	if L < len(myRune) {
		word := strings.ToLower(string(myRune[L:]))
		wordMap[word] += 1
	}
	return wordMap
}

