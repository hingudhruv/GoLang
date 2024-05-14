package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	dhruv := User{"Dhruv Hingu", "dmhingu@gmail.com", true, 21}
	fmt.Println(dhruv)
	fmt.Printf("Details of Dhruv are %+v\n", dhruv)
	fmt.Printf("Email is %v", dhruv.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
