// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Conveyor(nums []int, firstChan, secondChan chan int, wg *sync.WaitGroup) {

	wg.Add(2) // для двух горутин ниже

	go func() {
		defer wg.Done()
		for _, num := range nums {
			fmt.Printf("First channel: Receiving %d\n", num)
			firstChan <- num // отправляем в канал из слайса
			time.Sleep(200 * time.Millisecond)
		}
		close(firstChan) // закрываем канал, чтобы подать сигнал, что данные больше не будут поступать
	}()

	go func() {
		defer wg.Done()
		for num := range firstChan {
			square := num * num
			fmt.Printf("Second channel: Receiving square %d\n", square)
			secondChan <- square // пишем квадрат в второй канал
			time.Sleep(200 * time.Millisecond)
		}
		close(secondChan) // закрываем канал, чтобы подать сигнал, что данные больше не будут поступать
	}()

	for sq := range secondChan { // читаем из канала
		fmt.Printf("Outputting squares...%d\n", sq)
	}

	wg.Wait()
}

func main() {
	nums := rand.Perm(20)
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	Conveyor(nums, ch1, ch2, &wg)
}
