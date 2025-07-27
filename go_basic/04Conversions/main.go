package main

import (
	"fmt"
	"time"
)

func main() {

	// readerBhai := bufio.NewReader(os.Stdin)

	// fmt.Println("Enter your age:")
	// ageInput, _ := readerBhai.ReadString('\n')

	// fmt.Println("Enter your number:")
	// numberInput, _ := readerBhai.ReadString('\n')

	// fmt.Println("Your age is:", ageInput)
	// fmt.Println("Your number is:", numberInput)

	// var1, error := strconv.ParseFloat(strings.TrimSpace(ageInput), 64)

	// if error != nil {
	// 	fmt.Println("Error converting age input:", error)
	// } else {
	// 	fmt.Println("Converted age to float:", var1)
	// }
	// var2, error := strconv.ParseFloat(strings.TrimSpace(numberInput), 64)

	// if error != nil {
	// 	fmt.Println("Error converting number input:", error)
	// } else {
	// 	fmt.Println("Converted number to float:", var2)
	// }
	// sum := var1 + var2

	// fmt.Println("Sum:", sum)

	dates := time.Now()

	fmt.Println("Current date and time:", dates)
	fmt.Println("Current date and time:", dates.Format("2006-01-02 Monday"))

}
