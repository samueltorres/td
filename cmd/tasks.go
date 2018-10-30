package cmd

import (
	"strings"

	"github.com/samueltorres/td/pkg/printer"

	"github.com/spf13/cobra"
)

var verbose bool

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		cf, bs := LoadBoard()
		tks, err := bs.GetTasks(cf.CurrentBoard)
		CheckError(err)

		printer.PrintTasks(tks, verbose)
	},
}

var addTaskCommand = &cobra.Command{
	Use:   "add [task_description]",
	Short: "Adds a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cf, bs := LoadBoard()
		err := bs.AddTask(cf.CurrentBoard, strings.Join(args[0:], " "))
		CheckError(err)
	},
}

var taskDoneCommand = &cobra.Command{
	Use:   "done [task_id]",
	Short: "Sets a task to done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, bs := LoadBoard()
		err := bs.SetTaskStatus(c.CurrentBoard, args[0], true)
		CheckError(err)
	},
}

var taskUndoneCommand = &cobra.Command{
	Use:   "undone [task_id]",
	Short: "Sets a task to undone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, bs := LoadBoard()
		err := bs.SetTaskStatus(c.CurrentBoard, args[0], false)
		CheckError(err)
	},
}

var removeTaskCommand = &cobra.Command{
	Use:   "remove [task_id]",
	Short: "Removes a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, bs := LoadBoard()
		err := bs.RemoveTask(c.CurrentBoard, args[0])
		CheckError(err)
	},
}

func init() {
	listCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "lists tasks with verbose mode")

	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(addTaskCommand)
	rootCmd.AddCommand(taskDoneCommand)
	rootCmd.AddCommand(taskUndoneCommand)
	rootCmd.AddCommand(removeTaskCommand)
}
