package main

import (
	"strings"
	"unicode"
)

func Dictionary(s string) map[string]int {
	runeS := []rune(s)
	wordMap := map[string]int{}
	L := 0
	for r := range runeS {
		char := runeS[r]
		if !unicode.IsLetter(char) {
			word := strings.ToLower(string(runeS[L:r]))
			if len(word) > 0 {
				wordMap[word] += 1
			}
			L = r + 1
		}

	}
	if L < len(runeS) {
		word := strings.ToLower(string(runeS[L:]))
		wordMap[word] += 1
	}
	return wordMap
}

func Palindrom(s string) bool {
	runeS := []rune(s)
	L := 0
	r := len(runeS) - 1

	for L < r {
		for L < r && !unicode.IsLetter(runeS[L]) && !unicode.IsNumber(runeS[L]) {
			L++
		}

		for L < r && !unicode.IsLetter(runeS[r]) && !unicode.IsNumber(runeS[r]) {
			r--
		}

		if L < r {
			char1 := strings.ToLower(string(runeS[L])) 
			char2 := strings.ToLower(string(runeS[r]))

			if char1 != char2 {
				return false
			}
	
			L++
			r--
		} 
	}

	return true
}
