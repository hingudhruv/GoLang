package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in GO!!")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	//loops in go
	// this is comma separated format you cna skip a value by using a _
	for key, value := range languages {
		fmt.Printf("For Key %v value is %v \n", key, value)
	}
}
