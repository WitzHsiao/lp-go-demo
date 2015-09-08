package main

import (
	"fmt"
	"net/http"
)

type Message struct {
	Name    string
	Content string
}

var messages chan Message = make(chan Message)

func pollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msg := <-messages
	fmt.Println(msg)
	fmt.Fprintf(w, "%s: %s", msg.Name, msg.Content)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	messages <- Message{
		Name:    r.FormValue("Name"),
		Content: r.FormValue("Content"),
	}
	fmt.Fprintf(w, "send success")
}

func main() {
	http.HandleFunc("/lp", pollHandler)
	http.HandleFunc("/send", pushHandler)
	fmt.Println("Start server at port 8080")
	http.ListenAndServe(":8080", nil)
}
