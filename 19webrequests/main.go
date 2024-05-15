package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Type")

	defer response.Body.Close() // This is user's responsibility

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
