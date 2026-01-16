package internals

import (
	"fmt"
	"time"
)

func processOrder(orderID int, ch chan string) {
	var orderIDStr = fmt.Sprintf("%d", orderID)
	time.Sleep(time.Second) // Simulate processing time
	ch <- "Order " + orderIDStr + " processed"
}
func ECommerce() {
	ch := make(chan string)

	for i := 1; i <= 300; i++ {
		go processOrder(i, ch)
	}
	for i := 1; i <= 300; i++ {
		println(<-ch)
	}
}
