package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add          string
	Del          int
	EditWork     string
	ChangeBranch string
	AssignBranch string
	Toggle       int
	List         bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new work specify title")
	flag.StringVar(&cf.EditWork, "edit-work", "", "Edit a work by index & specify a new title. id:new_title")
	flag.StringVar(&cf.ChangeBranch, "change-branch", "", "Change a assigned branch by index & specify a new title. id:new_title")
	flag.StringVar(&cf.AssignBranch, "assign", "", "Assign the branch of the work by its id id:branch_name")
	flag.IntVar(&cf.Del, "del", -1, "Specify work by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify work by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all works")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(works *Works) {
	switch {
	case cf.List:
		works.print()
	case cf.Add != "":
		works.add(cf.Add)
	case cf.EditWork != "":
		parts := strings.SplitN(cf.EditWork, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for edit . Please id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error :invalid index for edit")
			os.Exit(1)
		}

		works.editWork(index, parts[1])

	case cf.ChangeBranch != "":
		parts := strings.SplitN(cf.ChangeBranch, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for edit , Please choose id:new_branch_name")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error:invalid index for edit")
			os.Exit(1)
		}

		works.editBranch(index, parts[1])

	case cf.AssignBranch != "":
		parts := strings.SplitN(cf.AssignBranch, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for assigning branch . Please id:assigned_branch_name")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error :invalid index for edit")
			os.Exit(1)
		}

		works.assignWork(index, parts[1])

	case cf.Toggle != -1:
		works.toggle(cf.Toggle)

	case cf.Del != -1:
		works.delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}

}
