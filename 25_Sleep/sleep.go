package main

import (
	"context"
	"fmt"
	"time"
)

func Sleep(seconds int) {
	start := time.Now()
	for time.Since(start) < time.Duration(seconds)*time.Second {
		continue
	}
}

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
