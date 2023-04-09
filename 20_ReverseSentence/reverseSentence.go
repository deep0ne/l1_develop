package main

import "strings"

func ReverseSentence(sentence string) string {
	var sb strings.Builder
	sb.Grow(len(sentence))

	words := strings.Split(sentence, " ")
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		if i != 0 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}
