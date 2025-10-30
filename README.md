Go Task Manager

1.TITLE AND OBJECTIVE
PROJECT;GO task manager
TECHNOLOGY USED; Go(GOLANG)

END GOAL;
This is a CLI application that demonstrates basic Go concepts such as:

- 🧩 Structs and slices  
- ⌨️ User input handling  
- 🔁 Control flow  
- 🧮 Basic CRUD logic  

The app allows users to:
- ✅ Add tasks with descriptions  
- 📋 View all tasks with completion status (✅ or ❌)  
- 📝 Mark tasks as complete  
- 🗑️ Delete tasks  
- 🚪 Exit the application gracefully  

2.QUICK SUMMARY OF THE TECHNOLOGY
WHAT IS GO? Go is an open-source programming language created by Google in 2009.It’s designed to make it simple, fast, and safe to build programs that run efficiently on modern computers and networks.

WHY USE GO
-Easy to read
-Catches many bugs before running
-It’s fast and memory efficient
-It has a powerful standard library;no need for many external packages

WHERE IT USED 
-Back-end development
-Cloud computing
-DevOps
-System programming

Real world examples; Big names use Go eg Google uses Go for backend systems and Uber for geolocation services

SYSTEM REQUIREMENTS
-OS: Windows, macOS, or Linus
-Memory: At least 2 GB RAM (4 GB+ recommended)
-Storage: ~1 GB free space
-Processor: Any modern CPU

You’ll also need:
A code editor e.g., Visual Studio Code
The Go compiler .....download from https://go.dev/dl/

4. INSTALLATION AND SETUP INSRTUCTIONS

STEP 1: INSTALL GO
## ⚙️ Installation & Setup Instructions

Follow these steps to install and run this Go project.

### 1. Install Go
If you don’t already have Go installed, download it from the official website:

 [https://go.dev/dl/](https://go.dev/dl/)

#After installation, verify it by running:

```bash
go version
 #You should see something like:
 go version go1.23.2 windows/amd64
2. Set up project workspace
#Create a New Go Project
mkdir go-toolkit
cd go-toolkit
#Initialize a Go Module
go mod init go-toolkit
You will see a file named;go.mod

5.MINIMAL WORKING SPACE
What the example does; 
It is a CLI application that demonstrates basic Go concepts like structs, slices, user input handling, and control flow. Basically it;
-Adds tasks with descriptions
-Lists all tasks with completion status (✅ for completed, ❌ for pending)
-Marks tasks as complete
-Deletes tasks from the list
-Exits the application
## 💻 Code Example

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
EXPECTED OUTPUT 
### 🧾 Expected Output

When you run the program using:

```bash
go run main.go
  You should see something like this
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


6.AI PROMPT JOURNAL

PROMPT 1;
I want to understand GO language as a a beginner in programming, Could you break down:How GO is implemented - The key syntax and structures I need to understand  - Common patterns and best practices
AI response summary;Go's philosophy is "simplicity and readability." The language is designed to be easy to learn while being powerful enough for production systems. Start with small programs, gradually incorporate more advanced features, and you'll quickly become productive!

PROMPT 2;
I want to understand GO language as a a beginner in programming, Could you break down:
-How GO is implemented - The key syntax and structures I need to understand  - Common patterns and best practices

AI response summary;Start with the basic structure and gradually add features. The key is to build small, testable components and integrate them together
PROMPT 3;
Why won’t VS Code let me edit?
AI response summary;Learned to reopen the folder, not the file directly.
PROMPT 4Create a simple Go project that adds, completes, and deletes tasks.
AI response;Got the base code for my CLI project.


7.COMMON ISSUES AND FIXESssue 1: "go: command not found"
Problem: Go not properly installed or not in PATH Solution:

-Reinstall Go 
-Add Go bin directory to system PATH
-Restart terminal/command prompt
Problem;cannot edit in read-only editor` 
 Opened file outside project folder - Reopened folder in VS Code using **File → Open Folder**
PROBLEM;cannot edit in read-only editor
-Opened file, not folder
-Used File → Open Folder in VS

8.PROJECT STRUCTURE
## 📂 Project Structure
go-task-manager/
├── go.mod # Defines the module path and Go version
├── main.go # Main Go source file containing all the logic
└── README.md # Project documentation and setup instructions
EXPLANATION

### 🧭 Explanation

- **main.go** → Contains your Go code (task manager logic).  
- **go.mod** → Keeps track of your module name and dependencies.  
- **README.md** → Explains how to use and run your project.  
- **/bin** → Optional folder where you can place built executables after running `go build`.

9.HOW TO RUN
## 🚀 How to Run

1. **Open the project folder** in VS Code (e.g., `go-task-manager`).
2. **Open the terminal** (`Ctrl + ~` or View → Terminal).
3. **Verify Go is installed** by running:
   ```bash
   go version
You should see something like go version go1.23.x windows/amd64.
4. Run the program:
go run main.go
5.Start interacting with the app by typing commands such as:
add Mental health matters
list
complete 1
delete 2
exit
Example Usage

After running go run main.go, you’ll see:
🧠 Welcome to Go Task Manager! Remember: Mental health matters 💚
Type one of the following commands:
add <task name> | complete <task number> | delete <task number> | list | exit
Example Session:
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


 OFFICIAL DOCUMENTATION
 -https://go.dev/doc/
 -http://www.cs.cmu.edu/afs/cs.cmu.edu/academic/class/15440-f11/go/doc/docs.html
 
 OFFICIAL TUTORIALS
 -A Tour of Go 
 -Effective Go 

COMMUNITY RESOURCES
-Go Wiki




 



