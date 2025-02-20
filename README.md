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

### 3  


