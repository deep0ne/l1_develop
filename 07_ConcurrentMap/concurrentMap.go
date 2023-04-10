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
	data map[string]int
}

func (m *ConcurrentMap) Put(str string, val int) {
	m.lock.Lock() // лочим мьютекс перед записью, чтобы избежать состояние гонки
	m.data[str] = val
	m.lock.Unlock()
}

func (m *ConcurrentMap) Get(str string) (int, bool) {
	m.lock.RLock() // запираем мьютекс для чтения
	v, ok := m.data[str]
	m.lock.RUnlock()
	return v, ok
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]int),
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
	fmt.Println(concurrentMap.data)
}
