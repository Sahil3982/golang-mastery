package worker

import (
	"fmt"
	"time"
	"github.com/sahil3982/order-process/model"
)

func Process(workerID int, jobs <-chan model.Order, results chan<- string) {

	for order := range jobs {
		time.Sleep(200 * time.Microsecond)
		results <- fmt.Sprintf(
			"Worker %d processed Order %d (₹%.2f)",
			workerID,
			order.ID,
			order.Amount,
		)
	}

}
