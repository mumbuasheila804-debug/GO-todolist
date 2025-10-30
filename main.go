package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

var tasks []Task

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("🗒️  No tasks yet.")
		return
	}
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Name)
	}
	fmt.Println()
}

func addTask(name string) {
	tasks = append(tasks, Task{Name: name, Completed: false})
	fmt.Println("✨ Task added successfully!")
}

func completeTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("⚠️  Invalid task number.")
		return
	}
	tasks[index].Completed = true
	fmt.Println("🎉 Task marked as complete!")
}

func deleteTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("⚠️  Invalid task number.")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("🗑️  Task deleted!")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🧠 Welcome to Go Task Manager!")
	fmt.Println("Type one of the following commands:")
	fmt.Println("add <task name> | complete <task number> | delete <task number> | list | exit")

	for {
		fmt.Print("\n> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("❗ Please provide a task name.")
				continue
			}
			addTask(parts[1])

		case "list":
			listTasks()

		case "complete":
			if len(parts) < 2 {
				fmt.Println("❗ Please provide task number.")
				continue
			}
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("❗ Invalid number.")
				continue
			}
			completeTask(num - 1)

		case "delete":
			if len(parts) < 2 {
				fmt.Println("❗ Please provide task number.")
				continue
			}
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("❗ Invalid number.")
				continue
			}
			deleteTask(num - 1)

		case "exit":
			fmt.Println("👋 Goodbye! Take care of your mental health 💚")
			return

		default:
			fmt.Println("❓ Unknown command. Try: add, list, complete, delete, exit")
		}
	}
}
