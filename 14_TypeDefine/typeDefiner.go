// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func DefineType(value any) {
	// TypeOf returns the reflection Type that represents the dynamic type of i.
	fmt.Printf("The type of value is %s\n", reflect.TypeOf(value))
}

func main() {

	c := make(chan int)
	someElements := []any{"hello", 42, c, true}
	for _, element := range someElements {
		DefineType(element)
	}

}
