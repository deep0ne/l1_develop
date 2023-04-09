package main

import "fmt"

func Intersection(firstSet, secondSet map[any]struct{}) []any {
	intersection := make([]any, 0)
	for key := range firstSet {
		if _, ok := secondSet[key]; ok {
			intersection = append(intersection, key)
		}
	}
	return intersection
}

func main() {
	firstSet := map[any]struct{}{
		"Hello":  {},
		"World!": {},
		"One":    {},
		"More":   {},
		"Time":   {},
	}
	secondSet := map[any]struct{}{
		"Hello":  {},
		"World!": {},
		"Second": {},
		"Time":   {},
	}

	firstSetInts := map[any]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	secondSetInts := map[any]struct{}{
		2: {},
	}

	fmt.Println(Intersection(firstSet, secondSet))         // [Time Hello World!]
	fmt.Println(Intersection(firstSetInts, secondSetInts)) // [2]
}
