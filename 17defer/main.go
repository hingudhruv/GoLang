package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")
	defer fmt.Println("World")
	defer fmt.Println("1.")
	defer fmt.Println("2.")
	defer fmt.Println("3...")
	fmt.Println("Hello")
}
