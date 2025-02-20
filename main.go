package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-task", show-task)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello, User! Welcome to ToDoList App!"
	fmt.Println(writer, greeting)
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
