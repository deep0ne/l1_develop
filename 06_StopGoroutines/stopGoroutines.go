package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 1. Через закрытие канала
func StopGoroutineWithSignal() {
	done := make(chan struct{}) // канал завершения
	go func() {
		for {
			select {
			case <-done: // если канал закрыт, завершаем
				fmt.Println("STOPPING GOROUTINE...")
				return
			default:
				fmt.Printf("Generating some number: %d\n", rand.Intn(10))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(done)                        // закрываем канал
	time.Sleep(100 * time.Millisecond) // чтобы вывелось сообщение
}

// 2. Через канал с булевым значением
func StopGoroutineWithBool() {
	done := make(chan bool) // создаём канал
	go func() {
		for {
			select {
			case <-done: // ждём, когда мы сможем из него что-то считать
				fmt.Println("STOPPING GOROUTINE...")
				return
			default:
				fmt.Printf("Generating some number: %d\n", rand.Intn(10))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	done <- true // отправляем значение, чтобы остановить горутину
	time.Sleep(100 * time.Millisecond)
}

func StopGoroutineWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("STOPPING GOROUTINE...")
				return
			default:
				fmt.Printf("Generating some number: %d\n", rand.Intn(10))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func main() {
	StopGoroutineWithSignal()
	StopGoroutineWithBool()
	StopGoroutineWithContext()
}
