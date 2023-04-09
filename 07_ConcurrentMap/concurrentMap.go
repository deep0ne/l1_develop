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

func (c *ConcurrentMap) Put(str string) {
	c.lock.Lock() // лочим мьютекс перед записью, чтобы избежать состояние гонки
	c.info[str]++
	c.lock.Unlock()
}

func (c *ConcurrentMap) Get(str string) int {
	c.lock.RLock()
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
		for i := 0; i < 1000; i++ {
			num := strconv.Itoa(rand.Intn(2))
			concurrentMap.Put(num)
		}
	}()
	wg.Wait()
	fmt.Println(concurrentMap.info)
}
