package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	works := Works{}
	storage := NewStorage[Works]("works.json")
	storage.Load(&works)

	if len(os.Args) < 2 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "-help")) {
		displayHelpMessage()
		os.Exit(0)
	}

	if err := Execute(&works); err != nil {
		fmt.Println("Error:", err)
		displayHelpMessage()
		os.Exit(1)
	}

	// cmdFlags.Execute(&works)

	storage.Save(works)
}

func displayHelpMessage() {
	logo := `
 _____        _             
|_   _|      | |            
  | | ___ ___| | _____ _ __ 
  | |/ __/ __| |/ / _ \ '__|
 _| |\__ \__ \   <  __/ |   
 \___/___/___/_|\_\___|_|   
`

	helpText := fmt.Sprintf(`
%s
Tasker CLI - Efficient Task Management
=====================================================

Commands:
  add         Add a new task
               Example: tasker -add "Complete project report"

  edit-work   Edit a task's title
               Format: tasker -edit-work "id:new title"
               Example: tasker -edit-work "2:Updated task description"

  change-branch Modify task's branch assignment
                 Format: tasker -change-branch "id:new branch"
                 Example: tasker -change-branch "1:Development"

  assign      Assign a branch to a specific task
               Format: tasker -assign "id:branch_name"
               Example: tasker -assign "3:Backend"

  del         Delete a task by index
               Example: tasker -del 2

  toggle      Toggle task completion status
               Example: tasker -toggle 1

  list        Display all tasks
               Example: tasker -list

Global Options:
  -h, -help    Show this help message

For more information, visit:
https://github.com/Sidontheway/Tasker_CLI
`, logo)

	fmt.Println(strings.TrimSpace(helpText))
}
