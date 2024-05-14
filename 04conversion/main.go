package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to GO !!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(nil)
	} else {
		fmt.Println("Added 1 to your rating ", numRating+1)
	}

}
