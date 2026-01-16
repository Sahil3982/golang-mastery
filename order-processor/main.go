package main

import (
	"fmt"
	"sync"

	"github.com/sahil3982/order-process/model"
	"github.com/sahil3982/order-process/worker"
)

func main() {
	totalOrder := 20
	workercount := 5
	jobs := make(chan model.Order, 20)
	results := make(chan string, 20)

	var wg sync.WaitGroup

	for i := 1; i <= workercount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker.Process(id, jobs, results)
		}(i)
	}

	go func() {
		for i := 1; i <= totalOrder; i++ {
			jobs <- model.Order{
				ID:     i,
				Amount: float64(100 + i),
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
