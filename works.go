package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

type Work struct {
	Title          string
	Completed      bool
	AssignedBranch string
}

type Works []Work

func (works *Works) add(title string) {
	work := Work{
		Title:          title,
		Completed:      false,
		AssignedBranch: "",
	}

	*works = append(*works, work)
}

func (works *Works) validateIndex(index int) error {
	if index < 0 || index >= len(*works) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (works *Works) isWorkAlreadyAssigned(index int) error {
	if err := works.validateIndex(index); err != nil {
		return err
	} else if (*works)[index].AssignedBranch != "" {
		err := errors.New("work is already assigned to another branch")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (works *Works) assignWork(index int, branch string) error {

	if err := works.isWorkAlreadyAssigned(index); err != nil {
		return err
	}

	(*works)[index].AssignedBranch = branch

	return nil
}

func (works *Works) delete(index int) error {
	t := *works
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*works = append(t[:index], t[index+1:]...)

	return nil
}

func (works *Works) toggle(index int) error {
	if err := works.validateIndex(index); err != nil {
		return err
	}

	isCompleted := (*works)[index].Completed

	(*works)[index].Completed = !isCompleted
	return nil
}

func (works *Works) editWork(index int, title string) error {
	if err := (*works).validateIndex(index); err != nil {
		return err
	}

	(*works)[index].Title = title
	return nil
}

func (works *Works) editBranch(index int, branch string) error {

	if err := works.validateIndex(index); err != nil {
		return err
	}

	(*works)[index].AssignedBranch = branch
	return nil
}

func (works *Works) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	// table.SetHeaders("#", "Works", "Completed Status", "Assigned Branch")

	table.SetHeaders(
		tml.Sprintf("<bg-black><bold>#</bold></bg-black>"),
		tml.Sprintf("<bg-black><bold>Works</bold></bg-black>"),
		tml.Sprintf("<bg-black><bold>Completed Status</bold></bg-black>"),
		tml.Sprintf("<bg-black><bold>Assigned Branch</bold></bg-black>"))

	for index, t := range *works {
		completed := "❌"

		if t.Completed {
			completed = "✅"
		}

		table.AddRow(
			tml.Sprintf("<italic>%s</italic>", strconv.Itoa(index)),
			tml.Sprintf("<italic>%s</italic>", t.Title),
			completed,
			tml.Sprintf("<italic>%s</italic>", t.AssignedBranch),
		)
	}
	table.Render()
}
