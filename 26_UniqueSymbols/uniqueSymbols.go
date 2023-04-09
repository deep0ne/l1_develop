// 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import "strings"

func UniqueSymbols(word string) bool {
	loweredWord := strings.ToLower(word) // приводим строку к нижнему регистру
	symbols := make(map[rune]struct{})   // делаем мапу с рунами
	for _, symbol := range loweredWord { // итерируемся по рунам строки
		if _, ok := symbols[symbol]; ok { // проверяем наличие руны в мапе
			return false // если оно там есть, значит встретили повторяющийся символ
		}
		symbols[symbol] = struct{}{} // если его там нет, добавляем в мапу для будущих проверок
	}
	return true // если не вернуло false, значит цикл кончился => повторяющиеся символы не встретились
}
