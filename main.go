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

Usage: tasker <command> [arguments]

Commands:
  add         Add a new task
              Example: tasker add "Complete project report"

  list        Display all tasks
              Example: tasker list

  edit-work   Update task title
              Example: tasker edit-work 1 "Updated task"

  change-branch  Modify task's branch
                 Example: tasker change-branch 1 "Development"

  assign      Assign branch to task
              Example: tasker assign 3 "Backend"

  toggle      Change task completion status
              Example: tasker toggle 1

  del         Delete a task
              Example: tasker del 2

Global Options:
  -h, -help ,help   Show this help message

For more information, visit:
https://github.com/Sidontheway/Tasker_CLI
`, logo)

	fmt.Println(strings.TrimSpace(helpText))
}
