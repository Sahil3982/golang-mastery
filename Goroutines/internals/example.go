package internals

import "fmt"

func worker(id int, ch chan int) {
	result := id * 2
	ch <- result
}

func Example() {
	ch := make(chan int)

	for i := 1; i <= 5; i++ {
		go worker(i, ch)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}