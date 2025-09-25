package main

import "strings"

func isPalandrom(text string) bool {
	lPtr := 0
	rPtr := len(text) - 1 

	text = strings.ToLower(text)
	for lPtr <= rPtr {
		if text[lPtr] != text[rPtr] {
			return false
		} else {
			lPtr++
			rPtr--
		}
	}
	return true
}
