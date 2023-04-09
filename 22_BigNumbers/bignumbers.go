package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func BigIntOperation(a, b *big.Int, operation string) *big.Int {
	switch operation {
	case "+":
		sum := new(big.Int).Add(a, b)
		return sum
	case "-":
		subtraction := new(big.Int).Sub(a, b)
		return subtraction
	case "*":
		product := new(big.Int).Mul(a, b)
		return product
	case "/":
		division := new(big.Int).Div(a, b)
		return division
	default:
		return new(big.Int)
	}
}

func BigFloatOperation(a, b *big.Float, operation string) *big.Float {
	switch operation {
	case "+":
		sum := new(big.Float).Add(a, b)
		return sum
	case "-":
		subtraction := new(big.Float).Sub(a, b)
		return subtraction
	case "*":
		product := new(big.Float).Mul(a, b)
		return product
	case "/":
		division := new(big.Float).Quo(a, b)
		return division
	default:
		return new(big.Float)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // сканим "а"
	aString := scanner.Text()

	scanner.Scan() // сканим "b"
	bString := scanner.Text()

	_, err := strconv.Atoi(aString)
	_, err2 := strconv.Atoi(bString)
	floatFlag := err != nil || err2 != nil // если хотя бы одно значение - флоат, работаем с флоатами

	var operation string // сканим операцию
	fmt.Scan(&operation)
	if floatFlag {
		a, success := new(big.Float).SetString(aString)
		if !success {
			log.Fatal("Did not convert")
		}

		b, success := new(big.Float).SetString(bString)
		if !success {
			log.Fatal("Did not convert")
		}

		fmt.Printf("%v %s %v = %s", a, operation, b, BigFloatOperation(a, b, operation).Text('f', 2))

	} else {
		a, success := new(big.Int).SetString(aString, 0)
		if !success {
			log.Fatal()
		}

		b, success := new(big.Int).SetString(bString, 0)
		if !success {
			log.Fatal()
		}

		fmt.Printf("%v %s %v = %s", a, operation, b, BigIntOperation(a, b, operation))
	}
}
