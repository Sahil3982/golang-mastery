package main

import "fmt"

func main() {
	fmt.Println("Hy Array")

	var friuts [9] string 
	friuts[0] = "Apple"
	friuts[1] = "Banana"
	friuts[2] = "Mango"

	fmt.Println(friuts)

	var vegitables = [3] string {"Carrot", "Potato", "Tomato"}
	fmt.Println(vegitables)
	fmt.Println("Lenght of", len(vegitables))
	fmt.Println("Capacity of", cap(vegitables))
	fmt.Println("Capacity of friuts", cap(friuts))
	fmt.Println("First element of friuts", len(friuts))

}
