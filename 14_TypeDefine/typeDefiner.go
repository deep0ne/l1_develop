package main

import (
	"fmt"
	"reflect"
)

func DefineType(value any) {
	fmt.Println(reflect.TypeOf(value))
}

func main() {
	var (
		i int
		s string
		b bool
		c chan int
	)

	DefineType(i)
	DefineType(s)
	DefineType(b)
	DefineType(c)

}
