package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hy Slices")

	var animals = []string{"Dog", "Cat", "Elephant"}
	fmt.Println(animals)
	fmt.Println("Length of animals:", len(animals))
	animals = append(animals, "Lion", "Tiger")
	var newAnimal = animals[:3]
	fmt.Println("After appending:", newAnimal)
	fmt.Println("New animals:", animals)
	fmt.Println("New length of animals:", len(newAnimal))
	fmt.Println("New length of animals:", len(animals))


	highScores := make([]int,5)

	highScores[0] = 100
	highScores[1] = 200		
	highScores[2] = 300
	highScores[3] = 400
	highScores[4] = 500
	highScores = append(highScores, 7845,545,56456,6,456,54,6,54,6,54,6,45,6,54,6,54)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)


	fmt.Println(highScores)


	var courses = [] string{"React","Node","Next","Python"}

	var index int = 2
	
	courses = append(courses[:index], courses[index+1:]... )


	fmt.Println(courses)


}
