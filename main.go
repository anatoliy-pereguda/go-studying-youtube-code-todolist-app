package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Tasks:</h1><ul>")
	for _, task := range taskItems {
		fmt.Fprintf(writer, "<li>%s</li>", task)
	}
	fmt.Fprintf(writer, "</ul>")
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our ToDoList App!"
	fmt.Fprintf(writer, "<h1>%s</h1>", greeting)
}
