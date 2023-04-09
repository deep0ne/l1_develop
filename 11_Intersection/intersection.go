// 11. Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func Intersection(firstSet, secondSet map[any]struct{}) []any {
	intersection := make([]any, 0) // слайс с найденными общими элементами
	for key := range firstSet {    // проходимся по ключам первого неупорядоченного множества
		if _, ok := secondSet[key]; ok { // если по этому ключу что-то есть во втором множестве
			intersection = append(intersection, key) // кладем в слайс с пересечениями
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
