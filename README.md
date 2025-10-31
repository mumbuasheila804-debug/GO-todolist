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

## ⚙️ Step 1: Install and Verify Go

### 🧩 Download Go
If you don’t already have Go installed, download it from the official page:  
🔗 [https://go.dev/doc/install](https://go.dev/doc/install)


Verify installation:

```bash
go version
```

**Expected output:**
```
go version go1.23.2 windows/amd64
```

 

### 📁 Step 2: Set Up Project Workspace

```bash
# Create project directory
mkdir go-task-manager
cd go-task-manager

# Initialize Go module (replace with your module path if publishing)
go mod init github.com/yourusername/go-task-manager
```

---
##🎯 Minimal Working Example
What This Example Does:
This CLI application demonstrates basic Go concepts like structs, slices, user input handling, and control flow. It allows you to:

✅ Add tasks with descriptions

📋 List all tasks with completion status (✅ for completed, ❌ for pending)

📝 Mark tasks as complete

🗑️ Delete tasks from the list

🚪 Exit the application
 
 ---

 
### 💻 Step 3: Add Your Code

Create `main.go` and paste:

```go
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
    fmt.Println("🧠 Welcome to Go Task Manager! Remember: Mental health matters 💚")
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
```


### 🧾 Expected Output

Run the program:

```bash
go run main.go
```

You should see something like:

```
🧠 Welcome to Go Task Manager! Remember: Mental health matters 💚
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
```

 ## 📝 AI Prompt Journal  

---

### 💡 **Prompt 1**
> 🗣️ *"I want to understand Go language as a beginner in programming. Could you break down:*  
> • *How Go is implemented*  
> • *The key syntax and structures I need to understand*  
> • *Common patterns and best practices"*

**🧠 AI Response Summary:**  
Go’s philosophy focuses on **simplicity and readability**.  
Start with the **basic structure** and gradually add new features.  
The key is to **build small, testable components** and integrate them step by step.

---

### 🧩 **Prompt 2**
> 🗣️ *"Why won’t VS Code let me edit?"*

**🧠 AI Response Summary:**  
Discovered that the issue happened because the **file was opened directly** instead of the **project folder**.  
Opening the entire folder in VS Code fixed the **read-only** editing issue.

---

### ⚙️ **Prompt 3**
> 🗣️ *"Create a simple Go project that adds, completes, and deletes tasks."*

**🧠 AI Response Summary:**  
Received a **Go CLI project** template with essential features —  
✅ *Add tasks*  
✅ *Complete tasks*  
✅ *List tasks*  
✅ *Delete tasks*  

It provided a strong starting point for building and learning core Go concepts.

---

✨ *End of Journal — documenting my progress as I learn Go and build real projects!*  

---

### 🔧 Common Issues and Fixes

**Issue 1:** `go: command not found`  
💡 **Fix:** Reinstall Go and ensure its bin folder is added to your system PATH.

**Issue 2:** “cannot edit in read-only editor”  
💡 **Fix:** Open your project folder in VS Code using **File → Open Folder**.

---

### 📂 Project Structure

```text
go-task-manager/
├── go.mod          # Defines module path and Go version
├── main.go         # Contains all application logic
└── README.md       # Project documentation
```

---
 

### 🚀 How to Run

1. Open your project folder in **VS Code**  
2. Open the terminal (**Ctrl + `**)  
3. Run the app:

```bash
go run main.go
```

4. Start typing commands such as:

```bash
add Mental health matters
list
complete 1
delete 2
exit
```

---
Example Usage
### 💬 Example Session

```
🧠 Welcome to Go Task Manager! Remember: Mental health matters 💚
Type one of the following commands:
add <task name> | complete <task number> | delete <task number> | list | exit

> add Drink water
✨ Task added successfully!
> list
1. ❌ Drink water
> complete 1
🎉 Task marked as complete!
> exit
👋 Goodbye! Take care of your mental health 💚
```

---

### 📚 Learning Resources

**Official Documentation**
- [Go Docs](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)

**Tutorials**
- [A Tour of Go](https://go.dev/tour/)

**Community**
- [Go Wiki](https://github.com/golang/go/wiki)



    
    







        
 









 



**
