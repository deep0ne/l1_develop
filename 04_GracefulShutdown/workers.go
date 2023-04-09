// 4. Релизовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читаю произвольные данные из канала и выводят в stdout
// Необходима возможность выбора количество воркеров при старте.
// Программа должна завершаться по нажатию CTRL+C.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// graceful shutdown
// https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/
// При получении сигнала завершаются работы воркеров.
// Такой способ позволяет реализовать "чистое" завершение - корректное очищение ресурсов и выход

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, done chan struct{}) {
	defer wg.Done()

	for {
		select {
		case j, ok := <-jobs: // если канал не закрыт, читаем и выводим
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, j)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		// ждём сигнала из канала, чтобы завершить цикл
		case <-done:
			return
		}
	}
}

func StartWorkers(numWorkers int) {
	jobs := make(chan int) // создаём канал с джобами

	var wg sync.WaitGroup       // создаём вейт группу
	done := make(chan struct{}) // канал завершения

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs, done) // запускаем воркеров
	}

	interrupt := make(chan os.Signal, 1) // канал для ожидания сигнала завершения
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt // блокируем выполнение
		close(done) // как только читаем сигнал из канала, закрываем канал завершения
	}()

	go func() {
		for {
			select {
			case <-done: // если закрыт, закрываем канал с джобами
				close(jobs)
				return
			default:
				jobs <- rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("All workers done")
}

func main() {
	StartWorkers(5)
}
