// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import "strings"

func ReverseSentence(sentence string) string {
	var sb strings.Builder // создаём билдер и наращиваем капасити для избежание аллокаций
	sb.Grow(len(sentence))

	words := strings.Split(sentence, " ")  // сплитим предложение
	for i := len(words) - 1; i >= 0; i-- { // идём с конца и пишем в билдер строки
		sb.WriteString(words[i])
		if i != 0 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}
