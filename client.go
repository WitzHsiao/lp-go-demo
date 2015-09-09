package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var dataChan chan string = make(chan string)

func requestData() {
	resp, err := http.Get("http://localhost:8080/lp")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	dataChan <- string(body)
}

func printData() {
	go func() {
		fmt.Println(<-dataChan)
	}()
	requestData()
}

func main() {
	fmt.Println("Start long polling...")
	for {
		printData()
	}
}
