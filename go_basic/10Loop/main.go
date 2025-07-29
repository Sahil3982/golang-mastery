package main

import "fmt"

func main() {
	fmt.Println("Hi loops")

	lost := []string{"sa", "sas", "fw", "wr", "rtot"}

	// for d :=0 ; d<len(lost); d++ {
	// 	fmt.Println(lost[d])
	// }

	// for i := range lost {
	// 	fmt.Println(i,lost[i])
	// }

	for key,_ := range lost{
		fmt.Println(key)
	}
}
