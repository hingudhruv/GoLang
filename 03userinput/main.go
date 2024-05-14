package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GO !!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating: ")

	// comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)
	fmt.Printf("The type of the rating %T ", input)

}
