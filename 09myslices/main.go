package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in GO !!")

	// When you don't define array size it is slice
	var fruitList = []string{"Apple", "Peach", "Banana"}
	fmt.Printf("Type of Fruit list is %T", fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 435
	highScores[3] = 132
	// highScores[4] = 323 [Throws error index out of range]

	fmt.Println(highScores)
	// append operation can be done because it reallocates the memory
	highScores = append(highScores, 4553, 1213)

	fmt.Println(highScores)

	//few more features of slice
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	// how to remove values from slices based on index

	var courses = []string{"reactjs", "Javascript", "Swift", "Python", "Ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	
}
