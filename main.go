package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Go to Jim"
var fullgolang = "Maintain a healthy diet"
var reward = "Get a great physique"
var taskItems = []string{shortGolang, fullgolang, reward}

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/hello-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for index, task := range taskItems {
		fmt.Fprintf(writer, "%d. %s\n", index+1, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var gretting = "###Welcome to our To-Do-List app###"
	fmt.Fprintln(writer, gretting)
}
