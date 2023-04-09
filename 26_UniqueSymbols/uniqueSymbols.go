package main

import "strings"

func UniqueSymbols(word string) bool {
	loweredWord := strings.ToLower(word)
	symbols := make(map[rune]struct{})
	for _, symbol := range loweredWord {
		if _, ok := symbols[symbol]; ok {
			return false
		}
		symbols[symbol] = struct{}{}
	}
	return true
}
