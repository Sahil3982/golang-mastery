package main

import "fmt"

func main() {
	fmt.Println("Hellw Map")

	language := make(map[string]string)

	language["js"] = "javascript"
	language["C#"] = "CSharp"
	language["JAva"] = "Java"
	language["RB"] = "Ruby"

	fmt.Println(language)
	delete(language,"C")

	fmt.Println(language)

	for _, value := range language{
		fmt.Println(value)
	}


}
