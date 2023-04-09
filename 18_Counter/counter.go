// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	Count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.Count, 1) // используем атомик, чтобы избежать состояние гонки
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup // создаём вейтгруппу, чтобы дождаться суммирования до конца
	for i := 0; i < 100; i++ {
		wg.Add(1) // для каждой горутины инкрементируем счётчик вейт группы
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count) // 100 * 100 = 10000
}
