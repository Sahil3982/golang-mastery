package main

import "fmt"

var isLogout bool = false 

const IsAdmin string = "admin"  /// if first later of cont variavke is capitalized, it is exported and can be accessed outside the package

func main(){
	var username string = "JohnDoe"
	fmt.Println(username);
	fmt.Printf("Type of username: %T \n", username);

	var age int = 30
	fmt.Println(age);
	fmt.Printf("Type of age: %T \n", age);

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn ,isLogout);
	fmt.Printf("Type of isLoggedIn: %T \n", isLoggedIn)
	fmt.Printf("Type of isLoggedIn: %T \n", isLogout)

	var check bool	
	fmt.Println(check);
	fmt.Printf("Type of check: %T \n", check);

	name := 84768
	fmt.Println(name);
	fmt.Printf("Type of name: %T \n", name);

	fmt.Println(IsAdmin);
	fmt.Printf("Type of IsAdmin: %T \n", IsAdmin);

}
