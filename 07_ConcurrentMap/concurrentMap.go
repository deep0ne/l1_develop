// 7.Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type ConcurrentMap struct {
	lock sync.RWMutex // используем RWMutex, коорый разделяет писателей и читателей
	info map[string]int
}

func (c *ConcurrentMap) Put(str string, val int) {
	c.lock.Lock() // лочим мьютекс перед записью, чтобы избежать состояние гонки
	c.info[str] = val
	c.lock.Unlock()
}

func (c *ConcurrentMap) Get(str string) int {
	c.lock.RLock() // запираем мьютекс для чтения
	v := c.info[str]
	c.lock.RUnlock()
	return v
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		info: make(map[string]int),
	}
}

func main() {
	var wg sync.WaitGroup
	concurrentMap := NewConcurrentMap()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			num := strconv.Itoa(rand.Intn(5))
			concurrentMap.Put(num, i)
		}
	}()
	wg.Wait()
	fmt.Println(concurrentMap.info)
}
