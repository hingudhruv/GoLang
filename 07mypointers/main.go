package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers.")

	// var ptr *int;

	myNumber := 23
	fmt.Println(myNumber)

	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value in the pointer address", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Updated value in the pointer address", *ptr)

}
