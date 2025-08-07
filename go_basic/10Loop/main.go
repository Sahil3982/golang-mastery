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

	for _, value := range lost {
		fmt.Println(value)
	}
	result := greet(43, 432) - sub(42, 12+3)
	fmt.Println(result)
	Sum, Message := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(Message, Sum)
}

func greet(a int16, b int16) int16 {
	fmt.Println("Hi greet ", a+b)
	return a + b
}

func sub(a int16, b int16) int16 {
	return a - b
}

func proAdder(value ...int) (int, string) {
	total := 0

	for _, value := range value {
		total = total + value
	}

	return total, "Sum of :"
}
