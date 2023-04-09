// 25. Реализовать собственную функцию sleep.

package main

import (
	"context"
	"fmt"
	"time"
)

// 1. через цикл
func Sleep(seconds int) {
	start := time.Now()
	for time.Since(start) < time.Duration(seconds)*time.Second {
		continue
	}
}

// 2. Через контекст с таймаутом
func SleepWithCtx(seconds int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	if <-ctx.Done(); true {
		return
	}
}

func main() {
	fmt.Println("Starting...")
	Sleep(5)
	fmt.Println("5 seconds passed...")

	fmt.Println("Starting again...")
	SleepWithCtx(5)
	fmt.Println("Another 5 seconds passed...")
}
