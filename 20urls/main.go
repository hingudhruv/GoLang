package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=bie23ce"

func main() {
	fmt.Println("Welcome to url handling using net module")

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println("Params is", val)
	}
	//do not forget to use & when defining URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "study",
		RawQuery: "username=dhruv",
	}
	anotheUrl := partsOfUrl.String()
	fmt.Println(anotheUrl)
}
