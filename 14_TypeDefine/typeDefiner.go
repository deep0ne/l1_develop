// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func DefineType(value any) {
	// TypeOf returns the reflection Type that represents the dynamic type of i.
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
