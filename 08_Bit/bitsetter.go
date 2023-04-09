package main

import "fmt"

func BitSet(num int64, pos uint) int64 {
	mask := int64(1) << pos // создаём маску при помощи побитового сдвига влево на pos позиций
	return num | mask       // применяем побитовый OR между маской и исходным числом, чтобы установить i-ый бит в 1
}

func BitClear(num int64, pos uint) int64 {
	mask := int64(1) << pos
	return num &^ mask // применяем AND с инвертированной маской ^
}

func main() {
	n := int64(32)
	pos := uint(4)
	n = BitSet(n, pos)
	fmt.Println(n)
	n = BitClear(n, pos)
	fmt.Println(n)
}
