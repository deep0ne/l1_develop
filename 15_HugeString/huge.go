// 15. К каким негативным последствиям может привести фрагмент кода и как это исправить?
// приведите корректный пример реализации

package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

/*

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

*/

/*
Возможно, проблема этого кода в том, что мы создаём очень большую строку, которая остаётся в памяти.
С использованием v[:100] создаётся новая строка, но она будет ссылаться на память исходной строки, что означает,
что даже если нам не нужна исходная строка, она по-прежнему будет существовать и память не будет освобождена до тех пор,
пока все ссылки на неё больше не будут использоваться

Это можно исправить с помощью буфера: мы создаём новую строку на основе v, но затем сбрасываем содержимое буфера,
поэтому ссылки на исходную строку v больше не будет и строка будет удалена сборщиком мусора
*/
var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func main() {
	m := 1 << 10
	v := createHugeString(m)
	buffer := bytes.NewBufferString(v[:100])
	justString := buffer.String()
	buffer.Reset()
	fmt.Println(justString)
}
