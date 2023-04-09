package main

import "fmt"

func CreateSet(words ...string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, word := range words {
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
