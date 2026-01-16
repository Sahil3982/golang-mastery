package internals

import (
	"fmt"
	"sync"
)

func workers(id int ,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workers(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
