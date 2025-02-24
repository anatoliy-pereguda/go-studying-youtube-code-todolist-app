# GO student-led project ToDoList App  

### HTML markup for better display  
```
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
```
