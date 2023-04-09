package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendRead(seconds int) {
	// создаём канал для отправки чисел
	stream := make(chan int)

	// первая горутина для отправки рандомного числа
	go func() {
		for {
			randomNum := rand.Intn(100)
			stream <- randomNum
			fmt.Printf("Send %d\n", randomNum)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// вторая горутина для считывания из канала
	go func() {
		for value := range stream {
			fmt.Printf("Receiving %d\n", value)
		}
	}()

	// по истечению n секунд закрываем канал
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("EXITING PROGRAMM...")
	close(stream)

}

func main() {
	SendRead(5)
}
