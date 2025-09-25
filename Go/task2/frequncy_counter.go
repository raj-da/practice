package main

import (
	"fmt"
	"unicode"
)

// frequncyCount takes a string and return frequncy count of non punctual characters
func frequncyCount(text string)  {
	charFrequncy := map[string]int {}

	for _, byte := range text {
		char := fmt.Sprintf("%c", byte)
		if !unicode.IsPunct([]rune(char)[0]) && char != " " {
			charFrequncy[char]++
		}
	}

	for _, byte := range text {
		char := fmt.Sprintf("%c", byte)
		if !unicode.IsPunct([]rune(char)[0]) && char != " "{
			println(char, ":", charFrequncy[char])
		}
	}
}