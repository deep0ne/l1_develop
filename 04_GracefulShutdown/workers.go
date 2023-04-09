// 4. Релизовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читаю произвольные данные из канала и выводят в stdout
// Необходима возможность выбора количество воркеров при старте.
// Программа должна завершаться по нажатию CTRL+C.

package main

import (
	"context"
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

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case j, ok := <-jobs: // если канал не закрыт, читаем и выводим
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, j)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		case <-ctx.Done():
			return
		}
	}
}

func StartWorkers(numWorkers int) {
	jobs := make(chan int) // создаём канал с джобами

	var wg sync.WaitGroup                                   // создаём вейт группу
	ctx, cancel := context.WithCancel(context.Background()) // в качестве альтернативы каналу завершения - контекст

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs, ctx) // запускаем воркеров
	}

	interrupt := make(chan os.Signal, 1) // канал для ожидания сигнала завершения
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt // блокируем выполнение
		cancel()    // как только читаем сигнал из канала, вызываем функцию отмены
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
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
