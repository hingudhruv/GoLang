package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to ")
	DecodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"ReactJs Bootcamp", 299, "Udemy.in", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "Udemy.in", "adc123", []string{"fullstack", "js"}},
		{"Angular", 299, "Udemy.in", "a23v23", nil},
	}

	finalJson, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"course": "ReactJs Bootcamp",
		"Price": 299,
        "website": "Udemy.in",
        "tags": ["web-dev","js"]
	}
	`)

	var lcoCourses course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Jason was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("Not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourses)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v Values is %v and Type is %T\n", k, v, v)
	}
}
