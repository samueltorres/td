package printer

import (
	"fmt"

	"github.com/samueltorres/td/pkg/boards"
)

// PrintBoards renders the list of boards
func PrintBoards(bs []string) {
	for _, b := range bs {
		fmt.Println(b)
	}
}

// PrintTasks prints the tasks to the console
func PrintTasks(tks []*boards.Task, verbose bool) {
	for _, t := range tks {
		st := "[ ]"
		if t.Status {
			st = "[x]"
		}

		if verbose {
			fmt.Println(fmt.Sprintf("%s  %s  (%s)", st, t.Description, t.ID))
		} else {
			fmt.Println(fmt.Sprintf("%s  %s", st, t.Description))
		}

	}
}

// PrintMessage prints a message to the console
func PrintMessage(m string) {
	fmt.Println(m)
}
