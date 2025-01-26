package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type CmdFlags struct {
// 	Add          string
// 	Del          int
// 	EditWork     string
// 	ChangeBranch string
// 	AssignBranch string
// 	Toggle       int
// 	List         bool
// }

// func NewCmdFlags() *CmdFlags {
// 	cf := CmdFlags{}

// 	flag.StringVar(&cf.Add, "add", "", "Add a new work specify title")
// 	flag.StringVar(&cf.EditWork, "edit-work", "", "Edit a work by index & specify a new title. id:new_title")
// 	flag.StringVar(&cf.ChangeBranch, "change-branch", "", "Change a assigned branch by index & specify a new title. id:new_title")
// 	flag.StringVar(&cf.AssignBranch, "assign", "", "Assign the branch of the work by its id id:branch_name")
// 	flag.IntVar(&cf.Del, "del", -1, "Specify work by index to delete")
// 	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify work by index to toggle complete true/false")
// 	flag.BoolVar(&cf.List, "list", false, "List all works")

// 	flag.Parse()

// 	return &cf
// }

func Execute(works *Works) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("no command provided")
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			return fmt.Errorf("add command requires a task description")
		}
		works.add(strings.Join(os.Args[2:], " "))

	case "list":
		works.print()

	case "edit-work":
		if len(os.Args) < 4 {
			return fmt.Errorf("edit-work requires index and new title")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid index")
		}
		works.editWork(index, strings.Join(os.Args[3:], " "))

	case "change-branch":
		if len(os.Args) < 4 {
			return fmt.Errorf("change-branch requires index and new branch")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid index")
		}
		works.editBranch(index, os.Args[3])

	case "assign":
		if len(os.Args) < 4 {
			return fmt.Errorf("assign requires index and branch name")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid index")
		}
		works.assignWork(index, os.Args[3])

	case "toggle":
		if len(os.Args) < 3 {
			return fmt.Errorf("toggle requires task index")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid index")
		}
		works.toggle(index)

	case "del":
		if len(os.Args) < 3 {
			return fmt.Errorf("del requires task index")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("invalid index")
		}
		works.delete(index)

	case "help":
		displayHelpMessage()

	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}
