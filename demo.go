package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url = "https://lco.dev"
//const myurl string = "https://jsonplaceholder.typicode.com"

func main() {
	//fmt.Println("Hello world")
	//response, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(response)
	//defer response.Body.Close()
	//databytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(databytes))
	//fmt.Println(myurl)
	//result, _ := url.Parse(myurl)
	//fmt.Println(result.Scheme)
	//qparams := result.Query()
	//fmt.Println(qparams)
	//for _, value := range qparams {
	//	fmt.Println(value)
	//}
	//partsOfurl := &url.URL{
	//	Scheme: "https",
	//	Host: ""
	//}
	//performGetRequest()

}
func performGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response length:", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func performPostRequest() {
	const url = "https://jsonplaceholder.typicode.com/posts"

}
