// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func CreateSet(words ...string) map[string]struct{} {
	set := make(map[string]struct{}) // значение в виде пустой структуры, которая не занимает памяти
	for _, word := range words {
		if _, ok := set[word]; ok { // проверяем, есть ли значение в нашем множестве
			continue // если есть, идём дальше по циклу
		}
		set[word] = struct{}{}
	}
	return set
}

func main() {
	set := CreateSet("cat", "cat", "dog", "cat", "tree")
	for word := range set {
		fmt.Println(word)
	}
}
