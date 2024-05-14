package main

import "fmt"

// Constant define with Public access [ by keeping first letter CAPITAL]
const LoginToken string = "somevalue"

func main() {
	// String Type
	var username string = "dhruv"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//Boolean Type
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//Integer Type
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	// To Know more visit: https://go.dev/ref/spec#Numeric_types

	//default values and alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	//implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style (using walarus [:=] operator)
	numberOfUsers := 300
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)
	// this can only be used inside the method/function.

	// Accessing the global value
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
