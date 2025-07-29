package main

import "fmt"

func main() {
	fmt.Println("Going to learn pointers in Go!")

	myNumber := 42
	
	var myPointer *int
	myPointer = &myNumber // myPointer now holds the address of myNumber

	fmt.Println("Value of myNumber:", &myNumber)
	fmt.Println("Value pointed to by myPointer:", *myPointer)

	*myPointer = *myPointer + 1 // Increment the value at the address myPointer points to
	fmt.Println("New value of myNumber after increment:", myNumber)
}