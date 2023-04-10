// 3. Дана последовательности чисел: 2 4 6 8 10. Найти сумму их квадратов с использованием конкурентных вычислений

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func ConcurrentSumMutex(nums []int) int {
	var (
		sum   int
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	wg.Add(len(nums)) // добавляем вейт группу, чтобы функция дождалась суммирования
	for _, num := range nums {
		go func(num int, mutex *sync.Mutex) {
			defer wg.Done()
			mutex.Lock() // начало критической секции
			sum += num * num
			mutex.Unlock() // конец критической секции
		}(num, &mutex)
	}
	wg.Wait()
	return sum
}

func ConcurrentSumChannel(nums []int) int {
	var (
		sum int
		wg  sync.WaitGroup
	)

	channel := make(chan int)

	for _, num := range nums {
		wg.Add(1)
		go func(num int) { // запускаем горутины и отправляем квадраты в канал
			defer wg.Done()
			channel <- num * num
		}(num)
	}

	go func() {
		wg.Wait()
		close(channel) // горутина, которая ждёт завершения других горутин и закрывает канал
	}()

	for num := range channel { // читаем из канала и складываем в переменную
		sum += num
	}

	return sum
}

func ConcurrentSumAtomic(nums []int) int {
	var (
		sum int64
		wg  sync.WaitGroup
	)
	wg.Add(1) // добавляем вейт группу, чтобы функция дождалась суммирования
	go func() {
		for _, num := range nums {
			atomic.AddInt64(&sum, int64(num*num)) // atomic int чтобы избежать гонку
		}
		wg.Done()
	}()
	wg.Wait()
	return int(sum)
}

// функция для проверки
func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num * num
	}
	return sum
}

func main() {
	m := rand.Perm(500)
	fmt.Println(ConcurrentSumMutex(m))
	fmt.Println(ConcurrentSumChannel(m))
	fmt.Println(ConcurrentSumAtomic(m))
	fmt.Println(sum(m))
}
