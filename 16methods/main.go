package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")

	dhruv := User{"Dhruv Hingu", "dmhingu@gmail.com", true, 21}
	fmt.Println(dhruv)
	fmt.Printf("Details of Dhruv are %+v\n", dhruv)
	fmt.Printf("Email is %v\n", dhruv.Email)
	dhruv.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// whenever we call a function the object copy is passed
func (u User) GetStatus() {
	fmt.Println("The Status of the User is", u.Status)
}

// so if we check in main function the values would not change
