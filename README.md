# 🧠 Go Task Manager

---

## 📋 Project Overview

**Title:** Go Task Manager  
**Objective:** CLI Task Management Application  
**Technology Used:** Go (Golang)  

---

## 🎯 Project Goal

A command-line task manager that demonstrates fundamental Go programming concepts:

### 🔧 Core Concepts Demonstrated
- 🧩 **Structs and slices**
- ⌨️ **User input handling** 
- 🔁 **Control flow**
- 🧮 **Basic CRUD operations**

### ✨ Application Features
- ✅ **Add** tasks with descriptions
- 📋 **View** all tasks with completion status (✅ or ❌)
- 📝 **Mark** tasks as complete
- 🗑️ **Delete** tasks
- 🚪 **Exit** application gracefully

---

## 🚀 About Go (Golang)

### ℹ️ What is Go?
Go is an open-source programming language created by Google in 2009. It's designed to make it simple, fast, and safe to build efficient programs for modern computers and networks.

### 💡 Why Use Go?
- ✅ **Easy to read** and maintain
- ✅ **Catches many bugs** before running
- ✅ **Fast and memory efficient**
- ✅ **Powerful standard library** (minimal external dependencies)

### 🌍 Where is Go Used?
- 🔧 **Back-end development**
- ☁️ **Cloud computing**
- ⚙️ **DevOps tools**
- 💻 **System programming**

### 🏢 Real-World Examples
**Big companies using Go:**
- **Google** - Backend systems
- **Uber** - Geolocation services
- **Twitch** - Chat systems
- **Dropbox** - File synchronization

---

## 💻 System Requirements

### 📋 Minimum Requirements
- **OS:** Windows, macOS, or Linux
- **Memory:** 2 GB RAM (4 GB+ recommended)
- **Storage:** ~1 GB free space  
- **Processor:** Any modern CPU

### 🔧 Required Tools
- **Code Editor:** Visual Studio Code (or any preferred editor)
- **Go Compiler:** Download from [https://go.dev/dl/](https://go.dev/dl/)

---

4. INSTALLATION AND SETUP INSRTUCTIONS
markdown
---

## 🛠️ Installation Guide

### Step 1: Install Go
If you don't already have Go installed, download it from the official website:  
**https://go.dev/dl/**

After installation, verify it by running:
```bash
go version
You should see something like:

text
go version go1.23.2 windows/amd64
Step 2: Set Up Project Workspace
bash
# Create a new Go project
mkdir go-toolkit
cd go-toolkit

# Initialize a Go module
go mod init go-toolkit
You will see a file named go.mod created in your directory.

🎯 Minimal Working Example
What This Example Does:
This CLI application demonstrates basic Go concepts like structs, slices, user input handling, and control flow. It allows you to:

✅ Add tasks with descriptions

📋 List all tasks with completion status (✅ for completed, ❌ for pending)

📝 Mark tasks as complete

🗑️ Delete tasks from the list

🚪 Exit the application

💻 Code Example
go
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
🧾 Expected Output
When you run the program using:

bash
go run main.go
You should see something like this:

text
🧠 Welcome to Go Task Manager!
Type one of the following commands:
add <task name> | complete <task number> | delete <task number> | list | exit

> add Mental health matters
✨ Task added successfully!

> add Take a 5-minute break
✨ Task added successfully!

> list

Your Tasks:
1. ❌ Mental health matters
2. ❌ Take a 5-minute break

> complete 1
🎉 Task marked as complete!

> list

Your Tasks:
1. ✅ Mental health matters
2. ❌ Take a 5-minute break

> delete 2
🗑️  Task deleted!

> exit
👋 Goodbye! Take care of your mental health 💚
📝 AI Prompt Journal
Prompt 1:
"I want to understand GO language as a beginner in programming, Could you break down: How GO is implemented - The key syntax and structures I need to understand - Common patterns and best practices"

AI Response Summary: Go's philosophy is "simplicity and readability." The language is designed to be easy to learn while being powerful enough for production systems. Start with small programs, gradually incorporate more advanced features, and you'll quickly become productive!

Prompt 2:
"I want to understand GO language as a beginner in programming, Could you break down: -How GO is implemented - The key syntax and structures I need to understand - Common patterns and best practices"

AI Response Summary: Start with the basic structure and gradually add features. The key is to build small, testable components and integrate them together.

Prompt 3:
"Why won't VS Code let me edit?"

AI Response Summary: Learned to reopen the folder, not the file directly.

Prompt 4:
"Create a simple Go project that adds, completes, and deletes tasks."

AI Response: Got the base code for my CLI project.

🔧 Common Issues and Fixes
Issue 1: "go: command not found"
Problem: Go not properly installed or not in PATH
Solution:

Reinstall Go

Add Go bin directory to system PATH

Restart terminal/command prompt

Issue 2: "cannot edit in read-only editor"
Problem: Opened file outside project folder
Solution: Reopened folder in VS Code using File → Open Folder

Issue 3: "cannot edit in read-only editor"
Problem: Opened file, not folder
Solution: Used File → Open Folder in VS Code

📂 Project Structure
text
go-task-manager/
├── go.mod          # Defines the module path and Go version
├── main.go         # Main Go source file containing all the logic
└── README.md       # Project documentation and setup instructions
🧭 Explanation
main.go → Contains your Go code (task manager logic)

go.mod → Keeps track of your module name and dependencies

README.md → Explains how to use and run your project

/bin → Optional folder where you can place built executables after running go build

🚀 How to Run
Open the project folder in VS Code (e.g., go-task-manager)

Open the terminal (Ctrl + ~ or View → Terminal)

Verify Go is installed by running:

bash
go version
You should see something like: go version go1.23.x windows/amd64

Run the program:

bash
go run main.go
Start interacting with the app by typing commands such as:

text
add Mental health matters
list
complete 1
delete 2
exit
Example Usage
After running go run main.go, you'll see:

text
🧠 Welcome to Go Task Manager! Remember: Mental health matters 💚
Type one of the following commands:
add <task name> | complete <task number> | delete <task number> | list | exit
Example Session:

text
> add Mental health matters
✨ Task added successfully!

> add Drink some water
✨ Task added successfully!

> list
Your Tasks:
1. ❌ Mental health matters
2. ❌ Drink some water

> complete 1
🎉 Task marked as complete!

> delete 2
🗑️  Task deleted!

> exit
👋 Goodbye! Take care of your mental health 💚
📚 Learning Resources
Official Documentation
https://go.dev/doc/

http://www.cs.cmu.edu/afs/cs.cmu.edu/academic/class/15440-f11/go/doc/docs.html

Official Tutorials
A Tour of Go

Effective Go

Community Resources
Go Wiki


    
    







        
 









 



