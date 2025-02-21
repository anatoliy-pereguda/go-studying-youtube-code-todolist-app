# GO YouTube Code TodoList App  
Go free YouTube Code from [TechWorld with Nana]( https://www.youtube.com/watch?v=XCZWyN9ZbEQ)  

## First sketch  
```
package main

import "fmt"

func main() {
	var shortGolang = "Watch GO Crash Course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a Cheesecake"

	fmt.Println("##### Welcome to ToDoList App! #####")

	fmt.Println("### List of my ToDos ###")
	fmt.Println(shortGolang)
	fmt.Println(fullGolang)
	fmt.Println(rewardDessert)

	fmt.Println()
	fmt.Println("### Tutorials ###")
	fmt.Println(shortGolang)
	fmt.Println(fullGolang)

	fmt.Println()
	fmt.Println("### Rewards ###")
	fmt.Println(rewardDessert)

	fmt.Println()
	fmt.Println("My Projects")
	fmt.Println(fullGolang)
}
```

### Second sketch  

```
package main

import "fmt"

func main() {
	fmt.Println("##### Welcome to ToDoList App! #####")

	var shortGolang = "Watch GO Crash Course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a Cheesecake"
	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	printTasks(taskItems)
	fmt.Println()
	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")
	fmt.Println("Updated list")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("### List of my ToDos ###")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
```

### HTTP sketch  

```
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
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}
```

### HTTP sketch с комментариями  

```
// это простое веб-приложение на Go для управления списком задач
// запуск: go run main.go
// в браузере
// http://localhost:8080/ → приветственное сообщение
// Hello user. Welcome to our Todolist App!
// http://localhost:8080/show-tasks → список задач
// Watch Go crash course
// Watch Nana's Golang Full Course
// Reward myself with a donut

// объявляем главный пакет программы, в Go каждая программа должна иметь пакет main, который является точкой входа в приложение
package main

// подключаем внешние пакеты
import (
	"fmt"      // стандартный пакет для форматированного ввода-вывода (например, печать текста)
	"net/http" // стандартный пакет для работы с HTTP-запросами (например, создание веб-серверов)
)

// определяем задачи в виде строковых переменных
var shortGolang = "Watch Go crash course"          
var fullGolang = "Watch Nana's Golang Full Course" 
var rewardDessert = "Reward myself with a donut"   

// создаем список задач в виде среза/слайса/динамического массива строк
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

// запускаем при старте главную функцию, это точка входа программы
func main() {
	// регистрируем обработчик для корневого пути "/" когда пользователь заходит на адрес http://localhost:8080/, вызывается функция helloUser
	http.HandleFunc("/", helloUser)

	// регистрируем обработчик для пути "/show-tasks" когда пользователь заходит на адрес http://localhost:8080/show-tasks, вызывается функция showTasks
	http.HandleFunc("/show-tasks", showTasks)

	// запускаем HTTP-сервер на порту 8080 параметр nil означает, что используется маршрутизатор по умолчанию
	http.ListenAndServe(":8080", nil)
}

// функция-обработчик для "/show-tasks"
// выводит список задач пользователю
// writer http.ResponseWriter — объект, через который отправляем ответ клиенту (здесь браузеру).
// request *http.Request — объект, содержащий информацию о входящем HTTP-запросе (здесь не используем)
func showTasks(writer http.ResponseWriter, request *http.Request) {
	// перебираем все задачи и отправляем их в HTTP-ответ
	for _, task := range taskItems { 
		fmt.Fprintln(writer, task) // выводит каждую задачу в новой строке
	}
}

// функция-обработчик для корневого маршрута "/"
// выводит приветственное сообщение пользователю
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!" // приветственное сообщение
	fmt.Fprintln(writer, greeting) // отправляем его в HTTP-ответ
}
```

### План по улучшению приложения  

- Добавить HTML-разметку для красивого отображения списка задач.
- Обработать ошибки (например, если сервер не может запуститься).
- Сделать задачи динамическими (например, добавить возможность редактирования или удаления задач).





