// 2. Написать программу, которая конкурентно рассчитывает значение квадратов чисел
// взятых из массива (2, 4, 6, 8, 10) и выводит их вадраты в stdout

package main

import (
	"fmt"
	"sync"
)

func ConcurrentSquaresWithWG(nums []int) {
	wg := sync.WaitGroup{}
	// т.к. запускаем горутину для каждого числа, создаём вейтгруппу с счётчиком 5
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}
	wg.Wait() // ожидаем, пока горутины завершат свою работу
}

func ConcurrentSquaresWithChannel(nums []int) {
	// создаём буферизированный канал, куда будем отправлять квадраты
	ch := make(chan int, len(nums))
	for _, num := range nums {
		// запускаем свою горутину для каждого num
		go func(num int) {
			ch <- num * num
		}(num)
	}

	//вычитываем данные из канала
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ConcurrentSquaresWithChannel(nums)
	ConcurrentSquaresWithWG(nums)
}
