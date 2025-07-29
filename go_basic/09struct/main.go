package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hi Struct")
	// Sahil := User{"Sahil", 33, "XYZ", true}
	// fmt.Println(Sahil)

	// number := 1
	// if number > 4 {
	// 	fmt.Println("Less")
	// } else {
	// 	fmt.Println("More")
	// }

	// if num := 6; num == 0 {
	// 	fmt.Println("Number is not zero")
	// } else {
	// 	fmt.Println("Number is  zero")

	// }

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(8) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice 1")
		fallthrough
	case 2:
		fmt.Println("Dice 2")
		fallthrough
	case 3:
		fmt.Println("Dice 3")
	case 4:
		fmt.Println("Dice 4")
	case 5:
		fmt.Println("Dice 5")
	case 6:
		fmt.Println("Dice 6")
	default:
		fmt.Println("None")
	}
}

type User struct {
	Name    string
	Age     int
	Address string
	Status  bool
}
