package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb")

	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content length :", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	//content is in byte format

	fmt.Println("Byte count is", byteCount)
	fmt.Println("Data is ", responseString.String())

	// The string conversion can also be done in this way,
	// but the former gives you more control
	// fmt.Println(string(content))
	fmt.Println("Bye")
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name":"dhruv",
			"age":10
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("name", "Dhruv")
	data.Add("age", "21")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
