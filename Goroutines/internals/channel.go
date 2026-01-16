package internals

import "fmt"

func Channel() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()

	value := <-ch
	fmt.Println(value)

}