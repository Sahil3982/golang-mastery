package internals

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func Goroutine() {
	go sayHello()
	fmt.Println("Hello from main")
	time.Sleep(time.Second)
}
