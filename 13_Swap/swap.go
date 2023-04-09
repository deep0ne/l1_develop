// 13. Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := 5, 10
	a, b = b, a // обычный свап

	fmt.Println(a == 10 && b == 5)
}

// func Swap(a, b *int) {
// 	*a, *b = *b, *a
// }
