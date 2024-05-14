package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in GoLang")

	var fruitList [4]string
	fruitList[1] = "Banana"
	fruitList[3] = "Apple"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of the list is ", len(fruitList))

	var vegList = [3]string{"potato", "onion", "capsicum"}

	fmt.Println("vegetable 2 is ", vegList[1])

}
