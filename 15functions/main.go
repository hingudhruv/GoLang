package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")

	result := adder(3, 4)
	fmt.Println("The Result is :", result)

	proResult, proMsg := proAdder(3, 2, 4, 1, 5, 2, 1)
	fmt.Println("The ProMsg is :", proMsg)
	fmt.Println("The ProResult is :", proResult)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Your Result is here"
}
