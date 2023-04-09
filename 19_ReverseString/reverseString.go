package main

import (
	"strings"
	"unicode/utf8"
)

func ReverseString(word string) string {
	if len(word) == 0 {
		return ""
	}

	var sb strings.Builder // создаём билдер для новой строки
	sb.Grow(len(word))     // наращиваем количество байт, чтобы избежать лишние переаллокации
	for len(word) > 0 {
		char, runeSize := utf8.DecodeLastRuneInString(word) // получаем доступ к последней руне
		sb.WriteRune(char)                                  // пишем последнюю руну
		word = word[:len(word)-runeSize]                    // уменьшаем строку
	}
	return sb.String()
}
